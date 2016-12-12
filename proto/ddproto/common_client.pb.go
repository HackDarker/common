// Code generated by protoc-gen-go.
// source: common_client.proto
// DO NOT EDIT!

/*
Package ddproto is a generated protocol buffer package.

It is generated from these files:
	common_client.proto
	common_server.proto
	common_server_poker.proto
	ddz_base.proto
	ddz_desk.proto
	ddz_hall.proto
	ddz_play.proto
	ddz_server.proto
	mahjong_desk.proto
	mahjong_server.proto
	zjh_base.proto
	zjh_desk.proto
	zjh_hall.proto
	zjh_play.proto

It has these top-level messages:
	ProtoHeader
	Heartbeat
	ServerInfo
	QuickConn
	AckQuickConn
	WeixinInfo
	CommonReqGameLogin
	CommonAckGameLogin
	CommonReqLogout
	CommonAckLogout
	CommonReqFeedback
	ClientBasePoker
	User
	TNotice
	GameSession
	CommonSrvGameUser
	CommonSrvPokerPai
	DdzBaseRoomTypeInfo
	DdzBasePlayerInfo
	DdzBaseDeskInfo
	DdzReqDissolveDesk
	DdzAckDissolveDesk
	DdzReqLeaveDesk
	DdzAckLeaveDesk
	DdzReqReady
	DdzAckReady
	DdzReqMessage
	DdzBcMessage
	DdzBaseWinCoinInfo
	DdzBaCurrentResult
	DdzBaseEndLotteryInfo
	DdzBcEndLottery
	DdzReqCreateDesk
	DdzAckCreateDesk
	DdzReqGameRecord
	DdzBaseUserRecord
	DdzBaseGameRecord
	DdzAckGameRecord
	DdzReqEnterDesk
	DdzAckEnterDesk
	DdzBcOpening
	DdzBcDealCards
	DdzReqShowHandPokers
	DdzAckShowHandPokers
	DdzReqJiaoDiZhu
	DdzAckJiaoDiZhu
	DdzReqRobDiZhu
	DdzAckRobDiZhu
	DdzReqDouble
	DdzAckDouble
	DdzBcStartPlay
	DdzReqMenuZhua
	DdzAckMenuZhua
	DdzReqSeeCards
	DdzAckSeeCards
	DdzReqPull
	DdzAckPull
	DdzReqOutCards
	DdzAckOutCards
	DdzReqActGuo
	DdzAckGuoAck
	DdzBcOverTurn
	DdzBcGameInfo
	DdzSrvOutPokerPais
	DdzSrvDeskTongJi
	DdzSrvDesk
	DdzSrvGameData
	DdzSrvBillBean
	DdzSrvBill
	DdzSrvUserStatistics
	DdzSrvUser
	DdzSrvRoom
	DdzSrvBak
	MjBasePlayOptions
	MjBaseCardInfo
	MjBaseRoomTypeInfo
	MjBaseComposeCard
	MjBasePlayerCard
	MjBasePlayerInfo
	MjBaseDeskGameInfo
	MjReqDissolveDesk
	MjAckDissolveDesk
	ZjhReqGetRoomList
	ZjhBaseRoomInfo
	ZjhAckRoomList
	ZjhReqEnterDesk
	ZjhBaseUserInfo
	ZjhDeskStateAck
	ZjhBcGameInfo
	ZjhBaseDeskInfo
	ZjhBasePlayerInfo
*/
package ddproto

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

// -------------------------------------------扑克牌相关的------------------------------------
// 房间类型枚举
// 扑克花色
type CommonEnumPokerColor int32

const (
	CommonEnumPokerColor_HONGTAO       CommonEnumPokerColor = 1
	CommonEnumPokerColor_FANGKUAI      CommonEnumPokerColor = 2
	CommonEnumPokerColor_HEITAO        CommonEnumPokerColor = 3
	CommonEnumPokerColor_MEIHUA        CommonEnumPokerColor = 4
	CommonEnumPokerColor_REDJOKER      CommonEnumPokerColor = 5
	CommonEnumPokerColor_BLACKBIGJOKER CommonEnumPokerColor = 6
)

