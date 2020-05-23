package services

import (
	"context"

	"grabber/database"
	pb "grabber/pb"

	log "github.com/sirupsen/logrus"
)

type AcademyServiceGrpcImpl struct {
	ctx context.Context
}

func NewAcademyServiceGrpcImpl(ctx context.Context) *AcademyServiceGrpcImpl {
	return &AcademyServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *AcademyServiceGrpcImpl) GetAcademy(ctx context.Context, in *pb.GetAcademyRequest) (*pb.Academy, error) {
	db := database.FromContext(serviceImpl.ctx)
	res, _ := db.GetAcademy(in.Id)

	log.Debug("Received academy", res.ID)

	isDisabled := 0

	if res.IsDisabled {
		isDisabled = 1
	}

	return &pb.Academy{
		Id:              res.ID.Hex(),
		Name:            res.Name,
		Abbreviation:    res.Abbreviation,
		Alias:           res.Alias,
		Website:         res.Website,
		Endpoint:        res.Endpoint,
		Version:         res.Version,
		DisabledSources: res.DisabledSources,
		IsDisabled:      int32(isDisabled),
	}, nil
}

func (serviceImpl *AcademyServiceGrpcImpl) ListAcademies(ctx context.Context, in *pb.ListAcademiesRequest) (*pb.ListAcademiesResponse, error) {
	db := database.FromContext(serviceImpl.ctx)
	academies, _, _ := db.GetAcademies()

	var res []*pb.Academy

	for _, t := range academies {
		isDisabled := 0

		if t.IsDisabled {
			isDisabled = 1
		}

		res = append(res, &pb.Academy{
			Id:              t.ID.Hex(),
			Name:            t.Name,
			Abbreviation:    t.Abbreviation,
			Alias:           t.Alias,
			Website:         t.Website,
			Endpoint:        t.Endpoint,
			Version:         t.Version,
			DisabledSources: t.DisabledSources,
			IsDisabled:      int32(isDisabled),
		})
	}
	return &pb.ListAcademiesResponse{
		Data: res,
	}, nil
}
