package fooddeliveryservice

import (
	"fmt"
	"time"
)

type OrderItem struct {
	menuItem *MenuItem
	quantity int
}

type Order struct {
	id            string
	orderItems    []*OrderItem
	restaurant    *Restaurant
	customer      *Customer
	deliveryAgent *DeliveryAgent
	totalAmount   int
	status        OrderStatus
}

func NewOrder(orderItems []*OrderItem, restaurant *Restaurant, customer *Customer) *Order {
	return &Order{
		id:         fmt.Sprintf("%d", time.Now().UnixNano()),
		orderItems: orderItems,
		restaurant: restaurant,
		customer:   customer,
	}
}
