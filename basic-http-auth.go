package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"github.com/boggydigital/nod"
	"net/http"
)

const defaultWWWAuth = `Basic realm="Restricted", charset="UTF-8"`

var usernames, passwords = make(map[string][32]byte), make(map[string][32]byte)

func SetUsername(role string, usr [32]byte) {
	usernames[role] = usr
}

func SetPassword(role string, pwd [32]byte) {
	passwords[role] = pwd
}

func BasicHttpAuth(next http.Handler, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if u, p, ok := r.BasicAuth(); ok {
			uh, ph := sha256.Sum256([]byte(u)), sha256.Sum256([]byte(p))

			for _, role := range roles {
				if usernameHash, ok := usernames[role]; ok {
					if passwordHash, ok := passwords[role]; ok {

						um := subtle.ConstantTimeCompare(uh[:], usernameHash[:]) == 1
						pm := subtle.ConstantTimeCompare(ph[:], passwordHash[:]) == 1

						if um && pm {
							next.ServeHTTP(w, r)
							return
						}
					}
				}
			}
		}

		w.Header().Set("WWW-Authenticate", defaultWWWAuth)
		http.Error(w, nod.ErrorStr("Unauthorized"), http.StatusUnauthorized)
	})
}
