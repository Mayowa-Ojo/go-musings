package middleware

import (
	"log"
	"net/http"
)

// MethodOverride -
func MethodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			method := r.FormValue("_method")
			log.Println(method)
			switch method {
			case "PUT", "PATCH", "DELETE":
				r.Method = method
			}
		}

		next.ServeHTTP(w, r)
	})
}
