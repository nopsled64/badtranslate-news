package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//TwitterConfig is this
type TwitterConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func readTwitterConfig(f string) TwitterConfig {
	file, _ := os.Open(f)
	defer file.Close()

	decoder := json.NewDecoder(file)

	twitterConfig := TwitterConfig{}
	err := decoder.Decode(&twitterConfig)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println(twitterConfig.AccessSecret) // output: [UserA, UserB]

	return twitterConfig
}

func postToTwitter(inputTweet string) error {

	twitterConfig := readTwitterConfig("twitter_config.json")

	config := oauth1.NewConfig(twitterConfig.ConsumerKey, twitterConfig.ConsumerSecret)
	token := oauth1.NewToken(twitterConfig.AccessToken, twitterConfig.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update(inputTweet, nil)
	if err != nil {
		fmt.Printf("Tweet posting error: %v\n", err)
		return fmt.Errorf("Tweet posting error: %v", err)
	}
	fmt.Println(tweet)
	fmt.Println(resp)

	return nil
}
