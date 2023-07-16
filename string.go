package sutils

import (
	"net/http"
)

type StringHandler struct {
	Val string
}

func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(sh.Val))
	if err != nil {
		panic(err)
	}
	w.(http.Flusher).Flush()
}
