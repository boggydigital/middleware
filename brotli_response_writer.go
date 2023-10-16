package middleware

import (
	"github.com/andybalholm/brotli"
	"io"
	"net/http"
	"sync"
)

var brPool = sync.Pool{
	New: func() interface{} {
		w := brotli.NewWriter(io.Discard)
		return w
	},
}

type brotliResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *brotliResponseWriter) WriteHeader(status int) {
	w.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(status)
}

func (w *brotliResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
