package services

import (
	"context"

	"grabber/grabber"
	"grabber/models"
	pb "grabber/pb"
)

type FacultyServiceGrpcImpl struct {
	ctx context.Context
}

func NewFacultyServiceGrpcImpl(ctx context.Context) *FacultyServiceGrpcImpl {
	return &FacultyServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *FacultyServiceGrpcImpl) GetFaculty(ctx context.Context, in *pb.GetFacultyRequest) (*pb.Faculty, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	faculty, _ := g.GetFacultyById(academy, in.Id)

	return faculty, nil
}

func (serviceImpl *FacultyServiceGrpcImpl) ListFaculties(ctx context.Context, _ *pb.ListFacultiesRequest) (*pb.ListFacultiesResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	faculties, _ := g.GetFaculties(academy, -1)

	return &pb.ListFacultiesResponse{
		Data: faculties,
		Meta: nil,
	}, nil
}
