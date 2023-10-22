package main

import (
	"fmt"
	"log"

	"github.com/grokify/gorod"
)

func main() {
	wantURL := "https://github.com/grokify"

	_, err := gorod.RetrieveWriteScreenshotFullPage(nil, wantURL, "example.png", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
