package payments

import (
	"fmt"

	"github.com/marianodsr/ecommerce/products"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func calculateAmount(items products.ItemList) int64 {

	price := 0

	for i := range items.Items {

		product := products.GetProductByID(items.Items[i].ID)
		price += product.Price

	}

	return int64(price)

}

//AttemptPayment func
func AttemptPayment(items products.ItemList) (string, error) {

	stripe.Key = "sk_test_51H3WbrDkOKK75jPQUHGQE8uBH5Y5YsDnKkySHxpeXkO8nUggbM0TgpCJyhvReOvEeRUWlVtWyLSDKsY7B5BuAFcj00xryRCTDT"

	amount := calculateAmount(items)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount * 100)),
		Currency: stripe.String(string(stripe.CurrencyARS)),
	}

	pi, err := paymentintent.New(params)
	fmt.Printf("Client secret: %v", pi.ClientSecret)
	if err != nil {
		fmt.Printf("ERROR CREANDO EL PAYMENT INTENT: %v", err)
		return "", err
	}

	return pi.ClientSecret, nil

}
