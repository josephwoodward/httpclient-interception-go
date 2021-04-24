package httpclientinterception

import (
	"net/http"
)

// NewInterceptorBuilder creates a new interceptionBuilder that allows you to configure requests to match
func NewInterceptorBuilder(o ...BuilderOption) *interceptionBuilder {
	builder := &configurationBuilder{}
	for _, o := range o {
		o(builder)
	}

	return &interceptionBuilder{builder: builder}
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

func ForHost(host string) BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(hostMatcher(host))
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

func RespondWithStatus(statusCode int) BuilderOption {
	return func(b *configurationBuilder) {
		b.Status = statusCode
	}
}
