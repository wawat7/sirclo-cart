package sirclo_cart

type Cart struct {
	Name     string
	Quantity int
}

var carts []Cart

func TambahProduk(name string, quantity int) {
	cart := Cart{
		Name:     name,
		Quantity: quantity,
	}

	carts = append(carts, cart)
}

func HapusProduk(name string) {
	var newCarts []Cart
	for _, cart := range carts {
		if cart.Name != name {
			newCarts = append(newCarts, cart)
		}
	}

	carts = newCarts
}

func TampilkanCart() []Cart {
	return carts
}
