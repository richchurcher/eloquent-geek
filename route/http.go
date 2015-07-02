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
var subrouter = Router.PathPrefix("/posts").Subrouter().StrictSlash(true)

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
	subrouter.Handle("/latest", API(post.PostLatest)).
		Methods("GET").
		Name("PostLatest")
	subrouter.Handle("/first", API(post.PostFirst)).
		Methods("GET").
		Name("PostFirst")
	subrouter.Handle("/{id:[0-9]+}/previous", API(post.PostPrevious)).
		Methods("GET").
		Name("PostPrevious")
	subrouter.Handle("/{id:[0-9]+}/next", API(post.PostNext)).
		Methods("GET").
		Name("PostNext")
}

// Wrap API requests
type API func(http.ResponseWriter, *http.Request, appengine.Context, string) *egerror.Error

// No sessions, just an API wrapper
func (h API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	start := time.Now()

	// Those handlers that don't need the id can simply ignore it...
	// having it passed in here makes for slightly easier testing.
	if err := h(w, r, c, mux.Vars(r)["id"]); err != nil {
		http.Error(w, err.Error(), err.Code)
	}

	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}
