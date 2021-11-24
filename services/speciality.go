package services

import (
	"context"
	"grabber/grabber"
	"grabber/models"
	pb "grabber/pb"
)

type SpecialityServiceGrpcImpl struct {
	ctx context.Context
}

func NewSpecialityServiceGrpcImpl(ctx context.Context) *SpecialityServiceGrpcImpl {
	return &SpecialityServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *SpecialityServiceGrpcImpl) GetSpeciality(ctx context.Context, in *pb.GetSpecialityRequest) (*pb.Speciality, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	speciality, _ := g.GetSpecialityById(academy, in.Id)

	return speciality, nil
}

func (serviceImpl *SpecialityServiceGrpcImpl) ListSpecialities(ctx context.Context, _ *pb.ListSpecialitiesRequest) (*pb.ListSpecialitiesResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	specialities, _ := g.GetSpecialities(academy)

	return &pb.ListSpecialitiesResponse{
		Data: specialities,
	}, nil
}

func (serviceImpl *SpecialityServiceGrpcImpl) ListFacultySpecialities(ctx context.Context, in *pb.ListFacultySpecialitiesRequest) (*pb.ListSpecialitiesResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	specialities, _ := g.GetFacultySpecialities(academy, in.FacultyId)

	return &pb.ListSpecialitiesResponse{
		Data: specialities,
	}, nil
}
