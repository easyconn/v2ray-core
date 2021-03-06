package http

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Config for HTTP proxy server.
type ServerConfig struct {
	Timeout          uint32            `protobuf:"varint,1,opt,name=timeout" json:"timeout,omitempty"`
	Accounts         map[string]string `protobuf:"bytes,2,rep,name=accounts" json:"accounts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AllowTransparent bool              `protobuf:"varint,3,opt,name=allow_transparent,json=allowTransparent" json:"allow_transparent,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServerConfig) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ServerConfig) GetAccounts() map[string]string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ServerConfig) GetAllowTransparent() bool {
	if m != nil {
		return m.AllowTransparent
	}
	return false
}

// ClientConfig for HTTP proxy client.
type ClientConfig struct {
}

func (m *ClientConfig) Reset()                    { *m = ClientConfig{} }
func (m *ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()               {}
func (*ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ServerConfig)(nil), "v2ray.core.proxy.http.ServerConfig")
	proto.RegisterType((*ClientConfig)(nil), "v2ray.core.proxy.http.ClientConfig")
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/http/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0x86, 0x49, 0xe7, 0xc7, 0x8c, 0x9b, 0xcc, 0xe0, 0xa0, 0x7a, 0x55, 0x76, 0x21, 0x05, 0x21,
	0xc5, 0x7a, 0x23, 0xee, 0x4a, 0x8b, 0xe0, 0x8d, 0x30, 0xaa, 0x78, 0xe1, 0x8d, 0xc4, 0x10, 0xb5,
	0xd8, 0xe6, 0x94, 0xb3, 0xd3, 0x6a, 0xfe, 0x92, 0xff, 0xca, 0x7f, 0x22, 0xcd, 0x9c, 0x1f, 0xb0,
	0xab, 0xe4, 0x9c, 0xf7, 0xc9, 0xc3, 0x4b, 0xf8, 0x61, 0x9b, 0xa2, 0x72, 0x52, 0x43, 0x95, 0x68,
	0x40, 0x93, 0xd4, 0x08, 0xef, 0x2e, 0x79, 0x21, 0xaa, 0x13, 0x0d, 0xf6, 0xa9, 0x78, 0x96, 0x35,
	0x02, 0x81, 0x18, 0x2f, 0x39, 0x34, 0xd2, 0x33, 0xb2, 0x63, 0x26, 0x9f, 0x8c, 0x0f, 0x6e, 0x0c,
	0xb6, 0x06, 0x33, 0x4f, 0x8b, 0x90, 0x6f, 0x52, 0x51, 0x19, 0x68, 0x28, 0x64, 0x11, 0x8b, 0x87,
	0xf9, 0x72, 0x14, 0xd7, 0xbc, 0xaf, 0xb4, 0x86, 0xc6, 0xd2, 0x3c, 0x0c, 0xa2, 0x5e, 0xbc, 0x9d,
	0x1e, 0xcb, 0x95, 0x52, 0xf9, 0x57, 0x28, 0xcf, 0xbf, 0xdf, 0x5c, 0x5a, 0x42, 0x97, 0xff, 0x28,
	0xc4, 0x11, 0xdf, 0x55, 0x65, 0x09, 0x6f, 0x0f, 0x84, 0xca, 0xce, 0x6b, 0x85, 0xc6, 0x52, 0xd8,
	0x8b, 0x58, 0xdc, 0xcf, 0x47, 0x3e, 0xb8, 0xfd, 0xdd, 0x1f, 0x4c, 0xf9, 0xf0, 0x9f, 0x47, 0x8c,
	0x78, 0xef, 0xd5, 0x38, 0x5f, 0x71, 0x2b, 0xef, 0xae, 0x62, 0x8f, 0xaf, 0xb7, 0xaa, 0x6c, 0x4c,
	0x18, 0xf8, 0xdd, 0x62, 0x38, 0x0b, 0x4e, 0xd9, 0x64, 0x87, 0x0f, 0xb2, 0xb2, 0x30, 0x96, 0x16,
	0x8d, 0x2e, 0xa6, 0x7c, 0x5f, 0x43, 0xb5, 0xba, 0xfb, 0x8c, 0xdd, 0xaf, 0x75, 0xe7, 0x47, 0x30,
	0xbe, 0x4b, 0x73, 0xe5, 0x64, 0xd6, 0xe5, 0x33, 0x9f, 0x5f, 0x11, 0xd5, 0x8f, 0x1b, 0xfe, 0x3b,
	0x4f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x08, 0xbc, 0x0a, 0x78, 0x01, 0x00, 0x00,
}
