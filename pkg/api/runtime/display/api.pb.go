// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: pkg/api/runtime/display/api.proto

package display

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type GraphResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Graph string `protobuf:"bytes,1,opt,name=graph,proto3" json:"graph,omitempty"`
}

func (x *GraphResponse) Reset() {
	*x = GraphResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_runtime_display_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphResponse) ProtoMessage() {}

func (x *GraphResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_runtime_display_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphResponse.ProtoReflect.Descriptor instead.
func (*GraphResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_runtime_display_api_proto_rawDescGZIP(), []int{0}
}

func (x *GraphResponse) GetGraph() string {
	if x != nil {
		return x.Graph
	}
	return ""
}

type UsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usage map[string]*ContainerStat `protobuf:"bytes,1,rep,name=usage,proto3" json:"usage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UsageResponse) Reset() {
	*x = UsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_runtime_display_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageResponse) ProtoMessage() {}

func (x *UsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_runtime_display_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageResponse.ProtoReflect.Descriptor instead.
func (*UsageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_runtime_display_api_proto_rawDescGZIP(), []int{1}
}

func (x *UsageResponse) GetUsage() map[string]*ContainerStat {
	if x != nil {
		return x.Usage
	}
	return nil
}

type ContainerStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stat    map[string]*Devices `protobuf:"bytes,1,rep,name=stat,proto3" json:"stat,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Project string              `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	User    string              `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Cluster string              `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Spec    map[string]*Spec    `protobuf:"bytes,5,rep,name=spec,proto3" json:"spec,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ContainerStat) Reset() {
	*x = ContainerStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_runtime_display_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerStat) ProtoMessage() {}

func (x *ContainerStat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_runtime_display_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerStat.ProtoReflect.Descriptor instead.
func (*ContainerStat) Descriptor() ([]byte, []int) {
	return file_pkg_api_runtime_display_api_proto_rawDescGZIP(), []int{2}
}

func (x *ContainerStat) GetStat() map[string]*Devices {
	if x != nil {
		return x.Stat
	}
	return nil
}

func (x *ContainerStat) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ContainerStat) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ContainerStat) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *ContainerStat) GetSpec() map[string]*Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type Devices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dev []*DeviceInfo `protobuf:"bytes,1,rep,name=dev,proto3" json:"dev,omitempty"`
}

func (x *Devices) Reset() {
	*x = Devices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_runtime_display_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Devices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Devices) ProtoMessage() {}

func (x *Devices) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_runtime_display_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Devices.ProtoReflect.Descriptor instead.
func (*Devices) Descriptor() ([]byte, []int) {
	return file_pkg_api_runtime_display_api_proto_rawDescGZIP(), []int{3}
}

func (x *Devices) GetDev() []*DeviceInfo {
	if x != nil {
		return x.Dev
	}
	return nil
}

type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CardIdx   string  `protobuf:"bytes,2,opt,name=card_idx,json=cardIdx,proto3" json:"card_idx,omitempty"`
	Gpu       float32 `protobuf:"fixed32,10,opt,name=gpu,proto3" json:"gpu,omitempty"`
	Mem       float32 `protobuf:"fixed32,11,opt,name=mem,proto3" json:"mem,omitempty"`
	Pids      []int32 `protobuf:"varint,12,rep,packed,name=pids,proto3" json:"pids,omitempty"`
	DeviceMem float32 `protobuf:"fixed32,13,opt,name=device_mem,json=deviceMem,proto3" json:"device_mem,omitempty"`
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_runtime_display_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_runtime_display_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_pkg_api_runtime_display_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceInfo) GetCardIdx() string {
	if x != nil {
		return x.CardIdx
	}
	return ""
}

func (x *DeviceInfo) GetGpu() float32 {
	if x != nil {
		return x.Gpu
	}
	return 0
}

func (x *DeviceInfo) GetMem() float32 {
	if x != nil {
		return x.Mem
	}
	return 0
}

func (x *DeviceInfo) GetPids() []int32 {
	if x != nil {
		return x.Pids
	}
	return nil
}

