package httpclientinterception

import (
	"net/http"
	"net/url"
	"testing"
)

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
		name       string
		configured map[string][]string
		got        map[string][]string
		match      bool
	}{
		{
			"Exact match",
			map[string][]string{"Content-Type": {"application/json"}},
			map[string][]string{"Content-Type": {"application/json"}},
			true,
		},
		{
			"Single header vs many",
			map[string][]string{"Content-Type": {"application/json"}},
			map[string][]string{"Content-Type": {"application/json", "text/html"}},
			true,
		},
		{
			"Missing header value",
			map[string][]string{"Content-Type": {""}},
			map[string][]string{"Content-Type": {"application/json", "text/html"}},
			false,
		},
		{
			"Multiple headers",
			map[string][]string{"Content-Type": {"application/json", "text/html"}},
			map[string][]string{"Content-Type": {""}},
			false,
		},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {

			r := &http.Request{Header: test.got}
			if headersMatcher(test.configured).Match(r) != test.match {
				t.Errorf("Wanted header, but was not")
			}

		})
	}
}
