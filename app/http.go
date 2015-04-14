package app

import (
	"log"
	"net/http"
	"time"

	"appengine"

	"github.com/gorilla/mux"
)

type Error struct {
	E       error
	Message string
	Code    int
}

func (e *Error) Error() string {
	return e.E.Error()
}

var Router = mux.NewRouter().StrictSlash(true)

func init() {
	http.Handle("/", Router)
}

// Wrap API requests
type API func(http.ResponseWriter, *http.Request, appengine.Context) *Error

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
