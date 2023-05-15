package middleware

import (
	"net/http"
	"time"
)

const (
	LastModifiedHeader    = "Last-Modified"
	IfModifiedSinceHeader = "If-Modified-Since"
)

func IsNotModified(ifModifiedSince, lastModified string) bool {
	//utcSince := time.Unix(since, 0).UTC()
	if ims, err := time.Parse(http.TimeFormat, ifModifiedSince); err == nil {
		if lm, err := time.Parse(http.TimeFormat, lastModified); err == nil {
			return lm.Unix() <= ims.Unix()
		}
	}
	return false
}
