# Release Notes: v0.1.0

**Release Date:** 2023-10-22

## Overview

GoRod v0.1.0 is the initial release of a helper library for the [Rod](https://github.com/go-rod/rod) web automation library. It provides utilities for interactive browser sessions and full-page screenshot capture.

## Key Features

### Interactive Browser Sessions

The `ForegroundBrowser` type enables browser automation with manual intervention points:

- Launch a visible browser window for user interaction
- Navigate to URLs with configurable delays
- Pause execution for manual actions (login, CAPTCHA, 2FA)
- Resume automation after user completes manual steps

```go
fb, err := gorod.NewForegroundBrowserPaused("https://example.com/login", 2, true)
if err != nil {
    log.Fatal(err)
}
defer fb.Close()
// Browser pauses here for user to log in manually
```

### Full-Page Screenshots

Capture complete page screenshots with HTTP status handling:

```go
page, err := gorod.RetrieveWriteScreenshotFullPage(
    nil,                    // browser (nil creates new)
    "https://example.com",  // URL
    "screenshot.png",       // output file
    nil,                    // options
)
```

### HTML File Retrieval

Fetch and save rendered HTML content:

```go
err := fb.GetWriteFileHTML(
    "https://example.com",
    "output.html",
    0644,
    false,          // force overwrite
    time.Second,    // write delay
)
```

## Installation

```bash
go get github.com/grokify/gorod
```

## Use Cases

- **Session capture**: Pause for manual login, then continue automation
- **CAPTCHA handling**: Let users solve CAPTCHAs before proceeding
- **Screenshot automation**: Capture full-page screenshots of rendered pages
- **Content archival**: Save rendered HTML from JavaScript-heavy sites

## CI/CD

GitHub Actions workflows for automated quality assurance:

- **CI**: Go test workflow
- **Linting**: golangci-lint static analysis
- **Dependabot**: Automated dependency updates

## Dependencies

- [go-rod/rod](https://github.com/go-rod/rod) - Browser automation library
- [grokify/mogo](https://github.com/grokify/mogo) - Go utility library

## Links

- [GitHub Repository](https://github.com/grokify/gorod)
- [GoDoc](https://pkg.go.dev/github.com/grokify/gorod)
- [Issue Tracker](https://github.com/grokify/gorod/issues)
