// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/proto/rocksdb.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type PutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DatabaseName  string                 `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"` // Name of the database to operate on
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	mi := &file_api_proto_rocksdb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_rocksdb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rocksdb_proto_rawDescGZIP(), []int{0}
}

func (x *PutRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *PutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type PutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutResponse) Reset() {
	*x = PutResponse{}
	mi := &file_api_proto_rocksdb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResponse) ProtoMessage() {}

func (x *PutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_rocksdb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResponse.ProtoReflect.Descriptor instead.
func (*PutResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rocksdb_proto_rawDescGZIP(), []int{1}
}

func (x *PutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PutResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DatabaseName  string                 `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"` // Name of the database to operate on
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_api_proto_rocksdb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_rocksdb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rocksdb_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []byte                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Found         bool                   `protobuf:"varint,2,opt,name=found,proto3" json:"found,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_api_proto_rocksdb_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_rocksdb_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rocksdb_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *GetResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

func (x *GetResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DatabaseName  string                 `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"` // Name of the database to operate on
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_api_proto_rocksdb_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_rocksdb_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rocksdb_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_api_proto_rocksdb_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_rocksdb_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rocksdb_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type StreamGetRequest struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	DatabaseName string                 `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"` // Name of the database to operate on
	// Types that are valid to be assigned to Query:
	//
	//	*StreamGetRequest_Prefix
	//	*StreamGetRequest_Keys
	Query         isStreamGetRequest_Query `protobuf_oneof:"query"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamGetRequest) Reset() {
	*x = StreamGetRequest{}
	mi := &file_api_proto_rocksdb_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamGetRequest) ProtoMessage() {}

func (x *StreamGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_rocksdb_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamGetRequest.ProtoReflect.Descriptor instead.
func (*StreamGetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rocksdb_proto_rawDescGZIP(), []int{6}
}

func (x *StreamGetRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *StreamGetRequest) GetQuery() isStreamGetRequest_Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *StreamGetRequest) GetPrefix() string {
	if x != nil {
		if x, ok := x.Query.(*StreamGetRequest_Prefix); ok {
			return x.Prefix
		}
	}
	return ""
}

func (x *StreamGetRequest) GetKeys() *KeySet {
	if x != nil {
		if x, ok := x.Query.(*StreamGetRequest_Keys); ok {
			return x.Keys
		}
	}
	return nil
}

type isStreamGetRequest_Query interface {
	isStreamGetRequest_Query()
}

type StreamGetRequest_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StreamGetRequest_Keys struct {
	Keys *KeySet `protobuf:"bytes,3,opt,name=keys,proto3,oneof"`
}

func (*StreamGetRequest_Prefix) isStreamGetRequest_Query() {}

func (*StreamGetRequest_Keys) isStreamGetRequest_Query() {}

type KeySet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Keys          []string               `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeySet) Reset() {
	*x = KeySet{}
	mi := &file_api_proto_rocksdb_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeySet) ProtoMessage() {}

func (x *KeySet) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_rocksdb_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeySet.ProtoReflect.Descriptor instead.
func (*KeySet) Descriptor() ([]byte, []int) {
	return file_api_proto_rocksdb_proto_rawDescGZIP(), []int{7}
}

func (x *KeySet) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type StreamGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamGetResponse) Reset() {
	*x = StreamGetResponse{}
	mi := &file_api_proto_rocksdb_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamGetResponse) ProtoMessage() {}

func (x *StreamGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_rocksdb_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamGetResponse.ProtoReflect.Descriptor instead.
func (*StreamGetResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rocksdb_proto_rawDescGZIP(), []int{8}
}

func (x *StreamGetResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StreamGetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *StreamGetResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_api_proto_rocksdb_proto protoreflect.FileDescriptor

var file_api_proto_rocksdb_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x63, 0x6b,
	0x73, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x6f, 0x63, 0x6b, 0x73,
	0x64, 0x62, 0x22, 0x59, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3d, 0x0a,
	0x0b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x43, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x46, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x40, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x81, 0x01, 0x0a,
	0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x25, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x48,
	0x00, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x22, 0x1c, 0x0a, 0x06, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x51,
	0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0xfd, 0x01, 0x0a, 0x0e, 0x52, 0x6f, 0x63, 0x6b, 0x73, 0x44, 0x42, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x13, 0x2e, 0x72, 0x6f,
	0x63, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x13, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x1b, 0x5a, 0x19, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_proto_rocksdb_proto_rawDescOnce sync.Once
	file_api_proto_rocksdb_proto_rawDescData []byte
)

func file_api_proto_rocksdb_proto_rawDescGZIP() []byte {
	file_api_proto_rocksdb_proto_rawDescOnce.Do(func() {
		file_api_proto_rocksdb_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_proto_rocksdb_proto_rawDesc), len(file_api_proto_rocksdb_proto_rawDesc)))
	})
	return file_api_proto_rocksdb_proto_rawDescData
}

var file_api_proto_rocksdb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_rocksdb_proto_goTypes = []any{
	(*PutRequest)(nil),        // 0: rocksdb.PutRequest
	(*PutResponse)(nil),       // 1: rocksdb.PutResponse
	(*GetRequest)(nil),        // 2: rocksdb.GetRequest
	(*GetResponse)(nil),       // 3: rocksdb.GetResponse
	(*DeleteRequest)(nil),     // 4: rocksdb.DeleteRequest
	(*DeleteResponse)(nil),    // 5: rocksdb.DeleteResponse
	(*StreamGetRequest)(nil),  // 6: rocksdb.StreamGetRequest
	(*KeySet)(nil),            // 7: rocksdb.KeySet
	(*StreamGetResponse)(nil), // 8: rocksdb.StreamGetResponse
}
var file_api_proto_rocksdb_proto_depIdxs = []int32{
	7, // 0: rocksdb.StreamGetRequest.keys:type_name -> rocksdb.KeySet
	0, // 1: rocksdb.RocksDBService.Put:input_type -> rocksdb.PutRequest
	2, // 2: rocksdb.RocksDBService.Get:input_type -> rocksdb.GetRequest
	4, // 3: rocksdb.RocksDBService.Delete:input_type -> rocksdb.DeleteRequest
	6, // 4: rocksdb.RocksDBService.StreamGet:input_type -> rocksdb.StreamGetRequest
	1, // 5: rocksdb.RocksDBService.Put:output_type -> rocksdb.PutResponse
	3, // 6: rocksdb.RocksDBService.Get:output_type -> rocksdb.GetResponse
	5, // 7: rocksdb.RocksDBService.Delete:output_type -> rocksdb.DeleteResponse
	8, // 8: rocksdb.RocksDBService.StreamGet:output_type -> rocksdb.StreamGetResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_rocksdb_proto_init() }
func file_api_proto_rocksdb_proto_init() {
	if File_api_proto_rocksdb_proto != nil {
		return
	}
	file_api_proto_rocksdb_proto_msgTypes[6].OneofWrappers = []any{
		(*StreamGetRequest_Prefix)(nil),
		(*StreamGetRequest_Keys)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_rocksdb_proto_rawDesc), len(file_api_proto_rocksdb_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_rocksdb_proto_goTypes,
		DependencyIndexes: file_api_proto_rocksdb_proto_depIdxs,
		MessageInfos:      file_api_proto_rocksdb_proto_msgTypes,
	}.Build()
	File_api_proto_rocksdb_proto = out.File
	file_api_proto_rocksdb_proto_goTypes = nil
	file_api_proto_rocksdb_proto_depIdxs = nil
}
