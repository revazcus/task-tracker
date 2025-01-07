package gateways

import (
	"context"
	"fmt"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	"google.golang.org/grpc"
)

type BaseGRPCGateway struct {
	connection *grpc.ClientConn
	url        string
	logger     loggerInterface.Logger
}

// NewBaseGRPCGateway TODO переписать на билдер
func NewBaseGRPCGateway(url string, logger loggerInterface.Logger) *BaseGRPCGateway {
	return &BaseGRPCGateway{
		url:    url,
		logger: logger,
	}
}

func (b *BaseGRPCGateway) Connection() *grpc.ClientConn {
	return b.connection
}

func (g *BaseGRPCGateway) Start() error {
	go g.worker()
	return nil
}

func (g *BaseGRPCGateway) worker() {
	if err := g.dial(); err != nil {
		g.logger.Error(context.Background(), fmt.Errorf("failed to dial: %e", err))
		return
	}
}

func (g *BaseGRPCGateway) dial() error {
	// TODO переписать
	connection, err := grpc.Dial(g.url, grpc.WithInsecure())
	if err != nil {
		return errors.NewError("SYS", fmt.Sprintf("failed to dial: %e", err))
	}
	g.connection = connection
	return nil
}
