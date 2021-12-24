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

	pb "github.com/tkeel-io/core-broker/api/topic/v1"
)

var MsgChan chan *pb.TopicEventRequest

func init() {
	MsgChan = make(chan *pb.TopicEventRequest, 100)
}

func Interface2string(in interface{}) (out string) {
	switch inString := in.(type) {
	case string:
		out = inString
	default:
		out = ""
	}
	return
}

type WsRequest struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
	Mode string `json:"mode,omitempty"`
}

const PubsubName = "core-broker-pubsub"

var Topic, _ = os.Hostname()

func GetSubscriptionID(entityID string) string {
	return entityID + "_" + Topic
}

func GetEntityID(subscriptionID string) string {
	return strings.Split(subscriptionID, "_")[0]
}
