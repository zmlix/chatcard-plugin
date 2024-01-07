// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: plugin/plugin.proto

package plugin

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

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAILED  Status = 1
	Status_PROCESS Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
		2: "PROCESS",
	}
	Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
		"PROCESS": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_plugin_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_plugin_plugin_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

type PluginOptionValueType int32

const (
	PluginOptionValueType_STRING PluginOptionValueType = 0
	PluginOptionValueType_NUMBER PluginOptionValueType = 1
	PluginOptionValueType_BOOL   PluginOptionValueType = 2
)

// Enum value maps for PluginOptionValueType.
var (
	PluginOptionValueType_name = map[int32]string{
		0: "STRING",
		1: "NUMBER",
		2: "BOOL",
	}
	PluginOptionValueType_value = map[string]int32{
		"STRING": 0,
		"NUMBER": 1,
		"BOOL":   2,
	}
)

func (x PluginOptionValueType) Enum() *PluginOptionValueType {
	p := new(PluginOptionValueType)
	*p = x
	return p
}

func (x PluginOptionValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginOptionValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_plugin_proto_enumTypes[1].Descriptor()
}

func (PluginOptionValueType) Type() protoreflect.EnumType {
	return &file_plugin_plugin_proto_enumTypes[1]
}

func (x PluginOptionValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginOptionValueType.Descriptor instead.
func (PluginOptionValueType) EnumDescriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{1}
}

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Call      string  `protobuf:"bytes,2,opt,name=call,proto3" json:"call,omitempty"`
	Arguments *string `protobuf:"bytes,3,opt,name=arguments,proto3,oneof" json:"arguments,omitempty"`
	Token     string  `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Options   *string `protobuf:"bytes,5,opt,name=options,proto3,oneof" json:"options,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CallRequest) GetCall() string {
	if x != nil {
		return x.Call
	}
	return ""
}

func (x *CallRequest) GetArguments() string {
	if x != nil && x.Arguments != nil {
		return *x.Arguments
	}
	return ""
}

func (x *CallRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CallRequest) GetOptions() string {
	if x != nil && x.Options != nil {
		return *x.Options
	}
	return ""
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

func (x *CallResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type PluginOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  PluginOptionValueType `protobuf:"varint,1,opt,name=type,proto3,enum=PluginOptionValueType" json:"type,omitempty"`
	Key   string                `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value string                `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PluginOption) Reset() {
	*x = PluginOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginOption) ProtoMessage() {}

func (x *PluginOption) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginOption.ProtoReflect.Descriptor instead.
func (*PluginOption) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *PluginOption) GetType() PluginOptionValueType {
	if x != nil {
		return x.Type
	}
	return PluginOptionValueType_STRING
}

func (x *PluginOption) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PluginOption) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Display *string         `protobuf:"bytes,2,opt,name=display,proto3,oneof" json:"display,omitempty"`
	Img     *string         `protobuf:"bytes,3,opt,name=img,proto3,oneof" json:"img,omitempty"`
	Version string          `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Info    *string         `protobuf:"bytes,5,opt,name=info,proto3,oneof" json:"info,omitempty"`
	Options []*PluginOption `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetDisplay() string {
	if x != nil && x.Display != nil {
		return *x.Display
	}
	return ""
}

func (x *Plugin) GetImg() string {
	if x != nil && x.Img != nil {
		return *x.Img
	}
	return ""
}

func (x *Plugin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Plugin) GetInfo() string {
	if x != nil && x.Info != nil {
		return *x.Info
	}
	return ""
}

func (x *Plugin) GetOptions() []*PluginOption {
	if x != nil {
		return x.Options
	}
	return nil
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{4}
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    Status    `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	Plugins   []*Plugin `protobuf:"bytes,2,rep,name=plugins,proto3" json:"plugins,omitempty"`
	Directory string    `protobuf:"bytes,3,opt,name=directory,proto3" json:"directory,omitempty"`
	Web       string    `protobuf:"bytes,4,opt,name=web,proto3" json:"web,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *ConnectResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

func (x *ConnectResponse) GetPlugins() []*Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

func (x *ConnectResponse) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

func (x *ConnectResponse) GetWeb() string {
	if x != nil {
		return x.Web
	}
	return ""
}

type DirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Paths []string `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *DirectoryRequest) Reset() {
	*x = DirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryRequest) ProtoMessage() {}

func (x *DirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryRequest.ProtoReflect.Descriptor instead.
func (*DirectoryRequest) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *DirectoryRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *DirectoryRequest) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

type DirectoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	Directory string `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *DirectoryResponse) Reset() {
	*x = DirectoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryResponse) ProtoMessage() {}

