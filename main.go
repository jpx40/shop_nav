package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(Index()).ServeHTTP(w, r)
	})

	r.Get("/home/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Query()
		templ.Handler(Hello("world2")).ServeHTTP(w, r)
	})
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(Hello("world")).ServeHTTP(w, r)
	})

	fmt.Println("Listening on :3001")

	s := "http://" + "localhost" + ":" + "3001"
	url := strings.TrimSpace(s)
	fmt.Println("", url)
	http.ListenAndServe(":3001", r)
}

func File()

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	_, err := os.Stat(path)

	if err == nil {
		if strings.ContainsAny(path, "{}*") {
			panic("FileServer does not permit any URL parameters.")
		}

		if path != "/" && path[len(path)-1] != '/' {
			r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
			path += "/"
		}
		path += "*"

		r.Get(path, func(w http.ResponseWriter, r *http.Request) {
			rctx := chi.RouteContext(r.Context())
			pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
			fs := http.StripPrefix(pathPrefix, http.FileServer(root))
			fs.ServeHTTP(w, r)
		})
	}
}
