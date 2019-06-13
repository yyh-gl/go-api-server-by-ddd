package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
