package main

import (
	"log"
	"net"
	"context"
	"github.com/Vetusq/gRPCserver/invoicer"
	"google.golang.org/grpc"

)

type MyInvoicerServer struct{
	compiled_proto.UnimplementedGreeterServer
}

func(s MyInvoicerServer) GetBalance(context.Context, *compiled_proto.RequestWalletInfo) (*compiled_proto.ResponceBalanceNonce, error){

	return &compiled_proto.RequestWalletInfo{
		Wallet: "0xeb2eb5c68156250c368914761bb8f1208d56acd0",
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
		log.Fatalf("Fail to server ", err)
	}
}