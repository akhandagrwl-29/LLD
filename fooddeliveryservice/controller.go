package fooddeliveryservice

import "fmt"

func Run() {
	foodDeliveryService := NewFoodDeliveryService()
	item1 := NewMenuItem("1", "Burger", "tasty burger", true, 10)
	item2 := NewMenuItem("1", "Pizza", "tasty pizza", true, 10)
	restaurant1 := NewRestaurant("R1", "Burger King", "Bengaluru")
	restaurant1.AddItem(item1)
	restaurant1.AddItem(item2)

	foodDeliveryService.AddRestaurant(restaurant1)

	deliveryAgent1 := NewDeliveryAgent("d1", "sam", "7886867895", true)
	foodDeliveryService.AddDeliveryAgent(deliveryAgent1)

	customer1 := NewCustomer("c1", "ram", "7886867895")
	foodDeliveryService.AddCustomer(customer1)

	fmt.Printf("Food delivery service is running with restuarants :%d\n", len(foodDeliveryService.GetAvailableRestaurants()))
	for _, restaurant := range foodDeliveryService.GetAvailableRestaurants() {
		fmt.Printf("Restaurant : %v\n and dishes present are : %+v\n", *restaurant, restaurant.GetItems())
	}

	order, err := foodDeliveryService.PlaceOrder(customer1.id, restaurant1.id, []*OrderItem{
		{
			menuItem: item1,
			quantity: 1,
		},
		{
			menuItem: item2,
			quantity: 2,
		},
	})
	if err != nil {
		fmt.Printf("Unable to place order error: %+v\n", err)
	}
	fmt.Printf("Order places successfully with order id :%v\n", order.id)

	foodDeliveryService.UpdateOrderStatus(order.id, OrderConfirmed)
	foodDeliveryService.UpdateOrderStatus(order.id, OrderOutForDelivery)
	foodDeliveryService.UpdateOrderStatus(order.id, orderDelivered)

}
