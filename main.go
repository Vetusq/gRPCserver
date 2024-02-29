package main

import (
	"log"
	"net"
	"context"
	"github.com/Vetusq/gRPCserver/invoicer"
	"google.golang.org/grpc"
	"context"
    "fmt"
    "math"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

type MyInvoicerServer struct{
	compiled_proto.UnimplementedGreeterServer
}

func(s MyInvoicerServer) GetBalance(context.Context, *invoicer.RequestWalletInfo) (*invoicer.ResponceBalanceNonce, error){
	
	client, err := ethclient.Dial("https://cloudflare-eth.com")
    if err != nil {
        log.Fatalf("Failed to connecting to client...",err)
    }
	walletAddress:= "0xeb2eb5c68156250c368914761bb8f1208d56acd0"
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
    account := common.HexToAddress(walletAddress)

	if re.MatchString(walletAddress) != true{
		log.Fatalf("Address not valid")
	}

	nonce , err:= client.ResponceBalanceNonce
	if err != nil {
        log.Fatal(err)
    }

	blockNumber := big.NewInt(5532993)
    balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
    if err != nil {
        log.Fatal(err)
    }

    fbalance := new(big.Float)
    fbalance.SetString(balanceAt.String())
    ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	
	return &invoicer.ResponceBalanceNonce{
		Balance: ethValue,
	}, nil
}

func main(){
	lis , err : = net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Fail to create listener ", err)
	}
	serverReg := grpc.NewServer()

	service := &MyInvoicerServer{}
	compiled_proto.RegisterGreeterServer(server , service)
	err = serverReg.Serve(lis)
	if err != nil {
		log.Fatalf("Fail to server %s", err)
	}
}