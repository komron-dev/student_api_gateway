package services

import (
	"fmt"
	"student_api_gateway/config"

	pbs "student_api_gateway/genproto/student_service"
	pbt "student_api_gateway/genproto/task_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	StudentService() pbs.StudentServiceClient
	TaskService() pbt.TaskServiceClient
}

type serviceManager struct {
	studentService pbs.StudentServiceClient
	taskService pbt.TaskServiceClient
}

func (s *serviceManager) StudentService() pbs.StudentServiceClient {
	return s.studentService
}

func (s *serviceManager) TaskService() pbt.TaskServiceClient {
	return s.taskService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connStudent, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.StudentServiceHost, conf.StudentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.TaskServiceHost, conf.TaskServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	serviceManager := &serviceManager{
		studentService: pbs.NewStudentServiceClient(connStudent),
		taskService: pbt.NewTaskServiceClient(connTask),
	}

	return serviceManager, nil
}
