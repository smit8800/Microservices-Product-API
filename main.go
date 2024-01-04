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

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gg)

	server :=

		http.ListenAndServe("127.0.0.1:8000", sm)
}
