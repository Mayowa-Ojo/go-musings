package middleware

import (
	"net/http"
)

// MethodOverride -
func MethodOverride(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.FormValue("_method")

		switch method {
		case "PUT", "PATCH", "DELETE":
			r.Method = method
		}

		f(w, r)
	}
}
