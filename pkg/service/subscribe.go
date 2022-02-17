package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	pb "github.com/tkeel-io/core-broker/api/subscribe/v1"
	"github.com/tkeel-io/core-broker/pkg/model"
	"github.com/tkeel-io/core-broker/pkg/pagination"
	"github.com/tkeel-io/core-broker/pkg/util"
	"github.com/tkeel-io/kit/log"
	"gorm.io/gorm"
)

const (
	SuccessStatus           = "SUCCESS"
	FailedStatus            = "FAILED"
	RepeatedInsertionStatus = "REPEATED INSERTION"
	ErrPartialFailure       = "PARTIAL FAILURE"

	subscribe   = true
	unsubscribe = false
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
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authorizedUser.ID}
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
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authorizedUser.ID}
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
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authorizedUser.ID}
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
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authorizedUser.ID}
	if model.DB().First(&subscribe).RowsAffected == 0 {
		err = errors.New("subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	resp := &pb.UnsubscribeEntitiesByIDsResponse{
		Id:     req.Id,
		Status: SuccessStatus,
	}

	tx := model.DB().Begin()
	for _, entityID := range req.Entities {
		subscribeEntity := model.SubscribeEntities{
			SubscribeID: subscribe.ID,
			EntityID:    entityID,
			UniqueKey:   fmt.Sprintf("%d:%s", subscribe.ID, entityID),
		}
		result := tx.
			Where("subscribe_id = ?", subscribeEntity.SubscribeID).
			Where("entity_id = ?", subscribeEntity.EntityID).
			Where("unique_key = ?", subscribeEntity.UniqueKey).
			Delete(&subscribeEntity)
		if result.Error != nil {
			log.Error("err:", result.Error)
			tx.Rollback()
			return nil, result.Error
		}
	}
	tx.Commit()

	return resp, nil
}

func (s *SubscribeService) ListSubscribeEntities(ctx context.Context, req *pb.ListSubscribeEntitiesRequest) (*pb.ListSubscribeEntitiesResponse, error) {
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authorizedUser.ID}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	page, err := pagination.Parse(req)
	if err != nil {
		log.Error("err:", err)
		return nil, err
	}

	var records []model.SubscribeEntities
	result := model.Paginate(&records, page, model.SubscribeEntities{SubscribeID: subscribe.ID})
	if result.Error != nil {
		log.Error("err:", result.Error)
		return nil, err
	}

	resp := &pb.ListSubscribeEntitiesResponse{}
	page.SetTotal(uint(len(cache)))
	err = page.FillResponse(resp)
	if err != nil {
		log.Error("err:", err)
		return nil, err
	}

	entitiesIDs := make([]string, 0, len(records))
	for i := range records {
		entitiesIDs = append(entitiesIDs, records[i].EntityID)
	}

	data, err := s.deviceEntities(entitiesIDs)
	if err != nil {
		log.Error("err:", err)
		return nil, err
	}

	resp.Data = data

	return resp, nil
}

func (s *SubscribeService) CreateSubscribe(ctx context.Context, req *pb.CreateSubscribeRequest) (*pb.CreateSubscribeResponse, error) {
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	sub := model.Subscribe{
		UserID:      authorizedUser.ID,
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
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authorizedUser.ID, TenantID: authorizedUser.TenantID}
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
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authorizedUser.ID, TenantID: authorizedUser.TenantID}
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
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}}
	validateSubscribeResult := model.DB().
		Where(&subscribe).
		Or("user_id = ?", authorizedUser.ID).
		Or("tenant_id = ?", authorizedUser.TenantID).
		First(&subscribe)
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
		CreatedAt:   subscribe.CreatedAt.Unix(),
		UpdatedAt:   subscribe.UpdatedAt.Unix(),
	}
	return resp, nil
}

func (s *SubscribeService) ListSubscribe(ctx context.Context, req *pb.ListSubscribeRequest) (*pb.ListSubscribeResponse, error) {
	authorizedUser, err := s.client.User(ctx)
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
	subscribeCondition := model.Subscribe{TenantID: authorizedUser.TenantID}
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
	if err = model.Count(&count, &subscribeCondition, &subscribeCondition).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("err:", err)
			return nil, err
		}
		count = 0
	}
	page.SetTotal(uint(count))
	resp := &pb.ListSubscribeResponse{}
	if err = page.FillResponse(resp); err != nil {
		log.Error("err:", err)
		return nil, err
	}

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

