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
	collCollections = "collections"
	collDegens      = "degens"
	collTokens      = "tokens"
)

type DegenDB struct {
	mongo *mongo.Client
}

func NewDegenDB() *DegenDB {
	mongo := connectToMongo()
	if mongo == nil {
		log.Errorf("❌ could not connect to MongoDB at: %s", viper.GetString("mongodb.uri"))

		return nil
	}

	// defer func() {
	// 	if err := mongo.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	ddb := &DegenDB{
		mongo: mongo,
	}

	// cleanup & initialize
	collectionsColl := ddb.mongo.Database(mongoDB).Collection(collCollections)
	degensColl := ddb.mongo.Database(mongoDB).Collection(collDegens)
	tokensColl := ddb.mongo.Database(mongoDB).Collection(collTokens)

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

	ddb.initializeDegensCollection()

	// check
	cursor, err := degensColl.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	var mongoDegen []Degen
	if err = cursor.All(context.TODO(), &mongoDegen); err != nil {
		panic(err)
	}
	for _, dgn := range mongoDegen {
		log.Printf("mongoDegen: %+v", dgn)
	}

	return ddb
}

func (ddb *DegenDB) initializeDegensCollection() {
	degensColl := ddb.mongo.Database(mongoDB).Collection(collDegens)

	ogDegens := []interface{}{
		Degen{
			Name:     "Ben",
			Accounts: Accounts{Twitter: "ben_leb", Telegram: "benleb", TelegramChatID: -1001808788625},
			Wallets: []Wallet{
				{Address: common.HexToAddress("0xCd8aF79Ba3974404e37f126a8E355690351Da8bD")},
			},
			Tags:      []Tag{"dakma"},
			CreatedAt: time.Now(),
		},
		Degen{
			Name:      "Luke",
			Accounts:  Accounts{Twitter: "0xlugges", Telegram: "luke_take_profits", TelegramChatID: 1337},
			Wallets:   []Wallet{{Address: common.HexToAddress("0x0DB54CC560Ae7832898e82E5E607E8142e519891")}},
			Tags:      []Tag{"dakma"},
			CreatedAt: time.Now(),
		},
	}

	// result, err := degensColl.UpdateMany(context.TODO(), bson.D{{}}, ogDegens, &options.UpdateOptions{
	result, err := degensColl.InsertMany(context.TODO(), ogDegens)
	if err != nil {
		panic(err)
	}

	log.Printf("result: %+v", result)
}

func (ddb *DegenDB) AddCollectionToken(collections interface{}, tokens interface{}) {
	collectionsColl := ddb.mongo.Database(mongoDB).Collection(collCollections)

	mongoCollections := make([]interface{}, 0)

	collectionList, ok := collections.([]Collection)
	if !ok {
		log.Printf("type asserting collections error")
		// panic(err)

		return
	}
	for _, collection := range collectionList {
		mongoCollections = append(mongoCollections, collection)
	}

	// result, err := degensColl.UpdateMany(context.TODO(), bson.D{{}}, ogDegens, &options.UpdateOptions{
	result, err := collectionsColl.InsertMany(context.TODO(), mongoCollections)
	if err != nil {
		log.Printf("add collections error: %+v", err)
		// panic(err)

		return
	}

	log.Printf("add collections result: %+v", result)

	mongoTokens := make([]interface{}, 0)

	tokenList, ok := tokens.([]Token)
	if !ok {
		log.Printf("ttype asserting tokens error: %+v", err)
		// panic(err)

		return
	}

	for _, token := range tokenList {
		mongoTokens = append(mongoTokens, token)
	}

	tokensColl := ddb.mongo.Database(mongoDB).Collection(collTokens)
	// result, err := degensColl.UpdateMany(context.TODO(), bson.D{{}}, ogDegens, &options.UpdateOptions{
	result, err = tokensColl.InsertMany(context.TODO(), mongoTokens)
	if err != nil {
		panic(err)
	}

	log.Printf("add tokens result: %+v", result)
}

func connectToMongo() *mongo.Client {
	uri := viper.GetString("mongodb.uri")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	log.Print("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}
