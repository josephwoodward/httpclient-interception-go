package httpclientinterception_test

import (
	. "httpclient-interception"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_Headers(t *testing.T) {
	t.Skip()

	// Arrange
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForHeaders("Content-Type", "application/json"),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	path, _ := url.Parse("/test/")
	request := &http.Request{URL: path}
	request.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(request)

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.Status)
	}

}

func Test_Host(t *testing.T) {

	// Arrange
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForHost("public.je-apis.com"),
		ForPath("/terms"),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	response, _ := client.Get("https://public.je-apis.com/terms")

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.Status)
	}

}

func Test_Port(t *testing.T) {

	// Arrange
	path := "http://tester.com:8080/test/"
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForPort("8080"),
		ForPath("/test/"),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	response, _ := client.Get(path)

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.Status)
	}

}

func Test_Path(t *testing.T) {

	// Arrange
	path := "/test/"
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForPath(path),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	response, _ := client.Get(path)

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.Status)
	}

}

func Test_MethodPut(t *testing.T) {

	// Arrange
	path := "/test/"
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForPut(),
		ForPath(path),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	url, _ := url.Parse(path)
	response, _ := client.Do(&http.Request{Method: http.MethodPut, URL: url})

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.Status)
	}
}

func Test_MethodPatch(t *testing.T) {

	// Arrange
	path := "/test/"
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForPatch(),
		ForPath(path),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	url, _ := url.Parse(path)
	response, _ := client.Do(&http.Request{Method: http.MethodPatch, URL: url})

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.Status)
	}

}

func Test_MethodDelete(t *testing.T) {

	// Arrange
	path := "/test/"
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForDelete(),
		ForPath(path),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	url, _ := url.Parse(path)
	response, _ := client.Do(&http.Request{Method: http.MethodDelete, URL: url})

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.Status)
	}

}

func Test_MissingRegistration(t *testing.T) {

	// Arrange
	opts := NewInterceptorOptions()
	builder := NewInterceptorBuilder(ForGet(), ForPath("/test"), RespondWithStatus(http.StatusInternalServerError))

	opts.OnMissingRegistration = func(r *http.Request) *http.Response {
		return &http.Response{StatusCode: http.StatusTeapot}
	}

	builder.RegisterOptions(opts)
	client := opts.Client()

	// Act
	response, _ := client.Get("http://localhost.com/test2")

	wanted := http.StatusTeapot
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.StatusCode)
	}
}

func Test_MethodPost(t *testing.T) {

	// Arrange
	path := "/test/"
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForPost(),
		ForPath(path),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	response, _ := client.Post(path, "application/json", ioutil.NopCloser(nil))

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.Status)
	}

}

func Test_MethodGet(t *testing.T) {

	// Arrange
	path := "/test/"
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForGet(),
		ForPath(path),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	response, _ := client.Get(path)

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status %v, got %v", wanted, response.Status)
	}

}

func Test_Https(t *testing.T) {

	// Arrange
	opts := NewInterceptorOptions()
	opts.PanicOnMissingRegistration = true

	builder := NewInterceptorBuilder(
		ForGet(),
		ForHttps(),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	path := "https://test.com"
	response, _ := client.Get(path)

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status %v, got %v", wanted, response.Status)
	}

}

func Test_Http(t *testing.T) {

	// Arrange
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForGet(),
		ForHttp(),
		RespondWithStatus(http.StatusInternalServerError))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	path := "https://test.com"
	response, _ := client.Get(path)

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status %v, got %v", wanted, response.Status)
	}

}

func Test_MethodLiteral(t *testing.T) {

	// Arrange
	path := "/test/"

	tests := []struct {
		want string
	}{
		{http.MethodPost},
		{http.MethodDelete},
		{http.MethodPatch},
		{http.MethodPut},
		{http.MethodPost},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {

			// Arrange
			opts := NewInterceptorOptions()
			builder := NewInterceptorBuilder(
				ForMethod(tt.want),
				ForPath(path),
				RespondWithStatus(http.StatusOK))

			builder.RegisterOptions(opts)

			client := opts.Client()

			// Act
			url, _ := url.Parse(path)
			response, _ := client.Do(&http.Request{Method: tt.want, URL: url})

			// Assert
			wanted := http.StatusOK
			if response.StatusCode != wanted {
				t.Errorf("Wanted status %v, got %v", wanted, response.Status)
			}
		})
	}

}

func Test_HttpServer(t *testing.T) {
	opts := NewInterceptorOptions()
	builder := NewInterceptorBuilder(
		ForGet(),
		ForPath("/test"),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	want := http.StatusOK
	srv := httptest.NewServer(opts.Handler())
	defer srv.Close()
	client := srv.Client()

	got, _ := client.Get(srv.URL + "/test")

	if want != got.StatusCode {
		t.Errorf("wanted: %v, got: %v", want, got.Status)
	}
}
