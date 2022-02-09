package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	pb "github.com/tkeel-io/core-broker/api/subscribe/v1"
	"github.com/tkeel-io/core-broker/pkg/model"
	"github.com/tkeel-io/kit/log"
	"gorm.io/gorm"
)

const (
	SuccessStatus               = "SUCCESS"
	RepeatedInsertionStatus     = "REPEATED INSERTION"
	InvalidRecordDeletionStatus = "INVALID RECORD DELETION"
	FailureStatus               = "FAILURE"
)

type SubscribeService struct {
	pb.UnimplementedSubscribeServer
	client *CoreClient
}

func NewSubscribeService() *SubscribeService {
	return &SubscribeService{client: NewCoreClient()}
}

func (s *SubscribeService) SubscribeEntitiesByIDs(ctx context.Context, req *pb.SubscribeEntitiesByIDsRequest) (*pb.SubscribeEntitiesByIDsResponse, error) {
	// verify Authentication in header and get user token map.
	tokenInfo, err := s.client.GetTokenMap(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: tokenInfo[Owner]}
	if model.DB().First(&subscribe).RowsAffected == 0 {
		err = errors.New("subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	resp := &pb.SubscribeEntitiesByIDsResponse{
		Id:     req.GetId(),
		Status: SuccessStatus,
	}
	if len(req.Entities) == 0 {
		return resp, nil
	}

	records := make([]model.SubscribeEntities, 0, len(req.Entities))
	for _, entityID := range req.Entities {
		subscribeEntity := model.SubscribeEntities{
			SubscribeID: subscribe.ID,
			EntityID:    entityID,
			UniqueKey:   fmt.Sprintf("%d:%s", subscribe.ID, entityID),
		}
		records = append(records, subscribeEntity)
	}
	result := model.DB().Create(&records)
	if result.Error != nil {
		log.Error("err:", result.Error)
		return nil, result.Error
	}
	if result.RowsAffected != int64(len(records)) {
		resp.Status = RepeatedInsertionStatus
	}

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
	// verify Authentication in header and get user token map.
	tokenInfo, err := s.client.GetTokenMap(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: tokenInfo[Owner]}
	if model.DB().First(&subscribe).RowsAffected == 0 {
		err = errors.New("subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	resp := &pb.UnsubscribeEntitiesByIDsResponse{
		Id:     req.Id,
		Status: SuccessStatus,
	}
	records := make([]model.SubscribeEntities, 0, len(req.Entities))
	for _, entityID := range req.Entities {
		subscribeEntity := model.SubscribeEntities{
			SubscribeID: subscribe.ID,
			EntityID:    entityID,
			UniqueKey:   fmt.Sprintf("%d:%s", subscribe.ID, entityID),
		}
		records = append(records, subscribeEntity)
	}

	result := model.DB().Delete(&records)
	if result.Error != nil {
		log.Error("err:", result.Error)
		return nil, result.Error
	}
	if result.RowsAffected != int64(len(records)) {
		resp.Status = InvalidRecordDeletionStatus
	}

	return resp, nil
}
func (s *SubscribeService) ListSubscribeEntities(ctx context.Context, req *pb.ListSubscribeEntitiesRequest) (*pb.ListSubscribeEntitiesResponse, error) {
	fmt.Println("list subscribe id: ", req.Id)
	fmt.Println(req.Page)
	resp := &pb.ListSubscribeEntitiesResponse{}
	return resp, nil
}
func (s *SubscribeService) CreateSubscribe(ctx context.Context, req *pb.CreateSubscribeRequest) (*pb.CreateSubscribeResponse, error) {
	// 1. verify Authentication in header and get user token map.
	tokenInfo, err := s.client.GetTokenMap(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	sub := model.Subscribe{
		UserID:      tokenInfo[Owner],
		Title:       req.Title,
		Description: req.Description,
	}

	resp := &pb.CreateSubscribeResponse{}
	if err = model.DB().Create(&sub).Error; err != nil {
		log.Error("err:", err)
		return nil, err
	}
	resp.Id = uint64(sub.ID)
	resp.Title = req.Title
	resp.Description = req.Description
	resp.Endpoint = sub.Endpoint
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
