// Code generated by protoc-gen-go.
// source: pdk_hall.proto
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

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

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

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of pdk_base_roomTypeInfo from pdk_base.proto

// Ignoring public import of pdk_base_playerInfo from pdk_base.proto

// Ignoring public import of pdk_base_playerRateInfo from pdk_base.proto

// Ignoring public import of pdk_base_commonRateInfo from pdk_base.proto

// Ignoring public import of pdk_base_timerInfo from pdk_base.proto

// Ignoring public import of pdk_base_deskInfo from pdk_base.proto

// Ignoring public import of pdk_enum_protoId from pdk_base.proto

// Ignoring public import of pdk_enum_errorCode from pdk_base.proto

// Ignoring public import of pdk_enum_paiType from pdk_base.proto

// Ignoring public import of pdk_enum_actType from pdk_base.proto

// Ignoring public import of pdk_enum_gameStatus from pdk_base.proto

// Ignoring public import of pdk_enum_playerStatus from pdk_base.proto

// Ignoring public import of pdk_enum_roomType from pdk_base.proto

// Ignoring public import of pdk_enum_enterType from pdk_base.proto

// Ignoring public import of pdk_enum_coinRoomLevel from pdk_base.proto

// Ignoring public import of pdk_enum_deskGameStatus from pdk_base.proto

// 创建房间
type PdkReqCreateDesk struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomTypeInfo     *PdkBaseRoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkReqCreateDesk) Reset()                    { *m = PdkReqCreateDesk{} }
func (m *PdkReqCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*PdkReqCreateDesk) ProtoMessage()               {}
func (*PdkReqCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{0} }

func (m *PdkReqCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkReqCreateDesk) GetRoomTypeInfo() *PdkBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 创建房间回复的信息
type PdkAckCreateDesk struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32               `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string              `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	UserBalance      *int64               `protobuf:"varint,4,opt,name=userBalance" json:"userBalance,omitempty"`
	CreateFee        *int64               `protobuf:"varint,5,opt,name=createFee" json:"createFee,omitempty"`
	RoomTypeInfo     *PdkBaseRoomTypeInfo `protobuf:"bytes,6,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkAckCreateDesk) Reset()                    { *m = PdkAckCreateDesk{} }
func (m *PdkAckCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*PdkAckCreateDesk) ProtoMessage()               {}
func (*PdkAckCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{1} }

func (m *PdkAckCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkAckCreateDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PdkAckCreateDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *PdkAckCreateDesk) GetUserBalance() int64 {
	if m != nil && m.UserBalance != nil {
		return *m.UserBalance
	}
	return 0
}

func (m *PdkAckCreateDesk) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *PdkAckCreateDesk) GetRoomTypeInfo() *PdkBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 游戏战绩
type PdkReqGameRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32       `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkReqGameRecord) Reset()                    { *m = PdkReqGameRecord{} }
func (m *PdkReqGameRecord) String() string            { return proto.CompactTextString(m) }
func (*PdkReqGameRecord) ProtoMessage()               {}
func (*PdkReqGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{2} }

func (m *PdkReqGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkReqGameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PdkReqGameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

//
type PdkBaseUserRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	NickName         *string      `protobuf:"bytes,3,opt,name=NickName" json:"NickName,omitempty"`
	WinAmount        *int64       `protobuf:"varint,4,opt,name=WinAmount" json:"WinAmount,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkBaseUserRecord) Reset()                    { *m = PdkBaseUserRecord{} }
func (m *PdkBaseUserRecord) String() string            { return proto.CompactTextString(m) }
func (*PdkBaseUserRecord) ProtoMessage()               {}
func (*PdkBaseUserRecord) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{3} }

func (m *PdkBaseUserRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkBaseUserRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkBaseUserRecord) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PdkBaseUserRecord) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

//
type PdkBaseGameRecord struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32               `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	DeskId           *int32               `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	BeginTime        *string              `protobuf:"bytes,4,opt,name=beginTime" json:"beginTime,omitempty"`
	Users            []*PdkBaseUserRecord `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkBaseGameRecord) Reset()                    { *m = PdkBaseGameRecord{} }
