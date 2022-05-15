package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/aoisensi/go-yasuna/pkg/yasuna"
)

func main() {
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
	saveRefreshToken(token.RefreshToken)
	twitter := yasuna.NewTwitter(nil, oauth2, token)
	tweet, err := twitter.PostTweet(&yasuna.PostTweetParams{Text: "Hello, world!"})
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully posted tweet: http://twitter.com/apexkilllogbot/status/%v\n", tweet.Data.ID)
}

func saveRefreshToken(token string) {
	body := bytes.NewBufferString(url.Values{
		"encrypted_value": {token},
	}.Encode())
	u := "https://api.github.com/repos/aoisensi/apex-kill-log-bot/actions/secrets/TWITTER_REFRESH_TOKEN"
	req, err := http.NewRequest("PUT", u, body)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Token "+os.Getenv("GITHUB_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	if 200 <= code && code < 300 {
		log.Println("Refresh token saved")
		return // OK
	}
	data, _ := io.ReadAll(resp.Body)
	panic("failed to save refresh token: " + string(data))
}