var CommonEnumPokerColor_name = map[int32]string{
	1: "HONGTAO",
	2: "FANGKUAI",
	3: "HEITAO",
	4: "MEIHUA",
	5: "REDJOKER",
	6: "BLACKBIGJOKER",
}
var CommonEnumPokerColor_value = map[string]int32{
	"HONGTAO":       1,
	"FANGKUAI":      2,
	"HEITAO":        3,
	"MEIHUA":        4,
	"REDJOKER":      5,
	"BLACKBIGJOKER": 6,
}

func (x CommonEnumPokerColor) Enum() *CommonEnumPokerColor {
	p := new(CommonEnumPokerColor)
	*p = x
	return p
}
func (x CommonEnumPokerColor) String() string {
	return proto.EnumName(CommonEnumPokerColor_name, int32(x))
}
func (x *CommonEnumPokerColor) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommonEnumPokerColor_value, data, "CommonEnumPokerColor")
	if err != nil {
		return err
	}
	*x = CommonEnumPokerColor(value)
	return nil
}
func (CommonEnumPokerColor) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// ProtoHeader 需要在每个 Message 中作为第一个字段
type ProtoHeader struct {
	Version          *string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Code             *int32  `protobuf:"varint,3,opt,name=code" json:"code,omitempty"`
	Error            *string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoHeader) Reset()                    { *m = ProtoHeader{} }
func (m *ProtoHeader) String() string            { return proto.CompactTextString(m) }
func (*ProtoHeader) ProtoMessage()               {}
func (*ProtoHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoHeader) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *ProtoHeader) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ProtoHeader) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *ProtoHeader) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

