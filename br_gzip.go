package middleware

// Original source: https://gist.github.com/CJEnright/bc2d8b8dc0c1389a9feeddb110f822d7
// License: MIT

import (
	"net/http"
	"strings"
)

func BrGzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		en := no

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "br") {
			if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				next.ServeHTTP(w, r)
				return
			} else {
				en = gz
			}
		} else {
			en = br
		}

		w.Header().Set("Content-Encoding", en.String())

		wc := en.GetPool()
		defer en.Put(wc)

		en.Reset(wc, w)
		defer wc.Close()

		next.ServeHTTP(&gzipResponseWriter{ResponseWriter: w, Writer: wc}, r)
	})
}
