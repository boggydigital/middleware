package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"github.com/boggydigital/nod"
	"net/http"
)

const defaultWWWAuth = `Basic realm="Restricted", charset="UTF-8"`

var usernameHash, passwordHash [32]byte

func SetUsername(usr [32]byte) {
	usernameHash = usr
}

func SetPassword(pwd [32]byte) {
	passwordHash = pwd
}

func BasicHttpAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if u, p, ok := r.BasicAuth(); ok {
			uh, ph := sha256.Sum256([]byte(u)), sha256.Sum256([]byte(p))

			um := subtle.ConstantTimeCompare(uh[:], usernameHash[:]) == 1
			pm := subtle.ConstantTimeCompare(ph[:], passwordHash[:]) == 1

			if um && pm {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", defaultWWWAuth)
		http.Error(w, nod.ErrorStr("Unauthorized"), http.StatusUnauthorized)
	})
}
