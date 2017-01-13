// Code generated by protoc-gen-go.
// source: bazil.org/bazil/server/control/wire/peer.proto
// DO NOT EDIT!

package wire

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PeerAddRequest struct {
	// Must be exactly 32 bytes long.
	Pub []byte `protobuf:"bytes,2,opt,name=pub,proto3" json:"pub,omitempty"`
}

func (m *PeerAddRequest) Reset()                    { *m = PeerAddRequest{} }
func (m *PeerAddRequest) String() string            { return proto.CompactTextString(m) }
func (*PeerAddRequest) ProtoMessage()               {}
func (*PeerAddRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PeerAddResponse struct {
}

func (m *PeerAddResponse) Reset()                    { *m = PeerAddResponse{} }
func (m *PeerAddResponse) String() string            { return proto.CompactTextString(m) }
func (*PeerAddResponse) ProtoMessage()               {}
func (*PeerAddResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type PeerLocationSetRequest struct {
	// Must be exactly 32 bytes long.
	Pub    []byte `protobuf:"bytes,1,opt,name=pub,proto3" json:"pub,omitempty"`
	Netloc string `protobuf:"bytes,2,opt,name=netloc" json:"netloc,omitempty"`
}

func (m *PeerLocationSetRequest) Reset()                    { *m = PeerLocationSetRequest{} }
func (m *PeerLocationSetRequest) String() string            { return proto.CompactTextString(m) }
func (*PeerLocationSetRequest) ProtoMessage()               {}
func (*PeerLocationSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type PeerLocationSetResponse struct {
}

func (m *PeerLocationSetResponse) Reset()                    { *m = PeerLocationSetResponse{} }
func (m *PeerLocationSetResponse) String() string            { return proto.CompactTextString(m) }
func (*PeerLocationSetResponse) ProtoMessage()               {}
func (*PeerLocationSetResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type PeerStorageAllowRequest struct {
	// Must be exactly 32 bytes long.
	Pub     []byte `protobuf:"bytes,1,opt,name=pub,proto3" json:"pub,omitempty"`
	Backend string `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
}

func (m *PeerStorageAllowRequest) Reset()                    { *m = PeerStorageAllowRequest{} }
func (m *PeerStorageAllowRequest) String() string            { return proto.CompactTextString(m) }
func (*PeerStorageAllowRequest) ProtoMessage()               {}
func (*PeerStorageAllowRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type PeerStorageAllowResponse struct {
}

func (m *PeerStorageAllowResponse) Reset()                    { *m = PeerStorageAllowResponse{} }
func (m *PeerStorageAllowResponse) String() string            { return proto.CompactTextString(m) }
func (*PeerStorageAllowResponse) ProtoMessage()               {}
func (*PeerStorageAllowResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type PeerVolumeAllowRequest struct {
	// Must be exactly 32 bytes long.
	Pub        []byte `protobuf:"bytes,1,opt,name=pub,proto3" json:"pub,omitempty"`
	VolumeName string `protobuf:"bytes,2,opt,name=volumeName" json:"volumeName,omitempty"`
}

func (m *PeerVolumeAllowRequest) Reset()                    { *m = PeerVolumeAllowRequest{} }
func (m *PeerVolumeAllowRequest) String() string            { return proto.CompactTextString(m) }
func (*PeerVolumeAllowRequest) ProtoMessage()               {}
func (*PeerVolumeAllowRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type PeerVolumeAllowResponse struct {
}

func (m *PeerVolumeAllowResponse) Reset()                    { *m = PeerVolumeAllowResponse{} }
func (m *PeerVolumeAllowResponse) String() string            { return proto.CompactTextString(m) }
func (*PeerVolumeAllowResponse) ProtoMessage()               {}
func (*PeerVolumeAllowResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*PeerAddRequest)(nil), "bazil.control.PeerAddRequest")
	proto.RegisterType((*PeerAddResponse)(nil), "bazil.control.PeerAddResponse")
	proto.RegisterType((*PeerLocationSetRequest)(nil), "bazil.control.PeerLocationSetRequest")
	proto.RegisterType((*PeerLocationSetResponse)(nil), "bazil.control.PeerLocationSetResponse")
	proto.RegisterType((*PeerStorageAllowRequest)(nil), "bazil.control.PeerStorageAllowRequest")
	proto.RegisterType((*PeerStorageAllowResponse)(nil), "bazil.control.PeerStorageAllowResponse")
	proto.RegisterType((*PeerVolumeAllowRequest)(nil), "bazil.control.PeerVolumeAllowRequest")
	proto.RegisterType((*PeerVolumeAllowResponse)(nil), "bazil.control.PeerVolumeAllowResponse")
}

var fileDescriptor1 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x40, 0xa9, 0xca, 0x8a, 0xa3, 0xb6, 0xb8, 0x07, 0xad, 0x82, 0x20, 0x39, 0x79, 0xda, 0x1c,
	0x44, 0xc4, 0x63, 0x3d, 0x8b, 0x88, 0x05, 0x0f, 0xde, 0x92, 0xec, 0x50, 0x16, 0xd3, 0x4c, 0x9c,
	0x64, 0x5b, 0xf0, 0xeb, 0x75, 0x93, 0x20, 0xd4, 0x8a, 0xb7, 0x90, 0x37, 0xef, 0x31, 0x0c, 0x34,
	0x5a, 0x7d, 0x76, 0xb6, 0x21, 0x5e, 0xc8, 0xf4, 0x92, 0x01, 0x79, 0x85, 0x2c, 0x0d, 0xb9, 0xc8,
	0x64, 0xe5, 0xba, 0x63, 0x94, 0x1e, 0x91, 0x1b, 0xcf, 0x14, 0xa9, 0x3e, 0xce, 0xf3, 0x05, 0x8b,
	0x4b, 0x18, 0x3f, 0x7f, 0xc3, 0x59, 0xdb, 0xbe, 0xe0, 0x47, 0x8f, 0x21, 0xd6, 0x87, 0xb0, 0xeb,
	0x7b, 0x3d, 0xdd, 0xb9, 0x1a, 0x5d, 0x1f, 0x89, 0x13, 0x98, 0xfc, 0xe0, 0xe0, 0xc9, 0x05, 0x14,
	0xb7, 0x70, 0x3a, 0x7c, 0x3d, 0x92, 0x51, 0xb1, 0x23, 0x37, 0xc7, 0xf8, 0xcb, 0x1c, 0x0d, 0x66,
	0x3d, 0x86, 0xca, 0x61, 0xb4, 0x64, 0x52, 0xe9, 0x40, 0x9c, 0xc3, 0xd9, 0x96, 0x56, 0x8a, 0x77,
	0x19, 0xcd, 0x23, 0xb1, 0x5a, 0xe0, 0xcc, 0x5a, 0x5a, 0xff, 0x99, 0x9c, 0xc0, 0xbe, 0x56, 0xe6,
	0x1d, 0x5d, 0x5b, 0x9a, 0x17, 0x30, 0xdd, 0x16, 0x4b, 0xf4, 0x3e, 0xaf, 0xf9, 0x4a, 0xb6, 0x5f,
	0xfe, 0xd3, 0xac, 0x01, 0x56, 0x69, 0xe4, 0x49, 0x2d, 0x71, 0x73, 0xd5, 0x0d, 0x35, 0x57, 0x1f,
	0xaa, 0xb7, 0xbd, 0xe1, 0xa0, 0xba, 0x4a, 0xc7, 0xbc, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc1,
	0xd6, 0x9b, 0x90, 0x7e, 0x01, 0x00, 0x00,
}
