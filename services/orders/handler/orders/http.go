package handler

import (
	"net/http"

	"github.com/ThuraMinThein/go_microservices/services/common/genproto/orders"
	"github.com/ThuraMinThein/go_microservices/services/common/util"
	"github.com/ThuraMinThein/go_microservices/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(ordersService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: ordersService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var req orders.CreateOrderRequest

	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	util.WriteJSON(w, http.StatusOK, res)
}
