There are only two really useful functions in here:

`RedirectSchemeHandler(scheme string, code int) http.Handler`

Which will respond to any request at the path it's assigned to with a response code and the same exactly URL except with the scheme specified (ie changing HTTP to HTTPS)

`NewSingleHostReverseProxy(target string)` which will change whatever path it's put on to target, but really without the user ever seeing, always passing through the rest of the URL.
