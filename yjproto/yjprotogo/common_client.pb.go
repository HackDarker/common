// Code generated by protoc-gen-go.
// source: common_client.proto
// DO NOT EDIT!

/*
Package yjprotogo is a generated protocol buffer package.

It is generated from these files:
	common_client.proto
	common_enum.proto
	common_server.proto
	hall.proto
	hall_base.proto
	hall_common_server_award.proto
	hall_common_server_model.proto
	hall_playback.proto
	mj_changchun_base.proto
	mj_changchun_play.proto
	sdy_base.proto
	sdy_desk.proto
	sdy_hall.proto
	sdy_play.proto
	tdk.proto
	tdk_base.proto
	tdk_data.proto

It has these top-level messages:
	ProtoHeader
	ServerInfo
	QuickConn
	AckQuickConn
	WeixinInfo
	CommonReqReg
	CommonAckReg
	CmOffline
	CmHearbeat
	GameSession
	HallReqEvent
	HallAckEvent
	HallLotteryItem
	HallSignLotteryInfo
	HallDrawLotteryInfo
	HallReqMailList
	HallAckMailList
	HallReqTask
	HallAckTask
	HallReqCheckTask
	HallAckCheckTask
	HallReqTaskSum
	HallAckTaskSum
	HallReqCheckBonus
	HallAckCheckBonus
	HallReqBagItems
	HallAckBagItems
	HallReqUserData
	HallAckUserData
	HallReqUpdateRealData
	HallAckUpdateRealData
	HallReqGoodsList
	HallAckGoodsList
	HallReqGoodsBuy
	HallAckGoodsBuy
	HallGoodsItemMsg
	HallReqRank
	HallAckRank
	HallReqDrawLottery
	HallAckDrawLottery
	HallReqDsLotteryInfo
	HallAckDsLotteryInfo
	HallReqFriendsList
	HallAckFriendsList
	HallReqRecommendUserList
	HallAckRecommendUserList
	HallReqFriendsSearch
	HallAckFriendsSearch
	HallReqAddFriend
	HallAckAddFriend
	HallReqDelFriend
	HallAckDelFriend
	HallFriendState
	HallUserInfo
	HallAckStrongboxInfo
	HallReqStrongboxInfo
	HallReqStrongboxAccess
	HallAckStrongboxAccess
	Game_GameRecord
	BeanUserRecord
	BeanGameRecord
	Game_AckGameRecord
	HallBeanBisai
	HallReqBisai
	HallReqFriendLotteryList
	HallAckFriendLotteryList
	HallReqFriendLotteryDraw
	HallAckFriendLotteryDraw
	HallReqDistanceMatched
	HallAckDistanceMatched
	HallItemEvent
	HallMailItem
	HallBagItem
	HallItemTask
	HallUserData
	HallRankItem
	CoinZone
	DiamondZone
	ExchangeZone
	BuyZone
	GoodsItem
	HallStrongboxInfo
	CommonReqGameLogin
	CommonReqGameLoginViaInput
	CommonAckGameLogin
	HallReqNewRoomList
	HallAckNewRoomList
	HallNewRoom_List
	Taward
	AwardMOnline
	User
	TNotice
	TGameCounts
	TUserTask
	Push
	PlaybackReqPage
	PlaybackAckPage
	PlaybackSnapshot
	PlaybackDeskInfo
	PlaybackItem
	PlaybackReqPageByGameid
	UserCoinBean
	RoomTypeInfo
	EndLotteryInfo
	PlayOptions
	DeskGameInfo
	PlayerInfo
	CardInfo
	ComposeCard
	PlayerCard
	WinCoinInfo
	Game_SendGameInfo
	Game_AckActHu
	P16AckGameOpening
	Game_SendCurrentResult
	P16AckRoomInit
	P16BeanDeskInfo
	P16ReqCreateDesk
	P16AckCreateDesk
	P16ReqEnterDesk
	P16AckEnterDesk
	P16ReqDissolveBeginGame
	P16AckDissolveBeginGame
	P16DissolveBeginGame
	P16AckPlayerEnter
	P16AckPlayerExit
	P16ReqReady
	P16AckReady
	P16AckOpening
	P16BcMoPai
	P16AckGameOverturn
	P16AckGameDealCards
	GangOverTurn
	BuGangOverTurn
	ChiOverTurn
	JiaoInfo
	JiaoPaiInfo
	P16AckGameSendOutCard
	Game_SendEndLottery
	P16ReqGameSendOutCard
	P16ReqChi
	P16AckChi
	P16ReqGameActPeng
	P16AckGameActPeng
	P16ReqGameActGang
	P16AckGameActGang
	P16ReqGameActBugang
	P16AckGameActBugang
	P16ReqGameActGuo
	P16AckGameActGuo
	P16ReqGameActHu
	P16ReqGameTing
	P16AckGameTing
	P16AckGameDabao
	P16AckGameChangbao
	SdyBaseUserPaiIds
	SdyBaseRoomTypeInfo
	SdyBaseTimerInfo
	SdyBasePlayerInfo
	SdyBaseCommonRateInfo
	SdyBaseDeskInfo
	SdyReqReady
	SdyAckReady
	SdyBaseWinCoinInfo
	SdyBcCurrentResult
	SdyBaseBill
	SdyBcEndLotteryInfo
	SdyReUserOutPai
	SdyReReady
	SdyReHuanDi
	SdyRePlay
	SdyReJiaoFen
	SdyReLenHandPokers
	SdyBcReconnectInfo
	SdyBcIsOnLine
	SdyReqCreateDesk
	SdyAckCreateDesk
	SdyReqEnterDesk
	SdyAckEnterDesk
	SdyBcUserBreak
	SdyReqDissolveDeskOwner
	SdyBcDissolveDeskOwner
	SdyReqApplyDissolveDeskApllyer
	SdyBcApplyDissolveDeskApplyer
	SdyReqApplyDissolveDeskUser
	SdyAckApplyDissolveDeskUser
	SdyReqLeaveDesk
	SdyAckLeaveDesk
	SdyReqSendMessage
	SdyAckSendMessage
	SdyBcOpening
	SdyDealCards
	SdyBcPlayerPokers
	SdyBcJiaoFen
	SdyReqJiaoFen
	SdyAckJiaoFen
	SdyBcJiaoFenResult
	SdyBcStartPlay
	SdyReqOutCards
	SdyAckOutCards
	SdyBcOverTurn
	SdyBcWhatPai
	SdyBcScorePai
	SdyBcGameInfo
	SdyBcDingZhu
	SdyReqDingZhu
	SdyBcDingZhuResult
	SdyAckDingZhu
	SdyBcHuanDi
	SdyReqHuanDi
	SdyAckHuanDi
	TdkHeartBeat
	TdkJoinRoom
	TdkJoinRoomRsp
	TdkLeaveRoom
	TdkLeaveRoomRsp
	TdkJoinDesk
	TdkJoinDeskRsp
	TdkCreateDesk
	TdkCreateDeskRsp
	TdkLeaveDesk
	TdkLeaveDeskRsp
	TdkUserReady
	TdkUserReadyRsp
	TdkReturnDesk
	TdkDissolveDeskReq
	TdkDissolveDeskUser
	TdkDisDeskDesicionReq
	TdkDisDeskDesicionRsp
	TdkDisDeskReqResult
	TdkDisDesk
	TdkSendPoker
	TdkStartBet
	TdkBet
	TdkBetRsp
	TdkFold
	TdkFoldRsp
	TdkQiJiao
	TdkQiJiaoRsp
	TdkFanTi
	TdkFanTiRsp
	TdkPass
	TdkPassRsp
	TdkKaiPai
	TdkKaiPaiRsp
	TdkRoundEnd
	TdkZhanJi
	TdkZhanJiRsp
	TdkMatchReq
	TdkMatchRsp
	TdkMatchAddUserRsp
	TdkJoinPlayingDeskRsp
	TdkEnterGame
	TdkEnterGameRsp
	TdkDeskUserData
	TdkUserPokerData
	TdkDisUserData
	TdkDeskConfig
	TdkZhanJiData
*/
package yjprotogo

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

