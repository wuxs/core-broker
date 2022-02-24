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

package types

import (
	"os"
	"strings"
	"sync"

	"github.com/tkeel-io/core-broker/pkg/util"

	pb "github.com/tkeel-io/core-broker/api/topic/v1"
)

var MsgChan = make(chan *pb.TopicEventRequest, 100)

func Interface2string(in interface{}) (out string) {
	switch inString := in.(type) {
	case string:
		out = inString
	default:
		out = ""
	}
	return
}

var entityMap sync.Map

type WsRequest struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
	Mode string `json:"mode,omitempty"`
}

const PubsubName = "core-broker-pubsub"

var Topic, _ = os.Hostname()

func GenerateSubscriptionID(entityID string) string {
	subID, ok := entityMap.Load(entityID)
	if ok {
		return subID.(string)
	} else {
		subID := entityID + "_" + Topic + util.GenerateRandString(10)
		entityMap.Store(entityID, subID)
		return subID
	}
}

func GetEntityID(subscriptionID string) string {
	return strings.Split(subscriptionID, "_")[0]
}

func DelSubscriptionID(entityID string) {
	entityMap.Delete(entityID)
}
