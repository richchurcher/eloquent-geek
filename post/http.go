package post

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"appengine"
	"appengine/datastore"

	"github.com/richchurcher/eloquent-geek/egerror"
)

func postUnmarshal(w http.ResponseWriter, r *http.Request) (*Post, error) {
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

func postEncode(w http.ResponseWriter, p Post) *egerror.Error {
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
		postEncode(w, *p)
	}
}

func PostIndex(w http.ResponseWriter, r *http.Request, c appengine.Context, _ string) *egerror.Error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	posts, err := All(c)
	if err != nil {
		return &egerror.Error{err, "Couldn't get posts.", http.StatusInternalServerError}
	}

	sliceEncode(w, *posts)
	return nil
}

func PostGet(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
	id, err := strconv.ParseInt(urlId, 10, 64)
	if err != nil {
		// Malformed URL. Most requests should never reach this point (they will
		// be directed to the NotFoundHandler by mux)
		return &egerror.Error{err, "Can't find that. Try again?", http.StatusInternalServerError}
	}
	p, err := GetByID(c, id)
	if err != nil {
		return &egerror.Error{err, "Couldn't get that post.", http.StatusInternalServerError}
	}

	writeSingle(w, p)
	return nil
}

func PostCreate(w http.ResponseWriter, r *http.Request, c appengine.Context, _ string) *egerror.Error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	p, err := postUnmarshal(w, r)
	if err != nil {
		return err.(*egerror.Error)
	}
	p.Save(c)
	w.WriteHeader(http.StatusCreated)
	postEncode(w, *p)
	return nil
}

func PostUpdate(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
	switch r.Method {
	case "OPTIONS":
		// Preflight request
		w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		break

	case "PUT":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		id, err := strconv.ParseInt(urlId, 10, 64)
		if err != nil {
			// Malformed URL. Most requests should never reach this point.
			return &egerror.Error{err, "Malformed URL.", http.StatusInternalServerError}
		}

		p, err := postUnmarshal(w, r)
		if err != nil {
			return err.(*egerror.Error)
		}
		p.key = datastore.NewKey(c, "Post", "", id, nil)
		if err := p.Save(c); err != nil {
			return &egerror.Error{err, "Unable to save post.", http.StatusInternalServerError}
		}

		w.WriteHeader(http.StatusOK)
		postEncode(w, *p)
		break
	}

	return nil
}

func PostDelete(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, err := strconv.ParseInt(urlId, 10, 64)
	if err != nil {
		// Malformed URL. Most requests should never reach this point.
		return &egerror.Error{err, "Malformed URL.", http.StatusInternalServerError}
	}

	if err := Delete(c, id); err != nil {
		return &egerror.Error{err, "Unable to delete post.", http.StatusInternalServerError}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func PostFirst(w http.ResponseWriter, r *http.Request, c appengine.Context, _ string) *egerror.Error {
	p, err := GetOne(c, "Created")
	if err != nil {
		return &egerror.Error{err, "Error retrieving first post", http.StatusInternalServerError}
	}

	writeSingle(w, p)
	return nil
}

func PostLatest(w http.ResponseWriter, r *http.Request, c appengine.Context, _ string) *egerror.Error {
	p, err := GetOne(c, "-Created")
	if err != nil {
		return &egerror.Error{err, "Error retrieving latest post", http.StatusInternalServerError}
	}

	writeSingle(w, p)
	return nil
}

func postAdjacent(w http.ResponseWriter, c appengine.Context, urlId string, op string) *egerror.Error {
	id, err := strconv.ParseInt(urlId, 10, 64)
	if err != nil {
		return &egerror.Error{err, "Can't find that. Try again?", http.StatusInternalServerError}
	}
	p, err := GetByID(c, id)
	if err != nil {
		return &egerror.Error{err, "Error retrieving next post.", http.StatusInternalServerError}
	}

	adjacent, err := GetAdjacent(c, op, p.Created)
	if err != nil {
		return &egerror.Error{err, "Error retrieving next post.", http.StatusInternalServerError}
	}

	if adjacent != nil {
		writeSingle(w, adjacent)
	} else {
		writeSingle(w, p)
	}
	return nil
}

func PostPrevious(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
	if err := postAdjacent(w, c, urlId, "<"); err != nil {
		return err
	}
	return nil
}

func PostNext(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
	if err := postAdjacent(w, c, urlId, ">"); err != nil {
		return err
	}
	return nil
}
