package degendb

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/nemo/marketplace"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kr/pretty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DegenDB struct {
	uri   string
	mongo *mongo.Client

	addressesColl   *mongo.Collection
	collectionsColl *mongo.Collection
	degensColl      *mongo.Collection
	tokensColl      *mongo.Collection

	ethClient *ethclient.Client
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

	log.Debug("mongodb ping successful")

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

	return &Address{ID: address.Hex()} // , IsContract: isContract, Type: addrType}
}

func (ddb *DegenDB) GetAccountType(address common.Address) AccountType {
	accountType := ExternallyOwnedAccount

	if ddb.IsContract(address) {
		accountType = ContractAccount
	}

	return accountType
}

func (ddb *DegenDB) IsContract(address common.Address) bool {
	// if its a marketplace address, its a contract
	if marketplace.Addresses().Contains(address) {
		return true
	}

	if ddb.ethClient == nil {
		return false
	}

	// ok 🙄 seems we really need to check via a node if its a eoa or contract
	codeAt, err := ddb.ethClient.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Debugf("❕ failed to get codeAt for %s: %s", address.String(), err)

		return false
	}

	log.Debugf("codeAt(%s): %+v", address.Hex(), codeAt)

	// if there is deployed code at the address, it's a contract
	return len(codeAt) > 0
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

// func (ddb *DegenDB) AddAddressWithFirst1KSlug(address common.Address, slug []Tag) {
// 	addressesColl := ddb.mongo.Database(mongoDB).Collection(collAddresses)

// 	addr := ddb.NewAddress(address)

// 	reallyTrue := true

// 	addressResult, err := addressesColl.UpdateByID(context.TODO(), addr.HexAddress, addr, &options.UpdateOptions{
// 		Upsert: &reallyTrue,
// 	})

// 	if err != nil {
// 		log.Error(err)
// 	}

// 	if addressResult == nil {
// 		log.Printf("addressResult.InsertedIDs is nil")

// 		continue
// 	}
// }

// /.
func (ddb *DegenDB) initializeDegensCollection() {
	var ogDegens []interface{}

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

		addressResult, err := ddb.addressesColl.InsertMany(context.TODO(), addresses)
		if err != nil {
			log.Errorf("InsertMany err: %+v", err)
		}

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

			_ = common.HexToAddress(hexAddress)
		}
	}

	if len(ogDegens) == 0 {
		log.Debugf("ogDegens is empty")

		return
	}

	_, err := ddb.degensColl.InsertMany(context.TODO(), ogDegens)
	if err != nil {
		log.Errorf("degensColl InsertMany err: %+v", err)
	}
}

func (ddb *DegenDB) AddCollections(collections interface{}) {
	mongoCollections := make([]interface{}, 0)

	collectionList, ok := collections.([]Collection)
	if !ok {
		log.Printf("type asserting collection error")

		return
	}
	for _, collection := range collectionList {
		mongoCollections = append(mongoCollections, collection)
	}

	result, err := ddb.collectionsColl.InsertMany(context.TODO(), mongoCollections)
	if err != nil {
		log.Printf("add collections error: %+v", err)

		return
	}

	log.Printf("add collections result: %+v", result)
}

func (ddb *DegenDB) AddCollectionToken(collections interface{}, tokens interface{}) {
	mongoCollections := make([]interface{}, 0)

	collectionList, ok := collections.([]Address)
	if !ok {
		log.Printf("type asserting collection token error")

		return
	}
	for _, collection := range collectionList {
		mongoCollections = append(mongoCollections, collection)
	}

	result, err := ddb.collectionsColl.InsertMany(context.TODO(), mongoCollections)
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

	result, err = ddb.tokensColl.InsertMany(context.TODO(), mongoTokens)
	if err != nil {
		log.Errorf("tokensColl.InsertMany err: %+v", err)
	}

	log.Printf("add tokens result: %+v", result)
}

// SaveContractInfo

