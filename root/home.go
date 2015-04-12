// Package root provides handlers for path "/"
package root

import (
	"html/template"
	"net/http"

	"appengine"

	"app"
	"post"

	"github.com/gorilla/sessions"
)

func init() {
	app.Router.Handle("/", app.Public(home))
	app.Router.NotFoundHandler = http.HandlerFunc(notFound)
}

func home(w http.ResponseWriter, r *http.Request, c appengine.Context, s *sessions.Session) *app.Error {
	p, err := post.All(c)
	if err != nil {
		return &app.Error{err, "Couldn't get posts from datastore.", http.StatusInternalServerError}
	}

	fm := template.FuncMap{
		"trustedHTML": post.TrustedHTML,
	}
	t, err := template.New("home").
		Funcs(fm).
		ParseFiles("templates/base.html", "templates/home.html")
	if err != nil {
		return &app.Error{err, "Template parse error.", http.StatusInternalServerError}
	}
	if err := t.ExecuteTemplate(w, "base", p); err != nil {
		return &app.Error{err, "Template execute error.", http.StatusInternalServerError}
	}
	return nil
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Can't find that.", http.StatusNotFound)
}
