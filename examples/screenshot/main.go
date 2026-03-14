package main

import (
	"fmt"
	"log"

	"github.com/grokify/gorod"
)

func main() {
	wantURL := "https://analysis.dinq.me/github?user=grokify"
	outfile := "dinkq-grokify.png"

	_, err := gorod.RetrieveWriteScreenshotFullPage(nil, wantURL, outfile, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
