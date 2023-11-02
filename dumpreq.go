package sutils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kr/pretty"
)

type Dumper struct{}

func (d *Dumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	output := FormatRequest(r)
	output2 := fmt.Sprintf("%# v", pretty.Formatter(r))

	w.Write([]byte(output))
	w.Write([]byte("\n\n\n\n"))
	w.Write([]byte(output2))
	w.(http.Flusher).Flush()
}

// FormatRequest generates ascii representation of a request
func FormatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
