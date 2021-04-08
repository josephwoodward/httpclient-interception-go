package httpclientinterception

import "net/http"

type matcher interface {
	Match(*http.Request) bool
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

// hostMatcher matches the request again the host name.
type hostMatcher string

func (host hostMatcher) Match(r *http.Request) bool {
	return r.URL.Host == string(host)
}

// headersMatcher
type headersMatcher map[string][]string

func (headers headersMatcher) Match(request *http.Request) bool {
	if len(request.Header) == 0 {
		return false
	}

	for headerKey, _ := range headers {
		rHeaderValues := request.Header.Values(headerKey)
		if len(rHeaderValues) == 0 {
			return false
		}

		for _, value := range headers[headerKey] {
			for i := 0; i < len(rHeaderValues); i++ {
				if value != rHeaderValues[i] {
					return false
				}
			}
		}
	}
	return true
}
