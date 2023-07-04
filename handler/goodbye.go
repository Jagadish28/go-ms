package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye from handler")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}

	log.Printf("Input Data: %s", d)
	fmt.Fprintf(rw, "bye %s", d)
}
