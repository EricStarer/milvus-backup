// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index_cgo_msg.proto

package indexcgopb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-backup/internal/proto/commonpb"
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

type TypeParams struct {
	Params               []*commonpb.KeyValuePair `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TypeParams) Reset()         { *m = TypeParams{} }
func (m *TypeParams) String() string { return proto.CompactTextString(m) }
func (*TypeParams) ProtoMessage()    {}
func (*TypeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{0}
}

func (m *TypeParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeParams.Unmarshal(m, b)
}
func (m *TypeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeParams.Marshal(b, m, deterministic)
}
func (m *TypeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeParams.Merge(m, src)
}
func (m *TypeParams) XXX_Size() int {
	return xxx_messageInfo_TypeParams.Size(m)
}
func (m *TypeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeParams.DiscardUnknown(m)
}

var xxx_messageInfo_TypeParams proto.InternalMessageInfo

func (m *TypeParams) GetParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.Params
	}
	return nil
}

type IndexParams struct {
	Params               []*commonpb.KeyValuePair `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IndexParams) Reset()         { *m = IndexParams{} }
func (m *IndexParams) String() string { return proto.CompactTextString(m) }
func (*IndexParams) ProtoMessage()    {}
func (*IndexParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{1}
}

func (m *IndexParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexParams.Unmarshal(m, b)
}
func (m *IndexParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexParams.Marshal(b, m, deterministic)
}
func (m *IndexParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexParams.Merge(m, src)
}
func (m *IndexParams) XXX_Size() int {
	return xxx_messageInfo_IndexParams.Size(m)
}
func (m *IndexParams) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexParams.DiscardUnknown(m)
}

var xxx_messageInfo_IndexParams proto.InternalMessageInfo

func (m *IndexParams) GetParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.Params
	}
	return nil
}

// TypeParams & IndexParams will be replaced by MapParams later
type MapParams struct {
	Params               []*commonpb.KeyValuePair `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MapParams) Reset()         { *m = MapParams{} }
func (m *MapParams) String() string { return proto.CompactTextString(m) }
func (*MapParams) ProtoMessage()    {}
func (*MapParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{2}
}

func (m *MapParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapParams.Unmarshal(m, b)
}
func (m *MapParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapParams.Marshal(b, m, deterministic)
}
func (m *MapParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapParams.Merge(m, src)
}
func (m *MapParams) XXX_Size() int {
	return xxx_messageInfo_MapParams.Size(m)
}
func (m *MapParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MapParams.DiscardUnknown(m)
}

var xxx_messageInfo_MapParams proto.InternalMessageInfo

func (m *MapParams) GetParams() []*commonpb.KeyValuePair {
	if m != nil {
		return m.Params
	}
	return nil
}

type MapParamsV2 struct {
	Params               map[string]string `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MapParamsV2) Reset()         { *m = MapParamsV2{} }
func (m *MapParamsV2) String() string { return proto.CompactTextString(m) }
func (*MapParamsV2) ProtoMessage()    {}
func (*MapParamsV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{3}
}

func (m *MapParamsV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapParamsV2.Unmarshal(m, b)
}
func (m *MapParamsV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapParamsV2.Marshal(b, m, deterministic)
}
func (m *MapParamsV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapParamsV2.Merge(m, src)
}
func (m *MapParamsV2) XXX_Size() int {
	return xxx_messageInfo_MapParamsV2.Size(m)
}
func (m *MapParamsV2) XXX_DiscardUnknown() {
	xxx_messageInfo_MapParamsV2.DiscardUnknown(m)
}

var xxx_messageInfo_MapParamsV2 proto.InternalMessageInfo

func (m *MapParamsV2) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type Binary struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Binary) Reset()         { *m = Binary{} }
func (m *Binary) String() string { return proto.CompactTextString(m) }
func (*Binary) ProtoMessage()    {}
func (*Binary) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{4}
}

func (m *Binary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Binary.Unmarshal(m, b)
}
func (m *Binary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Binary.Marshal(b, m, deterministic)
}
func (m *Binary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Binary.Merge(m, src)
}
func (m *Binary) XXX_Size() int {
	return xxx_messageInfo_Binary.Size(m)
}
func (m *Binary) XXX_DiscardUnknown() {
	xxx_messageInfo_Binary.DiscardUnknown(m)
}

var xxx_messageInfo_Binary proto.InternalMessageInfo

