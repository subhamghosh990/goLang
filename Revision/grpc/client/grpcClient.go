package grpcClient

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	abc "abc"

	"google.golang.org/grpc"
)

var abcClient abc.ABCServiceClient
var conn *grpc.ClientConn

const (
	ip   = "127.0.0.1"
	port = 990
)

func startClient() bool {

	servAdd := fmt.Sprintf("%s:%d", ip, port)
	timeOut := 10 * time.Second

	conn, err := grpc.Dial(servAdd, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(timeOut))
	if err != nil {
		fmt.Println("grpc dial err : ", err)
		return false
	}
	abcClient = abc.NewABCServiceClient(conn)
	return true
}

func sendRpcDataToserver() bool {
	if conn == nil {
		if !startClient() {
			return false
		}
	}
	data := &abc.DataResq{Id: 100}
	timeOut := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	resp, err := abcClient.SendRpcData(ctx, data)
	if err != nil {
		fmt.Println("grpc message send err : ", err)
		return false
	}
	respData, _ := json.Marshal(resp)
	fmt.Println("response received : ", string(respData))
	return true
}

func CreateClient(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
start:
	var res bool
	if startClient() {
		res = sendRpcDataToserver()
	}
	if res {
		return
	} else {
		goto start
	}
}
