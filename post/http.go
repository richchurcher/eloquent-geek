package post

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"appengine"
	"appengine/datastore"

	"app"

	"github.com/gorilla/mux"
)

func init() {
	s := app.Router.PathPrefix("/post/").Subrouter()
	s.Handle("/", app.API(postIndex)).
		Methods("GET").
		Name("PostIndex")
	s.Handle("/{id:[0-9]+}", app.API(postGet)).
		Methods("GET").
		Name("PostGet")
	s.Handle("/update/{id:[0-9]+}", app.API(postUpdate)).
		Methods("PUT", "OPTIONS"). // NOTE: OPTIONS crucial here, allows preflight request
		Name("PostUpdate")
	s.Handle("/create", app.API(postCreate)).
		Methods("POST").
		Name("PostCreate")
}

func postUnmarshal(w http.ResponseWriter, r *http.Request) (*Post, error) {
	var p Post
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1049576))
	if err != nil {
		return nil, &app.Error{err, "Couldn't read request body.", http.StatusInternalServerError}
	}
	if err := r.Body.Close(); err != nil {
		return nil, &app.Error{err, "Couldn't close request body.", http.StatusInternalServerError}
	}
	if err := json.Unmarshal(body, &p); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			return nil, &app.Error{err, "Couldn't encode error message.", http.StatusInternalServerError}
		}
	}
	return &p, nil
}

// Add an ID field for export to JSON
func postEncode(w http.ResponseWriter, p Post) *app.Error {
	// Add exported ID field for JSON response
	p.Number = p.ID()

	if err := json.NewEncoder(w).Encode(p); err != nil {
		return &app.Error{err, "Couldn't encode JSON.", http.StatusInternalServerError}
	}
	return nil
}

func sliceEncode(w http.ResponseWriter, p []Post) *app.Error {
	for i, _ := range p {
		p[i].Number = p[i].ID()
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		return &app.Error{err, "Couldn't encode JSON.", http.StatusInternalServerError}
	}
	return nil
}

func postIndex(w http.ResponseWriter, r *http.Request, c appengine.Context) *app.Error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	posts, err := All(c)
	if err != nil {
		return &app.Error{err, "Couldn't get posts.", http.StatusInternalServerError}
	}

	sliceEncode(w, *posts)
	return nil
}

func postGet(w http.ResponseWriter, r *http.Request, c appengine.Context) *app.Error {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		// Malformed URL. Most requests should never reach this point (they will
		// be directed to the NotFoundHandler by mux)
		return &app.Error{err, "Can't find that. Try again?", http.StatusInternalServerError}
	}
	p, err := GetByNumber(c, id)
	if err != nil {
		return &app.Error{err, "Couldn't get that post.", http.StatusInternalServerError}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	postEncode(w, *p)
	return nil
}

func postCreate(w http.ResponseWriter, r *http.Request, c appengine.Context) *app.Error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	p, err := postUnmarshal(w, r)
	if err != nil {
		return err.(*app.Error)
	}
	p.Save(c)
	w.WriteHeader(http.StatusCreated)
	postEncode(w, *p)
	return nil
}

func postUpdate(w http.ResponseWriter, r *http.Request, c appengine.Context) *app.Error {
	switch r.Method {
	case "OPTIONS":
		// Preflight request
		w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		break

	case "PUT":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			// Malformed URL. Most requests should never reach this point.
			return &app.Error{err, "Malformed URL.", http.StatusInternalServerError}
		}

		p, err := postUnmarshal(w, r)
		if err != nil {
			return err.(*app.Error)
		}
		p.key = datastore.NewKey(c, "Post", "", id, nil)
		if err := p.Save(c); err != nil {
			return &app.Error{err, "Unable to save post.", http.StatusInternalServerError}
		}

		w.WriteHeader(http.StatusOK)
		postEncode(w, *p)
		break
	}

	return nil
}
