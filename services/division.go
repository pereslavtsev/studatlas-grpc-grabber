package services

import (
	"context"

	"grabber/database"
	"grabber/grabber"
	pb "grabber/pb"

	log "github.com/sirupsen/logrus"
)

type DivisionServiceGrpcImpl struct {
	ctx context.Context
}

func NewDivisionServiceGrpcImpl(ctx context.Context) *DivisionServiceGrpcImpl {
	return &DivisionServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *DivisionServiceGrpcImpl) GetDivision(_ context.Context, in *pb.GetDivisionRequest) (*pb.Division, error) {
	db := database.FromContext(serviceImpl.ctx)
	academy, err := db.GetAcademy(in.AcademyId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	g := grabber.FromContext(serviceImpl.ctx)

	division, err := g.GetDivisionById(academy, in.Id)

	return division, nil
}

func (serviceImpl *DivisionServiceGrpcImpl) ListDivisions(_ context.Context, in *pb.ListDivisionsRequest) (*pb.ListDivisionsResponse, error) {
	db := database.FromContext(serviceImpl.ctx)
	academy, err := db.GetAcademy(in.AcademyId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	g := grabber.FromContext(serviceImpl.ctx)

	divisions, err := g.GetDivisions(academy, -1)

	return &pb.ListDivisionsResponse{
		Data: divisions,
	}, nil
}
