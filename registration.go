package httpclientinterception

import "net/http"

// PanicOnMissingRegistration causes HttpClientInterception to panic if not registration is found
type PanicOnMissingRegistration bool

// OnMissingRegistration is invoked before the request is handled by the http.Client when a registration is missing
type OnMissingRegistration func(r *http.Request) *http.Response

type registrationOptions struct {
	interceptorBuilder *InterceptorBuilder
	builders           []*configurationBuilder
	PanicOnMissingRegistration
	OnMissingRegistration
}

func (o *registrationOptions) Client() *http.Client {
	transport := &interceptorTransport{
		RoundTripper:               http.DefaultTransport,
		PanicOnMissingRegistration: o.PanicOnMissingRegistration,
		OnMissingRegistration:      o.OnMissingRegistration,
		config:                     *o.builders[0],
	}
	return &http.Client{Transport: transport}
}

func (o *registrationOptions) Transport() http.RoundTripper {
	return http.DefaultTransport
}

// NewInterceptorOptions
func NewInterceptorOptions() *registrationOptions {
	return &registrationOptions{}
}
