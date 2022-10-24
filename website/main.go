package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 NotFound", http.StatusNotFound)
	}
	if r.Method != "POST" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}
	name := r.FormValue("user")
	passwd := r.FormValue("passwd")
	_, err := fmt.Fprintln(w, name, passwd)
	if err != nil {
		return
	}
}

func main() {
	staticServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", staticServer)
	http.HandleFunc("/form", formHandler)

	fmt.Println("start server at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("listen 8080 port failed", err)
	}
}
