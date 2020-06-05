package services

import (
	"context"
	"grabber/grabber"
	"grabber/models"
	pb "grabber/pb"
)

type ReportServiceGrpcImpl struct {
	ctx context.Context
}

func NewReportServiceGrpcImpl(ctx context.Context) *ReportServiceGrpcImpl {
	return &ReportServiceGrpcImpl{
		ctx: ctx,
	}
}

func (serviceImpl *ReportServiceGrpcImpl) ListFacultyReports(ctx context.Context, in *pb.ListFacultyReportsRequest) (*pb.ListReportsResponse, error) {
	academy := ctx.Value("academy").(*models.Academy)
	g := grabber.FromContext(serviceImpl.ctx)

	reports, _ := g.GetFacultyReports(academy, in.FacultyId)

	return &pb.ListReportsResponse{
		Data: reports,
	}, nil
}
