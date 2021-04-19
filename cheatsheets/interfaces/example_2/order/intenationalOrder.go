package order

import (
	"fmt"
)

var international = &InternationalOrder{}

type InternationalOrder struct {
	Order
}

func NewInternationalOrder() *InternationalOrder {
	international.products = append(international.products, GetProductDetail("Lap Top", 450, 1, 450.50))
	international.products = append(international.products, GetProductDetail("Video Game", 600, 2, 1200.50))
	international.Client = SetClient("Carl", "Smith", "carlsmith@gmail.com", "9658521365")
	international.ShippingAddress = SetShippingAddress("Colfax Avenue", "Seattle", "USA", "45712")
	return international
}

func (into *InternationalOrder) FillOrderSummary() {
	var extraFee float32 = 0.5
	var taxes float32 = 0.25
	var shippingCost float32 = 35
	subtotal = CalculateSubTotal(into.products)

	totalBeforeTax = (subtotal + shippingCost)
	totalTaxes = (taxes * subtotal)
	totalExtraFee = (totalTaxes * extraFee)
	total = (subtotal + totalTaxes) + totalExtraFee
	into.Summary = Summary{
		total:          total,
		subtotal:       subtotal,
		totalBeforeTax: totalBeforeTax,
	}

}

func (into *InternationalOrder) Notify() {
	email := into.Client.email
	name := into.Client.name
	phone := into.Client.phone

	fmt.Println()
	fmt.Println("---International Order---")
	fmt.Println("Notifying: ", name)
	fmt.Println("Sending email notification to :", email)
	fmt.Println("Sending sms notification to :", phone)
	fmt.Println("Sending whatsapp notification to :", phone)
}

func (into *InternationalOrder) PrintOrderDetails() {
	fmt.Println()
	fmt.Println("International Summary")
	fmt.Println("Order details: ")
	fmt.Println("-- Total Before Taxes: ", into.Summary.totalBeforeTax)
	fmt.Println("-- SubTotal: ", into.Summary.subtotal)
	fmt.Println("-- Total: ", into.Summary.total)
	fmt.Printf("-- Delivery Address to: %s %s %s \n", into.ShippingAddress.street, into.ShippingAddress.city, into.ShippingAddress.country)
	fmt.Printf("-- Client: %s %s \n", into.Client.name, into.Client.lastName)
}