func (m *PdkBaseGameRecord) String() string            { return proto.CompactTextString(m) }
func (*PdkBaseGameRecord) ProtoMessage()               {}
func (*PdkBaseGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{4} }

func (m *PdkBaseGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkBaseGameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PdkBaseGameRecord) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PdkBaseGameRecord) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *PdkBaseGameRecord) GetUsers() []*PdkBaseUserRecord {
	if m != nil {
		return m.Users
	}
	return nil
}

//
type PdkAckGameRecord struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32              `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Records          []*PdkBaseGameRecord `protobuf:"bytes,3,rep,name=records" json:"records,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkAckGameRecord) Reset()                    { *m = PdkAckGameRecord{} }
func (m *PdkAckGameRecord) String() string            { return proto.CompactTextString(m) }
func (*PdkAckGameRecord) ProtoMessage()               {}
func (*PdkAckGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{5} }

func (m *PdkAckGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkAckGameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkAckGameRecord) GetRecords() []*PdkBaseGameRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

// 客户端请求进入 room, 服务器返回DdzSendGameInfo
type PdkReqEnterDesk struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MatchId          *int32                `protobuf:"varint,2,opt,name=matchId" json:"matchId,omitempty"`
	TableId          *int32                `protobuf:"varint,3,opt,name=tableId" json:"tableId,omitempty"`
	PassWord         *string               `protobuf:"bytes,4,opt,name=PassWord" json:"PassWord,omitempty"`
	UserId           *uint32               `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	EnterType        *PdkEnumEnterType     `protobuf:"varint,6,opt,name=enterType,enum=ddproto.PdkEnumEnterType" json:"enterType,omitempty"`
	CoinRoomLevel    *PdkEnumCoinRoomLevel `protobuf:"varint,7,opt,name=coinRoomLevel,enum=ddproto.PdkEnumCoinRoomLevel" json:"coinRoomLevel,omitempty"`
	IsContinuous     *bool                 `protobuf:"varint,8,opt,name=isContinuous" json:"isContinuous,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PdkReqEnterDesk) Reset()                    { *m = PdkReqEnterDesk{} }
func (m *PdkReqEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*PdkReqEnterDesk) ProtoMessage()               {}
func (*PdkReqEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{6} }

func (m *PdkReqEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkReqEnterDesk) GetMatchId() int32 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *PdkReqEnterDesk) GetTableId() int32 {
	if m != nil && m.TableId != nil {
		return *m.TableId
	}
	return 0
}

