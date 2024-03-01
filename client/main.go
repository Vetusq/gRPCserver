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

	walletAddress := "0xef89e95c889f349f8ae3c226d87c94a96f6a9bfc"

	req := &pb.RequestWalletInfo{
		Wallet: walletAddress,
	}

	resp, err := client.GetBalance(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetBalance: %s", err)
	}

	ethBalanceBigInt := big.NewInt(resp.Balance)
	fbalance := new(big.Float)
	fbalance.SetString(ethBalanceBigInt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Printf("Balance: %.5f ETH\n", ethValue)
	fmt.Printf("Nonce: %d\n", resp.Nonce)
}
