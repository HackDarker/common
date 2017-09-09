// Code generated by protoc-gen-go. DO NOT EDIT.
// source: phz_base.proto

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

// Ignoring public import of common_req_kickout from common_client.proto

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

// Ignoring public import of common_req_upload_location from common_client.proto

// Ignoring public import of common_bc_leaveTimeout from common_client.proto

// Ignoring public import of common_desk_by_agent from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

type PhzEnumProtoId int32

const (
	PhzEnumProtoId_PHZ_PID_HEARTBEAT             PhzEnumProtoId = 0
	PhzEnumProtoId_PHZ_PID_CREATEDESK_REQ        PhzEnumProtoId = 1
	PhzEnumProtoId_PHZ_PID_ENTERDESK_REQ         PhzEnumProtoId = 2
	PhzEnumProtoId_PHZ_PID_DESK_ACK              PhzEnumProtoId = 3
	PhzEnumProtoId_PHZ_PID_GAMEINFO_REQ          PhzEnumProtoId = 4
	PhzEnumProtoId_PHZ_PID_GAMEINFO              PhzEnumProtoId = 5
	PhzEnumProtoId_PHZ_PID_READY_REQ             PhzEnumProtoId = 6
	PhzEnumProtoId_PHZ_PID_READY_ACK             PhzEnumProtoId = 7
	PhzEnumProtoId_PHZ_PID_OPENING               PhzEnumProtoId = 8
	PhzEnumProtoId_PHZ_PID_SENDCARDS             PhzEnumProtoId = 9
	PhzEnumProtoId_PHZ_PID_OUTCARD_REQ           PhzEnumProtoId = 10
	PhzEnumProtoId_PHZ_PID_OUTCARD_ACK           PhzEnumProtoId = 11
	PhzEnumProtoId_PHZ_PID_MOPAI                 PhzEnumProtoId = 12
	PhzEnumProtoId_PHZ_PID_OVERTURN              PhzEnumProtoId = 13
	PhzEnumProtoId_PHZ_PID_CANPENG               PhzEnumProtoId = 14
	PhzEnumProtoId_PHZ_PID_PENG_REQ              PhzEnumProtoId = 15
	PhzEnumProtoId_PHZ_PID_PENG_ACK              PhzEnumProtoId = 16
	PhzEnumProtoId_PHZ_PID_CHI_REQ               PhzEnumProtoId = 17
	PhzEnumProtoId_PHZ_PID_CANCHI                PhzEnumProtoId = 18
	PhzEnumProtoId_PHZ_PID_CHI_ACK               PhzEnumProtoId = 19
	PhzEnumProtoId_PHZ_PID_TIPAI                 PhzEnumProtoId = 20
	PhzEnumProtoId_PHZ_PID_WEIPAI                PhzEnumProtoId = 21
	PhzEnumProtoId_PHZ_PID_PAOPAI                PhzEnumProtoId = 22
	PhzEnumProtoId_PHZ_PID_HUPAI_REQ             PhzEnumProtoId = 23
	PhzEnumProtoId_PHZ_PID_CANHUPAI              PhzEnumProtoId = 24
	PhzEnumProtoId_PHZ_PID_HUPAI_ACK             PhzEnumProtoId = 25
	PhzEnumProtoId_PHZ_PID_PASS_REQ              PhzEnumProtoId = 26
	PhzEnumProtoId_PHZ_PID_PASS_ACK              PhzEnumProtoId = 27
	PhzEnumProtoId_PHZ_PID_CURRENTRESULT         PhzEnumProtoId = 28
	PhzEnumProtoId_PHZ_PID_ENDRESULT             PhzEnumProtoId = 29
	PhzEnumProtoId_PHZ_PID_APPLYDISSOLVE_REQ     PhzEnumProtoId = 30
	PhzEnumProtoId_PHZ_PID_APPLYDISSOLVE_ACK     PhzEnumProtoId = 31
	PhzEnumProtoId_PHZ_PID_APPLYDISSOLVEBACK_REQ PhzEnumProtoId = 32
	PhzEnumProtoId_PHZ_PID_APPLYDISSOLVEBACK_ACK PhzEnumProtoId = 33
	PhzEnumProtoId_PHZ_PID_DISSOLVEDESK_REQ      PhzEnumProtoId = 34
	PhzEnumProtoId_PHZ_PID_DISSOLVEDESK_ACK      PhzEnumProtoId = 35
)

