package service

import (
	"context"

	"github.com/pkg/errors"
	pb "github.com/tkeel-io/core-broker/api/subscribe/v1"
	"github.com/tkeel-io/core-broker/pkg/deviceutil"
	"github.com/tkeel-io/core-broker/pkg/model"
	"github.com/tkeel-io/core-broker/pkg/pagination"
	"github.com/tkeel-io/core-broker/pkg/subscribeuril"
	"github.com/tkeel-io/kit/log"
	"gorm.io/gorm"
)

const (
	SuccessStatus           = "SUCCESS"
	RepeatedInsertionStatus = "REPEATED INSERTION"
	ErrPartialFailure       = "PARTIAL FAILURE"
)

type SubscribeService struct {
	pb.UnimplementedSubscribeServer
	client *CoreClient
}

func NewSubscribeService() *SubscribeService {
	if err := model.Setup(); err != nil {
		log.Fatal(err)
	}

	return &SubscribeService{client: NewCoreClient()}
}

func (s *SubscribeService) SubscribeEntitiesByIDs(ctx context.Context, req *pb.SubscribeEntitiesByIDsRequest) (*pb.SubscribeEntitiesByIDsResponse, error) {
	// verify Authentication in header and get user token map.
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authUser.ID}
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
	result := model.DB().Preload("Subscribe").Create(&records)
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
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authUser.ID}
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
	ids, err := s.getDeviceEntitiesIDsFromGroups(ctx, req.Groups, authUser.Token)
	if err != nil {
		err = errors.Wrap(err, "get device entities IDs from groups IDs error")
		log.Error("err:", err)
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("no device found")
	}
	records := s.createSubscribeEntitiesRecords(ids, &subscribe)
	log.Info("create subscribe entities records:", records)
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
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authUser.ID}
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
	ids, err := s.getDeviceEntitiesIDsFromTemplates(ctx, req.Models, authUser.Token)
	if err != nil {
		err = errors.Wrap(err, "get device entities IDs from models IDs error")
		log.Error("err:", err)
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("no device found")
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
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authUser.ID}
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
			Subscribe: subscribe,
			EntityID:  entityID,
			UniqueKey: subscribeuril.GenerateSubscribeTopic(subscribe.ID, entityID),
		}
		result := tx.
			Where("subscribe_id = ?", subscribeEntity.Subscribe.ID).
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
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authUser.ID}
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
	page.SetTotal(uint(len(records)))
	err = page.FillResponse(resp)
	if err != nil {
		log.Error("err:", err)
		return nil, err
	}

	entitiesIDs := make([]string, 0, len(records))
	for i := range records {
		entitiesIDs = append(entitiesIDs, records[i].EntityID)
	}

	data, err := s.deviceEntities(entitiesIDs, authUser.Token)
	if err != nil {
		log.Error("err:", err)
		return nil, err
	}

	resp.Data = data

	return resp, nil
}

func (s *SubscribeService) CreateSubscribe(ctx context.Context, req *pb.CreateSubscribeRequest) (*pb.CreateSubscribeResponse, error) {
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	sub := model.Subscribe{
		UserID:      authUser.ID,
		Title:       req.Title,
		Description: req.Description,
	}

	// TODO: lock the table
	var count string
	findResult := model.DB().Model(&model.Subscribe{}).Select("1").
		Where(&model.Subscribe{UserID: authUser.ID, IsDefault: true}).
		Limit(1).
		Find(&count)
	if errors.Is(
		findResult.Error,
		gorm.ErrRecordNotFound,
	) || findResult.RowsAffected == 0 {
		sub.IsDefault = true
	}

	if err = model.DB().Create(&sub).Error; err != nil {
		log.Error("err:", err)
		return nil, err
	}

	return &pb.CreateSubscribeResponse{
		Id:          uint64(sub.ID),
		Title:       sub.Title,
		Description: sub.Description,
		Endpoint:    sub.Endpoint,
		IsDefault:   sub.IsDefault,
	}, nil
}

func (s *SubscribeService) UpdateSubscribe(ctx context.Context, req *pb.UpdateSubscribeRequest) (*pb.UpdateSubscribeResponse, error) {
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authUser.ID}
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
		IsDefault:   subscribe.IsDefault,
	}
	return resp, nil
}

func (s *SubscribeService) DeleteSubscribe(ctx context.Context, req *pb.DeleteSubscribeRequest) (*pb.DeleteSubscribeResponse, error) {
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authUser.ID}
	validateSubscribeResult := model.DB().Model(&subscribe).Where(&subscribe).First(&subscribe)
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
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authUser.ID}
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
		CreatedAt:   subscribe.CreatedAt.Unix(),
		UpdatedAt:   subscribe.UpdatedAt.Unix(),
		IsDefault:   subscribe.IsDefault,
	}
	return resp, nil
}

