package main

import (
	"log"
	"net/http"
	"os"
	"project/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gg := handlers.NewGoodbye(l)

	ss := http.NewServeMux()
	ss.Handle("/", hh)
	ss.Handle("/goodbye", gg)

	http.ListenAndServe("127.0.0.1:8000", ss)
}
