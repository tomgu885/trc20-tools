package tron

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc"

	//"github.com/weibi/tron-protocol/api"
)

var (
	tronGridUrl = ""
	apiUrl      = ""
	restyClient *resty.Client
	//walletClient api.WalletClient

	fullApi    string
	fullGrpc   string
	tokenAddr  = ""
	grpcClient *client.GrpcClient
)

// shasta https://api.shasta.trongrid.io
// prod

func init() {
	restyClient = resty.New()
	//grpcClient = client.NewGrpcClient("")
}

// https://developers.tron.network/docs/trongrid
// shasta FullNode GRPC - grpc.shasta.trongrid.io:50051
// mainnet grpc.trongrid.io:50051
func Init(testing bool) {
	if testing { // shasta
		tronGridUrl = "grpc.shasta.trongrid.io:50051"
		fullApi = "https://api.shasta.trongrid.io/"
		tokenAddr = "TSf7K8HVrBWCEfWdpvYLa4K7PMJhRNtkWN"
	} else {
		tronGridUrl = "grpc.trongrid.io:50051"
		fullApi = "https://api.trongrid.io/"
		tokenAddr = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	}
	fmt.Println("tronGridUrl:", tronGridUrl)
	grpcClient = client.NewGrpcClient(tronGridUrl)
	grpcClient.SetAPIKey("5b744b2e-15cf-4311-be05-49ce94c1757b")
	grpcClient.Start(grpc.WithInsecure())
}
