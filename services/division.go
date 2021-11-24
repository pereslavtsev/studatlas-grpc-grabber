package services

import (
	"context"

	"grabber/grabber"
	"grabber/models"
	pb "grabber/pb"
)

type DivisionServiceGrpcImpl struct {
	ctx context.Context
}

func NewDivisionServiceGrpcImpl(ctx context.Context) *DivisionServiceGrpcImpl {
	return &DivisionServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *DivisionServiceGrpcImpl) GetDivision(ctx context.Context, in *pb.GetDivisionRequest) (*pb.Division, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	division, _ := g.GetDivisionById(academy, in.Id)

	return division, nil
}

func (serviceImpl *DivisionServiceGrpcImpl) ListDivisions(ctx context.Context, _ *pb.ListDivisionsRequest) (*pb.ListDivisionsResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)
	divisions, _ := g.GetDivisions(academy)

	return &pb.ListDivisionsResponse{
		Data: divisions,
	}, nil
}
