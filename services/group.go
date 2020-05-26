package services

import (
	"context"
	"grabber/grabber"
	"grabber/models"

	pb "grabber/pb"
)

type GroupServiceGrpcImpl struct {
	ctx context.Context
}

func NewGroupServiceGrpcImpl(ctx context.Context) *GroupServiceGrpcImpl {
	return &GroupServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *GroupServiceGrpcImpl) GetGroup(ctx context.Context, in *pb.GetGroupRequest) (*pb.Group, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	group, _ := g.GetGroupById(academy, in.Id)

	return group, nil
}

func (serviceImpl *GroupServiceGrpcImpl) ListGroups(ctx context.Context, _ *pb.ListGroupsRequest) (*pb.ListGroupsResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	groups, _ := g.GetGroups(academy)

	return &pb.ListGroupsResponse{
		Data: groups,
	}, nil
}

func (serviceImpl *GroupServiceGrpcImpl) ListFacultyGroups(ctx context.Context, in *pb.ListFacultyGroupsRequest) (*pb.ListGroupsResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	groups, _ := g.GetFacultyGroups(academy, in.FacultyId)

	return &pb.ListGroupsResponse{
		Data: groups,
	}, nil
}

func (serviceImpl *GroupServiceGrpcImpl) ListSpecialityGroups(ctx context.Context, in *pb.ListSpecialityGroupsRequest) (*pb.ListGroupsResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	groups, _ := g.GetSpecialityGroups(academy, in.SpecialityId)

	return &pb.ListGroupsResponse{
		Data: groups,
	}, nil
}
