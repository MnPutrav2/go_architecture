package method

import (
	"fmt"
	"net/http"

	"github.com/MnPutrav2/go_architecture/util"
)

func GET(mux *http.ServeMux, pattern string, next http.HandlerFunc, middle ...util.Middleware) {
	pt := fmt.Sprintf("GET %s", pattern)

	for _, handle := range middle {
		next = handle(next)
	}

	mux.HandleFunc(pt, next)
}

func DELETE(mux *http.ServeMux, pattern string, next http.HandlerFunc, middle ...util.Middleware) {
	pt := fmt.Sprintf("DELETE %s", pattern)

	for _, handle := range middle {
		next = handle(next)
	}

	mux.HandleFunc(pt, next)
}

func POST(mux *http.ServeMux, pattern string, next http.HandlerFunc, middle ...util.Middleware) {
	pt := fmt.Sprintf("POST %s", pattern)

	for _, handle := range middle {
		next = handle(next)
	}

	mux.HandleFunc(pt, next)
}
