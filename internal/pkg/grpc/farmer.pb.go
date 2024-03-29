// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: chicken_farmer/v1/farmer.proto

package grpc

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FarmerName string `protobuf:"bytes,1,opt,name=farmer_name,json=farmerName,proto3" json:"farmer_name,omitempty"`
	FarmName   string `protobuf:"bytes,2,opt,name=farm_name,json=farmName,proto3" json:"farm_name,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetFarmerName() string {
	if x != nil {
		return x.FarmerName
	}
	return ""
}

func (x *RegisterRequest) GetFarmName() string {
	if x != nil {
		return x.FarmName
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FarmerId string `protobuf:"bytes,1,opt,name=farmer_id,json=farmerId,proto3" json:"farmer_id,omitempty"`
	FarmId   string `protobuf:"bytes,2,opt,name=farm_id,json=farmId,proto3" json:"farm_id,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetFarmerId() string {
	if x != nil {
		return x.FarmerId
	}
	return ""
}

func (x *RegisterResponse) GetFarmId() string {
	if x != nil {
		return x.FarmId
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FarmerName string `protobuf:"bytes,1,opt,name=farmer_name,json=farmerName,proto3" json:"farmer_name,omitempty"`
	Password   string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetFarmerName() string {
	if x != nil {
		return x.FarmerName
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

type GrantGoldEggsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   uint32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	FarmerId string `protobuf:"bytes,2,opt,name=farmer_id,json=farmerId,proto3" json:"farmer_id,omitempty"`
}

func (x *GrantGoldEggsRequest) Reset() {
	*x = GrantGoldEggsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantGoldEggsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantGoldEggsRequest) ProtoMessage() {}

func (x *GrantGoldEggsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantGoldEggsRequest.ProtoReflect.Descriptor instead.
func (*GrantGoldEggsRequest) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{4}
}

func (x *GrantGoldEggsRequest) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *GrantGoldEggsRequest) GetFarmerId() string {
	if x != nil {
		return x.FarmerId
	}
	return ""
}

type GrantGoldEggsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GrantGoldEggsResponse) Reset() {
	*x = GrantGoldEggsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantGoldEggsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantGoldEggsResponse) ProtoMessage() {}

func (x *GrantGoldEggsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantGoldEggsResponse.ProtoReflect.Descriptor instead.
func (*GrantGoldEggsResponse) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{5}
}

type SpendGoldEggsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   uint32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	FarmerId string `protobuf:"bytes,2,opt,name=farmer_id,json=farmerId,proto3" json:"farmer_id,omitempty"`
}

func (x *SpendGoldEggsRequest) Reset() {
	*x = SpendGoldEggsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpendGoldEggsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpendGoldEggsRequest) ProtoMessage() {}

func (x *SpendGoldEggsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpendGoldEggsRequest.ProtoReflect.Descriptor instead.
func (*SpendGoldEggsRequest) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{6}
}

func (x *SpendGoldEggsRequest) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SpendGoldEggsRequest) GetFarmerId() string {
	if x != nil {
		return x.FarmerId
	}
	return ""
}

type SpendGoldEggsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SpendGoldEggsResponse) Reset() {
	*x = SpendGoldEggsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpendGoldEggsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpendGoldEggsResponse) ProtoMessage() {}

func (x *SpendGoldEggsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpendGoldEggsResponse.ProtoReflect.Descriptor instead.
func (*SpendGoldEggsResponse) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{7}
}

type GetGoldEggsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FarmerId string `protobuf:"bytes,1,opt,name=farmer_id,json=farmerId,proto3" json:"farmer_id,omitempty"`
}

func (x *GetGoldEggsRequest) Reset() {
	*x = GetGoldEggsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoldEggsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoldEggsRequest) ProtoMessage() {}

func (x *GetGoldEggsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoldEggsRequest.ProtoReflect.Descriptor instead.
func (*GetGoldEggsRequest) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{8}
}

func (x *GetGoldEggsRequest) GetFarmerId() string {
	if x != nil {
		return x.FarmerId
	}
	return ""
}

type GetGoldEggsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount uint32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GetGoldEggsResponse) Reset() {
	*x = GetGoldEggsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoldEggsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoldEggsResponse) ProtoMessage() {}

func (x *GetGoldEggsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chicken_farmer_v1_farmer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoldEggsResponse.ProtoReflect.Descriptor instead.
func (*GetGoldEggsResponse) Descriptor() ([]byte, []int) {
	return file_chicken_farmer_v1_farmer_proto_rawDescGZIP(), []int{9}
}

func (x *GetGoldEggsResponse) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_chicken_farmer_v1_farmer_proto protoreflect.FileDescriptor

var file_chicken_farmer_v1_farmer_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x63, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x72, 0x6d,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a,
	0x2a, 0x92, 0x41, 0x27, 0x0a, 0x25, 0xd2, 0x01, 0x0b, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x09, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0xd2, 0x01, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x48, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x66, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x61, 0x72, 0x6d, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x72, 0x6d,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x3a, 0x1e, 0x92, 0x41, 0x1b, 0x0a, 0x19, 0xd2, 0x01, 0x0b, 0x66, 0x61, 0x72, 0x6d,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x2e, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x4b, 0x0a, 0x14, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x47, 0x6f, 0x6c, 0x64, 0x45,
	0x67, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x17, 0x0a, 0x15, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x14, 0x53, 0x70, 0x65, 0x6e,
	0x64, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x72, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x72,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x47, 0x6f,
	0x6c, 0x64, 0x45, 0x67, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0xfd, 0x01, 0x0a, 0x13, 0x46, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x63, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66,
	0x61, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x68, 0x69, 0x63, 0x6b,
	0x65, 0x6e, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x92,
	0x41, 0x02, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x1a, 0x13, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x6c, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x63, 0x68,
	0x69, 0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63,
	0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0x92, 0x41, 0x02, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x32, 0xc2, 0x02, 0x0a, 0x14, 0x46, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x69,
	0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x61,
	0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x47, 0x6f, 0x6c,
	0x64, 0x45, 0x67, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x64, 0x0a, 0x0d, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73,
	0x12, 0x27, 0x2e, 0x63, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x69, 0x63,
	0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70,
	0x65, 0x6e, 0x64, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6c, 0x64,
	0x45, 0x67, 0x67, 0x73, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66,
	0x61, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6c, 0x64,
	0x45, 0x67, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x68,
	0x69, 0x63, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x87, 0x01, 0x5a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x92, 0x41, 0x71, 0x5a, 0x5d,
	0x0a, 0x5b, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x4d,
	0x08, 0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x10, 0x0a,
	0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chicken_farmer_v1_farmer_proto_rawDescOnce sync.Once
	file_chicken_farmer_v1_farmer_proto_rawDescData = file_chicken_farmer_v1_farmer_proto_rawDesc
)