func (m *PdkReqEnterDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

func (m *PdkReqEnterDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkReqEnterDesk) GetEnterType() PdkEnumEnterType {
	if m != nil && m.EnterType != nil {
		return *m.EnterType
	}
	return PdkEnumEnterType_PDK_FRIEND
}

func (m *PdkReqEnterDesk) GetCoinRoomLevel() PdkEnumCoinRoomLevel {
	if m != nil && m.CoinRoomLevel != nil {
		return *m.CoinRoomLevel
	}
	return PdkEnumCoinRoomLevel_PDK_LV1
}

func (m *PdkReqEnterDesk) GetIsContinuous() bool {
	if m != nil && m.IsContinuous != nil {
		return *m.IsContinuous
	}
	return false
}

type PdkAckEnterDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	IsContinuous     *bool        `protobuf:"varint,2,opt,name=isContinuous" json:"isContinuous,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PdkAckEnterDesk) Reset()                    { *m = PdkAckEnterDesk{} }
func (m *PdkAckEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*PdkAckEnterDesk) ProtoMessage()               {}
func (*PdkAckEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor32, []int{7} }

func (m *PdkAckEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PdkAckEnterDesk) GetIsContinuous() bool {
	if m != nil && m.IsContinuous != nil {
		return *m.IsContinuous
	}
	return false
}

func init() {
	proto.RegisterType((*PdkReqCreateDesk)(nil), "ddproto.pdk_req_createDesk")
	proto.RegisterType((*PdkAckCreateDesk)(nil), "ddproto.pdk_ack_createDesk")
	proto.RegisterType((*PdkReqGameRecord)(nil), "ddproto.pdk_req_gameRecord")
	proto.RegisterType((*PdkBaseUserRecord)(nil), "ddproto.pdk_base_userRecord")
	proto.RegisterType((*PdkBaseGameRecord)(nil), "ddproto.pdk_base_gameRecord")
	proto.RegisterType((*PdkAckGameRecord)(nil), "ddproto.pdk_ack_gameRecord")
	proto.RegisterType((*PdkReqEnterDesk)(nil), "ddproto.pdk_req_enterDesk")
	proto.RegisterType((*PdkAckEnterDesk)(nil), "ddproto.pdk_ack_enterDesk")
}

var fileDescriptor32 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x26, 0x09, 0xfd, 0x73, 0x77, 0x2b, 0xe1, 0x45, 0xab, 0xa8, 0x54, 0x10, 0xe5, 0xd4, 0x03,
	0xea, 0xa1, 0x07, 0x24, 0x8e, 0x94, 0x1f, 0x6d, 0x25, 0xb4, 0xaa, 0xac, 0x45, 0x7b, 0xac, 0xdc,
	0x64, 0xd8, 0x9a, 0xc6, 0x76, 0x89, 0x13, 0x10, 0x2f, 0xc0, 0x03, 0x20, 0xf1, 0x38, 0x3c, 0x03,
	0xef, 0xc3, 0x09, 0xd9, 0xce, 0x2f, 0x15, 0x07, 0x22, 0x2e, 0x51, 0xbe, 0xf1, 0x8c, 0xfd, 0x7d,
	0xdf, 0xcc, 0xa0, 0xc9, 0x31, 0x3e, 0x6c, 0xf7, 0x34, 0x49, 0x16, 0xc7, 0x54, 0x66, 0x12, 0x0f,
	0xe2, 0xd8, 0xfc, 0x4c, 0x2f, 0x22, 0xc9, 0xb9, 0x14, 0xdb, 0x28, 0x61, 0x20, 0x32, 0x7b, 0x3a,
	0x35, 0xd9, 0x3b, 0xaa, 0xc0, 0xe2, 0xf0, 0xab, 0x83, 0xb0, 0x0e, 0xa5, 0xf0, 0x71, 0x1b, 0xa5,
	0x40, 0x33, 0x78, 0x05, 0xea, 0x80, 0x9f, 0xa2, 0xfe, 0x1e, 0x68, 0x0c, 0xa9, 0xef, 0x04, 0xce,
	0x7c, 0xbc, 0x7c, 0xb8, 0x28, 0x6e, 0x5d, 0x6c, 0xf4, 0xf7, 0xca, 0x9c, 0x91, 0x22, 0x07, 0xaf,
	0xd0, 0x59, 0x2a, 0x25, 0xbf, 0xf9, 0x72, 0x84, 0xb5, 0x78, 0x2f, 0x7d, 0xd7, 0xd4, 0x3c, 0xae,
	0x6a, 0xca, 0x37, 0xb7, 0xcd, 0x2c, 0xd2, 0xaa, 0x09, 0x7f, 0x15, 0x44, 0x68, 0x74, 0xe8, 0x4e,
	0xe4, 0x12, 0xf5, 0x63, 0x50, 0x87, 0x75, 0x6c, 0x28, 0xf4, 0x48, 0x81, 0xf0, 0x14, 0x0d, 0x8f,
	0x54, 0xa9, 0xcf, 0x32, 0x8d, 0x7d, 0x2f, 0x70, 0xe6, 0x23, 0x52, 0x61, 0x1c, 0xa0, 0x71, 0xae,
	0x20, 0x5d, 0xd1, 0x84, 0x8a, 0x08, 0xfc, 0xfb, 0x81, 0x33, 0xf7, 0x48, 0x33, 0x84, 0x67, 0x68,
	0x64, 0x19, 0xbd, 0x01, 0xf0, 0x7b, 0xe6, 0xbc, 0x0e, 0x9c, 0x88, 0xef, 0x77, 0x10, 0xff, 0xa1,
	0x6e, 0xc2, 0x1d, 0xe5, 0x40, 0x20, 0xd2, 0xcc, 0xfe, 0x4d, 0xfb, 0x04, 0xb9, 0xac, 0xd4, 0xed,
	0xb2, 0x58, 0x7b, 0xa1, 0x45, 0xac, 0xad, 0xe2, 0x73, 0x52, 0xa0, 0xf0, 0xbb, 0x83, 0x2e, 0x2a,
	0x4e, 0x3a, 0xd6, 0xe9, 0xb5, 0x4b, 0xd4, 0x7f, 0x67, 0x6f, 0x77, 0xed, 0xed, 0x16, 0x69, 0xa7,
	0xaf, 0x59, 0x74, 0xb8, 0xa6, 0x1c, 0x4a, 0xa7, 0x4b, 0xac, 0x7d, 0xbc, 0x65, 0xe2, 0x05, 0x97,
	0xb9, 0xc8, 0x0a, 0x9f, 0xeb, 0x40, 0xf8, 0xa3, 0xc9, 0xeb, 0x7f, 0xba, 0x50, 0x4c, 0x84, 0xd7,
	0x9a, 0x88, 0x19, 0x1a, 0xed, 0xe0, 0x8e, 0x89, 0x1b, 0xc6, 0x6d, 0xcf, 0x47, 0xa4, 0x0e, 0xe0,
	0x25, 0xea, 0x69, 0x67, 0x94, 0xdf, 0x0b, 0xbc, 0xf9, 0x78, 0x39, 0x3b, 0x6d, 0x66, 0x6d, 0x1c,
	0xb1, 0xa9, 0xe1, 0xb7, 0xc6, 0x00, 0x77, 0xa6, 0x5f, 0x37, 0xcd, 0x6d, 0x36, 0x0d, 0x3f, 0x43,
	0x83, 0xd4, 0xdc, 0xa7, 0x7c, 0xef, 0x6f, 0x94, 0xea, 0x47, 0x49, 0x99, 0x1c, 0xfe, 0x74, 0xd1,
	0x83, 0x72, 0xb2, 0x40, 0x64, 0x90, 0x76, 0x58, 0x2a, 0x1f, 0x0d, 0x38, 0xcd, 0xa2, 0x7d, 0xb5,
	0x55, 0x25, 0xd4, 0x27, 0x19, 0xdd, 0x25, 0x50, 0xb9, 0x5b, 0x42, 0x3d, 0x06, 0x1b, 0xaa, 0xd4,
	0xad, 0x5e, 0x38, 0xeb, 0x6e, 0x85, 0x1b, 0x1a, 0x7b, 0x2d, 0x8d, 0xcf, 0xd1, 0xc8, 0x50, 0xd4,
	0x5b, 0x61, 0xb6, 0x68, 0xb2, 0x7c, 0xd4, 0x52, 0x09, 0x22, 0xe7, 0xdb, 0x2a, 0x85, 0xd4, 0xd9,
	0xf8, 0x35, 0x3a, 0x8f, 0x24, 0x13, 0x44, 0x4a, 0xfe, 0x16, 0x3e, 0x41, 0xe2, 0x0f, 0x4c, 0xf9,
	0x93, 0xd3, 0xf2, 0x56, 0x1a, 0x69, 0x57, 0xe1, 0x10, 0x9d, 0x31, 0xf5, 0x52, 0x8a, 0x8c, 0x89,
	0x5c, 0xe6, 0xca, 0x1f, 0x06, 0xce, 0x7c, 0x48, 0x5a, 0xb1, 0x10, 0xac, 0xa1, 0xba, 0xcb, 0x5d,
	0x0d, 0xfd, 0xf3, 0x19, 0xf7, 0xf4, 0x99, 0x95, 0x7b, 0xe5, 0x6d, 0xee, 0x6d, 0x9c, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x26, 0xa3, 0xa8, 0x31, 0xde, 0x05, 0x00, 0x00,
}
