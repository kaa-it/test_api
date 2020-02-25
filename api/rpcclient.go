package api

import (
	rpc "astro_pro/rpc/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

type RpcClient struct {
	conn *grpc.ClientConn
	c    rpc.RpcClient
}

func NewRpcClient() *RpcClient {
	viper.SetConfigName("astro_pro.conf")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	rpcPort := "5555"

	if viper.IsSet("backend.rpc_port") {
		rpcPort = viper.GetString("backend.rpc_port")
	}

	rpcClient := &RpcClient{}

	rpcClient.conn, err = grpc.Dial("backend:"+rpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	rpcClient.c = rpc.NewRpcClient(rpcClient.conn)

	return rpcClient
}
