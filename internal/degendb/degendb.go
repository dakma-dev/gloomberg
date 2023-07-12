package degendb

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoDB         = "dev-degendb"
	collCollections = "collections"
	collDegens      = "degens"
	collTokens      = "tokens"
	collWalllets    = "wallets"
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
		walletsColl := ddb.mongo.Database(mongoDB).Collection(collWalllets)

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

		err = walletsColl.Drop(context.Background())
		if err != nil {
			log.Error(err)
		}

		// initialize og degens
		ddb.initializeDegensCollection()

		// check
		cursor, err := degensColl.Find(context.Background(), bson.D{})
		if err != nil {
			log.Error(err)
		}

		// query
		var mongoDegen []Degen
		if err = cursor.All(context.TODO(), &mongoDegen); err != nil {
			log.Error(err)
		}

		// print
		for _, dgn := range mongoDegen {
			log.Printf("mongoDegen: %+v", dgn)
		}
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

func (ddb *DegenDB) initializeDegensCollection() {
	degensColl := ddb.mongo.Database(mongoDB).Collection(collDegens)
	walletsColl := ddb.mongo.Database(mongoDB).Collection(collWalllets)

	ogDegens := []interface{}{}

	for _, degen := range ogDegens {
		degen, ok := degen.(*Degen)
		if !ok {
			log.Printf("type asserting degen error")

			continue
		}
		// og := ogDegens[degenID].(Degen)

		wallets := make([]interface{}, 0)

		for _, wallet := range degen.RawWallets {
			// wallets[degen.Name] = append(wallets[degen.Name], wallet.ID)
			wallets = append(wallets, wallet)
		}

		walletsResult, err := walletsColl.InsertMany(context.TODO(), wallets)
		if err != nil {
			log.Error(err)
		}

		for _, walletID := range walletsResult.InsertedIDs {
			wID, ok := walletID.(primitive.ObjectID)
			if !ok {
				log.Printf("type asserting walletID error")

				continue
			}

			degen.Wallets = append(degen.Wallets, wID)
		}
	}

	degensResult, err := degensColl.InsertMany(context.TODO(), ogDegens)
	if err != nil {
		log.Error(err)
	}

	log.Printf("result: %+v", degensResult)
}

func (ddb *DegenDB) AddCollectionToken(collections interface{}, tokens interface{}) {
	collectionsColl := ddb.mongo.Database(mongoDB).Collection(collCollections)

	mongoCollections := make([]interface{}, 0)

	collectionList, ok := collections.([]Collection)
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
