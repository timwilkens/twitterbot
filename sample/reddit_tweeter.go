package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/timwilkens/twitterbot"
)

type RedditJson struct {
	Data Data `json:"data"`
}

type Data struct {
	Children []Kid `json:"children"`
}

type Kid struct {
	ChildData ChildData `json:"data"`
}

type ChildData struct {
	Title string `json:"title"`
}

type RedditTweeter struct{}

func (r RedditTweeter) NextTweet() string {
	resp, err := http.Get("http://www.reddit.com/.json")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var redditData RedditJson
	err = json.Unmarshal(body, &redditData)
	if err != nil {
		return ""
	}

	if len(redditData.Data.Children) > 1 {
		return redditData.Data.Children[0].ChildData.Title
	} else {
		return ""
	}
}

func main() {
	configPtr := flag.String("config", "", "Location of config")
	flag.Parse()
	if *configPtr == "" {
		fmt.Println("Usage: --config=/path/to/config")
		return
	}

	bot, err := twitterbot.New(*configPtr)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := RedditTweeter{}
	bot.Run(r)
}