// 这里都是通用过的一些协议
type Heartbeat struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Heartbeat) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 服务器信息
type ServerInfo struct {
	Ip               *string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port             *int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Status           *int32  `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServerInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *ServerInfo) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ServerInfo) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

// 接入服务器
type QuickConn struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ChannelId        *string      `protobuf:"bytes,2,opt,name=channelId" json:"channelId,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	CurrVersion      *int32       `protobuf:"varint,4,opt,name=currVersion" json:"currVersion,omitempty"`
	LanguageId       *int32       `protobuf:"varint,5,opt,name=languageId" json:"languageId,omitempty"`
	UserId           *uint32      `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *QuickConn) Reset()                    { *m = QuickConn{} }
func (m *QuickConn) String() string            { return proto.CompactTextString(m) }
func (*QuickConn) ProtoMessage()               {}
func (*QuickConn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QuickConn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *QuickConn) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

func (m *QuickConn) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *QuickConn) GetCurrVersion() int32 {
	if m != nil && m.CurrVersion != nil {
		return *m.CurrVersion
	}
	return 0
}

func (m *QuickConn) GetLanguageId() int32 {
	if m != nil && m.LanguageId != nil {
		return *m.LanguageId
	}
	return 0
}

func (m *QuickConn) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type AckQuickConn struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	GameServer       *ServerInfo  `protobuf:"bytes,2,opt,name=gameServer" json:"gameServer,omitempty"`
	ReleaseTag       *int32       `protobuf:"varint,3,opt,name=releaseTag" json:"releaseTag,omitempty"`
	CurrVersion      *int32       `protobuf:"varint,4,opt,name=currVersion" json:"currVersion,omitempty"`
	IsUpdate         *int32       `protobuf:"varint,5,opt,name=isUpdate" json:"isUpdate,omitempty"`
	DownloadUrl      *string      `protobuf:"bytes,6,opt,name=downloadUrl" json:"downloadUrl,omitempty"`
	VersionInfo      *string      `protobuf:"bytes,7,opt,name=versionInfo" json:"versionInfo,omitempty"`
	IsMaintain       *int32       `protobuf:"varint,8,opt,name=isMaintain" json:"isMaintain,omitempty"`
	MaintainMsg      *string      `protobuf:"bytes,9,opt,name=maintainMsg" json:"maintainMsg,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AckQuickConn) Reset()                    { *m = AckQuickConn{} }
func (m *AckQuickConn) String() string            { return proto.CompactTextString(m) }
func (*AckQuickConn) ProtoMessage()               {}
func (*AckQuickConn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AckQuickConn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AckQuickConn) GetGameServer() *ServerInfo {
	if m != nil {
		return m.GameServer
	}
	return nil
}

func (m *AckQuickConn) GetReleaseTag() int32 {
	if m != nil && m.ReleaseTag != nil {
		return *m.ReleaseTag
	}
	return 0
}

func (m *AckQuickConn) GetCurrVersion() int32 {
	if m != nil && m.CurrVersion != nil {
		return *m.CurrVersion
	}
	return 0
}

func (m *AckQuickConn) GetIsUpdate() int32 {
	if m != nil && m.IsUpdate != nil {
		return *m.IsUpdate
	}
	return 0
}

func (m *AckQuickConn) GetDownloadUrl() string {
	if m != nil && m.DownloadUrl != nil {
		return *m.DownloadUrl
	}
	return ""
}

func (m *AckQuickConn) GetVersionInfo() string {
	if m != nil && m.VersionInfo != nil {
		return *m.VersionInfo
	}
	return ""
}

func (m *AckQuickConn) GetIsMaintain() int32 {
	if m != nil && m.IsMaintain != nil {
		return *m.IsMaintain
	}
	return 0
}

func (m *AckQuickConn) GetMaintainMsg() string {
	if m != nil && m.MaintainMsg != nil {
		return *m.MaintainMsg
	}
	return ""
}

// 微信信息
type WeixinInfo struct {
	OpenId           *string `protobuf:"bytes,1,opt,name=openId" json:"openId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	HeadUrl          *string `protobuf:"bytes,3,opt,name=headUrl" json:"headUrl,omitempty"`
	Sex              *int32  `protobuf:"varint,4,opt,name=sex" json:"sex,omitempty"`
	City             *string `protobuf:"bytes,5,opt,name=city" json:"city,omitempty"`
	UnionId          *string `protobuf:"bytes,6,opt,name=unionId" json:"unionId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WeixinInfo) Reset()                    { *m = WeixinInfo{} }
func (m *WeixinInfo) String() string            { return proto.CompactTextString(m) }
func (*WeixinInfo) ProtoMessage()               {}
func (*WeixinInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WeixinInfo) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *WeixinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *WeixinInfo) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *WeixinInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *WeixinInfo) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *WeixinInfo) GetUnionId() string {
	if m != nil && m.UnionId != nil {
		return *m.UnionId
	}
	return ""
}

