package gorod

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/grokify/mogo/os/osutil"
)

type ForegroundBrowser struct {
	Launcher *launcher.Launcher
	Browser  *rod.Browser
}

func (fb *ForegroundBrowser) Close() {
	if fb.Browser != nil {
		fb.Browser.MustClose()
	}
	if fb.Launcher != nil {
		fb.Launcher.Cleanup()
	}
}

func NewForegroundBrowserPaused(navURL string, delaySeconds uint, paused bool) ForegroundBrowser {
	// Headless runs the browser on foreground, you can also use flag "-rod=show"
	// Devtools opens the tab in each new tab opened automatically
	l := launcher.New().
		Headless(false).
		Devtools(true)

	// defer l.Cleanup() // remove launcher.FlagUserDataDir

	url := l.MustLaunch()

	// Trace shows verbose debug information for each action executed
	// SlowMotion is a debug related function that waits 2 seconds between
	// each action, making it easier to inspect what your code is doing.
	browser := rod.New().
		ControlURL(url).
		Trace(true).
		SlowMotion(time.Duration(delaySeconds) * time.Second).
		MustConnect()

	// ServeMonitor plays screenshots of each tab. This feature is extremely
	// useful when debugging with headless mode.
	// You can also enable it with flag "-rod=monitor"
	launcher.Open(browser.ServeMonitor(""))

	// defer browser.MustClose()

	browser.MustPage(navURL)

	// page := browser.MustPage("https://ebay.com/")
	// page.MustElement("input").MustInput("git").MustType(input.Enter)
	// text := page.MustElement(".codesearch-results p").MustText()
	// fmt.Println(text)

	if paused {
		fmt.Println("Press the Enter Key after logging in!")
		var input string
		fmt.Scanln(&input)
	}

	return ForegroundBrowser{
		Launcher: l,
		Browser:  browser}
}

func (fb *ForegroundBrowser) GetWriteFileHTML(url, filename string, perm os.FileMode, force bool, writeDelay time.Duration) error {
	if !force && osutil.MustFileSize(filename) > 0 {
		return nil
	}
	page := fb.Browser.MustPage(url)
	ht, err := page.HTML()
	if err != nil {
		return err
	}
	time.Sleep(writeDelay)
	return os.WriteFile(filename, []byte(ht), perm)
}

// GetWriteFileMulti stores HTML and PNG screenshots.
func (fb *ForegroundBrowser) GetWriteFileMulti(srcURL, filenameHTML string, perm os.FileMode, force bool, opts *proto.PageCaptureScreenshot) error {
	filenameSS := filenameHTML + ".png"
	existsSS, err := osutil.Exists(filenameHTML)
	var page *rod.Page
	if force || err != nil || !existsSS {
		page, err = RetrieveWriteScreenshotFullPage(fb.Browser, srcURL, filenameSS, opts)
		if err != nil {
			return err
		}
	} else {
		page = fb.Browser.MustPage(srcURL)
	}
	defer page.Close()
	existsHTML, err := osutil.Exists(filenameHTML)
	if !force && err == nil && existsHTML {
		return nil
	}
	ht, err := page.HTML()
	if err != nil {
		return err
	}
	return os.WriteFile(filenameHTML, []byte(ht), perm)
}