var PhzEnumProtoId_name = map[int32]string{
	0:  "PHZ_PID_HEARTBEAT",
	1:  "PHZ_PID_CREATEDESK_REQ",
	2:  "PHZ_PID_ENTERDESK_REQ",
	3:  "PHZ_PID_DESK_ACK",
	4:  "PHZ_PID_GAMEINFO_REQ",
	5:  "PHZ_PID_GAMEINFO",
	6:  "PHZ_PID_READY_REQ",
	7:  "PHZ_PID_READY_ACK",
	8:  "PHZ_PID_OPENING",
	9:  "PHZ_PID_SENDCARDS",
	10: "PHZ_PID_OUTCARD_REQ",
	11: "PHZ_PID_OUTCARD_ACK",
	12: "PHZ_PID_MOPAI",
	13: "PHZ_PID_OVERTURN",
	14: "PHZ_PID_CANPENG",
	15: "PHZ_PID_PENG_REQ",
	16: "PHZ_PID_PENG_ACK",
	17: "PHZ_PID_CHI_REQ",
	18: "PHZ_PID_CANCHI",
	19: "PHZ_PID_CHI_ACK",
	20: "PHZ_PID_TIPAI",
	21: "PHZ_PID_WEIPAI",
	22: "PHZ_PID_PAOPAI",
	23: "PHZ_PID_HUPAI_REQ",
	24: "PHZ_PID_CANHUPAI",
	25: "PHZ_PID_HUPAI_ACK",
	26: "PHZ_PID_PASS_REQ",
	27: "PHZ_PID_PASS_ACK",
	28: "PHZ_PID_CURRENTRESULT",
	29: "PHZ_PID_ENDRESULT",
	30: "PHZ_PID_APPLYDISSOLVE_REQ",
	31: "PHZ_PID_APPLYDISSOLVE_ACK",
	32: "PHZ_PID_APPLYDISSOLVEBACK_REQ",
	33: "PHZ_PID_APPLYDISSOLVEBACK_ACK",
	34: "PHZ_PID_DISSOLVEDESK_REQ",
	35: "PHZ_PID_DISSOLVEDESK_ACK",
}
var PhzEnumProtoId_value = map[string]int32{
	"PHZ_PID_HEARTBEAT":             0,
	"PHZ_PID_CREATEDESK_REQ":        1,
	"PHZ_PID_ENTERDESK_REQ":         2,
	"PHZ_PID_DESK_ACK":              3,
	"PHZ_PID_GAMEINFO_REQ":          4,
	"PHZ_PID_GAMEINFO":              5,
	"PHZ_PID_READY_REQ":             6,
	"PHZ_PID_READY_ACK":             7,
	"PHZ_PID_OPENING":               8,
	"PHZ_PID_SENDCARDS":             9,
	"PHZ_PID_OUTCARD_REQ":           10,
	"PHZ_PID_OUTCARD_ACK":           11,
	"PHZ_PID_MOPAI":                 12,
	"PHZ_PID_OVERTURN":              13,
	"PHZ_PID_CANPENG":               14,
	"PHZ_PID_PENG_REQ":              15,
	"PHZ_PID_PENG_ACK":              16,
	"PHZ_PID_CHI_REQ":               17,
	"PHZ_PID_CANCHI":                18,
	"PHZ_PID_CHI_ACK":               19,
	"PHZ_PID_TIPAI":                 20,
	"PHZ_PID_WEIPAI":                21,
	"PHZ_PID_PAOPAI":                22,
	"PHZ_PID_HUPAI_REQ":             23,
	"PHZ_PID_CANHUPAI":              24,
	"PHZ_PID_HUPAI_ACK":             25,
	"PHZ_PID_PASS_REQ":              26,
	"PHZ_PID_PASS_ACK":              27,
	"PHZ_PID_CURRENTRESULT":         28,
	"PHZ_PID_ENDRESULT":             29,
	"PHZ_PID_APPLYDISSOLVE_REQ":     30,
	"PHZ_PID_APPLYDISSOLVE_ACK":     31,
	"PHZ_PID_APPLYDISSOLVEBACK_REQ": 32,
	"PHZ_PID_APPLYDISSOLVEBACK_ACK": 33,
	"PHZ_PID_DISSOLVEDESK_REQ":      34,
	"PHZ_PID_DISSOLVEDESK_ACK":      35,
}