// 登录的结构
type CommonReqGameLogin struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	ProtoVersion     *int32       `protobuf:"varint,3,opt,name=protoVersion" json:"protoVersion,omitempty"`
	WxInfo           *WeixinInfo  `protobuf:"bytes,4,opt,name=wxInfo" json:"wxInfo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonReqGameLogin) Reset()                    { *m = CommonReqGameLogin{} }
func (m *CommonReqGameLogin) String() string            { return proto.CompactTextString(m) }
func (*CommonReqGameLogin) ProtoMessage()               {}
func (*CommonReqGameLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CommonReqGameLogin) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommonReqGameLogin) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *CommonReqGameLogin) GetProtoVersion() int32 {
	if m != nil && m.ProtoVersion != nil {
		return *m.ProtoVersion
	}
	return 0
}

func (m *CommonReqGameLogin) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

// 游戏登录回复
type CommonAckGameLogin struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	NickName         *string      `protobuf:"bytes,3,opt,name=nickName" json:"nickName,omitempty"`
	RoomPassword     *string      `protobuf:"bytes,4,opt,name=roomPassword" json:"roomPassword,omitempty"`
	CostCreateRoom   *int64       `protobuf:"varint,5,opt,name=costCreateRoom" json:"costCreateRoom,omitempty"`
	CostRebuy        *int64       `protobuf:"varint,6,opt,name=costRebuy" json:"costRebuy,omitempty"`
	Championship     *bool        `protobuf:"varint,7,opt,name=championship" json:"championship,omitempty"`
	Chip             *int64       `protobuf:"varint,8,opt,name=chip" json:"chip,omitempty"`
	MailCount        *int32       `protobuf:"varint,9,opt,name=mailCount" json:"mailCount,omitempty"`
	Notice           *string      `protobuf:"bytes,10,opt,name=notice" json:"notice,omitempty"`
	GameStatus       *int32       `protobuf:"varint,11,opt,name=gameStatus" json:"gameStatus,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonAckGameLogin) Reset()                    { *m = CommonAckGameLogin{} }
func (m *CommonAckGameLogin) String() string            { return proto.CompactTextString(m) }
func (*CommonAckGameLogin) ProtoMessage()               {}
func (*CommonAckGameLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CommonAckGameLogin) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommonAckGameLogin) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *CommonAckGameLogin) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *CommonAckGameLogin) GetRoomPassword() string {
	if m != nil && m.RoomPassword != nil {
		return *m.RoomPassword
	}
	return ""
}

func (m *CommonAckGameLogin) GetCostCreateRoom() int64 {
	if m != nil && m.CostCreateRoom != nil {
		return *m.CostCreateRoom
	}
	return 0
}

func (m *CommonAckGameLogin) GetCostRebuy() int64 {
	if m != nil && m.CostRebuy != nil {
		return *m.CostRebuy
	}
	return 0
}

func (m *CommonAckGameLogin) GetChampionship() bool {
	if m != nil && m.Championship != nil {
		return *m.Championship
	}
	return false
}

func (m *CommonAckGameLogin) GetChip() int64 {
	if m != nil && m.Chip != nil {
		return *m.Chip
	}
	return 0
}

func (m *CommonAckGameLogin) GetMailCount() int32 {
	if m != nil && m.MailCount != nil {
		return *m.MailCount
	}
	return 0
}

func (m *CommonAckGameLogin) GetNotice() string {
	if m != nil && m.Notice != nil {
		return *m.Notice
	}
	return ""
}

func (m *CommonAckGameLogin) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

// 客户端请求
type CommonReqLogout struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonReqLogout) Reset()                    { *m = CommonReqLogout{} }
func (m *CommonReqLogout) String() string            { return proto.CompactTextString(m) }
func (*CommonReqLogout) ProtoMessage()               {}
func (*CommonReqLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CommonReqLogout) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommonReqLogout) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 服务器推
type CommonAckLogout struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonAckLogout) Reset()                    { *m = CommonAckLogout{} }
func (m *CommonAckLogout) String() string            { return proto.CompactTextString(m) }
func (*CommonAckLogout) ProtoMessage()               {}
func (*CommonAckLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CommonAckLogout) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommonAckLogout) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 反馈信息的协议
type CommonReqFeedback struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Message          *string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonReqFeedback) Reset()                    { *m = CommonReqFeedback{} }
func (m *CommonReqFeedback) String() string            { return proto.CompactTextString(m) }
func (*CommonReqFeedback) ProtoMessage()               {}
func (*CommonReqFeedback) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CommonReqFeedback) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommonReqFeedback) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ClientBasePoker struct {
	Suit             *CommonEnumPokerColor `protobuf:"varint,1,opt,name=suit,enum=ddproto.CommonEnumPokerColor" json:"suit,omitempty"`
	Num              *int32                `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
	Id               *int32                `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *ClientBasePoker) Reset()                    { *m = ClientBasePoker{} }
