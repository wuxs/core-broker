// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 0.1.0

package v1

import (
	context "context"
	go_restful "github.com/emicklei/go-restful"
	errors "github.com/tkeel-io/kit/errors"
	result "github.com/tkeel-io/kit/result"
	protojson "google.golang.org/protobuf/encoding/protojson"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.anypb.result.protojson.go_restful.errors.emptypb.

var (
	_ = protojson.MarshalOptions{}
	_ = anypb.Any{}
	_ = emptypb.Empty{}
)

type SubscribeHTTPServer interface {
	CreateSubscribe(context.Context, *CreateSubscribeRequest) (*CreateSubscribeResponse, error)
	DeleteSubscribe(context.Context, *DeleteSubscribeRequest) (*DeleteSubscribeResponse, error)
	GetSubscribe(context.Context, *GetSubscribeRequest) (*GetSubscribeResponse, error)
	ListSubscribe(context.Context, *ListSubscribeRequest) (*ListSubscribeResponse, error)
	ListSubscribeEntities(context.Context, *ListSubscribeEntitiesRequest) (*ListSubscribeEntitiesResponse, error)
	SubscribeEntitiesByGroups(context.Context, *SubscribeEntitiesByGroupsRequest) (*SubscribeEntitiesByGroupsResponse, error)
	SubscribeEntitiesByIDs(context.Context, *SubscribeEntitiesByIDsRequest) (*SubscribeEntitiesByIDsResponse, error)
	SubscribeEntitiesByModels(context.Context, *SubscribeEntitiesByModelsRequest) (*SubscribeEntitiesByModelsResponse, error)
	UnsubscribeEntitiesByIDs(context.Context, *UnsubscribeEntitiesByIDsRequest) (*UnsubscribeEntitiesByIDsResponse, error)
	UpdateSubscribe(context.Context, *UpdateSubscribeRequest) (*UpdateSubscribeResponse, error)
}

type SubscribeHTTPHandler struct {
	srv SubscribeHTTPServer
}

func newSubscribeHTTPHandler(s SubscribeHTTPServer) *SubscribeHTTPHandler {
	return &SubscribeHTTPHandler{srv: s}
}

func (h *SubscribeHTTPHandler) CreateSubscribe(req *go_restful.Request, resp *go_restful.Response) {
	in := CreateSubscribeRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.CreateSubscribe(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *SubscribeHTTPHandler) DeleteSubscribe(req *go_restful.Request, resp *go_restful.Response) {
	in := DeleteSubscribeRequest{}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.DeleteSubscribe(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *SubscribeHTTPHandler) GetSubscribe(req *go_restful.Request, resp *go_restful.Response) {
	in := GetSubscribeRequest{}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.GetSubscribe(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *SubscribeHTTPHandler) ListSubscribe(req *go_restful.Request, resp *go_restful.Response) {
	in := ListSubscribeRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.ListSubscribe(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *SubscribeHTTPHandler) ListSubscribeEntities(req *go_restful.Request, resp *go_restful.Response) {
	in := ListSubscribeEntitiesRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.ListSubscribeEntities(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *SubscribeHTTPHandler) SubscribeEntitiesByGroups(req *go_restful.Request, resp *go_restful.Response) {
	in := SubscribeEntitiesByGroupsRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.SubscribeEntitiesByGroups(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *SubscribeHTTPHandler) SubscribeEntitiesByIDs(req *go_restful.Request, resp *go_restful.Response) {
	in := SubscribeEntitiesByIDsRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.SubscribeEntitiesByIDs(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *SubscribeHTTPHandler) SubscribeEntitiesByModels(req *go_restful.Request, resp *go_restful.Response) {
	in := SubscribeEntitiesByModelsRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.SubscribeEntitiesByModels(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *SubscribeHTTPHandler) UnsubscribeEntitiesByIDs(req *go_restful.Request, resp *go_restful.Response) {
	in := UnsubscribeEntitiesByIDsRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.UnsubscribeEntitiesByIDs(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func (h *SubscribeHTTPHandler) UpdateSubscribe(req *go_restful.Request, resp *go_restful.Response) {
	in := UpdateSubscribeRequest{}
	if err := transportHTTP.GetBody(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteHeaderAndJson(http.StatusBadRequest,
			result.Set(http.StatusBadRequest, err.Error(), nil), "application/json")
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.UpdateSubscribe(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteHeaderAndJson(httpCode,
			result.Set(httpCode, tErr.Message, out), "application/json")
		return
	}
	anyOut, err := anypb.New(out)
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}

	outB, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}.Marshal(&result.Http{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: anyOut,
	})
	if err != nil {
		resp.WriteHeaderAndJson(http.StatusInternalServerError,
			result.Set(http.StatusInternalServerError, err.Error(), nil), "application/json")
		return
	}
	resp.WriteHeader(http.StatusOK)

	var remain int
	for {
		outB = outB[remain:]
		remain, err = resp.Write(outB)
		if err != nil {
			return
		}
		if remain == 0 {
			break
		}
	}
}

func RegisterSubscribeHTTPServer(container *go_restful.Container, srv SubscribeHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newSubscribeHTTPHandler(srv)
	ws.Route(ws.POST("/subscribe/{id}/entities").
		To(handler.SubscribeEntitiesByIDs))
	ws.Route(ws.POST("/subscribe/{id}/groups").
		To(handler.SubscribeEntitiesByGroups))
	ws.Route(ws.POST("/subscribe/{id}/models").
		To(handler.SubscribeEntitiesByModels))
	ws.Route(ws.POST("/subscribe/{id}/entities/delete").
		To(handler.UnsubscribeEntitiesByIDs))
	ws.Route(ws.POST("/subscribe/{id}/entities/list").
		To(handler.ListSubscribeEntities))
	ws.Route(ws.POST("/subscribe").
		To(handler.CreateSubscribe))
	ws.Route(ws.PUT("/subscribe/{id}").
		To(handler.UpdateSubscribe))
	ws.Route(ws.DELETE("/subscribe/{id}").
		To(handler.DeleteSubscribe))
	ws.Route(ws.GET("/subscribe/{id}").
		To(handler.GetSubscribe))
	ws.Route(ws.POST("/subscribe/list").
		To(handler.ListSubscribe))
}