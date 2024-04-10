package prodcut_core

import (
	"context"

	pb "course.project/authentication/proto/core"
)

type CoreServerImpl struct{}

func NewServer() CoreServerImpl {
	return CoreServerImpl{}
}

func (c CoreServerImpl) AddProduct(ctx context.Context, in *pb.Product) (*pb.Error, error)
