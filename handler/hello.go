package handler

import (
	"fmt"
	"io/ioutil"
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
	h.l.Println("Hello from handler")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}

	log.Printf("Input Data: %s", d)
	fmt.Fprintf(rw, "hello %s", d)
}
