package httpclientinterception

import (
	"net/http"
)

type matcher interface {
	Match(*http.Request) bool
}

type matchMetadata struct {
}

// methodMatcher matches the request against the method value.
type methodMatcher string

func (method methodMatcher) Match(r *http.Request) bool {
	return r.Method == string(method)
}

// pathMatcher matches the request against a path value.
type pathMatcher string

func (path pathMatcher) Match(r *http.Request) bool {
	return r.URL.Path == string(path)
}

// queryStringMatcher matches the request against a query string.
type queryStringMatcher string

func (query queryStringMatcher) Match(r *http.Request) bool {
	//// TODO: Does order of query values matter?
	//// r should exist in `query`
	//v, err := url.ParseQuery(string(query))
	//if err != nil {
	//	return true
	//}
	//
	//rQueryString := r.URL.Query()
	//
	//for key, _ := range rQueryString {
	//	if v != nil {
	//		fmt.Println(v)
	//	}
	//}
	//
	return true
}

// schemeMatcher matches the request against a path value.
type schemeMatcher string

func (scheme schemeMatcher) Match(r *http.Request) bool {
	return r.URL.Scheme == string(scheme)
}

// portMatcher matches the request against a path value.
type portMatcher string

func (port portMatcher) Match(r *http.Request) bool {
	return r.URL.Port() == string(port)
}

// hostMatcher matches the request against a host name.
type hostMatcher string

func (host hostMatcher) Match(r *http.Request) bool {
	if string(host) == "*" {
		return true
	}

	return r.URL.Host == string(host)
}

// headersMatcher
type headersMatcher map[string][]string

func (headers headersMatcher) Match(request *http.Request) bool {
	if len(request.Header) == 0 {
		return false
	}

	var matches = make(graph)
	for key, headerValues := range headers {
		requestHeaderValues := request.Header.Values(key)

		if len(headerValues) == 0 || len(requestHeaderValues) == 0 {
			return true
		}

		for _, value := range headerValues {
			for _, headerValue := range requestHeaderValues {
				m := something{key: key, namespace: key + "." + value, given: headerValue, expected: value}
				if value == headerValue {
					m.matched = true
					matches[key] = m
					break
				}

				m.matched = false
				matches[key] = m
			}
		}
	}

	return matches.isMatch()
}

// TODO: Not sure what to call this yet?
type something struct {
	key       string
	matched   bool
	expected  string
	given     string
	namespace string
}

type graph map[string]something

func (g graph) isMatch() bool {

	for _, s2 := range g {
		if s2.matched == false {
			return false
		}
	}

	return true
}
