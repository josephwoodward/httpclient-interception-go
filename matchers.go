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
