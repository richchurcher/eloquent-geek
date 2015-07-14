package post

import (
	"time"

	"appengine"
	"appengine/datastore"
)

type Post struct {
	Title    string    `json:"title"`
	Body     string    `json:"body" datastore:",noindex"`
	Tags     []string  `json:"tags"`
	Style    string    `json:"style"`
	Image    string    `json:"image"`
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
