package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	pb "github.com/tkeel-io/core-broker/api/subscribe/v1"
	"github.com/tkeel-io/core-broker/pkg/model"
	"github.com/tkeel-io/core-broker/pkg/pagination"
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
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
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

	records := s.createSubscribeEntitiesRecords(req.Entities, &subscribe)
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
	tokenInfo, err := s.client.GetTokenMap(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: tokenInfo[Owner]}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	resp := &pb.SubscribeEntitiesByGroupsResponse{
		Id:     req.Id,
		Status: SuccessStatus,
	}
	ids, err := s.getDeviceEntitiesIDsFromGroups(ctx, req.Groups)
	if err != nil {
		err = errors.Wrap(err, "get device entities IDs from groups IDs error")
		log.Error("err:", err)
		return nil, err
	}
	records := s.createSubscribeEntitiesRecords(ids, &subscribe)
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
func (s *SubscribeService) SubscribeEntitiesByModels(ctx context.Context, req *pb.SubscribeEntitiesByModelsRequest) (*pb.SubscribeEntitiesByModelsResponse, error) {
	tokenInfo, err := s.client.GetTokenMap(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: tokenInfo[Owner]}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	resp := &pb.SubscribeEntitiesByModelsResponse{
		Id:     req.Id,
		Status: SuccessStatus,
	}
	ids, err := s.getDeviceEntitiesIDsFromModels(ctx, req.Models)
	if err != nil {
		err = errors.Wrap(err, "get device entities IDs from models IDs error")
		log.Error("err:", err)
		return nil, err
	}
	records := s.createSubscribeEntitiesRecords(ids, &subscribe)
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
	// TODO: implement
	fmt.Println("list subscribe id: ", req.Id)
	fmt.Println(req.PageNum, req.PageSize)
	resp := &pb.ListSubscribeEntitiesResponse{}
	return resp, nil
}
func (s *SubscribeService) CreateSubscribe(ctx context.Context, req *pb.CreateSubscribeRequest) (*pb.CreateSubscribeResponse, error) {
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
	tokenInfo, err := s.client.GetTokenMap(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: tokenInfo[Owner]}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	if req.Title != "" {
		subscribe.Title = req.Title
	}
	if req.Description != "" {
		subscribe.Description = req.Description
	}
	if err = model.DB().Save(&subscribe).Error; err != nil {
		err = errors.Wrap(err, "update subscribe info err")
		log.Error("err:", err)
		return nil, err
	}

	resp := &pb.UpdateSubscribeResponse{
		Id:          uint64(subscribe.ID),
		Title:       subscribe.Title,
		Description: subscribe.Description,
		Endpoint:    subscribe.Endpoint,
	}
	return resp, nil
}
func (s *SubscribeService) DeleteSubscribe(ctx context.Context, req *pb.DeleteSubscribeRequest) (*pb.DeleteSubscribeResponse, error) {
	tokenInfo, err := s.client.GetTokenMap(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: tokenInfo[Owner]}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	if err = model.DB().Delete(&subscribe).Error; err != nil {
		err = errors.Wrap(err, "delete subscribe err")
		log.Error("err:", err)
		return nil, err
	}

	return &pb.DeleteSubscribeResponse{Id: req.Id}, nil
}
func (s *SubscribeService) GetSubscribe(ctx context.Context, req *pb.GetSubscribeRequest) (*pb.GetSubscribeResponse, error) {
	tokenInfo, err := s.client.GetTokenMap(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: tokenInfo[Owner]}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	var count int64
	model.DB().Model(&model.SubscribeEntities{}).Where("subscribe_id = ?", subscribe.ID).Count(&count)

	resp := &pb.GetSubscribeResponse{
		Id:          uint64(subscribe.ID),
		Title:       subscribe.Title,
		Description: subscribe.Description,
		Endpoint:    subscribe.Endpoint,
		Count:       uint64(count),
	}
	return resp, nil
}
func (s *SubscribeService) ListSubscribe(ctx context.Context, req *pb.ListSubscribeRequest) (*pb.ListSubscribeResponse, error) {
	tokenInfo, err := s.client.GetTokenMap(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	page, err := pagination.Parse(req)
	if err != nil {
		log.Error("parse request page info error:", err)
		return nil, err
	}
	var subscribes []model.Subscribe
	result := &gorm.DB{Error: errors.New("db query error")}
	subscribeCondition := model.Subscribe{UserID: tokenInfo[Owner]}
	if page.Required() {
		result = model.Paginate(&subscribes, page, &subscribeCondition)
	} else {
		result = model.ListAll(&subscribes, &subscribeCondition)
	}
	if result.Error != nil {
		log.Error("err:", result.Error)
		return nil, result.Error
	}

	var count int64
	if err = model.Count(&count, &subscribeCondition).Error; err != nil {
		log.Error("err:", err)
		return nil, err
	}
	resp := &pb.ListSubscribeResponse{}
	page.FillResponse(resp, count)

	data := make([]*pb.SubscribeObject, 0, len(subscribes))
	for i := range subscribes {
		data = append(data, &pb.SubscribeObject{
			Id:          uint64(subscribes[i].ID),
			Title:       subscribes[i].Title,
			Description: subscribes[i].Description,
			Endpoint:    subscribes[i].Endpoint,
		})
	}
	resp.Data = data

	return resp, nil
}

// createSubscribeEntitiesRecords create SubscribeEntities(subscribe_entities table) records.
func (s *SubscribeService) createSubscribeEntitiesRecords(entityIDs []string, subscribe *model.Subscribe) []model.SubscribeEntities {
	records := make([]model.SubscribeEntities, 0, len(entityIDs))
	for _, entityID := range entityIDs {
		subscribeEntity := model.SubscribeEntities{
			SubscribeID: subscribe.ID,
			EntityID:    entityID,
			UniqueKey:   fmt.Sprintf("%d:%s", subscribe.ID, entityID),
		}
		records = append(records, subscribeEntity)
	}
	return records
}

// TODO: implement getDeviceEntitiesIDsFromGroups
func (s *SubscribeService) getDeviceEntitiesIDsFromGroups(ctx context.Context, groups []string) ([]string, error) {
	panic("implement me")
}

// TODO: implement getDeviceEntitiesIDsFromModels
func (s *SubscribeService) getDeviceEntitiesIDsFromModels(ctx context.Context, models []string) ([]string, error) {
	panic("implement me")
}
