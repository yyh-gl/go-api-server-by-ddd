package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yyh-gl/go-api-server-by-ddd/infra"
)

func main() {
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:3000")
	fmt.Println("========================")
	log.Fatal(http.ListenAndServe(":3000", infra.Router))
}
