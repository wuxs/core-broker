package service

import pb "github.com/tkeel-io/core-broker/api/topic/v1"

var MsgChan chan *pb.TopicEventRequest

func init() {
	MsgChan = make(chan *pb.TopicEventRequest, 100)
}

func interface2string(in interface{}) (out string) {
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
