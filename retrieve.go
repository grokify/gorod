package gorod

import (
	"fmt"
	"net/http"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/grokify/mogo/net/http/httputilmore"
)

/*

Rod notes:

https://github.com/go-rod/rod/issues/32 - full poage screenshot

https://github.com/go-rod/rod/issues/213 - respons headers

*/

/*

open the login window page := browser.Page("http://login.com")
change the size of the window with page.Window(0,0, width_only_for_captcha, height_only_for_captcha)
scroll to the captcha el := page.Element("/xpath/to/captcha").ScrollIntoView() to
wait for the user to finish the captcha el.WaitInvisible()

*/

func RetrieveWriteScreenshotFullPage(browser *rod.Browser, srcURL, filename string, opts *proto.PageCaptureScreenshot) (*rod.Page, error) {
	if browser == nil {
		browser = rod.New().MustConnect()
	}
	page := browser.MustPage("")

	var e proto.NetworkResponseReceived

	wait := page.WaitEvent(&e)
	page.MustNavigate(srcURL)
	wait()

	if e.Response.Status == http.StatusNotFound {
		return page, httputilmore.ErrStatus404
	} else if e.Response.Status != http.StatusOK {
		return page, fmt.Errorf("status not 200 [%d]", e.Response.Status)
	}

	page.MustWaitLoad()

	img, err := page.Screenshot(true, nil)
	if err != nil {
		return page, err
	}
	err = utils.OutputFile(filename, img)
	return page, err
}
