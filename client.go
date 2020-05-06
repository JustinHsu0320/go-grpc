package main

import (
	"context"
	"log"

	"github.com/JustinHsu0320/go-grpc-tutorial/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	// 初始化 gRPC 客戶端
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	// 初始化 chat 服務之客戶端
	c := chat.NewChatServiceClient(conn)

	// 遠端調用 chat 結構體 > SayHello 方法
	message := &chat.Message{
		Body: "Hello from the client",
	}
	response, err := c.SayHello(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	// 打印結果
	log.Printf("Response from Server: %s", response.Body)
}
