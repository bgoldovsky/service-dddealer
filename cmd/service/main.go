package main

import (
	"context"
	"log"
	"net"

	"github.com/bgoldovsky/dealer/service-dealer/internal/handlers/dealer"

	"github.com/bgoldovsky/dealer/service-dealer/internal/generated/rpc"

	"github.com/bgoldovsky/dealer/service-dealer/internal/app/repository/order"
	orderSrv "github.com/bgoldovsky/dealer/service-dealer/internal/app/services/order"
	"github.com/bgoldovsky/dealer/service-dealer/internal/interceptors"

	"github.com/bgoldovsky/dealer/service-dealer/internal/cfg"

	"github.com/bgoldovsky/dealer/service-dealer/internal/app/logger"
	"github.com/bgoldovsky/dealer/service-dealer/internal/database"
	"google.golang.org/grpc"
)

func main() {
	// Repo
	db := database.NewDatabase(context.Background(), cfg.GetConnString())
	ordersRepo := order.New(db)

	// Services
	orderService := orderSrv.New(ordersRepo)

	// Handlers
	dealerHandler := dealer.New(orderService)

	// Registration
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LoggingInterceptor))
	rpc.RegisterDealerServer(server, dealerHandler)

	// Start
	port := cfg.GetGrpcPort()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	logger.Log.WithField("port", port).Infoln("gRPC server starts..")
	log.Fatal(server.Serve(listener))
}
