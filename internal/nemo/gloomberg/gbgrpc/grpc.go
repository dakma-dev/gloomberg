package gbgrpc

import (
	context "context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg/gbgrpc/gen"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/seawa"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GloombergGRPC struct {
	gb *gloomberg.Gloomberg
	sw *seawa.SeaWatcher

	gen.UnimplementedGloombergServer

	// grpcServer *grpc.Server
}

// connect to grpc server.
func listen(serverAddress string, gloombergGRPC *GloombergGRPC) error {
	// configure listener
	grpcListener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}

	// configure server
	var opts []grpc.ServerOption
	if creds, err := gloomberg.GetTLSCredentialsWithoutClientAuth(); err == nil {
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	// run grpc server
	grpcServer := grpc.NewServer(opts...)

	gen.RegisterGloombergServer(grpcServer, gloombergGRPC)

	return grpcServer.Serve(grpcListener)
}

func StartServer(gb *gloomberg.Gloomberg, sw *seawa.SeaWatcher) {
	gloombergGRPC := &GloombergGRPC{
		gb: gb,
		sw: sw,
	}

	// get config
	listenHost := viper.GetString("grpc.server.listenAddress")
	port := viper.GetUint16("grpc.server.port")
	serverAddress := fmt.Sprintf("%s:%d", listenHost, port)

	go func() {
		for {
			log.Error(listen(serverAddress, gloombergGRPC))

			time.Sleep(time.Second * 2)
		}
	}()
	// // configure listener
	// grpcListener, err := net.Listen("tcp", serverAddress)
	// if err != nil {
	// 	log.Errorf("failed to listen: %v", err)
	// }

	// // configure server
	// var opts []grpc.ServerOption
	// if creds, err := gloomberg.GetTLSCredentialsWithoutClientAuth(); err == nil {
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }

	// // run grpc server
	// go func() {
	// 	grpcServer := grpc.NewServer(opts...)

	// 	gen.RegisterGloombergServer(grpcServer, gloombergGRPC)

	// 	go log.Fatal(grpcServer.Serve(grpcListener))
	// }()

	sw.Prf("grpc server running on %+v", style.BoldAlmostWhite(serverAddress))
}

func (gg *GloombergGRPC) Subscribe(_ context.Context, req *gen.SubscriptionRequest) (*emptypb.Empty, error) {
	availableEventTypes := []gen.EventType{gen.EventType_ITEM_LISTED, gen.EventType_METADATA_UPDATED, gen.EventType_ITEM_RECEIVED_BID, gen.EventType_COLLECTION_OFFER} //nolint:nosnakecase // ItemMetadataUpdated} // ItemMetadataUpdated, ItemCancelled

	req.EventTypes = availableEventTypes

	gg.sw.Prf("received subscription request for %s collections/slugs (%s types each)...", style.BoldAlmostWhite(fmt.Sprint(len(req.Collections))), style.BoldAlmostWhite(fmt.Sprint(len(req.EventTypes))))

	newEventSubscriptions := uint64(0)

	go func() {
		newEventSubscriptions = gg.sw.SubscribeForSlugs(req.Collections, req.EventTypes)

		gg.sw.Prf(
			"successfully subscribed to %s new collections/slugs | total subscribed collections: %s",
			style.AlmostWhiteStyle.Render(fmt.Sprint(newEventSubscriptions)),
			style.AlmostWhiteStyle.Render(fmt.Sprint(len(gg.sw.ActiveSubscriptions()))),
		)
	}()

	return &emptypb.Empty{}, nil
}

