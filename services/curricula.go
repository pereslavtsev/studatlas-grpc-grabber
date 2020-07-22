package services

import (
	"context"
	"grabber/grabber"
	"grabber/models"
	pb "grabber/pb"
)

type CurriculaServiceGrpcImpl struct {
	ctx context.Context
}

func NewCurriculaServiceGrpcImpl(ctx context.Context) *CurriculaServiceGrpcImpl {
	return &CurriculaServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *CurriculaServiceGrpcImpl) ListFacultyCurricula(ctx context.Context, in *pb.ListFacultyCurriculaRequest) (*pb.ListCurriculaResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	reports, _ := g.GetFacultyCurricula(academy, in)

	return &pb.ListCurriculaResponse{
		Data: reports,
	}, nil
}
