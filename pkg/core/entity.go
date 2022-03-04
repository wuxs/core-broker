package core

import (
	"context"
	"encoding/json"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/tkeel-io/kit/log"
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

	log.Infof("invoke patch entity %s \n By %s \n Content.Data:%v", patchEntityURL, http.MethodPut, string(content.Data))
	if re, err := c.daprClient.InvokeMethodWithContent(ctx, AppID, patchEntityURL, http.MethodPut, content); err != nil {
		log.Errorf("invoke %s \n and Request Body:%v \n Response Content: %s \n err:%v", patchEntityURL, content, string(re), err)
		return err
	}
	return nil
}

func (c Client) GetDeviceEntity(entityID string) (*Entity, error) {
	ctx := context.Background()
	queryEntityURL := QueryDeviceEntityURL(entityID)

	log.Debugf("invoke get device entity %s", queryEntityURL)
	resp, err := c.daprClient.InvokeMethod(ctx, AppID, queryEntityURL, http.MethodGet)
	if err != nil {
		log.Errorf("invoke %s \n response content: %s \n err:%v", queryEntityURL, string(resp), err)
		return nil, err
	}

	response := &GetEntityResponse{}
	if err = json.Unmarshal(resp, response); err != nil {
		log.Errorf("unmarshal response content: %s \n err:%v", string(resp), err)
		return nil, err
	}

	log.Debug("get entity response raw:", string(resp))
	log.Debug("get entity:", response)

	return &response.Data, nil
}
