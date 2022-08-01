package main

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"example/go-grpc/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	resp, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello(): %s", err)
	}

	log.Printf("Response from server: %s", resp.Body)
}