package infra

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yyh-gl/go-api-server-by-ddd/handler"
)

var Router *httprouter.Router

func init() {
	router := httprouter.New()

	router.GET("/books", handler.BookIndex)

	Router = router
}
