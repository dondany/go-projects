// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/todo.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TodoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ListId        int32                  `protobuf:"varint,2,opt,name=listId,proto3" json:"listId,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Completed     bool                   `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	mi := &file_proto_todo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoResponse) GetListId() int32 {
	if x != nil {
		return x.ListId
	}
	return 0
}

func (x *TodoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TodoResponse) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *TodoResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type TodoListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Todos         []*TodoResponse        `protobuf:"bytes,3,rep,name=todos,proto3" json:"todos,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodoListResponse) Reset() {
	*x = TodoListResponse{}
	mi := &file_proto_todo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListResponse) ProtoMessage() {}

func (x *TodoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListResponse.ProtoReflect.Descriptor instead.
func (*TodoListResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoListResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoListResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TodoListResponse) GetTodos() []*TodoResponse {
	if x != nil {
		return x.Todos
	}
	return nil
}

func (x *TodoListResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type TodoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ListId        int32                  `protobuf:"varint,2,opt,name=listId,proto3" json:"listId,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodoRequest) Reset() {
	*x = TodoRequest{}
	mi := &file_proto_todo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoRequest) ProtoMessage() {}

func (x *TodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoRequest.ProtoReflect.Descriptor instead.
func (*TodoRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{2}
}

func (x *TodoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoRequest) GetListId() int32 {
	if x != nil {
		return x.ListId
	}
	return 0
}

func (x *TodoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TodoUpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Completed     bool                   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodoUpdateRequest) Reset() {
	*x = TodoUpdateRequest{}
	mi := &file_proto_todo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodoUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoUpdateRequest) ProtoMessage() {}

func (x *TodoUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoUpdateRequest.ProtoReflect.Descriptor instead.
func (*TodoUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{3}
}

func (x *TodoUpdateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TodoUpdateRequest) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type TodoListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodoListRequest) Reset() {
	*x = TodoListRequest{}
	mi := &file_proto_todo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListRequest) ProtoMessage() {}

func (x *TodoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListRequest.ProtoReflect.Descriptor instead.
func (*TodoListRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{4}
}

func (x *TodoListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateTodoListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTodoListRequest) Reset() {
	*x = UpdateTodoListRequest{}
	mi := &file_proto_todo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTodoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoListRequest) ProtoMessage() {}

func (x *UpdateTodoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoListRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoListRequest) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTodoListRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTodoListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TodoListFilter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodoListFilter) Reset() {
	*x = TodoListFilter{}
	mi := &file_proto_todo_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodoListFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListFilter) ProtoMessage() {}

func (x *TodoListFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListFilter.ProtoReflect.Descriptor instead.
func (*TodoListFilter) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{6}
}

type TodoListsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Lists         []*TodoListResponse    `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodoListsResponse) Reset() {
	*x = TodoListsResponse{}
	mi := &file_proto_todo_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodoListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListsResponse) ProtoMessage() {}

func (x *TodoListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListsResponse.ProtoReflect.Descriptor instead.
func (*TodoListsResponse) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{7}
}

func (x *TodoListsResponse) GetLists() []*TodoListResponse {
	if x != nil {
		return x.Lists
	}
	return nil
}

type ID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ID) Reset() {
	*x = ID{}
	mi := &file_proto_todo_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todo_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_proto_todo_proto_rawDescGZIP(), []int{8}
}

func (x *ID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_todo_proto protoreflect.FileDescriptor

var file_proto_todo_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x10,
	0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x49, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x54, 0x6f, 0x64, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x54, 0x6f,
	0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3b, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x10,
	0x0a, 0x0e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x41, 0x0a, 0x11, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x6c, 0x69,
	0x73, 0x74, 0x73, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd9, 0x03, 0x0a, 0x0b, 0x54, 0x6f,
	0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a,
	0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x08, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x49, 0x44, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x08, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x49, 0x44, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x64, 0x6f, 0x12, 0x11, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x64, 0x6f, 0x12, 0x08, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x49, 0x44, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_todo_proto_rawDescOnce sync.Once
	file_proto_todo_proto_rawDescData []byte
)

func file_proto_todo_proto_rawDescGZIP() []byte {
	file_proto_todo_proto_rawDescOnce.Do(func() {
		file_proto_todo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_todo_proto_rawDesc), len(file_proto_todo_proto_rawDesc)))
	})
	return file_proto_todo_proto_rawDescData
}

var file_proto_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_todo_proto_goTypes = []any{
	(*TodoResponse)(nil),          // 0: todo.TodoResponse
	(*TodoListResponse)(nil),      // 1: todo.TodoListResponse
	(*TodoRequest)(nil),           // 2: todo.TodoRequest
	(*TodoUpdateRequest)(nil),     // 3: todo.TodoUpdateRequest
	(*TodoListRequest)(nil),       // 4: todo.TodoListRequest
	(*UpdateTodoListRequest)(nil), // 5: todo.UpdateTodoListRequest
	(*TodoListFilter)(nil),        // 6: todo.TodoListFilter
	(*TodoListsResponse)(nil),     // 7: todo.TodoListsResponse
	(*ID)(nil),                    // 8: todo.ID
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_proto_todo_proto_depIdxs = []int32{
	9,  // 0: todo.TodoResponse.createdAt:type_name -> google.protobuf.Timestamp
	0,  // 1: todo.TodoListResponse.todos:type_name -> todo.TodoResponse
	9,  // 2: todo.TodoListResponse.createdAt:type_name -> google.protobuf.Timestamp
	1,  // 3: todo.TodoListsResponse.lists:type_name -> todo.TodoListResponse
	6,  // 4: todo.TodoService.GetTodoLists:input_type -> todo.TodoListFilter
	4,  // 5: todo.TodoService.CreateTodoList:input_type -> todo.TodoListRequest
	8,  // 6: todo.TodoService.GetTodoList:input_type -> todo.ID
	5,  // 7: todo.TodoService.UpdateTodoList:input_type -> todo.UpdateTodoListRequest
	8,  // 8: todo.TodoService.DeleteTodoList:input_type -> todo.ID
	2,  // 9: todo.TodoService.CreateTodo:input_type -> todo.TodoRequest
	3,  // 10: todo.TodoService.UpdateTodo:input_type -> todo.TodoUpdateRequest
	8,  // 11: todo.TodoService.DeleteTodo:input_type -> todo.ID
	7,  // 12: todo.TodoService.GetTodoLists:output_type -> todo.TodoListsResponse
	1,  // 13: todo.TodoService.CreateTodoList:output_type -> todo.TodoListResponse
	1,  // 14: todo.TodoService.GetTodoList:output_type -> todo.TodoListResponse
	1,  // 15: todo.TodoService.UpdateTodoList:output_type -> todo.TodoListResponse
	10, // 16: todo.TodoService.DeleteTodoList:output_type -> google.protobuf.Empty
	0,  // 17: todo.TodoService.CreateTodo:output_type -> todo.TodoResponse
	0,  // 18: todo.TodoService.UpdateTodo:output_type -> todo.TodoResponse
	10, // 19: todo.TodoService.DeleteTodo:output_type -> google.protobuf.Empty
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_todo_proto_init() }
func file_proto_todo_proto_init() {
	if File_proto_todo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_todo_proto_rawDesc), len(file_proto_todo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todo_proto_goTypes,
		DependencyIndexes: file_proto_todo_proto_depIdxs,
		MessageInfos:      file_proto_todo_proto_msgTypes,
	}.Build()
	File_proto_todo_proto = out.File
	file_proto_todo_proto_goTypes = nil
	file_proto_todo_proto_depIdxs = nil
}
