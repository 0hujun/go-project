package routers

import (
	"github.com/gorilla/mux"
	"github.com/hujun-petal/go-project/bookstore/pkg/controllers"
)

var RegisterBookStoreRouters = func(r *mux.Router) {
	r.HandleFunc("/", controllers.HomePage).Methods("GET")
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
