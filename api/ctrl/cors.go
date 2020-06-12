package ctrl

import (
	"net/http"
	"strings"
)

// cors enabling middleware
func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods",
			strings.Join([]string{
				http.MethodPost,
				http.MethodGet,
				http.MethodOptions,
				http.MethodPut,
				http.MethodDelete,
			}, ","))
		w.Header().Set("Access-Control-Allow-Headers",
			strings.Join([]string{
				"Accept",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"X-CSRF-Token",
				"Authorization",
			}, ","))
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
