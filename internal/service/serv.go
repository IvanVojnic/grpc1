package service

import (
	pr "EFms1GRPS/proto"
	"context"
)

type Ms1Server struct {
	pr.UnimplementedMs1Server
}

func NewMs1Server() *Ms1Server {
	return &Ms1Server{}
}

func (s *Ms1Server) Add(ctx context.Context, req *pr.AddRequest) (*pr.AddResponse, error) {
	return &pr.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