func (m *ClientBasePoker) String() string            { return proto.CompactTextString(m) }
func (*ClientBasePoker) ProtoMessage()               {}
func (*ClientBasePoker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ClientBasePoker) GetSuit() CommonEnumPokerColor {
	if m != nil && m.Suit != nil {
		return *m.Suit
	}
	return CommonEnumPokerColor_HONGTAO
}

func (m *ClientBasePoker) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *ClientBasePoker) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtoHeader)(nil), "ddproto.ProtoHeader")
	proto.RegisterType((*Heartbeat)(nil), "ddproto.Heartbeat")
	proto.RegisterType((*ServerInfo)(nil), "ddproto.ServerInfo")
	proto.RegisterType((*QuickConn)(nil), "ddproto.QuickConn")
	proto.RegisterType((*AckQuickConn)(nil), "ddproto.AckQuickConn")
	proto.RegisterType((*WeixinInfo)(nil), "ddproto.WeixinInfo")
	proto.RegisterType((*CommonReqGameLogin)(nil), "ddproto.common_req_gameLogin")
	proto.RegisterType((*CommonAckGameLogin)(nil), "ddproto.common_ack_gameLogin")
	proto.RegisterType((*CommonReqLogout)(nil), "ddproto.common_req_logout")
	proto.RegisterType((*CommonAckLogout)(nil), "ddproto.common_ack_logout")
	proto.RegisterType((*CommonReqFeedback)(nil), "ddproto.common_req_feedback")
	proto.RegisterType((*ClientBasePoker)(nil), "ddproto.client_base_poker")
	proto.RegisterEnum("ddproto.CommonEnumPokerColor", CommonEnumPokerColor_name, CommonEnumPokerColor_value)
}

