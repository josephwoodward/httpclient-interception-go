package httpclientinterception

import (
	"net/http"
)

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

// NewInterceptorBuilder creates a new Interceptor Builder that allows you to configure requests to match.
func NewInterceptorBuilder(o ...BuilderOption) *InterceptorBuilder {
	builder := &configurationBuilder{}
	for _, o := range o {
		o(builder)
	}

	return &InterceptorBuilder{builder: builder}
}

type configurationBuilder struct {
	Status   int
	Host     string
	Headers  []string
	err      error
	matchers []matcher
}

func (c *configurationBuilder) addMatcher(m matcher) {
	c.matchers = append(c.matchers, m)
}

type BuilderOption func(b *configurationBuilder)

// ForGet matches against GET requests
func ForGet() BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(methodMatcher(http.MethodGet))
	}
}

// ForPost matches against POST requests
func ForPost() BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(methodMatcher(http.MethodPost))
	}
}

// ForPut matches against PUT requests
func ForPut() BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(methodMatcher(http.MethodPut))
	}
}

// ForDelete matches against DELETE requests
func ForDelete() BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(methodMatcher(http.MethodDelete))
	}
}

// ForPatch matches against PATCH requests
func ForPatch() BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(methodMatcher(http.MethodPatch))
	}
}

// ForMethod allows you to pass a
func ForMethod(method string) BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(methodMatcher(method))
	}
}

func ForPath(path string) BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(pathMatcher(path))
	}
}

func ForPort(port string) BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(portMatcher(port))
	}
}

func RespondWithStatus(statusCode int) BuilderOption {
	return func(b *configurationBuilder) {
		b.Status = statusCode
	}
}

func ForHost(host string) BuilderOption {
	return func(b *configurationBuilder) {
		b.Host = host
	}
}

func ForHeaders(headers ...string) BuilderOption {
	return func(b *configurationBuilder) {
		b.Headers = headers
	}
}

func ForHttps() BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(schemeMatcher("https"))
	}
}

func ForHttp() BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(schemeMatcher("http"))
	}
}

type interceptorTransport struct {
	RoundTripper http.RoundTripper
	config       configurationBuilder
	PanicOnMissingRegistration
	OnMissingRegistration
}

func (i *interceptorTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	pass := true

	for _, m := range i.config.matchers {
		if !m.Match(req) {
			pass = false
		}
	}

	var response *http.Response
	if pass == true {
		response = &http.Response{StatusCode: i.config.Status}
	}

	if i.OnMissingRegistration != nil {
		response = i.OnMissingRegistration(req)
	}

	if response != nil {
		return response, nil
	}

	if i.PanicOnMissingRegistration {
		panic("Missing registration")
	}

	return i.RoundTripper.RoundTrip(req)
}