func (m *Binary) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Binary) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type BinarySet struct {
	Datas                []*Binary `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BinarySet) Reset()         { *m = BinarySet{} }
func (m *BinarySet) String() string { return proto.CompactTextString(m) }
func (*BinarySet) ProtoMessage()    {}
func (*BinarySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_c009bd9544a7343c, []int{5}
}

func (m *BinarySet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BinarySet.Unmarshal(m, b)
}
func (m *BinarySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BinarySet.Marshal(b, m, deterministic)
}
func (m *BinarySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinarySet.Merge(m, src)
}
func (m *BinarySet) XXX_Size() int {
	return xxx_messageInfo_BinarySet.Size(m)
}
func (m *BinarySet) XXX_DiscardUnknown() {
	xxx_messageInfo_BinarySet.DiscardUnknown(m)
}

var xxx_messageInfo_BinarySet proto.InternalMessageInfo

func (m *BinarySet) GetDatas() []*Binary {
	if m != nil {
		return m.Datas
	}
	return nil
}

func init() {
	proto.RegisterType((*TypeParams)(nil), "milvus.proto.indexcgo.TypeParams")
	proto.RegisterType((*IndexParams)(nil), "milvus.proto.indexcgo.IndexParams")
	proto.RegisterType((*MapParams)(nil), "milvus.proto.indexcgo.MapParams")
	proto.RegisterType((*MapParamsV2)(nil), "milvus.proto.indexcgo.MapParamsV2")
	proto.RegisterMapType((map[string]string)(nil), "milvus.proto.indexcgo.MapParamsV2.ParamsEntry")
	proto.RegisterType((*Binary)(nil), "milvus.proto.indexcgo.Binary")
	proto.RegisterType((*BinarySet)(nil), "milvus.proto.indexcgo.BinarySet")
}

func init() { proto.RegisterFile("index_cgo_msg.proto", fileDescriptor_c009bd9544a7343c) }

var fileDescriptor_c009bd9544a7343c = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xc6, 0x0a, 0x7d, 0xdd, 0x41, 0xa2, 0x42, 0x19, 0x08, 0xb3, 0xa7, 0x5d, 0x4c,
	0x65, 0xbb, 0x38, 0x11, 0x91, 0x81, 0x53, 0x11, 0x61, 0x54, 0xd9, 0xc1, 0xcb, 0x48, 0xbb, 0xd0,
	0x85, 0xb5, 0x49, 0x68, 0xd3, 0x61, 0xf7, 0x29, 0xfc, 0xc8, 0xd2, 0xa6, 0x95, 0x29, 0xca, 0x0e,
	0xbb, 0xbd, 0xf7, 0xe7, 0xfd, 0x7e, 0xc9, 0x7b, 0x70, 0xcc, 0xf8, 0x92, 0x7e, 0x2c, 0xc2, 0x48,
	0x2c, 0x92, 0x2c, 0xc2, 0x32, 0x15, 0x4a, 0xa0, 0xd3, 0x84, 0xc5, 0x9b, 0x3c, 0xd3, 0x1d, 0xae,
	0x26, 0xc2, 0x48, 0xf4, 0xba, 0xa1, 0x48, 0x12, 0xc1, 0x75, 0xec, 0x3e, 0x00, 0xbc, 0x15, 0x92,
	0xce, 0x48, 0x4a, 0x92, 0x0c, 0x8d, 0xc1, 0x94, 0x55, 0xe5, 0x18, 0xfd, 0xf6, 0xc0, 0x1e, 0x9e,
	0xe3, 0x1f, 0x8e, 0x9a, 0x7c, 0xa6, 0xc5, 0x9c, 0xc4, 0x39, 0x9d, 0x11, 0x96, 0xfa, 0x35, 0xe0,
	0x3e, 0x82, 0xfd, 0x54, 0x3e, 0x71, 0xb8, 0x69, 0x0a, 0xd6, 0x0b, 0x91, 0x87, 0x7b, 0x3e, 0x0d,
	0xb0, 0xbf, 0x45, 0xf3, 0x21, 0x9a, 0xfe, 0x52, 0x61, 0xfc, 0xe7, 0x81, 0xf0, 0x0e, 0x83, 0x75,
	0x71, 0xcf, 0x55, 0x5a, 0x34, 0xde, 0xde, 0x18, 0xec, 0x9d, 0x18, 0x1d, 0x41, 0x7b, 0x4d, 0x0b,
	0xc7, 0xe8, 0x1b, 0x03, 0xcb, 0x2f, 0x4b, 0x74, 0x02, 0x9d, 0x4d, 0xf9, 0x1b, 0xa7, 0x55, 0x65,
	0xba, 0xb9, 0x6e, 0x5d, 0x19, 0xee, 0x25, 0x98, 0x13, 0xc6, 0xc9, 0x7e, 0xaa, 0x5b, 0x53, 0xee,
	0x1d, 0x58, 0x9a, 0x78, 0xa5, 0x0a, 0x8d, 0xa0, 0xb3, 0x24, 0x8a, 0x34, 0x0b, 0x9c, 0xfd, 0xb3,
	0x80, 0x06, 0x7c, 0x3d, 0x3b, 0xb9, 0x7d, 0xbf, 0x89, 0x98, 0x5a, 0xe5, 0x41, 0x79, 0xac, 0x2d,
	0x8b, 0x63, 0xb6, 0x55, 0x34, 0x5c, 0x79, 0x9a, 0xbd, 0x08, 0x48, 0xb8, 0xce, 0xa5, 0xc7, 0xb8,
	0xa2, 0x29, 0x27, 0xb1, 0x57, 0xb9, 0xbc, 0xc6, 0x25, 0x83, 0xc0, 0xac, 0x92, 0xd1, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa0, 0x39, 0x85, 0x0a, 0x64, 0x02, 0x00, 0x00,
}
