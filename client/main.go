package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

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

	walletAddress := "0x00091B44f98a9DfBaF12CfF719bbA49EC41e0000"
	// tokenAddress := "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619"

	req := &pb.RequestWalletInfo{
		Address: walletAddress,
	}
	// req := pb.RequestWalletTokenInfo{
	// 	Address:      walletAddress,
	// 	Tokenaddress: tokenAddress,
	// }
	resp, err := client.GetBalance(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetBalance: %s", err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(resp.Balance)
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Printf("Balance: %.5f de \n", ethValue)
	fmt.Printf("Nonce: %d\n", resp.Nonce)
}
