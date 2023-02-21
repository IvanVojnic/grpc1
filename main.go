package main

import (
	"EFms1GRPS/internal/service"
	pr "EFms1GRPS/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := service.NewMs1Server()
	pr.RegisterMs1Server(s, srv)
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		defer logrus.Fatalf("error while listening port: %e", err)
	}

	if errServ := s.Serve(listen); errServ != nil {
		defer logrus.Fatalf("error while listening server: %e", err)
	}
}
