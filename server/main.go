package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"regexp"

	"github.com/Vetusq/gRPCserver/invoicer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
)

type MyInvoicerServer struct {
	invoicer.UnimplementedGreeterServer
}

func (s MyInvoicerServer) GetBalance(ctx context.Context, req *invoicer.RequestWalletInfo) (*invoicer.ResponceBalanceNonce, error) {
	client, err := ethclient.Dial("https://polygon-rpc.com")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to client: %w", err)
	}

	walletAddress := req.Address
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	if !re.MatchString(walletAddress) {
		log.Fatalf("Address not valid %s", err)
	}

	account := common.HexToAddress(walletAddress)
	nonce, err := client.NonceAt(ctx, account, nil)
	if err != nil {
		log.Fatalf("Failed to get nonce %s", err)
	}

	balanceAt, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatalf("Failed to get balance %s", err)
	}
	balanceInt := balanceAt.Int64()

	return &invoicer.ResponceBalanceNonce{
		Balance: balanceInt,
		Nonce:   nonce,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	server := grpc.NewServer()
	invoicer.RegisterGreeterServer(server, &MyInvoicerServer{})

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