func (x *DirectoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryResponse.ProtoReflect.Descriptor instead.
func (*DirectoryResponse) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{7}
}

func (x *DirectoryResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

func (x *DirectoryResponse) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

var File_plugin_plugin_proto protoreflect.FileDescriptor

var file_plugin_plugin_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x61, 0x6c,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x21, 0x0a,
	0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x4b, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62, 0x0a, 0x0c,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xcb, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x88, 0x01, 0x01, 0x12, 0x15,
	0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x69,
	0x6d, 0x67, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x69, 0x6d, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x10,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x85, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x65, 0x62, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x65, 0x62, 0x22, 0x3e, 0x0a, 0x10, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0x52, 0x0a, 0x11, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2a, 0x2e, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x2a, 0x39, 0x0a, 0x15,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x02, 0x32, 0x98, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12,
	0x0c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x32,
	0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x63, 0x68, 0x61, 0x74, 0x63, 0x61, 0x72, 0x64, 0x2d, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x3b,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_plugin_proto_rawDescOnce sync.Once
	file_plugin_plugin_proto_rawDescData = file_plugin_plugin_proto_rawDesc
)

func file_plugin_plugin_proto_rawDescGZIP() []byte {
	file_plugin_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_plugin_proto_rawDescData)
	})
	return file_plugin_plugin_proto_rawDescData
}

var file_plugin_plugin_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_plugin_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_plugin_plugin_proto_goTypes = []interface{}{
	(Status)(0),                // 0: Status
	(PluginOptionValueType)(0), // 1: PluginOptionValueType
	(*CallRequest)(nil),        // 2: CallRequest
	(*CallResponse)(nil),       // 3: CallResponse
	(*PluginOption)(nil),       // 4: PluginOption
	(*Plugin)(nil),             // 5: Plugin
	(*ConnectRequest)(nil),     // 6: ConnectRequest
	(*ConnectResponse)(nil),    // 7: ConnectResponse
	(*DirectoryRequest)(nil),   // 8: DirectoryRequest
	(*DirectoryResponse)(nil),  // 9: DirectoryResponse
}
var file_plugin_plugin_proto_depIdxs = []int32{
	0, // 0: CallResponse.status:type_name -> Status
	1, // 1: PluginOption.type:type_name -> PluginOptionValueType
	4, // 2: Plugin.options:type_name -> PluginOption
	0, // 3: ConnectResponse.status:type_name -> Status
	5, // 4: ConnectResponse.plugins:type_name -> Plugin
	0, // 5: DirectoryResponse.status:type_name -> Status
	6, // 6: PluginService.Connect:input_type -> ConnectRequest
	2, // 7: PluginService.Call:input_type -> CallRequest
	8, // 8: PluginService.Directory:input_type -> DirectoryRequest
	7, // 9: PluginService.Connect:output_type -> ConnectResponse
	3, // 10: PluginService.Call:output_type -> CallResponse
	9, // 11: PluginService.Directory:output_type -> DirectoryResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_plugin_plugin_proto_init() }
func file_plugin_plugin_proto_init() {
	if File_plugin_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_plugin_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
		file_plugin_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginOption); i {
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
		file_plugin_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_plugin_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_plugin_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_plugin_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryRequest); i {
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
		file_plugin_plugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryResponse); i {
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
	file_plugin_plugin_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_plugin_plugin_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugin_plugin_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_plugin_proto_depIdxs,
		EnumInfos:         file_plugin_plugin_proto_enumTypes,
		MessageInfos:      file_plugin_plugin_proto_msgTypes,
	}.Build()
	File_plugin_plugin_proto = out.File
	file_plugin_plugin_proto_rawDesc = nil
	file_plugin_plugin_proto_goTypes = nil
	file_plugin_plugin_proto_depIdxs = nil
}
