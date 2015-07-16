package post

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/richchurcher/eloquent-geek/egerror"
)

func unmarshal(w http.ResponseWriter, r *http.Request) (*Post, error) {
	var p Post
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1049576))
	if err != nil {
		return nil, &egerror.Error{err, "Couldn't read request body.", http.StatusInternalServerError}
	}
	if err := r.Body.Close(); err != nil {
		return nil, &egerror.Error{err, "Couldn't close request body.", http.StatusInternalServerError}
	}
	if err := json.Unmarshal(body, &p); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			return nil, &egerror.Error{err, "Couldn't encode error message.", http.StatusInternalServerError}
		}
	}
	return &p, nil
}

func encode(w http.ResponseWriter, p Post) *egerror.Error {
	// Add exported ID field for JSON response
	p.ID = p.GetID()

	if err := json.NewEncoder(w).Encode(p); err != nil {
		return &egerror.Error{err, "Couldn't encode JSON.", http.StatusInternalServerError}
	}
	return nil
}

func sliceEncode(w http.ResponseWriter, p []Post) *egerror.Error {
	for i, _ := range p {
		p[i].ID = p[i].GetID()
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		return &egerror.Error{err, "Couldn't encode JSON.", http.StatusInternalServerError}
	}
	return nil
}

func writeSingle(w http.ResponseWriter, p *Post) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if p == nil {
		// Request was ok, but no post in datastore
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		encode(w, *p)
	}
}
