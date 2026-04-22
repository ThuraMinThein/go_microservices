package types

import (
	"context"

	"github.com/ThuraMinThein/go_microservices/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
