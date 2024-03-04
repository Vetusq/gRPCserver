package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"regexp"

	"github.com/Vetusq/gRPCserver/invoicer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
)

type MyInvoicerServer struct {
	invoicer.UnimplementedGreeterServer
}

func (s MyInvoicerServer) StreamGetBalance(stream *invoicer.RequestWalletTokenInfo) error {

	client, err := ethclient.Dial("https://polygon-rpc.com")
	if err != nil {
		log.Printf("failed to connect to client: %s", err)
		return nil, fmt.Errorf("failed to connect to client: %w", err)
	}

	walletAddress := reqWalletreqTokenBalance.Balance
	tokenSmartContract := reqTokenBalance.Tokenaddress

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	if !re.MatchString(walletAddress) {
		log.Printf("Wallet address not valid: %s", err)
		return nil, fmt.Errorf("wallet address not valid: %w", err)
	}

	account := common.HexToAddress(walletAddress)

	tokenAddress := common.HexToAddress(tokenSmartContract)

	abi, err := NewERC20(tokenAddress, client)
	if err != nil {
		log.Printf("Failed to connect abi %s", err)
		return nil, fmt.Errorf("failed to connect abi %s", err)
	}

	balance, err := abi.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		log.Fatalf("Failed to get balance %s", err)
	}
	fmt.Println(balance)
	balanceString := balance.String()
	fmt.Println(balance)
	return &invoicer.RequestWalletTokenInfo{
		Balance: balanceString,
	}, nil
}

func (s MyInvoicerServer) GetBalance(ctx context.Context, req *invoicer.RequestWalletInfo) (*invoicer.ResponceBalanceNonce, error) {

	client, err := ethclient.Dial("https://polygon-rpc.com")
	if err != nil {
		log.Printf("failed to connect to client: %s", err)
		return nil, fmt.Errorf("failed to connect to client: %w", err)
	}

	walletAddress := req.Address
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	if !re.MatchString(walletAddress) {
		log.Printf("Wallet address not valid: %s", err)
		return nil, fmt.Errorf("wallet address not valid: %w", err)
	}

	account := common.HexToAddress(walletAddress)
	nonce, err := client.NonceAt(ctx, account, nil)
	if err != nil {
		log.Printf("Failed to get nonce %s", err)
		return nil, fmt.Errorf("failed to get nonce %w", err)
	}

	tokenSmartContract := "0x081Ec4c0e30159C8259BAD8F4887f83010a681DC"
	tokenAddress := common.HexToAddress(tokenSmartContract)

	abi, err := NewERC20(tokenAddress, client)
	if err != nil {
		log.Printf("Failed to connect abi %s", err)
		return nil, fmt.Errorf("failed to connect abi %s", err)
	}

	balance, err := abi.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		log.Fatalf("Failed to get balance %s", err)
	}
	fmt.Println(balance)
	balanceString := balance.String()
	fmt.Println(balance)
	return &invoicer.ResponceBalanceNonce{
		Balance: balanceString,
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
