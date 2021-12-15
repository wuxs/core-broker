/*
Copyright 2021 The tKeel Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"context"

	v1 "github.com/tkeel-io/core-broker/api/openapi/v1"
	"github.com/tkeel-io/core-broker/pkg/util"
	openapi_v1 "github.com/tkeel-io/tkeel-interface/openapi/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// OpenapiService is a openapi service.
type OpenapiService struct {
	v1.UnimplementedOpenapiServer
}

// NewOpenapiService new a openapi service.
func NewOpenapiService() *OpenapiService {
	return &OpenapiService{
		UnimplementedOpenapiServer: v1.UnimplementedOpenapiServer{},
	}
}

// AddonsIdentify implements AddonsIdentify.OpenapiServer.
func (s *OpenapiService) AddonsIdentify(ctx context.Context, in *openapi_v1.AddonsIdentifyRequest) (*openapi_v1.AddonsIdentifyResponse, error) {
	return &openapi_v1.AddonsIdentifyResponse{
		Res: util.GetV1ResultBadRequest("not declare addons"),
	}, nil
}

// Identify implements Identify.OpenapiServer.
func (s *OpenapiService) Identify(ctx context.Context, in *emptypb.Empty) (*openapi_v1.IdentifyResponse, error) {
	return &openapi_v1.IdentifyResponse{
		Res:          util.GetV1ResultOK(),
		PluginId:     "tkeel-hello",
		Version:      "v0.2.0",
		TkeelVersion: "v0.2.0",
	}, nil
}

// Status implements Status.OpenapiServer.
func (s *OpenapiService) Status(ctx context.Context, in *emptypb.Empty) (*openapi_v1.StatusResponse, error) {
	return &openapi_v1.StatusResponse{
		Res:    util.GetV1ResultOK(),
		Status: openapi_v1.PluginStatus_RUNNING,
	}, nil
}

// TenantBind implements TenantBind.OpenapiServer.
func (s *OpenapiService) TenantBind(ctx context.Context, in *openapi_v1.TenantBindRequst) (*openapi_v1.TenantBindResponse, error) {
	return &openapi_v1.TenantBindResponse{
		Res: util.GetV1ResultOK(),
	}, nil
}

// TenantUnbind implements TenantUnbind.OpenapiServer.
func (s *OpenapiService) TenantUnbind(ctx context.Context, in *openapi_v1.TenantUnbindRequst) (*openapi_v1.TenantUnbindResponse, error) {
	return &openapi_v1.TenantUnbindResponse{
		Res: util.GetV1ResultOK(),
	}, nil
}
