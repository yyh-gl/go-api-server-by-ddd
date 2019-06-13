package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:3000")
	fmt.Println("========================")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
