// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_server_msg.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMCARD_BILL_TYPE from common_enum.proto

// 游戏结束后的统计数据消息
type CommonSrvMsgGameCount struct {
	UserId           *uint32               `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	GameId           *CommonEnumGame       `protobuf:"varint,2,opt,name=game_id,json=gameId,enum=ddproto.CommonEnumGame" json:"game_id,omitempty"`
	GameNumber       *int32                `protobuf:"varint,3,opt,name=game_number,json=gameNumber" json:"game_number,omitempty"`
	RoomType         *COMMON_ENUM_ROOMTYPE `protobuf:"varint,4,opt,name=room_type,json=roomType,enum=ddproto.COMMON_ENUM_ROOMTYPE" json:"room_type,omitempty"`
	RoomLevel        *int32                `protobuf:"varint,5,opt,name=room_level,json=roomLevel" json:"room_level,omitempty"`
	Bill             *float64              `protobuf:"fixed64,6,opt,name=bill" json:"bill,omitempty"`
	IsWine           *bool                 `protobuf:"varint,7,opt,name=is_wine,json=isWine" json:"is_wine,omitempty"`
	Data             *string               `protobuf:"bytes,8,opt,name=data" json:"data,omitempty"`
	StartTime        *int64                `protobuf:"varint,9,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime          *int64                `protobuf:"varint,10,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	CoinFee          *int32                `protobuf:"varint,11,opt,name=coin_fee,json=coinFee" json:"coin_fee,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *CommonSrvMsgGameCount) Reset()                    { *m = CommonSrvMsgGameCount{} }
func (m *CommonSrvMsgGameCount) String() string            { return proto.CompactTextString(m) }
func (*CommonSrvMsgGameCount) ProtoMessage()               {}
func (*CommonSrvMsgGameCount) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *CommonSrvMsgGameCount) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetGameId() CommonEnumGame {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return CommonEnumGame_GID_SRC
}

func (m *CommonSrvMsgGameCount) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetRoomType() COMMON_ENUM_ROOMTYPE {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return COMMON_ENUM_ROOMTYPE_DESK_FRIEND
}

func (m *CommonSrvMsgGameCount) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetBill() float64 {
	if m != nil && m.Bill != nil {
		return *m.Bill
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetIsWine() bool {
	if m != nil && m.IsWine != nil {
		return *m.IsWine
	}
	return false
}

func (m *CommonSrvMsgGameCount) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *CommonSrvMsgGameCount) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *CommonSrvMsgGameCount) GetCoinFee() int32 {
	if m != nil && m.CoinFee != nil {
		return *m.CoinFee
	}
	return 0
}

func init() {
	proto.RegisterType((*CommonSrvMsgGameCount)(nil), "ddproto.common_srv_msg_game_count")
}

func init() { proto.RegisterFile("common_server_msg.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0xcf, 0x6b, 0xc2, 0x30,
	0x14, 0xc7, 0x17, 0x7f, 0xf4, 0x47, 0x64, 0x83, 0xe5, 0x62, 0x1c, 0xc8, 0xca, 0x4e, 0x3d, 0x79,
	0xf0, 0xb8, 0xeb, 0x70, 0x20, 0xac, 0x56, 0x82, 0x63, 0xec, 0x14, 0xaa, 0x79, 0x93, 0x40, 0x93,
	0x48, 0xda, 0x3a, 0xfc, 0xd7, 0x77, 0x1a, 0x79, 0xea, 0xd8, 0xa5, 0xbc, 0xf7, 0xfd, 0xf1, 0x79,
	0x0d, 0x1d, 0xef, 0x9c, 0x31, 0xce, 0xca, 0x06, 0xfc, 0x11, 0xbc, 0x34, 0xcd, 0x7e, 0x76, 0xf0,
	0xae, 0x75, 0x2c, 0x56, 0x0a, 0x87, 0x87, 0xfb, 0x4b, 0x02, 0x6c, 0x67, 0xce, 0xde, 0xd3, 0x4f,
	0x8f, 0x4e, 0xae, 0x3d, 0x7f, 0x0c, 0x25, 0xb9, 0xaf, 0x0c, 0xc8, 0x9d, 0xeb, 0x6c, 0xcb, 0xc6,
	0x34, 0xee, 0x1a, 0xf0, 0x52, 0x2b, 0x4e, 0x32, 0x92, 0xdf, 0x8a, 0x28, 0xac, 0x4b, 0xc5, 0xe6,
	0x34, 0xc6, 0x98, 0x56, 0xbc, 0x97, 0x91, 0xfc, 0x6e, 0x3e, 0x99, 0x5d, 0x8e, 0xcc, 0xfe, 0xdd,
	0x40, 0x94, 0x88, 0xc2, 0x77, 0xa9, 0xd8, 0x23, 0x1d, 0x61, 0xc7, 0x76, 0x66, 0x0b, 0x9e, 0xf7,
	0x33, 0x92, 0x0f, 0x05, 0x0d, 0xd2, 0x0a, 0x15, 0xf6, 0x4c, 0x53, 0xef, 0x9c, 0x91, 0xed, 0xe9,
	0x00, 0x7c, 0x80, 0xd8, 0xe9, 0x1f, 0xf6, 0xa5, 0x2c, 0x8a, 0x72, 0x25, 0x17, 0xab, 0xf7, 0x42,
	0x8a, 0xb2, 0x2c, 0x36, 0x9f, 0xeb, 0x85, 0x48, 0x42, 0x7e, 0x73, 0x3a, 0x00, 0x9b, 0x52, 0x8a,
	0xdd, 0x1a, 0x8e, 0x50, 0xf3, 0x21, 0xb2, 0x91, 0xf6, 0x16, 0x04, 0xc6, 0xe8, 0x60, 0xab, 0xeb,
	0x9a, 0x47, 0x19, 0xc9, 0x89, 0xc0, 0x39, 0x3c, 0x4e, 0x37, 0xf2, 0x5b, 0x5b, 0xe0, 0x71, 0x46,
	0xf2, 0x44, 0x44, 0xba, 0xf9, 0xd0, 0x16, 0x42, 0x58, 0x55, 0x6d, 0xc5, 0x93, 0x8c, 0xe4, 0xa9,
	0xc0, 0x39, 0xf0, 0x9b, 0xb6, 0xf2, 0xad, 0x6c, 0xb5, 0x01, 0x9e, 0x66, 0x24, 0xef, 0x8b, 0x14,
	0x95, 0x8d, 0x36, 0xc0, 0x26, 0x34, 0x01, 0xab, 0xce, 0x26, 0x45, 0x33, 0x06, 0xab, 0xae, 0xd6,
	0xce, 0x69, 0x2b, 0xbf, 0x00, 0xf8, 0x08, 0xff, 0x2b, 0x0e, 0xfb, 0x2b, 0xc0, 0xfa, 0xe6, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x83, 0xce, 0x3c, 0x81, 0xb4, 0x01, 0x00, 0x00,
}
