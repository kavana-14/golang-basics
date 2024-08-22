package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v \n", err)
		return
	}

	fmt.Fprintf(w, "Form submitted successfully \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name: %s \n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
	fmt.Fprintf(w, "Email: %s\n", email)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello!\n")
}

func main() {
		fileServer := http.FileServer(http.Dir("./static"))
		http.Handle("/", fileServer)
		http.HandleFunc("/form", formHandler)
		http.HandleFunc("/hello", helloHandler)

		fmt.Println("Starting server at port 8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
}
