package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ThuraMinThein/go_microservices/services/common/genproto/orders"
	"github.com/ThuraMinThein/go_microservices/services/common/util"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 123,
			ProductID:  456,
			Quantity:   2,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		o, err := c.GetOrders(ctx, &orders.GetOrderRequest{
			CustomerID: 42,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		util.WriteJSON(w, http.StatusOK, o)
	})

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
