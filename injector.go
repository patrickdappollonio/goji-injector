package injector

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

func WithName(name string, x interface{}) func(*web.C, http.Handler) http.Handler {
	return func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			c.Env[name] = x
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
