// Sample translate-quickstart translates "Hello, world!" into Russian.
package main

import (
	"fmt"
	"html"
	"os"
)

func main() {
	badText, err := badTranslate("Michael Gove admits he was lucky to avoid jail over cocaine use")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Badly translated passage: %v\n", html.UnescapeString(badText))
}
