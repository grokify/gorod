package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:     "gorod",
	Short:   "Headless browser CLI tool",
	Long:    "A CLI tool for fetching and rendering web pages using a headless browser.",
	Version: version,
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

// Fetch command flags
var (
	fetchURL        string
	fetchOutput     string
	fetchSelector   string
	fetchWaitStable bool
	fetchTimeout    int
	fetchHeadless   bool
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch and render a web page",
	Long: `Fetch a URL using a headless browser, wait for JavaScript to render,
and output the page content.

Examples:
  gorod fetch -u https://example.com
  gorod fetch -u https://example.com -o html
  gorod fetch -u https://example.com -s ".content" -w`,
	RunE: runFetch,
}

func init() {
	fetchCmd.Flags().StringVarP(&fetchURL, "url", "u", "", "URL to fetch (required)")
	fetchCmd.Flags().StringVarP(&fetchOutput, "output", "o", "text", "Output format: text, html")
	fetchCmd.Flags().StringVarP(&fetchSelector, "selector", "s", "", "CSS selector to extract (default: body)")
	fetchCmd.Flags().BoolVarP(&fetchWaitStable, "wait-stable", "w", false, "Wait for page to be stable")
	fetchCmd.Flags().IntVarP(&fetchTimeout, "timeout", "t", 30, "Timeout in seconds")
	fetchCmd.Flags().BoolVar(&fetchHeadless, "headless", true, "Run in headless mode")

	if err := fetchCmd.MarkFlagRequired("url"); err != nil {
		panic(err)
	}
}

func runFetch(cmd *cobra.Command, args []string) error {
	// Create launcher
	l := launcher.New().Headless(fetchHeadless)
	defer l.Cleanup()

	url := l.MustLaunch()

	// Create browser
	browser := rod.New().
		ControlURL(url).
		MustConnect()
	defer browser.MustClose()

	// Set timeout
	browser = browser.Timeout(time.Duration(fetchTimeout) * time.Second)

	// Navigate to page
	page := browser.MustPage(fetchURL)

	// Wait for page load
	page.MustWaitLoad()

	// Optionally wait for stability (useful for JS-heavy pages)
	if fetchWaitStable {
		page.MustWaitStable()
	}

	// Get content
	selector := fetchSelector
	if selector == "" {
		selector = "body"
	}

	el := page.MustElement(selector)

	var content string
	switch fetchOutput {
	case "html":
		content = el.MustHTML()
	case "text":
		fallthrough
	default:
		content = el.MustText()
	}

	fmt.Println(content)
	return nil
}
