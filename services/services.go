package services

import (
	"fmt"

	"github.com/muhriddinsalohiddin/api-gateway/config"
	pb "github.com/muhriddinsalohiddin/api-gateway/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	TaskService() pb.TaskServiceClient
}

type serviceManager struct {
	taskService pb.TaskServiceClient
}

func (s *serviceManager) TaskService() pb.TaskServiceClient {
	return s.taskService
}
func NewServiceManager(cfg *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.TaskServiceHost, cfg.TaskServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	serviceManager := &serviceManager{
		taskService: pb.NewTaskServiceClient(connTask),
	}
	return serviceManager, nil
}