func file_chicken_farmer_v1_farmer_proto_rawDescGZIP() []byte {
	file_chicken_farmer_v1_farmer_proto_rawDescOnce.Do(func() {
		file_chicken_farmer_v1_farmer_proto_rawDescData = protoimpl.X.CompressGZIP(file_chicken_farmer_v1_farmer_proto_rawDescData)
	})
	return file_chicken_farmer_v1_farmer_proto_rawDescData
}

var file_chicken_farmer_v1_farmer_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_chicken_farmer_v1_farmer_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),       // 0: chicken_farmer.v1.RegisterRequest
	(*RegisterResponse)(nil),      // 1: chicken_farmer.v1.RegisterResponse
	(*LoginRequest)(nil),          // 2: chicken_farmer.v1.LoginRequest
	(*LoginResponse)(nil),         // 3: chicken_farmer.v1.LoginResponse
	(*GrantGoldEggsRequest)(nil),  // 4: chicken_farmer.v1.GrantGoldEggsRequest
	(*GrantGoldEggsResponse)(nil), // 5: chicken_farmer.v1.GrantGoldEggsResponse
	(*SpendGoldEggsRequest)(nil),  // 6: chicken_farmer.v1.SpendGoldEggsRequest
	(*SpendGoldEggsResponse)(nil), // 7: chicken_farmer.v1.SpendGoldEggsResponse
	(*GetGoldEggsRequest)(nil),    // 8: chicken_farmer.v1.GetGoldEggsRequest
	(*GetGoldEggsResponse)(nil),   // 9: chicken_farmer.v1.GetGoldEggsResponse
}
var file_chicken_farmer_v1_farmer_proto_depIdxs = []int32{
	0, // 0: chicken_farmer.v1.FarmerPublicService.Register:input_type -> chicken_farmer.v1.RegisterRequest
	2, // 1: chicken_farmer.v1.FarmerPublicService.Login:input_type -> chicken_farmer.v1.LoginRequest
	4, // 2: chicken_farmer.v1.FarmerPrivateService.GrantGoldEggs:input_type -> chicken_farmer.v1.GrantGoldEggsRequest
	6, // 3: chicken_farmer.v1.FarmerPrivateService.SpendGoldEggs:input_type -> chicken_farmer.v1.SpendGoldEggsRequest
	8, // 4: chicken_farmer.v1.FarmerPrivateService.GetGoldEggs:input_type -> chicken_farmer.v1.GetGoldEggsRequest
	1, // 5: chicken_farmer.v1.FarmerPublicService.Register:output_type -> chicken_farmer.v1.RegisterResponse
	3, // 6: chicken_farmer.v1.FarmerPublicService.Login:output_type -> chicken_farmer.v1.LoginResponse
	5, // 7: chicken_farmer.v1.FarmerPrivateService.GrantGoldEggs:output_type -> chicken_farmer.v1.GrantGoldEggsResponse
	7, // 8: chicken_farmer.v1.FarmerPrivateService.SpendGoldEggs:output_type -> chicken_farmer.v1.SpendGoldEggsResponse
	9, // 9: chicken_farmer.v1.FarmerPrivateService.GetGoldEggs:output_type -> chicken_farmer.v1.GetGoldEggsResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chicken_farmer_v1_farmer_proto_init() }
func file_chicken_farmer_v1_farmer_proto_init() {
	if File_chicken_farmer_v1_farmer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chicken_farmer_v1_farmer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_chicken_farmer_v1_farmer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_chicken_farmer_v1_farmer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_chicken_farmer_v1_farmer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_chicken_farmer_v1_farmer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantGoldEggsRequest); i {
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
		file_chicken_farmer_v1_farmer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantGoldEggsResponse); i {
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
		file_chicken_farmer_v1_farmer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpendGoldEggsRequest); i {
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
		file_chicken_farmer_v1_farmer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpendGoldEggsResponse); i {
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
		file_chicken_farmer_v1_farmer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoldEggsRequest); i {
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
		file_chicken_farmer_v1_farmer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoldEggsResponse); i {
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
			RawDescriptor: file_chicken_farmer_v1_farmer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_chicken_farmer_v1_farmer_proto_goTypes,
		DependencyIndexes: file_chicken_farmer_v1_farmer_proto_depIdxs,
		MessageInfos:      file_chicken_farmer_v1_farmer_proto_msgTypes,
	}.Build()
	File_chicken_farmer_v1_farmer_proto = out.File
	file_chicken_farmer_v1_farmer_proto_rawDesc = nil
	file_chicken_farmer_v1_farmer_proto_goTypes = nil
	file_chicken_farmer_v1_farmer_proto_depIdxs = nil
}
