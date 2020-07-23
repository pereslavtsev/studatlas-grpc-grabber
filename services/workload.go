package services

import (
	"context"
	"grabber/grabber"
	"grabber/models"
	pb "grabber/pb"
)

type WorkloadServiceGrpcImpl struct {
	ctx context.Context
}

func NewWorkloadServiceGrpcImpl(ctx context.Context) *WorkloadServiceGrpcImpl {
	return &WorkloadServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *WorkloadServiceGrpcImpl) ListGroupsWorkloads(ctx context.Context, in *pb.ListGroupsWorkloadsRequest) (*pb.ListGroupsWorkloadsResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	data, _ := g.GetGroupsWorkloads(academy, in)

	return &pb.ListGroupsWorkloadsResponse{
		Data: data,
	}, nil
}

func (serviceImpl *WorkloadServiceGrpcImpl) ListTeachersWorkloads(ctx context.Context, in *pb.ListTeachersWorkloadsRequest) (*pb.ListTeachersWorkloadsResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	data, _ := g.GetTeachersWorkloads(academy, in)

	return &pb.ListTeachersWorkloadsResponse{
		Data: data,
	}, nil
}

func (serviceImpl *WorkloadServiceGrpcImpl) GetGroupWorkload(ctx context.Context, in *pb.GetGroupWorkloadRequest) (*pb.GetGroupWorkloadResponse, error) {
	//_ := ctx.Value("academy").(*models.Academy)
	// := grabber.FromContext(serviceImpl.ctx)

	//, _ := g.GetTeachersWorkloads(academy, in)

	return &pb.GetGroupWorkloadResponse{
		Data: nil,
	}, nil
}
