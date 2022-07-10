package middleware

import (
	"fmt"
	"github.com/boggydigital/nod"
	"net/http"
)

func httpMethodOnly(httpMethod string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != httpMethod {
			err := fmt.Errorf("expected %s got %s", httpMethod, r.Method)
			http.Error(w, nod.Error(err).Error(), http.StatusMethodNotAllowed)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func GetMethodOnly(next http.Handler) http.Handler {
	return httpMethodOnly(http.MethodGet, next)
}

func PutMethodOnly(next http.Handler) http.Handler {
	return httpMethodOnly(http.MethodPut, next)
}

func DeleteMethodOnly(next http.Handler) http.Handler {
	return httpMethodOnly(http.MethodDelete, next)
}
