// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: base/pagination.proto

package pagination

import (
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

// Enum to specify the order of results.
type Order int32

const (
	// Unspecified order (default value). The server will determine the order.
	Order_ORDER_UNSPECIFIED Order = 0
	// Ascending order. Results will be arranged in ascending order.
	Order_ASCENDING Order = 1
	// Descending order. Results will be arranged in descending order.
	Order_DESCENDING Order = 2
)

// Enum value maps for Order.
var (
	Order_name = map[int32]string{
		0: "ORDER_UNSPECIFIED",
		1: "ASCENDING",
		2: "DESCENDING",
	}
	Order_value = map[string]int32{
		"ORDER_UNSPECIFIED": 0,
		"ASCENDING":         1,
		"DESCENDING":        2,
	}
)

func (x Order) Enum() *Order {
	p := new(Order)
	*p = x
	return p
}

func (x Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order) Descriptor() protoreflect.EnumDescriptor {
	return file_base_pagination_proto_enumTypes[0].Descriptor()
}

func (Order) Type() protoreflect.EnumType {
	return &file_base_pagination_proto_enumTypes[0]
}

func (x Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order.Descriptor instead.
func (Order) EnumDescriptor() ([]byte, []int) {
	return file_base_pagination_proto_rawDescGZIP(), []int{0}
}

// Request message with pagination parameters.
type RequestWithPagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional: The maximum number of items to return in the response.
	// If set, the server will limit the number of results to this value.
	// If not provided, the server will determine the appropriate limit.
	Limit *int32 `protobuf:"varint,1,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	// Optional: Cursor to retrieve the next page of results.
	// When paginating through a large set of results, the server may return a
	// cursor in the response. This cursor can be used to retrieve the next page
	// of results. If not provided, the server will start from the beginning.
	Cursor *string `protobuf:"bytes,2,opt,name=cursor,proto3,oneof" json:"cursor,omitempty"`
	// Optional: Order in which the results should be sorted.
	// You can specify whether the results should be in ascending or descending order.
	// If not provided, the server may use a default order.
	OrderBy *Order `protobuf:"varint,3,opt,name=order_by,json=orderBy,proto3,enum=base.v1.pagination.Order,oneof" json:"order_by,omitempty"`
}

func (x *RequestWithPagination) Reset() {
	*x = RequestWithPagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestWithPagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestWithPagination) ProtoMessage() {}

func (x *RequestWithPagination) ProtoReflect() protoreflect.Message {
	mi := &file_base_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestWithPagination.ProtoReflect.Descriptor instead.
func (*RequestWithPagination) Descriptor() ([]byte, []int) {
	return file_base_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *RequestWithPagination) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *RequestWithPagination) GetCursor() string {
	if x != nil && x.Cursor != nil {
		return *x.Cursor
	}
	return ""
}

func (x *RequestWithPagination) GetOrderBy() Order {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return Order_ORDER_UNSPECIFIED
}

// Response message with pagination data.
type ResponseWithPagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional: Cursor to retrieve the next page of results.
	// If the server has more results to display, it will include a cursor in the
	// response. The client can use this cursor to fetch the next page of data.
	// The cursor should be a 32-character UUID (e.g., "4dd5f17d-07a9-4c15-aa34-19d98d4c8d6a").
	NextCursor *string `protobuf:"bytes,1,opt,name=next_cursor,json=nextCursor,proto3,oneof" json:"next_cursor,omitempty"`
	// Optional: Cursor to retrieve the previous page of results.
	// If the client wants to navigate back to the previous page of results,
	// it can use this cursor to fetch the previous set of data.
	// The cursor should be a 32-character UUID (e.g., "9a380e19-63b4-4e47-b5f7-8f8994e5e65d").
	PreviousCursor *string `protobuf:"bytes,2,opt,name=previous_cursor,json=previousCursor,proto3,oneof" json:"previous_cursor,omitempty"`
}

func (x *ResponseWithPagination) Reset() {
	*x = ResponseWithPagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseWithPagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseWithPagination) ProtoMessage() {}

func (x *ResponseWithPagination) ProtoReflect() protoreflect.Message {
	mi := &file_base_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseWithPagination.ProtoReflect.Descriptor instead.
func (*ResponseWithPagination) Descriptor() ([]byte, []int) {
	return file_base_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseWithPagination) GetNextCursor() string {
	if x != nil && x.NextCursor != nil {
		return *x.NextCursor
	}
	return ""
}

func (x *ResponseWithPagination) GetPreviousCursor() string {
	if x != nil && x.PreviousCursor != nil {
		return *x.PreviousCursor
	}
	return ""
}

var File_base_pagination_proto protoreflect.FileDescriptor

var file_base_pagination_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xac, 0x01, 0x0a, 0x15,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x02, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x16, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x69, 0x74, 0x68, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6e, 0x65,
	0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2a, 0x3d, 0x0a,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x0d, 0x5a, 0x0b,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_base_pagination_proto_rawDescOnce sync.Once
	file_base_pagination_proto_rawDescData = file_base_pagination_proto_rawDesc
)

func file_base_pagination_proto_rawDescGZIP() []byte {
	file_base_pagination_proto_rawDescOnce.Do(func() {
		file_base_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_pagination_proto_rawDescData)
	})
	return file_base_pagination_proto_rawDescData
}

var file_base_pagination_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_base_pagination_proto_goTypes = []interface{}{
	(Order)(0),                     // 0: base.v1.pagination.Order
	(*RequestWithPagination)(nil),  // 1: base.v1.pagination.RequestWithPagination
	(*ResponseWithPagination)(nil), // 2: base.v1.pagination.ResponseWithPagination
}
var file_base_pagination_proto_depIdxs = []int32{
	0, // 0: base.v1.pagination.RequestWithPagination.order_by:type_name -> base.v1.pagination.Order
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_base_pagination_proto_init() }
func file_base_pagination_proto_init() {
	if File_base_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestWithPagination); i {
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
		file_base_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseWithPagination); i {
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
	file_base_pagination_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_base_pagination_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_pagination_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_pagination_proto_goTypes,
		DependencyIndexes: file_base_pagination_proto_depIdxs,
		EnumInfos:         file_base_pagination_proto_enumTypes,
		MessageInfos:      file_base_pagination_proto_msgTypes,
	}.Build()
	File_base_pagination_proto = out.File
	file_base_pagination_proto_rawDesc = nil
	file_base_pagination_proto_goTypes = nil
	file_base_pagination_proto_depIdxs = nil
}
