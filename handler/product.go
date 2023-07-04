package handler

import (
	"fmt"
	"go-ms/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, _ := strconv.Atoi(idString)
		p.l.Println("id", id)
		p.updateProduct(id, rw, r)
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
	// rw.Write(d)

}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("post method")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusBadGateway)
	}

	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)

}

func (p *Product) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	fmt.Println("post method")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusBadGateway)
	}

	p.l.Printf("Prod: %#v", prod)

	err = data.UpdateProduct(id, prod)

	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not fount", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not fount", http.StatusInternalServerError)
		return
	}

}
