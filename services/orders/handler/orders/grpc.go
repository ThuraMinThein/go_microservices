package handler

import (
	"context"

	"github.com/ThuraMinThein/go_microservices/services/common/genproto/orders"
	"github.com/ThuraMinThein/go_microservices/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGRPCHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGRPCOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	grpcHandler := &OrdersGRPCHandler{
		ordersService: ordersService,
	}

	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrdersGRPCHandler) GetOrders(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	o := h.ordersService.GetOrders(ctx)

	res := &orders.GetOrderResponse{
		Orders: o,
	}
	return res, nil
}

func (h *OrdersGRPCHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 1,
		ProductID:  1,
		Quantity:   1,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	return res, nil
}
