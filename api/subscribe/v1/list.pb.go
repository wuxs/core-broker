// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/subscribe/v1/list.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum      uint64 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize     uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	OrderBy      string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	IsDescending bool   `protobuf:"varint,4,opt,name=is_descending,json=isDescending,proto3" json:"is_descending,omitempty"`
	KeyWords     string `protobuf:"bytes,5,opt,name=key_words,json=keyWords,proto3" json:"key_words,omitempty"`
	SearchKey    string `protobuf:"bytes,6,opt,name=search_key,json=searchKey,proto3" json:"search_key,omitempty"`
	//user define
	PacketId int64 `protobuf:"varint,10,opt,name=packet_id,json=packetId,proto3" json:"packet_id,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subscribe_v1_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscribe_v1_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_api_subscribe_v1_list_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetPageNum() uint64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListRequest) GetIsDescending() bool {
	if x != nil {
		return x.IsDescending
	}
	return false
}

func (x *ListRequest) GetKeyWords() string {
	if x != nil {
		return x.KeyWords
	}
	return ""
}

func (x *ListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

func (x *ListRequest) GetPacketId() int64 {
	if x != nil {
		return x.PacketId
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageNum  uint64 `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	LastPage uint64 `protobuf:"varint,3,opt,name=last_page,json=lastPage,proto3" json:"last_page,omitempty"`
	PageSize uint64 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	//user define
	PacketId uint64 `protobuf:"varint,10,opt,name=packet_id,json=packetId,proto3" json:"packet_id,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subscribe_v1_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_subscribe_v1_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_api_subscribe_v1_list_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListResponse) GetPageNum() uint64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListResponse) GetLastPage() uint64 {
	if x != nil {
		return x.LastPage
	}
	return 0
}

func (x *ListResponse) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListResponse) GetPacketId() uint64 {
	if x != nil {
		return x.PacketId
	}
	return 0
}

var File_api_subscribe_v1_list_proto protoreflect.FileDescriptor

var file_api_subscribe_v1_list_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe9, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x13, 0xe0, 0x41, 0x02, 0x92, 0x41, 0x0d, 0x32, 0x0b, 0x50, 0x61, 0x67, 0x65,
	0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x2e, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x11, 0xe0, 0x41, 0x02, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x50, 0x61, 0x67,
	0x65, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x2b, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x10, 0xe0, 0x41, 0x01, 0x92, 0x41, 0x0a, 0x32, 0x08, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x20, 0x62, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x3a, 0x0a,
	0x0d, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x15, 0xe0, 0x41, 0x01, 0x92, 0x41, 0x0f, 0x32, 0x0d, 0x49, 0x73,
	0x20, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x69, 0x73, 0x44,
	0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x09, 0x6b, 0x65, 0x79,
	0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xe0, 0x41,
	0x01, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x4b, 0x65, 0x79, 0x20, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xe0,
	0x41, 0x01, 0x92, 0x41, 0x0c, 0x32, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x4b, 0x65,
	0x79, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x09,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x11, 0xe0, 0x41, 0x01, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x20,
	0x49, 0x44, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0xf3, 0x01, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41,
	0x02, 0x92, 0x41, 0x07, 0x32, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x13, 0xe0, 0x41, 0x02, 0x92, 0x41, 0x0d, 0x32, 0x0b, 0x50, 0x61,
	0x67, 0x65, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x2e, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x11, 0xe0, 0x41, 0x02, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x4c,
	0x61, 0x73, 0x74, 0x20, 0x70, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x11, 0xe0, 0x41, 0x02, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x50,
	0x61, 0x67, 0x65, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x04, 0x42, 0x11, 0xe0, 0x41, 0x01, 0x92, 0x41, 0x0b, 0x32, 0x09, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x49, 0x44, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_subscribe_v1_list_proto_rawDescOnce sync.Once
	file_api_subscribe_v1_list_proto_rawDescData = file_api_subscribe_v1_list_proto_rawDesc
)

func file_api_subscribe_v1_list_proto_rawDescGZIP() []byte {
	file_api_subscribe_v1_list_proto_rawDescOnce.Do(func() {
		file_api_subscribe_v1_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_subscribe_v1_list_proto_rawDescData)
	})
	return file_api_subscribe_v1_list_proto_rawDescData
}

var file_api_subscribe_v1_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_subscribe_v1_list_proto_goTypes = []interface{}{
	(*ListRequest)(nil),  // 0: api.subscribe.v1.ListRequest
	(*ListResponse)(nil), // 1: api.subscribe.v1.ListResponse
}
var file_api_subscribe_v1_list_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_subscribe_v1_list_proto_init() }
func file_api_subscribe_v1_list_proto_init() {
	if File_api_subscribe_v1_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_subscribe_v1_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_subscribe_v1_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_subscribe_v1_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_subscribe_v1_list_proto_goTypes,
		DependencyIndexes: file_api_subscribe_v1_list_proto_depIdxs,
		MessageInfos:      file_api_subscribe_v1_list_proto_msgTypes,
	}.Build()
	File_api_subscribe_v1_list_proto = out.File
	file_api_subscribe_v1_list_proto_rawDesc = nil
	file_api_subscribe_v1_list_proto_goTypes = nil
	file_api_subscribe_v1_list_proto_depIdxs = nil
}
