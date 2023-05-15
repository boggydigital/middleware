package middleware

import (
	"net/http"
	"time"
)

const (
	LastModifiedHeader    = "Last-Modified"
	IfModifiedSinceHeader = "If-Modified-Since"
)

func IsNotModified(ifModifiedSince string, since int64) bool {
	utcSince := time.Unix(since, 0).UTC()
	if ims, err := time.Parse(http.TimeFormat, ifModifiedSince); err == nil {
		return utcSince.Unix() <= ims.UTC().Unix()
	}
	return false
}
