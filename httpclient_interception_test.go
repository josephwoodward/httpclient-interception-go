package httpclientinterception_test

import (
	"encoding/json"
	. "httpclient-interception"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func Test_Headers(t *testing.T) {

	// Arrange
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		//TODO: Ordering is irrelevant?
		ForHeaders(map[string][]string{"Content-Type": {"text/html", "application/json"}}),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	path, _ := url.Parse("http://localhost/test/")
	request := &http.Request{URL: path}
	request.Header = http.Header{}

	request.Header = map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Content-Type":    {"application/json", "text/html"},
	}

	response, _ := client.Do(request)

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted to match headers but could not")
	}
}

func Test_ResponseBody(t *testing.T) {

	// Arrange
	b := &testBody{Page: 5, Name: "John Wick"}
	content, _ := json.Marshal(b)

	opts := NewInterceptorOptions()
	builder := NewInterceptorBuilder(
		ForHeaders(map[string][]string{"Content-Type": {"text/html", "application/json"}}),
		RespondsWithJsonContent(content),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	path, _ := url.Parse("http://localhost/test/")
	request := &http.Request{URL: path}
	request.Header = http.Header{}

	request.Header = map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Content-Type":    {"application/json", "text/html"},
	}

	r, _ := client.Do(request)

	// Assert
	wanted := http.StatusOK
	if r.StatusCode != wanted {
		t.Errorf("Wanted to match headers but could not")
	}

	var body []byte
	_, err := r.Body.Read(body)
	if err != nil {
		t.Fatalf("Error parsing response body")
	}
	if reflect.DeepEqual(r.Body, body) {
		t.Errorf("test")
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

func Test_Query(t *testing.T) {

	// Arrange
	path := "/test/"
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForPath(path),
		ForQuery("per_page=2"),
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
	builder := NewInterceptorBuilder(
		ForGet(),
		ForPath("/test"),
		RespondWithStatus(http.StatusInternalServerError))

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

func Test_AnyHost(t *testing.T) {

	// Arrange
	opts := NewInterceptorOptions()
	opts.PanicOnMissingRegistration = true

	builder := NewInterceptorBuilder(
		ForHttps(),
		ForAnyHost(),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	client := opts.Client()

	// Act
	response, _ := client.Get("https://example.com/test")

	// Assert
	wanted := http.StatusOK
	if response.StatusCode != wanted {
		t.Errorf("Wanted status: %v, got: %v", wanted, response.Status)
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

// TODO - Add the ability to scope configurations
func Test_ScopedInterceptions(t *testing.T) {

	t.Skip()

	// Arrange
	opts := NewInterceptorOptions()

	builder := NewInterceptorBuilder(
		ForGet(),
		ForHttp(),
		RespondWithStatus(http.StatusInternalServerError))

	builder.RegisterOptions(opts)

	cleanUp := opts.BeginScope()
	defer cleanUp()

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

func Test_MultipleRequestRegistrations(t *testing.T) {

	// Arrange
	opts := NewInterceptorOptions()

	builder1 := NewInterceptorBuilder(
		ForGet(),
		ForHttps(),
		ForPath("/builder1"),
		RespondWithStatus(http.StatusForbidden))

	builder2 := NewInterceptorBuilder(
		ForGet(),
		ForHttps(),
		ForPath("/builder2"),
		RespondWithStatus(http.StatusTeapot))

	builder1.RegisterOptions(opts)
	builder2.RegisterOptions(opts)

	client := opts.Client()

	// Act
	path := "https://test.com"
	_, _ = client.Get(path + "/noop1")
	response1, _ := client.Get(path + "/builder1")

	_, _ = client.Get(path + "/noop2")
	response2, _ := client.Get(path + "/builder2")

	// Assert
	wanted := http.StatusForbidden
	if response1.StatusCode != wanted {
		t.Errorf("Response 1 wanted status %v, got %v", wanted, response1.Status)
	}

	wanted = http.StatusTeapot
	if response2.StatusCode != wanted {
		t.Errorf("Response 2 wanted status %v, got %v", wanted, response2.Status)
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

type testBody struct {
	Page int    `json:"page"`
	Name string `json:"name"`
}
