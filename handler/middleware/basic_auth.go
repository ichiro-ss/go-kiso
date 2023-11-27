package middleware

import (
	"crypto/subtle"
	"net/http"
	"os"
	"strconv"
)

func BasicAuth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		os.Setenv("BASIC_AUTH_USER_ID", "aaa")
		os.Setenv("BASIC_AUTH_PASSWORD", "pass")

		basicAuthUserID := os.Getenv("BASIC_AUTH_USER_ID")
		basicAuthPassword := os.Getenv("BASIC_AUTH_PASSWORD")
		// r.SetBasicAuth("", "")
		userID, password, ok := r.BasicAuth()

		// 単なる文字列比較だと時間によって長さが推測されてしまう
		if ok && subtle.ConstantTimeCompare([]byte(userID), []byte(basicAuthUserID)) == 1 && subtle.ConstantTimeCompare([]byte(password), []byte(basicAuthPassword)) == 1 {
			h.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			err_str := strconv.Itoa(http.StatusUnauthorized) + ":" + "Unauthorized"
			http.Error(w, err_str, http.StatusUnauthorized)

		}
	}
	return http.HandlerFunc(fn)
}
