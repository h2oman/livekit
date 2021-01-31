// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: internal.proto

package livekit

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip      string     `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	NumCpus uint32     `protobuf:"varint,3,opt,name=num_cpus,json=numCpus,proto3" json:"num_cpus,omitempty"`
	Stats   *NodeStats `protobuf:"bytes,4,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Node) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Node) GetNumCpus() uint32 {
	if x != nil {
		return x.NumCpus
	}
	return 0
}

func (x *Node) GetStats() *NodeStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type NodeStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// when server was started
	StartedAt int64 `protobuf:"varint,1,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// when server last reported its status
	UpdatedAt    int64  `protobuf:"varint,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	NumRooms     uint32 `protobuf:"varint,3,opt,name=num_rooms,json=numRooms,proto3" json:"num_rooms,omitempty"`
	NumClients   uint32 `protobuf:"varint,4,opt,name=num_clients,json=numClients,proto3" json:"num_clients,omitempty"`
	NumTracksIn  uint32 `protobuf:"varint,5,opt,name=num_tracks_in,json=numTracksIn,proto3" json:"num_tracks_in,omitempty"`
	NumTracksOut uint32 `protobuf:"varint,6,opt,name=num_tracks_out,json=numTracksOut,proto3" json:"num_tracks_out,omitempty"`
}

func (x *NodeStats) Reset() {
	*x = NodeStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStats) ProtoMessage() {}

func (x *NodeStats) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStats.ProtoReflect.Descriptor instead.
func (*NodeStats) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{1}
}

func (x *NodeStats) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *NodeStats) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *NodeStats) GetNumRooms() uint32 {
	if x != nil {
		return x.NumRooms
	}
	return 0
}

func (x *NodeStats) GetNumClients() uint32 {
	if x != nil {
		return x.NumClients
	}
	return 0
}

func (x *NodeStats) GetNumTracksIn() uint32 {
	if x != nil {
		return x.NumTracksIn
	}
	return 0
}

func (x *NodeStats) GetNumTracksOut() uint32 {
	if x != nil {
		return x.NumTracksOut
	}
	return 0
}

// message to RTC nodes
type RTCNodeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantKey string `protobuf:"bytes,1,opt,name=participant_key,json=participantKey,proto3" json:"participant_key,omitempty"`
	// Types that are assignable to Message:
	//	*RTCNodeMessage_StartSession
	//	*RTCNodeMessage_Request
	Message isRTCNodeMessage_Message `protobuf_oneof:"message"`
}

func (x *RTCNodeMessage) Reset() {
	*x = RTCNodeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RTCNodeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RTCNodeMessage) ProtoMessage() {}

func (x *RTCNodeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RTCNodeMessage.ProtoReflect.Descriptor instead.
func (*RTCNodeMessage) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{2}
}

func (x *RTCNodeMessage) GetParticipantKey() string {
	if x != nil {
		return x.ParticipantKey
	}
	return ""
}

func (m *RTCNodeMessage) GetMessage() isRTCNodeMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *RTCNodeMessage) GetStartSession() *StartSession {
	if x, ok := x.GetMessage().(*RTCNodeMessage_StartSession); ok {
		return x.StartSession
	}
	return nil
}

func (x *RTCNodeMessage) GetRequest() *SignalRequest {
	if x, ok := x.GetMessage().(*RTCNodeMessage_Request); ok {
		return x.Request
	}
	return nil
}

type isRTCNodeMessage_Message interface {
	isRTCNodeMessage_Message()
}

type RTCNodeMessage_StartSession struct {
	StartSession *StartSession `protobuf:"bytes,2,opt,name=start_session,json=startSession,proto3,oneof"`
}

type RTCNodeMessage_Request struct {
	Request *SignalRequest `protobuf:"bytes,3,opt,name=request,proto3,oneof"`
}

func (*RTCNodeMessage_StartSession) isRTCNodeMessage_Message() {}

func (*RTCNodeMessage_Request) isRTCNodeMessage_Message() {}

// message to Signal nodes
type SignalNodeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// Types that are assignable to Message:
	//	*SignalNodeMessage_Response
	//	*SignalNodeMessage_EndSession
	Message isSignalNodeMessage_Message `protobuf_oneof:"message"`
}

func (x *SignalNodeMessage) Reset() {
	*x = SignalNodeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalNodeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalNodeMessage) ProtoMessage() {}

func (x *SignalNodeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalNodeMessage.ProtoReflect.Descriptor instead.
func (*SignalNodeMessage) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{3}
}

