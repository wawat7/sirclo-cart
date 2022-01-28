package sirclo_cart

type Cart struct {
	Name     string
	Quantity int
}

var carts []Cart

func TambahProduk(name string, quantity int) {

	var newCarts []Cart
	if len(carts) != 0 {
		nameFound := false
		for _, cart := range carts {
			if cart.Name == name {
				nameFound = true
				cart.Quantity += quantity
			}
			newCarts = append(newCarts, cart)
		}

		if nameFound == false {
			cart := Cart{
				Name:     name,
				Quantity: quantity,
			}
			newCarts = append(newCarts, cart)

		}
		carts = newCarts
		return
	} else {
		cart := Cart{
			Name:     name,
			Quantity: quantity,
		}

		carts = append(carts, cart)
		return
	}

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
