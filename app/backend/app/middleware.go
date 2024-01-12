package app

import (
	"net/http"
	"strings"
)

// allowContentTypeMiddleware リクエストの Content-Type が指定したものでなければ 415 で応答するミドルウェアを返す。
func allowContentTypeMiddleware(contentTypes ...string) func(next http.Handler) http.Handler {
	allowedContentTypes := make(map[string]struct{}, len(contentTypes))
	for _, ctype := range contentTypes {
		allowedContentTypes[strings.TrimSpace(strings.ToLower(ctype))] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength == 0 {
				next.ServeHTTP(w, r)
				return
			}

			s := strings.ToLower(strings.TrimSpace(r.Header.Get("Content-Type")))
			if i := strings.Index(s, ";"); i > -1 {
				s = s[0:i]
			}

			if _, ok := allowedContentTypes[s]; ok {
				next.ServeHTTP(w, r)
				return
			}

			writeHTTPError(w, http.StatusUnsupportedMediaType)
		})
	}
}
