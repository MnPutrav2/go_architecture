package method

import (
	"fmt"
	"net/http"

	"github.com/MnPutrav2/go_architecture/app/util"
)

func GET(mux *http.ServeMux, pattern string, next http.HandlerFunc, middle ...util.Middleware) {
	pt := fmt.Sprintf("GET %s", fmt.Sprintf("/api%s", pattern))

	for _, handle := range middle {
		next = handle(next)
	}

	mux.HandleFunc(pt, next)
}

func DELETE(mux *http.ServeMux, pattern string, next http.HandlerFunc, middle ...util.Middleware) {
	pt := fmt.Sprintf("DELETE %s", fmt.Sprintf("/api%s", pattern))

	for _, handle := range middle {
		next = handle(next)
	}

	mux.HandleFunc(pt, next)
}

func POST(mux *http.ServeMux, pattern string, next http.HandlerFunc, middle ...util.Middleware) {
	pt := fmt.Sprintf("POST %s", fmt.Sprintf("/api%s", pattern))

	for _, handle := range middle {
		next = handle(next)
	}

	mux.HandleFunc(pt, next)
}

func PUT(mux *http.ServeMux, pattern string, next http.HandlerFunc, middle ...util.Middleware) {
	pt := fmt.Sprintf("PATCH %s", fmt.Sprintf("/api%s", pattern))

	for _, handle := range middle {
		next = handle(next)
	}

	mux.HandleFunc(pt, next)
}

func PATCH(mux *http.ServeMux, pattern string, next http.HandlerFunc, middle ...util.Middleware) {
	pt := fmt.Sprintf("PATCH %s", fmt.Sprintf("/api%s", pattern))

	for _, handle := range middle {
		next = handle(next)
	}

	mux.HandleFunc(pt, next)
}
