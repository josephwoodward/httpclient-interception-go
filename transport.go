package httpclientinterception

import "net/http"

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
