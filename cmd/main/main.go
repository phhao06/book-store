package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/dialect/mysql"
	"github.com/phhao06/book-store/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle(r)
	log.Fatal(http.ListenAndServe("localhost: 9010", r))
}