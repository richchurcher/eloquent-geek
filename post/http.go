package post

import (
	"net/http"
	"strconv"

	"appengine"
	"appengine/datastore"

	"github.com/richchurcher/eloquent-geek/egerror"
)

func Index(w http.ResponseWriter, r *http.Request, c appengine.Context, _ string) *egerror.Error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	posts, err := All(c)
	if err != nil {
		return &egerror.Error{err, "Couldn't get posts.", http.StatusInternalServerError}
	}

	sliceEncode(w, *posts)
	return nil
}

func Get(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
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

func Create(w http.ResponseWriter, r *http.Request, c appengine.Context, _ string) *egerror.Error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	p, err := unmarshal(w, r)
	if err != nil {
		return err.(*egerror.Error)
	}
	p.Save(c)
	w.WriteHeader(http.StatusCreated)
	encode(w, *p)
	return nil
}

func Update(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
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

		p, err := unmarshal(w, r)
		if err != nil {
			return err.(*egerror.Error)
		}
		p.key = datastore.NewKey(c, "Post", "", id, nil)
		if err := p.Save(c); err != nil {
			return &egerror.Error{err, "Unable to save post.", http.StatusInternalServerError}
		}

		w.WriteHeader(http.StatusOK)
		encode(w, *p)
		break
	}

	return nil
}

func Delete(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id, err := strconv.ParseInt(urlId, 10, 64)
	if err != nil {
		// Malformed URL. Most requests should never reach this point.
		return &egerror.Error{err, "Malformed URL.", http.StatusInternalServerError}
	}

	if err := DeleteById(c, id); err != nil {
		return &egerror.Error{err, "Unable to delete post.", http.StatusInternalServerError}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func First(w http.ResponseWriter, r *http.Request, c appengine.Context, _ string) *egerror.Error {
	p, err := GetOne(c, "Created")
	if err != nil {
		return &egerror.Error{err, "Error retrieving first post", http.StatusInternalServerError}
	}

	writeSingle(w, p)
	return nil
}

func Latest(w http.ResponseWriter, r *http.Request, c appengine.Context, _ string) *egerror.Error {
	p, err := GetOne(c, "-Created")
	if err != nil {
		return &egerror.Error{err, "Error retrieving latest post", http.StatusInternalServerError}
	}

	writeSingle(w, p)
	return nil
}

func adjacent(w http.ResponseWriter, c appengine.Context, urlId string, op string) *egerror.Error {
	id, err := strconv.ParseInt(urlId, 10, 64)
	if err != nil {
		return &egerror.Error{err, "Can't find that. Try again?", http.StatusInternalServerError}
	}
	p, err := GetByID(c, id)
	if err != nil {
		return &egerror.Error{err, "Error retrieving next post.", http.StatusInternalServerError}
	}

	adj, err := GetAdjacent(c, op, p.Created)
	if err != nil {
		return &egerror.Error{err, "Error retrieving next post.", http.StatusInternalServerError}
	}

	if adj != nil {
		writeSingle(w, adj)
	} else {
		writeSingle(w, p)
	}
	return nil
}

func Prev(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
	if err := adjacent(w, c, urlId, "<"); err != nil {
		return err
	}
	return nil
}

func Next(w http.ResponseWriter, r *http.Request, c appengine.Context, urlId string) *egerror.Error {
	if err := adjacent(w, c, urlId, ">"); err != nil {
		return err
	}
	return nil
}
