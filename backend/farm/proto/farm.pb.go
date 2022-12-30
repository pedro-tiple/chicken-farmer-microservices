// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: farm/proto/farm.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Farm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Day        uint32  `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	GoldenEggs uint32  `protobuf:"varint,3,opt,name=golden_eggs,json=goldenEggs,proto3" json:"golden_eggs,omitempty"`
	Barns      []*Barn `protobuf:"bytes,4,rep,name=barns,proto3" json:"barns,omitempty"`
}

func (x *Farm) Reset() {
	*x = Farm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Farm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Farm) ProtoMessage() {}

func (x *Farm) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Farm.ProtoReflect.Descriptor instead.
func (*Farm) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{0}
}

func (x *Farm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Farm) GetDay() uint32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *Farm) GetGoldenEggs() uint32 {
	if x != nil {
		return x.GoldenEggs
	}
	return 0
}

func (x *Farm) GetBarns() []*Barn {
	if x != nil {
		return x.Barns
	}
	return nil
}

type Barn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Feed          uint32     `protobuf:"varint,2,opt,name=feed,proto3" json:"feed,omitempty"`
	HasAutoFeeder bool       `protobuf:"varint,3,opt,name=has_auto_feeder,json=hasAutoFeeder,proto3" json:"has_auto_feeder,omitempty"`
	Chickens      []*Chicken `protobuf:"bytes,4,rep,name=chickens,proto3" json:"chickens,omitempty"`
}

func (x *Barn) Reset() {
	*x = Barn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Barn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Barn) ProtoMessage() {}

func (x *Barn) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Barn.ProtoReflect.Descriptor instead.
func (*Barn) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{1}
}

func (x *Barn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Barn) GetFeed() uint32 {
	if x != nil {
		return x.Feed
	}
	return 0
}

func (x *Barn) GetHasAutoFeeder() bool {
	if x != nil {
		return x.HasAutoFeeder
	}
	return false
}

func (x *Barn) GetChickens() []*Chicken {
	if x != nil {
		return x.Chickens
	}
	return nil
}

