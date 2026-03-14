package gorod

import (
	"net/http"

	"github.com/go-rod/rod/lib/proto"
	"github.com/grokify/mogo/net/http/httputilmore"
)

type Cookies []*proto.NetworkCookie

func (c Cookies) HTTPCookies() []*http.Cookie {
	httpCookies := []*http.Cookie{}
	for _, rodCookie := range c {
		if rodCookie == nil {
			continue
		}
		httpCookies = append(httpCookies, &http.Cookie{
			Name:  rodCookie.Name,
			Value: rodCookie.Value,
		})
	}
	return httpCookies
}

func (c Cookies) String() string {
	return httputilmore.Cookies(c.HTTPCookies()).String()
}
