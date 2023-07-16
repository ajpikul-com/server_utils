package server_utils

import (
	"net/http"
)

type HandlerItself func(http.ResponseWriter, *http.Request)

func (h HandlerItself) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	(h)(w, r)
}
