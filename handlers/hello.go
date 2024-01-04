package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World!")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	if string(d) != "" {
		h.l.Println("Data: ", string(d))
		fmt.Fprintf(rw, "Hello  %s ", string(d))
	} else {
		h.l.Println("No data")
		fmt.Fprintf(rw, "Hello")
	}
}
