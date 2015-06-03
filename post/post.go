package post

import (
	"html/template"
	"time"

	"appengine"
	"appengine/datastore"
)

type Post struct {
	Title    string    `json:"title"`
	Body     string    `json:"body" datastore:",noindex"`
	Tags     []string  `json:"tags"`
	Style    string    `json:"style"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	ID       int64     `json:"id" datastore:"-"` // present in JSON response, absent in datastore
	key      *datastore.Key
}

func (p *Post) Save(c appengine.Context) error {
	if p.key == nil {
		// New post, was not retrieved from datastore
		p.key = datastore.NewIncompleteKey(c, "Post", nil)
		p.Created = time.Now()
	}
	p.Modified = time.Now()

	k, err := datastore.Put(c, p.key, p)
	if err != nil {
		return err
	}
	p.key = k
	return nil
}

func (p *Post) GetID() int64 {
	if p.key == nil {
		return 0
	}
	return p.key.IntID()
}

func GetByID(c appengine.Context, id int64) (*Post, error) {
	k := datastore.NewKey(c, "Post", "", id, nil)
	p := new(Post)
	if err := datastore.Get(c, k, p); err != nil {
		return nil, err
	}
	p.key = k
	return p, nil
}

func All(c appengine.Context) (*[]Post, error) {
	var p []Post
	q := datastore.NewQuery("Post").Order("-Created")
	k, err := q.GetAll(c, &p)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(p); i++ {
		p[i].key = k[i]
	}

	// Always return at least an empty slice
	if p == nil {
		p = []Post{}
	}
	return &p, nil
}

func Delete(c appengine.Context, id int64) error {
	k := datastore.NewKey(c, "Post", "", id, nil)
	if err := datastore.Delete(c, k); err != nil {
		return err
	}
	return nil
}

func TrustedHTML(b []byte) template.HTML {
	return template.HTML(b)
}