// 注册类型
type CommonEnumReg int32

const (
	CommonEnumReg_RET_TYPE_TOURIST CommonEnumReg = 1
	CommonEnumReg_RET_TYPE_WEIXIN  CommonEnumReg = 2
)

var CommonEnumReg_name = map[int32]string{
	1: "RET_TYPE_TOURIST",
	2: "RET_TYPE_WEIXIN",
}
var CommonEnumReg_value = map[string]int32{
	"RET_TYPE_TOURIST": 1,
	"RET_TYPE_WEIXIN":  2,
}

func (x CommonEnumReg) Enum() *CommonEnumReg {
	p := new(CommonEnumReg)
	*p = x
	return p
}
func (x CommonEnumReg) String() string {
	return proto.EnumName(CommonEnumReg_name, int32(x))
}
func (x *CommonEnumReg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommonEnumReg_value, data, "CommonEnumReg")
	if err != nil {
		return err
	}
	*x = CommonEnumReg(value)
	return nil
}
func (CommonEnumReg) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 客户端系统类型
type CommonEnumOsType int32

const (
	CommonEnumOsType_OS_IOS     CommonEnumOsType = 1
	CommonEnumOsType_OS_ANDROID CommonEnumOsType = 2
	CommonEnumOsType_OS_WEB     CommonEnumOsType = 3
)

