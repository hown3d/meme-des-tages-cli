package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/hown3d/meme-des-tages-cli/reddit"
	"github.com/pkg/browser"
)

var subreddit *string = flag.String("subreddit", "dankmemes", "subreddit to look in")

func main() {
	flag.Parse()
	client := reddit.NewClient()
	u, err := client.GetTopPost(context.TODO(), *subreddit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("found url %s\n", u)
	err = browser.OpenURL(u.String())
	if err != nil {
		log.Fatal(err)
	}
}
