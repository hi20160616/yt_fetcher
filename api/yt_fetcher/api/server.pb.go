// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: api/yt_fetcher/api/server.proto

package api

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

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastUpdated string   `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Vids        []string `protobuf:"bytes,4,rep,name=vids,proto3" json:"vids,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_yt_fetcher_api_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_api_yt_fetcher_api_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_api_yt_fetcher_api_server_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Channel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Channel) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *Channel) GetVids() []string {
	if x != nil {
		return x.Vids
	}
	return nil
}

type Channels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channels []*Channel `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
}

func (x *Channels) Reset() {
	*x = Channels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_yt_fetcher_api_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channels) ProtoMessage() {}

func (x *Channels) ProtoReflect() protoreflect.Message {
	mi := &file_api_yt_fetcher_api_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channels.ProtoReflect.Descriptor instead.
func (*Channels) Descriptor() ([]byte, []int) {
	return file_api_yt_fetcher_api_server_proto_rawDescGZIP(), []int{1}
}

func (x *Channels) GetChannels() []*Channel {
	if x != nil {
		return x.Channels
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Duration    string `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Cid         string `protobuf:"bytes,5,opt,name=cid,proto3" json:"cid,omitempty"`
	Cname       string `protobuf:"bytes,6,opt,name=cname,proto3" json:"cname,omitempty"`
	LastUpdated string `protobuf:"bytes,7,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_yt_fetcher_api_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_api_yt_fetcher_api_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_api_yt_fetcher_api_server_proto_rawDescGZIP(), []int{2}
}

func (x *Video) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Video) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Video) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *Video) GetCname() string {
	if x != nil {
		return x.Cname
	}
	return ""
}

func (x *Video) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

type Videos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After  string   `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"`
	Before string   `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	Videos []*Video `protobuf:"bytes,3,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *Videos) Reset() {
	*x = Videos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_yt_fetcher_api_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Videos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Videos) ProtoMessage() {}

func (x *Videos) ProtoReflect() protoreflect.Message {
	mi := &file_api_yt_fetcher_api_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Videos.ProtoReflect.Descriptor instead.
func (*Videos) Descriptor() ([]byte, []int) {
	return file_api_yt_fetcher_api_server_proto_rawDescGZIP(), []int{3}
}

func (x *Videos) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *Videos) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

func (x *Videos) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

var File_api_yt_fetcher_api_server_proto protoreflect.FileDescriptor

var file_api_yt_fetcher_api_server_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x22, 0x64, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x76, 0x69, 0x64, 0x73, 0x22, 0x3f, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x08,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x65, 0x0a, 0x06, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x79, 0x74, 0x5f, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x32, 0xde, 0x03, 0x0a, 0x0e, 0x59, 0x6f, 0x75,
	0x74, 0x75, 0x62, 0x65, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x73, 0x12, 0x17, 0x2e, 0x79, 0x74, 0x5f,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x1a, 0x17, 0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x17, 0x2e, 0x79, 0x74,
	0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x16, 0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x54,
	0x6f, 0x12, 0x16, 0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x1a, 0x16, 0x2e, 0x79, 0x74, 0x5f, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12,
	0x15, 0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x1a, 0x15, 0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x43, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x17, 0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x12, 0x18, 0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x1a, 0x18, 0x2e, 0x79, 0x74,
	0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x17, 0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x17,
	0x2e, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x00, 0x42, 0x65, 0x0a, 0x1b, 0x69, 0x6f, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x42, 0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47,
	0x75, 0x69, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x32, 0x30, 0x31, 0x36, 0x30, 0x36,
	0x31, 0x36, 0x2f, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x79, 0x74, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_yt_fetcher_api_server_proto_rawDescOnce sync.Once
	file_api_yt_fetcher_api_server_proto_rawDescData = file_api_yt_fetcher_api_server_proto_rawDesc
)

func file_api_yt_fetcher_api_server_proto_rawDescGZIP() []byte {
	file_api_yt_fetcher_api_server_proto_rawDescOnce.Do(func() {
		file_api_yt_fetcher_api_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_yt_fetcher_api_server_proto_rawDescData)
	})
	return file_api_yt_fetcher_api_server_proto_rawDescData
}

var file_api_yt_fetcher_api_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_yt_fetcher_api_server_proto_goTypes = []interface{}{
	(*Channel)(nil),  // 0: yt_fetcher.api.Channel
	(*Channels)(nil), // 1: yt_fetcher.api.Channels
	(*Video)(nil),    // 2: yt_fetcher.api.Video
	(*Videos)(nil),   // 3: yt_fetcher.api.Videos
}
var file_api_yt_fetcher_api_server_proto_depIdxs = []int32{
	0, // 0: yt_fetcher.api.Channels.channels:type_name -> yt_fetcher.api.Channel
	2, // 1: yt_fetcher.api.Videos.videos:type_name -> yt_fetcher.api.Video
	0, // 2: yt_fetcher.api.YoutubeFetcher.GetVideoIds:input_type -> yt_fetcher.api.Channel
	0, // 3: yt_fetcher.api.YoutubeFetcher.GetVideos:input_type -> yt_fetcher.api.Channel
	3, // 4: yt_fetcher.api.YoutubeFetcher.GetVideosFromTo:input_type -> yt_fetcher.api.Videos
	2, // 5: yt_fetcher.api.YoutubeFetcher.GetVideo:input_type -> yt_fetcher.api.Video
	0, // 6: yt_fetcher.api.YoutubeFetcher.GetSetCname:input_type -> yt_fetcher.api.Channel
	1, // 7: yt_fetcher.api.YoutubeFetcher.GetChannels:input_type -> yt_fetcher.api.Channels
	0, // 8: yt_fetcher.api.YoutubeFetcher.GetChannel:input_type -> yt_fetcher.api.Channel
	0, // 9: yt_fetcher.api.YoutubeFetcher.GetVideoIds:output_type -> yt_fetcher.api.Channel
	3, // 10: yt_fetcher.api.YoutubeFetcher.GetVideos:output_type -> yt_fetcher.api.Videos
	3, // 11: yt_fetcher.api.YoutubeFetcher.GetVideosFromTo:output_type -> yt_fetcher.api.Videos
	2, // 12: yt_fetcher.api.YoutubeFetcher.GetVideo:output_type -> yt_fetcher.api.Video
	0, // 13: yt_fetcher.api.YoutubeFetcher.GetSetCname:output_type -> yt_fetcher.api.Channel
	1, // 14: yt_fetcher.api.YoutubeFetcher.GetChannels:output_type -> yt_fetcher.api.Channels
	0, // 15: yt_fetcher.api.YoutubeFetcher.GetChannel:output_type -> yt_fetcher.api.Channel
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_yt_fetcher_api_server_proto_init() }
func file_api_yt_fetcher_api_server_proto_init() {
	if File_api_yt_fetcher_api_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_yt_fetcher_api_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_api_yt_fetcher_api_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channels); i {
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
		file_api_yt_fetcher_api_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_api_yt_fetcher_api_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Videos); i {
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
			RawDescriptor: file_api_yt_fetcher_api_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_yt_fetcher_api_server_proto_goTypes,
		DependencyIndexes: file_api_yt_fetcher_api_server_proto_depIdxs,
		MessageInfos:      file_api_yt_fetcher_api_server_proto_msgTypes,
	}.Build()
	File_api_yt_fetcher_api_server_proto = out.File
	file_api_yt_fetcher_api_server_proto_rawDesc = nil
	file_api_yt_fetcher_api_server_proto_goTypes = nil
	file_api_yt_fetcher_api_server_proto_depIdxs = nil
}