func (ddb *DegenDB) First1KTxAlreadyFetchedFor(address common.Address) bool {
	addrFilter := bson.D{{Key: "_id", Value: address.Hex()}}

	var addr Address
	err := ddb.addressesColl.FindOne(context.Background(), addrFilter).Decode(&addr)
	if err != nil {
		// check for ErrNoDocuments
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Debugf("🥊 First1KTxAlreadyFetchedFor - mongo.ErrNoDocuments | %+v", address.Hex())

			// ddb.AddAddress(address)
		} else {
			log.Errorf("First1KTxAlreadyFetchedFor err: %+v", err)
		}

		return false
	}

	log.Debugf("addr: %+v | addr.First1KFetchedAt: %+v | addr.First1KFetchedAt.IsZero(): %+v", pretty.Sprint(addr), addr.First1KFetchedAt, addr.First1KFetchedAt.IsZero())

	if addr.First1KFetchedAt.IsZero() {
		return false
	}

	return true
}

// func (ddb *DegenDB) SetFirst1KFetchDate(address common.Address) {
// 	addrFilter := bson.D{{Key: "_id", Value: address.Hex()}}

// 	upd := bson.D{
// 		{Key: "$currentDate", Value: bson.D{{Key: "first1kFetchedAt", Value: true}}},
// 	}

// 	addressResult := ddb.addressesColl.FindOneAndUpdate(context.Background(), addrFilter, upd, &options.FindOneAndUpdateOptions{})

// 	if addressResult == nil {
// 		log.Error("addressResult.InsertedIDs is nil")

// 		return
// 	}
// }

// func (ddb *DegenDB) SaveAddressWithFirst1KSlugs(address common.Address, first1kSlugs []Tag) {
// 	addr := &Address{ID: address.Hex()}

// 	// use lower-case tags without special chars
// 	lowerTags := make([]Tag, 0)
// 	for _, tag := range first1kSlugs {
// 		lowerTag := strings.ToLower(string(tag))
// 		// TODO: find a better way to prepare the filename
// 		// pretty sure there exists something already
// 		specialCharPattern := regexp.MustCompile(`[!\/\[\]\':;.,<>?@#$%^&*()_+|{}~]`)
// 		sanitizedLowerTag := strings.ReplaceAll(specialCharPattern.ReplaceAllString(lowerTag, ""), " ", "_")

// 		lowerTags = append(lowerTags, Tag(sanitizedLowerTag))
// 	}

// 	addrFilter := bson.D{{Key: "_id", Value: address.Hex()}}

// 	upd := bson.D{
// 		{Key: "$set", Value: addr},
// 		{Key: "$setOnInsert", Value: bson.D{{Key: "created_at", Value: time.Now()}, {Key: "type", Value: ddb.GetAccountType(address)}}},
// 		{Key: "$currentDate", Value: bson.D{{Key: "updated_at", Value: true}}},
// 		{Key: "$addToSet", Value: bson.D{{Key: "first1K", Value: bson.D{{Key: "$each", Value: lowerTags}}}}}, // bson.D{{"First1K", bson.D{{"$each", first1kSlugs}}
// 	}

// 	upsert := true
// 	after := options.After

// 	result := ddb.addressesColl.FindOneAndUpdate(context.TODO(), addrFilter, upd, &options.FindOneAndUpdateOptions{
// 		ReturnDocument: &after,
// 		Upsert:         &upsert,
// 	})

// 	if result.Err() != nil && !errors.Is(result.Err(), mongo.ErrNoDocuments) {
// 		log.Errorf("result err: %+v", result.Err())

// 		return
// 	}

// 	if result == nil {
// 		log.Error("addressResult.InsertedIDs is nil")

// 		return
// 	} else if result.Err() != nil {
// 		// check for ErrNoDocuments
// 		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
// 			log.Debugf("addressResult.Err() is mongo.ErrNoDocuments")
// 		} else {
// 			log.Errorf("result err: %+v", result.Err())
// 		}

// 		return
// 	}

// 	// var decoded interface{}
// 	// addressResult.Decode(&decoded)

// 	log.Debug("addressResult: %+v", result)
// }

func (ddb *DegenDB) SaveAddressFirst1KFetchedAt(address common.Address, contractInfo *ContractInfo) (result *mongo.SingleResult) {
	upd := bson.E{Key: "$currentDate", Value: bson.D{{Key: "first1kFetchedAt", Value: true}}}

	return ddb.SaveAddress(address, contractInfo, upd)
}

