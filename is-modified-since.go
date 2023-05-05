package middleware

import "time"

const (
	LastModifiedHeader    = "Last-Modified"
	IfModifiedSinceHeader = "If-Modified-Since"
)

func IsNotModified(ifModifiedSince string, since int64) bool {
	utcSince := time.Unix(since, 0).UTC()
	if ims, err := time.Parse(time.RFC1123, ifModifiedSince); err == nil {
		return utcSince.Unix() <= ims.UTC().Unix()
	}
	return false
}
