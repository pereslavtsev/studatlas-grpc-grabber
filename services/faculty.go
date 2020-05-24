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

func (serviceImpl *FacultyServiceGrpcImpl) GetFaculty(_ context.Context, _ *pb.GetFacultyRequest) (*pb.Faculty, error) {
	var faculties *pb.Faculty
	return faculties, nil
}

func (serviceImpl *FacultyServiceGrpcImpl) ListFaculties(_ context.Context, in *pb.ListFacultiesRequest) (*pb.ListFacultiesResponse, error) {
	db := database.FromContext(serviceImpl.ctx)
	academy, err := db.GetAcademy(in.AcademyId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	g := grabber.FromContext(serviceImpl.ctx)

	faculties, err := g.GetFaculties(academy)

	return &pb.ListFacultiesResponse{
		Data: faculties,
		Meta: nil,
	}, nil
}
