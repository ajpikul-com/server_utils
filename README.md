# Utilities

#### `FormatRequest(r *http.Request) string`

Dumps Request

#### `HandlerItself func(http.ResponseWriter, *http.Request)`

Convert a `func(http.ResponseWriter, *http.Request)` to `HandlerItself` to so that it becomes a self-calling `http.Handler`

#### `RedirectSchemeHandler(scheme string, code int) http.Handler`

Tells requests to always use correct `scheme`.

#### `NewSingleHostReverseProxy(target string) (*ReverseProxy, err)`

Redirects everything behind-the-scenes to `target`

#### `StringHandler struct`
```go
type StringHandler struct {
	Val string
}
```

Will serve `Val` as a webpage
