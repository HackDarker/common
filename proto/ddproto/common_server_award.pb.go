// Code generated by protoc-gen-go.
// source: common_server_award.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// todo 等待增加
type Award int32

const (
	Award_taskType_time Award = 1
	Award_taskType_2    Award = 2
)

var Award_name = map[int32]string{
	1: "taskType_time",
	2: "taskType_2",
}
var Award_value = map[string]int32{
	"taskType_time": 1,
	"taskType_2":    2,
}

func (x Award) Enum() *Award {
	p := new(Award)
	*p = x
	return p
}
func (x Award) String() string {
	return proto.EnumName(Award_name, int32(x))
}
func (x *Award) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Award_value, data, "Award")
	if err != nil {
		return err
	}
	*x = Award(value)
	return nil
}
func (Award) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

// 用户奖励，保存到数据库的结构
type Taward struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TaskType         *int32  `protobuf:"varint,2,opt,name=taskType" json:"taskType,omitempty"`
	AwardType        *int32  `protobuf:"varint,3,opt,name=awardType" json:"awardType,omitempty"`
	UserId           *uint32 `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	Memo             *string `protobuf:"bytes,5,opt,name=memo" json:"memo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Taward) Reset()                    { *m = Taward{} }
func (m *Taward) String() string            { return proto.CompactTextString(m) }
func (*Taward) ProtoMessage()               {}
func (*Taward) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Taward) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Taward) GetTaskType() int32 {
	if m != nil && m.TaskType != nil {
		return *m.TaskType
	}
	return 0
}

func (m *Taward) GetAwardType() int32 {
	if m != nil && m.AwardType != nil {
		return *m.AwardType
	}
	return 0
}

func (m *Taward) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Taward) GetMemo() string {
	if m != nil && m.Memo != nil {
		return *m.Memo
	}
	return ""
}

// 在线奖励
type AwardMOnline struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Capital          *int64  `protobuf:"varint,2,opt,name=capital" json:"capital,omitempty"`
	Rate             *int64  `protobuf:"varint,3,opt,name=rate" json:"rate,omitempty"`
	BeginTime        *string `protobuf:"bytes,4,opt,name=beginTime" json:"beginTime,omitempty"`
	DurationSec      *int64  `protobuf:"varint,5,opt,name=durationSec" json:"durationSec,omitempty"`
	Level            *int32  `protobuf:"varint,6,opt,name=level" json:"level,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AwardMOnline) Reset()                    { *m = AwardMOnline{} }
func (m *AwardMOnline) String() string            { return proto.CompactTextString(m) }
func (*AwardMOnline) ProtoMessage()               {}
func (*AwardMOnline) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *AwardMOnline) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *AwardMOnline) GetCapital() int64 {
	if m != nil && m.Capital != nil {
		return *m.Capital
	}
	return 0
}

func (m *AwardMOnline) GetRate() int64 {
	if m != nil && m.Rate != nil {
		return *m.Rate
	}
	return 0
}

func (m *AwardMOnline) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *AwardMOnline) GetDurationSec() int64 {
	if m != nil && m.DurationSec != nil {
		return *m.DurationSec
	}
	return 0
}

func (m *AwardMOnline) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*Taward)(nil), "ddproto.Taward")
	proto.RegisterType((*AwardMOnline)(nil), "ddproto.award_m_online")
	proto.RegisterEnum("ddproto.Award", Award_name, Award_value)
}

var fileDescriptor5 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0x49, 0xbb, 0xed, 0xda, 0x91, 0x2d, 0xeb, 0x20, 0x12, 0xc5, 0x43, 0xd9, 0x53, 0xd9,
	0x83, 0x07, 0xdf, 0xc2, 0x6b, 0xec, 0xbd, 0xc4, 0x66, 0x90, 0x60, 0xd3, 0x94, 0x34, 0xbb, 0xe2,
	0xc5, 0x27, 0xf2, 0x21, 0x65, 0xa7, 0xee, 0x87, 0xb7, 0xff, 0x07, 0xff, 0xf0, 0xcb, 0xc0, 0x7d,
	0xe7, 0x9d, 0xf3, 0x43, 0x3b, 0x51, 0xd8, 0x53, 0x68, 0xf5, 0xa7, 0x0e, 0xe6, 0x69, 0x0c, 0x3e,
	0x7a, 0x5c, 0x1a, 0xc3, 0x62, 0xf3, 0x0d, 0x79, 0xc3, 0x05, 0x96, 0x90, 0x58, 0x23, 0x45, 0x25,
	0xea, 0x4c, 0x25, 0xd6, 0xe0, 0x03, 0x5c, 0x45, 0x3d, 0x7d, 0x34, 0x5f, 0x23, 0xc9, 0x84, 0xd3,
	0x93, 0xc7, 0x47, 0x28, 0x78, 0xc4, 0x65, 0xca, 0xe5, 0x39, 0xc0, 0x3b, 0xc8, 0x77, 0x13, 0x85,
	0x17, 0x23, 0x17, 0x95, 0xa8, 0x57, 0xea, 0xcf, 0x21, 0xc2, 0xc2, 0x91, 0xf3, 0x32, 0xab, 0x44,
	0x5d, 0x28, 0xd6, 0x9b, 0x1f, 0x01, 0x25, 0x2f, 0x5b, 0xd7, 0xfa, 0xa1, 0xb7, 0xc3, 0xe5, 0x5c,
	0xfc, 0x9b, 0x4b, 0x58, 0x76, 0x7a, 0xb4, 0x51, 0xf7, 0xcc, 0x93, 0xaa, 0xa3, 0x3d, 0x3c, 0x1c,
	0x74, 0x9c, 0x49, 0x52, 0xc5, 0xfa, 0x80, 0xf8, 0x46, 0xef, 0x76, 0x68, 0xac, 0x23, 0xe6, 0x28,
	0xd4, 0x39, 0xc0, 0x0a, 0xae, 0xcd, 0x2e, 0xe8, 0x68, 0xfd, 0xf0, 0x4a, 0x1d, 0x13, 0xa5, 0xea,
	0x32, 0xc2, 0x5b, 0xc8, 0x7a, 0xda, 0x53, 0x2f, 0x73, 0xfe, 0xde, 0x6c, 0xb6, 0x5b, 0xc8, 0xe6,
	0x6b, 0xdd, 0xc0, 0xea, 0x78, 0x8d, 0x36, 0x5a, 0x47, 0x6b, 0x81, 0x25, 0xc0, 0x29, 0x7a, 0x5e,
	0x27, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x93, 0x74, 0x68, 0x0b, 0x7f, 0x01, 0x00, 0x00,
}
