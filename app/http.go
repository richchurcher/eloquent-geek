package app

import (
	"net/http"

	"github.com/richchurcher/eloquent-geek/route"
)

func init() {
	http.Handle("/", route.Router)
}
