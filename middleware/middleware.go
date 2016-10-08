package middleware

import (
	"context"
	"net/http"
)

type key int

const (
	responseWrittenKey key = 0
)

// Middleware is a middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain wraps a handler in Middleware
// Middleware can be terminated by setting the value of the Context's
// responseWrittenKey to true
func Chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, responseWrittenKey, false)
	wrapped := handler
	for _, middleware := range middlewares {
		// This is a bit of a black box. What we are doing is making a new handler
		// that only runs its wrapped (i.e. next along the chain) handler
		// if the previous handler didn't set the Context's responseWrittenKey
		wrapped = func(handler http.HandlerFunc) http.HandlerFunc {
			return middleware(func(rw http.ResponseWriter, r *http.Request) {
				done, ok := r.Context().Value(responseWrittenKey).(bool)
				if !ok {
					// log.Println("Cast failed")
					return
				}
				if done {
					// log.Println("Done, do nothing")
					cancel()
					return
				}
				// log.Println("Not done, run handler")
				handler(rw, r)
			})
		}(wrapped)
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		r = r.WithContext(ctx)
		wrapped(rw, r)
		cancel()
	}
}