var CommonEnumOsType_name = map[int32]string{
	1: "OS_IOS",
	2: "OS_ANDROID",
	3: "OS_WEB",
}
var CommonEnumOsType_value = map[string]int32{
	"OS_IOS":     1,
	"OS_ANDROID": 2,
	"OS_WEB":     3,
}

func (x CommonEnumOsType) Enum() *CommonEnumOsType {
	p := new(CommonEnumOsType)
	*p = x
	return p
}
func (x CommonEnumOsType) String() string {
	return proto.EnumName(CommonEnumOsType_name, int32(x))
}
func (x *CommonEnumOsType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommonEnumOsType_value, data, "CommonEnumOsType")
	if err != nil {
		return err
	}
	*x = CommonEnumOsType(value)
	return nil
}
func (CommonEnumOsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

// 服务器信息
type ServerInfo struct {
	Ip               *string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port             *int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	ReleaseTag       *int32  `protobuf:"varint,3,opt,name=releaseTag" json:"releaseTag,omitempty"`
	CurrVersion      *int32  `protobuf:"varint,4,opt,name=currVersion" json:"currVersion,omitempty"`
	IsUpdate         *int32  `protobuf:"varint,5,opt,name=isUpdate" json:"isUpdate,omitempty"`
	DownloadUrl      *string `protobuf:"bytes,6,opt,name=downloadUrl" json:"downloadUrl,omitempty"`
	VersionInfo      *string `protobuf:"bytes,7,opt,name=versionInfo" json:"versionInfo,omitempty"`
	IsMaintain       *int32  `protobuf:"varint,8,opt,name=isMaintain" json:"isMaintain,omitempty"`
	MaintainMsg      *string `protobuf:"bytes,9,opt,name=maintainMsg" json:"maintainMsg,omitempty"`
	Status           *int32  `protobuf:"varint,10,opt,name=status" json:"status,omitempty"`
	GameId           *int32  `protobuf:"varint,11,opt,name=gameId" json:"gameId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

func (m *ServerInfo) GetReleaseTag() int32 {
	if m != nil && m.ReleaseTag != nil {
		return *m.ReleaseTag
	}
	return 0
}

func (m *ServerInfo) GetCurrVersion() int32 {
	if m != nil && m.CurrVersion != nil {
		return *m.CurrVersion
	}
	return 0
}

func (m *ServerInfo) GetIsUpdate() int32 {
	if m != nil && m.IsUpdate != nil {
		return *m.IsUpdate
	}
	return 0
}

func (m *ServerInfo) GetDownloadUrl() string {
	if m != nil && m.DownloadUrl != nil {
		return *m.DownloadUrl
	}
	return ""
}

func (m *ServerInfo) GetVersionInfo() string {
	if m != nil && m.VersionInfo != nil {
		return *m.VersionInfo
	}
	return ""
}

func (m *ServerInfo) GetIsMaintain() int32 {
	if m != nil && m.IsMaintain != nil {
		return *m.IsMaintain
	}
	return 0
}

func (m *ServerInfo) GetMaintainMsg() string {
	if m != nil && m.MaintainMsg != nil {
		return *m.MaintainMsg
	}
	return ""
}

func (m *ServerInfo) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *ServerInfo) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
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
func (*QuickConn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
	Header            *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	GameServer        []*ServerInfo `protobuf:"bytes,2,rep,name=gameServer" json:"gameServer,omitempty"`
	ServerListVersion *int32        `protobuf:"varint,3,opt,name=serverListVersion" json:"serverListVersion,omitempty"`
	XXX_unrecognized  []byte        `json:"-"`
}

func (m *AckQuickConn) Reset()                    { *m = AckQuickConn{} }
func (m *AckQuickConn) String() string            { return proto.CompactTextString(m) }
func (*AckQuickConn) ProtoMessage()               {}
func (*AckQuickConn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AckQuickConn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AckQuickConn) GetGameServer() []*ServerInfo {
	if m != nil {
		return m.GameServer
	}
	return nil
}

func (m *AckQuickConn) GetServerListVersion() int32 {
	if m != nil && m.ServerListVersion != nil {
		return *m.ServerListVersion
	}
	return 0
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
func (*WeixinInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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

// 请求
type CommonReqReg struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RegType          *int32       `protobuf:"varint,2,opt,name=regType" json:"regType,omitempty"`
	WxInfo           *WeixinInfo  `protobuf:"bytes,3,opt,name=wxInfo" json:"wxInfo,omitempty"`
	ClientOSType     *int32       `protobuf:"varint,4,opt,name=clientOSType" json:"clientOSType,omitempty"`
	ChannelId        *string      `protobuf:"bytes,5,opt,name=channelId" json:"channelId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonReqReg) Reset()                    { *m = CommonReqReg{} }
func (m *CommonReqReg) String() string            { return proto.CompactTextString(m) }
func (*CommonReqReg) ProtoMessage()               {}
func (*CommonReqReg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CommonReqReg) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommonReqReg) GetRegType() int32 {
	if m != nil && m.RegType != nil {
		return *m.RegType
	}
	return 0
}

func (m *CommonReqReg) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *CommonReqReg) GetClientOSType() int32 {
	if m != nil && m.ClientOSType != nil {
		return *m.ClientOSType
	}
	return 0
}

