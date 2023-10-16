package middleware

import (
	"compress/gzip"
	"github.com/andybalholm/brotli"
	"io"
)

type enc int

const (
	no enc = iota
	br
	gz
)

func (e enc) String() string {
	switch e {
	case br:
		return "br"
	case gz:
		return "gzip"
	case no:
		fallthrough
	default:
		return ""
	}
}

func (e enc) GetPool() io.WriteCloser {
	switch e {
	case br:
		return brPool.Get().(*brotli.Writer)
	case gz:
		return gzPool.Get().(*gzip.Writer)
	default:
		return nil
	}
}

func (e enc) Put(wc io.WriteCloser) {
	switch e {
	case br:
		brPool.Put(wc)
	case gz:
		gzPool.Put(wc)
	default:
		// do nothing
	}
}

func (e enc) Reset(wc io.WriteCloser, w io.Writer) {
	switch e {
	case br:
		wc.(*brotli.Writer).Reset(w)
	case gz:
		wc.(*gzip.Writer).Reset(w)
	default:
		// do nothing
	}
}
