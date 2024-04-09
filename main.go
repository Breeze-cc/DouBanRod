package main

import (
	"douban/method"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	u := launcher.New().
		Set("window-size", "1600,900").
		Set("no-sandbox").
		Set("ignore-certificate-errors").
		Set("ignore-certificate-errors-spki-list").
		Set("ignore-ssl-errors").
		Headless(false)
	uURL := u.MustLaunch()
	page := rod.New().ControlURL(uURL).MustConnect().MustPage()
	method.GETNewMovies(page)
	//time.Sleep(time.Hour)
}
