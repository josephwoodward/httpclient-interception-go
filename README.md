# HttpClient Interception for Go

## Introduction

Based off of the [C# version](https://github.com/justeat/httpclient-interception), HttpClientInterception is a tool to help test HTTP services without the need to starting up a HTTP Server.

### Example

```go
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

```
