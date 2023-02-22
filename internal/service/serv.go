package service

import (
	pr "EFms1GRPS/proto"
	"context"
	"io"
	"sync"
	"time"
)

type Ms1Server struct {
	pr.UnimplementedMs1Server
	mu       sync.Mutex // protects routeNotes
	numCheck []*pr.IsEvenNumResponse
}

func NewMs1Server() *Ms1Server {
	return &Ms1Server{}
}

func (s *Ms1Server) Add(ctx context.Context, req *pr.AddRequest) (*pr.AddResponse, error) {
	return &pr.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
func (s *Ms1Server) IsEvan(stream pr.Ms1_IsEvanServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		var check pr.IsEvenNumResponse
		if in.Num%2 == 0 {
			check = pr.IsEvenNumResponse{IsEvan: true}
		} else {
			check = pr.IsEvenNumResponse{IsEvan: false}
		}
		if errSend := stream.Send(&check); errSend != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
}
