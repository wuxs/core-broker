package service

import (
	"context"
	"fmt"

	pb "github.com/tkeel-io/core-broker/api/subscribe/v1"
)

type SubscribeService struct {
	pb.UnimplementedSubscribeServer
}

func NewSubscribeService() *SubscribeService {
	return &SubscribeService{}
}

func (s *SubscribeService) SubscribeEntitiesByIDs(ctx context.Context, req *pb.SubscribeEntitiesByIDsRequest) (*pb.SubscribeEntitiesByIDsResponse, error) {
	fmt.Println("entities: ", req.Entities)
	resp := &pb.SubscribeEntitiesByIDsResponse{}
	resp.Id = req.Id
	return resp, nil
}
func (s *SubscribeService) SubscribeEntitiesByGroups(ctx context.Context, req *pb.SubscribeEntitiesByGroupsRequest) (*pb.SubscribeEntitiesByGroupsResponse, error) {
	fmt.Println("groups: ", req.Groups)
	resp := &pb.SubscribeEntitiesByGroupsResponse{}
	resp.Id = req.Id
	return resp, nil
}
func (s *SubscribeService) SubscribeEntitiesByModels(ctx context.Context, req *pb.SubscribeEntitiesByModelsRequest) (*pb.SubscribeEntitiesByModelsResponse, error) {
	fmt.Println("models: ", req.Models)
	resp := &pb.SubscribeEntitiesByModelsResponse{}
	resp.Id = req.Id
	return resp, nil
}
func (s *SubscribeService) UnsubscribeEntitiesByIDs(ctx context.Context, req *pb.UnsubscribeEntitiesByIDsRequest) (*pb.UnsubscribeEntitiesByIDsResponse, error) {
	fmt.Println("entities: ", req.Entities)
	resp := &pb.UnsubscribeEntitiesByIDsResponse{}
	return resp, nil
}
func (s *SubscribeService) ListSubscribeEntities(ctx context.Context, req *pb.ListSubscribeEntitiesRequest) (*pb.ListSubscribeEntitiesResponse, error) {
	fmt.Println("list subscribe id: ", req.Id)
	fmt.Println(req.Page)
	resp := &pb.ListSubscribeEntitiesResponse{}
	return resp, nil
}
func (s *SubscribeService) CreateSubscribe(ctx context.Context, req *pb.CreateSubscribeRequest) (*pb.CreateSubscribeResponse, error) {
	fmt.Println(req.Name)
	fmt.Println(req.Description)
	resp := &pb.CreateSubscribeResponse{}
	resp.Id = "sub1234"
	resp.Name = req.Name
	resp.Description = req.Description
	resp.Endpoint = "amqp://xxxx"
	return resp, nil
}
func (s *SubscribeService) UpdateSubscribe(ctx context.Context, req *pb.UpdateSubscribeRequest) (*pb.UpdateSubscribeResponse, error) {
	fmt.Println(req.Name)
	fmt.Println(req.Description)
	resp := &pb.UpdateSubscribeResponse{}
	resp.Id = req.Id
	resp.Name = req.Name
	resp.Description = req.Description
	resp.Endpoint = "amqp://xxxx"
	return resp, nil
}
func (s *SubscribeService) DeleteSubscribe(ctx context.Context, req *pb.DeleteSubscribeRequest) (*pb.DeleteSubscribeResponse, error) {
	fmt.Println("id:", req.Id)
	return &pb.DeleteSubscribeResponse{}, nil
}
func (s *SubscribeService) GetSubscribe(ctx context.Context, req *pb.GetSubscribeRequest) (*pb.GetSubscribeResponse, error) {
	fmt.Println(req.Id)
	resp := &pb.GetSubscribeResponse{}
	return resp, nil
}
func (s *SubscribeService) ListSubscribe(ctx context.Context, req *pb.ListSubscribeRequest) (*pb.ListSubscribeResponse, error) {
	return &pb.ListSubscribeResponse{}, nil
}
