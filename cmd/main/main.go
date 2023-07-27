package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phhao06/book-store/pkg/routers"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is listen at 9010")
	log.Fatal(http.ListenAndServe("localhost: 9010", r))
}
