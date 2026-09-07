package domain

import "time"

type OrderStatus string

const(
	OrderAwaitingShipment OrderStatus = "awaiting_shipment"
	OrderCancelled OrderStatus = "cancelled"
)

type OrderItem struct {
	SKU string
	Quantity int
	Price float64
}

type Order struct {
	ID string
	Market string
	Currency string
	Amount float64
	Status OrderStatus
	FulfillmentWH string
	Cancellation bool
	Items []OrderItem
}

type InventoryBalance struct {
	WarehouseID string
	SKU string
	Available int
}

type InventoryTransfer struct {
	ID string
	SKU string
	FormWareHouse string
	ToWareHouse string
	Quantity string
	Status string
	CreateAt time.Time
}