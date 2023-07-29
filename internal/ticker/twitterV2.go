package ticker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/spf13/viper"
)

type TwitterCredentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

type TwitterClient struct {
	twitterClient *twitter.Client
}

type authorize struct {
	Token string
}

func (a authorize) Add(*http.Request) {
}

func (c *TwitterClient) PostTweetV2(msg string) {
	if !viper.GetBool("twitter.enabled") {
		log.Printf("twitter not enabled")

		return
	}

	if len(msg) > 400 {
		log.Printf("tweet too long: %v", len(msg))

		return
	}

	rateLimit := callout(c.twitterClient, msg)
	fmt.Printf("Create Tweet Response: %v\n", *rateLimit)
}

func NewTwitterClient(cred *TwitterCredentials) *TwitterClient {
	if !viper.GetBool("twitter.enabled") {
		log.Printf("twitter not enabled")

		return nil
	}

	client := getClient(cred)

	return &TwitterClient{
		twitterClient: client,
	}
}

func getClient(cred *TwitterCredentials) *twitter.Client {
	//  Supported authentication types are [OAuth 1.0a User Context, OAuth 2.0 User Context]
	config := oauth1.NewConfig(cred.ConsumerKey, cred.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, &oauth1.Token{
		Token:       cred.AccessToken,
		TokenSecret: cred.AccessTokenSecret,
	})

	client := &twitter.Client{
		Authorizer: authorize{
			// Token: viper.GetString("twitter.bearer_token"),
		},
		Client: httpClient,
		Host:   "https://api.twitter.com",
	}

	return client
}

func callout(client *twitter.Client, msg string) *twitter.CreateTweetResponse {
	request := twitter.CreateTweetRequest{
		DirectMessageDeepLink: "",
		ForSuperFollowersOnly: false,
		QuoteTweetID:          "",
		Text:                  msg,
		ReplySettings:         "",
		Geo:                   nil,
		Media:                 nil,
		Poll:                  nil,
		Reply:                 nil,
	}

	tweet, err := client.CreateTweet(context.Background(), request)
	if err != nil {
		log.Printf("create Tweets error: %v", err)

		return nil
	}

	enc, err := json.MarshalIndent(tweet, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))

	return tweet
}
