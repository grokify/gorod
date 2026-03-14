# GoRod

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/gorod/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/gorod/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/gorod/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/gorod/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/gorod/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/gorod/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/gorod
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/gorod
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/gorod
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/gorod
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fgorod
 [loc-svg]: https://tokei.rs/b1/github/grokify/gorod
 [repo-url]: https://github.com/grokify/gorod
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/gorod/blob/master/LICENSE

GoRod is a Go package providing helper utilities for the [Rod](https://github.com/go-rod/rod) web automation library. It simplifies common browser automation tasks like interactive sessions, cookie extraction, element waiting, and screenshot capture.

## Installation

```bash
go get github.com/grokify/gorod
```

### CLI Tool

```bash
go install github.com/grokify/gorod/cmd/gorod@latest
```

## Features

- **Interactive Browser Sessions** - Launch browsers that pause for manual interaction (e.g., login, CAPTCHA)
- **Cookie Extraction** - Extract browser cookies and convert to standard `http.Cookie` format
- **Wait Helpers** - Wait for elements to be visible or clickable with timeout
- **Screenshot Capture** - Full-page screenshots with HTTP status handling
- **CLI Tool** - Command-line tool for headless page fetching

## Quick Start

### Interactive Browser Session

Launch a browser, navigate to a URL, and pause for user interaction:

```go
package main

import (
    "fmt"
    "log"

    "github.com/grokify/gorod"
)

func main() {
    // Open browser, navigate to URL, pause for login
    fb, err := gorod.NewForegroundBrowserPaused("https://example.com/login", 2, true)
    if err != nil {
        log.Fatal(err)
    }
    defer fb.Close()

    // Extract cookies after user logs in
    cookies, err := fb.Cookies()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Session cookies: %s\n", cookies.String())
}
```

### Cookie Extraction

Convert Rod cookies to standard HTTP cookies:

```go
cookies, _ := fb.Cookies()

// Get as []*http.Cookie
httpCookies := cookies.HTTPCookies()

// Get as cookie header string
headerValue := cookies.String()
```

### Wait for Elements

Wait for elements to be visible or clickable:

```go
import (
    "time"
    "github.com/go-rod/rod"
    "github.com/grokify/gorod"
)

page := browser.MustPage("https://example.com")

// Wait for element to be visible
el, err := gorod.WaitVisible(page, "#login-button", 10*time.Second)
if err != nil {
    log.Fatal("Login button not found")
}

// Wait for element to be clickable (visible and enabled)
el, err = gorod.WaitClickable(page, "#submit", 10*time.Second)
if err != nil {
    log.Fatal("Submit button not clickable")
}
el.MustClick()
```

### Full-Page Screenshot

Capture a full-page screenshot with HTTP status handling:

```go
import "github.com/grokify/gorod"

page, err := gorod.RetrieveWriteScreenshotFullPage(
    nil,                    // browser (nil creates new)
    "https://example.com",  // URL
    "screenshot.png",       // output file
    nil,                    // screenshot options
)
if err != nil {
    log.Fatal(err)
}
defer page.Close()
```

## CLI Usage

The `gorod` CLI provides headless browser operations from the command line.

### Fetch Command

Render a page and extract content:

```bash
# Get page text
gorod fetch -u https://example.com

# Get page HTML
gorod fetch -u https://example.com -o html

# Extract specific element
gorod fetch -u https://example.com -s ".main-content" -o html

# Wait for page stability (JS-heavy pages)
gorod fetch -u https://example.com -w

# With timeout
gorod fetch -u https://example.com -t 60
```

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--url` | `-u` | (required) | URL to fetch |
| `--output` | `-o` | `text` | Output format: `text` or `html` |
| `--selector` | `-s` | `body` | CSS selector to extract |
| `--wait-stable` | `-w` | `false` | Wait for page stability |
| `--timeout` | `-t` | `30` | Timeout in seconds |
| `--headless` | | `true` | Run in headless mode |

## API Reference

### Types

#### ForegroundBrowser

Interactive browser session with pause capability.

```go
type ForegroundBrowser struct {
    Launcher *launcher.Launcher
    Browser  *rod.Browser
}

// Create new browser, navigate to URL, optionally pause
func NewForegroundBrowserPaused(navURL string, delaySeconds int, paused bool) (ForegroundBrowser, error)

// Close browser and cleanup
func (fb *ForegroundBrowser) Close()

// Extract cookies from browser
func (fb *ForegroundBrowser) Cookies() (Cookies, error)

// Fetch and write HTML to file
func (fb *ForegroundBrowser) GetWriteFileHTML(url, filename string, perm os.FileMode, force bool, writeDelay time.Duration) error
```

#### Cookies

Browser cookie collection with conversion utilities.

```go
type Cookies []*proto.NetworkCookie

// Convert to standard HTTP cookies
func (c Cookies) HTTPCookies() []*http.Cookie

// Serialize to cookie header string
func (c Cookies) String() string
```

### Functions

#### Wait Helpers

```go
// Wait for element to become visible
func WaitVisible(page *rod.Page, selector string, timeout time.Duration) (*rod.Element, error)

// Wait for element to be visible and enabled
func WaitClickable(page *rod.Page, selector string, timeout time.Duration) (*rod.Element, error)
```

#### Screenshot

```go
// Capture full-page screenshot with HTTP status handling
func RetrieveWriteScreenshotFullPage(browser *rod.Browser, srcURL, filename string, opts *proto.PageCaptureScreenshot) (*rod.Page, error)
```

## Examples

See the [`examples/`](examples/) directory:

- [`examples/cookies/`](examples/cookies/) - Cookie extraction demo
- [`examples/screenshot/`](examples/screenshot/) - Screenshot capture demo

## Use Cases

- **Session capture**: Extract authenticated session cookies for API testing
- **Web scraping**: Render JavaScript-heavy pages and extract content
- **E2E testing helpers**: Wait utilities for reliable element interaction
- **Screenshot automation**: Capture full-page screenshots for visual testing
- **Manual authentication flows**: Pause for CAPTCHA, 2FA, or complex logins

## Dependencies

- [go-rod/rod](https://github.com/go-rod/rod) - Main browser automation library
- [grokify/mogo](https://github.com/grokify/mogo) - Go utility library

## License

MIT License - see [LICENSE](LICENSE) for details.
