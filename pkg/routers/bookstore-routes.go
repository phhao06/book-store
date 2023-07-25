package routers

import (
	"github.com/gorilla/mux"
	"github.com/phhao06/book-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(route *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Method("POST")
	router.HandleFunc("/book/", controllers.GetBook).Method("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Method("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Method("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Method("DELETE")
}