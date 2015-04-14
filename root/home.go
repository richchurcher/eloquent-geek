// Package root provides handlers for path "/"
package root

import (
	//"html/template"
	"net/http"

	//"appengine"

	"app"

	//"github.com/gorilla/sessions"
)

func init() {
	app.Router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	}))
	app.Router.NotFoundHandler = http.HandlerFunc(notFound)
}

//func home(w http.ResponseWriter, r *http.Request, c appengine.Context, s *sessions.Session) *app.Error {
	//t, err := template.New("home").
		//ParseFiles("templates/base.html", "templates/home.html")
	//if err != nil {
		//return &app.Error{err, "Template parse error.", http.StatusInternalServerError}
	//}
	//if err := t.ExecuteTemplate(w, "base", nil); err != nil {
		//return &app.Error{err, "Template execute error.", http.StatusInternalServerError}
	//}
	//return nil
//}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Can't find that.", http.StatusNotFound)
}
