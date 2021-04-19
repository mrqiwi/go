package order

import (
	"fmt"
)

var national = &NationalOrder{}

type NationalOrder struct {
	Order
}

func NewNationalOrder() *NationalOrder {
	national.products = append(national.products, GetProductDetail("Sugar", 12, 3, 36))
	national.products = append(national.products, GetProductDetail("Cereal", 16, 2, 36))
	national.Client = SetClient("Phill", "Heat", "phill@gmail.com", "8415748569")
	national.ShippingAddress = SetShippingAddress("North Ave", "San Antonio", "USA", "854789")
	return national
}

func (nato *NationalOrder) FillOrderSummary() {
	var taxes float32 = 0.20
	var shippingCost float32 = 5
	subtotal = CalculateSubTotal(nato.products)

	totalBeforeTax = (subtotal + shippingCost)
	totalTaxes = (taxes * subtotal)
	total = (subtotal + totalTaxes)

	nato.Summary = Summary{
		total,
		subtotal,
		totalBeforeTax,
	}
}

func (nato *NationalOrder) Notify() {
	email := nato.Client.email
	fmt.Println("---National Order---")
	fmt.Println("Sending email notification to:", email)
}

func (nato *NationalOrder) PrintOrderDetails() {
	fmt.Println()
	fmt.Println("National Summary")
	fmt.Println("Order details: ")
	fmt.Println("Total: ", nato.Summary.total)
	fmt.Printf("Delivery Address to: %s %s %s \n", nato.ShippingAddress.street, nato.ShippingAddress.city, nato.ShippingAddress.country)
}
