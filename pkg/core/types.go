package core

const AppID = "core"
const MimeJson = "application/json"

type GetEntityResponse struct {
	Code string
	Data Entity
	Msg  string
}

type ListEntity struct {
	Items    []Entity
	PageSize int32
	PageNum  int32
	Total    string
}

type Entity struct {
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
	ID            string `json:"_id"`
	Owner         string `json:"_owner"`
	Source        string `json:"_source"`
	Status        string `json:"_status"`
	SubscribeAddr string `json:"_subscribeAddr"`
	CreatedAt     int64  `json:"_createdAt"`
	UpdatedAt     int64  `json:"_updatedAt"`
}
