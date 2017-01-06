// Code generated by protoc-gen-go.
// source: zjh_desk.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of EProtoId from zjh_base.proto

// Ignoring public import of zjh_enum_playerGameStatus from zjh_base.proto

// Ignoring public import of zjh_enum_deskState from zjh_base.proto

// Ignoring public import of zjh_enum_userState from zjh_base.proto

// Ignoring public import of zjh_enum_roomType from zjh_base.proto

// 1.请求：房间列表
type ZjhReqGetRoomList struct {
	UserId           *int32           `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	RoomType         *ZjhEnumRoomType `protobuf:"varint,2,opt,name=roomType,enum=ddproto.ZjhEnumRoomType" json:"roomType,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ZjhReqGetRoomList) Reset()                    { *m = ZjhReqGetRoomList{} }
func (m *ZjhReqGetRoomList) String() string            { return proto.CompactTextString(m) }
func (*ZjhReqGetRoomList) ProtoMessage()               {}
func (*ZjhReqGetRoomList) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

func (m *ZjhReqGetRoomList) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ZjhReqGetRoomList) GetRoomType() ZjhEnumRoomType {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return ZjhEnumRoomType_ROOM_TYPE_FRIEND
}

// 服务器返回的经典模式场次
type ZjhBaseRoomInfo struct {
	Header           *ProtoHeader     `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomId           *int32           `protobuf:"varint,2,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *ZjhEnumRoomType `protobuf:"varint,3,opt,name=roomType,enum=ddproto.ZjhEnumRoomType" json:"roomType,omitempty"`
	RoomLevel        *int32           `protobuf:"varint,4,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string          `protobuf:"bytes,5,opt,name=roomTitle" json:"roomTitle,omitempty"`
	BaseValue        *int64           `protobuf:"varint,6,opt,name=baseValue" json:"baseValue,omitempty"`
	MaxValue         *int64           `protobuf:"varint,7,opt,name=maxValue" json:"maxValue,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ZjhBaseRoomInfo) Reset()                    { *m = ZjhBaseRoomInfo{} }
func (m *ZjhBaseRoomInfo) String() string            { return proto.CompactTextString(m) }
func (*ZjhBaseRoomInfo) ProtoMessage()               {}
func (*ZjhBaseRoomInfo) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{1} }

func (m *ZjhBaseRoomInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBaseRoomInfo) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhBaseRoomInfo) GetRoomType() ZjhEnumRoomType {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return ZjhEnumRoomType_ROOM_TYPE_FRIEND
}

func (m *ZjhBaseRoomInfo) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *ZjhBaseRoomInfo) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func (m *ZjhBaseRoomInfo) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *ZjhBaseRoomInfo) GetMaxValue() int64 {
	if m != nil && m.MaxValue != nil {
		return *m.MaxValue
	}
	return 0
}

// 2返回：大厅列表
type ZjhAckRoomList struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Rooms            []*ZjhBaseRoomInfo `protobuf:"bytes,2,rep,name=rooms" json:"rooms,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ZjhAckRoomList) Reset()                    { *m = ZjhAckRoomList{} }
func (m *ZjhAckRoomList) String() string            { return proto.CompactTextString(m) }
func (*ZjhAckRoomList) ProtoMessage()               {}
func (*ZjhAckRoomList) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{2} }

func (m *ZjhAckRoomList) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhAckRoomList) GetRooms() []*ZjhBaseRoomInfo {
	if m != nil {
		return m.Rooms
	}
	return nil
}

// ===============================第二步：加入该房间下系统分配的一个牌桌==============================
// 请求：进入选择的场景房间，自动进入一个有空位的房间，如果没有空位则新建一个空房间进入。
type ZjhReqEnterDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomId           *int32       `protobuf:"varint,2,opt,name=roomId" json:"roomId,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhReqEnterDesk) Reset()                    { *m = ZjhReqEnterDesk{} }
func (m *ZjhReqEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*ZjhReqEnterDesk) ProtoMessage()               {}
func (*ZjhReqEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{3} }