var fileDescriptor0 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x6e, 0xdb, 0x3a,
	0x10, 0x85, 0xe3, 0xa7, 0xc6, 0x49, 0xae, 0xa3, 0x04, 0x81, 0x76, 0xf7, 0x42, 0xb7, 0x40, 0x8b,
	0x02, 0x0d, 0xd0, 0x2c, 0xba, 0x77, 0x5c, 0x37, 0x56, 0x9d, 0x57, 0xdd, 0xb8, 0x5d, 0x74, 0x61,
	0xd0, 0x12, 0xa3, 0x10, 0x96, 0x48, 0x95, 0xa4, 0x9a, 0xe4, 0x0b, 0xba, 0xea, 0x97, 0xf6, 0x0f,
	0xba, 0xea, 0x90, 0xa2, 0x13, 0x23, 0x40, 0x8b, 0xb4, 0x59, 0x18, 0x10, 0xc7, 0xc3, 0x39, 0x67,
	0xce, 0x9c, 0x21, 0x6c, 0xc7, 0x22, 0xcf, 0x05, 0x9f, 0xc5, 0x19, 0xa3, 0x5c, 0xef, 0x15, 0x52,
	0x68, 0xe1, 0xb7, 0x93, 0xc4, 0x7e, 0x84, 0x63, 0xe8, 0x9e, 0x99, 0x8f, 0x11, 0x25, 0x09, 0x95,
	0xfe, 0x3f, 0xd0, 0xfe, 0x42, 0xa5, 0x62, 0x82, 0x07, 0xb5, 0xff, 0x6a, 0xcf, 0x3c, 0x7f, 0x13,
	0x5a, 0xa5, 0xa2, 0x32, 0x4a, 0x82, 0x35, 0x3c, 0x6f, 0xf8, 0xeb, 0xd0, 0x88, 0x45, 0x42, 0x83,
	0x3a, 0x9e, 0x9a, 0xfe, 0x06, 0x34, 0xa9, 0x94, 0x42, 0x06, 0x0d, 0x93, 0x1c, 0xbe, 0x04, 0x0f,
	0xeb, 0x48, 0x3d, 0xa7, 0x44, 0xfb, 0x4f, 0xa0, 0x75, 0x69, 0x8b, 0xda, 0x4a, 0xdd, 0xfd, 0x9d,
	0x3d, 0x87, 0xb9, 0xb7, 0x02, 0x18, 0xbe, 0x02, 0x78, 0x4f, 0x25, 0x62, 0x46, 0xfc, 0x42, 0xf8,
	0x00, 0x6b, 0xac, 0x70, 0xc8, 0x88, 0x54, 0x08, 0xa9, 0x2d, 0x6e, 0xd3, 0xf0, 0x50, 0x9a, 0xe8,
	0x52, 0x55, 0xc8, 0xe1, 0xb7, 0x1a, 0x78, 0xef, 0x4a, 0x16, 0x2f, 0x06, 0x82, 0xf3, 0x87, 0x61,
	0xf9, 0x5b, 0xe0, 0xc5, 0x97, 0x84, 0x73, 0x9a, 0xb9, 0x76, 0x6c, 0x7b, 0x29, 0xc9, 0x29, 0x9e,
	0xab, 0x86, 0xb6, 0xa1, 0x1b, 0x97, 0x52, 0x7e, 0x70, 0x1a, 0x34, 0x6c, 0x10, 0x69, 0x65, 0x84,
	0xa7, 0x25, 0x49, 0x4d, 0x62, 0x73, 0xc9, 0xc7, 0xe9, 0xd2, 0x32, 0xba, 0x84, 0xdf, 0x6b, 0xb0,
	0xde, 0x8f, 0x17, 0x7f, 0x4a, 0xe9, 0x29, 0x80, 0xc1, 0xaf, 0x24, 0xb0, 0x9c, 0xba, 0xfb, 0xdb,
	0xb7, 0x99, 0x2b, 0xca, 0x20, 0x07, 0x49, 0x33, 0x4a, 0x14, 0x3d, 0x27, 0xe9, 0xef, 0xc8, 0xf6,
	0xa0, 0xc3, 0xd4, 0xb4, 0x48, 0x88, 0xa6, 0x8e, 0x2a, 0xa6, 0x25, 0xe2, 0x8a, 0x67, 0x82, 0x24,
	0x53, 0x99, 0x59, 0xbe, 0x9e, 0x09, 0xba, 0x41, 0x9b, 0xf2, 0x41, 0xdb, 0x06, 0x11, 0x84, 0xa9,
	0x63, 0xc2, 0xb8, 0xc6, 0x5f, 0xd0, 0x59, 0xde, 0xce, 0x5d, 0xe4, 0x58, 0xa5, 0x81, 0x67, 0x07,
	0xbd, 0x00, 0xf8, 0x48, 0xd9, 0x35, 0xb3, 0x97, 0x8d, 0x16, 0xa2, 0xa0, 0x1c, 0xb5, 0xa8, 0x26,
	0x87, 0x14, 0x38, 0xca, 0x70, 0x82, 0x8d, 0x39, 0x99, 0xd1, 0x56, 0x46, 0x0c, 0x03, 0x5f, 0xb7,
	0x81, 0x2e, 0xd4, 0x15, 0xbd, 0x76, 0x94, 0x8d, 0xa7, 0x98, 0xbe, 0xb1, 0x74, 0x6d, 0x6e, 0xc9,
	0x0d, 0xaf, 0x4a, 0x5a, 0x2f, 0xfc, 0x5a, 0x83, 0x1d, 0xe7, 0x61, 0x49, 0x3f, 0xcf, 0x8c, 0x5e,
	0x47, 0x22, 0x65, 0x0f, 0x95, 0xf8, 0xbe, 0x83, 0x77, 0x60, 0xdd, 0x26, 0x2d, 0x65, 0xab, 0xb4,
	0xfc, 0x1f, 0x5a, 0x57, 0xd7, 0x56, 0x8a, 0xc6, 0xbd, 0x21, 0xdc, 0x35, 0x1a, 0xfe, 0xb8, 0x63,
	0x42, 0xe2, 0xc5, 0xa3, 0x99, 0xac, 0xea, 0x54, 0xc9, 0x82, 0xdc, 0xa4, 0x10, 0xf9, 0x19, 0x51,
	0xea, 0x4a, 0xc8, 0xa4, 0x5a, 0x2b, 0x7f, 0x17, 0x36, 0x63, 0xa1, 0xf4, 0x40, 0xe2, 0x5a, 0xd1,
	0x09, 0xfe, 0x6f, 0x95, 0xaa, 0x5b, 0x3f, 0x63, 0x7c, 0x42, 0xe7, 0xe5, 0x8d, 0xd5, 0xaa, 0x6e,
	0x0a, 0xa0, 0xc5, 0xf3, 0x02, 0x1b, 0x53, 0x97, 0xb8, 0x4a, 0x66, 0xae, 0x1d, 0x2b, 0xb0, 0x39,
	0x75, 0x96, 0xd7, 0x70, 0xa2, 0xd9, 0x40, 0x94, 0x5c, 0xdb, 0x79, 0x5a, 0x37, 0x73, 0xa1, 0x59,
	0x4c, 0x03, 0x58, 0x1a, 0xc1, 0xda, 0xb2, 0xda, 0xb8, 0xae, 0xdd, 0xb8, 0x08, 0xb6, 0x56, 0xa6,
	0x90, 0x89, 0x54, 0x94, 0xfa, 0xef, 0x1a, 0x5f, 0x29, 0x65, 0x64, 0x7c, 0x54, 0xa9, 0xa3, 0xdb,
	0xf7, 0xcd, 0xb0, 0xba, 0xa0, 0x34, 0x99, 0x63, 0xcd, 0x07, 0x16, 0x43, 0xab, 0xe5, 0x54, 0x29,
	0xdc, 0xeb, 0xca, 0xa7, 0xe1, 0x27, 0x24, 0x66, 0x9f, 0xc9, 0xd9, 0x1c, 0x37, 0x6d, 0x56, 0x88,
	0x05, 0x66, 0xbd, 0x80, 0x86, 0x2a, 0x99, 0xb6, 0x95, 0x36, 0xf7, 0xff, 0xbd, 0xad, 0xe4, 0x70,
	0x29, 0x2f, 0xf3, 0x2a, 0x73, 0x20, 0x32, 0x21, 0x8d, 0xb5, 0x31, 0xe2, 0x9e, 0x2d, 0xf3, 0xa0,
	0xb9, 0xb7, 0xe5, 0x79, 0x06, 0xbb, 0xbf, 0xbc, 0xd2, 0x1e, 0x9d, 0x9e, 0x1c, 0x9e, 0xf7, 0x4f,
	0x7b, 0x35, 0x1c, 0x56, 0xe7, 0x4d, 0xff, 0xe4, 0x70, 0x3c, 0xed, 0x47, 0xbd, 0x35, 0x2c, 0xd0,
	0x1a, 0x0d, 0x23, 0xf3, 0x4f, 0xdd, 0x7c, 0x1f, 0x0f, 0xa3, 0xd1, 0xb4, 0xdf, 0x6b, 0x98, 0xac,
	0xc9, 0xf0, 0xf5, 0xdb, 0xd3, 0xf1, 0x70, 0xd2, 0x6b, 0xe2, 0x48, 0x37, 0x0e, 0x8e, 0xfa, 0x83,
	0xf1, 0x41, 0x74, 0x58, 0x85, 0x5a, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x31, 0x30, 0xfd, 0xc8,
	0xf7, 0x05, 0x00, 0x00,
}
