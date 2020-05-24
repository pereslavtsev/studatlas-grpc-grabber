package services

import (
	"context"

	"grabber/database"
	m "grabber/models"
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

func serializeAcademy(res *m.Academy) *pb.Academy {
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
	}
}

func (serviceImpl *AcademyServiceGrpcImpl) GetAcademy(ctx context.Context, in *pb.GetAcademyRequest) (*pb.Academy, error) {
	db := database.FromContext(serviceImpl.ctx)

	res, err := db.GetAcademy(in.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return serializeAcademy(res), nil
}

func (serviceImpl *AcademyServiceGrpcImpl) ListAcademies(ctx context.Context, in *pb.ListAcademiesRequest) (*pb.ListAcademiesResponse, error) {
	db := database.FromContext(serviceImpl.ctx)
	res, count, err := db.GetAcademies()

	if count == 0 {
		log.Warn("No academies founded")
	}

	if err != nil {
		log.Error(err)
		return nil, err
	}

	var academies []*pb.Academy

	for _, t := range res {
		academies = append(academies, serializeAcademy(t))
	}
	return &pb.ListAcademiesResponse{
		Data: academies,
	}, nil
}
