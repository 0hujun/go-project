package main

import (
	"github.com/gorilla/mux"
	"github.com/hujun-petal/go-project/bookstore/pkg/routers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookStoreRouters(r)
	http.Handle("/", r)
	log.Println("start service at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
