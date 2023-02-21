package service

import (
	pr "EFms1GRPS/proto"
	"context"
	"io"
	"sync"
)

type Ms1Server struct {
	pr.UnimplementedMs1Server
	mu         sync.Mutex // protects routeNotes
	routeNotes map[string][]*pr.Ms1Server
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

		s.mu.Lock()
		//s.routeNotes[key] = append(s.routeNotes[key], in.)
		//rn := make([]*pb.RouteNote, len(s.routeNotes[key]))
		//copy(rn, s.routeNotes[key])
		s.mu.Unlock()

		for _, note := range rn {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}