func (x *SignalNodeMessage) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (m *SignalNodeMessage) GetMessage() isSignalNodeMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *SignalNodeMessage) GetResponse() *SignalResponse {
	if x, ok := x.GetMessage().(*SignalNodeMessage_Response); ok {
		return x.Response
	}
	return nil
}

func (x *SignalNodeMessage) GetEndSession() *EndSession {
	if x, ok := x.GetMessage().(*SignalNodeMessage_EndSession); ok {
		return x.EndSession
	}
	return nil
}

type isSignalNodeMessage_Message interface {
	isSignalNodeMessage_Message()
}

type SignalNodeMessage_Response struct {
	Response *SignalResponse `protobuf:"bytes,2,opt,name=response,proto3,oneof"`
}

type SignalNodeMessage_EndSession struct {
	EndSession *EndSession `protobuf:"bytes,3,opt,name=end_session,json=endSession,proto3,oneof"`
}

func (*SignalNodeMessage_Response) isSignalNodeMessage_Message() {}

func (*SignalNodeMessage_EndSession) isSignalNodeMessage_Message() {}

type StartSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomName     string `protobuf:"bytes,1,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	Identity     string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	ConnectionId string `protobuf:"bytes,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
}

func (x *StartSession) Reset() {
	*x = StartSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSession) ProtoMessage() {}

func (x *StartSession) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSession.ProtoReflect.Descriptor instead.
func (*StartSession) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{4}
}

func (x *StartSession) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *StartSession) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *StartSession) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

type EndSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndSession) Reset() {
	*x = EndSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSession) ProtoMessage() {}

func (x *EndSession) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSession.ProtoReflect.Descriptor instead.
func (*EndSession) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{5}
}

var File_internal_proto protoreflect.FileDescriptor

var file_internal_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x1a, 0x09, 0x72, 0x74, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x70, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x6e, 0x75, 0x6d, 0x43, 0x70, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x22, 0xd1, 0x01, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x6e, 0x75, 0x6d, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75,
	0x6d, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x6e, 0x75, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6e,
	0x75, 0x6d, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x49, 0x6e, 0x12,
	0x24, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x5f, 0x6f, 0x75,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x73, 0x4f, 0x75, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x0e, 0x52, 0x54, 0x43, 0x4e, 0x6f, 0x64,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x3c, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x32, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb2,
	0x01, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x6e,
	0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x6c, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_rawDescOnce sync.Once
	file_internal_proto_rawDescData = file_internal_proto_rawDesc
)

func file_internal_proto_rawDescGZIP() []byte {
	file_internal_proto_rawDescOnce.Do(func() {
		file_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_rawDescData)
	})
	return file_internal_proto_rawDescData
}

var file_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_proto_goTypes = []interface{}{
	(*Node)(nil),              // 0: livekit.Node
	(*NodeStats)(nil),         // 1: livekit.NodeStats
	(*RTCNodeMessage)(nil),    // 2: livekit.RTCNodeMessage
	(*SignalNodeMessage)(nil), // 3: livekit.SignalNodeMessage
	(*StartSession)(nil),      // 4: livekit.StartSession
	(*EndSession)(nil),        // 5: livekit.EndSession
	(*SignalRequest)(nil),     // 6: livekit.SignalRequest
	(*SignalResponse)(nil),    // 7: livekit.SignalResponse
}
var file_internal_proto_depIdxs = []int32{
	1, // 0: livekit.Node.stats:type_name -> livekit.NodeStats
	4, // 1: livekit.RTCNodeMessage.start_session:type_name -> livekit.StartSession
	6, // 2: livekit.RTCNodeMessage.request:type_name -> livekit.SignalRequest
	7, // 3: livekit.SignalNodeMessage.response:type_name -> livekit.SignalResponse
	5, // 4: livekit.SignalNodeMessage.end_session:type_name -> livekit.EndSession
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_proto_init() }
func file_internal_proto_init() {
	if File_internal_proto != nil {
		return
	}
	file_rtc_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStats); i {
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
		file_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RTCNodeMessage); i {
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
		file_internal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalNodeMessage); i {
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
		file_internal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartSession); i {
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
		file_internal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSession); i {
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
	file_internal_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*RTCNodeMessage_StartSession)(nil),
		(*RTCNodeMessage_Request)(nil),
	}
	file_internal_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SignalNodeMessage_Response)(nil),
		(*SignalNodeMessage_EndSession)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_proto_goTypes,
		DependencyIndexes: file_internal_proto_depIdxs,
		MessageInfos:      file_internal_proto_msgTypes,
	}.Build()
	File_internal_proto = out.File
	file_internal_proto_rawDesc = nil
	file_internal_proto_goTypes = nil
	file_internal_proto_depIdxs = nil
}
