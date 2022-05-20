package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codertjay/go-postgres/router"
)


func main() {
	r := router.Router()
	fmt.Printf("Starting server on the port  8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}