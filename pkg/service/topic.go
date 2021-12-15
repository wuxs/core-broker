/*
Copyright 2021 The tKeel Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
