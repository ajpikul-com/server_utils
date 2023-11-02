package sutils

import (
	"net/http"
	"regexp"
	"strconv"
)

// redirectSchemeHandler holds the new scheme (usually HTTPS) and redirect code (30x) to use during redirection
type redirectSchemeHandler struct {
	scheme string
	code   int
}

// ServeHTTP rewrites the requests URL and appropriately and then calls Redirect.
func (rsh *redirectSchemeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	newURL := r.URL
	if r.URL.IsAbs() == false {
		r.URL.Host = r.Host
	}
	newURL.Scheme = rsh.scheme
	http.Redirect(w, r, newURL.String(), rsh.code)
}

// RedirectSchemeHandler returns a new http.Handler
func RedirectSchemeHandler(scheme string, code int) http.Handler {
	defaultLogger.Info("Creating new scheme redirect to " + scheme + " with code " + strconv.Itoa(code))
	return &redirectSchemeHandler{scheme, code}
}

type urlRewriteHandler struct {
	prefix   string
	rewrite  string
	code     int
	compiled *regexp.Regexp
	handler  http.Handler
}

// ServeHTTP rewrites the requests URL and appropriately and then calls Redirect.
func (urw *urlRewriteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	newURL := r.URL
	newURL.Host = urw.compiled.ReplaceAllString(r.URL.Host, urw.rewrite)
	if newURL.Host != r.URL.Host {
		http.Redirect(w, r, newURL.String(), urw.code)
		return
	}
	urw.handler.ServeHTTP(w, r)

}

// URLRewriteHandler returns a new http.Handler
func URLRewriteHandler(prefix string, rewrite string, code int, handler http.Handler) http.Handler {
	compiled := regexp.MustCompile(`(?i)^` + prefix)
	return &urlRewriteHandler{prefix, rewrite, code, compiled, handler}
}

/*
if strings.HasPrefix(r.URL.Host, grupopikul.com)
if strings.HasPrefix(r.URL.Host, pikulgroup.com)
if strings.HasPrefix(r.URL.Host, www.pikulgroup.com)
*/
