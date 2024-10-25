package httputil

import (
	"net/http"
	"strings"
)

func CSPJoin(directives []string) string {
	return strings.Join(directives, "; ")
}

type StaticCSPHeader struct {
	CSPDirectives CSPDirectives
}

func (m StaticCSPHeader) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", m.CSPDirectives.String())
		next.ServeHTTP(w, r)
	})
}
