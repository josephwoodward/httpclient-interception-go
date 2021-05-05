# HttpClient Interception for Go

### Introduction

Based off of the [C# version](https://github.com/justeat/httpclient-interception), HttpClientInterception is a tool to help test HTTP services without the need to start up an HTTP Server.

Note: This project is still under development.

### Example

```go

import (
    . "httpclient-interception"
    ...
)


func Test_ExampleMatch(t *testing.T) {

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

```

HttpClientInterception can also be used to create an `http.Handler` for servers:

```go
func Test_HttpServer(t *testing.T) {
	opts := NewInterceptorOptions()
	builder := NewInterceptorBuilder(
		ForGet(),
		ForPath("/test"),
		RespondWithStatus(http.StatusOK))

	builder.RegisterOptions(opts)

	// Create a `http.Handler` for our test server
	srv := httptest.NewServer(opts.Handler())
	defer srv.Close()
	client := srv.Client()

	got, _ := client.Get(srv.URL + "/test")

    want := http.StatusOK	
	if got.StatusCode != want {
		t.Errorf("wanted: %v, got: %v", want, got.StatusCode)
	}
}
```
