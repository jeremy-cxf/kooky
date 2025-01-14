package kooky_test

import (
	"github.com/jeremy-cxf/kooky"
	_ "github.com/jeremy-cxf/kooky/browser/all" // register cookiestore finders
)

var cookieName = `NID`

func ExampleFilterCookies() {
	cookies := kooky.ReadCookies() // automatic read

	cookies = kooky.FilterCookies(
		cookies,
		kooky.Valid,                    // remove expired cookies
		kooky.DomainContains(`google`), // cookie domain has to contain "google"
		kooky.Name(cookieName),         // cookie name is "NID"
		kooky.Debug,                    // print cookies after applying previous filter
	)
}
