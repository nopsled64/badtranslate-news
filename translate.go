package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func badTranslate(t string) (string, error) {
	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return "", err
	}

	languages := []string{"ru", "af", "nl", "fy", "ko", "pl", "en"}

	for i, l := range languages {
		// Sets the target language.
		target, err := language.Parse(l)
		if err != nil {
			log.Fatalf("Failed to parse target language: %v", err)
			return "", err
		}

		// Translates the text into Russian.
		translations, err := client.Translate(ctx, []string{t}, target, nil)
		if err != nil {
			log.Fatalf("Failed to translate text: %v", err)
			return "", err
		}

		t = translations[0].Text

		fmt.Printf("Translation %v: translation into %v:%v\n", i, l, t)
		fmt.Println()

	}
	return t, nil
}
