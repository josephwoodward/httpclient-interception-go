# HttpClient Interception for Go

### Introduction

Based off of the [C# version](https://github.com/justeat/httpclient-interception), HttpClientInterception is a tool to help test HTTP services without the need to start up an HTTP Server.

### Example

```go

import (
    . "httpclient-interception"
    ...
)


func ExampleMatch(t *testing.T) {

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
