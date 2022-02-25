package core

import (
	"context"
	"encoding/json"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
)

func (c Client) PatchEntity(entityID string, data []map[string]interface{}) error {
	ctx := context.Background()
	patchEntityURL := PatchEntityURL(entityID)

	contentData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	content := &dapr.DataContent{
		Data:        contentData,
		ContentType: MimeJson,
	}
	_, err = c.daprClient.InvokeMethodWithContent(ctx, AppID, patchEntityURL, http.MethodPut, content)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) GetEntity(entityID string) (*Entity, error) {
	ctx := context.Background()
	queryEntityURL := QueryEntityURL(entityID)

	resp, err := c.daprClient.InvokeMethod(ctx, AppID, queryEntityURL, http.MethodGet)
	if err != nil {
		return nil, err
	}

	response := &GetEntityResponse{}
	if err = json.Unmarshal(resp, response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
