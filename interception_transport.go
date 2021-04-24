package httpclientinterception

import "net/http"

type interceptorTransport struct {
	RoundTripper http.RoundTripper
	config       configurationBuilder
	PanicOnMissingRegistration
	OnMissingRegistration
}

func (t *interceptorTransport) RoundTrip(request *http.Request) (*http.Response, error) {

	pass := true

	for _, m := range t.config.matchers {
		if !m.Match(request) {
			pass = false
		}
	}

	var response *http.Response
	if pass == true {
		response = &http.Response{StatusCode: t.config.Status}
	}

	if t.OnMissingRegistration != nil {
		response = t.OnMissingRegistration(request)
	}

	if response != nil {
		return response, nil
	}

	if t.PanicOnMissingRegistration {
		panic("Missing registration")
	}

	return t.RoundTripper.RoundTrip(request)
}
