package services

import (
	"context"

	"grabber/database"
	"grabber/grabber"
	pb "grabber/pb"

	log "github.com/sirupsen/logrus"
)

type FacultyServiceGrpcImpl struct {
	ctx context.Context
}

func NewFacultyServiceGrpcImpl(ctx context.Context) *FacultyServiceGrpcImpl {
	return &FacultyServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *FacultyServiceGrpcImpl) GetFaculty(_ context.Context, in *pb.GetFacultyRequest) (*pb.Faculty, error) {
	db := database.FromContext(serviceImpl.ctx)
	academy, err := db.GetAcademy(in.AcademyId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	g := grabber.FromContext(serviceImpl.ctx)

	faculty, err := g.GetFacultyById(academy, in.Id)

	return faculty, nil
}

func (serviceImpl *FacultyServiceGrpcImpl) ListFaculties(_ context.Context, in *pb.ListFacultiesRequest) (*pb.ListFacultiesResponse, error) {
	db := database.FromContext(serviceImpl.ctx)
	academy, err := db.GetAcademy(in.AcademyId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	g := grabber.FromContext(serviceImpl.ctx)

	faculties, err := g.GetFaculties(academy, -1)

	return &pb.ListFacultiesResponse{
		Data: faculties,
		Meta: nil,
	}, nil
}
