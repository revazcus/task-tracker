package grpc

import (
	"context"
	"fmt"
	"github.com/revazcus/task-tracker/backend/common/gateways/user1/gateways/user/protoModel"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	"google.golang.org/grpc"
	"net"
)

type UserServer struct {
	Port           string
	listener       net.Listener
	userController *UserController
	logger         loggerInterface.Logger
}

// NewUserServer TODO переписать на билдер
func NewUserServer(port string, userController *UserController, logger loggerInterface.Logger) *UserServer {
	return &UserServer{
		Port:           port,
		userController: userController,
		logger:         logger,
	}
}

func (s *UserServer) Start() error {
	var err error
	if s.listener, err = net.Listen("tcp", s.Port); err != nil {
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
