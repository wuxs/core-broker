package core

import "fmt"

func PatchEntityURL(entityID string) string {
	return fmt.Sprintf("v1/entities/%s/patch?owner=admin&source=dm", entityID)
}

func QueryDeviceEntityURL(entityID string) string {
	return fmt.Sprintf("v1/entities/%s/properties?type=DEVICE&owner=admin&source=dm", entityID)
}

func CreateEntityURL(entityID, userID, source string) string {
	return fmt.Sprintf("v1/entities?id=%s&owner=%s&source=%s", entityID, userID, source)
}
