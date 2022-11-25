package grpcServar

import (
	"context"
	"fmt"
	"net"
	"sync"

	abc "abc"

	"google.golang.org/grpc"
)

type Server struct {
}

var grpcServer *grpc.Server

func startGRPCServer() bool {
	grpcServer = grpc.NewServer()
	s := Server{}
	lis, err := net.Listen("tcp", ":990")
	if err != nil {
		fmt.Println("Not able to start grpc port : ", err)
		return false
	}

	abc.RegisterABCServiceServer(grpcServer, &s)
	if err = grpcServer.Serve(lis); err != nil {
		fmt.Println("Not able to start grpc server : ", err)
		return false
	}
	return true
}

func stopGRPCServer() {
	grpcServer.Stop()
}

func (s *Server) SendRpcData(ctx context.Context, req *abc.DataResq) (*abc.DataResq, error) {
	fmt.Println("received request from client : ", req)
	resp := &abc.DataResq{Id: req.Id * 100}
	return resp, nil
}

func CreateServer(wg *sync.WaitGroup) {
	if wg != nil {
		if grpcServer != nil {
			defer stopGRPCServer()
		}
		defer wg.Done()
	}
	if !startGRPCServer() {
		return
	}

}
