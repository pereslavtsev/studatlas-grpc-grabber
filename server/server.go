package server

import (
	"context"
	"google.golang.org/grpc/metadata"
	"grabber/config"
	"net"

	pb "grabber/pb"
	"grabber/services"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type contextKeyType string

var contextKey = contextKeyType("server")

// Server is the gRPC-server operations wrapper
type Server struct {
	ctx    context.Context
	server *grpc.Server
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// authentication (token verification)
	_, _ = metadata.FromIncomingContext(ctx)
	log.WithField("req", req).Infof("Request method: %s", info.FullMethod)

	m, err := handler(ctx, req)
	return m, err
}

// Init initializes the gRPC-server wrapper and create a new value context with the object in it.
func Init(ctx context.Context) context.Context {
	log.Infof(`Init %s...`, contextKey)

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))

	pb.RegisterAcademyServiceServer(s, services.NewAcademyServiceGrpcImpl(ctx))
	pb.RegisterDivisionServiceServer(s, services.NewDivisionServiceGrpcImpl(ctx))
	pb.RegisterFacultyServiceServer(s, services.NewFacultyServiceGrpcImpl(ctx))

	// Register reflection service on gRPC server.
	reflection.Register(s)

	return context.WithValue(ctx, contextKey, &Server{
		ctx:    ctx,
		server: s,
	})
}

// FromContext returns the server wrapper from a given context.
func FromContext(ctx context.Context) *Server {
	db, ok := ctx.Value(contextKey).(*Server)
	//log.WithField("ctx", ctx).Debug("ctx")
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
