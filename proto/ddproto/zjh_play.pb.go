// Code generated by protoc-gen-go.
// source: zjh_play.proto
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

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of EProtoId from zjh_base.proto

// Ignoring public import of zjh_enum_playerGameStatus from zjh_base.proto

// Ignoring public import of zjh_enum_deskState from zjh_base.proto

// Ignoring public import of zjh_enum_userState from zjh_base.proto

// Ignoring public import of zjh_enum_roomType from zjh_base.proto

// Ignoring public import of zjh_req_getRoomList from zjh_desk.proto

// Ignoring public import of zjh_base_roomInfo from zjh_desk.proto

// Ignoring public import of zjh_ack_roomList from zjh_desk.proto

// Ignoring public import of zjh_req_enterDesk from zjh_desk.proto

// Ignoring public import of zjh_base_userInfo from zjh_desk.proto

// Ignoring public import of ZjhDeskStateAck from zjh_desk.proto

// 游戏信息(广播)（接收服务端消息）(别的玩家已看牌是独立协议更好吧? 否则通过此协议下发是否过于冗余?)
type ZjhBcGameInfo struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfo       []*ZjhBasePlayerInfo `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	ZjhDeskInfo      *ZjhBaseDeskInfo     `protobuf:"bytes,3,opt,name=zjhDeskInfo" json:"zjhDeskInfo,omitempty"`
	SenderUserId     *uint32              `protobuf:"varint,4,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *int32               `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhBcGameInfo) Reset()                    { *m = ZjhBcGameInfo{} }
func (m *ZjhBcGameInfo) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcGameInfo) ProtoMessage()               {}
func (*ZjhBcGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *ZjhBcGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcGameInfo) GetPlayerInfo() []*ZjhBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *ZjhBcGameInfo) GetZjhDeskInfo() *ZjhBaseDeskInfo {
	if m != nil {
		return m.ZjhDeskInfo
	}
	return nil
}

func (m *ZjhBcGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *ZjhBcGameInfo) GetIsReconnect() int32 {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return 0
}

