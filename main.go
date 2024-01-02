package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
		}
		log.Println("Data: ", string(d))
		fmt.Fprintf(w, "Hello  %s ", string(d))
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GoodBye World!")
	})
	http.ListenAndServe("127.0.0.1:8000", nil)
}
