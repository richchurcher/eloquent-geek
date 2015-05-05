package post

import (
	"reflect"
	"testing"

	"appengine/aetest"
	"appengine/datastore"
)

// A completely pointless test... that works!!!!!
func TestSavePost(t *testing.T) {
	// Arrange
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	p := Post{
		Title: "Title",
		Body:  "Body",
		Tags:  []string{"tagone", "tagtwo", "tagthree"},
		Style: "",
	}

	// Act
	p.Save(c)
	id := p.GetID()
	k := datastore.NewKey(c, "Post", "", id, nil)
	expected := new(Post)
	datastore.Get(c, k, expected)
	expected.key = k

	// Assert
	if actual, _ := GetByID(c, k.IntID()); !reflect.DeepEqual(actual, expected) {
		t.Errorf("composeMessage() = %+v, expected %+v", actual, expected)
	}
}
