package main

import (
	"fmt"
	"log"

	"github.com/grokify/gorod"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	URL string `short:"u" long:"url" description:"URL for webhook" required:"true"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	fb, err := gorod.NewForegroundBrowserPaused(opts.URL, 2, true)
	if err != nil {
		log.Fatal(err)
	} else {
		defer fb.Close()
	}
	c, err := fb.Cookies()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cookies: %s\n", c.String())

	fmt.Println("DONE")
}
