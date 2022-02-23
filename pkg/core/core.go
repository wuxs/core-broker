package core

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tkeel-io/core-broker/pkg/subscribeuril"
	"net/http"
	"strconv"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/pkg/errors"
	types "github.com/tkeel-io/core-broker/pkg/types"
)

type Client struct {
	daprClient dapr.Client
}

func NewCoreClient() (*Client, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "init dapr client error")
	}
	return &Client{daprClient: client}, nil
}

type SubscriptionData struct {
	Mode       string `json:"mode,omitempty"`
	Source     string `json:"source,omitempty"`
	Filter     string `json:"filter,omitempty"`
	Topic      string `json:"topic,omitempty"`
	PubsubName string `json:"pubsub_name,omitempty"`
}

func (c *Client) Subscribe(entityID string, topic string) error {
	ctx := context.Background()

	subscriptionID := types.GenerateSubscriptionID(entityID)
	methodName := fmt.Sprintf("v1/subscriptions?id=%s&owner=admin&source=dm&type=SUBSCRIPTION", subscriptionID)
	filter := buildSubscribeQuery(subscriptionID, entityID)
	subscriptionData := SubscriptionData{
		Mode:       "realtime",
		Source:     "ignore",
		Filter:     filter,
		Topic:      types.Topic,
		PubsubName: types.PubsubName,
	}
	if topic != "" {
		subscriptionID = encodeSubscribeID(subscribeuril.GetSubscribeID(topic))
		methodName = fmt.Sprintf("v1/subscriptions?id=%s&owner=admin&source=dm&type=SUBSCRIPTION", subscriptionID)
		filter = buildSubscribeQuery(subscriptionID, entityID)
		subscriptionData = SubscriptionData{
			Mode:       "realtime",
			Source:     "ignore",
			Filter:     filter,
			Topic:      topic,
			PubsubName: types.PubsubName,
		}
	}

	contentData, err := json.Marshal(subscriptionData)
	if err != nil {
		return errors.Wrap(err, "subscriptionData marshal error")
	}

	content := &dapr.DataContent{
		Data:        contentData,
		ContentType: MimeJson,
	}
	c.daprClient.InvokeMethodWithContent(ctx, AppID, methodName, http.MethodPost, content)
	return nil
}

func (c *Client) UnSubscribe(entityID string) error {
	ctx := context.Background()
	subscriptionID := types.GenerateSubscriptionID(entityID)
	methodName := fmt.Sprintf("v1/subscriptions/%s?owner=admin&source=dm&type=SUBSCRIPTION", subscriptionID)
	c.daprClient.InvokeMethod(ctx, AppID, methodName, http.MethodDelete)
	return nil
}

const ql = "insert into %s select %s.*"

func buildSubscribeQuery(to string, from string) string {
	return fmt.Sprintf(ql, to, from)
}

func encodeSubscribeID(id uint) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", id)))
}

func decodeSubscriptionIDToSubscribeID(id string) uint {
	data, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return 0
	}
	t, _ := strconv.ParseUint(string(data), 10, 64)
	return uint(t)
}
