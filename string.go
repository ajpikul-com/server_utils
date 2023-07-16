package server_utils

import (
	"net/http"
)

type StringHandler struct {
	val string
}

func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(sh.val))
	if err != nil {
		panic(err)
	}
	w.(http.Flusher).Flush()
}
