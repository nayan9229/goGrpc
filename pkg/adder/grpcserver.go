package adder

import (
	"context"

	"github.com/nayan9229/goGrpc/api/proto"
)

// GRPCServer struct
type GRPCServer struct {
	api.UnimplementedAdderServer
}

// Add method for calculate X + Y
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{
		Result: req.GetX() + req.GetY(),
	}, nil
}