func (s *SubscribeService) MySubscribed(ctx context.Context, req *pb.MySubscribedRequest) (*pb.MySubscribedResponse, error) {
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}

	page, err := pagination.Parse(req)
	if err != nil {
		log.Error("parse request page info error:", err)
		return nil, err
	}

	var subscribesIDs []uint
	if err = model.DB().Model(&model.SubscribeEntities{}).
		Where("user_id = ?", authorizedUser.ID).
		Pluck("subscribe_id", &subscribesIDs).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("err:", err)
			return nil, err
		}
	}

	var subscribes []model.Subscribe
	if page.Required() {
		result := model.Paginate(&subscribes, page, subscribesIDs)
		if result.Error != nil {
			log.Error("err:", result.Error)
			return nil, result.Error
		}
	} else {
		result := model.ListAll(&subscribes, subscribesIDs)
		if result.Error != nil {
			log.Error("err:", result.Error)
			return nil, result.Error
		}
	}

	resp := &pb.MySubscribedResponse{}
	if err = page.FillResponse(resp); err != nil {
		log.Error("err:", err)
		return nil, err
	}

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

func (s *SubscribeService) Subscribe(ctx context.Context, req *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}

	resp := &pb.SubscribeResponse{Status: FailedStatus}
	condition := model.SubscribeUsers{SubscribeID: uint(req.Id), UserID: authorizedUser.ID}
	switch req.Option {
	case subscribe:
		if err = model.DB().Where(&condition).FirstOrCreate(&condition).Error; err != nil {
			log.Error("err:", err)
			return nil, err
		}
		resp.Status = SuccessStatus
	case unsubscribe:
		if err = model.DB().Where(&condition).Delete(&condition).Error; err != nil {
			log.Error("err:", err)
			return nil, err
		}
		resp.Status = SuccessStatus
	}

	return resp, nil
}

func (s *SubscribeService) ChangeSubscribed(ctx context.Context, req *pb.ChangeSubscribedRequest) (*pb.ChangeSubscribedResponse, error) {
	authorizedUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authorizedUser.ID, TenantID: authorizedUser.TenantID}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}
	targetSubscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.TargetID)}, UserID: authorizedUser.ID, TenantID: authorizedUser.TenantID}
	validateSubscribeResult = model.DB().First(&targetSubscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	errs := []error{}
	for i := range req.SelectedIDs {
		entityID := req.SelectedIDs[i]
		subscribeEntity := model.SubscribeEntities{
			SubscribeID: subscribe.ID,
			EntityID:    entityID,
			UniqueKey:   fmt.Sprintf("%d:%s", subscribe.ID, entityID),
		}
		targetSubscribeEntity := model.SubscribeEntities{
			SubscribeID: targetSubscribe.ID,
			EntityID:    entityID,
			UniqueKey:   fmt.Sprintf("%d:%s", targetSubscribe.ID, entityID),
		}
		if err = model.DB().Model(&subscribeEntity).Updates(targetSubscribeEntity).Error; err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) == len(req.SelectedIDs) {
		err = errors.Wrap(errs[0], "change subscribed failed")
		log.Error("err:", err)
		return nil, err
	}

	resp := &pb.ChangeSubscribedResponse{Status: SuccessStatus}
	if len(errs) < len(req.SelectedIDs) {
		resp.Status = ErrPartialFailure
	}

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

var cache = map[string]*pb.Entity{}

// TODO: implement getDeviceEntitiesIDsFromGroups
func (s *SubscribeService) getDeviceEntitiesIDsFromGroups(ctx context.Context, groups []string) ([]string, error) {
	var data []string
	for i := range groups {
		device := pb.Entity{}
		device.ID = util.GenerateRandString(10)
		device.Name = util.GenerateRandString(5)
		device.Group = groups[i]
		cache[device.ID] = &device
		data = append(data, device.ID)
	}
	return data, nil
}

// TODO: implement getDeviceEntitiesIDsFromModels
func (s *SubscribeService) getDeviceEntitiesIDsFromModels(ctx context.Context, models []string) ([]string, error) {
	var data []string
	for i := range models {
		device := pb.Entity{}
		device.ID = util.GenerateRandString(10)
		device.Name = util.GenerateRandString(5)
		device.Template = models[i]
		cache[device.ID] = &device
		data = append(data, device.ID)
	}
	return data, nil
}

// TODO: implement deviceEntities
func (s SubscribeService) deviceEntities(ids []string) ([]*pb.Entity, error) {
	entities := make([]*pb.Entity, 0, len(ids))
	for _, id := range ids {
		entity := cache[id]
		entities = append(entities, entity)
	}
	return entities, nil
}
