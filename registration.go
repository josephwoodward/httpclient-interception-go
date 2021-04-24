package httpclientinterception

import "net/http"

type interceptionOptions struct {
	interceptorBuilder *interceptionBuilder
	builders           []*configurationBuilder
	PanicOnMissingRegistration
	OnMissingRegistration
}

// NewInterceptorOptions creates a new interceptionOptions object used to configure your interceptor
func NewInterceptorOptions() *interceptionOptions {
	return &interceptionOptions{}
}

// PanicOnMissingRegistration causes HttpClientInterception to panic if not registration is found
type PanicOnMissingRegistration bool

// OnMissingRegistration is invoked before the request is handled by the http.Client when a registration is missing
type OnMissingRegistration func(r *http.Request) *http.Response

func (o *interceptionOptions) Client() *http.Client {
	transport := &interceptorTransport{
		RoundTripper:               http.DefaultTransport,
		PanicOnMissingRegistration: o.PanicOnMissingRegistration,
		OnMissingRegistration:      o.OnMissingRegistration,
		config:                     *o.builders[0],
	}
	return &http.Client{Transport: transport}
}

func (o *interceptionOptions) Transport() http.RoundTripper {
	return http.DefaultTransport
}

func (o *interceptionOptions) Handler() http.Handler {
	return &interceptionHandler{
		config:                     o.builders[0],
		PanicOnMissingRegistration: o.PanicOnMissingRegistration,
		OnMissingRegistration:      o.OnMissingRegistration,
	}
}
