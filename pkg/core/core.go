package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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

func (c *Client) Subscribe(entityID string) error {
	appID := "core"
	ctx := context.Background()
	subscriptionID := types.GetSubscriptionID(entityID)
	methodName := fmt.Sprintf("v1/subscriptions?id=%s&owner=admin&source=dm&type=SUBSCRIPTION", subscriptionID)
	filter := fmt.Sprintf("insert into %s select %s.status", subscriptionID, entityID)

	subscriptionData := SubscriptionData{
		Mode:       "realtime",
		Source:     "ignore",
		Filter:     filter,
		Topic:      types.Topic,
		PubsubName: types.PubsubName,
	}

	contentData, err := json.Marshal(subscriptionData)
	if err != nil {
		return errors.Wrap(err, "subscriptionData marshal error")
	}

	content := &dapr.DataContent{
		Data:        contentData,
		ContentType: "application/json",
	}
	c.daprClient.InvokeMethodWithContent(ctx, appID, methodName, http.MethodPost, content)
	return nil
}

func (c *Client) UnSubscribe(entityID string) error {
	appID := "core"
	ctx := context.Background()
	subscriptionID := types.GetSubscriptionID(entityID)
	methodName := fmt.Sprintf("v1/subscriptions/%s?owner=admin&source=dm&type=SUBSCRIPTION", subscriptionID)
	c.daprClient.InvokeMethod(ctx, appID, methodName, http.MethodDelete)
	return nil
}