func (s *SubscribeService) ListSubscribe(ctx context.Context, req *pb.ListSubscribeRequest) (*pb.ListSubscribeResponse, error) {
	authUser, err := s.client.User(ctx)
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
	subscribeCondition := model.Subscribe{UserID: authUser.ID}
	if page.Required() {
		result = model.Paginate(&subscribes, page, &subscribeCondition)
	} else {
		result = model.ListAll(&subscribes, &subscribeCondition)
	}
	if result.Error != nil {
		log.Error("err:", result.Error)
		return nil, result.Error
	}

	data := make([]*pb.SubscribeObject, 0, len(subscribes))
	for i := range subscribes {
		data = append(data, &pb.SubscribeObject{
			Id:          uint64(subscribes[i].ID),
			Title:       subscribes[i].Title,
			Description: subscribes[i].Description,
			Endpoint:    subscribes[i].Endpoint,
			IsDefault:   subscribes[i].IsDefault,
		})
	}

	resp := &pb.ListSubscribeResponse{}
	// for template create default subscribe
	if len(subscribes) == 0 {
		createRequest := &pb.CreateSubscribeRequest{
			Title:       "Default Title",
			Description: "This is default subscribe.",
		}
		subscribeResponse, err := s.CreateSubscribe(ctx, createRequest)
		if err != nil {
			log.Error("create default subscribe failed:", err)
			return nil, err
		}
		data = append(data, &pb.SubscribeObject{
			Id:          subscribeResponse.Id,
			Title:       subscribeResponse.Title,
			Description: subscribeResponse.Description,
			Endpoint:    subscribeResponse.Endpoint,
			IsDefault:   subscribeResponse.IsDefault,
		})
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
	if err = page.FillResponse(resp); err != nil {
		log.Error("err:", err)
		return nil, err
	}

	resp.Data = data

	return resp, nil
}

func (s *SubscribeService) ChangeSubscribed(ctx context.Context, req *pb.ChangeSubscribedRequest) (*pb.ChangeSubscribedResponse, error) {
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	if len(req.SelectedIds) == 0 {
		return nil, errors.New("selectedIds is empty")
	}

	if req.TargetId == 0 {
		return nil, errors.New("targetId is empty")
	}

	subscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.Id)}, UserID: authUser.ID}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}
	targetSubscribe := model.Subscribe{Model: gorm.Model{ID: uint(req.TargetId)}, UserID: authUser.ID}
	validateSubscribeResult = model.DB().First(&targetSubscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user ID mismatch")
		log.Error("err:", err)
		return nil, err
	}

	errs := []error{}
	for i := range req.SelectedIds {
		entityID := req.SelectedIds[i]
		subscribeEntity := model.SubscribeEntities{
			SubscribeID: subscribe.ID,
			Subscribe:   subscribe,
			EntityID:    entityID,
			UniqueKey:   subscribeuril.GenerateSubscribeTopic(subscribe.ID, entityID),
		}
		targetSubscribeEntity := model.SubscribeEntities{
			Subscribe:   targetSubscribe,
			SubscribeID: targetSubscribe.ID,
			EntityID:    entityID,
			UniqueKey:   subscribeuril.GenerateSubscribeTopic(targetSubscribe.ID, entityID),
		}
		if err := model.DB().Debug().Where(targetSubscribeEntity).First(&targetSubscribeEntity).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			errs = append(errs, errors.New("target subscribe entity already exists"))
			continue
		}
		if err = model.DB().Debug().Model(&subscribeEntity).Where(subscribeEntity).Updates(targetSubscribeEntity).Error; err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) == len(req.SelectedIds) && len(errs) > 0 {
		err = errors.Wrap(errs[0], "change subscribed failed")
		log.Error("err:", err)
		return nil, err
	}

	resp := &pb.ChangeSubscribedResponse{Status: SuccessStatus}
	if len(errs) != 0 {
		resp.Status = ErrPartialFailure
	}

	return resp, nil
}

func (s *SubscribeService) ValidateSubscribed(ctx context.Context, req *pb.ValidateSubscribedRequest) (*pb.ValidateSubscribedResponse, error) {
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	if req.Topic == "" {
		return nil, errors.New("topic is empty")
	}

	subscribe := model.Subscribe{Endpoint: req.Topic, UserID: authUser.ID}
	validateSubscribeResult := model.DB().First(&subscribe)
	if validateSubscribeResult.RowsAffected == 0 {
		err = errors.Wrap(validateSubscribeResult.Error, "subscribe and user mismatch")
		log.Error("invalid error:", err)
		return nil, err
	}
	resp := &pb.ValidateSubscribedResponse{Status: SuccessStatus}

	return resp, nil
}

