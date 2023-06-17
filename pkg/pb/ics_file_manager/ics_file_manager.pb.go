// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: ics_file_manager/ics_file_manager.proto

package icsfilemanager

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ICSFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GameId int32  `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ICSFile) Reset() {
	*x = ICSFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ICSFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ICSFile) ProtoMessage() {}

func (x *ICSFile) ProtoReflect() protoreflect.Message {
	mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ICSFile.ProtoReflect.Descriptor instead.
func (*ICSFile) Descriptor() ([]byte, []int) {
	return file_ics_file_manager_ics_file_manager_proto_rawDescGZIP(), []int{0}
}

func (x *ICSFile) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ICSFile) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *ICSFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateICSFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IcsFile *ICSFile `protobuf:"bytes,1,opt,name=ics_file,json=icsFile,proto3" json:"ics_file,omitempty"`
}

func (x *CreateICSFileRequest) Reset() {
	*x = CreateICSFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateICSFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateICSFileRequest) ProtoMessage() {}

func (x *CreateICSFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateICSFileRequest.ProtoReflect.Descriptor instead.
func (*CreateICSFileRequest) Descriptor() ([]byte, []int) {
	return file_ics_file_manager_ics_file_manager_proto_rawDescGZIP(), []int{1}
}

func (x *CreateICSFileRequest) GetIcsFile() *ICSFile {
	if x != nil {
		return x.IcsFile
	}
	return nil
}

type DeleteICSFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteICSFileRequest) Reset() {
	*x = DeleteICSFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteICSFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteICSFileRequest) ProtoMessage() {}

func (x *DeleteICSFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteICSFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteICSFileRequest) Descriptor() ([]byte, []int) {
	return file_ics_file_manager_ics_file_manager_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteICSFileRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetICSFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetICSFileRequest) Reset() {
	*x = GetICSFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetICSFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetICSFileRequest) ProtoMessage() {}

func (x *GetICSFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetICSFileRequest.ProtoReflect.Descriptor instead.
func (*GetICSFileRequest) Descriptor() ([]byte, []int) {
	return file_ics_file_manager_ics_file_manager_proto_rawDescGZIP(), []int{3}
}

func (x *GetICSFileRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetICSFileByGameIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId int32 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *GetICSFileByGameIDRequest) Reset() {
	*x = GetICSFileByGameIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetICSFileByGameIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetICSFileByGameIDRequest) ProtoMessage() {}

func (x *GetICSFileByGameIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetICSFileByGameIDRequest.ProtoReflect.Descriptor instead.
func (*GetICSFileByGameIDRequest) Descriptor() ([]byte, []int) {
	return file_ics_file_manager_ics_file_manager_proto_rawDescGZIP(), []int{4}
}

func (x *GetICSFileByGameIDRequest) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

type ListICSFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IcsFiles []*ICSFile `protobuf:"bytes,1,rep,name=ics_files,json=icsFiles,proto3" json:"ics_files,omitempty"`
}

func (x *ListICSFilesResponse) Reset() {
	*x = ListICSFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListICSFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListICSFilesResponse) ProtoMessage() {}

func (x *ListICSFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ics_file_manager_ics_file_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListICSFilesResponse.ProtoReflect.Descriptor instead.
func (*ListICSFilesResponse) Descriptor() ([]byte, []int) {
	return file_ics_file_manager_ics_file_manager_proto_rawDescGZIP(), []int{5}
}

func (x *ListICSFilesResponse) GetIcsFiles() []*ICSFile {
	if x != nil {
		return x.IcsFiles
	}
	return nil
}

var File_ics_file_manager_ics_file_manager_proto protoreflect.FileDescriptor

var file_ics_file_manager_ics_file_manager_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x69, 0x63, 0x73, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x07, 0x49, 0x43, 0x53, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x4c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x43, 0x53, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x69, 0x63, 0x73, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x63, 0x73,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x43,
	0x53, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x69, 0x63, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x26,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x43, 0x53, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x43, 0x53,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x49, 0x43, 0x53, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x22, 0x4e, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x43, 0x53, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x69, 0x63, 0x73,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69,
	0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x49, 0x43, 0x53, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x69, 0x63, 0x73, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x32, 0xb4, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x43, 0x53, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x26,
	0x2e, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x43, 0x53, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x43, 0x53, 0x46, 0x69, 0x6c,
	0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x43, 0x53,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x2e, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x43,
	0x53, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x43, 0x53,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x43, 0x53, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x63, 0x73, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x43, 0x53,
	0x46, 0x69, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x43, 0x53,
	0x46, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x2b, 0x2e, 0x69,
	0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x43, 0x53, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x47, 0x61, 0x6d, 0x65,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x63, 0x73, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x43, 0x53,
	0x46, 0x69, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x43,
	0x53, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26,
	0x2e, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x43, 0x53, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6b, 0x69, 0x74, 0x61, 0x35, 0x36, 0x33,
	0x37, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x2d, 0x69, 0x63, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x69, 0x63,
	0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3b, 0x69,
	0x63, 0x73, 0x66, 0x69, 0x6c, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ics_file_manager_ics_file_manager_proto_rawDescOnce sync.Once
	file_ics_file_manager_ics_file_manager_proto_rawDescData = file_ics_file_manager_ics_file_manager_proto_rawDesc
)

func file_ics_file_manager_ics_file_manager_proto_rawDescGZIP() []byte {
	file_ics_file_manager_ics_file_manager_proto_rawDescOnce.Do(func() {
		file_ics_file_manager_ics_file_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_ics_file_manager_ics_file_manager_proto_rawDescData)
	})
	return file_ics_file_manager_ics_file_manager_proto_rawDescData
}

var file_ics_file_manager_ics_file_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ics_file_manager_ics_file_manager_proto_goTypes = []interface{}{
	(*ICSFile)(nil),                   // 0: ics_file_manager.ICSFile
	(*CreateICSFileRequest)(nil),      // 1: ics_file_manager.CreateICSFileRequest
	(*DeleteICSFileRequest)(nil),      // 2: ics_file_manager.DeleteICSFileRequest
	(*GetICSFileRequest)(nil),         // 3: ics_file_manager.GetICSFileRequest
	(*GetICSFileByGameIDRequest)(nil), // 4: ics_file_manager.GetICSFileByGameIDRequest
	(*ListICSFilesResponse)(nil),      // 5: ics_file_manager.ListICSFilesResponse
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_ics_file_manager_ics_file_manager_proto_depIdxs = []int32{
	0, // 0: ics_file_manager.CreateICSFileRequest.ics_file:type_name -> ics_file_manager.ICSFile
	0, // 1: ics_file_manager.ListICSFilesResponse.ics_files:type_name -> ics_file_manager.ICSFile
	1, // 2: ics_file_manager.Service.CreateICSFile:input_type -> ics_file_manager.CreateICSFileRequest
	2, // 3: ics_file_manager.Service.DeleteICSFile:input_type -> ics_file_manager.DeleteICSFileRequest
	3, // 4: ics_file_manager.Service.GetICSFile:input_type -> ics_file_manager.GetICSFileRequest
	4, // 5: ics_file_manager.Service.GetICSFileByGameID:input_type -> ics_file_manager.GetICSFileByGameIDRequest
	6, // 6: ics_file_manager.Service.ListICSFiles:input_type -> google.protobuf.Empty
	0, // 7: ics_file_manager.Service.CreateICSFile:output_type -> ics_file_manager.ICSFile
	6, // 8: ics_file_manager.Service.DeleteICSFile:output_type -> google.protobuf.Empty
	0, // 9: ics_file_manager.Service.GetICSFile:output_type -> ics_file_manager.ICSFile
	0, // 10: ics_file_manager.Service.GetICSFileByGameID:output_type -> ics_file_manager.ICSFile
	5, // 11: ics_file_manager.Service.ListICSFiles:output_type -> ics_file_manager.ListICSFilesResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ics_file_manager_ics_file_manager_proto_init() }
func file_ics_file_manager_ics_file_manager_proto_init() {
	if File_ics_file_manager_ics_file_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ics_file_manager_ics_file_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ICSFile); i {
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
		file_ics_file_manager_ics_file_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateICSFileRequest); i {
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
		file_ics_file_manager_ics_file_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteICSFileRequest); i {
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
		file_ics_file_manager_ics_file_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetICSFileRequest); i {
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
		file_ics_file_manager_ics_file_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetICSFileByGameIDRequest); i {
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
		file_ics_file_manager_ics_file_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListICSFilesResponse); i {
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
			RawDescriptor: file_ics_file_manager_ics_file_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ics_file_manager_ics_file_manager_proto_goTypes,
		DependencyIndexes: file_ics_file_manager_ics_file_manager_proto_depIdxs,
		MessageInfos:      file_ics_file_manager_ics_file_manager_proto_msgTypes,
	}.Build()
	File_ics_file_manager_ics_file_manager_proto = out.File
	file_ics_file_manager_ics_file_manager_proto_rawDesc = nil
	file_ics_file_manager_ics_file_manager_proto_goTypes = nil
	file_ics_file_manager_ics_file_manager_proto_depIdxs = nil
}