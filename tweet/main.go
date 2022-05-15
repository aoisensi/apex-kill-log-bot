package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aoisensi/go-yasuna/pkg/yasuna"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	token := &yasuna.Token{
		AccessToken: os.Getenv("TWITTER_ACCESS_TOKEN"),
	}
	userID, _ := strconv.ParseInt(os.Getenv("TWITTER_USER_ID"), 10, 64)
	twitter := yasuna.NewTwitter(nil, nil, token)
	data, err := twitter.GetUserFollowers(userID, yasuna.WithMaxResults(1000))
	if err != nil {
		panic(err)
	}

	followers := data.Data
	if len(followers) < 2 {
		log.Fatal("low followers")
	}

	killer := randomUser(followers)
	deader := randomUser(followers)
	for killer.ID == deader.ID {
		deader = randomUser(followers)
	}

	rate := random.Intn(MaxRate)

	var kill Kill

	for _, k := range kills {
		rate -= k.Rate
		if rate <= 0 {
			kill = k
			break
		}
	}

	text := fmt.Sprintf(kill.Format, killer.Username, deader.Username)

	tweet, err := twitter.PostTweet(&yasuna.PostTweetParams{Text: text})
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully posted tweet: http://twitter.com/ApexKillLogBot/status/%v\n", tweet.Data.ID)
}

func randomUser(users []*yasuna.User) *yasuna.User {
	return users[random.Intn(len(users))]
}