func (gg *GloombergGRPC) GetEvents(req *gen.SubscriptionRequest, stream gen.Gloomberg_GetEventsServer) error { //nolint:nosnakecase
	// availableEventTypes := []gen.EventType{gen.EventType_ITEM_LISTED, gen.EventType_METADATA_UPDATED, gen.EventType_ITEM_RECEIVED_BID, gen.EventType_COLLECTION_OFFER} //nolint:nosnakecase // ItemMetadataUpdated} // ItemMetadataUpdated, ItemCancelled

	// req.EventTypes = availableEventTypes

	// gg.sw.Prf("received subscription request for %s collections/slugs (%s types each)...", style.BoldAlmostWhite(fmt.Sprint(len(req.Collections))), style.BoldAlmostWhite(fmt.Sprint(len(req.EventTypes))))

	// go func() {
	// 	newEventSubscriptions := gg.sw.SubscribeForSlugs(req.Collections, req.EventTypes)

	// 	gg.sw.Prf(
	// 		"successfully subscribed to %s new collections/slugs | total subscribed collections: %s",
	// 		style.AlmostWhiteStyle.Render(fmt.Sprint(newEventSubscriptions)),
	// 		style.AlmostWhiteStyle.Render(fmt.Sprint(len(gg.sw.ActiveSubscriptions()))),
	// 	)
	// }()

	go gg.Subscribe(context.Background(), req) //nolint:errcheck

	chanItemListed := gg.gb.SubscribeItemListed()
	defer gg.gb.UnsubscribeItemListed(chanItemListed)

	for event := range chanItemListed {
		itemName := event.Payload.Item.Name

		// fix missing collection name in item name (example: INS1D3RS)
		if strings.HasPrefix(event.Payload.Item.Name, "#") {
			if collectionName, _ := gg.gb.Rueidi.GetCachedContractName(context.Background(), event.Payload.Item.ContractAddress()); collectionName != "" {
				itemName = fmt.Sprintf("%s %s%s", collectionName, "#", event.Payload.Item.TokenID())
			}
		}

		// transform *models.ItemListed event to ItemListed grpc message
		itemListed := &gen.ItemListed{
			EventType: gen.EventType(gen.EventType_value[event.EventType]), //nolint:nosnakecase
			SentAt:    &timestamppb.Timestamp{Seconds: event.SentAt.Unix()},

			Payload: &gen.ItemListed_ItemListedPayload{ //nolint:nosnakecase
				Item: &gen.ItemListed_Item{ //nolint:nosnakecase
					Chain:     &gen.ItemListed_Chain{Name: "ethereum"}, //nolint:nosnakecase
					NftId:     event.Payload.Item.String(),
					Permalink: event.Payload.Item.Permalink,
					Metadata: &gen.ItemListed_Metadata{ //nolint:nosnakecase
						Name:         itemName,
						ImageUrl:     event.Payload.Item.ImageURL,
						AnimationUrl: event.Payload.Item.AnimationURL,
						MetadataUrl:  event.Payload.Item.MetadataURL,
					},
				},
				BasePrice:      event.Payload.BasePrice.String(),
				Collection:     &gen.ItemListed_Collection{Slug: event.Payload.Slug}, //nolint:nosnakecase
				IsPrivate:      event.Payload.IsPrivate,
				ListingDate:    &timestamppb.Timestamp{Seconds: event.Payload.ListingDate.Unix()},
				EventTimestamp: &timestamppb.Timestamp{Seconds: event.Payload.EventTimestamp.Unix()},
				Quantity:       uint32(event.Payload.Quantity),
				Maker:          &gen.ItemListed_Account{Address: event.Payload.Maker.Address.String()}, //nolint:nosnakecase
				Taker:          &gen.ItemListed_Account{Address: event.Payload.Taker.Address.String()}, //nolint:nosnakecase
				ExpirationDate: &timestamppb.Timestamp{Seconds: event.Payload.ExpirationDate.Unix()},
				OrderHash:      event.Payload.OrderHash.String(),
				PaymentToken: &gen.ItemListed_PaymentToken{ //nolint:nosnakecase
					Address:  event.Payload.Address.String(),
					Symbol:   event.Payload.Symbol,
					Name:     event.Payload.Name,
					Decimals: uint32(event.Payload.Decimals),
					UsdPrice: event.Payload.UsdPrice,
				},
			},
		}

		ev := &gen.Event{
			EventType: gen.EventType_ITEM_LISTED, //nolint:nosnakecase
			Payload: &gen.EventPayload{
				Kind: &gen.EventPayload_ItemListed{ //nolint:nosnakecase
					ItemListed: itemListed,
				},
			},
		}

		// gg.sw.Prf("sending event: %+v", ev)

		if err := stream.Send(ev); err != nil {
			log.Printf("❌ error sending event to grpc client: %s", err)

			return err
		}

		// output to terminal
		collectionPrimaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(event.Payload.Item.NftID.ContractAddress().Hash().Big().Int64()))
		collectionSecondaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(event.Payload.Item.NftID.ContractAddress().Big().Int64() ^ 2))

		price := price.NewPrice(event.Payload.BasePrice)
		fmtCurrencySymbol := collectionPrimaryStyle.Bold(false).Render("Ξ")
		fmtPrice := style.BoldAlmostWhite(fmt.Sprintf("%5.2f", price.Ether())) + fmtCurrencySymbol

		fmtItemName := strings.ReplaceAll(collectionPrimaryStyle.Bold(true).Render(itemName), "#", collectionSecondaryStyle.Render("#"))
		fmtItemName = strings.ReplaceAll(fmtItemName, event.Payload.Item.TokenID().String(), collectionPrimaryStyle.Bold(true).Render(event.Payload.Item.TokenID().String()))

		fmtItemLink := style.TerminalLink(event.Payload.Item.Permalink, fmtItemName)

		eventType := degendb.EventType(degendb.GetEventType(event.EventType))

		gg.sw.Prf("%s %s %s", eventType.Icon(), fmtPrice, fmtItemLink)
	}

	return nil
}
