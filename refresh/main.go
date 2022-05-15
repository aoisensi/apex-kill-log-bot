package main

import (
	"fmt"
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
		Scope:        yasuna.ScopeAll,
	}
	if err := oauth2.Refresh(token); err != nil {
		panic(err)
	}
	fmt.Printf("::add-mask::%s\n", token.AccessToken)
	fmt.Printf("::add-mask::%s\n", token.RefreshToken)
	fmt.Printf("::set-output name=access_token::%s\n", token.AccessToken)
	fmt.Printf("::set-output name=refresh_token::%s\n", token.RefreshToken)
}
