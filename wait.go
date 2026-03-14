package gorod

import (
	"errors"
	"time"

	"github.com/go-rod/rod"
)

// WaitVisible waits for the element matching selector to become visible.
func WaitVisible(page *rod.Page, selector string, timeout time.Duration) (*rod.Element, error) {
	deadline := time.Now().Add(timeout)

	for {
		el, err := page.Element(selector)
		if err == nil && el != nil {
			visible, _ := el.Visible()
			if visible {
				return el, nil
			}
		}

		if time.Now().After(deadline) {
			return nil, errors.New("timeout waiting for element to be visible: " + selector)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// WaitClickable waits for the element to be both visible and enabled.
func WaitClickable(page *rod.Page, selector string, timeout time.Duration) (*rod.Element, error) {
	deadline := time.Now().Add(timeout)

	for {
		el, err := page.Element(selector)
		if err == nil && el != nil {
			visible, _ := el.Visible()
			disabled, _ := el.Attribute("disabled")
			if visible && disabled == nil {
				return el, nil
			}
		}

		if time.Now().After(deadline) {
			return nil, errors.New("timeout waiting for element to be clickable: " + selector)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

/*
// WaitClick waits for the element to be clickable, then clicks it.
func WaitClick(page *rod.Page, selector string, timeout time.Duration) error {
	el, err := WaitClickable(page, selector, timeout)
	if err != nil {
		return err
	}
	return el.Click()
}
*/
