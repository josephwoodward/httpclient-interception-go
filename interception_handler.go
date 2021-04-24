package httpclientinterception

import (
	"net/http"
	"strconv"
)

type interceptionHandler struct {
	config                     *configurationBuilder
	PanicOnMissingRegistration PanicOnMissingRegistration
	OnMissingRegistration      OnMissingRegistration
}

func (h *interceptionHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	pass := true

	for _, m := range h.config.matchers {
		if !m.Match(request) {
			pass = false
		}
	}

	if pass == true {
		writer.Header().Set("Status-Code", strconv.Itoa(h.config.Status))
	}

	//if h.OnMissingRegistration != nil {
	//	response = h.OnMissingRegistration(request)
	//}

	if h.PanicOnMissingRegistration {
		panic("Missing registration")
	}
}
