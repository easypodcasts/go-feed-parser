package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/easypodcasts/go-feed-parser/build"
	"github.com/mmcdole/gofeed"
)

func main() {
	url := envString("FEED_URL", "")
	fFeedURL := flag.String("feed.url", url, "feed url to parse")
	version := flag.Bool("v", false, "prints current version")
	flag.Parse()

	if *version {
		fmt.Println(build.Version)
		os.Exit(0)
	}

	if *fFeedURL == "" {
		flag.Usage()
		os.Exit(0)
	}

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(*fFeedURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(feed)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}

	return e
}
