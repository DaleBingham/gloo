// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: endpoint.proto

package v1 // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=ep
// @solo-kit:resource.plural_name=endpoints
// @solo-kit:resource.resource_groups=api.gloo.solo.io
//
// Endpoints represent dynamically discovered address/ports where an upstream service is listening
type Endpoint struct {
	// List of the upstreams the endpoint belongs to
	Upstreams []*core.ResourceRef `protobuf:"bytes,1,rep,name=upstreams" json:"upstreams,omitempty"`
	// Address of the endpoint (ip or hostname)
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// listening port for the endpoint
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoint_4e122dd0536d8c7e, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (dst *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(dst, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetUpstreams() []*core.ResourceRef {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Endpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "gloo.solo.io.Endpoint")
}
func (this *Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint)
	if !ok {
		that2, ok := that.(Endpoint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Upstreams) != len(that1.Upstreams) {
		return false
	}
	for i := range this.Upstreams {
		if !this.Upstreams[i].Equal(that1.Upstreams[i]) {
			return false
		}
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("endpoint.proto", fileDescriptor_endpoint_4e122dd0536d8c7e) }

var fileDescriptor_endpoint_4e122dd0536d8c7e = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x31, 0xad, 0x68, 0xeb, 0x02, 0x83, 0x85, 0x90, 0xe9, 0x00, 0x11, 0x53, 0x06, 0xb0,
	0x69, 0x91, 0x80, 0xb9, 0x12, 0x23, 0x8b, 0x47, 0x36, 0x37, 0xb9, 0x35, 0xa6, 0x3f, 0xd7, 0xb2,
	0x1d, 0x9e, 0x09, 0x89, 0x17, 0xe1, 0x29, 0x18, 0x78, 0x12, 0x94, 0xc4, 0x01, 0x21, 0x31, 0x30,
	0xf9, 0x5e, 0x9d, 0xef, 0xc8, 0xe7, 0x5c, 0x7a, 0x08, 0xdb, 0xd2, 0xa1, 0xdd, 0x46, 0xe1, 0x3c,
	0x46, 0x64, 0xfb, 0x66, 0x8d, 0x28, 0x02, 0xae, 0x51, 0x58, 0x9c, 0x1c, 0x19, 0x34, 0xd8, 0x08,
	0xb2, 0x9e, 0x5a, 0x66, 0x32, 0x35, 0x36, 0x3e, 0x55, 0x0b, 0x51, 0xe0, 0x46, 0xd6, 0xe4, 0xa5,
	0xc5, 0xf6, 0x5d, 0xd9, 0x28, 0xb5, 0xb3, 0xf2, 0x65, 0x2a, 0x37, 0x10, 0x75, 0xa9, 0xa3, 0x4e,
	0x96, 0x8b, 0x7f, 0x58, 0x3c, 0x2c, 0x5b, 0xfa, 0xfc, 0x8d, 0xd0, 0xe1, 0x7d, 0xca, 0xc5, 0x6e,
	0xe9, 0xa8, 0x72, 0x21, 0x7a, 0xd0, 0x9b, 0xc0, 0x49, 0xd6, 0xcb, 0xc7, 0xb3, 0x13, 0x51, 0xa0,
	0x87, 0x2e, 0xa5, 0x50, 0x10, 0xb0, 0xf2, 0x05, 0x28, 0x58, 0xaa, 0x1f, 0x96, 0x71, 0x3a, 0xd0,
	0x65, 0xe9, 0x21, 0x04, 0xbe, 0x9b, 0x91, 0x7c, 0xa4, 0xba, 0x95, 0x31, 0xda, 0x77, 0xe8, 0x23,
	0xef, 0x65, 0x24, 0x3f, 0x50, 0xcd, 0xcc, 0xee, 0xe8, 0xb0, 0xcb, 0xcc, 0x07, 0x19, 0xc9, 0xc7,
	0xb3, 0xe3, 0xdf, 0xbf, 0x3c, 0x24, 0x75, 0xde, 0x7f, 0xff, 0x38, 0xdb, 0x51, 0xdf, 0xf4, 0xfc,
	0xe6, 0xf5, 0xf3, 0x94, 0x3c, 0x5e, 0xfd, 0xd1, 0xb0, 0xbe, 0xa5, 0x74, 0x1e, 0x9f, 0xa1, 0x88,
	0x21, 0x6d, 0x2b, 0x93, 0xfa, 0x2e, 0xf6, 0x9a, 0xb2, 0xd7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0xd3, 0x54, 0xed, 0x83, 0x01, 0x00, 0x00,
}