// 新加入玩家后 桌子广播需要的协议
type ZjhBaseDeskInfo struct {
	GameStatus       *int32           `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	RoomInfo         *ZjhBaseRoomInfo `protobuf:"bytes,2,opt,name=roomInfo" json:"roomInfo,omitempty"`
	PlayerNum        *int32           `protobuf:"varint,3,opt,name=playerNum" json:"playerNum,omitempty"`
	ActiveUserId     *uint32          `protobuf:"varint,4,opt,name=activeUserId" json:"activeUserId,omitempty"`
	ActionTime       *int32           `protobuf:"varint,5,opt,name=actionTime" json:"actionTime,omitempty"`
	NInitActionTime  *int32           `protobuf:"varint,6,opt,name=nInitActionTime" json:"nInitActionTime,omitempty"`
	InitRoomCoin     *int64           `protobuf:"varint,7,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32           `protobuf:"varint,8,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32           `protobuf:"varint,9,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	RoomNumber       *string          `protobuf:"bytes,10,opt,name=roomNumber" json:"roomNumber,omitempty"`
	RoomOwnerUserId  *uint32          `protobuf:"varint,11,opt,name=roomOwnerUserId" json:"roomOwnerUserId,omitempty"`
	PlayRate         *int32           `protobuf:"varint,12,opt,name=playRate" json:"playRate,omitempty"`
	CurrRoundCount   *int32           `protobuf:"varint,13,opt,name=currRoundCount" json:"currRoundCount,omitempty"`
	TotalRoundCount  *int32           `protobuf:"varint,14,opt,name=totalRoundCount" json:"totalRoundCount,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ZjhBaseDeskInfo) Reset()                    { *m = ZjhBaseDeskInfo{} }
func (m *ZjhBaseDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*ZjhBaseDeskInfo) ProtoMessage()               {}
func (*ZjhBaseDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *ZjhBaseDeskInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetRoomInfo() *ZjhBaseRoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

func (m *ZjhBaseDeskInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetActionTime() int32 {
	if m != nil && m.ActionTime != nil {
		return *m.ActionTime
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetNInitActionTime() int32 {
	if m != nil && m.NInitActionTime != nil {
		return *m.NInitActionTime
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *ZjhBaseDeskInfo) GetRoomOwnerUserId() uint32 {
	if m != nil && m.RoomOwnerUserId != nil {
		return *m.RoomOwnerUserId
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetCurrRoundCount() int32 {
	if m != nil && m.CurrRoundCount != nil {
		return *m.CurrRoundCount
	}
	return 0
}

func (m *ZjhBaseDeskInfo) GetTotalRoundCount() int32 {
	if m != nil && m.TotalRoundCount != nil {
		return *m.TotalRoundCount
	}
	return 0
}

// 有玩家进入房间(服务器广播)
type ZjhBcNewPlayerEnter struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfo       *ZjhBasePlayerInfo `protobuf:"bytes,2,opt,name=playerInfo" json:"playerInfo,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ZjhBcNewPlayerEnter) Reset()                    { *m = ZjhBcNewPlayerEnter{} }
func (m *ZjhBcNewPlayerEnter) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcNewPlayerEnter) ProtoMessage()               {}
func (*ZjhBcNewPlayerEnter) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *ZjhBcNewPlayerEnter) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcNewPlayerEnter) GetPlayerInfo() *ZjhBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

type ZjhBasePlayerInfo struct {
	IsFirst          *bool                    `protobuf:"varint,1,opt,name=isFirst" json:"isFirst,omitempty"`
	PlayerPokers     []*ClientBasePoker       `protobuf:"bytes,2,rep,name=playerPokers" json:"playerPokers,omitempty"`
	Coin             *int64                   `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	NickName         *string                  `protobuf:"bytes,4,opt,name=nickName" json:"nickName,omitempty"`
	Sex              *int32                   `protobuf:"varint,5,opt,name=sex" json:"sex,omitempty"`
	UserId           *uint32                  `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	BReady           *int32                   `protobuf:"varint,8,opt,name=bReady" json:"bReady,omitempty"`
	Status           *ZjhEnumPlayerGameStatus `protobuf:"varint,9,opt,name=status,enum=ddproto.ZjhEnumPlayerGameStatus" json:"status,omitempty"`
	WxInfo           *WeixinInfo              `protobuf:"bytes,10,opt,name=wxInfo" json:"wxInfo,omitempty"`
	OnlineStatus     *int32                   `protobuf:"varint,11,opt,name=onlineStatus" json:"onlineStatus,omitempty"`
	IsCheckedPokers  *bool                    `protobuf:"varint,12,opt,name=isCheckedPokers" json:"isCheckedPokers,omitempty"`
	SeatIndex        *int32                   `protobuf:"varint,13,opt,name=seatIndex" json:"seatIndex,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ZjhBasePlayerInfo) Reset()                    { *m = ZjhBasePlayerInfo{} }
func (m *ZjhBasePlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*ZjhBasePlayerInfo) ProtoMessage()               {}
func (*ZjhBasePlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *ZjhBasePlayerInfo) GetIsFirst() bool {
	if m != nil && m.IsFirst != nil {
		return *m.IsFirst
	}
	return false
}

func (m *ZjhBasePlayerInfo) GetPlayerPokers() []*ClientBasePoker {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

func (m *ZjhBasePlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *ZjhBasePlayerInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetBReady() int32 {
	if m != nil && m.BReady != nil {
		return *m.BReady
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetStatus() ZjhEnumPlayerGameStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ZjhEnumPlayerGameStatus_ZJH_TEMP
}

func (m *ZjhBasePlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *ZjhBasePlayerInfo) GetOnlineStatus() int32 {
	if m != nil && m.OnlineStatus != nil {
		return *m.OnlineStatus
	}
	return 0
}

func (m *ZjhBasePlayerInfo) GetIsCheckedPokers() bool {
	if m != nil && m.IsCheckedPokers != nil {
		return *m.IsCheckedPokers
	}
	return false
}

func (m *ZjhBasePlayerInfo) GetSeatIndex() int32 {
	if m != nil && m.SeatIndex != nil {
		return *m.SeatIndex
	}
	return 0
}

// 开局（接收服务端消息）: 客户端收到消息后就开始【霸底】、【发牌】的动画, 然后处理第一个 OverTurn
type ZjhBcOpening struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfoList   []*ZjhBasePlayerInfo `protobuf:"bytes,2,rep,name=playerInfoList" json:"playerInfoList,omitempty"`
	BaseAnte         *int64               `protobuf:"varint,3,opt,name=baseAnte" json:"baseAnte,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhBcOpening) Reset()                    { *m = ZjhBcOpening{} }
func (m *ZjhBcOpening) String() string            { return proto.CompactTextString(m) }
func (*ZjhBcOpening) ProtoMessage()               {}
func (*ZjhBcOpening) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *ZjhBcOpening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ZjhBcOpening) GetPlayerInfoList() []*ZjhBasePlayerInfo {
	if m != nil {
		return m.PlayerInfoList
	}
	return nil
}

func (m *ZjhBcOpening) GetBaseAnte() int64 {
	if m != nil && m.BaseAnte != nil {
		return *m.BaseAnte
	}
	return 0
}

func init() {
	proto.RegisterType((*ZjhBcGameInfo)(nil), "ddproto.zjh_bc_gameInfo")
	proto.RegisterType((*ZjhBaseDeskInfo)(nil), "ddproto.zjh_base_deskInfo")
	proto.RegisterType((*ZjhBcNewPlayerEnter)(nil), "ddproto.zjh_bc_newPlayerEnter")
	proto.RegisterType((*ZjhBasePlayerInfo)(nil), "ddproto.zjh_base_playerInfo")
	proto.RegisterType((*ZjhBcOpening)(nil), "ddproto.zjh_bc_opening")
}

var fileDescriptor14 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x4d, 0x93, 0x26, 0x93, 0x34, 0x85, 0x2d, 0x20, 0x2b, 0xaa, 0x90, 0x65, 0x21, 0x64,
	0x09, 0xd4, 0x43, 0x0f, 0x1c, 0x50, 0x85, 0x54, 0xb5, 0x7c, 0x44, 0x42, 0xc5, 0x5a, 0x40, 0x1c,
	0x23, 0xc7, 0x1e, 0x9a, 0x6d, 0xe2, 0xdd, 0xca, 0xbb, 0xa6, 0x69, 0x6f, 0x88, 0x3f, 0x02, 0x7f,
	0x82, 0xdf, 0x87, 0x76, 0xbd, 0xf1, 0x47, 0x0a, 0x12, 0x15, 0x97, 0xc8, 0xf3, 0xf2, 0xc6, 0xf3,
	0xe6, 0xcd, 0x8c, 0x61, 0x78, 0x7d, 0x3e, 0x9b, 0x5c, 0x2c, 0xa2, 0xab, 0xfd, 0x8b, 0x4c, 0x28,
	0x41, 0xb6, 0x92, 0xc4, 0x3c, 0x8c, 0x76, 0x63, 0x91, 0xa6, 0x82, 0x4f, 0xe2, 0x05, 0x43, 0xae,
	0x8a, 0x7f, 0x47, 0x86, 0x3d, 0x8d, 0x24, 0xd6, 0xe3, 0x04, 0xe5, 0xbc, 0x88, 0xfd, 0x6f, 0x1b,
	0xb0, 0x63, 0x28, 0xf1, 0xe4, 0x2c, 0x4a, 0x71, 0xcc, 0xbf, 0x08, 0xf2, 0x0c, 0x3a, 0x33, 0x8c,
	0x12, 0xcc, 0x5c, 0xc7, 0x73, 0x82, 0xfe, 0xc1, 0xfd, 0x7d, 0x5b, 0x62, 0x3f, 0xd4, 0xbf, 0x6f,
	0xcd, 0x7f, 0xd4, 0x72, 0xc8, 0x21, 0x80, 0x56, 0x83, 0x99, 0xce, 0x75, 0x37, 0xbc, 0x56, 0xd0,
	0x3f, 0xd8, 0x2b, 0x33, 0x56, 0xe5, 0x27, 0x15, 0x87, 0xd6, 0xf8, 0xe4, 0x10, 0xfa, 0xd7, 0xe7,
	0xb3, 0x13, 0x94, 0x73, 0x93, 0xde, 0x32, 0x05, 0x47, 0x37, 0xd3, 0x13, 0xcb, 0xa0, 0x75, 0x3a,
	0xf1, 0x61, 0x20, 0x91, 0x27, 0x98, 0x7d, 0x92, 0x98, 0x8d, 0x13, 0x77, 0xd3, 0x73, 0x82, 0x6d,
	0xda, 0xc0, 0x88, 0x07, 0x7d, 0x26, 0x29, 0xc6, 0x82, 0x73, 0x8c, 0x95, 0xdb, 0xf6, 0x9c, 0xa0,
	0x4d, 0xeb, 0x90, 0xff, 0x73, 0x13, 0xee, 0xdd, 0x28, 0x44, 0x1e, 0x01, 0xbc, 0x89, 0x52, 0xfc,
	0xa0, 0x22, 0x95, 0x4b, 0xe3, 0x44, 0x9b, 0xd6, 0x10, 0xf2, 0x1c, 0xba, 0x99, 0x10, 0xa9, 0xed,
	0xfa, 0x2f, 0xb2, 0x57, 0x0c, 0x5a, 0x72, 0xc9, 0x1e, 0xf4, 0x8a, 0xfe, 0x4f, 0xf3, 0xd4, 0xf4,
	0xdb, 0xa6, 0x15, 0xa0, 0x3b, 0x8a, 0x62, 0xc5, 0xbe, 0x62, 0xb3, 0xa3, 0x3a, 0xa6, 0x95, 0xe9,
	0x58, 0xf0, 0x8f, 0x2c, 0x45, 0xdb, 0x50, 0x0d, 0x21, 0x01, 0xec, 0xf0, 0x31, 0x67, 0xea, 0xa8,
	0x22, 0x75, 0x0c, 0x69, 0x1d, 0xd6, 0xd5, 0x18, 0x67, 0x8a, 0x0a, 0x91, 0x1e, 0x0b, 0xc6, 0xdd,
	0x2d, 0xcf, 0x09, 0x5a, 0xb4, 0x81, 0x91, 0xc7, 0xb0, 0x1d, 0xe7, 0x59, 0x16, 0x2e, 0xa2, 0xab,
	0x63, 0x91, 0x73, 0xe5, 0x76, 0xcd, 0xbb, 0x9a, 0x20, 0x79, 0x02, 0x43, 0x25, 0x54, 0xb4, 0xa8,
	0x68, 0x3d, 0x43, 0x5b, 0x43, 0xb5, 0x76, 0xed, 0xc4, 0x69, 0x9e, 0x4e, 0x31, 0x73, 0xc1, 0x73,
	0x82, 0x1e, 0xad, 0x21, 0x5a, 0xbb, 0x8e, 0xde, 0x5f, 0xf2, 0x72, 0xa8, 0x7d, 0x63, 0xc1, 0x3a,
	0x4c, 0x46, 0xd0, 0xd5, 0xb6, 0xd1, 0x48, 0xa1, 0x3b, 0x30, 0xb5, 0xca, 0x58, 0xab, 0xd1, 0xf2,
	0xa8, 0xc8, 0x79, 0x52, 0xa8, 0xd9, 0x2e, 0xd4, 0x34, 0x51, 0x5d, 0xcd, 0xe8, 0xab, 0x11, 0x87,
	0x85, 0x53, 0x6b, 0xb0, 0xff, 0xdd, 0x81, 0x07, 0xf6, 0x4e, 0x38, 0x5e, 0x86, 0x66, 0x5e, 0xaf,
	0xb8, 0xc2, 0xec, 0x3f, 0xaf, 0xc5, 0xb9, 0xcd, 0xb5, 0xf8, 0xbf, 0x5a, 0xb0, 0xfb, 0x07, 0x0e,
	0x71, 0x61, 0x8b, 0xc9, 0xd7, 0x2c, 0x93, 0xca, 0x88, 0xe8, 0xd2, 0x55, 0x48, 0x5e, 0xc2, 0xa0,
	0xe0, 0x85, 0x62, 0x8e, 0x99, 0xb4, 0xf7, 0x59, 0x6d, 0x6a, 0xf1, 0xb1, 0xb0, 0x2f, 0xd4, 0x14,
	0xda, 0xe0, 0x13, 0x02, 0x9b, 0xb1, 0xde, 0x8c, 0x96, 0xd9, 0x0c, 0xf3, 0xac, 0x9d, 0xe7, 0x2c,
	0x9e, 0x9f, 0x46, 0x29, 0x9a, 0xfd, 0xec, 0xd1, 0x32, 0x26, 0x77, 0xa1, 0x25, 0x71, 0x69, 0x97,
	0x52, 0x3f, 0x92, 0x87, 0xd0, 0xc9, 0x8b, 0x41, 0x76, 0xcc, 0x20, 0x6d, 0xa4, 0xf1, 0x29, 0xc5,
	0x28, 0xb9, 0xb2, 0x0b, 0x65, 0x23, 0xf2, 0x02, 0x3a, 0xb2, 0xb8, 0x39, 0xbd, 0x41, 0xc3, 0x03,
	0xbf, 0xe1, 0x0e, 0xf2, 0x3c, 0xb5, 0x9d, 0x57, 0xb7, 0x48, 0x6d, 0x06, 0x79, 0x0a, 0x9d, 0xcb,
	0xa5, 0x71, 0x16, 0x8c, 0xb3, 0xbb, 0x65, 0xee, 0x67, 0x64, 0x4b, 0xc6, 0x8d, 0xa1, 0x96, 0xa2,
	0x97, 0x5f, 0xf0, 0x05, 0xe3, 0xab, 0x13, 0xef, 0x1b, 0x19, 0x0d, 0x4c, 0x2f, 0x08, 0x93, 0xc7,
	0x33, 0x8c, 0xe7, 0x98, 0x58, 0x07, 0x07, 0xc6, 0xe0, 0x75, 0x58, 0x9f, 0xb5, 0xc4, 0x48, 0x8d,
	0x79, 0x82, 0x4b, 0xbb, 0x6d, 0x15, 0xe0, 0xff, 0x70, 0x8a, 0xef, 0xf6, 0x34, 0x9e, 0x88, 0x0b,
	0xe4, 0x8c, 0x9f, 0xdd, 0x72, 0x6f, 0x4e, 0x60, 0x58, 0xcd, 0xfb, 0x1d, 0x93, 0xea, 0x9f, 0xbe,
	0xb4, 0x6b, 0x39, 0x7a, 0x72, 0x9a, 0x72, 0xc4, 0x15, 0xda, 0x89, 0x96, 0x71, 0x78, 0x27, 0x74,
	0xc2, 0x8d, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0x6e, 0xf2, 0x74, 0x5e, 0x06, 0x00, 0x00,
}
