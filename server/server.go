package server

import (
	"context"
	"grabber/services"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"grabber/config"
	"grabber/database"
	pb "grabber/pb"
)

var contextKey = "server"

// Server is the gRPC-server operations wrapper
type Server struct {
	ctx    context.Context
	server *grpc.Server
}

func (s *Server) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Error("metadata.FromIncomingContext")
		return nil, nil
	}

	db := database.FromContext(s.ctx)
	ctx = context.WithValue(ctx, "database", db) // pass DB to the context

	switch info.FullMethod {
	case "/grabber.AcademyService/GetAcademy":
	case "/grabber.AcademyService/ListAcademies":
		break
	default:
		if len(md.Get("academy_id")) == 0 {
			log.Error(`Request has no "academy_id" parameter in a request metadata"`)
			return nil, nil
		}

		// Pass an academy instance to the context
		academy, _ := db.GetAcademy(md.Get("academy_id")[0])
		ctx = context.WithValue(ctx, "academy", academy)
	}

	log.WithFields(log.Fields{
		"req": req,
		//"metadata": md,
	}).Infof("Request method: %s", info.FullMethod)

	m, err := handler(ctx, req)
	return m, err
}

// Init initializes the gRPC-server wrapper and create a new value context with the object in it.
func Init(ctx context.Context) context.Context {
	log.Infof(`Init %s...`, contextKey)

	val := &Server{
		ctx: ctx,
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(val.UnaryInterceptor))

	pb.RegisterAcademyServiceServer(s, services.NewAcademyServiceGrpcImpl())
	pb.RegisterCurriculaServiceServer(s, services.NewCurriculaServiceGrpcImpl(ctx))
	pb.RegisterDivisionServiceServer(s, services.NewDivisionServiceGrpcImpl(ctx))
	pb.RegisterFacultyServiceServer(s, services.NewFacultyServiceGrpcImpl(ctx))
	pb.RegisterGroupServiceServer(s, services.NewGroupServiceGrpcImpl(ctx))
	pb.RegisterReportServiceServer(s, services.NewReportServiceGrpcImpl(ctx))
	pb.RegisterSpecialityServiceServer(s, services.NewSpecialityServiceGrpcImpl(ctx))
	pb.RegisterWorkloadServiceServer(s, services.NewWorkloadServiceGrpcImpl(ctx))

	// Register reflection service on gRPC server.
	reflection.Register(s)

	val.server = s

	return context.WithValue(ctx, contextKey, val)
}

// FromContext returns the server wrapper from a given context.
func FromContext(ctx context.Context) *Server {
	db, ok := ctx.Value(contextKey).(*Server)
	if !ok {
		log.Fatal("calling server.FromContext from a non-database context")
	}
	return db
}

func (s *Server) Start() {
	log.Infof(`Starting %s...`, contextKey)
	lis, err := net.Listen("tcp", config.FromContext(s.ctx).Address)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	log.WithContext(s.ctx).Infof("Listen on %s", lis.Addr())
	if err := s.server.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