func (x *DeviceInfo) GetDeviceMem() float32 {
	if x != nil {
		return x.DeviceMem
	}
	return 0
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_runtime_display_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_runtime_display_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_runtime_display_api_proto_rawDescGZIP(), []int{5}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gpu float32 `protobuf:"fixed32,1,opt,name=gpu,proto3" json:"gpu,omitempty"`
	Mem float32 `protobuf:"fixed32,2,opt,name=mem,proto3" json:"mem,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_runtime_display_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_runtime_display_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_pkg_api_runtime_display_api_proto_rawDescGZIP(), []int{6}
}

func (x *Spec) GetGpu() float32 {
	if x != nil {
		return x.Gpu
	}
	return 0
}

func (x *Spec) GetMem() float32 {
	if x != nil {
		return x.Mem
	}
	return 0
}

var File_pkg_api_runtime_display_api_proto protoreflect.FileDescriptor

var file_pkg_api_runtime_display_api_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x22, 0x9a,
	0x01, 0x0a, 0x0d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x50, 0x0a, 0x0a, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd6, 0x02, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x12, 0x34, 0x0a,
	0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x73,
	0x74, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x1a, 0x49, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x09,
	0x53, 0x70, 0x65, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x07, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x25, 0x0a, 0x03, 0x64, 0x65, 0x76, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x03, 0x64, 0x65, 0x76, 0x22, 0x8e, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x78,
	0x12, 0x10, 0x0a, 0x03, 0x67, 0x70, 0x75, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x67,
	0x70, 0x75, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x65, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x6d, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x69, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x22, 0x2b, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x10, 0x0a, 0x03,
	0x67, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x67, 0x70, 0x75, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x65, 0x6d,
	0x32, 0xf8, 0x01, 0x0a, 0x0a, 0x47, 0x50, 0x55, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12,
	0x4c, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x4d, 0x0a,
	0x0b, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x18, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x12, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x19, 0x5a, 0x17, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_runtime_display_api_proto_rawDescOnce sync.Once
	file_pkg_api_runtime_display_api_proto_rawDescData = file_pkg_api_runtime_display_api_proto_rawDesc
)

func file_pkg_api_runtime_display_api_proto_rawDescGZIP() []byte {
	file_pkg_api_runtime_display_api_proto_rawDescOnce.Do(func() {
		file_pkg_api_runtime_display_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_runtime_display_api_proto_rawDescData)
	})
	return file_pkg_api_runtime_display_api_proto_rawDescData
}

var file_pkg_api_runtime_display_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_api_runtime_display_api_proto_goTypes = []interface{}{
	(*GraphResponse)(nil),   // 0: display.GraphResponse
	(*UsageResponse)(nil),   // 1: display.UsageResponse
	(*ContainerStat)(nil),   // 2: display.ContainerStat
	(*Devices)(nil),         // 3: display.Devices
	(*DeviceInfo)(nil),      // 4: display.DeviceInfo
	(*VersionResponse)(nil), // 5: display.VersionResponse
	(*Spec)(nil),            // 6: display.Spec
	nil,                     // 7: display.UsageResponse.UsageEntry
	nil,                     // 8: display.ContainerStat.StatEntry
	nil,                     // 9: display.ContainerStat.SpecEntry
	(*empty.Empty)(nil),     // 10: google.protobuf.Empty
}
var file_pkg_api_runtime_display_api_proto_depIdxs = []int32{
	7,  // 0: display.UsageResponse.usage:type_name -> display.UsageResponse.UsageEntry
	8,  // 1: display.ContainerStat.stat:type_name -> display.ContainerStat.StatEntry
	9,  // 2: display.ContainerStat.spec:type_name -> display.ContainerStat.SpecEntry
	4,  // 3: display.Devices.dev:type_name -> display.DeviceInfo
	2,  // 4: display.UsageResponse.UsageEntry.value:type_name -> display.ContainerStat
	3,  // 5: display.ContainerStat.StatEntry.value:type_name -> display.Devices
	6,  // 6: display.ContainerStat.SpecEntry.value:type_name -> display.Spec
	10, // 7: display.GPUDisplay.PrintGraph:input_type -> google.protobuf.Empty
	10, // 8: display.GPUDisplay.PrintUsages:input_type -> google.protobuf.Empty
	10, // 9: display.GPUDisplay.Version:input_type -> google.protobuf.Empty
	0,  // 10: display.GPUDisplay.PrintGraph:output_type -> display.GraphResponse
	1,  // 11: display.GPUDisplay.PrintUsages:output_type -> display.UsageResponse
	5,  // 12: display.GPUDisplay.Version:output_type -> display.VersionResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_api_runtime_display_api_proto_init() }
func file_pkg_api_runtime_display_api_proto_init() {
	if File_pkg_api_runtime_display_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_runtime_display_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphResponse); i {
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
		file_pkg_api_runtime_display_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageResponse); i {
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
		file_pkg_api_runtime_display_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerStat); i {
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
		file_pkg_api_runtime_display_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Devices); i {
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
		file_pkg_api_runtime_display_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
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
		file_pkg_api_runtime_display_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_pkg_api_runtime_display_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
			RawDescriptor: file_pkg_api_runtime_display_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_runtime_display_api_proto_goTypes,
		DependencyIndexes: file_pkg_api_runtime_display_api_proto_depIdxs,
		MessageInfos:      file_pkg_api_runtime_display_api_proto_msgTypes,
	}.Build()
	File_pkg_api_runtime_display_api_proto = out.File
	file_pkg_api_runtime_display_api_proto_rawDesc = nil
	file_pkg_api_runtime_display_api_proto_goTypes = nil
	file_pkg_api_runtime_display_api_proto_depIdxs = nil
}
