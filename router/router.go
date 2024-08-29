package router

import (
	"net/http"

	"brijesh.dev/toolkit/middleware"
)

type router struct {
	mux        *http.ServeMux
	middleware []middleware.Middleware
}

func NewRouter() *router {
	return &router{
		mux:        http.NewServeMux(),
		middleware: []middleware.Middleware{},
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	finalHandler := http.Handler(r.mux)
	for i := len(r.middleware) - 1; i >= 0; i-- {
		finalHandler = r.middleware[i](finalHandler)
	}
	finalHandler.ServeHTTP(w, req)
}

func (r *router) Use(mw middleware.Middleware) {
	r.middleware = append(r.middleware, mw)
}

func (r *router) Handle(method, pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, req)
	})
}

func (r *router) GET(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodGet, pattern, handler)
}

func (r *router) POST(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodPost, pattern, handler)
}

func (r *router) PUT(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodPut, pattern, handler)
}

func (r *router) DELETE(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodDelete, pattern, handler)
}

func (r *router) PATCH(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodPatch, pattern, handler)
}

func (r *router) HEAD(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodHead, pattern, handler)
}

func (r *router) OPTIONS(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodOptions, pattern, handler)
}

func (r *router) CONNECT(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodConnect, pattern, handler)
}

func (r *router) TRACE(pattern string, handler http.HandlerFunc) {
	r.Handle(http.MethodTrace, pattern, handler)
}
