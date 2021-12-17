package service

import (
	"context"
	"fmt"
	"os"

	pb "github.com/tkeel-io/core-broker/api/dapr"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SubscribeService struct {
	pb.UnimplementedSubscribeServer
}

func NewSubscribeService() *SubscribeService {
	return &SubscribeService{}
}

func (s *SubscribeService) GetSubscribe(ctx context.Context, req *emptypb.Empty) (*pb.ListTopicSubscriptionsResponse, error) {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		hostName = "abc"
	}
	resp := &pb.ListTopicSubscriptionsResponse{}
	resp.Subscriptions = append(resp.Subscriptions, &pb.TopicSubscription{
		PubsubName: "core-broker-pubsub",
		Topic:      hostName,
		Metadata:   map[string]string{},
		Route:      "/v1/topic",
	})

	return resp, nil
}
