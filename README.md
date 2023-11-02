# Utilities

#### `FormatRequest(r *http.Request) string`

Dumps Request into string

```go
dumper := &sutils.Dumper{}
```

`Dumper` has a `ServeHTTP` which just dumps everything, including `http.Request`.

#### `HandlerItself func(http.ResponseWriter, *http.Request)`

Convert a `func(http.ResponseWriter, *http.Request)` to `HandlerItself` to so that it becomes a self-calling `http.Handler`

#### `RedirectSchemeHandler(scheme string, code int) http.Handler`

Tells requests to always use correct `scheme`.

```go
// Only 300 error codes work as expected, browser thing
NewHostRewriteHandler(desired_host, path_prefix_to_add, code_to_return)

&RedirectHandler{Path: "", Code: http.StatusMovedPermanently}
// Does exacetly what you think
```

#### `NewSingleHostReverseProxy(target string) (*ReverseProxy, err)`

Redirects everything behind-the-scenes to `target`

#### `StringHandler struct`
```go
type StringHandler struct {
	Val string
}
```

Will serve `Val` as a webpage
