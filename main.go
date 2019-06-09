// Sample translate-quickstart translates "Hello, world!" into Russian.
package main

import (
	"fmt"
	"html"
	"os"
	"time"
)

func main() {

	rss, err := getNewsRSS()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	for i := 0; i < len(rss.Channels[0].Items); i++ {
		fmt.Println("News Title: " + rss.Channels[0].Items[i].Title)
		fmt.Println("	Description: " + rss.Channels[0].Items[i].Description)
		reportDate, err := time.Parse(time.RFC1123, rss.Channels[0].Items[i].PubDate)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("	Date of report: " + reportDate.String())
	}

	var actualNews = rss.Channels[0].Items[0].Title + " - " + rss.Channels[0].Items[0].Description

	fmt.Println("Here is the news report: ", actualNews)

	badText, err := badTranslate(actualNews)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Badly translated passage: %v\n", html.UnescapeString(badText))

	var tweetToPost = html.UnescapeString(badText)

	fmt.Printf("This would have been posted: \"%v\"", tweetToPost)

	postToTwitter(tweetToPost)
	// Add some eror handling for post length
	// add some error handling for not reposting the same thing
}
