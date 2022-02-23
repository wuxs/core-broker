package deviceutil

import (
	"encoding/json"
)

type SearchResponse struct {
	Code string
	Data ListDeviceObject
	Msg  string
}

type ListDeviceObject struct {
	ListDeviceObject ListEntity
}

type SearchEntityResponse struct {
	Code string
	Data ListEntity
	Msg  string
}

type ListEntity struct {
	Items    []Object
	PageSize int32
	PageNum  int32
	Total    string
}

type Object struct {
	Config     interface{}
	Id         string
	Mappers    []interface{}
	Owner      string
	Properties Property
	Source     string
	Type       string
}

type Property struct {
	Group     Group     `json:"group,omitempty"`
	BasicInfo BasicInfo `json:"basicInfo"`
	SysField  SysField  `json:"sysField"`
}

type Group struct {
	Description string                 `json:"description"`
	Ext         map[string]interface{} `json:"ext"`
	Name        string                 `json:"name"`
}

type BasicInfo struct {
	Description  string
	Name         string
	TemplateID   string `json:"templateId"`
	TemplateName string `json:"templateName"`
	ParentID     string `json:"parentId"`
	ParentName   string `json:"parentName"`
}

type SysField struct {
	ID        string `json:"_id"`
	Owner     string `json:"_owner"`
	Source    string `json:"_source"`
	Status    string `json:"_status"`
	CreatedAt int64  `json:"_createdAt"`
	UpdatedAt int64  `json:"_updatedAt"`
}

func ParseSearchResponse(bytes []byte) (*SearchResponse, error) {
	var response SearchResponse
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func ParseSearchEntityResponse(bytes []byte) (*SearchEntityResponse, error) {
	var response SearchEntityResponse
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
