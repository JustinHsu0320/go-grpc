package main

import (
	"log"
	"net"

	"github.com/JustinHsu0320/go-grpc-tutorial/chat"
	"google.golang.org/grpc"
)

func main() {
	// gRPC Server 將使用 port 9001
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen on port 9001: %v", err)
	}

	// gRPC Server 初始
	grpcServer := grpc.NewServer()

	// 於 gRPC Server 註冊 chat.Server 結構體，以暴露其 SayHello 方法
	s := &chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, s)

	// gRPC Server 使用 port 9001
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9001: %v", err)
	}
}
