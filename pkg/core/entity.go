package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
)

func (c Client) PatchEntity(entityID string, data []map[string]interface{}) error {
	ctx := context.Background()
	patchEntityURL := fmt.Sprintf("v1/entities/%s/patch?owner=admin&source=dm", entityID)

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
