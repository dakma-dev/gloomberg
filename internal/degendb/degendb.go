package degendb

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoDB         = "dev-degendb"
	collAddresses   = "addresses"
	collCollections = "collections"
	collDegens      = "degens"
	collTokens      = "tokens"
)

type DegenDB struct {
	uri   string
	mongo *mongo.Client
}

func NewDegenDB() *DegenDB {
	ddb := &DegenDB{
		uri: viper.GetString("mongodb.uri"),
	}

	mongoClient, err := ddb.connect()
	if err != nil {
		log.Errorf("❌ could not connect to mongoDB at %s: %s", ddb.uri, err)

		return nil
	}

	log.Infof("✅ connected to mongoDB at %s", ddb.uri)

	ddb.mongo = mongoClient

	if viper.GetBool("mongodb.initialize") {
		// cleanup & initialize
		collectionsColl := ddb.mongo.Database(mongoDB).Collection(collCollections)
		degensColl := ddb.mongo.Database(mongoDB).Collection(collDegens)
		tokensColl := ddb.mongo.Database(mongoDB).Collection(collTokens)
		addressesColl := ddb.mongo.Database(mongoDB).Collection(collAddresses)

		err := collectionsColl.Drop(context.Background())
		if err != nil {
			log.Error(err)
		}

		err = degensColl.Drop(context.Background())
		if err != nil {
			log.Error(err)
		}

		err = tokensColl.Drop(context.Background())
		if err != nil {
			log.Error(err)
		}

		err = addressesColl.Drop(context.Background())
		if err != nil {
			log.Error(err)
		}

		// initialize og degens
		ddb.initializeDegensCollection()

		// // check
		// cursor, err := degensColl.Find(context.Background(), bson.D{})
		// if err != nil {
		// 	log.Error(err)
		// }

		// // query
		// var mongoDegen []Degen
		// if err = cursor.All(context.TODO(), &mongoDegen); err != nil {
		// 	log.Error(err)
		// }

		// // print
		// for _, dgn := range mongoDegen {
		// 	log.Printf("mongoDegen: %+v", dgn)
		// }
	}

	return ddb
}

