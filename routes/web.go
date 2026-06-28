package route

import "net/http"

func RouteWeb() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// register route in here

	})
}
