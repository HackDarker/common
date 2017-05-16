// Code generated by protoc-gen-go.
// source: mj_changchun_base.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DdzEnumProtoId int32

const (
	DdzEnumProtoId_P16ID_HEARTBEAT     DdzEnumProtoId = 0
	DdzEnumProtoId_P16ID_REQCREATEDESK DdzEnumProtoId = 1
	DdzEnumProtoId_P16ID_ACKCREATEDESK DdzEnumProtoId = 2
	DdzEnumProtoId_P16ID_REQREADY      DdzEnumProtoId = 3
	DdzEnumProtoId_P16ID_ACKREADY      DdzEnumProtoId = 4
)

var DdzEnumProtoId_name = map[int32]string{
	0: "P16ID_HEARTBEAT",
	1: "P16ID_REQCREATEDESK",
	2: "P16ID_ACKCREATEDESK",
	3: "P16ID_REQREADY",
	4: "P16ID_ACKREADY",
}
var DdzEnumProtoId_value = map[string]int32{
	"P16ID_HEARTBEAT":     0,
	"P16ID_REQCREATEDESK": 1,
	"P16ID_ACKCREATEDESK": 2,
	"P16ID_REQREADY":      3,
	"P16ID_ACKREADY":      4,
}

func (x DdzEnumProtoId) Enum() *DdzEnumProtoId {
	p := new(DdzEnumProtoId)
	*p = x
	return p
}
func (x DdzEnumProtoId) String() string {
	return proto.EnumName(DdzEnumProtoId_name, int32(x))
}
func (x *DdzEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DdzEnumProtoId_value, data, "DdzEnumProtoId")
	if err != nil {
		return err
	}
	*x = DdzEnumProtoId(value)
	return nil
}
func (DdzEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterEnum("yjprotogo.DdzEnumProtoId", DdzEnumProtoId_name, DdzEnumProtoId_value)
}

var fileDescriptor1 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0xcd, 0x8a, 0x4f,
	0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xce, 0x28, 0xcd, 0x8b, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xac, 0xcc, 0x02, 0x33, 0xd2, 0xf3, 0xb5, 0x1a, 0x19, 0xb9, 0x04,
	0x52, 0x52, 0xaa, 0xe2, 0x53, 0xf3, 0x4a, 0x73, 0xe3, 0xc1, 0x82, 0x9e, 0x29, 0x42, 0xc2, 0x5c,
	0xfc, 0x01, 0x86, 0x66, 0x9e, 0x2e, 0xf1, 0x1e, 0xae, 0x8e, 0x41, 0x21, 0x4e, 0xae, 0x8e, 0x21,
	0x02, 0x0c, 0x42, 0xe2, 0x5c, 0xc2, 0x10, 0xc1, 0x20, 0xd7, 0x40, 0xe7, 0x20, 0xa0, 0x98, 0xab,
	0x8b, 0x6b, 0xb0, 0xb7, 0x00, 0x23, 0x42, 0xc2, 0xd1, 0xd9, 0x1b, 0x49, 0x82, 0x49, 0x48, 0x88,
	0x8b, 0x0f, 0xae, 0x03, 0x28, 0xee, 0x12, 0x29, 0xc0, 0x8c, 0x10, 0x03, 0x2a, 0x86, 0x88, 0xb1,
	0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xdb, 0x51, 0xb5, 0xa8, 0x00, 0x00, 0x00,
}