func (ddb *DegenDB) connect() (*mongo.Client, error) {
	if ddb.uri == "" {
		log.Error("You must set your 'MONGODB_URI' environmental variable. See https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(ddb.uri))
	if err != nil {
		return nil, err
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		return nil, err
	}

	log.Print("mongodb ping successful")

	return client, nil
}

func (ddb *DegenDB) Disconnect() error {
	err := ddb.mongo.Disconnect(context.TODO())

	return err
}

func (ddb *DegenDB) NewAddresses(addresses []common.Address) []*Address {
	addrs := make([]*Address, 0)

	for _, addr := range addresses {
		addrs = append(addrs, ddb.NewAddress(addr))
	}

	return addrs
}

func (ddb *DegenDB) NewAddress(address common.Address) *Address {
	// codeAt, err := pool.GetCodeAt(context.Background(), address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// isContract := len(codeAt) > 0

	// addrType := "eoa"
	// if isContract {
	// 	addrType = "contract"
	// }

	// log.Printf("degendb| address %s is a %s address", style.AlmostWhiteStyle.Render(address.Hex()), style.AlmostWhiteStyle.Render(addrType))

	return &Address{HexAddress: address.Hex(), Address: address} // , IsContract: isContract, Type: addrType}
}

func (ddb *DegenDB) NewDegen(name string, addresses []common.Address, twitter string, telegram string, telegramID int64, tags []Tag) *Degen {
	addrs := ddb.NewAddresses(addresses)

	degen := &Degen{
		Name:      name,
		Addresses: addrs,
		Tags:      tags,
		Accounts:  Accounts{Twitter: twitter, Telegram: telegram, TelegramChatID: telegramID},
		CreatedAt: time.Now(),
	}

	return degen
}

// /.
func (ddb *DegenDB) initializeDegensCollection() {
	degensColl := ddb.mongo.Database(mongoDB).Collection(collDegens)
	addressesColl := ddb.mongo.Database(mongoDB).Collection(collAddresses)

	beAddr1 := common.HexToAddress("0x37416906c8011358DaB16F0d73BeEbf580d4AFa8")
	beAddr2 := common.HexToAddress("0x8fEE0de24CB0B8df2423dfC68113C133Cd7650b3")
	beAddr3 := common.HexToAddress("0xCd8aF79Ba3974404e37f126a8E355690351Da8bD")
	luAddr1 := common.HexToAddress("0x0DB54CC560Ae7832898e82E5E607E8142e519891")
	luAddr2 := common.HexToAddress("0x9654F22b9dEBac18396b4815C138A450786a7045")
	luAddr3 := common.HexToAddress("0xB364600e673E63FCbA4Ed1012F55DEb31eFa14ac")

	ogDegens := []interface{}{
		&Degen{
			Name:     "Ben",
			Accounts: Accounts{Twitter: "ben_leb", Telegram: "benleb", TelegramChatID: -1001808788625},
			Addresses: []*Address{
				{HexAddress: beAddr1.Hex(), Address: beAddr1, Name: "gnova"},
				{HexAddress: beAddr2.Hex(), Address: beAddr2, Name: "neva"},
				{HexAddress: beAddr3.Hex(), Address: beAddr3, Name: "drastic"},
			},
			Tags:      []Tag{"og", "dakma", "dev"},
			CreatedAt: time.Now(),
		},
		&Degen{
			Name:     "Luke",
			Accounts: Accounts{Twitter: "0xlugges", Telegram: "luke_take_profits"},
			Addresses: []*Address{
				{HexAddress: luAddr1.Hex(), Address: luAddr1},
				{HexAddress: luAddr2.Hex(), Address: luAddr2},
				{HexAddress: luAddr3.Hex(), Address: luAddr3},
			},
			Tags:      []Tag{"og", "dakma", "dev"},
			CreatedAt: time.Now(),
		},
	}

	for _, degen := range ogDegens {
		degen, ok := degen.(*Degen)
		if !ok {
			log.Printf("type asserting degen error")

			continue
		}
		// og := ogDegens[degenID].(Degen)

		addresses := make([]interface{}, 0)

		for _, address := range degen.Addresses {
			addresses = append(addresses, address)
		}

		addressResult, err := addressesColl.InsertMany(context.TODO(), addresses)
		if err != nil {
			log.Error(err)
		}

		log.Printf("addressResult: %#v", addressResult)

		if addressResult == nil {
			log.Printf("addressResult.InsertedIDs is nil")

			continue
		}

		for _, addr := range addressResult.InsertedIDs {
			hexAddress, ok := addr.(string)
			if !ok {
				log.Printf("type asserting common.Address error")

				continue
			}

			address := common.HexToAddress(hexAddress)

			log.Printf("  👚 addr: %#v", address)
			// log.Printf("  👚 common.BytesToAddress(addr.Data): %#v", common.BytesToAddress(addr.Data))
		}
	}

	degensResult, err := degensColl.InsertMany(context.TODO(), ogDegens)
	if err != nil {
		log.Error(err)
	}

	log.Printf("degensResult: %#v", degensResult.InsertedIDs)
}

func (ddb *DegenDB) AddCollectionToken(collections interface{}, tokens interface{}) {
	collectionsColl := ddb.mongo.Database(mongoDB).Collection(collCollections)

	mongoCollections := make([]interface{}, 0)

	collectionList, ok := collections.([]Address)
	if !ok {
		log.Printf("type asserting collections error")

		return
	}
	for _, collection := range collectionList {
		mongoCollections = append(mongoCollections, collection)
	}

	// result, err := degensColl.UpdateMany(context.TODO(), bson.D{{}}, ogDegens, &options.UpdateOptions{
	result, err := collectionsColl.InsertMany(context.TODO(), mongoCollections)
	if err != nil {
		log.Printf("add collections error: %+v", err)

		return
	}

	log.Printf("add collections result: %+v", result)

	mongoTokens := make([]interface{}, 0)

	tokenList, ok := tokens.([]Token)
	if !ok {
		log.Printf("ttype asserting tokens error: %+v", err)

		return
	}

	for _, token := range tokenList {
		mongoTokens = append(mongoTokens, token)
	}

	tokensColl := ddb.mongo.Database(mongoDB).Collection(collTokens)
	// result, err := degensColl.UpdateMany(context.TODO(), bson.D{{}}, ogDegens, &options.UpdateOptions{
	result, err = tokensColl.InsertMany(context.TODO(), mongoTokens)
	if err != nil {
		log.Error(err)
	}

	log.Printf("add tokens result: %+v", result)
}
