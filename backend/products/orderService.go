package products

//CreateOrder func

var orderRepo orderRepository

//CreateOrder func
func CreateOrder(name string, email string, adress string, products []Product) {

	order := &Order{
		ClientName:   name,
		ClientEmail:  email,
		ClientAdress: adress,
		Products:     products,
	}

	orderRepo.saveOrder(order)

}

func Migrate() {

	orderRepo.migrate()
}
