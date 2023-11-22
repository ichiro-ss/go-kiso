package middleware

import (
	"context"
	"net/http"

	"github.com/mileusna/useragent"
)

type ctxKey string

const (
	k ctxKey = "OS"
)

func PutOsNameOnContext(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		ctx := r.Context()
		ctx = context.WithValue(ctx, k, ua.OS)

		// val, ok := r.Context().Value(k).(string)
		// fmt.Println("--before")
		// if ok {
		// 	fmt.Println(val)
		// }

		r = r.Clone(ctx)

		// val, ok := r.Context().Value(k).(string)
		// fmt.Println("--after")
		// if ok {
		// 	fmt.Println(val)
		// }

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
