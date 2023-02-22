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
	log.Fatal(http.ListenAndServe("192.168.124.10:80", r))
}
