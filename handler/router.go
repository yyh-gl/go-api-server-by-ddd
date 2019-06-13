package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var Router *httprouter.Router

func init() {
	router := httprouter.New()

	router.GET("/", Index)

	Router = router
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
