// Code generated by protoc-gen-go.
// source: client/pkg/shard/shard.proto
// DO NOT EDIT!

/*
Package shard is a generated protocol buffer package.

It is generated from these files:
	client/pkg/shard/shard.proto

It has these top-level messages:
	ServerState
	FrontendState
	ServerRole
	Addresses
	StartRegister
	FinishRegister
	Version
	StartAssignRoles
	FinishAssignRoles
	FailedToAssignRoles
	SetServerState
	SetFrontendState
	AddServerRole
	RemoveServerRole
	SetServerRole
	DeleteServerRole
	SetAddresses
	GetAddress
	GetShardToAddress
*/
package shard

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ServerState struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Version int64  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *ServerState) Reset()                    { *m = ServerState{} }
func (m *ServerState) String() string            { return proto.CompactTextString(m) }
func (*ServerState) ProtoMessage()               {}
func (*ServerState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FrontendState struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Version int64  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *FrontendState) Reset()                    { *m = FrontendState{} }
func (m *FrontendState) String() string            { return proto.CompactTextString(m) }
func (*FrontendState) ProtoMessage()               {}
func (*FrontendState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ServerRole struct {
	Address string          `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Version int64           `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	Shards  map[uint64]bool `protobuf:"bytes,3,rep,name=shards" json:"shards,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ServerRole) Reset()                    { *m = ServerRole{} }
func (m *ServerRole) String() string            { return proto.CompactTextString(m) }
func (*ServerRole) ProtoMessage()               {}
func (*ServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServerRole) GetShards() map[uint64]bool {
	if m != nil {
		return m.Shards
	}
	return nil
}

type Addresses struct {
	Version   int64             `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Addresses map[uint64]string `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Addresses) Reset()                    { *m = Addresses{} }
func (m *Addresses) String() string            { return proto.CompactTextString(m) }
func (*Addresses) ProtoMessage()               {}
func (*Addresses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Addresses) GetAddresses() map[uint64]string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type StartRegister struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *StartRegister) Reset()                    { *m = StartRegister{} }
func (m *StartRegister) String() string            { return proto.CompactTextString(m) }
func (*StartRegister) ProtoMessage()               {}
func (*StartRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type FinishRegister struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FinishRegister) Reset()                    { *m = FinishRegister{} }
func (m *FinishRegister) String() string            { return proto.CompactTextString(m) }
func (*FinishRegister) ProtoMessage()               {}
func (*FinishRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Version struct {
	Result int64  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type StartAssignRoles struct {
}

func (m *StartAssignRoles) Reset()                    { *m = StartAssignRoles{} }
func (m *StartAssignRoles) String() string            { return proto.CompactTextString(m) }
func (*StartAssignRoles) ProtoMessage()               {}
func (*StartAssignRoles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type FinishAssignRoles struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *FinishAssignRoles) Reset()                    { *m = FinishAssignRoles{} }
func (m *FinishAssignRoles) String() string            { return proto.CompactTextString(m) }
func (*FinishAssignRoles) ProtoMessage()               {}
func (*FinishAssignRoles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type FailedToAssignRoles struct {
	ServerStates map[string]*ServerState `protobuf:"bytes,1,rep,name=server_states" json:"server_states,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NumShards    uint64                  `protobuf:"varint,2,opt,name=num_shards" json:"num_shards,omitempty"`
	NumReplicas  uint64                  `protobuf:"varint,3,opt,name=num_replicas" json:"num_replicas,omitempty"`
}

func (m *FailedToAssignRoles) Reset()                    { *m = FailedToAssignRoles{} }
func (m *FailedToAssignRoles) String() string            { return proto.CompactTextString(m) }
func (*FailedToAssignRoles) ProtoMessage()               {}
func (*FailedToAssignRoles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FailedToAssignRoles) GetServerStates() map[string]*ServerState {
	if m != nil {
		return m.ServerStates
	}
	return nil
}

type SetServerState struct {
	ServerState *ServerState `protobuf:"bytes,1,opt,name=serverState" json:"serverState,omitempty"`
}

func (m *SetServerState) Reset()                    { *m = SetServerState{} }
func (m *SetServerState) String() string            { return proto.CompactTextString(m) }
func (*SetServerState) ProtoMessage()               {}
func (*SetServerState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SetServerState) GetServerState() *ServerState {
	if m != nil {
		return m.ServerState
	}
	return nil
}

type SetFrontendState struct {
	FrontendState *FrontendState `protobuf:"bytes,1,opt,name=frontendState" json:"frontendState,omitempty"`
}

func (m *SetFrontendState) Reset()                    { *m = SetFrontendState{} }
func (m *SetFrontendState) String() string            { return proto.CompactTextString(m) }
func (*SetFrontendState) ProtoMessage()               {}
func (*SetFrontendState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SetFrontendState) GetFrontendState() *FrontendState {
	if m != nil {
		return m.FrontendState
	}
	return nil
}

type AddServerRole struct {
	ServerRole *ServerRole `protobuf:"bytes,1,opt,name=serverRole" json:"serverRole,omitempty"`
	Error      string      `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *AddServerRole) Reset()                    { *m = AddServerRole{} }
func (m *AddServerRole) String() string            { return proto.CompactTextString(m) }
func (*AddServerRole) ProtoMessage()               {}
func (*AddServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AddServerRole) GetServerRole() *ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type RemoveServerRole struct {
	ServerRole *ServerRole `protobuf:"bytes,1,opt,name=serverRole" json:"serverRole,omitempty"`
	Error      string      `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *RemoveServerRole) Reset()                    { *m = RemoveServerRole{} }
func (m *RemoveServerRole) String() string            { return proto.CompactTextString(m) }
func (*RemoveServerRole) ProtoMessage()               {}
func (*RemoveServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *RemoveServerRole) GetServerRole() *ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type SetServerRole struct {
	ServerRole *ServerRole `protobuf:"bytes,2,opt,name=serverRole" json:"serverRole,omitempty"`
}

func (m *SetServerRole) Reset()                    { *m = SetServerRole{} }
func (m *SetServerRole) String() string            { return proto.CompactTextString(m) }
func (*SetServerRole) ProtoMessage()               {}
func (*SetServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *SetServerRole) GetServerRole() *ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type DeleteServerRole struct {
	ServerRole *ServerRole `protobuf:"bytes,2,opt,name=serverRole" json:"serverRole,omitempty"`
}

func (m *DeleteServerRole) Reset()                    { *m = DeleteServerRole{} }
func (m *DeleteServerRole) String() string            { return proto.CompactTextString(m) }
func (*DeleteServerRole) ProtoMessage()               {}
func (*DeleteServerRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *DeleteServerRole) GetServerRole() *ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type SetAddresses struct {
	Addresses *Addresses `protobuf:"bytes,1,opt,name=addresses" json:"addresses,omitempty"`
}

func (m *SetAddresses) Reset()                    { *m = SetAddresses{} }
func (m *SetAddresses) String() string            { return proto.CompactTextString(m) }
func (*SetAddresses) ProtoMessage()               {}
func (*SetAddresses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *SetAddresses) GetAddresses() *Addresses {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type GetAddress struct {
	Shard   uint64 `protobuf:"varint,1,opt,name=shard" json:"shard,omitempty"`
	Version int64  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	Result  string `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Ok      bool   `protobuf:"varint,4,opt,name=ok" json:"ok,omitempty"`
	Error   string `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
}

func (m *GetAddress) Reset()                    { *m = GetAddress{} }
func (m *GetAddress) String() string            { return proto.CompactTextString(m) }
func (*GetAddress) ProtoMessage()               {}
func (*GetAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type GetShardToAddress struct {
	Version int64             `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Result  map[uint64]string `protobuf:"bytes,2,rep,name=result" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error   string            `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *GetShardToAddress) Reset()                    { *m = GetShardToAddress{} }
func (m *GetShardToAddress) String() string            { return proto.CompactTextString(m) }
func (*GetShardToAddress) ProtoMessage()               {}
func (*GetShardToAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *GetShardToAddress) GetResult() map[uint64]string {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerState)(nil), "shard.ServerState")
	proto.RegisterType((*FrontendState)(nil), "shard.FrontendState")
	proto.RegisterType((*ServerRole)(nil), "shard.ServerRole")
	proto.RegisterType((*Addresses)(nil), "shard.Addresses")
	proto.RegisterType((*StartRegister)(nil), "shard.StartRegister")
	proto.RegisterType((*FinishRegister)(nil), "shard.FinishRegister")
	proto.RegisterType((*Version)(nil), "shard.Version")
	proto.RegisterType((*StartAssignRoles)(nil), "shard.StartAssignRoles")
	proto.RegisterType((*FinishAssignRoles)(nil), "shard.FinishAssignRoles")
	proto.RegisterType((*FailedToAssignRoles)(nil), "shard.FailedToAssignRoles")
	proto.RegisterType((*SetServerState)(nil), "shard.SetServerState")
	proto.RegisterType((*SetFrontendState)(nil), "shard.SetFrontendState")
	proto.RegisterType((*AddServerRole)(nil), "shard.AddServerRole")
	proto.RegisterType((*RemoveServerRole)(nil), "shard.RemoveServerRole")
	proto.RegisterType((*SetServerRole)(nil), "shard.SetServerRole")
	proto.RegisterType((*DeleteServerRole)(nil), "shard.DeleteServerRole")
	proto.RegisterType((*SetAddresses)(nil), "shard.SetAddresses")
	proto.RegisterType((*GetAddress)(nil), "shard.GetAddress")
	proto.RegisterType((*GetShardToAddress)(nil), "shard.GetShardToAddress")
}

var fileDescriptor0 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x51, 0x6f, 0x12, 0x41,
	0x10, 0xce, 0x41, 0xa1, 0x32, 0xf4, 0x10, 0xce, 0x3e, 0x90, 0x46, 0x23, 0xae, 0x1a, 0x49, 0xb4,
	0xa0, 0xd4, 0x18, 0xf5, 0xc5, 0x34, 0x5a, 0xea, 0x33, 0x18, 0x7d, 0x6c, 0xce, 0x32, 0xd2, 0x0b,
	0xd7, 0x3b, 0xb2, 0xbb, 0x90, 0xf4, 0xd5, 0x67, 0x7f, 0x80, 0x7f, 0xca, 0xff, 0xe4, 0xee, 0xec,
	0x72, 0xb7, 0x70, 0x54, 0x6b, 0x7c, 0x21, 0xdc, 0xec, 0x7c, 0xdf, 0x7c, 0xfb, 0xed, 0xcc, 0xc0,
	0xdd, 0xf3, 0x38, 0xc2, 0x44, 0xf6, 0xe7, 0xb3, 0x69, 0x5f, 0x5c, 0x84, 0x7c, 0x62, 0x7e, 0x7b,
	0x73, 0x9e, 0xca, 0x34, 0xa8, 0xd0, 0x07, 0xeb, 0x43, 0x7d, 0x8c, 0x7c, 0x89, 0x7c, 0x2c, 0x43,
	0x89, 0xc1, 0x6d, 0xd8, 0x0d, 0x27, 0x13, 0x8e, 0x42, 0xb4, 0xbd, 0x8e, 0xd7, 0xad, 0xe9, 0x80,
	0x3a, 0x14, 0x51, 0x9a, 0xb4, 0x4b, 0x2a, 0x50, 0x66, 0x2f, 0xc0, 0x1f, 0xf2, 0x34, 0x91, 0x98,
	0x4c, 0x6e, 0x0a, 0xf9, 0xe1, 0x01, 0x98, 0x22, 0xa3, 0x34, 0xbe, 0x01, 0x20, 0x38, 0x84, 0x2a,
	0xa9, 0x13, 0xed, 0x72, 0xa7, 0xdc, 0xad, 0x0f, 0xee, 0xf5, 0x8c, 0xf2, 0x9c, 0xa4, 0x37, 0xa6,
	0xf3, 0x93, 0x44, 0xf2, 0xab, 0x83, 0x43, 0x75, 0x87, 0xfc, 0x33, 0xa8, 0x43, 0x79, 0x86, 0x57,
	0xc4, 0xbd, 0x13, 0xf8, 0x50, 0x59, 0x86, 0xf1, 0x02, 0x89, 0xf9, 0xd6, 0xdb, 0xd2, 0x6b, 0x8f,
	0x7d, 0xf7, 0xa0, 0x76, 0x6c, 0x04, 0xa0, 0x70, 0x8b, 0x7b, 0x54, 0x7c, 0x00, 0xb5, 0x70, 0x75,
	0xaa, 0x50, 0xba, 0xfe, 0x7d, 0x5b, 0x3f, 0x43, 0xe5, 0xff, 0x8c, 0x82, 0xe7, 0xd0, 0x58, 0x8f,
	0xfc, 0x41, 0x44, 0x8d, 0x44, 0x74, 0xc0, 0x57, 0xf6, 0x71, 0x39, 0xc2, 0x69, 0x24, 0x24, 0xf2,
	0x82, 0x2b, 0x4c, 0x71, 0x0e, 0xa3, 0x24, 0x12, 0x17, 0xd7, 0xa6, 0x68, 0x5e, 0xe4, 0x3c, 0xe5,
	0x86, 0x97, 0x75, 0x61, 0xf7, 0xb3, 0xb9, 0x4a, 0xd0, 0x80, 0xaa, 0xca, 0x5b, 0xc4, 0xd2, 0x5e,
	0x6a, 0x23, 0x33, 0x80, 0x26, 0x55, 0x3f, 0x16, 0x22, 0x9a, 0x26, 0xda, 0x50, 0xc1, 0x18, 0xb4,
	0x4c, 0x3d, 0x27, 0x98, 0xe3, 0x8c, 0xa6, 0x5f, 0x1e, 0xdc, 0x19, 0x86, 0x51, 0x8c, 0x93, 0x4f,
	0xa9, 0x9b, 0xf6, 0x1e, 0x7c, 0x41, 0x6f, 0x73, 0x26, 0x74, 0x4f, 0x68, 0x7d, 0xda, 0xb7, 0x67,
	0xd6, 0xb7, 0x2d, 0x90, 0x9e, 0xd3, 0x75, 0xd6, 0xb2, 0x00, 0x20, 0x59, 0x5c, 0x9e, 0xd9, 0x97,
	0x2f, 0x91, 0x73, 0xfb, 0xb0, 0xa7, 0x63, 0x1c, 0xe7, 0x71, 0x74, 0x1e, 0xea, 0x7e, 0x50, 0xd1,
	0x83, 0x53, 0x68, 0x15, 0xe1, 0x8e, 0xe3, 0xb5, 0xe0, 0x81, 0xeb, 0x78, 0x7d, 0x10, 0xac, 0x35,
	0x10, 0xa1, 0xe8, 0x15, 0xde, 0x40, 0x63, 0x8c, 0xd2, 0x1d, 0x80, 0x27, 0x50, 0x17, 0xf9, 0x27,
	0xb1, 0x6d, 0x85, 0xb3, 0x77, 0xca, 0x42, 0x94, 0xeb, 0xa3, 0xf0, 0x14, 0xfc, 0x6f, 0x6e, 0xc0,
	0xc2, 0xf7, 0x57, 0x36, 0xb8, 0x67, 0xec, 0x04, 0x7c, 0xd5, 0x33, 0xce, 0x5c, 0x3c, 0x06, 0x10,
	0xd9, 0x97, 0x85, 0xb6, 0x0a, 0x9d, 0xbf, 0xf9, 0x94, 0x1f, 0xa1, 0x39, 0xc2, 0xcb, 0x74, 0x89,
	0xff, 0xcd, 0xf4, 0x4a, 0xb5, 0xe4, 0xca, 0x8c, 0x2d, 0x34, 0xa5, 0x6b, 0x68, 0x94, 0x89, 0xcd,
	0x0f, 0x18, 0xa3, 0xc4, 0x7f, 0x87, 0x1e, 0xc1, 0x9e, 0x2a, 0x99, 0x0f, 0xe3, 0x43, 0x77, 0xf6,
	0x8c, 0xee, 0xe6, 0xe6, 0xec, 0xb1, 0x2f, 0x00, 0xa7, 0x19, 0x48, 0x5f, 0x82, 0x12, 0xec, 0xa8,
	0x15, 0x76, 0x49, 0x3e, 0x09, 0x65, 0xea, 0x0c, 0x80, 0x52, 0x3a, 0x6b, 0xef, 0xe8, 0x6d, 0x90,
	0x1b, 0x50, 0x21, 0x03, 0x7e, 0x7a, 0xd0, 0x52, 0xcc, 0xb4, 0x4b, 0x54, 0xb3, 0xda, 0x02, 0x85,
	0x05, 0xf1, 0x32, 0x63, 0x34, 0xdb, 0xe1, 0x91, 0x55, 0x58, 0x80, 0xf6, 0x46, 0x94, 0x66, 0xda,
	0x33, 0xab, 0x45, 0x32, 0xf4, 0xce, 0x72, 0x4f, 0xff, 0xb2, 0x2e, 0xbe, 0x56, 0x69, 0x69, 0x1f,
	0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x65, 0xa7, 0x5f, 0x72, 0xd4, 0x05, 0x00, 0x00,
}
