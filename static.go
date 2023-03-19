package middleware

import (
	"bytes"
	"github.com/boggydigital/nod"
	"io"
	"net/http"
)

var staticContent = make(map[string][]byte)

func getStaticContent(w http.ResponseWriter, r *http.Request) bool {
	key := r.URL.Path
	if r.URL.RawQuery != "" {
		key += "?" + r.URL.RawQuery
	}
	if bs, ok := staticContent[key]; ok {
		br := bytes.NewReader(bs)
		if _, err := io.Copy(w, br); err != nil {
			http.Error(w, nod.Error(err).Error(), http.StatusInternalServerError)
			return false
		}
		return true
	}
	return false
}

func SetStaticContent(path string, content []byte) {
	staticContent[path] = content
}

func Static(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if getStaticContent(w, r) {
			nod.Log("serving static content: %s", r.URL.String())
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
