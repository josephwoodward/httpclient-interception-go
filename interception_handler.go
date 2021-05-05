package httpclientinterception

import (
	"net/http"
	"strconv"
)

type interceptionHandler struct {
	config configurationBuilder
	PanicOnMissingRegistration
	OnMissingRegistration
}

func (o *interceptionHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	matched := true

	// TODO: Move into interception as a method
	for _, m := range o.config.matchers {
		if !m.Match(request) {
			matched = false
		}
	}

	if matched == true {
		writer.Header().Set("Status-Code", strconv.Itoa(o.config.Status))
	}

	if o.OnMissingRegistration != nil {
		_ = o.OnMissingRegistration(request)
	}

	if o.PanicOnMissingRegistration {
		panic("Missing registration")
	}
}
