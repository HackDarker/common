// Code generated by protoc-gen-go.
// source: common_server.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 用户的seesion 主要是游戏的
type GameSession struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	RoomId           *int32  `protobuf:"varint,2,opt,name=RoomId" json:"RoomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,3,opt,name=DeskId" json:"DeskId,omitempty"`
	GameStatus       *int32  `protobuf:"varint,4,opt,name=GameStatus" json:"GameStatus,omitempty"`
	GameId           *int32  `protobuf:"varint,5,opt,name=GameId" json:"GameId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,6,opt,name=GameNumber" json:"GameNumber,omitempty"`
	GameCustomStatus *int32  `protobuf:"varint,7,opt,name=GameCustomStatus" json:"GameCustomStatus,omitempty"`
	IsBreak          *bool   `protobuf:"varint,8,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool   `protobuf:"varint,9,opt,name=isLeave" json:"isLeave,omitempty"`
	RoomType         *int32  `protobuf:"varint,10,opt,name=roomType" json:"roomType,omitempty"`
	RoomPassword     *string `protobuf:"bytes,11,opt,name=roomPassword" json:"roomPassword,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,12,opt,name=RoomLevel" json:"RoomLevel,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameSession) Reset()                    { *m = GameSession{} }
func (m *GameSession) String() string            { return proto.CompactTextString(m) }
func (*GameSession) ProtoMessage()               {}
func (*GameSession) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *GameSession) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *GameSession) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *GameSession) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *GameSession) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *GameSession) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameSession) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *GameSession) GetGameCustomStatus() int32 {
	if m != nil && m.GameCustomStatus != nil {
		return *m.GameCustomStatus
	}
	return 0
}

func (m *GameSession) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *GameSession) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *GameSession) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *GameSession) GetRoomPassword() string {
	if m != nil && m.RoomPassword != nil {
		return *m.RoomPassword
	}
	return ""
}

func (m *GameSession) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func init() {
	proto.RegisterType((*GameSession)(nil), "yjprotogo.GameSession")
}

var fileDescriptor2 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xd1, 0x4a, 0xf3, 0x40,
	0x10, 0x85, 0x49, 0xff, 0xbf, 0x6d, 0x32, 0xad, 0x20, 0x2b, 0xc8, 0x20, 0x22, 0xa1, 0x57, 0xc1,
	0x0b, 0x1f, 0x42, 0x05, 0x09, 0x14, 0x91, 0x55, 0xaf, 0x65, 0x75, 0x07, 0xa9, 0x75, 0x3b, 0x65,
	0x27, 0x89, 0xf4, 0x99, 0x7d, 0x09, 0xd9, 0xdd, 0xc4, 0x2a, 0xde, 0xcd, 0xf7, 0x9d, 0x3d, 0xe7,
	0x62, 0xe1, 0xe8, 0x85, 0x9d, 0xe3, 0xcd, 0x93, 0x90, 0xef, 0xc8, 0x5f, 0x6c, 0x3d, 0x37, 0xac,
	0x8a, 0xdd, 0x5b, 0x3c, 0x5e, 0x79, 0xf1, 0x39, 0x82, 0xd9, 0x8d, 0x71, 0x74, 0x4f, 0x22, 0x2b,
	0xde, 0xa8, 0x63, 0x98, 0x3c, 0x0a, 0xf9, 0xda, 0x62, 0x56, 0x66, 0xd5, 0x81, 0xee, 0x29, 0x78,
	0xcd, 0xec, 0x6a, 0x8b, 0xa3, 0x32, 0xab, 0xc6, 0xba, 0xa7, 0xe0, 0xaf, 0x49, 0xd6, 0xb5, 0xc5,
	0x7f, 0xc9, 0x27, 0x52, 0x67, 0x00, 0x71, 0xb6, 0x31, 0x4d, 0x2b, 0xf8, 0x3f, 0x66, 0x3f, 0x4c,
	0xe8, 0x05, 0xaa, 0x2d, 0x8e, 0x53, 0x2f, 0xd1, 0xd0, 0xbb, 0x6d, 0xdd, 0x33, 0x79, 0x9c, 0xec,
	0x7b, 0xc9, 0xa8, 0x73, 0x38, 0x0c, 0x74, 0xd5, 0x4a, 0xc3, 0xae, 0x5f, 0x9f, 0xc6, 0x57, 0x7f,
	0xbc, 0x42, 0x98, 0xae, 0xe4, 0xd2, 0x93, 0x59, 0x63, 0x5e, 0x66, 0x55, 0xae, 0x07, 0x4c, 0xc9,
	0x92, 0x4c, 0x47, 0x58, 0x0c, 0x49, 0x44, 0x75, 0x02, 0xb9, 0x67, 0x76, 0x0f, 0xbb, 0x2d, 0x21,
	0xc4, 0xdd, 0x6f, 0x56, 0x0b, 0x98, 0x87, 0xfb, 0xce, 0x88, 0x7c, 0xb0, 0xb7, 0x38, 0x2b, 0xb3,
	0xaa, 0xd0, 0xbf, 0x9c, 0x3a, 0x85, 0x22, 0xfc, 0xcc, 0x92, 0x3a, 0x7a, 0xc7, 0x79, 0x1c, 0xd8,
	0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x1f, 0xda, 0xd0, 0x8e, 0x01, 0x00, 0x00,
}