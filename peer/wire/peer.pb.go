// Code generated by protoc-gen-go.
// source: bazil.org/bazil/peer/wire/peer.proto
// DO NOT EDIT!

/*
Package wire is a generated protocol buffer package.

It is generated from these files:
	bazil.org/bazil/peer/wire/peer.proto

It has these top-level messages:
	PingRequest
	PingResponse
	ObjectPutRequest
	ObjectPutResponse
	ObjectGetRequest
	ObjectGetResponse
	VolumeConnectRequest
	VolumeConnectResponse
	VolumeSyncPullRequest
	VolumeSyncPullItem
	Dirent
	File
	Dir
	Tombstone
*/
package wire

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bazil_cas "bazil.org/bazil/cas/wire"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VolumeSyncPullItem_Error int32

const (
	VolumeSyncPullItem_SUCCESS VolumeSyncPullItem_Error = 0
	// The path in the request did not refer to a directory.
	VolumeSyncPullItem_NOT_A_DIRECTORY VolumeSyncPullItem_Error = 1
)

var VolumeSyncPullItem_Error_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_A_DIRECTORY",
}
var VolumeSyncPullItem_Error_value = map[string]int32{
	"SUCCESS":         0,
	"NOT_A_DIRECTORY": 1,
}

func (x VolumeSyncPullItem_Error) String() string {
	return proto.EnumName(VolumeSyncPullItem_Error_name, int32(x))
}
func (VolumeSyncPullItem_Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PingResponse struct {
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ObjectPutRequest struct {
	// Only set in the first streamed message.
	Key  []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ObjectPutRequest) Reset()                    { *m = ObjectPutRequest{} }
func (m *ObjectPutRequest) String() string            { return proto.CompactTextString(m) }
func (*ObjectPutRequest) ProtoMessage()               {}
func (*ObjectPutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ObjectPutResponse struct {
}

func (m *ObjectPutResponse) Reset()                    { *m = ObjectPutResponse{} }
func (m *ObjectPutResponse) String() string            { return proto.CompactTextString(m) }
func (*ObjectPutResponse) ProtoMessage()               {}
func (*ObjectPutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ObjectGetRequest struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *ObjectGetRequest) Reset()                    { *m = ObjectGetRequest{} }
func (m *ObjectGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ObjectGetRequest) ProtoMessage()               {}
func (*ObjectGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ObjectGetResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ObjectGetResponse) Reset()                    { *m = ObjectGetResponse{} }
func (m *ObjectGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ObjectGetResponse) ProtoMessage()               {}
func (*ObjectGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type VolumeConnectRequest struct {
	VolumeName string `protobuf:"bytes,1,opt,name=volumeName" json:"volumeName,omitempty"`
}

func (m *VolumeConnectRequest) Reset()                    { *m = VolumeConnectRequest{} }
func (m *VolumeConnectRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeConnectRequest) ProtoMessage()               {}
func (*VolumeConnectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type VolumeConnectResponse struct {
	VolumeID []byte `protobuf:"bytes,1,opt,name=volumeID,proto3" json:"volumeID,omitempty"`
}

func (m *VolumeConnectResponse) Reset()                    { *m = VolumeConnectResponse{} }
func (m *VolumeConnectResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeConnectResponse) ProtoMessage()               {}
func (*VolumeConnectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type VolumeSyncPullRequest struct {
	VolumeID []byte `protobuf:"bytes,1,opt,name=volumeID,proto3" json:"volumeID,omitempty"`
	Path     string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *VolumeSyncPullRequest) Reset()                    { *m = VolumeSyncPullRequest{} }
func (m *VolumeSyncPullRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeSyncPullRequest) ProtoMessage()               {}
func (*VolumeSyncPullRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type VolumeSyncPullItem struct {
	// This is used to work around gRPC fixed error codes and error
	// strings.
	//
	// It can only be present in the first streamed message.
	// All other fields are to be ignored.
	Error VolumeSyncPullItem_Error `protobuf:"varint,1,opt,name=error,enum=bazil.peer.VolumeSyncPullItem_Error" json:"error,omitempty"`
	// Logical clocks in Dirents use small integers to identify peers.
	// This map connects those identifiers to globally unique peer
	// public keys.
	//
	// This can only be present in the first streamed message.
	Peers map[uint32][]byte `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Logical clock for the directory itself.
	//
	// This can only be present in the first streamed message.
	DirClock []byte `protobuf:"bytes,4,opt,name=dirClock,proto3" json:"dirClock,omitempty"`
	// Directory entries. More entries may follow in later streamed
	// messages. The entries are required to be in lexicographical
	// (bytewise) order, across all messages.
	Children []*Dirent `protobuf:"bytes,3,rep,name=children" json:"children,omitempty"`
}

func (m *VolumeSyncPullItem) Reset()                    { *m = VolumeSyncPullItem{} }
func (m *VolumeSyncPullItem) String() string            { return proto.CompactTextString(m) }
func (*VolumeSyncPullItem) ProtoMessage()               {}
func (*VolumeSyncPullItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *VolumeSyncPullItem) GetPeers() map[uint32][]byte {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *VolumeSyncPullItem) GetChildren() []*Dirent {
	if m != nil {
		return m.Children
	}
	return nil
}

type Dirent struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*Dirent_File
	//	*Dirent_Dir
	//	*Dirent_Tombstone
	Type  isDirent_Type `protobuf_oneof:"type"`
	Clock []byte        `protobuf:"bytes,4,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (m *Dirent) Reset()                    { *m = Dirent{} }
func (m *Dirent) String() string            { return proto.CompactTextString(m) }
func (*Dirent) ProtoMessage()               {}
func (*Dirent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type isDirent_Type interface {
	isDirent_Type()
}

type Dirent_File struct {
	File *File `protobuf:"bytes,2,opt,name=file,oneof"`
}
type Dirent_Dir struct {
	Dir *Dir `protobuf:"bytes,3,opt,name=dir,oneof"`
}
type Dirent_Tombstone struct {
	Tombstone *Tombstone `protobuf:"bytes,5,opt,name=tombstone,oneof"`
}

func (*Dirent_File) isDirent_Type()      {}
func (*Dirent_Dir) isDirent_Type()       {}
func (*Dirent_Tombstone) isDirent_Type() {}

func (m *Dirent) GetType() isDirent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Dirent) GetFile() *File {
	if x, ok := m.GetType().(*Dirent_File); ok {
		return x.File
	}
	return nil
}

func (m *Dirent) GetDir() *Dir {
	if x, ok := m.GetType().(*Dirent_Dir); ok {
		return x.Dir
	}
	return nil
}

func (m *Dirent) GetTombstone() *Tombstone {
	if x, ok := m.GetType().(*Dirent_Tombstone); ok {
		return x.Tombstone
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Dirent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _Dirent_OneofMarshaler, _Dirent_OneofUnmarshaler, []interface{}{
		(*Dirent_File)(nil),
		(*Dirent_Dir)(nil),
		(*Dirent_Tombstone)(nil),
	}
}

func _Dirent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Dirent)
	// type
	switch x := m.Type.(type) {
	case *Dirent_File:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.File); err != nil {
			return err
		}
	case *Dirent_Dir:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Dir); err != nil {
			return err
		}
	case *Dirent_Tombstone:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tombstone); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Dirent.Type has unexpected type %T", x)
	}
	return nil
}

func _Dirent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Dirent)
	switch tag {
	case 2: // type.file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(File)
		err := b.DecodeMessage(msg)
		m.Type = &Dirent_File{msg}
		return true, err
	case 3: // type.dir
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Dir)
		err := b.DecodeMessage(msg)
		m.Type = &Dirent_Dir{msg}
		return true, err
	case 5: // type.tombstone
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Tombstone)
		err := b.DecodeMessage(msg)
		m.Type = &Dirent_Tombstone{msg}
		return true, err
	default:
		return false, nil
	}
}

