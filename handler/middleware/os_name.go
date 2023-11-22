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

func AnalizeOS(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		ctx := r.Context()
		ctx = context.WithValue(ctx, k, ua.OS)
		r = r.Clone(ctx)

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
