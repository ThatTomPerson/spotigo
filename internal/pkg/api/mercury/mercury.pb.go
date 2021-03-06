// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mercury.proto

/*
Package mercury is a generated protocol buffer package.

It is generated from these files:
	mercury.proto

It has these top-level messages:
	MercuryMultiGetRequest
	MercuryMultiGetReply
	MercuryRequest
	MercuryReply
	Header
	UserField
*/
package mercury

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

type MercuryReply_CachePolicy int32

const (
	MercuryReply_CACHE_NO      MercuryReply_CachePolicy = 1
	MercuryReply_CACHE_PRIVATE MercuryReply_CachePolicy = 2
	MercuryReply_CACHE_PUBLIC  MercuryReply_CachePolicy = 3
)

var MercuryReply_CachePolicy_name = map[int32]string{
	1: "CACHE_NO",
	2: "CACHE_PRIVATE",
	3: "CACHE_PUBLIC",
}
var MercuryReply_CachePolicy_value = map[string]int32{
	"CACHE_NO":      1,
	"CACHE_PRIVATE": 2,
	"CACHE_PUBLIC":  3,
}

func (x MercuryReply_CachePolicy) Enum() *MercuryReply_CachePolicy {
	p := new(MercuryReply_CachePolicy)
	*p = x
	return p
}
func (x MercuryReply_CachePolicy) String() string {
	return proto.EnumName(MercuryReply_CachePolicy_name, int32(x))
}
func (x *MercuryReply_CachePolicy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MercuryReply_CachePolicy_value, data, "MercuryReply_CachePolicy")
	if err != nil {
		return err
	}
	*x = MercuryReply_CachePolicy(value)
	return nil
}
func (MercuryReply_CachePolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type MercuryMultiGetRequest struct {
	Request          []*MercuryRequest `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *MercuryMultiGetRequest) Reset()                    { *m = MercuryMultiGetRequest{} }
func (m *MercuryMultiGetRequest) String() string            { return proto.CompactTextString(m) }
func (*MercuryMultiGetRequest) ProtoMessage()               {}
func (*MercuryMultiGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MercuryMultiGetRequest) GetRequest() []*MercuryRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type MercuryMultiGetReply struct {
	Reply            []*MercuryReply `protobuf:"bytes,1,rep,name=reply" json:"reply,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *MercuryMultiGetReply) Reset()                    { *m = MercuryMultiGetReply{} }
func (m *MercuryMultiGetReply) String() string            { return proto.CompactTextString(m) }
func (*MercuryMultiGetReply) ProtoMessage()               {}
func (*MercuryMultiGetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MercuryMultiGetReply) GetReply() []*MercuryReply {
	if m != nil {
		return m.Reply
	}
	return nil
}

type MercuryRequest struct {
	Uri              *string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	ContentType      *string `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Body             []byte  `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Etag             []byte  `protobuf:"bytes,4,opt,name=etag" json:"etag,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MercuryRequest) Reset()                    { *m = MercuryRequest{} }
func (m *MercuryRequest) String() string            { return proto.CompactTextString(m) }
func (*MercuryRequest) ProtoMessage()               {}
func (*MercuryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MercuryRequest) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *MercuryRequest) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *MercuryRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *MercuryRequest) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

type MercuryReply struct {
	StatusCode       *int32                    `protobuf:"zigzag32,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	StatusMessage    *string                   `protobuf:"bytes,2,opt,name=status_message,json=statusMessage" json:"status_message,omitempty"`
	CachePolicy      *MercuryReply_CachePolicy `protobuf:"varint,3,opt,name=cache_policy,json=cachePolicy,enum=MercuryReply_CachePolicy" json:"cache_policy,omitempty"`
	Ttl              *int32                    `protobuf:"zigzag32,4,opt,name=ttl" json:"ttl,omitempty"`
	Etag             []byte                    `protobuf:"bytes,5,opt,name=etag" json:"etag,omitempty"`
	ContentType      *string                   `protobuf:"bytes,6,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Body             []byte                    `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *MercuryReply) Reset()                    { *m = MercuryReply{} }
func (m *MercuryReply) String() string            { return proto.CompactTextString(m) }
func (*MercuryReply) ProtoMessage()               {}
func (*MercuryReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MercuryReply) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *MercuryReply) GetStatusMessage() string {
	if m != nil && m.StatusMessage != nil {
		return *m.StatusMessage
	}
	return ""
}

func (m *MercuryReply) GetCachePolicy() MercuryReply_CachePolicy {
	if m != nil && m.CachePolicy != nil {
		return *m.CachePolicy
	}
	return MercuryReply_CACHE_NO
}

func (m *MercuryReply) GetTtl() int32 {
	if m != nil && m.Ttl != nil {
		return *m.Ttl
	}
	return 0
}

func (m *MercuryReply) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

func (m *MercuryReply) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *MercuryReply) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Header struct {
	Uri              *string      `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	ContentType      *string      `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Method           *string      `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	StatusCode       *int32       `protobuf:"zigzag32,4,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	UserFields       []*UserField `protobuf:"bytes,6,rep,name=user_fields,json=userFields" json:"user_fields,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Header) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Header) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *Header) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *Header) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *Header) GetUserFields() []*UserField {
	if m != nil {
		return m.UserFields
	}
	return nil
}

