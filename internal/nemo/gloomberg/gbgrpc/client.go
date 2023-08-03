package gbgrpc

import (
	context "context"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg/gbgrpc/gen"
	seawaModels "github.com/benleb/gloomberg/internal/seawa/models"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

var GRPCClient gen.GloombergClient

// type grpcClient struct {
// 	gen.GloombergClient
// }

// NewClient to grpc server.
func NewClient(grpcAddress string) gen.GloombergClient {
	if grpcAddress == "" {
		grpcAddress = fmt.Sprintf("%s:%d", viper.GetString("grpc.client.host"), viper.GetUint("grpc.client.port"))
	}

	// grpc options
	var opts []grpc.DialOption

	if creds := gloomberg.GetTLSClientCredentials(); creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// connect
	gloomberg.Prf("connecting to gRPC %s...", style.BoldAlmostWhite(grpcAddress))
	conn, err := grpc.Dial(grpcAddress, opts...)
	if err != nil {
		gbl.Log.Warnf("fail to dial: %v", err)

		return nil
	}

	// time needed to establish connection
	time.Sleep(time.Millisecond * 337)

	// check connection state
	if conn.GetState() != connectivity.Ready {
		return nil
	}

	// return client
	GRPCClient = gen.NewGloombergClient(conn)

	return GRPCClient // gen.NewGloombergClient(conn)
}

func FetchEvents(gb *gloomberg.Gloomberg) {
	// server to connect to
	grpcAddress := fmt.Sprintf("%s:%d", viper.GetString("grpc.client.host"), viper.GetUint("grpc.client.port"))

	// grpc options
	var opts []grpc.DialOption

	if creds := gloomberg.GetTLSClientCredentials(); creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	failedAttempts := 0

	for {
		if failedAttempts > 0 {
			// exponential backoff
			waitTime := time.Second * time.Duration(math.Pow(2.0, float64(int(math.Min(float64(failedAttempts), 5)))))

			log.Printf("retrying to connect to gRPC %s in %.0f seconds...", style.BoldAlmostWhite(grpcAddress), waitTime.Seconds())

			time.Sleep(waitTime)
		}

		// connect
		gloomberg.Prf("connecting to gRPC %s...", style.BoldAlmostWhite(grpcAddress))
		conn, err := grpc.Dial(grpcAddress, opts...)
		if err != nil {
			gbl.Log.Warnf("fail to dial: %v", err)

			continue
		}

		// time needed to establish connection
		time.Sleep(time.Millisecond * 337)

		// check connection state
		if conn.GetState() != connectivity.Ready {
			log.Errorf("fail to connect to gRPC %s", style.BoldAlmostWhite(grpcAddress))
		}

		// return client
		grpcClient := gen.NewGloombergClient(conn)

		if grpcClient == nil {
			log.Errorf("fail to connect to gRPC %s", style.BoldAlmostWhite(grpcAddress))

			failedAttempts++

			continue
		}

		// get eventstream
		gloomberg.Prf("subscribing via grpc to: %s", style.BoldAlmostWhite(degendb.Listing.OpenseaEventName()))

		subsriptionRequest := &gen.SubscriptionRequest{EventTypes: []gen.EventType{gen.EventType_ITEM_LISTED}, Collections: gb.CollectionDB.OpenseaSlugs()} //nolint:nosnakecase
		stream, err := grpcClient.GetEvents(context.Background(), subsriptionRequest)
		if err != nil {
			log.Errorf("getting stream failed: %v, retrying", err)

			failedAttempts++

			continue
		}

		// reset failed attempts
		failedAttempts = 0

		for {
			// get events
			event, err := stream.Recv()
			log.Debugf("🐔 client received: %+v", event)

			if err != nil {
				if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) || errors.Is(err, io.ErrClosedPipe) {
					log.Errorf("io.EOF error: %v", err)
				}

				if event == nil {
					log.Errorf("receiving event failed: %v | %+v", err, event)
				}

				err := stream.CloseSend()
				if err != nil {
					log.Errorf("closing stream failed: %v", err)
				}

				break
			}

			basePrice, ok := new(big.Int).SetString(event.Payload.GetItemListed().Payload.BasePrice, 10)
			if !ok {
				log.Errorf("error parsing base price: %v", err)

				continue
			}

			var itemListed seawaModels.ItemListed

			log.Debugf("🐔 creating itemListed struct: %+v", protojson.Format(event))

			// transform event back to seawaModel.ItemListed
			itemListed = seawaModels.ItemListed{
				EventType: strings.ToLower(event.EventType.String()),
				SentAt:    event.Payload.GetItemListed().SentAt.AsTime(),
				Payload: seawaModels.ItemListedPayload{
					Item: seawaModels.Item{
						NftID:     *seawaModels.ParseNftID(event.Payload.GetItemListed().Payload.Item.NftId),
						Chain:     seawaModels.Chain{Name: event.Payload.GetItemListed().Payload.Item.Chain.Name},
						Permalink: event.Payload.GetItemListed().Payload.Item.Permalink,
						Metadata: seawaModels.Metadata{
							Name:         event.Payload.GetItemListed().Payload.Item.Metadata.Name,
							ImageURL:     event.Payload.GetItemListed().Payload.Item.Metadata.ImageUrl,
							AnimationURL: event.Payload.GetItemListed().Payload.Item.Metadata.AnimationUrl,
							MetadataURL:  event.Payload.GetItemListed().Payload.Item.Metadata.MetadataUrl,
						},
					},
					IsPrivate:   event.Payload.GetItemListed().Payload.IsPrivate,
					ListingDate: event.Payload.GetItemListed().Payload.ListingDate.AsTime(),
					EventPayload: seawaModels.EventPayload{
						EventTimestamp:     event.Payload.GetItemListed().Payload.EventTimestamp.AsTime(),
						BasePrice:          basePrice,
						Maker:              seawaModels.Account{Address: common.HexToAddress(event.Payload.GetItemListed().Payload.Maker.Address)},
						Taker:              seawaModels.Account{Address: common.HexToAddress(event.Payload.GetItemListed().Payload.Taker.Address)},
						Quantity:           int(event.Payload.GetItemListed().Payload.Quantity),
						OrderHash:          common.HexToHash(event.Payload.GetItemListed().Payload.OrderHash),
						ExpirationDate:     event.Payload.GetItemListed().Payload.ExpirationDate.AsTime(),
						CollectionCriteria: seawaModels.CollectionCriteria{Slug: event.Payload.GetItemListed().Payload.Collection.Slug},
						PaymentToken:       seawaModels.PaymentToken{Address: common.HexToAddress(event.Payload.GetItemListed().Payload.PaymentToken.Address), Symbol: event.Payload.GetItemListed().Payload.PaymentToken.Symbol, Decimals: int(event.Payload.GetItemListed().Payload.PaymentToken.Decimals)},
					},
				},
			}

			// send event to the eventhub
			gb.In.ItemListed <- &itemListed

			log.Debugf("🐔 sent to eventHub: %+v", itemListed)
		}
	}
}
