package post_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"appengine/aetest"

	. "github.com/richchurcher/eloquent-geek/post"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Post", func() {

	var (
		r     *http.Request
		w     *httptest.ResponseRecorder
		c     aetest.Context
		p, rp *Post
	)

	BeforeEach(func() {
		c, _ = aetest.NewContext(nil)
		w = httptest.NewRecorder()
	})

	Describe("POST route", func() {
		Context("Given a single test post", func() {

			BeforeEach(func() {
				p = &Post{
					Title: "Title",
					Body:  "Body",
					Tags:  []string{"tagone", "tagtwo", "tagthree"},
					Style: "",
				}
				rp = new(Post)
				SeedSinglePost(w, r, c, p)
				ProcessPostResponse(w, p, rp)
			})

			It("should issue HTTP 201", func() {
				Expect(w.Code).To(Equal(201))
			})

			It("should return a valid Post object as JSON", func() {
				Expect(rp).To(Equal(p))
			})

		})
	})

	Describe("GET route", func() {
		Context("Given a single test post", func() {

			var (
				posts       []Post
				rposts      []Post
				getRecorder *httptest.ResponseRecorder
			)

			BeforeEach(func() {
				p = &Post{
					Title: "Title",
					Body:  "Body",
					Tags:  []string{"tagone", "tagtwo", "tagthree"},
					Style: "",
				}
				rp = new(Post)
				SeedSinglePost(w, r, c, p)
				ProcessPostResponse(w, p, rp)

				posts = []Post{*p}
				getRequest, _ := http.NewRequest("GET", "post", nil)
				getRecorder = httptest.NewRecorder()

				err := PostIndex(getRecorder, getRequest, c, "")
				if err != nil {
					log.Fatal(err)
				}
				json.Unmarshal(getRecorder.Body.Bytes(), &rposts)
			})

			It("should issue HTTP 200", func() {
				Expect(getRecorder.Code).To(Equal(200))
			})

			It("should return valid Post object(s) as JSON", func() {
				Expect(rposts).To(Equal(posts))
			})
		})
	})

	Describe("DELETE route", func() {
		Context("Given a single test post", func() {

			var deleteRecorder *httptest.ResponseRecorder

			BeforeEach(func() {
				p = &Post{
					Title: "Title",
					Body:  "Body",
					Tags:  []string{"tagone", "tagtwo", "tagthree"},
					Style: "",
				}
				rp = new(Post)
				SeedSinglePost(w, r, c, p)
				ProcessPostResponse(w, p, rp)

				urlId := fmt.Sprintf("%d", rp.ID)
				deleteRequest, _ := http.NewRequest("DELETE", "post/"+urlId, nil)
				deleteRecorder = httptest.NewRecorder()

				err := PostDelete(deleteRecorder, deleteRequest, c, urlId)
				if err != nil {
					log.Fatal(err)
				}
			})

			It("should issue HTTP 204", func() {
				Expect(deleteRecorder.Code).To(Equal(204))
			})
		})
	})

	Describe("GET (single) route", func() {
		Context("Given a single test post", func() {

			var getRecorder *httptest.ResponseRecorder

			BeforeEach(func() {
				p = &Post{
					Title: "Title",
					Body:  "Body",
					Tags:  []string{"tagone", "tagtwo", "tagthree"},
					Style: "",
				}
				rp = new(Post)
				SeedSinglePost(w, r, c, p)
				ProcessPostResponse(w, p, rp)

				urlId := fmt.Sprintf("%d", rp.ID)
				getRequest, _ := http.NewRequest("GET", "post/"+urlId, nil)
				getRecorder = httptest.NewRecorder()

				err := PostGet(getRecorder, getRequest, c, urlId)
				if err != nil {
					log.Fatal(err)
				}
				json.Unmarshal(getRecorder.Body.Bytes(), &rp)
				p.Created = rp.Created
				p.Modified = rp.Modified
			})

			It("should issue HTTP 200", func() {
				Expect(getRecorder.Code).To(Equal(200))
			})

			It("should return a single Post as JSON", func() {
				Expect(rp).To(Equal(p))
			})
		})
	})

	Describe("PUT route", func() {
		Context("Given a single test post", func() {

			var putRecorder *httptest.ResponseRecorder

			BeforeEach(func() {
				p = &Post{
					Title: "Title",
					Body:  "Body",
					Tags:  []string{"tagone", "tagtwo", "tagthree"},
					Style: "",
				}
				rp = new(Post)
				SeedSinglePost(w, r, c, p)
				ProcessPostResponse(w, p, rp)

				putRecorder = httptest.NewRecorder()

				p.Body = "Expected"
				PutPost(putRecorder, r, c, p)
				json.Unmarshal(putRecorder.Body.Bytes(), &rp)
				p.Created = rp.Created
				p.Modified = rp.Modified
			})

			It("should issue HTTP 200", func() {
				Expect(putRecorder.Code).To(Equal(200))
			})

			It("should return a single Post as JSON", func() {
				Expect(rp).To(Equal(p))
			})
		})
	})
})

func SeedSinglePost(w *httptest.ResponseRecorder, r *http.Request, c aetest.Context, p *Post) error {
	payload, _ := json.Marshal(p)
	r, _ = http.NewRequest("POST", "post", bytes.NewBuffer(payload))
	r.Header.Set("Content-Type", "application/json")

	err := PostCreate(w, r, c, "")
	if err != nil {
		return err
	}

	return nil
}

func PutPost(w *httptest.ResponseRecorder, r *http.Request, c aetest.Context, p *Post) error {
	urlId := fmt.Sprintf("%d", p.ID)
	payload, _ := json.Marshal(p)
	r, _ = http.NewRequest("PUT", "post/"+urlId, bytes.NewBuffer(payload))
	r.Header.Set("Content-Type", "application/json")

	err := PostUpdate(w, r, c, fmt.Sprintf("%d", p.ID))
	if err != nil {
		return err
	}

	return nil
}

// Unmarshal a post, equalise datastore-only values
func ProcessPostResponse(w *httptest.ResponseRecorder, p *Post, rp *Post) {
	json.Unmarshal(w.Body.Bytes(), &rp)
	p.ID = rp.ID
	p.Created = rp.Created
	p.Modified = rp.Modified
}
