package twitterbot

import (
	"fmt"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

type TBot struct {
	api  *anaconda.TwitterApi
	keys *Keys
}

func NewBot(config string) (*TBot, error) {
	keys, err := ReadConfig(config)
	if err != nil {
		return nil, err
	}

	anaconda.SetConsumerKey(keys.consumerPublic)
	anaconda.SetConsumerSecret(keys.consumerSecret)
	api := anaconda.NewTwitterApi(keys.accessPublic, keys.accessSecret)

	return &TBot{api, keys}, nil
}

type TweetCreator interface {
	NextTweet() string
}

func (t *TBot) RunBot(creator TweetCreator) {
	var previousTweet string

	for {
		tweet := creator.NextTweet()
		if previousTweet == "" || previousTweet != tweet {
			fmt.Println("[" + time.Now().Format(time.RFC850) + "] Posting " + tweet)
			t.api.PostTweet(tweet, nil)
			previousTweet = tweet
		}
		fmt.Println("[" + time.Now().Format(time.RFC850) + "] Sleeping...")
		time.Sleep(10 * time.Minute)
	}
}
