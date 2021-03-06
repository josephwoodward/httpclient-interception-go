package httpclientinterception

import (
	"net/http"
	"net/url"
	"testing"
)

func Test_MatchQuery(t *testing.T) {

	tt := []struct {
		want string
	}{
		{"?test=1&test2=abc"},
	}

	for _, test := range tt {
		url, _ := url.Parse("test.com/" + test.want)
		r := &http.Request{URL: url}
		if queryStringMatcher(test.want).Match(r) != true {
			t.Errorf("Wanted: %v, but was not", test.want)
		}
	}
}

func Test_MatchMethod(t *testing.T) {

	tt := []struct {
		want string
	}{
		{http.MethodPost},
		{http.MethodHead},
		{http.MethodDelete},
		{http.MethodPatch},
		{http.MethodPut},
		{http.MethodPost},
	}

	for _, test := range tt {
		r := &http.Request{Method: test.want}
		if methodMatcher(test.want).Match(r) != true {
			t.Errorf("Wanted: %v, but was not", test.want)
		}
	}

}

func Test_MatchPath(t *testing.T) {

	var tt = []struct {
		want string
	}{
		{"/hello/"},
		{"/hello"},
		{"/hello.html"},
		{"hello.html"},
	}
	for _, test := range tt {
		r := &http.Request{URL: &url.URL{Path: test.want}}
		if pathMatcher(test.want).Match(r) != true {
			t.Errorf("Wanted: %v, but was not", test.want)
		}
	}
}

func Test_MatchHeaders(t *testing.T) {

	var tt = []struct {
		name     string
		expected map[string][]string
		got      map[string][]string
		match    bool
	}{
		{
			"Exact match should pass",
			map[string][]string{"Content-Type": {"application/json"}},
			map[string][]string{"Content-Type": {"application/json"}},
			true,
		},
		{
			"Single header vs many should pass",
			map[string][]string{"Content-Type": {"application/json"}},
			map[string][]string{"Content-Type": {"application/json", "text/html"}},
			true,
		},
		{
			"Missing header value should fail",
			map[string][]string{"Content-Type": {""}},
			map[string][]string{"Content-Type": {"application/json", "text/html"}},
			false,
		},
		{
			"Multiple headers vs none should fail",
			map[string][]string{"Content-Type": {"application/json", "text/html"}},
			map[string][]string{"Content-Type": {""}},
			false,
		},
		{
			"Multiple headers vs none should pass",
			map[string][]string{"Referer": {"https://developer.mozilla.org/testpage.html"}},
			map[string][]string{"Referer": {"https://developer.mozilla.org/testpage.html"}, "Content-Type": {"application/json", "text/html"}, "Connection": {"keep-alive"}},
			true,
		},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {

			r := &http.Request{Header: test.got}
			if headersMatcher(test.expected).Match(r) != test.match {
				t.Errorf("Wanted header, but was not")
			}

		})
	}
}
