package handler

import (
	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

func init() {
	router := httprouter.New()

	router.GET("/books", BookIndex)

	Router = router
}
