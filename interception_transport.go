package httpclientinterception

import (
	"net/http"
)

type interceptorTransport struct {
	RoundTripper http.RoundTripper
	config       configurationBuilder
	PanicOnMissingRegistration
	OnMissingRegistration
}

func (o *interceptorTransport) RoundTrip(request *http.Request) (*http.Response, error) {

	matched := true

	// TODO: Move into interceptorTransport as a method
	for _, m := range o.config.matchers {
		if !m.Match(request) {
			matched = false
		}
	}

	var response *http.Response

	if matched == true {
		response = &http.Response{StatusCode: o.config.Status}
	}

	if response == nil && o.OnMissingRegistration != nil {
		response = o.OnMissingRegistration(request)
	}

	if response != nil {
		return response, nil
	}

	if o.PanicOnMissingRegistration {
		panic("Missing registration")
	}

	return o.RoundTripper.RoundTrip(request)
}