func (m *CommonReqReg) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

// 回复
type CommonAckReg struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonAckReg) Reset()                    { *m = CommonAckReg{} }
func (m *CommonAckReg) String() string            { return proto.CompactTextString(m) }
func (*CommonAckReg) ProtoMessage()               {}
func (*CommonAckReg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CommonAckReg) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CommonAckReg) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 表示玩家掉线
type CmOffline struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CmOffline) Reset()                    { *m = CmOffline{} }
func (m *CmOffline) String() string            { return proto.CompactTextString(m) }
func (*CmOffline) ProtoMessage()               {}
func (*CmOffline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// 表示心跳
type CmHearbeat struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CmHearbeat) Reset()                    { *m = CmHearbeat{} }
func (m *CmHearbeat) String() string            { return proto.CompactTextString(m) }
func (*CmHearbeat) ProtoMessage()               {}
func (*CmHearbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CmHearbeat) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoHeader)(nil), "yjprotogo.ProtoHeader")
	proto.RegisterType((*ServerInfo)(nil), "yjprotogo.ServerInfo")
	proto.RegisterType((*QuickConn)(nil), "yjprotogo.QuickConn")
	proto.RegisterType((*AckQuickConn)(nil), "yjprotogo.AckQuickConn")
	proto.RegisterType((*WeixinInfo)(nil), "yjprotogo.WeixinInfo")
	proto.RegisterType((*CommonReqReg)(nil), "yjprotogo.common_req_reg")
	proto.RegisterType((*CommonAckReg)(nil), "yjprotogo.common_ack_reg")
	proto.RegisterType((*CmOffline)(nil), "yjprotogo.cm_offline")
	proto.RegisterType((*CmHearbeat)(nil), "yjprotogo.cm_hearbeat")
	proto.RegisterEnum("yjprotogo.CommonEnumReg", CommonEnumReg_name, CommonEnumReg_value)
	proto.RegisterEnum("yjprotogo.CommonEnumOsType", CommonEnumOsType_name, CommonEnumOsType_value)
}