func (ddb *DegenDB) SaveAddressFirst1KSlugs(address common.Address, first1kSlugs []Tag) (result *mongo.SingleResult) {
	// use lower-case tags without special chars
	lowerTags := make([]Tag, 0)
	for _, tag := range first1kSlugs {
		lowerTag := strings.ToLower(string(tag))
		// TODO: find a better way to prepare the filename
		// pretty sure there exists something already
		specialCharPattern := regexp.MustCompile(`[!\/\[\]\':;.,<>?@#$%^&*()_+|{}~]`)
		sanitizedLowerTag := strings.ReplaceAll(specialCharPattern.ReplaceAllString(lowerTag, ""), " ", "_")

		lowerTags = append(lowerTags, Tag(sanitizedLowerTag))
	}

	upd := bson.E{Key: "$addToSet", Value: bson.D{{Key: "first1K", Value: bson.D{{Key: "$each", Value: lowerTags}}}}}

	return ddb.SaveAddress(address, nil, upd)
}

func (ddb *DegenDB) SaveAddress(address common.Address, contractInfo *ContractInfo, updEntries ...primitive.E) (result *mongo.SingleResult) {
	// use lower-case tags without special chars
	lowerTags := make([]Tag, 0)
	for _, tag := range lowerTags {
		lowerTag := strings.ToLower(string(tag))
		// TODO: find a better way to prepare the filename
		// pretty sure there exists something already
		specialCharPattern := regexp.MustCompile(`[!\/\[\]\':;.,<>?@#$%^&*()_+|{}~]`)
		sanitizedLowerTag := strings.ReplaceAll(specialCharPattern.ReplaceAllString(lowerTag, ""), " ", "_")

		lowerTags = append(lowerTags, Tag(sanitizedLowerTag))
	}

	// filter to find the address
	addrFilter := bson.D{{Key: "_id", Value: address.Hex()}}

	// create a new address
	newAddress := &Address{ID: address.Hex()}

	// optional collection info
	if contractInfo != nil {
		newAddress.Name = contractInfo.Name
		newAddress.SlugOpenSea = contractInfo.Collection
		// addr.ContractStandard = contractInfo.ContractStandard
		// addr.Supply = contractInfo.Supply
	}

	// base data for upsert
	upd := bson.D{
		{Key: "$set", Value: newAddress},
		{Key: "$setOnInsert", Value: bson.D{{Key: "created_at", Value: time.Now()}, {Key: "type", Value: ddb.GetAccountType(address)}}},
		{Key: "$currentDate", Value: bson.D{{Key: "updated_at", Value: true}}},
		// {Key: "$addToSet", Value: bson.D{{Key: "first1K", Value: bson.D{{Key: "$each", Value: lowerTags}}}}}, // bson.D{{"First1K", bson.D{{"$each", first1kSlugs}}
	}

	// add variable updates
	upd = append(upd, updEntries...)

	// upsert options
	upsert := true
	after := options.After

	result = ddb.addressesColl.FindOneAndUpdate(context.TODO(), addrFilter, upd, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	})

	if result.Err() != nil && !errors.Is(result.Err(), mongo.ErrNoDocuments) {
		log.Errorf("result err: %+v", result.Err())

		return nil
	}

	if result == nil {
		log.Error("addressResult.InsertedIDs is nil")

		return nil
	} else if result.Err() != nil {
		// check for ErrNoDocuments
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			log.Debugf("addressResult.Err() is mongo.ErrNoDocuments")
		} else {
			log.Errorf("result err: %+v", result.Err())
		}

		return result
	}

	// var decoded interface{}
	// addressResult.Decode(&decoded)

	log.Debug("addressResult: %+v", result)

	return result
}

// func (ddb *DegenDB) AddAddress(contractInfo *ContractInfo, address common.Address) error {
// 	// addr := ddb.NewAddress(address)
// 	addr := &Address{ID: address.Hex(), Type: ddb.GetAccountType(address), CreatedAt: time.Now(), UpdatedAt: time.Now()}

// 	if contractInfo != nil {
// 		addr.Name = contractInfo.Name
// 		addr.SlugOpenSea = contractInfo.Collection
// 		// addr.ContractStandard = contractInfo.ContractStandard
// 		// addr.Supply = contractInfo.Supply
// 	}

// 	result, err := ddb.addressesColl.InsertOne(context.TODO(), addr, &options.InsertOneOptions{})
// 	if err != nil {
// 		log.Errorf("result err: %+v", err)

// 		return err
// 	}

// 	log.Debugf("addAddress result: %+v", result)

// 	return nil
// }
