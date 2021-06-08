package httpclientinterception

import (
	"fmt"
	"net/http"
	"strings"
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
	err      error
	matchers []matcher
	Content  []byte
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

func ForQuery(query string) BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(pathMatcher(query))
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

func ForAnyHost() BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(hostMatcher("*"))
	}
}

func ForHeaders(headers map[string][]string) BuilderOption {
	return func(b *configurationBuilder) {
		b.addMatcher(headersMatcher(headers))
	}
}

//func ForHeader(pairs ...string) BuilderOption {
//	return func(b *configurationBuilder) {
//		m, err := mapFromPairsToString(pairs...)
//		b.err = err
//		b.addMatcher(headersMatcher(m))
//	}
//}

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

func RespondsWithJsonContent(content []byte) BuilderOption {
	return func(b *configurationBuilder) {
		b.Content = content
	}
}

func RespondWithStatus(statusCode int) BuilderOption {
	return func(b *configurationBuilder) {
		b.Status = statusCode
	}
}

// checkPairs returns the count of strings passed in, and an error if
// the count is not an even number.
func checkPairs(pairs ...string) (int, error) {
	length := len(pairs)
	if length%2 != 0 {
		return length, fmt.Errorf(
			"mux: number of parameters must be multiple of 2, got %v", pairs)
	}
	return length, nil
}

// mapFromPairsToString converts variadic string parameters to a
// string to string map.
func mapFromPairsToString(pairs ...string) (map[string][]string, error) {
	length, err := checkPairs(pairs...)
	if err != nil {
		return nil, err
	}
	m := make(map[string][]string, length/2)
	for i := 0; i < length; i += 2 {
		p := pairs[i+1]
		s := strings.Split(p, ",")

		fmt.Println(p)
		m[pairs[i]] = s
	}
	return m, nil
}
