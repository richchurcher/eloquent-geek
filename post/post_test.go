package post

import (
	"reflect"
	"testing"

	"appengine/aetest"
	"appengine/datastore"
)

func TestSavePost(t *testing.T) {
	// Arrange
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	key := datastore.NewIncompleteKey(c, "Post", nil)
	p := Post{
		"Title",
		"Body",
		[]string{"tagone", "tagtwo", "tagthree"},
		"",
	}

	// Act
	p.Save(c)

	// Assert
	if actual := GetByID(c, key.IntID()); !reflect.DeepEqual(actual, expected) {
		t.Errorf("composeMessage() = %+v, expected %+v", actual, expected)
	}

}
