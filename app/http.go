package app

import (
	"log"
	"net/http"
	"time"

	"appengine"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Error struct {
	E       error
	Message string
	Code    int
}

func (e *Error) Error() string {
	return e.E.Error()
}

var Router = mux.NewRouter()

func init() {
	http.Handle("/", Router)
}

// Public wraps handlers. It will check for a session, but it does not require
// one in order to proceed
type Public func(http.ResponseWriter, *http.Request, appengine.Context, *sessions.Session) *Error

// ServeHTTP implements interface http.Handler for the Public wrapper.
func (h Public) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	store := sessions.NewCookieStore([]byte("VQzwNVF6fLqiMpMQFwB19dlKh1-7_t3u_qNuyl6RpB5J5W3xJjdsk91azpqH6C13z5E-Xs9zv3DzTXjxDEJ7Jw=="))
	s, _ := store.Get(r, "main")
	start := time.Now()

	if err := h(w, r, c, s); err != nil {
		http.Error(w, err.Error(), err.Code)
	}

	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}

type Private func(http.ResponseWriter, *http.Request, appengine.Context, *sessions.Session) *Error

func (h Private) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	store := sessions.NewCookieStore([]byte("VQzwNVF6fLqiMpMQFwB19dlKh1-7_t3u_qNuyl6RpB5J5W3xJjdsk91azpqH6C13z5E-Xs9zv3DzTXjxDEJ7Jw=="))
	s, _ := store.Get(r, "main")

	if err := h(w, r, c, s); err != nil {
		http.Error(w, err.Error(), err.Code)
	}
}

// Wrap API requests
type API func(http.ResponseWriter, *http.Request, appengine.Context) *Error

// ServeHTTP implements interface http.Handler for the API wrapper.
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
