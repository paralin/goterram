// Code generated by protoc-gen-go.
// source: github.com/paralin/goterram/components/sprite/sprite.proto
// DO NOT EDIT!

/*
Package sprite is a generated protocol buffer package.

It is generated from these files:
	github.com/paralin/goterram/components/sprite/sprite.proto

It has these top-level messages:
	SpriteData
	SpriteComponentInit
*/
package sprite

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

type SpriteData struct {
}

func (m *SpriteData) Reset()                    { *m = SpriteData{} }
func (m *SpriteData) String() string            { return proto.CompactTextString(m) }
func (*SpriteData) ProtoMessage()               {}
func (*SpriteData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SpriteComponentInit struct {
	Data *SpriteData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *SpriteComponentInit) Reset()                    { *m = SpriteComponentInit{} }
func (m *SpriteComponentInit) String() string            { return proto.CompactTextString(m) }
func (*SpriteComponentInit) ProtoMessage()               {}
func (*SpriteComponentInit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SpriteComponentInit) GetData() *SpriteData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SpriteData)(nil), "sprite.SpriteData")
	proto.RegisterType((*SpriteComponentInit)(nil), "sprite.SpriteComponentInit")
}

func init() {
	proto.RegisterFile("github.com/paralin/goterram/components/sprite/sprite.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2b, 0x2d, 0x4e, 0x2d, 0xca, 0x4f, 0xca, 0x2f,
	0xc9, 0x4c, 0x2e, 0xd6, 0x4f, 0xcf, 0x2f, 0x49, 0x2d, 0x2a, 0x4a, 0xcc, 0xd5, 0x07, 0xca, 0x14,
	0xe4, 0xe7, 0xa5, 0xe6, 0x95, 0x14, 0xeb, 0x17, 0x17, 0x14, 0x65, 0x96, 0xa4, 0x42, 0x29, 0xbd,
	0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x36, 0x08, 0x4f, 0x89, 0x87, 0x8b, 0x2b, 0x18, 0xcc, 0x72,
	0x49, 0x2c, 0x49, 0x54, 0xb2, 0xe5, 0x12, 0x86, 0xf0, 0x9c, 0x61, 0xba, 0x3d, 0xf3, 0x32, 0x4b,
	0x84, 0xd4, 0xb8, 0x58, 0x52, 0x80, 0xd2, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x42, 0x7a,
	0x50, 0x93, 0x10, 0x1a, 0x83, 0xc0, 0xf2, 0x49, 0x6c, 0x60, 0xb3, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf8, 0xbc, 0x75, 0xd6, 0x9e, 0x00, 0x00, 0x00,
}
