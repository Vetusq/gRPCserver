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

	walletAddress := "0xa0006e2f7973678e60a2e445ab8353abad6854d5"

	tokenAddress := []string{
		"0x2f6f07cdcf3588944bf4c42ac74ff24bf56e7590", //stg
		"0x7ceb23fd6bc0add59e62ac25578270cff1b9f619", // weth
		"0xc168e40227e4ebd8c1cae80f7a55a4f0e6d66c97", // DFYN
		"0x16eccfdbb4ee1a85a33f3a9b21175cd7ae753db4", // Pos
		"0x081Ec4c0e30159C8259BAD8F4887f83010a681DC", //De
	}

	req1 := &pb.RequestWalletInfo{
		Address: walletAddress,
	}
	resp1, err := client.GetBalance(context.Background(), req1)
	if err != nil {
		log.Fatalf("Error calling GetBalance: %s", err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(resp1.Balance)
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Printf("Balance: %.5f weth \n", ethValue)
	fmt.Printf("Nonce: %d\n", resp1.Nonce)

	// ---

	stream, err := client.StreamGetBalance(context.Background())
	if err != nil {
		log.Fatalf("Error calling StreamGetBalance: %s", err)
	}

	for i := range tokenAddress {

		req2 := &pb.RequestWalletTokenInfo{Address: walletAddress, Tokenaddress: tokenAddress[i]}

		err := stream.Send(req2)
		if err != nil {
			log.Fatalf("Error send wallet info : %v", err)
			continue
		}

		resp, err := stream.Recv()

		if err != nil {
			log.Printf("Error receiving balance: %v", err)
			continue
		}

		fbalance := new(big.Float)
		fbalance.SetString(resp.Balance)
		ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
		tokenName := "UnknownToken"
		switch tokenAddress[i] {
		case "0x2f6f07cdcf3588944bf4c42ac74ff24bf56e7590":
			tokenName = "STG"
		case "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619":
			tokenName = "WETH"
		case "0xc168e40227e4ebd8c1cae80f7a55a4f0e6d66c97":
			tokenName = "DFYN"
		case "0x16eccfdbb4ee1a85a33f3a9b21175cd7ae753db4":
			tokenName = "Pos"
		case "0x081Ec4c0e30159C8259BAD8F4887f83010a681DC":
			tokenName = "De"
		}
		fmt.Printf("Balance: %.5f-%s\n", ethValue, tokenName)

	}

}
