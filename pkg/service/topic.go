package service

import (
	"context"

	pb "github.com/tkeel-io/core-broker/api/topic/v1"
	"github.com/tkeel-io/kit/log"
)

const (
	// SubscriptionResponseStatusSuccess means message is processed successfully.
	SubscriptionResponseStatusSuccess = "SUCCESS"
	// SubscriptionResponseStatusRetry means message to be retried by Dapr.
	SubscriptionResponseStatusRetry = "RETRY"
	// SubscriptionResponseStatusDrop means warning is logged and message is dropped.
	SubscriptionResponseStatusDrop = "DROP"
)

type TopicService struct {
	pb.UnimplementedTopicServer
}

func NewTopicService() *TopicService {
	return &TopicService{}
}

func (s *TopicService) TopicEventHandler(ctx context.Context, req *pb.TopicEventRequest) (*pb.TopicEventResponse, error) {
	MsgChan <- req
	log.Debug("topic event")
	return &pb.TopicEventResponse{Status: SubscriptionResponseStatusSuccess}, nil
}
