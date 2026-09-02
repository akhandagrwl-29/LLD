package fooddeliveryservice

import (
	"errors"
	"fmt"
)

type FoodDeliveryService struct {
	restaurants    map[string]*Restaurant
	customers      map[string]*Customer
	deliveryAgents map[string]*DeliveryAgent
	orders         map[string]*Order
}

var foodDeliveryService *FoodDeliveryService

func NewFoodDeliveryService() *FoodDeliveryService {
	if foodDeliveryService != nil {
		return foodDeliveryService
	}
	return &FoodDeliveryService{
		restaurants:    make(map[string]*Restaurant),
		customers:      make(map[string]*Customer),
		deliveryAgents: make(map[string]*DeliveryAgent),
		orders:         make(map[string]*Order),
	}
}

func (fds *FoodDeliveryService) AddRestaurant(restaurant *Restaurant) {
	fds.restaurants[restaurant.id] = restaurant
}

func (fds *FoodDeliveryService) AddCustomer(customer *Customer) {
	fds.customers[customer.id] = customer
}

func (fds *FoodDeliveryService) AddDeliveryAgent(agent *DeliveryAgent) {
	fds.deliveryAgents[agent.id] = agent
}

func (fds *FoodDeliveryService) GetAvailableRestaurants() []*Restaurant {
	var restaurants []*Restaurant
	for _, restaurant := range fds.restaurants {
		restaurants = append(restaurants, restaurant)
	}
	return restaurants
}

func (fds *FoodDeliveryService) GetRestaurantMenu(restaurantId string) []*MenuItem {
	menuItems := make([]*MenuItem, 0)
	restaurant, exists := fds.restaurants[restaurantId]
	if !exists {
		return menuItems
	}
	return restaurant.GetItems()
}

func (fds *FoodDeliveryService) PlaceOrder(customerId string, restaurantId string, items []*OrderItem) (*Order, error) {
	customer, exists := fds.customers[customerId]
	if !exists {
		return nil, errors.New("customer not found")
	}

	restaurant, exists := fds.restaurants[restaurantId]
	if !exists {
		return nil, errors.New("restaurant not found")
	}

	order := NewOrder(items, restaurant, customer)
	fds.orders[order.id] = order
	fmt.Printf("Order places succesfully\n")
	return order, nil
}

func (fds *FoodDeliveryService) UpdateOrderStatus(orderId string, status OrderStatus) error {
	order, exists := fds.orders[orderId]
	if !exists {
		return errors.New("order not found")
	}

	if status == OrderConfirmed {
		fds.assignDeliveryAgent(order)
	}
	fds.notifyCustomer(order)
	order.status = status
	return nil
}

func (fds *FoodDeliveryService) assignDeliveryAgent(order *Order) {
	for _, agent := range fds.deliveryAgents {
		if agent.isAvailable {
			order.deliveryAgent = agent
			agent.isAvailable = false
			break
		}
	}
	fmt.Printf("delivery agent assigned\n")
}

func (fds *FoodDeliveryService) notifyCustomer(order *Order) {
	fmt.Printf("Customer Notification: Your order with order id: %s and status %s\n", order.id, order.status)
}

func (fds *FoodDeliveryService) CancelOrder(orderId string) error {
	order, exists := fds.orders[orderId]
	if !exists {
		return errors.New("order not found")
	}
	if order.status == orderDelivered || order.status == OrderCancelled {
		return errors.New("order is already cancelled or delivered")
	}
	order.status = OrderCancelled
	deliveryAgent, exists := fds.deliveryAgents[order.deliveryAgent.id]
	if !exists {
		return errors.New("delivery agent not found")
	}
	deliveryAgent.isAvailable = true
	fds.notifyCustomer(order)
	// notify restaurant
	fmt.Printf("Order is cancelled\n")
	return nil
}