func (x PhzEnumProtoId) Enum() *PhzEnumProtoId {
	p := new(PhzEnumProtoId)
	*p = x
	return p
}
func (x PhzEnumProtoId) String() string {
	return proto.EnumName(PhzEnumProtoId_name, int32(x))
}
func (x *PhzEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumProtoId_value, data, "PhzEnumProtoId")
	if err != nil {
		return err
	}
	*x = PhzEnumProtoId(value)
	return nil
}
func (PhzEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{0} }

type PhzEnumRoomType int32

const (
	PhzEnumRoomType_PHZ_ROOMTYPE_FRIEND PhzEnumRoomType = 1
)

var PhzEnumRoomType_name = map[int32]string{
	1: "PHZ_ROOMTYPE_FRIEND",
}
var PhzEnumRoomType_value = map[string]int32{
	"PHZ_ROOMTYPE_FRIEND": 1,
}

func (x PhzEnumRoomType) Enum() *PhzEnumRoomType {
	p := new(PhzEnumRoomType)
	*p = x
	return p
}
func (x PhzEnumRoomType) String() string {
	return proto.EnumName(PhzEnumRoomType_name, int32(x))
}
func (x *PhzEnumRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumRoomType_value, data, "PhzEnumRoomType")
	if err != nil {
		return err
	}
	*x = PhzEnumRoomType(value)
	return nil
}
func (PhzEnumRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{1} }

type PhzEnumTipaiType int32

const (
	PhzEnumTipaiType_PHZ_TIPAITYPE_MING PhzEnumTipaiType = 1
	PhzEnumTipaiType_PHZ_TIPAITYPE_AN   PhzEnumTipaiType = 2
)

var PhzEnumTipaiType_name = map[int32]string{
	1: "PHZ_TIPAITYPE_MING",
	2: "PHZ_TIPAITYPE_AN",
}
var PhzEnumTipaiType_value = map[string]int32{
	"PHZ_TIPAITYPE_MING": 1,
	"PHZ_TIPAITYPE_AN":   2,
}

func (x PhzEnumTipaiType) Enum() *PhzEnumTipaiType {
	p := new(PhzEnumTipaiType)
	*p = x
	return p
}
func (x PhzEnumTipaiType) String() string {
	return proto.EnumName(PhzEnumTipaiType_name, int32(x))
}
func (x *PhzEnumTipaiType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumTipaiType_value, data, "PhzEnumTipaiType")
	if err != nil {
		return err
	}
	*x = PhzEnumTipaiType(value)
	return nil
}
func (PhzEnumTipaiType) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{2} }

type PhzEnumPengType int32

const (
	PhzEnumPengType_PHZ_PENGTYPE_WEI        PhzEnumPengType = 1
	PhzEnumPengType_PHZ_PENGTYPE_CHOUWEI    PhzEnumPengType = 2
	PhzEnumPengType_PHZ_PENGTYPE_NORMALPENG PhzEnumPengType = 3
)

