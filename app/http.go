package app

import "github.com/richchurcher/eloquent-geek/route"

func init() {
	ittp.Handle("/", route.Router)
}
