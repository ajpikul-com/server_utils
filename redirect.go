package sutils

import (
	"net/http"
	"strconv"
)

// redirectSchemeHandler holds the new scheme (usually HTTPS) and redirect code (30x) to use during redirection
type redirectSchemeHandler struct {
	scheme string
	code   int
}

// ServeHTTP rewrites the requests URL and appropriately and then calls Redirect.
func (rsh *redirectSchemeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Host = r.Host
	r.URL.Scheme = rsh.scheme
	http.Redirect(w, r, r.URL.String(), rsh.code)
}

// RedirectSchemeHandler returns a new http.Handler
func RedirectSchemeHandler(scheme string, code int) http.Handler {
	defaultLogger.Info("Creating new scheme redirect to " + scheme + " with code " + strconv.Itoa(code))
	return &redirectSchemeHandler{scheme, code}
}

type HostRewriteHandler struct {
	toHost     string
	code       int
	pathprefix string
}

// ServeHTTP takes regex to match a domain, replace found regex with rewrite, and adds path prefix to path
func (hrw *HostRewriteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Host = hrw.toHost
	if hrw.pathprefix != "" {
		if r.URL.Path[0] != '/' {
			r.URL.Path = hrw.pathprefix + "/ " + r.URL.Path
		} else {
			r.URL.Path = hrw.pathprefix + r.URL.Path
		}
	}
	defaultLogger.Info("New string: " + r.URL.String())
	http.Redirect(w, r, r.URL.String(), hrw.code)
}

// HostRewriteHandler returns a new http.Handler
func NewHostRewriteHandler(toHost string, pathprefix string, code int) http.Handler {
	return &HostRewriteHandler{toHost, code, pathprefix}
}

type RedirectHandler struct {
	Path string
	Code int
}

func (rh *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, rh.Path, rh.Code)
}

/*
if strings.HasPrefix(r.URL.Host, grupopikul.com)
if strings.HasPrefix(r.URL.Host, pikulgroup.com)
if strings.HasPrefix(r.URL.Host, www.pikulgroup.com)
*/
