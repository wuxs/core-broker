package core

import "fmt"

func PatchEntityURL(entityID string) string {
	return fmt.Sprintf("v1/entities/%s/patch?owner=admin&source=dm", entityID)
}

func QueryEntityURL(entityID string) string {
	return fmt.Sprintf("v1/entities/%s/properties?owner=admin&source=dm", entityID)
}
