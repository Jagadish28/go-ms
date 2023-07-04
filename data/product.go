package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func GetProducts() Products {
	return products
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	products = append(products, p)
}

func UpdateProduct(id int, p *Product) error {
	pos, _, err := findProduct(id)

	fmt.Println("pos-->", pos)
	fmt.Println("products-->", products)
	fmt.Println("id-->", id)
	fmt.Println("p-->", p)

	if err != nil {
		return ErrorProductNotFound
	}

	p.ID = id
	products[pos] = p
	return nil
}
func getNextID() int {
	lp := products[len(products)-1]
	fmt.Printf("lp vsl %v", lp.ID)
	return lp.ID + 1

}

var ErrorProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, *Product, error) {
	for i, p := range products {
		if p.ID == id {
			return i, p, nil
		}
	}

	return 0, nil, ErrorProductNotFound

}

var products = []*Product{
	&Product{
		ID:          1,
		Name:        "Nescafe",
		Description: " A good cofee",
		Price:       34.3,
		SKU:         "test123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	}, &Product{
		ID:          2,
		Name:        "Bru",
		Description: " AlsomA good cofee",
		Price:       31.3,
		SKU:         "test1234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
