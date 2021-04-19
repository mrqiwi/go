package order

import (
	"fmt"
)

func New() {
	fmt.Println("Order package!!")
	natOrd := NewNationalOrder()
	intOrd := NewInternationalOrder()
	ords := []Operations{natOrd, intOrd}
	ProcessOrder(ords)
}

type Product struct {
	name  string
	price int
}

type ProductDetail struct {
	Product
	amount int
	total  float32
}

type Summary struct {
	total          float32
	subtotal       float32
	totalBeforeTax float32
}

type ShippingAddress struct {
	street  string
	city    string
	country string
	cp      string
}

type Client struct {
	name     string
	lastName string
	email    string
	phone    string
}

type Order struct {
	products []*ProductDetail
	Summary
	ShippingAddress
	Client
}

type Processer interface {
	FillOrderSummary()
}

type Printer interface {
	PrintOrderDetails()
}

type Notifier interface {
	Notify()
}

type Operations interface {
	Processer
	Printer
	Notifier
}

func ProcessOrder(orders []Operations) {
	for _, order := range orders {
		order.FillOrderSummary()
		order.Notify()
		order.PrintOrderDetails()
	}
}
