package server_utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ReverseProxy is a wrapper for httputil's reverse proxy that just does the tedious working of parsing a url string. It redirects one server to another.
type ReverseProxy struct {
	http.Handler
}

// NewSingleHostReverseProxy is the constructor for the ReverseProxy struct that actually does the work. I use it when I'm rnning a 3rd party server behind a firewall on my box.
func NewSingleHostReverseProxy(target string) (*ReverseProxy, error) {
	defaultLogger.Info("Creating new single host reverse proxy to " + target)
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return &ReverseProxy{
		Handler: httputil.NewSingleHostReverseProxy(targetURL),
	}, nil
}
