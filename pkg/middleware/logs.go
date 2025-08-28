package middleware

import (
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start := time.Now()
		wrapper := &WrapperWriter{ResponseWriter: w, StatusCode: http.StatusOK}
		next.ServeHTTP(wrapper, r)
		//log.Printf("IP %s - Method: %s, Path: %s, Duration: %s, StatusCode: %d\n", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start), wrapper.StatusCode)
	})
}
