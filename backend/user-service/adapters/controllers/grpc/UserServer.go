package grpc

import (
	"common/gateways/user/protoModel"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
	"net"
)

type UserServer struct {
	port           string
	listener       net.Listener
	userController *UserController
	logger         loggerInterface.Logger
}

func (s *UserServer) Start() error {
	var err error
	if s.listener, err = net.Listen("tcp", s.port); err != nil {
		return errors.NewError("SYS", fmt.Sprintf("failed to listen: %e", err))
	}
	go s.worker()
	return nil
}

func (s *UserServer) worker() {
	grpcServer := grpc.NewServer()
	protoModel.RegisterUserServiceServer(grpcServer, s.userController)
	if err := grpcServer.Serve(s.listener); err != nil {
		s.logger.Error(context.Background(), fmt.Errorf("failed to serve: %w", err))
		return
	}
	grpcServer.Stop()
}
