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
	subrouter.Handle("/", API(post.Index)).
		Methods("GET").
		Name("Index")
	subrouter.Handle("/{id:[0-9]+}", API(post.Get)).
		Methods("GET").
		Name("Get")
	//subrouter.Handle("/{id:[0-9]+}", API(post.Update)).
	//Methods("PUT", "OPTIONS"). // NOTE: OPTIONS crucial here, allows preflight request
	//Name("Update")
	//subrouter.Handle("/", API(post.Create)).
	//Methods("POST").
	//Name("Create")
	//subrouter.Handle("/{id:[0-9]+}", API(post.Delete)).
	//Methods("DELETE").
	//Name("Delete")
	subrouter.Handle("/latest", API(post.Latest)).
		Methods("GET").
		Name("Latest")
	subrouter.Handle("/first", API(post.First)).
		Methods("GET").
		Name("First")
	subrouter.Handle("/{id:[0-9]+}/previous", API(post.Prev)).
		Methods("GET").
		Name("Prev")
	subrouter.Handle("/{id:[0-9]+}/next", API(post.Next)).
		Methods("GET").
		Name("Next")
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
