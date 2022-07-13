package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/form", formhandler)

	fmt.Println("server starting at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "HELLO TO THIS SERVER")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform();err=%v", err)
		return
	}

	fmt.Fprintf(w, "POST METHOD SUCCESSFULLY\n")
	name := r.FormValue("name")
	Branch := r.FormValue("Branch")
	Age := r.FormValue("Age")
	Roolno := r.FormValue("Roolno")
	fmt.Fprintf(w, "name=%v\n", name)
	fmt.Fprintf(w, "Branch=%v\n", Branch)
	fmt.Fprintf(w, "Age=%v\n", Age)
	fmt.Fprintf(w, "Roolno=%v\n", Roolno)
}