var fileDescriptor0 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0xb8, 0x49, 0xeb, 0x49, 0xda, 0xba, 0x6e, 0x41, 0xe6, 0x86, 0x2c, 0x81, 0xa0,
	0x87, 0x1e, 0x90, 0xb8, 0xc1, 0xa1, 0xd0, 0x48, 0xb5, 0xa0, 0x4d, 0x69, 0x5c, 0x0a, 0x27, 0x6b,
	0xb1, 0xb7, 0xee, 0x12, 0x7b, 0xd7, 0xec, 0xda, 0xb4, 0x91, 0x78, 0x07, 0xce, 0xbc, 0x17, 0x0f,
	0xc4, 0xec, 0x66, 0x13, 0x19, 0x0e, 0xa8, 0x3d, 0x58, 0xda, 0x19, 0xcf, 0xcc, 0x7e, 0xff, 0xcc,
	0xd8, 0xb0, 0x9b, 0x89, 0xaa, 0x12, 0x3c, 0xcd, 0x4a, 0x46, 0x79, 0x73, 0x50, 0x4b, 0xd1, 0x88,
	0xc0, 0x9b, 0x7f, 0x35, 0x87, 0x42, 0x44, 0xef, 0x60, 0x78, 0xa6, 0x8f, 0xc7, 0x94, 0xe4, 0x54,
	0x06, 0xdb, 0xb0, 0xfe, 0x9d, 0x4a, 0xc5, 0x04, 0x0f, 0x9d, 0xc7, 0xce, 0x33, 0x2f, 0xd8, 0x82,
	0x41, 0xab, 0xa8, 0x8c, 0xf3, 0xb0, 0x87, 0xf6, 0x66, 0x30, 0x82, 0xb5, 0x4c, 0xe4, 0x34, 0x74,
	0xd1, 0xea, 0x07, 0x9b, 0xd0, 0xa7, 0x52, 0x0a, 0x19, 0xae, 0xe9, 0xe0, 0xe8, 0xb7, 0x03, 0x30,
	0xa5, 0x12, 0x2b, 0xc4, 0xfc, 0x4a, 0x04, 0x00, 0x3d, 0x56, 0xdb, 0x3a, 0x98, 0x57, 0x0b, 0xd9,
	0x98, 0x2a, 0xfd, 0x00, 0x5f, 0x49, 0x5a, 0x52, 0xa2, 0x68, 0x42, 0x0a, 0x5b, 0x6b, 0x17, 0x86,
	0x59, 0x2b, 0xe5, 0x47, 0x7b, 0xfd, 0x9a, 0x71, 0xfa, 0xb0, 0xc1, 0xd4, 0x45, 0x9d, 0x93, 0x86,
	0x86, 0xfd, 0x65, 0x58, 0x2e, 0x6e, 0x78, 0x29, 0x48, 0x7e, 0x21, 0xcb, 0x70, 0x60, 0xaa, 0xa3,
	0xd3, 0x62, 0xeb, 0x8b, 0xc3, 0x75, 0xe3, 0xc4, 0x4b, 0x98, 0x3a, 0x21, 0x8c, 0x37, 0xf8, 0x84,
	0x1b, 0xcb, 0xec, 0xca, 0x7a, 0x4e, 0x54, 0x11, 0x7a, 0x4b, 0x8d, 0xaa, 0x21, 0x4d, 0xab, 0x42,
	0x30, 0x41, 0x68, 0x17, 0xa4, 0xa2, 0xa8, 0x79, 0xa8, 0xed, 0xe8, 0xa7, 0x03, 0xde, 0x87, 0x96,
	0x65, 0xb3, 0xb7, 0x82, 0xf3, 0xe0, 0x29, 0x0c, 0xae, 0x4d, 0xb3, 0x8c, 0xb2, 0xe1, 0x8b, 0x87,
	0x07, 0xab, 0x6e, 0x1e, 0x74, 0x5b, 0xb9, 0x03, 0x5e, 0x76, 0x4d, 0x38, 0xa7, 0xa5, 0x6d, 0x9e,
	0xd7, 0x29, 0xfc, 0x1f, 0xc9, 0x88, 0x5d, 0x12, 0x5e, 0xb4, 0xa4, 0xd0, 0x81, 0xfd, 0x25, 0x91,
	0x9d, 0x82, 0xd6, 0xbb, 0x19, 0xfd, 0x80, 0xd1, 0x61, 0x36, 0xbb, 0x3f, 0xd3, 0x73, 0x00, 0x0d,
	0xb0, 0x98, 0x11, 0x42, 0xb9, 0x18, 0xfb, 0xa0, 0x13, 0xdb, 0x19, 0xde, 0x23, 0xd8, 0x51, 0xc6,
	0x7a, 0xcf, 0x54, 0xb3, 0x24, 0x34, 0xd8, 0xd1, 0x0c, 0xe0, 0x92, 0xb2, 0x5b, 0x66, 0x9a, 0xad,
	0xd9, 0x44, 0x4d, 0x39, 0xb2, 0x2d, 0x26, 0x8d, 0x23, 0xe3, 0xc8, 0x75, 0x8a, 0xf7, 0x58, 0xd9,
	0xb8, 0x54, 0x9a, 0x4e, 0x8f, 0xcb, 0x35, 0x8e, 0x21, 0xb8, 0x8a, 0xde, 0x5a, 0xbd, 0x7a, 0xa3,
	0x58, 0x33, 0x37, 0x4a, 0x4d, 0x6c, 0xcb, 0xf5, 0x1c, 0x17, 0x52, 0xbd, 0xe8, 0x97, 0x03, 0x5b,
	0x76, 0x87, 0x25, 0xfd, 0x86, 0x4f, 0x71, 0x67, 0xb5, 0x58, 0x0b, 0xc3, 0x93, 0x79, 0x4d, 0xed,
	0xda, 0x3d, 0x81, 0xc1, 0xcd, 0xad, 0xd9, 0x10, 0xd7, 0x24, 0x76, 0xa5, 0x77, 0x14, 0xed, 0xc1,
	0x68, 0xf1, 0xb9, 0x4c, 0xa6, 0x26, 0x79, 0xc1, 0xf9, 0xd7, 0x3c, 0x0d, 0x6c, 0x74, 0xbc, 0x42,
	0x23, 0xd9, 0xec, 0x5e, 0x68, 0xff, 0x7c, 0x56, 0xd1, 0x08, 0x20, 0xab, 0x52, 0x71, 0x75, 0x55,
	0x32, 0x4e, 0xa3, 0x97, 0xb8, 0x17, 0x55, 0x8a, 0x85, 0xe4, 0x17, 0x4a, 0x9a, 0xbb, 0x16, 0xdd,
	0x7f, 0x05, 0xdb, 0x16, 0x87, 0xf2, 0xb6, 0x32, 0x3c, 0x7b, 0xe0, 0x9f, 0x8f, 0x93, 0x34, 0xf9,
	0x7c, 0x36, 0x4e, 0x93, 0xc9, 0xc5, 0x79, 0x3c, 0x4d, 0x7c, 0x07, 0xf7, 0x6e, 0x7b, 0xe5, 0xbd,
	0x1c, 0xc7, 0x9f, 0xe2, 0x53, 0xbf, 0xb7, 0xff, 0x7a, 0xf5, 0xaf, 0x30, 0xd9, 0x42, 0xa5, 0x0d,
	0x8a, 0xc7, 0x6d, 0x1c, 0x4c, 0xa6, 0x69, 0x3c, 0x99, 0x62, 0xde, 0x16, 0x00, 0x9e, 0x0f, 0x4f,
	0x8f, 0xce, 0x27, 0xf1, 0x91, 0xdf, 0xb3, 0xef, 0x2e, 0xc7, 0x6f, 0x7c, 0xf7, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7c, 0x9c, 0xe6, 0xd6, 0x69, 0x04, 0x00, 0x00,
}