func (s *SubscribeService) SubscribeByDevice(ctx context.Context, req *pb.SubscribeByDeviceRequest) (*pb.SubscribeByDeviceResponse, error) {
	authUser, err := s.client.User(ctx)
	if nil != err {
		log.Error("err:", err)
		return nil, err
	}
	if req.Id == "" {
		return nil, errors.New("invalid device id")
	}
	if req.SubscribeIds == nil || len(req.SubscribeIds) == 0 {
		return nil, errors.New("invalid subscribe ids")
	}

	var count int
	validateSubscribeResult := model.DB().Select("1").
		Where("id IN ?", req.SubscribeIds).
		Where("user_id = ?", authUser).Find(&count)
	if validateSubscribeResult.RowsAffected != int64(len(req.SubscribeIds)) {
		err = errors.Wrap(validateSubscribeResult.Error, "device and user mismatch")
		log.Error("err:", err)
		return nil, err
	}
	subscribeEntities := make([]model.SubscribeEntities, len(req.SubscribeIds))
	for i := range req.SubscribeIds {
		subscribeEntities[i] = model.SubscribeEntities{
			SubscribeID: uint(req.SubscribeIds[i]),
			EntityID:    req.Id,
			UniqueKey:   subscribeuril.GenerateSubscribeTopic(uint(req.SubscribeIds[i]), req.Id),
		}
	}
	if err = model.DB().Debug().Create(&subscribeEntities).Error; err != nil {
		log.Error("create err:", err)
		return nil, err
	}

	resp := &pb.SubscribeByDeviceResponse{Status: SuccessStatus}
	return resp, nil
}

// createSubscribeEntitiesRecords create SubscribeEntities(subscribe_entities table) records.
func (s *SubscribeService) createSubscribeEntitiesRecords(entityIDs []string, subscribe *model.Subscribe) []model.SubscribeEntities {
	records := make([]model.SubscribeEntities, 0, len(entityIDs))
	for _, entityID := range entityIDs {
		subscribeEntity := model.SubscribeEntities{
			Subscribe: *subscribe,
			EntityID:  entityID,
			UniqueKey: subscribeuril.GenerateSubscribeTopic(subscribe.ID, entityID),
		}
		records = append(records, subscribeEntity)
	}
	return records
}

func (s *SubscribeService) getDeviceEntitiesIDsFromGroups(ctx context.Context, groups []string, token string) ([]string, error) {
	var data []string
	dc := deviceutil.NewClient(token)
	for i := range groups {
		bytes, err := dc.Search(deviceutil.DeviceSearch, deviceutil.Conditions{deviceutil.GroupQuery(groups[i]), deviceutil.DeviceTypeQuery()})
		if err != nil {
			log.Error("query device by device group err:", err)
			return nil, err
		}
		log.Info("query device by device group:", string(bytes))
		resp, err := deviceutil.ParseSearchResponse(bytes)
		if err != nil {
			log.Error("parse device search response err:", err)
			return nil, err
		}

		for _, device := range resp.Data.ListDeviceObject.Items {
			data = append(data, device.Id)
		}
	}
	return data, nil
}

func (s *SubscribeService) getDeviceEntitiesIDsFromTemplates(ctx context.Context, templates []string, token string) ([]string, error) {
	var data []string
	dc := deviceutil.NewClient(token)
	for i := range templates {
		bytes, err := dc.Search(deviceutil.DeviceSearch, deviceutil.Conditions{deviceutil.TemplateQuery(templates[i])})
		if err != nil {
			log.Error("query device by device group err:", err)
			return nil, err
		}
		resp, err := deviceutil.ParseSearchResponse(bytes)
		if err != nil {
			log.Error("parse device search response err:", err)
			return nil, err
		}

		for _, device := range resp.Data.ListDeviceObject.Items {
			data = append(data, device.Id)
		}
	}
	return data, nil
}

func (s SubscribeService) deviceEntities(ids []string, token string) ([]*pb.Entity, error) {
	entities := make([]*pb.Entity, 0, len(ids))
	client := deviceutil.NewClient(token)
	for _, id := range ids {
		bytes, err := client.Search(deviceutil.EntitySearch, deviceutil.Conditions{deviceutil.DeviceQuery(id)})
		if err != nil {
			log.Error("query device by device id err:", err)
			return nil, err
		}
		resp, err := deviceutil.ParseSearchEntityResponse(bytes)
		if err != nil {
			log.Error("parse device search response err:", err)
			return nil, err
		}
		if len(resp.Data.Items) == 0 {
			log.Error("device not found:", id)
			return nil, errors.New("device not found")
		}
		entity := &pb.Entity{
			ID:        id,
			Name:      resp.Data.Items[0].Properties.BasicInfo.Name,
			Template:  resp.Data.Items[0].Properties.BasicInfo.TemplateName,
			Group:     resp.Data.Items[0].Properties.BasicInfo.ParentName,
			Status:    resp.Data.Items[0].Properties.SysField.Status,
			UpdatedAt: resp.Data.Items[0].Properties.SysField.UpdatedAt,
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
