package main

import (
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/mcp-bank/broker-service/internal/broker"
	"github.com/mcp-bank/proto/gen/brokerv1"
	"google.golang.org/grpc"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		slog.Error("failed to listen",
			"err", err)
		return
	}
	server := grpc.NewServer()
	brokerv1.RegisterBrokerServiceServer(server, broker.New())
	go func() {
		if err = server.Serve(lis); err != nil {
			slog.Error("failed to serve",
				"err", err)
			return
		}
	}()
	<-shutdown
	slog.Info("graceful shutdown")
	server.GracefulStop()
}
