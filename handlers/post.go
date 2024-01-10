package handlers

import (
	"fmt"
	"net/http"
	"project/data"
)

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	p.l.Println("Create Called -------------------------------------->")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	fmt.Print("After fetching value from Prod -------------------------------->")
	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
