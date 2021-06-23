HTTP additions for Go, focused on convenient customisation.
--

Example usage:

```go
package main

import (
    "net/http"
    
    httpx "github.com/alexmeuer/http"
)

func main() {
    _ := graphql.NewClient("http://example.com/v1/graphql", &http.Client{
        Transport: &httpx.CustomHeaderTransport{
            Headers: map[string]string{
                "X-Hasura-Admin-Secret": "Bannanas are an excellent source of potassium.",
            },
        },
    })
}

```