type File struct {
	Manifest *bazil_cas.Manifest `protobuf:"bytes,1,opt,name=manifest" json:"manifest,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *File) GetManifest() *bazil_cas.Manifest {
	if m != nil {
		return m.Manifest
	}
	return nil
}

type Dir struct {
}

func (m *Dir) Reset()                    { *m = Dir{} }
func (m *Dir) String() string            { return proto.CompactTextString(m) }
func (*Dir) ProtoMessage()               {}
func (*Dir) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type Tombstone struct {
}

func (m *Tombstone) Reset()                    { *m = Tombstone{} }
func (m *Tombstone) String() string            { return proto.CompactTextString(m) }
func (*Tombstone) ProtoMessage()               {}
func (*Tombstone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*PingRequest)(nil), "bazil.peer.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "bazil.peer.PingResponse")
	proto.RegisterType((*ObjectPutRequest)(nil), "bazil.peer.ObjectPutRequest")
	proto.RegisterType((*ObjectPutResponse)(nil), "bazil.peer.ObjectPutResponse")
	proto.RegisterType((*ObjectGetRequest)(nil), "bazil.peer.ObjectGetRequest")
	proto.RegisterType((*ObjectGetResponse)(nil), "bazil.peer.ObjectGetResponse")
	proto.RegisterType((*VolumeConnectRequest)(nil), "bazil.peer.VolumeConnectRequest")
	proto.RegisterType((*VolumeConnectResponse)(nil), "bazil.peer.VolumeConnectResponse")
	proto.RegisterType((*VolumeSyncPullRequest)(nil), "bazil.peer.VolumeSyncPullRequest")
	proto.RegisterType((*VolumeSyncPullItem)(nil), "bazil.peer.VolumeSyncPullItem")
	proto.RegisterType((*Dirent)(nil), "bazil.peer.Dirent")
	proto.RegisterType((*File)(nil), "bazil.peer.File")
	proto.RegisterType((*Dir)(nil), "bazil.peer.Dir")
	proto.RegisterType((*Tombstone)(nil), "bazil.peer.Tombstone")
	proto.RegisterEnum("bazil.peer.VolumeSyncPullItem_Error", VolumeSyncPullItem_Error_name, VolumeSyncPullItem_Error_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Peer service

type PeerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	ObjectPut(ctx context.Context, opts ...grpc.CallOption) (Peer_ObjectPutClient, error)
	ObjectGet(ctx context.Context, in *ObjectGetRequest, opts ...grpc.CallOption) (Peer_ObjectGetClient, error)
	VolumeConnect(ctx context.Context, in *VolumeConnectRequest, opts ...grpc.CallOption) (*VolumeConnectResponse, error)
	VolumeSyncPull(ctx context.Context, in *VolumeSyncPullRequest, opts ...grpc.CallOption) (Peer_VolumeSyncPullClient, error)
}

type peerClient struct {
	cc *grpc.ClientConn
}

func NewPeerClient(cc *grpc.ClientConn) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/bazil.peer.Peer/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) ObjectPut(ctx context.Context, opts ...grpc.CallOption) (Peer_ObjectPutClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Peer_serviceDesc.Streams[0], c.cc, "/bazil.peer.Peer/ObjectPut", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerObjectPutClient{stream}
	return x, nil
}

type Peer_ObjectPutClient interface {
	Send(*ObjectPutRequest) error
	CloseAndRecv() (*ObjectPutResponse, error)
	grpc.ClientStream
}

type peerObjectPutClient struct {
	grpc.ClientStream
}

func (x *peerObjectPutClient) Send(m *ObjectPutRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *peerObjectPutClient) CloseAndRecv() (*ObjectPutResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ObjectPutResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peerClient) ObjectGet(ctx context.Context, in *ObjectGetRequest, opts ...grpc.CallOption) (Peer_ObjectGetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Peer_serviceDesc.Streams[1], c.cc, "/bazil.peer.Peer/ObjectGet", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerObjectGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Peer_ObjectGetClient interface {
	Recv() (*ObjectGetResponse, error)
	grpc.ClientStream
}

type peerObjectGetClient struct {
	grpc.ClientStream
}

func (x *peerObjectGetClient) Recv() (*ObjectGetResponse, error) {
	m := new(ObjectGetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peerClient) VolumeConnect(ctx context.Context, in *VolumeConnectRequest, opts ...grpc.CallOption) (*VolumeConnectResponse, error) {
	out := new(VolumeConnectResponse)
	err := grpc.Invoke(ctx, "/bazil.peer.Peer/VolumeConnect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) VolumeSyncPull(ctx context.Context, in *VolumeSyncPullRequest, opts ...grpc.CallOption) (Peer_VolumeSyncPullClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Peer_serviceDesc.Streams[2], c.cc, "/bazil.peer.Peer/VolumeSyncPull", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerVolumeSyncPullClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Peer_VolumeSyncPullClient interface {
	Recv() (*VolumeSyncPullItem, error)
	grpc.ClientStream
}

type peerVolumeSyncPullClient struct {
	grpc.ClientStream
}

func (x *peerVolumeSyncPullClient) Recv() (*VolumeSyncPullItem, error) {
	m := new(VolumeSyncPullItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Peer service

type PeerServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	ObjectPut(Peer_ObjectPutServer) error
	ObjectGet(*ObjectGetRequest, Peer_ObjectGetServer) error
	VolumeConnect(context.Context, *VolumeConnectRequest) (*VolumeConnectResponse, error)
	VolumeSyncPull(*VolumeSyncPullRequest, Peer_VolumeSyncPullServer) error
}

func RegisterPeerServer(s *grpc.Server, srv PeerServer) {
	s.RegisterService(&_Peer_serviceDesc, srv)
}

func _Peer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PeerServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Peer_ObjectPut_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeerServer).ObjectPut(&peerObjectPutServer{stream})
}

type Peer_ObjectPutServer interface {
	SendAndClose(*ObjectPutResponse) error
	Recv() (*ObjectPutRequest, error)
	grpc.ServerStream
}

type peerObjectPutServer struct {
	grpc.ServerStream
}

func (x *peerObjectPutServer) SendAndClose(m *ObjectPutResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *peerObjectPutServer) Recv() (*ObjectPutRequest, error) {
	m := new(ObjectPutRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Peer_ObjectGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ObjectGetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PeerServer).ObjectGet(m, &peerObjectGetServer{stream})
}

type Peer_ObjectGetServer interface {
	Send(*ObjectGetResponse) error
	grpc.ServerStream
}

type peerObjectGetServer struct {
	grpc.ServerStream
}

func (x *peerObjectGetServer) Send(m *ObjectGetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Peer_VolumeConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(VolumeConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PeerServer).VolumeConnect(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Peer_VolumeSyncPull_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VolumeSyncPullRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PeerServer).VolumeSyncPull(m, &peerVolumeSyncPullServer{stream})
}

type Peer_VolumeSyncPullServer interface {
	Send(*VolumeSyncPullItem) error
	grpc.ServerStream
}

type peerVolumeSyncPullServer struct {
	grpc.ServerStream
}

func (x *peerVolumeSyncPullServer) Send(m *VolumeSyncPullItem) error {
	return x.ServerStream.SendMsg(m)
}

var _Peer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bazil.peer.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Peer_Ping_Handler,
		},
		{
			MethodName: "VolumeConnect",
			Handler:    _Peer_VolumeConnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ObjectPut",
			Handler:       _Peer_ObjectPut_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ObjectGet",
			Handler:       _Peer_ObjectGet_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "VolumeSyncPull",
			Handler:       _Peer_VolumeSyncPull_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x94, 0x4f, 0x73, 0xd2, 0x40,
	0x18, 0xc6, 0x49, 0x93, 0x60, 0x79, 0x53, 0x28, 0x2e, 0x76, 0xcc, 0x30, 0x5a, 0x21, 0x83, 0x23,
	0xed, 0x68, 0x70, 0xe8, 0xc1, 0x8e, 0x9e, 0x2c, 0xa0, 0x32, 0xa3, 0x85, 0x01, 0xac, 0xa3, 0x97,
	0x4e, 0x08, 0xdb, 0x36, 0x36, 0x24, 0x98, 0x2c, 0x75, 0xf0, 0x53, 0x78, 0xf0, 0xe6, 0x97, 0x75,
	0xff, 0x24, 0x10, 0xfe, 0xb4, 0x9e, 0x92, 0xec, 0xfe, 0xde, 0x67, 0x9f, 0xdd, 0x7d, 0x9f, 0x40,
	0x65, 0x68, 0xfd, 0x72, 0x5c, 0xd3, 0x0f, 0x2e, 0x6b, 0xfc, 0xad, 0x36, 0xc1, 0x38, 0xa8, 0xfd,
	0x74, 0x02, 0xcc, 0xdf, 0xcc, 0x49, 0xe0, 0x13, 0x1f, 0x81, 0xa0, 0xd8, 0x48, 0xf1, 0xd9, 0x6a,
	0x85, 0x6d, 0x85, 0xa2, 0x60, 0x6c, 0x79, 0xce, 0x05, 0x0e, 0x89, 0x28, 0x32, 0xb2, 0xa0, 0x75,
	0x1d, 0xef, 0xb2, 0x87, 0x7f, 0x4c, 0xe9, 0xa0, 0x91, 0x83, 0x1d, 0xf1, 0x19, 0x4e, 0x7c, 0x2f,
	0xc4, 0xc6, 0x0b, 0xc8, 0x77, 0x86, 0xdf, 0xb1, 0x4d, 0xba, 0x53, 0x12, 0x31, 0x48, 0x03, 0xf9,
	0x1a, 0xcf, 0x74, 0xa9, 0x24, 0x55, 0x77, 0xd0, 0x0e, 0x28, 0x23, 0x8b, 0x58, 0xfa, 0x16, 0xfb,
	0x32, 0x0a, 0x70, 0x3f, 0x81, 0x47, 0x1a, 0x4f, 0x62, 0x8d, 0xf7, 0x78, 0xa3, 0x86, 0x51, 0x8e,
	0xab, 0x38, 0x20, 0xaa, 0xe6, 0xc2, 0x02, 0x39, 0x84, 0x07, 0x67, 0xbe, 0x3b, 0x1d, 0xe3, 0x86,
	0xef, 0x79, 0x94, 0x8c, 0x75, 0x10, 0xc0, 0x0d, 0x1f, 0x3f, 0xb5, 0xc6, 0x98, 0xb3, 0x19, 0xe3,
	0x00, 0xf6, 0x56, 0xd8, 0x48, 0x32, 0x0f, 0xdb, 0x02, 0x6e, 0x37, 0x23, 0xd9, 0x57, 0x31, 0xda,
	0x9f, 0x79, 0x76, 0x77, 0xea, 0xba, 0xb1, 0xee, 0x1a, 0xca, 0xfc, 0x4c, 0x2c, 0x72, 0xc5, 0x37,
	0x9a, 0x31, 0xfe, 0x6c, 0x01, 0x5a, 0xae, 0x6c, 0x13, 0x3c, 0x46, 0x47, 0xa0, 0xe2, 0x20, 0xf0,
	0x03, 0x5e, 0x93, 0xab, 0x57, 0xcc, 0xc5, 0x95, 0x98, 0xeb, 0xb8, 0xd9, 0x62, 0x2c, 0x3a, 0x06,
	0x95, 0x01, 0x21, 0x95, 0x96, 0xab, 0x5a, 0xfd, 0xe0, 0x3f, 0x45, 0x5d, 0xc6, 0xb6, 0x3c, 0x12,
	0xcc, 0x98, 0xcb, 0x91, 0x13, 0x34, 0x5c, 0xdf, 0xbe, 0xd6, 0x15, 0xee, 0xb2, 0x02, 0xdb, 0xf6,
	0x95, 0xe3, 0x8e, 0x02, 0xec, 0xe9, 0x32, 0x97, 0x43, 0x49, 0xb9, 0x26, 0xed, 0x00, 0x8f, 0x14,
	0x9f, 0x03, 0x24, 0x54, 0x12, 0x77, 0x91, 0x45, 0x59, 0x50, 0x6f, 0x2c, 0x77, 0x8a, 0xc5, 0x85,
	0xbe, 0xde, 0x3a, 0x96, 0xe8, 0x79, 0xaa, 0xc2, 0xa8, 0x06, 0xf7, 0xfa, 0x9f, 0x1b, 0x8d, 0x56,
	0xbf, 0x9f, 0x4f, 0xa1, 0x02, 0xec, 0x9e, 0x76, 0x06, 0xe7, 0x6f, 0xcf, 0x9b, 0xed, 0x5e, 0xab,
	0x31, 0xe8, 0xf4, 0xbe, 0xe6, 0x25, 0xe3, 0xaf, 0x04, 0x69, 0xb1, 0x06, 0x3b, 0x2f, 0x6f, 0x7e,
	0x27, 0xa8, 0x04, 0xca, 0x85, 0xe3, 0x0a, 0x55, 0xad, 0x9e, 0x4f, 0x7a, 0x7a, 0x47, 0xc7, 0x3f,
	0xa4, 0xd0, 0x3e, 0xc8, 0x74, 0x2f, 0xd4, 0x34, 0x03, 0x76, 0x57, 0x4c, 0xd3, 0xf9, 0x43, 0xc8,
	0x10, 0x7f, 0x3c, 0x0c, 0x89, 0xef, 0x61, 0x5d, 0xe5, 0xd4, 0x5e, 0x92, 0x1a, 0xc4, 0x93, 0x94,
	0xa5, 0x9b, 0xb0, 0x17, 0x87, 0x72, 0x92, 0x06, 0x85, 0xcc, 0x26, 0xac, 0x99, 0x15, 0xb6, 0x18,
	0x7a, 0x0a, 0xdb, 0x71, 0x0a, 0xb8, 0x3d, 0xad, 0x5e, 0x88, 0x94, 0x68, 0x4a, 0xcc, 0x4f, 0xd1,
	0x94, 0xa1, 0x82, 0x4c, 0x97, 0x36, 0x34, 0xc8, 0xcc, 0xb5, 0xeb, 0xbf, 0x65, 0x50, 0xd8, 0xd1,
	0xa1, 0x37, 0xf4, 0x49, 0x83, 0x82, 0x1e, 0x26, 0x3d, 0x24, 0x92, 0x54, 0xd4, 0xd7, 0x27, 0xa2,
	0x3c, 0xa4, 0xd0, 0x47, 0xc8, 0xcc, 0x63, 0x82, 0x1e, 0x25, 0xc1, 0xd5, 0xb0, 0x15, 0x1f, 0xdf,
	0x32, 0x1b, 0x6b, 0x55, 0xa5, 0x85, 0x1a, 0x8d, 0xcf, 0x26, 0xb5, 0x45, 0xec, 0x36, 0xa9, 0x25,
	0x32, 0x67, 0xa4, 0x5e, 0x4a, 0xe8, 0x0c, 0xb2, 0x4b, 0xe9, 0x41, 0xa5, 0xf5, 0x7e, 0x5c, 0x0e,
	0x61, 0xb1, 0x7c, 0x07, 0x31, 0xdf, 0xf3, 0x17, 0xc8, 0x2d, 0x37, 0x33, 0x2a, 0xdf, 0xde, 0xe8,
	0xb1, 0xf2, 0xfe, 0xdd, 0x59, 0x60, 0x86, 0x4f, 0xd2, 0xdf, 0x14, 0xf6, 0x63, 0x1b, 0xa6, 0xf9,
	0x0f, 0xed, 0xe8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xa5, 0xeb, 0x1b, 0x2d, 0x05, 0x00,
	0x00,
}