var PhzEnumPengType_name = map[int32]string{
	1: "PHZ_PENGTYPE_WEI",
	2: "PHZ_PENGTYPE_CHOUWEI",
	3: "PHZ_PENGTYPE_NORMALPENG",
}
var PhzEnumPengType_value = map[string]int32{
	"PHZ_PENGTYPE_WEI":        1,
	"PHZ_PENGTYPE_CHOUWEI":    2,
	"PHZ_PENGTYPE_NORMALPENG": 3,
}

func (x PhzEnumPengType) Enum() *PhzEnumPengType {
	p := new(PhzEnumPengType)
	*p = x
	return p
}
func (x PhzEnumPengType) String() string {
	return proto.EnumName(PhzEnumPengType_name, int32(x))
}
func (x *PhzEnumPengType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumPengType_value, data, "PhzEnumPengType")
	if err != nil {
		return err
	}
	*x = PhzEnumPengType(value)
	return nil
}
func (PhzEnumPengType) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{3} }

type PhzBaseCreateOption struct {
	UserCount        *int32           `protobuf:"varint,1,opt,name=userCount" json:"userCount,omitempty"`
	RoomType         *PhzEnumRoomType `protobuf:"varint,2,opt,name=roomType,enum=ddproto.PhzEnumRoomType" json:"roomType,omitempty"`
	BoardsCout       *int32           `protobuf:"varint,3,opt,name=boardsCout" json:"boardsCout,omitempty"`
	HuXi             *int32           `protobuf:"varint,4,opt,name=huXi" json:"huXi,omitempty"`
	IsDaiKai         *bool            `protobuf:"varint,5,opt,name=isDaiKai" json:"isDaiKai,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PhzBaseCreateOption) Reset()                    { *m = PhzBaseCreateOption{} }
func (m *PhzBaseCreateOption) String() string            { return proto.CompactTextString(m) }
func (*PhzBaseCreateOption) ProtoMessage()               {}
func (*PhzBaseCreateOption) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{0} }

func (m *PhzBaseCreateOption) GetUserCount() int32 {
	if m != nil && m.UserCount != nil {
		return *m.UserCount
	}
	return 0
}

func (m *PhzBaseCreateOption) GetRoomType() PhzEnumRoomType {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return PhzEnumRoomType_PHZ_ROOMTYPE_FRIEND
}

func (m *PhzBaseCreateOption) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *PhzBaseCreateOption) GetHuXi() int32 {
	if m != nil && m.HuXi != nil {
		return *m.HuXi
	}
	return 0
}

func (m *PhzBaseCreateOption) GetIsDaiKai() bool {
	if m != nil && m.IsDaiKai != nil {
		return *m.IsDaiKai
	}
	return false
}

type PhzBaseRoomInfo struct {
	OwnerId          *uint32          `protobuf:"varint,1,opt,name=ownerId" json:"ownerId,omitempty"`
	Password         *string          `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	UserCount        *int32           `protobuf:"varint,3,opt,name=userCount" json:"userCount,omitempty"`
	BoardsCout       *int32           `protobuf:"varint,4,opt,name=boardsCout" json:"boardsCout,omitempty"`
	CurrentRound     *int32           `protobuf:"varint,5,opt,name=currentRound" json:"currentRound,omitempty"`
	RoomType         *PhzEnumRoomType `protobuf:"varint,6,opt,name=roomType,enum=ddproto.PhzEnumRoomType" json:"roomType,omitempty"`
	Huxi             *int32           `protobuf:"varint,7,opt,name=huxi" json:"huxi,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PhzBaseRoomInfo) Reset()                    { *m = PhzBaseRoomInfo{} }
func (m *PhzBaseRoomInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzBaseRoomInfo) ProtoMessage()               {}
func (*PhzBaseRoomInfo) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{1} }

func (m *PhzBaseRoomInfo) GetOwnerId() uint32 {
	if m != nil && m.OwnerId != nil {
		return *m.OwnerId
	}
	return 0
}

func (m *PhzBaseRoomInfo) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *PhzBaseRoomInfo) GetUserCount() int32 {
	if m != nil && m.UserCount != nil {
		return *m.UserCount
	}
	return 0
}

func (m *PhzBaseRoomInfo) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *PhzBaseRoomInfo) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *PhzBaseRoomInfo) GetRoomType() PhzEnumRoomType {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return PhzEnumRoomType_PHZ_ROOMTYPE_FRIEND
}

func (m *PhzBaseRoomInfo) GetHuxi() int32 {
	if m != nil && m.Huxi != nil {
		return *m.Huxi
	}
	return 0
}

type PhzBaseDeskInfo struct {
	RoomInfo         *PhzBaseRoomInfo `protobuf:"bytes,1,opt,name=roomInfo" json:"roomInfo,omitempty"`
	GameStatus       *int32           `protobuf:"varint,2,opt,name=gameStatus" json:"gameStatus,omitempty"`
	RemainPaiCount   *int32           `protobuf:"varint,3,opt,name=remainPaiCount" json:"remainPaiCount,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PhzBaseDeskInfo) Reset()                    { *m = PhzBaseDeskInfo{} }
func (m *PhzBaseDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzBaseDeskInfo) ProtoMessage()               {}
func (*PhzBaseDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{2} }

func (m *PhzBaseDeskInfo) GetRoomInfo() *PhzBaseRoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

func (m *PhzBaseDeskInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PhzBaseDeskInfo) GetRemainPaiCount() int32 {
	if m != nil && m.RemainPaiCount != nil {
		return *m.RemainPaiCount
	}
	return 0
}

type PhzBasePlayerInfo struct {
	UserId           *uint32     `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool       `protobuf:"varint,2,opt,name=isOwner" json:"isOwner,omitempty"`
	IsBanker         *bool       `protobuf:"varint,3,opt,name=isBanker" json:"isBanker,omitempty"`
	IsReady          *bool       `protobuf:"varint,4,opt,name=isReady" json:"isReady,omitempty"`
	IsLeave          *bool       `protobuf:"varint,5,opt,name=isLeave" json:"isLeave,omitempty"`
	WxInfo           *WeixinInfo `protobuf:"bytes,6,opt,name=wxInfo" json:"wxInfo,omitempty"`
	Score            *int64      `protobuf:"varint,7,opt,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PhzBasePlayerInfo) Reset()                    { *m = PhzBasePlayerInfo{} }
func (m *PhzBasePlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzBasePlayerInfo) ProtoMessage()               {}
func (*PhzBasePlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{3} }

func (m *PhzBasePlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PhzBasePlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PhzBasePlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *PhzBasePlayerInfo) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

func (m *PhzBasePlayerInfo) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *PhzBasePlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *PhzBasePlayerInfo) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func init() {
	proto.RegisterType((*PhzBaseCreateOption)(nil), "ddproto.phz_base_createOption")
	proto.RegisterType((*PhzBaseRoomInfo)(nil), "ddproto.phz_base_roomInfo")
	proto.RegisterType((*PhzBaseDeskInfo)(nil), "ddproto.phz_base_deskInfo")
	proto.RegisterType((*PhzBasePlayerInfo)(nil), "ddproto.phz_base_playerInfo")
	proto.RegisterEnum("ddproto.PhzEnumProtoId", PhzEnumProtoId_name, PhzEnumProtoId_value)
	proto.RegisterEnum("ddproto.PhzEnumRoomType", PhzEnumRoomType_name, PhzEnumRoomType_value)
	proto.RegisterEnum("ddproto.PhzEnumTipaiType", PhzEnumTipaiType_name, PhzEnumTipaiType_value)
	proto.RegisterEnum("ddproto.PhzEnumPengType", PhzEnumPengType_name, PhzEnumPengType_value)
}

func init() { proto.RegisterFile("phz_base.proto", fileDescriptor43) }

var fileDescriptor43 = []byte{
	// 871 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x52, 0xdb, 0x46,
	0x14, 0x8e, 0x30, 0x36, 0x66, 0x13, 0x9c, 0x65, 0xcd, 0x8f, 0x70, 0x20, 0x25, 0xee, 0x4c, 0x87,
	0xa1, 0x1d, 0x2e, 0x7a, 0xd1, 0xfb, 0x45, 0xde, 0x60, 0x0d, 0xb6, 0xa4, 0xae, 0xe4, 0x50, 0x7a,
	0xe3, 0x2a, 0xd6, 0x36, 0xd1, 0x24, 0x96, 0x3c, 0x92, 0x5c, 0xa0, 0xaf, 0xd1, 0xe7, 0xe9, 0x8b,
	0xf4, 0x45, 0x72, 0xdb, 0xd9, 0x23, 0xaf, 0xb2, 0xc2, 0x49, 0x67, 0x7a, 0xa7, 0xf3, 0x7d, 0xdf,
	0x9e, 0x9f, 0xef, 0x1c, 0x1b, 0x75, 0x16, 0xef, 0xff, 0x9c, 0xbe, 0x0d, 0x73, 0x71, 0xb1, 0xc8,
	0xd2, 0x22, 0x25, 0x5b, 0x51, 0x04, 0x1f, 0xbd, 0xee, 0x2c, 0x9d, 0xcf, 0xd3, 0x64, 0x3a, 0xfb,
	0x18, 0x8b, 0xa4, 0x28, 0xd9, 0xfe, 0xdf, 0x06, 0xda, 0x57, 0x0f, 0xa6, 0xb3, 0x4c, 0x84, 0x85,
	0x70, 0x17, 0x45, 0x9c, 0x26, 0xe4, 0x18, 0x6d, 0x2f, 0x73, 0x91, 0x59, 0xe9, 0x32, 0x29, 0x4c,
	0xe3, 0xd4, 0x38, 0x6b, 0xf2, 0xcf, 0x00, 0xf9, 0x09, 0xb5, 0xb3, 0x34, 0x9d, 0x07, 0x0f, 0x0b,
	0x61, 0x6e, 0x9c, 0x1a, 0x67, 0x9d, 0x1f, 0x7b, 0x17, 0xab, 0x42, 0x17, 0x32, 0x9f, 0x48, 0x96,
	0xf3, 0xa9, 0x52, 0xf0, 0x4a, 0x4b, 0x5e, 0x22, 0xf4, 0x36, 0x0d, 0xb3, 0x28, 0xb7, 0xd2, 0x65,
	0x61, 0x36, 0x20, 0xad, 0x86, 0x10, 0x82, 0x36, 0xdf, 0x2f, 0x7f, 0x89, 0xcd, 0x4d, 0x60, 0xe0,
	0x9b, 0xf4, 0x50, 0x3b, 0xce, 0x07, 0x61, 0x7c, 0x1d, 0xc6, 0x66, 0xf3, 0xd4, 0x38, 0x6b, 0xf3,
	0x2a, 0xee, 0x7f, 0x32, 0xd0, 0x6e, 0xd5, 0xbf, 0xac, 0x62, 0x27, 0xbf, 0xa7, 0xc4, 0x44, 0x5b,
	0xe9, 0x5d, 0x22, 0x32, 0x3b, 0x82, 0xce, 0x77, 0xb8, 0x0a, 0x65, 0xae, 0x45, 0x98, 0xe7, 0x77,
	0x69, 0x16, 0x41, 0xdf, 0xdb, 0xbc, 0x8a, 0xeb, 0x13, 0x37, 0x1e, 0x4f, 0x5c, 0xef, 0x7c, 0x73,
	0xad, 0xf3, 0x3e, 0x7a, 0x36, 0x5b, 0x66, 0x99, 0x48, 0x0a, 0x9e, 0x2e, 0x93, 0x08, 0x3a, 0x6d,
	0xf2, 0x1a, 0x56, 0x73, 0xad, 0xf5, 0x3f, 0x5c, 0x03, 0x57, 0xee, 0x63, 0x73, 0x4b, 0xb9, 0x72,
	0x1f, 0xf7, 0xff, 0xd2, 0x27, 0x8f, 0x44, 0xfe, 0x01, 0x26, 0x5f, 0x55, 0x90, 0xdf, 0x30, 0xfa,
	0xd3, 0x47, 0x15, 0x6a, 0x3e, 0xf1, 0x4a, 0x2b, 0xa7, 0x7b, 0x17, 0xce, 0x85, 0x5f, 0x84, 0xc5,
	0x32, 0x07, 0x67, 0x9a, 0x5c, 0x43, 0xc8, 0x77, 0xa8, 0x93, 0x89, 0x79, 0x18, 0x27, 0x5e, 0x18,
	0xeb, 0x06, 0x3d, 0x42, 0xfb, 0xff, 0x18, 0xa8, 0x5b, 0xd5, 0x59, 0x7c, 0x0c, 0x1f, 0x44, 0x06,
	0xf9, 0x0f, 0x50, 0x4b, 0x5a, 0x59, 0x2d, 0x64, 0x15, 0xc9, 0x4d, 0xc5, 0xb9, 0x2b, 0x97, 0x03,
	0x45, 0xdb, 0x5c, 0x85, 0xe5, 0xd6, 0x2f, 0xc3, 0xe4, 0x83, 0xc8, 0xa0, 0x16, 0x6c, 0xbd, 0x8c,
	0xcb, 0x57, 0x5c, 0x84, 0xd1, 0x03, 0x2c, 0x02, 0x5e, 0x41, 0x58, 0x32, 0x23, 0x11, 0xfe, 0x21,
	0x56, 0xa7, 0xa2, 0x42, 0xf2, 0x3d, 0x6a, 0xdd, 0xdd, 0x83, 0x2f, 0x2d, 0xf0, 0xa5, 0x5b, 0xf9,
	0x72, 0x23, 0xe2, 0xfb, 0x38, 0x01, 0x43, 0x56, 0x12, 0xb2, 0x87, 0x9a, 0xf9, 0x2c, 0xcd, 0x04,
	0x38, 0xde, 0xe0, 0x65, 0x70, 0xfe, 0xa9, 0x85, 0x70, 0xb5, 0x26, 0x78, 0x6b, 0x47, 0x64, 0x1f,
	0xed, 0x7a, 0xc3, 0x5f, 0xa7, 0x9e, 0x3d, 0x98, 0x0e, 0x19, 0xe5, 0xc1, 0x25, 0xa3, 0x01, 0x7e,
	0x42, 0x7a, 0xe8, 0x40, 0xc1, 0x16, 0x67, 0x34, 0x60, 0x03, 0xe6, 0x5f, 0x4f, 0x39, 0xfb, 0x19,
	0x1b, 0xe4, 0x08, 0xed, 0x2b, 0x8e, 0x39, 0x01, 0xe3, 0x15, 0xb5, 0x41, 0xf6, 0x10, 0x56, 0x14,
	0xa0, 0xd4, 0xba, 0xc6, 0x0d, 0x62, 0xa2, 0x3d, 0x85, 0x5e, 0xd1, 0x31, 0xb3, 0x9d, 0xd7, 0x2e,
	0xe8, 0x37, 0x75, 0xbd, 0x62, 0x70, 0x53, 0xef, 0x89, 0x33, 0x3a, 0xb8, 0x05, 0x71, 0x6b, 0x1d,
	0x96, 0xd9, 0xb7, 0x48, 0x17, 0x3d, 0x57, 0xb0, 0xeb, 0x31, 0xc7, 0x76, 0xae, 0x70, 0x5b, 0xd7,
	0xfa, 0xcc, 0x19, 0x58, 0x94, 0x0f, 0x7c, 0xbc, 0x4d, 0x0e, 0x51, 0xb7, 0xd2, 0x4e, 0x02, 0x89,
	0x42, 0x6e, 0xf4, 0x25, 0x42, 0x66, 0x7f, 0x4a, 0x76, 0xd1, 0x8e, 0x22, 0xc6, 0xae, 0x47, 0x6d,
	0xfc, 0x4c, 0x6f, 0xda, 0x7d, 0xc3, 0x78, 0x30, 0xe1, 0x0e, 0xde, 0xd1, 0xdb, 0xb0, 0xa8, 0xe3,
	0x31, 0xe7, 0x0a, 0x77, 0x74, 0xa9, 0x44, 0xa0, 0xd8, 0xf3, 0x35, 0x54, 0x56, 0xc2, 0xb5, 0x04,
	0x43, 0x1b, 0xa4, 0xbb, 0x84, 0xa0, 0x8e, 0x96, 0xd5, 0x1a, 0xda, 0x98, 0x3c, 0x16, 0xca, 0xd7,
	0x5d, 0xbd, 0xcf, 0xc0, 0x96, 0x7d, 0xee, 0xe9, 0x6f, 0x6f, 0x18, 0x60, 0xfb, 0x3a, 0xe6, 0x51,
	0x98, 0xe7, 0xa0, 0x76, 0x02, 0x13, 0x8f, 0x96, 0xa5, 0x0f, 0xf5, 0x2e, 0x2d, 0xea, 0x00, 0x83,
	0xcd, 0x75, 0xb1, 0x2c, 0x7f, 0x54, 0x1b, 0x89, 0xfa, 0x3e, 0xa4, 0xe8, 0xad, 0xa1, 0x52, 0xfb,
	0x42, 0xbf, 0x1f, 0x6b, 0xc2, 0x39, 0x73, 0x02, 0xce, 0xfc, 0xc9, 0x28, 0xc0, 0xc7, 0x7a, 0x76,
	0xe6, 0x0c, 0x56, 0xf0, 0x09, 0x39, 0x41, 0x47, 0x0a, 0xa6, 0x9e, 0x37, 0xba, 0x1d, 0xd8, 0xbe,
	0xef, 0x8e, 0xde, 0x30, 0x28, 0xf3, 0xf2, 0xeb, 0xb4, 0xac, 0xf7, 0x0d, 0x79, 0x85, 0x4e, 0xbe,
	0x48, 0x5f, 0x52, 0xab, 0xbc, 0xdb, 0xd3, 0xff, 0x96, 0xc8, 0x2c, 0xaf, 0xc8, 0x31, 0x32, 0xab,
	0xd3, 0x5e, 0xb1, 0xd5, 0xe1, 0xf7, 0xbf, 0xca, 0xca, 0xb7, 0xdf, 0x9e, 0xff, 0x50, 0xfe, 0xd7,
	0xd5, 0xfe, 0x1f, 0xd5, 0xc9, 0x71, 0xd7, 0x1d, 0x07, 0xb7, 0x1e, 0x9b, 0xbe, 0xe6, 0x36, 0x73,
	0x06, 0xd8, 0x38, 0xbf, 0x44, 0xa4, 0x52, 0x17, 0xf1, 0x22, 0x8c, 0x41, 0x7e, 0x80, 0x88, 0x94,
	0xc3, 0x72, 0x41, 0x3f, 0x96, 0x97, 0x6e, 0x28, 0x8f, 0x3f, 0xe3, 0xd4, 0xc1, 0x1b, 0xe7, 0xbf,
	0x69, 0x15, 0x17, 0x22, 0x79, 0x07, 0x29, 0xd4, 0x3a, 0x98, 0x73, 0x05, 0xca, 0x1b, 0x66, 0x63,
	0xa3, 0xfa, 0x75, 0x2a, 0xd4, 0x1a, 0xba, 0x13, 0xc9, 0x6c, 0x90, 0x17, 0xe8, 0xb0, 0xc6, 0x38,
	0x2e, 0x1f, 0xd3, 0x11, 0x9c, 0x76, 0xc3, 0x7b, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25,
	0xb0, 0x49, 0x60, 0xab, 0x07, 0x00, 0x00,
}
