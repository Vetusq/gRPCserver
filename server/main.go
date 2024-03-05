package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"

	"github.com/Vetusq/gRPCserver/invoicer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
)

type MyInvoicerServer struct {
	invoicer.UnimplementedGreeterServer
}

func (s MyInvoicerServer) StreamGetBalance(stream invoicer.Greeter_StreamGetBalanceServer) error {

	client, err := ethclient.Dial("https://polygon-rpc.com")
	if err != nil {
		log.Printf("failed to connect to client: %s", err)
		return fmt.Errorf("failed to connect to client: %w", err)
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Failed to receive request: %v", err)
			return err
		}

		walletAddress := req.Address
		fmt.Println("Addres geting on server ", walletAddress)
		tokenAddress := common.HexToAddress(req.Tokenaddress)
		fmt.Println("Token address", tokenAddress)

		walletAddressString := string(walletAddress)

		re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
		if !re.MatchString(walletAddressString) {
			log.Printf("Wallet address not valid: %s", err)
			return fmt.Errorf("wallet address not valid")
		}
		account := common.HexToAddress(walletAddressString)

		abi, err := NewERC20(tokenAddress, client)
		if err != nil {
			log.Printf("Failed to connect abi %s", err)
			return fmt.Errorf("failed to connect abi: %w", err)
		}

		balance, err := abi.BalanceOf(&bind.CallOpts{}, account)
		if err != nil {
			log.Printf("Failed to get balance %s", err)
			return fmt.Errorf("failed to get balance: %w", err)
		}

		balanceString := balance.String()
		fmt.Println("Balance of wallet:", balanceString)

		if err := stream.Send(&invoicer.ResponceBalance{Balance: balanceString}); err != nil {
			log.Printf("Failed to send balance update: %v", err)
			return err
		}
	}
}

func (s MyInvoicerServer) GetBalance(ctx context.Context, req *invoicer.RequestWalletInfo) (*invoicer.ResponceBalanceNonce, error) {

	client, err := ethclient.Dial("https://polygon-rpc.com")
	if err != nil {
		log.Printf("failed to connect to client: %s", err)
		return nil, fmt.Errorf("failed to connect to client: %w", err)
	}

	walletAddress := req.Address

	hash := crypto.Keccak256Hash(walletAddress)

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), req.Signature)
	if err != nil {
		log.Fatal(err)
	}

	matchesSignature := bytes.Equal(sigPublicKey, req.Publickeybytes)
	fmt.Println("Signature confirm state =", matchesSignature)

	walletAddressString := string(walletAddress)
	account := common.HexToAddress(walletAddressString)
	nonce, err := client.NonceAt(ctx, account, nil)
	if err != nil {
		log.Printf("Failed to get nonce %s", err)
		return nil, fmt.Errorf("failed to get nonce %w", err)
	}

	tokenSmartContract := "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"
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
		Balance:      balanceString,
		Nonce:        nonce,
		VerifiedSign: matchesSignature,
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