type UserField struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            []byte  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserField) Reset()                    { *m = UserField{} }
func (m *UserField) String() string            { return proto.CompactTextString(m) }
func (*UserField) ProtoMessage()               {}
func (*UserField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserField) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *UserField) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*MercuryMultiGetRequest)(nil), "MercuryMultiGetRequest")
	proto.RegisterType((*MercuryMultiGetReply)(nil), "MercuryMultiGetReply")
	proto.RegisterType((*MercuryRequest)(nil), "MercuryRequest")
	proto.RegisterType((*MercuryReply)(nil), "MercuryReply")
	proto.RegisterType((*Header)(nil), "Header")
	proto.RegisterType((*UserField)(nil), "UserField")
	proto.RegisterEnum("MercuryReply_CachePolicy", MercuryReply_CachePolicy_name, MercuryReply_CachePolicy_value)
}

func init() { proto.RegisterFile("mercury.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdf, 0x6e, 0x94, 0x40,
	0x14, 0xc6, 0xc3, 0xfe, 0xa1, 0xee, 0x01, 0x56, 0x76, 0xd2, 0x34, 0x78, 0xe5, 0x8a, 0x31, 0x59,
	0x63, 0xc2, 0x45, 0xbd, 0xd4, 0x0b, 0x2b, 0xa9, 0xb6, 0x89, 0xab, 0xcd, 0xa4, 0xf5, 0x96, 0x20,
	0x1c, 0x5b, 0x52, 0xe8, 0xe0, 0xfc, 0x31, 0x99, 0xe7, 0xf1, 0x11, 0x7c, 0x41, 0x33, 0x0c, 0x94,
	0xae, 0x1b, 0x6f, 0xbc, 0xfb, 0xce, 0xef, 0x9c, 0x9c, 0xf3, 0xcd, 0x07, 0x10, 0x34, 0xc8, 0x0b,
	0xc5, 0x75, 0xd2, 0x72, 0x26, 0x59, 0x9c, 0xc2, 0xd1, 0xd6, 0x82, 0xad, 0xaa, 0x65, 0xf5, 0x11,
	0x25, 0xc5, 0x1f, 0x0a, 0x85, 0x24, 0x2f, 0xe1, 0x80, 0x5b, 0x19, 0x39, 0xeb, 0xe9, 0xc6, 0x3b,
	0x7e, 0x9c, 0xf4, 0x93, 0xfd, 0x04, 0x1d, 0xfa, 0xf1, 0x1b, 0x38, 0xdc, 0x5b, 0xd2, 0xd6, 0x9a,
	0x3c, 0x87, 0x39, 0x37, 0xa2, 0x5f, 0x10, 0x8c, 0x0b, 0xda, 0x5a, 0x53, 0xdb, 0x8b, 0x1b, 0x58,
	0xee, 0xee, 0x25, 0x21, 0x4c, 0x15, 0xaf, 0x22, 0x67, 0xed, 0x6c, 0x16, 0xd4, 0x48, 0xf2, 0x0c,
	0xfc, 0x82, 0xdd, 0x49, 0xbc, 0x93, 0x99, 0xd4, 0x2d, 0x46, 0x93, 0xae, 0xe5, 0xf5, 0xec, 0x52,
	0xb7, 0x48, 0x08, 0xcc, 0xbe, 0xb1, 0x52, 0x47, 0xd3, 0xb5, 0xb3, 0xf1, 0x69, 0xa7, 0x0d, 0x43,
	0x99, 0x5f, 0x47, 0x33, 0xcb, 0x8c, 0x8e, 0x7f, 0x4f, 0xc0, 0x7f, 0x68, 0x83, 0x3c, 0x05, 0x4f,
	0xc8, 0x5c, 0x2a, 0x91, 0x15, 0xac, 0xc4, 0xee, 0xea, 0x8a, 0x82, 0x45, 0x29, 0x2b, 0x91, 0xbc,
	0x80, 0x65, 0x3f, 0xd0, 0xa0, 0x10, 0xf9, 0xf5, 0x70, 0x3e, 0xb0, 0x74, 0x6b, 0x21, 0x79, 0x0b,
	0x7e, 0x91, 0x17, 0x37, 0x98, 0xb5, 0xac, 0xae, 0x0a, 0x6b, 0x64, 0x79, 0xfc, 0x64, 0xe7, 0xcd,
	0x49, 0x6a, 0x26, 0x2e, 0xba, 0x01, 0xea, 0x15, 0x63, 0x61, 0xde, 0x2c, 0x65, 0xdd, 0x39, 0x5d,
	0x51, 0x23, 0xef, 0xcd, 0xcf, 0x47, 0xf3, 0x7b, 0x39, 0xb8, 0xff, 0xce, 0xe1, 0x60, 0xcc, 0x21,
	0x7e, 0x07, 0xde, 0x83, 0xc3, 0xc4, 0x87, 0x47, 0xe9, 0x49, 0x7a, 0x76, 0x9a, 0x7d, 0xfe, 0x12,
	0x3a, 0x64, 0x05, 0x81, 0xad, 0x2e, 0xe8, 0xf9, 0xd7, 0x93, 0xcb, 0xd3, 0x70, 0x42, 0x42, 0xf0,
	0x7b, 0x74, 0xf5, 0xfe, 0xd3, 0x79, 0x1a, 0x4e, 0xe3, 0x5f, 0x0e, 0xb8, 0x67, 0x98, 0x97, 0xc8,
	0xff, 0xef, 0xeb, 0x1c, 0x81, 0xdb, 0xa0, 0xbc, 0x61, 0x65, 0x17, 0xcb, 0x82, 0xf6, 0xd5, 0xdf,
	0xe1, 0xcf, 0xf6, 0xc2, 0x7f, 0x05, 0x9e, 0x12, 0xc8, 0xb3, 0xef, 0x15, 0xd6, 0xa5, 0x88, 0xdc,
	0xee, 0x47, 0x82, 0xe4, 0x4a, 0x20, 0xff, 0x60, 0x10, 0x05, 0x35, 0x48, 0x11, 0xbf, 0x86, 0xc5,
	0x7d, 0xc3, 0xf8, 0xbc, 0x45, 0x3d, 0xf8, 0xbc, 0x45, 0x4d, 0x0e, 0x61, 0xfe, 0x33, 0xaf, 0x95,
	0x35, 0xe8, 0x53, 0x5b, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xc0, 0xa9, 0x39, 0x11, 0x03,
	0x00, 0x00,
}
