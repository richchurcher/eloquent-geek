package route

import (
	"log"
	"net/http"
	"time"

	"appengine"

	"github.com/gorilla/mux"
	"github.com/richchurcher/eloquent-geek/egerror"
	"github.com/richchurcher/eloquent-geek/post"
)

var Router = mux.NewRouter()
var subrouter = Router.PathPrefix("/post").Subrouter().StrictSlash(true)

func init() {
	subrouter.Handle("/", API(post.PostIndex)).
		Methods("GET").
		Name("PostIndex")
	subrouter.Handle("/{id:[0-9]+}", API(post.PostGet)).
		Methods("GET").
		Name("PostGet")
	subrouter.Handle("/{id:[0-9]+}", API(post.PostUpdate)).
		Methods("PUT", "OPTIONS"). // NOTE: OPTIONS crucial here, allows preflight request
		Name("PostUpdate")
	subrouter.Handle("/", API(post.PostCreate)).
		Methods("POST").
		Name("PostCreate")
	subrouter.Handle("/{id:[0-9]+}", API(post.PostDelete)).
		Methods("DELETE").
		Name("PostDelete")
}

// Wrap API requests
type API func(http.ResponseWriter, *http.Request, appengine.Context) *egerror.Error

// No sessions, just an API wrapper
func (h API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	start := time.Now()

	if err := h(w, r, c); err != nil {
		http.Error(w, err.Error(), err.Code)
	}

	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}
