// Sample translate-quickstart translates "Hello, world!" into Russian.
package main

import (
	"fmt"
	"os"
)

func main() {
	// badText, err := badTranslate("Michael Gove admits he was lucky to avoid jail over cocaine use")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("Badly translated passage: %v\n", html.UnescapeString(badText))

	rss, err := getNewsRSS()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	for i := 0; i < len(rss.Channels[0].Items); i++ {
		fmt.Println("News Title: " + rss.Channels[0].Items[i].Title)
		fmt.Println("	Description: " + rss.Channels[0].Items[i].Description)
	}
}
