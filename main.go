package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aoisensi/go-yasuna/pkg/yasuna"
)

func main() {
	fmt.Println(os.Getenv("TWITTER_CLIENT_ID"))
	fmt.Println(os.Getenv("TWITTER_CLIENT_SECRET"))
	fmt.Println(os.Getenv("TWITTER_REFRESH_TOKEN"))
	oauth2 := &yasuna.OAuth2{
		ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
	}
	token := &yasuna.Token{
		RefreshToken: os.Getenv("TWITTER_REFRESH_TOKEN"),
	}
	if err := oauth2.Refresh(token); err != nil {
		panic(err)
	}
	fmt.Printf("::add-mask::%s\n", token.AccessToken)
	fmt.Printf("::set-output name=refresh_token::%s\n", token.AccessToken)
	twitter := yasuna.NewTwitter(nil, oauth2, token)
	tweet, err := twitter.PostTweet(&yasuna.PostTweetParams{Text: "Hello, world!"})
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully posted tweet: http://twitter.com/ApexKillLogBot/status/%v\n", tweet.Data.ID)
}
