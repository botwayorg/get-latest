package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/botwayorg/get-latest/api"
)

func New() *Router {
	r := &Router{
		router: mux.NewRouter(),
	}

	r.initRoutes()

	return r
}

// Router
type Router struct {
	router *mux.Router
}

// ServeHTTP
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *Router) initRoutes() {
	// Root Page
	rootRouter := mux.NewRouter()

	rootRouter.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("welcome, this is get-latest api, please use /:user/:repo to get latest version"))
	}).Methods(http.MethodGet)

	rootRouter.HandleFunc("/{user}/{repo}", api.Latest()).Methods(http.MethodGet)
	rootRouter.HandleFunc("/{user}/{repo}/{no-v}", api.Latest()).Methods(http.MethodGet)
	rootRouter.HandleFunc("/{user}/{repo}/{token}", api.Latest()).Methods(http.MethodGet)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(CORS))
	n.Use(negroni.HandlerFunc(Secure))

	r.router.PathPrefix("/").Handler(n.With(
		LimitHandler(),
		negroni.Wrap(rootRouter),
	))
}
