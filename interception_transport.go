package httpclientinterception

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type matchResult struct {
	totalRegistrations int
	matchers           regMatch
}

type regMatch struct {
	mt      map[int][]matchType
	success bool
}

type matchType struct {
	matched bool
	matcher
}

type interceptorTransport struct {
	RoundTripper http.RoundTripper
	config       []*configurationBuilder
	PanicOnMissingRegistration
	OnMissingRegistration
	result matchResult
	winner []winner
}

type winner struct {
	builder *configurationBuilder
	success bool
}

func (t *interceptorTransport) getResult() *winner {
	for _, w := range t.winner {
		if w.success {
			return &w
		}
	}

	return nil
}

func (t *interceptorTransport) RoundTrip(request *http.Request) (*http.Response, error) {

	t.winner = nil

	// TODO: Move into interceptorTransport as a method
	// Loop through registrations
	for _, builder := range t.config {

		// Loop through matchers
		winner := winner{success: true, builder: builder}
		for _, m := range builder.matchers {
			if !m.Match(request) {
				winner.success = false
			}
		}
		t.winner = append(t.winner, winner)

	}

	var response *http.Response

	result := t.getResult()
	if result != nil {
		response = &http.Response{
			StatusCode: result.builder.Status,
			Body:       ioutil.NopCloser(strings.NewReader(string(result.builder.Content))), // r type is io.ReadCloser
		}
	}

	if response == nil && t.OnMissingRegistration != nil {
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
