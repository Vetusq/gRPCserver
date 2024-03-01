package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Vetusq/gRPCserver/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := "localhost:8089"
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	walletAddress := "0x84d34f4f83a87596cd3fb6887cff8f17bf5a7b83"

	req := &pb.RequestWalletInfo{
		Wallet: walletAddress,
	}

	resp, err := client.GetBalance(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetBalance: %v", err)
	}

	fmt.Printf("Balance: %d\n WEI\n", resp.Balance)
	fmt.Printf("Nonce: %d\n", resp.Nonce)
}
