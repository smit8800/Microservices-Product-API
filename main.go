package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Wold!")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
		}
		log.Println("Data: ", d)
		fmt.Fprintf(w, "Hello  %s ", d)
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/form.html")
	})

	http.HandleFunc("/form_pressed", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/welcome.html")
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GoodBye Wold!")
	})
	http.ListenAndServe("127.0.0.1:8000", nil)
}