type Chicken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DateOfBirth    uint32 `protobuf:"varint,2,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	RestingUntil   uint32 `protobuf:"varint,5,opt,name=resting_until,json=restingUntil,proto3" json:"resting_until,omitempty"`
	NormalEggsLaid uint32 `protobuf:"varint,3,opt,name=normal_eggs_laid,json=normalEggsLaid,proto3" json:"normal_eggs_laid,omitempty"`
	GoldEggsLaid   uint32 `protobuf:"varint,4,opt,name=gold_eggs_laid,json=goldEggsLaid,proto3" json:"gold_eggs_laid,omitempty"`
}

func (x *Chicken) Reset() {
	*x = Chicken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chicken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chicken) ProtoMessage() {}

func (x *Chicken) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chicken.ProtoReflect.Descriptor instead.
func (*Chicken) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{2}
}

func (x *Chicken) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Chicken) GetDateOfBirth() uint32 {
	if x != nil {
		return x.DateOfBirth
	}
	return 0
}

func (x *Chicken) GetRestingUntil() uint32 {
	if x != nil {
		return x.RestingUntil
	}
	return 0
}

func (x *Chicken) GetNormalEggsLaid() uint32 {
	if x != nil {
		return x.NormalEggsLaid
	}
	return 0
}

func (x *Chicken) GetGoldEggsLaid() uint32 {
	if x != nil {
		return x.GoldEggsLaid
	}
	return 0
}

type GetFarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFarmRequest) Reset() {
	*x = GetFarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFarmRequest) ProtoMessage() {}

func (x *GetFarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFarmRequest.ProtoReflect.Descriptor instead.
func (*GetFarmRequest) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{3}
}

type GetFarmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Farm *Farm `protobuf:"bytes,1,opt,name=farm,proto3" json:"farm,omitempty"`
}

func (x *GetFarmResponse) Reset() {
	*x = GetFarmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFarmResponse) ProtoMessage() {}

func (x *GetFarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFarmResponse.ProtoReflect.Descriptor instead.
func (*GetFarmResponse) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{4}
}

func (x *GetFarmResponse) GetFarm() *Farm {
	if x != nil {
		return x.Farm
	}
	return nil
}

type BuyBarnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuyBarnRequest) Reset() {
	*x = BuyBarnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyBarnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyBarnRequest) ProtoMessage() {}

func (x *BuyBarnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyBarnRequest.ProtoReflect.Descriptor instead.
func (*BuyBarnRequest) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{5}
}

type BuyBarnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuyBarnResponse) Reset() {
	*x = BuyBarnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyBarnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyBarnResponse) ProtoMessage() {}

func (x *BuyBarnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyBarnResponse.ProtoReflect.Descriptor instead.
func (*BuyBarnResponse) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{6}
}

type BuyFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarnID string `protobuf:"bytes,1,opt,name=barnID,proto3" json:"barnID,omitempty"`
	Amount uint32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BuyFeedRequest) Reset() {
	*x = BuyFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyFeedRequest) ProtoMessage() {}

func (x *BuyFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyFeedRequest.ProtoReflect.Descriptor instead.
func (*BuyFeedRequest) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{7}
}

func (x *BuyFeedRequest) GetBarnID() string {
	if x != nil {
		return x.BarnID
	}
	return ""
}

func (x *BuyFeedRequest) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BuyFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuyFeedResponse) Reset() {
	*x = BuyFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyFeedResponse) ProtoMessage() {}

func (x *BuyFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyFeedResponse.ProtoReflect.Descriptor instead.
func (*BuyFeedResponse) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{8}
}

type BuyChickenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarnID string `protobuf:"bytes,1,opt,name=barnID,proto3" json:"barnID,omitempty"`
}

func (x *BuyChickenRequest) Reset() {
	*x = BuyChickenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyChickenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyChickenRequest) ProtoMessage() {}

func (x *BuyChickenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyChickenRequest.ProtoReflect.Descriptor instead.
func (*BuyChickenRequest) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{9}
}

func (x *BuyChickenRequest) GetBarnID() string {
	if x != nil {
		return x.BarnID
	}
	return ""
}

type BuyChickenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BuyChickenResponse) Reset() {
	*x = BuyChickenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyChickenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyChickenResponse) ProtoMessage() {}

func (x *BuyChickenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyChickenResponse.ProtoReflect.Descriptor instead.
func (*BuyChickenResponse) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{10}
}

func (x *BuyChickenResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FeedChickenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChickenID string `protobuf:"bytes,1,opt,name=chickenID,proto3" json:"chickenID,omitempty"`
}

func (x *FeedChickenRequest) Reset() {
	*x = FeedChickenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedChickenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedChickenRequest) ProtoMessage() {}

func (x *FeedChickenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedChickenRequest.ProtoReflect.Descriptor instead.
func (*FeedChickenRequest) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{11}
}

func (x *FeedChickenRequest) GetChickenID() string {
	if x != nil {
		return x.ChickenID
	}
	return ""
}

type FeedChickenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeedChickenResponse) Reset() {
	*x = FeedChickenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedChickenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedChickenResponse) ProtoMessage() {}

func (x *FeedChickenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedChickenResponse.ProtoReflect.Descriptor instead.
func (*FeedChickenResponse) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{12}
}

type FeedChickensOfBarnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BarnID string `protobuf:"bytes,1,opt,name=barnID,proto3" json:"barnID,omitempty"`
}

func (x *FeedChickensOfBarnRequest) Reset() {
	*x = FeedChickensOfBarnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedChickensOfBarnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedChickensOfBarnRequest) ProtoMessage() {}

func (x *FeedChickensOfBarnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedChickensOfBarnRequest.ProtoReflect.Descriptor instead.
func (*FeedChickensOfBarnRequest) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{13}
}

func (x *FeedChickensOfBarnRequest) GetBarnID() string {
	if x != nil {
		return x.BarnID
	}
	return ""
}

type FeedChickensOfBarnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeedChickensOfBarnResponse) Reset() {
	*x = FeedChickensOfBarnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_farm_proto_farm_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedChickensOfBarnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedChickensOfBarnResponse) ProtoMessage() {}

func (x *FeedChickensOfBarnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_farm_proto_farm_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedChickensOfBarnResponse.ProtoReflect.Descriptor instead.
func (*FeedChickensOfBarnResponse) Descriptor() ([]byte, []int) {
	return file_farm_proto_farm_proto_rawDescGZIP(), []int{14}
}

var File_farm_proto_farm_proto protoreflect.FileDescriptor

var file_farm_proto_farm_proto_rawDesc = []byte{
	0x0a, 0x15, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x72,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x6f, 0x0a,
	0x04, 0x46, 0x61, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x67,
	0x6f, 0x6c, 0x64, 0x65, 0x6e, 0x5f, 0x65, 0x67, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x67, 0x6f, 0x6c, 0x64, 0x65, 0x6e, 0x45, 0x67, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x05,
	0x62, 0x61, 0x72, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x42, 0x61, 0x72, 0x6e, 0x52, 0x05, 0x62, 0x61, 0x72, 0x6e, 0x73, 0x22, 0x7d,
	0x0a, 0x04, 0x42, 0x61, 0x72, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x61,
	0x73, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x46, 0x65, 0x65, 0x64,
	0x65, 0x72, 0x12, 0x29, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x69, 0x63,
	0x6b, 0x65, 0x6e, 0x52, 0x08, 0x63, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0xb2, 0x01,
	0x0a, 0x07, 0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x74,
	0x69, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x65, 0x67, 0x67,
	0x73, 0x5f, 0x6c, 0x61, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x45, 0x67, 0x67, 0x73, 0x4c, 0x61, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x67, 0x6f, 0x6c, 0x64, 0x5f, 0x65, 0x67, 0x67, 0x73, 0x5f, 0x6c, 0x61, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x67, 0x6f, 0x6c, 0x64, 0x45, 0x67, 0x67, 0x73, 0x4c, 0x61,
	0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x66, 0x61, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x61, 0x72,
	0x6d, 0x52, 0x04, 0x66, 0x61, 0x72, 0x6d, 0x22, 0x10, 0x0a, 0x0e, 0x42, 0x75, 0x79, 0x42, 0x61,
	0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x42, 0x75, 0x79,
	0x42, 0x61, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x0a, 0x0e,
	0x42, 0x75, 0x79, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x61, 0x72, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x61, 0x72, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x11,
	0x0a, 0x0f, 0x42, 0x75, 0x79, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2b, 0x0a, 0x11, 0x42, 0x75, 0x79, 0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x72, 0x6e, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x72, 0x6e, 0x49, 0x44, 0x22, 0x24,
	0x0a, 0x12, 0x42, 0x75, 0x79, 0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x12, 0x46, 0x65, 0x65, 0x64, 0x43, 0x68, 0x69, 0x63,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68,
	0x69, 0x63, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x22, 0x15, 0x0a, 0x13, 0x46, 0x65, 0x65, 0x64,
	0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x33, 0x0a, 0x19, 0x46, 0x65, 0x65, 0x64, 0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x73, 0x4f,
	0x66, 0x42, 0x61, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x61, 0x72, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x61,
	0x72, 0x6e, 0x49, 0x44, 0x22, 0x1c, 0x0a, 0x1a, 0x46, 0x65, 0x65, 0x64, 0x43, 0x68, 0x69, 0x63,
	0x6b, 0x65, 0x6e, 0x73, 0x4f, 0x66, 0x42, 0x61, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x9f, 0x03, 0x0a, 0x0b, 0x46, 0x61, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x12, 0x14, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61,
	0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07,
	0x42, 0x75, 0x79, 0x42, 0x61, 0x72, 0x6e, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42,
	0x75, 0x79, 0x42, 0x61, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x75, 0x79, 0x42, 0x61, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x42, 0x75, 0x79, 0x46, 0x65, 0x65,
	0x64, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x75, 0x79, 0x46, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42,
	0x75, 0x79, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x0a, 0x42, 0x75, 0x79, 0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x12, 0x17,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x75, 0x79, 0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42,
	0x75, 0x79, 0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x43, 0x68, 0x69, 0x63, 0x6b,
	0x65, 0x6e, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x43, 0x68,
	0x69, 0x63, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x12, 0x46, 0x65, 0x65,
	0x64, 0x43, 0x68, 0x69, 0x63, 0x6b, 0x65, 0x6e, 0x73, 0x4f, 0x66, 0x42, 0x61, 0x72, 0x6e, 0x12,
	0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x43, 0x68, 0x69, 0x63, 0x6b,
	0x65, 0x6e, 0x73, 0x4f, 0x66, 0x42, 0x61, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x43, 0x68, 0x69, 0x63,
	0x6b, 0x65, 0x6e, 0x73, 0x4f, 0x66, 0x42, 0x61, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x64, 0x72, 0x6f, 0x2d, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2f, 0x66,
	0x61, 0x72, 0x6d, 0x2d, 0x66, 0x61, 0x72, 0x6d, 0x65, 0x72, 0x2d, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x66, 0x61, 0x72, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_farm_proto_farm_proto_rawDescOnce sync.Once
	file_farm_proto_farm_proto_rawDescData = file_farm_proto_farm_proto_rawDesc
)

func file_farm_proto_farm_proto_rawDescGZIP() []byte {
	file_farm_proto_farm_proto_rawDescOnce.Do(func() {
		file_farm_proto_farm_proto_rawDescData = protoimpl.X.CompressGZIP(file_farm_proto_farm_proto_rawDescData)
	})
	return file_farm_proto_farm_proto_rawDescData
}

var file_farm_proto_farm_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_farm_proto_farm_proto_goTypes = []interface{}{
	(*Farm)(nil),                       // 0: grpc.Farm
	(*Barn)(nil),                       // 1: grpc.Barn
	(*Chicken)(nil),                    // 2: grpc.Chicken
	(*GetFarmRequest)(nil),             // 3: grpc.GetFarmRequest
	(*GetFarmResponse)(nil),            // 4: grpc.GetFarmResponse
	(*BuyBarnRequest)(nil),             // 5: grpc.BuyBarnRequest
	(*BuyBarnResponse)(nil),            // 6: grpc.BuyBarnResponse
	(*BuyFeedRequest)(nil),             // 7: grpc.BuyFeedRequest
	(*BuyFeedResponse)(nil),            // 8: grpc.BuyFeedResponse
	(*BuyChickenRequest)(nil),          // 9: grpc.BuyChickenRequest
	(*BuyChickenResponse)(nil),         // 10: grpc.BuyChickenResponse
	(*FeedChickenRequest)(nil),         // 11: grpc.FeedChickenRequest
	(*FeedChickenResponse)(nil),        // 12: grpc.FeedChickenResponse
	(*FeedChickensOfBarnRequest)(nil),  // 13: grpc.FeedChickensOfBarnRequest
	(*FeedChickensOfBarnResponse)(nil), // 14: grpc.FeedChickensOfBarnResponse
}
var file_farm_proto_farm_proto_depIdxs = []int32{
	1,  // 0: grpc.Farm.barns:type_name -> grpc.Barn
	2,  // 1: grpc.Barn.chickens:type_name -> grpc.Chicken
	0,  // 2: grpc.GetFarmResponse.farm:type_name -> grpc.Farm
	3,  // 3: grpc.FarmService.GetFarm:input_type -> grpc.GetFarmRequest
	5,  // 4: grpc.FarmService.BuyBarn:input_type -> grpc.BuyBarnRequest
	7,  // 5: grpc.FarmService.BuyFeed:input_type -> grpc.BuyFeedRequest
	9,  // 6: grpc.FarmService.BuyChicken:input_type -> grpc.BuyChickenRequest
	11, // 7: grpc.FarmService.FeedChicken:input_type -> grpc.FeedChickenRequest
	13, // 8: grpc.FarmService.FeedChickensOfBarn:input_type -> grpc.FeedChickensOfBarnRequest
	4,  // 9: grpc.FarmService.GetFarm:output_type -> grpc.GetFarmResponse
	6,  // 10: grpc.FarmService.BuyBarn:output_type -> grpc.BuyBarnResponse
	8,  // 11: grpc.FarmService.BuyFeed:output_type -> grpc.BuyFeedResponse
	10, // 12: grpc.FarmService.BuyChicken:output_type -> grpc.BuyChickenResponse
	12, // 13: grpc.FarmService.FeedChicken:output_type -> grpc.FeedChickenResponse
	14, // 14: grpc.FarmService.FeedChickensOfBarn:output_type -> grpc.FeedChickensOfBarnResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_farm_proto_farm_proto_init() }
func file_farm_proto_farm_proto_init() {
	if File_farm_proto_farm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_farm_proto_farm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Farm); i {
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
		file_farm_proto_farm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Barn); i {
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
		file_farm_proto_farm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chicken); i {
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
		file_farm_proto_farm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFarmRequest); i {
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
		file_farm_proto_farm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFarmResponse); i {
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
		file_farm_proto_farm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyBarnRequest); i {
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
		file_farm_proto_farm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyBarnResponse); i {
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
		file_farm_proto_farm_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyFeedRequest); i {
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
		file_farm_proto_farm_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyFeedResponse); i {
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
		file_farm_proto_farm_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyChickenRequest); i {
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
		file_farm_proto_farm_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyChickenResponse); i {
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
		file_farm_proto_farm_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedChickenRequest); i {
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
		file_farm_proto_farm_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedChickenResponse); i {
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
		file_farm_proto_farm_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedChickensOfBarnRequest); i {
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
		file_farm_proto_farm_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedChickensOfBarnResponse); i {
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
			RawDescriptor: file_farm_proto_farm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_farm_proto_farm_proto_goTypes,
		DependencyIndexes: file_farm_proto_farm_proto_depIdxs,
		MessageInfos:      file_farm_proto_farm_proto_msgTypes,
	}.Build()
	File_farm_proto_farm_proto = out.File
	file_farm_proto_farm_proto_rawDesc = nil
	file_farm_proto_farm_proto_goTypes = nil
	file_farm_proto_farm_proto_depIdxs = nil
}