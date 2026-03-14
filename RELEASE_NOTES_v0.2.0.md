# Release Notes: v0.2.0

**Release Date:** 2026-03-14

## Overview

GoRod v0.2.0 introduces a CLI tool for headless browser operations, cookie extraction utilities, and wait helper functions for reliable element interaction. These additions make GoRod more versatile for both library usage and command-line workflows.

## Highlights

- **CLI Tool**: New `gorod fetch` command for headless page rendering from the command line
- **Cookie Extraction**: Extract browser cookies and convert to standard `http.Cookie` format
- **Wait Helpers**: Reliable element waiting with `WaitVisible()` and `WaitClickable()` functions

## New Features

### CLI Tool

The new `gorod` CLI provides headless browser operations from the command line:

```bash
# Install
go install github.com/grokify/gorod/cmd/gorod@latest

# Fetch page text
gorod fetch -u https://example.com

# Fetch page HTML
gorod fetch -u https://example.com -o html

# Extract specific element
gorod fetch -u https://example.com -s ".main-content" -o html

# Wait for page stability (JS-heavy pages)
gorod fetch -u https://example.com -w --timeout 60
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--url` | `-u` | (required) | URL to fetch |
| `--output` | `-o` | `text` | Output format: `text` or `html` |
| `--selector` | `-s` | `body` | CSS selector to extract |
| `--wait-stable` | `-w` | `false` | Wait for page stability |
| `--timeout` | `-t` | `30` | Timeout in seconds |
| `--headless` | | `true` | Run in headless mode |

### Cookie Extraction

New `Cookies` type for extracting and converting browser cookies:

```go
fb, _ := gorod.NewForegroundBrowserPaused("https://example.com", 2, true)
defer fb.Close()

// Extract cookies after user interaction
cookies, err := fb.Cookies()
if err != nil {
    log.Fatal(err)
}

// Convert to standard http.Cookie slice
httpCookies := cookies.HTTPCookies()

// Get as cookie header string
headerValue := cookies.String()
fmt.Printf("Cookie: %s\n", headerValue)
```

### Wait Helpers

New functions for reliable element interaction:

```go
// Wait for element to become visible
el, err := gorod.WaitVisible(page, "#login-form", 10*time.Second)
if err != nil {
    log.Fatal("Login form not found")
}

// Wait for element to be visible AND enabled (clickable)
btn, err := gorod.WaitClickable(page, "#submit-button", 5*time.Second)
if err != nil {
    log.Fatal("Submit button not clickable")
}
btn.MustClick()
```

### Improved ForegroundBrowser

- **Better input handling**: Uses `bufio.Reader` for more reliable pause/resume behavior
- **Page stability**: Adds `MustWaitStable()` after navigation for JS-heavy pages
- **Error handling**: New `ErrBrowserNotInitialized` error for nil browser detection

## Installation

```bash
# Library
go get github.com/grokify/gorod

# CLI
go install github.com/grokify/gorod/cmd/gorod@latest
```

## Examples

New examples demonstrating v0.2.0 features:

- [`examples/cookies/`](examples/cookies/) - Cookie extraction and conversion
- [`examples/screenshot/`](examples/screenshot/) - Screenshot capture (updated)

## Dependencies

- Update `github.com/grokify/mogo` to v0.73.5
- Add `github.com/spf13/cobra` for CLI framework
- Add `github.com/jessevdk/go-flags` for example CLI parsing
- Add `fetchup` replace directive for rod v0.116.2 compatibility

## Upgrade Guide

This release is backwards compatible with v0.1.x. New features are additive.

**Behavioral change:** `NewForegroundBrowserPaused()` now calls `MustWaitStable()` after page navigation, which may slightly increase wait times but improves reliability on JavaScript-heavy pages.

## What's Next

Planned for future releases:

- Additional wait helpers (WaitText, WaitAttribute)
- Form filling utilities
- Multiple browser profile support
- Screenshot comparison tools

## Links

- [GitHub Repository](https://github.com/grokify/gorod)
- [Documentation](https://pkg.go.dev/github.com/grokify/gorod)
- [Changelog](https://github.com/grokify/gorod/blob/main/CHANGELOG.md)
- [Issue Tracker](https://github.com/grokify/gorod/issues)
