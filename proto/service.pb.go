// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Book struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Book) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

type BookServiceReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookServiceReq) Reset()         { *m = BookServiceReq{} }
func (m *BookServiceReq) String() string { return proto.CompactTextString(m) }
func (*BookServiceReq) ProtoMessage()    {}
func (*BookServiceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *BookServiceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookServiceReq.Unmarshal(m, b)
}
func (m *BookServiceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookServiceReq.Marshal(b, m, deterministic)
}
func (m *BookServiceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookServiceReq.Merge(m, src)
}
func (m *BookServiceReq) XXX_Size() int {
	return xxx_messageInfo_BookServiceReq.Size(m)
}
func (m *BookServiceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BookServiceReq.DiscardUnknown(m)
}

var xxx_messageInfo_BookServiceReq proto.InternalMessageInfo

type BookServiceReq_Create struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookServiceReq_Create) Reset()         { *m = BookServiceReq_Create{} }
func (m *BookServiceReq_Create) String() string { return proto.CompactTextString(m) }
func (*BookServiceReq_Create) ProtoMessage()    {}
func (*BookServiceReq_Create) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1, 0}
}

func (m *BookServiceReq_Create) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookServiceReq_Create.Unmarshal(m, b)
}
func (m *BookServiceReq_Create) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookServiceReq_Create.Marshal(b, m, deterministic)
}
func (m *BookServiceReq_Create) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookServiceReq_Create.Merge(m, src)
}
func (m *BookServiceReq_Create) XXX_Size() int {
	return xxx_messageInfo_BookServiceReq_Create.Size(m)
}
func (m *BookServiceReq_Create) XXX_DiscardUnknown() {
	xxx_messageInfo_BookServiceReq_Create.DiscardUnknown(m)
}

var xxx_messageInfo_BookServiceReq_Create proto.InternalMessageInfo

func (m *BookServiceReq_Create) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BookServiceReq_Create) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

type BookServiceReq_Find struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookServiceReq_Find) Reset()         { *m = BookServiceReq_Find{} }
func (m *BookServiceReq_Find) String() string { return proto.CompactTextString(m) }
func (*BookServiceReq_Find) ProtoMessage()    {}
func (*BookServiceReq_Find) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1, 1}
}

func (m *BookServiceReq_Find) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookServiceReq_Find.Unmarshal(m, b)
}
func (m *BookServiceReq_Find) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookServiceReq_Find.Marshal(b, m, deterministic)
}
func (m *BookServiceReq_Find) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookServiceReq_Find.Merge(m, src)
}
func (m *BookServiceReq_Find) XXX_Size() int {
	return xxx_messageInfo_BookServiceReq_Find.Size(m)
}
func (m *BookServiceReq_Find) XXX_DiscardUnknown() {
	xxx_messageInfo_BookServiceReq_Find.DiscardUnknown(m)
}

var xxx_messageInfo_BookServiceReq_Find proto.InternalMessageInfo

func (m *BookServiceReq_Find) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Book)(nil), "yiplee.twirp.gateway.service.Book")
	proto.RegisterType((*BookServiceReq)(nil), "yiplee.twirp.gateway.service.BookServiceReq")
	proto.RegisterType((*BookServiceReq_Create)(nil), "yiplee.twirp.gateway.service.BookServiceReq.Create")
	proto.RegisterType((*BookServiceReq_Find)(nil), "yiplee.twirp.gateway.service.BookServiceReq.Find")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xa9, 0xcc, 0x2c, 0xc8, 0x49,
	0x4d, 0xd5, 0x2b, 0x29, 0xcf, 0x2c, 0x2a, 0xd0, 0x4b, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0xd4,
	0x83, 0xaa, 0x51, 0x72, 0xe0, 0x62, 0x71, 0xca, 0xcf, 0xcf, 0x16, 0xe2, 0xe3, 0x62, 0xca, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0xc9, 0xf9, 0x65, 0xa9,
	0x45, 0x12, 0xcc, 0x60, 0x41, 0x08, 0x47, 0x29, 0x86, 0x8b, 0x0f, 0x64, 0x42, 0x30, 0xc4, 0xc0,
	0xa0, 0xd4, 0x42, 0x29, 0x23, 0x2e, 0x36, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0xb8, 0x29, 0x8c,
	0xd8, 0x4c, 0x61, 0x42, 0x32, 0x45, 0x4a, 0x8c, 0x8b, 0xc5, 0x2d, 0x33, 0x2f, 0x05, 0xdd, 0x1d,
	0x46, 0xe7, 0x19, 0xb9, 0xb8, 0x91, 0x8c, 0x17, 0x4a, 0x84, 0x9b, 0x6d, 0xac, 0x87, 0xcf, 0x63,
	0x7a, 0xa8, 0x6e, 0xd2, 0x83, 0x68, 0x92, 0x52, 0x22, 0xac, 0x49, 0x28, 0x16, 0xea, 0x14, 0x43,
	0x92, 0x2c, 0x00, 0x69, 0x21, 0xc6, 0x78, 0x27, 0xf6, 0x28, 0x56, 0x70, 0xc4, 0x24, 0xb1, 0x81,
	0x29, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xa2, 0x63, 0x5e, 0xb0, 0x01, 0x00, 0x00,
}
