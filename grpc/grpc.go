package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RegisterGRPC(host string, port int, register func(s *grpc.Server)) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Listening service failed%v", err)
		return err
	}
	s := grpc.NewServer()
	reflection.Register(s)
	register(s)
	log.Println(listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to start the service%v", err)
		return err
	}
	return nil
}