func (m *ZjhReqEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhReqEnterDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhReqEnterDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 用户信息
type ZjhBaseUserInfo struct {
	Uid              *int32            `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	NickName         *string           `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	Coin             *int64            `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Chip             *int32            `protobuf:"varint,4,opt,name=chip" json:"chip,omitempty"`
	State            *ZjhEnumUserState `protobuf:"varint,5,opt,name=state,enum=ddproto.ZjhEnumUserState" json:"state,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *ZjhBaseUserInfo) Reset()                    { *m = ZjhBaseUserInfo{} }
func (m *ZjhBaseUserInfo) String() string            { return proto.CompactTextString(m) }
func (*ZjhBaseUserInfo) ProtoMessage()               {}
func (*ZjhBaseUserInfo) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{4} }

func (m *ZjhBaseUserInfo) GetUid() int32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *ZjhBaseUserInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *ZjhBaseUserInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *ZjhBaseUserInfo) GetChip() int32 {
	if m != nil && m.Chip != nil {
		return *m.Chip
	}
	return 0
}

func (m *ZjhBaseUserInfo) GetState() ZjhEnumUserState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ZjhEnumUserState_USER_IS_GAMING
}

// 返回：该桌面的用户和游戏状态
type ZjhDeskStateAck struct {
	State            *ZjhEnumDeskState  `protobuf:"varint,1,opt,name=state,enum=ddproto.ZjhEnumDeskState" json:"state,omitempty"`
	UserList         []*ZjhBaseUserInfo `protobuf:"bytes,2,rep,name=userList" json:"userList,omitempty"`
	DeskId           *int32             `protobuf:"varint,4,opt,name=deskId" json:"deskId,omitempty"`
	RoomInfo         *ZjhBaseRoomInfo   `protobuf:"bytes,5,opt,name=roomInfo" json:"roomInfo,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ZjhDeskStateAck) Reset()                    { *m = ZjhDeskStateAck{} }
func (m *ZjhDeskStateAck) String() string            { return proto.CompactTextString(m) }
func (*ZjhDeskStateAck) ProtoMessage()               {}
func (*ZjhDeskStateAck) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{5} }

func (m *ZjhDeskStateAck) GetState() ZjhEnumDeskState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ZjhEnumDeskState_DESK_IS_GAMING
}

func (m *ZjhDeskStateAck) GetUserList() []*ZjhBaseUserInfo {
	if m != nil {
		return m.UserList
	}
	return nil
}

func (m *ZjhDeskStateAck) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ZjhDeskStateAck) GetRoomInfo() *ZjhBaseRoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ZjhReqGetRoomList)(nil), "ddproto.zjh_req_getRoomList")
	proto.RegisterType((*ZjhBaseRoomInfo)(nil), "ddproto.zjh_base_roomInfo")
	proto.RegisterType((*ZjhAckRoomList)(nil), "ddproto.zjh_ack_roomList")
	proto.RegisterType((*ZjhReqEnterDesk)(nil), "ddproto.zjh_req_enterDesk")
	proto.RegisterType((*ZjhBaseUserInfo)(nil), "ddproto.zjh_base_userInfo")
	proto.RegisterType((*ZjhDeskStateAck)(nil), "ddproto.ZjhDeskStateAck")
}

var fileDescriptor20 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x66, 0xe3, 0x3a, 0x6d, 0xb6, 0x22, 0x94, 0x2d, 0x42, 0x56, 0xe0, 0x60, 0xf9, 0xe4, 0x03,
	0x8a, 0x20, 0x07, 0xee, 0x48, 0x1c, 0x88, 0x54, 0xa1, 0x68, 0xa9, 0x38, 0x70, 0x89, 0x8c, 0x77,
	0x20, 0xae, 0x7f, 0x36, 0xb5, 0xd7, 0x08, 0x78, 0x01, 0x5e, 0x82, 0x47, 0xe2, 0xa1, 0xd0, 0xcc,
	0xee, 0xda, 0x14, 0x14, 0x04, 0x52, 0x2f, 0xd6, 0xfc, 0x7c, 0xfe, 0x76, 0xe6, 0xfb, 0x86, 0xcf,
	0xbf, 0x5e, 0xed, 0xb6, 0x0a, 0xba, 0x72, 0xb9, 0x6f, 0xb5, 0xd1, 0xe2, 0x58, 0x29, 0x0a, 0x16,
	0xe7, 0xb9, 0xae, 0x6b, 0xdd, 0x6c, 0xf3, 0xaa, 0x80, 0xc6, 0xd8, 0xee, 0x82, 0xd0, 0xef, 0xb3,
	0x0e, 0x6c, 0x9e, 0x00, 0x3f, 0xc7, 0x4a, 0x0b, 0xd7, 0xdb, 0x8f, 0x60, 0xa4, 0xd6, 0xf5, 0x45,
	0xd1, 0x19, 0xf1, 0x90, 0x4f, 0xfb, 0x0e, 0xda, 0xb5, 0x8a, 0x58, 0xcc, 0xd2, 0x50, 0xba, 0x4c,
	0x3c, 0xe7, 0x27, 0xad, 0xd6, 0xf5, 0xe5, 0x97, 0x3d, 0x44, 0x93, 0x98, 0xa5, 0xf3, 0xd5, 0x62,
	0xe9, 0xde, 0x5b, 0x22, 0x0f, 0x34, 0x7d, 0xbd, 0xf5, 0x08, 0x39, 0x60, 0x93, 0x6f, 0x13, 0x7e,
	0xdf, 0xbf, 0x4c, 0xfd, 0x75, 0xf3, 0x41, 0x8b, 0x27, 0x7c, 0xba, 0x83, 0x4c, 0x41, 0x4b, 0xaf,
	0x9c, 0xae, 0x1e, 0x0c, 0x5c, 0x1b, 0xfc, 0xbe, 0xa2, 0x9e, 0x74, 0x18, 0x9c, 0x89, 0xfe, 0x54,
	0xf4, 0x72, 0x28, 0x5d, 0x76, 0x63, 0xa6, 0xe0, 0xdf, 0x67, 0x12, 0x8f, 0xf9, 0x0c, 0xe3, 0x0b,
	0xf8, 0x04, 0x55, 0x74, 0x44, 0x94, 0x63, 0xc1, 0x77, 0x2f, 0x0b, 0x53, 0x41, 0x14, 0xc6, 0x2c,
	0x9d, 0xc9, 0xb1, 0x80, 0x5d, 0x5c, 0xe5, 0x6d, 0x56, 0xf5, 0x10, 0x4d, 0x63, 0x96, 0x06, 0x72,
	0x2c, 0x88, 0x05, 0x3f, 0xa9, 0xb3, 0xcf, 0xb6, 0x79, 0x4c, 0xcd, 0x21, 0x4f, 0x5a, 0x7e, 0x86,
	0x43, 0x65, 0x79, 0x49, 0x33, 0x91, 0xda, 0xff, 0xa7, 0xc3, 0x53, 0x1e, 0xe2, 0x9f, 0x5d, 0x34,
	0x89, 0x83, 0xf4, 0xf4, 0xb7, 0x65, 0x6f, 0x08, 0x2c, 0x2d, 0x30, 0xb9, 0xb6, 0xe2, 0xa3, 0xc9,
	0xd0, 0x18, 0x68, 0x5f, 0x42, 0x57, 0xde, 0x92, 0xf8, 0xe3, 0xa1, 0xa0, 0xf4, 0x77, 0xfd, 0xa1,
	0x24, 0xdf, 0xd9, 0x2f, 0x86, 0x53, 0x0d, 0x0d, 0x3f, 0xe3, 0x41, 0x5f, 0xf8, 0x9b, 0xc2, 0x10,
	0xa5, 0x6a, 0x8a, 0xbc, 0x7c, 0x9d, 0xd5, 0xf6, 0xa0, 0x66, 0x72, 0xc8, 0x85, 0xe0, 0x47, 0xb9,
	0x2e, 0x1a, 0x62, 0x0e, 0x24, 0xc5, 0x54, 0xdb, 0x15, 0x7b, 0xe7, 0x17, 0xc5, 0xe2, 0x19, 0x0f,
	0x3b, 0x93, 0x19, 0x6b, 0xd3, 0x7c, 0xf5, 0xe8, 0x4f, 0xf7, 0x71, 0x80, 0x37, 0x08, 0x91, 0x16,
	0x99, 0xfc, 0x60, 0xfc, 0xde, 0xbb, 0xab, 0x1d, 0x0a, 0x41, 0xf5, 0x17, 0x79, 0x39, 0xd2, 0xb0,
	0x43, 0x34, 0xca, 0xc3, 0x1d, 0x0d, 0x9e, 0x1e, 0x52, 0xa3, 0x89, 0x87, 0xdd, 0xf0, 0xdb, 0xcb,
	0x01, 0x8b, 0xaa, 0x21, 0xd7, 0x5a, 0xb9, 0x3d, 0x5c, 0xe6, 0x4f, 0x19, 0xd1, 0xb4, 0xcc, 0xdf,
	0xdd, 0x1d, 0xb0, 0x9b, 0x3b, 0x1b, 0xf6, 0x33, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xc9, 0x89, 0x3f,
	0x08, 0x04, 0x00, 0x00,
}
