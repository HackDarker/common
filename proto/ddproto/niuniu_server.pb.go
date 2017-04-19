// Code generated by protoc-gen-go.
// source: niuniu_server.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// =================================公共================================
// 牛牛牌的类型
type NiuniuEnum_PokerType int32

const (
	NiuniuEnum_PokerType_NO_NIU           NiuniuEnum_PokerType = 1
	NiuniuEnum_PokerType_NIU_1            NiuniuEnum_PokerType = 2
	NiuniuEnum_PokerType_NIU_2            NiuniuEnum_PokerType = 3
	NiuniuEnum_PokerType_NIU_3            NiuniuEnum_PokerType = 4
	NiuniuEnum_PokerType_NIU_4            NiuniuEnum_PokerType = 5
	NiuniuEnum_PokerType_NIU_5            NiuniuEnum_PokerType = 6
	NiuniuEnum_PokerType_NIU_6            NiuniuEnum_PokerType = 7
	NiuniuEnum_PokerType_NIU_7            NiuniuEnum_PokerType = 8
	NiuniuEnum_PokerType_NIU_8            NiuniuEnum_PokerType = 9
	NiuniuEnum_PokerType_NIU_9            NiuniuEnum_PokerType = 10
	NiuniuEnum_PokerType_NIU_NIU          NiuniuEnum_PokerType = 11
	NiuniuEnum_PokerType_YIN_NIU          NiuniuEnum_PokerType = 12
	NiuniuEnum_PokerType_JIN_NIU          NiuniuEnum_PokerType = 13
	NiuniuEnum_PokerType_WU_XIAO_NIU      NiuniuEnum_PokerType = 14
	NiuniuEnum_PokerType_NIU_ZHA_DAN      NiuniuEnum_PokerType = 15
	NiuniuEnum_PokerType_NIU_YI_TIAO_LONG NiuniuEnum_PokerType = 16
)

var NiuniuEnum_PokerType_name = map[int32]string{
	1:  "NO_NIU",
	2:  "NIU_1",
	3:  "NIU_2",
	4:  "NIU_3",
	5:  "NIU_4",
	6:  "NIU_5",
	7:  "NIU_6",
	8:  "NIU_7",
	9:  "NIU_8",
	10: "NIU_9",
	11: "NIU_NIU",
	12: "YIN_NIU",
	13: "JIN_NIU",
	14: "WU_XIAO_NIU",
	15: "NIU_ZHA_DAN",
	16: "NIU_YI_TIAO_LONG",
}
var NiuniuEnum_PokerType_value = map[string]int32{
	"NO_NIU":           1,
	"NIU_1":            2,
	"NIU_2":            3,
	"NIU_3":            4,
	"NIU_4":            5,
	"NIU_5":            6,
	"NIU_6":            7,
	"NIU_7":            8,
	"NIU_8":            9,
	"NIU_9":            10,
	"NIU_NIU":          11,
	"YIN_NIU":          12,
	"JIN_NIU":          13,
	"WU_XIAO_NIU":      14,
	"NIU_ZHA_DAN":      15,
	"NIU_YI_TIAO_LONG": 16,
}

func (x NiuniuEnum_PokerType) Enum() *NiuniuEnum_PokerType {
	p := new(NiuniuEnum_PokerType)
	*p = x
	return p
}
func (x NiuniuEnum_PokerType) String() string {
	return proto.EnumName(NiuniuEnum_PokerType_name, int32(x))
}
func (x *NiuniuEnum_PokerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnum_PokerType_value, data, "NiuniuEnum_PokerType")
	if err != nil {
		return err
	}
	*x = NiuniuEnum_PokerType(value)
	return nil
}
func (NiuniuEnum_PokerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor29, []int{0} }

// 用户游戏状态
type NiuniuEnumUserState int32

const (
	NiuniuEnumUserState_NIU_USER_STATUS_NO_READY   NiuniuEnumUserState = 1
	NiuniuEnumUserState_NIU_USER_STATUS_IS_READY   NiuniuEnumUserState = 2
	NiuniuEnumUserState_NIU_USER_STATUS_IS_GAMMING NiuniuEnumUserState = 3
)

var NiuniuEnumUserState_name = map[int32]string{
	1: "NIU_USER_STATUS_NO_READY",
	2: "NIU_USER_STATUS_IS_READY",
	3: "NIU_USER_STATUS_IS_GAMMING",
}
var NiuniuEnumUserState_value = map[string]int32{
	"NIU_USER_STATUS_NO_READY":   1,
	"NIU_USER_STATUS_IS_READY":   2,
	"NIU_USER_STATUS_IS_GAMMING": 3,
}

func (x NiuniuEnumUserState) Enum() *NiuniuEnumUserState {
	p := new(NiuniuEnumUserState)
	*p = x
	return p
}
func (x NiuniuEnumUserState) String() string {
	return proto.EnumName(NiuniuEnumUserState_name, int32(x))
}
func (x *NiuniuEnumUserState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumUserState_value, data, "NiuniuEnumUserState")
	if err != nil {
		return err
	}
	*x = NiuniuEnumUserState(value)
	return nil
}
func (NiuniuEnumUserState) EnumDescriptor() ([]byte, []int) { return fileDescriptor29, []int{1} }

// 房间状态
type NiuniuEnumDeskState int32

const (
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER NiuniuEnumDeskState = 1
	NiuniuEnumDeskState_NIU_DESK_STATUS_NO_READY   NiuniuEnumDeskState = 2
	NiuniuEnumDeskState_NIU_DESK_STATUS_IS_READY   NiuniuEnumDeskState = 3
	NiuniuEnumDeskState_NIU_DESK_STATUS_IS_GAMMING NiuniuEnumDeskState = 4
)

var NiuniuEnumDeskState_name = map[int32]string{
	1: "NIU_DESK_STATUS_WAIT_ENTER",
	2: "NIU_DESK_STATUS_NO_READY",
	3: "NIU_DESK_STATUS_IS_READY",
	4: "NIU_DESK_STATUS_IS_GAMMING",
}
var NiuniuEnumDeskState_value = map[string]int32{
	"NIU_DESK_STATUS_WAIT_ENTER": 1,
	"NIU_DESK_STATUS_NO_READY":   2,
	"NIU_DESK_STATUS_IS_READY":   3,
	"NIU_DESK_STATUS_IS_GAMMING": 4,
}

func (x NiuniuEnumDeskState) Enum() *NiuniuEnumDeskState {
	p := new(NiuniuEnumDeskState)
	*p = x
	return p
}
func (x NiuniuEnumDeskState) String() string {
	return proto.EnumName(NiuniuEnumDeskState_name, int32(x))
}
func (x *NiuniuEnumDeskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumDeskState_value, data, "NiuniuEnumDeskState")
	if err != nil {
		return err
	}
	*x = NiuniuEnumDeskState(value)
	return nil
}
func (NiuniuEnumDeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor29, []int{2} }

// 坐庄规则
type NiuniuEnumBankerRule int32

const (
	NiuniuEnumBankerRule_DING_ZHUANG       NiuniuEnumBankerRule = 1
	NiuniuEnumBankerRule_SUI_JI_ZUO_ZHUANG NiuniuEnumBankerRule = 2
	NiuniuEnumBankerRule_QIANG_ZHUANG      NiuniuEnumBankerRule = 3
)

var NiuniuEnumBankerRule_name = map[int32]string{
	1: "DING_ZHUANG",
	2: "SUI_JI_ZUO_ZHUANG",
	3: "QIANG_ZHUANG",
}
var NiuniuEnumBankerRule_value = map[string]int32{
	"DING_ZHUANG":       1,
	"SUI_JI_ZUO_ZHUANG": 2,
	"QIANG_ZHUANG":      3,
}

func (x NiuniuEnumBankerRule) Enum() *NiuniuEnumBankerRule {
	p := new(NiuniuEnumBankerRule)
	*p = x
	return p
}
func (x NiuniuEnumBankerRule) String() string {
	return proto.EnumName(NiuniuEnumBankerRule_name, int32(x))
}
func (x *NiuniuEnumBankerRule) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumBankerRule_value, data, "NiuniuEnumBankerRule")
	if err != nil {
		return err
	}
	*x = NiuniuEnumBankerRule(value)
	return nil
}
func (NiuniuEnumBankerRule) EnumDescriptor() ([]byte, []int) { return fileDescriptor29, []int{3} }

// 打出去的牌
type NiuniuSrvPoker struct {
	Pais             []*CommonSrvPokerPai  `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *NiuniuEnum_PokerType `protobuf:"varint,3,opt,name=type,enum=ddproto.NiuniuEnum_PokerType" json:"type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuSrvPoker) Reset()                    { *m = NiuniuSrvPoker{} }
func (m *NiuniuSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvPoker) ProtoMessage()               {}
func (*NiuniuSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{0} }

func (m *NiuniuSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *NiuniuSrvPoker) GetType() NiuniuEnum_PokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return NiuniuEnum_PokerType_NO_NIU
}

// 对局账单信息
type NiuniuUserBill struct {
	Score            *int32 `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
	CountHasNiu      *int32 `protobuf:"varint,2,opt,name=count_has_niu,json=countHasNiu" json:"count_has_niu,omitempty"`
	CountNoNiu       *int32 `protobuf:"varint,3,opt,name=count_no_niu,json=countNoNiu" json:"count_no_niu,omitempty"`
	CountWin         *int32 `protobuf:"varint,4,opt,name=count_win,json=countWin" json:"count_win,omitempty"`
	CountLost        *int32 `protobuf:"varint,5,opt,name=count_lost,json=countLost" json:"count_lost,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NiuniuUserBill) Reset()                    { *m = NiuniuUserBill{} }
func (m *NiuniuUserBill) String() string            { return proto.CompactTextString(m) }
func (*NiuniuUserBill) ProtoMessage()               {}
func (*NiuniuUserBill) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{1} }

func (m *NiuniuUserBill) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *NiuniuUserBill) GetCountHasNiu() int32 {
	if m != nil && m.CountHasNiu != nil {
		return *m.CountHasNiu
	}
	return 0
}

func (m *NiuniuUserBill) GetCountNoNiu() int32 {
	if m != nil && m.CountNoNiu != nil {
		return *m.CountNoNiu
	}
	return 0
}

func (m *NiuniuUserBill) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *NiuniuUserBill) GetCountLost() int32 {
	if m != nil && m.CountLost != nil {
		return *m.CountLost
	}
	return 0
}

// desk 配置选项
type NiuniuSrvDeskOption struct {
	MinUser          *int32                `protobuf:"varint,1,opt,name=minUser" json:"minUser,omitempty"`
	MaxUser          *int32                `protobuf:"varint,2,opt,name=maxUser" json:"maxUser,omitempty"`
	MaxCircle        *int32                `protobuf:"varint,3,opt,name=maxCircle" json:"maxCircle,omitempty"`
	HasFlower        *bool                 `protobuf:"varint,4,opt,name=hasFlower" json:"hasFlower,omitempty"`
	BankRule         *NiuniuEnumBankerRule `protobuf:"varint,5,opt,name=bankRule,enum=ddproto.NiuniuEnumBankerRule" json:"bankRule,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuSrvDeskOption) Reset()                    { *m = NiuniuSrvDeskOption{} }
func (m *NiuniuSrvDeskOption) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvDeskOption) ProtoMessage()               {}
func (*NiuniuSrvDeskOption) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{2} }

func (m *NiuniuSrvDeskOption) GetMinUser() int32 {
	if m != nil && m.MinUser != nil {
		return *m.MinUser
	}
	return 0
}

func (m *NiuniuSrvDeskOption) GetMaxUser() int32 {
	if m != nil && m.MaxUser != nil {
		return *m.MaxUser
	}
	return 0
}

func (m *NiuniuSrvDeskOption) GetMaxCircle() int32 {
	if m != nil && m.MaxCircle != nil {
		return *m.MaxCircle
	}
	return 0
}

func (m *NiuniuSrvDeskOption) GetHasFlower() bool {
	if m != nil && m.HasFlower != nil {
		return *m.HasFlower
	}
	return false
}

func (m *NiuniuSrvDeskOption) GetBankRule() NiuniuEnumBankerRule {
	if m != nil && m.BankRule != nil {
		return *m.BankRule
	}
	return NiuniuEnumBankerRule_DING_ZHUANG
}

// desk 的信息
type NiuniuClientDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	DeskNumber       *string              `protobuf:"bytes,2,opt,name=deskNumber" json:"deskNumber,omitempty"`
	GameNumber       *int32               `protobuf:"varint,3,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32               `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	Status           *NiuniuEnumDeskState `protobuf:"varint,5,opt,name=status,enum=ddproto.NiuniuEnumDeskState" json:"status,omitempty"`
	LastWiner        *uint32              `protobuf:"varint,6,opt,name=lastWiner" json:"lastWiner,omitempty"`
	CircleNo         *int32               `protobuf:"varint,8,opt,name=circleNo" json:"circleNo,omitempty"`
	Owner            *uint32              `protobuf:"varint,10,opt,name=owner" json:"owner,omitempty"`
	CurrBanker       *uint32              `protobuf:"varint,11,opt,name=currBanker" json:"currBanker,omitempty"`
	DeskOption       *NiuniuSrvDeskOption `protobuf:"bytes,12,opt,name=deskOption" json:"deskOption,omitempty"`
	Users            []*NiuniuClientUser  `protobuf:"bytes,13,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *NiuniuClientDesk) Reset()                    { *m = NiuniuClientDesk{} }
func (m *NiuniuClientDesk) String() string            { return proto.CompactTextString(m) }
func (*NiuniuClientDesk) ProtoMessage()               {}
func (*NiuniuClientDesk) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{3} }

func (m *NiuniuClientDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *NiuniuClientDesk) GetDeskNumber() string {
	if m != nil && m.DeskNumber != nil {
		return *m.DeskNumber
	}
	return ""
}

func (m *NiuniuClientDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *NiuniuClientDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuClientDesk) GetStatus() NiuniuEnumDeskState {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER
}

func (m *NiuniuClientDesk) GetLastWiner() uint32 {
	if m != nil && m.LastWiner != nil {
		return *m.LastWiner
	}
	return 0
}

func (m *NiuniuClientDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *NiuniuClientDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *NiuniuClientDesk) GetCurrBanker() uint32 {
	if m != nil && m.CurrBanker != nil {
		return *m.CurrBanker
	}
	return 0
}

func (m *NiuniuClientDesk) GetDeskOption() *NiuniuSrvDeskOption {
	if m != nil {
		return m.DeskOption
	}
	return nil
}

func (m *NiuniuClientDesk) GetUsers() []*NiuniuClientUser {
	if m != nil {
		return m.Users
	}
	return nil
}

// 用户属性
type NiuniuClientUser struct {
	UserId           *uint32              `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Bill             *NiuniuUserBill      `protobuf:"bytes,6,opt,name=bill" json:"bill,omitempty"`
	Status           *NiuniuEnumUserState `protobuf:"varint,7,opt,name=status,enum=ddproto.NiuniuEnumUserState" json:"status,omitempty"`
	IsOnline         *bool                `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	Index            *int32               `protobuf:"varint,12,opt,name=index" json:"index,omitempty"`
	Pokers           *NiuniuSrvPoker      `protobuf:"bytes,13,opt,name=pokers" json:"pokers,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *NiuniuClientUser) Reset()                    { *m = NiuniuClientUser{} }
func (m *NiuniuClientUser) String() string            { return proto.CompactTextString(m) }
func (*NiuniuClientUser) ProtoMessage()               {}
func (*NiuniuClientUser) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{4} }

func (m *NiuniuClientUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *NiuniuClientUser) GetBill() *NiuniuUserBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *NiuniuClientUser) GetStatus() NiuniuEnumUserState {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return NiuniuEnumUserState_NIU_USER_STATUS_NO_READY
}

func (m *NiuniuClientUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *NiuniuClientUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *NiuniuClientUser) GetPokers() *NiuniuSrvPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

// desk 的信息
type NiuniuSrvDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	DeskNumber       *string              `protobuf:"bytes,2,opt,name=deskNumber" json:"deskNumber,omitempty"`
	GameNumber       *int32               `protobuf:"varint,3,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32               `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	Status           *NiuniuEnumDeskState `protobuf:"varint,5,opt,name=status,enum=ddproto.NiuniuEnumDeskState" json:"status,omitempty"`
	LastWiner        *uint32              `protobuf:"varint,6,opt,name=lastWiner" json:"lastWiner,omitempty"`
	CircleNo         *int32               `protobuf:"varint,8,opt,name=circleNo" json:"circleNo,omitempty"`
	Owner            *uint32              `protobuf:"varint,10,opt,name=owner" json:"owner,omitempty"`
	CurrBanker       *uint32              `protobuf:"varint,11,opt,name=currBanker" json:"currBanker,omitempty"`
	DeskOption       *NiuniuSrvDeskOption `protobuf:"bytes,12,opt,name=deskOption" json:"deskOption,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *NiuniuSrvDesk) Reset()                    { *m = NiuniuSrvDesk{} }
func (m *NiuniuSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvDesk) ProtoMessage()               {}
func (*NiuniuSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{5} }

func (m *NiuniuSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *NiuniuSrvDesk) GetDeskNumber() string {
	if m != nil && m.DeskNumber != nil {
		return *m.DeskNumber
	}
	return ""
}

func (m *NiuniuSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *NiuniuSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuSrvDesk) GetStatus() NiuniuEnumDeskState {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER
}

func (m *NiuniuSrvDesk) GetLastWiner() uint32 {
	if m != nil && m.LastWiner != nil {
		return *m.LastWiner
	}
	return 0
}

func (m *NiuniuSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *NiuniuSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *NiuniuSrvDesk) GetCurrBanker() uint32 {
	if m != nil && m.CurrBanker != nil {
		return *m.CurrBanker
	}
	return 0
}

func (m *NiuniuSrvDesk) GetDeskOption() *NiuniuSrvDeskOption {
	if m != nil {
		return m.DeskOption
	}
	return nil
}

// 用户属性
type NiuniuSrvUser struct {
	UserId           *uint32              `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Bill             *NiuniuUserBill      `protobuf:"bytes,6,opt,name=bill" json:"bill,omitempty"`
	Status           *NiuniuEnumUserState `protobuf:"varint,7,opt,name=status,enum=ddproto.NiuniuEnumUserState" json:"status,omitempty"`
	IsOnline         *bool                `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	Index            *int32               `protobuf:"varint,12,opt,name=index" json:"index,omitempty"`
	Pokers           *NiuniuSrvPoker      `protobuf:"bytes,13,opt,name=pokers" json:"pokers,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *NiuniuSrvUser) Reset()                    { *m = NiuniuSrvUser{} }
func (m *NiuniuSrvUser) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvUser) ProtoMessage()               {}
func (*NiuniuSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{6} }

func (m *NiuniuSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *NiuniuSrvUser) GetBill() *NiuniuUserBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *NiuniuSrvUser) GetStatus() NiuniuEnumUserState {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return NiuniuEnumUserState_NIU_USER_STATUS_NO_READY
}

func (m *NiuniuSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *NiuniuSrvUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *NiuniuSrvUser) GetPokers() *NiuniuSrvPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

// room 的信息
type NiuniuSrvRoom struct {
	RoomId           *int32  `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NiuniuSrvRoom) Reset()                    { *m = NiuniuSrvRoom{} }
func (m *NiuniuSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvRoom) ProtoMessage()               {}
func (*NiuniuSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{7} }

func (m *NiuniuSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func init() {
	proto.RegisterType((*NiuniuSrvPoker)(nil), "ddproto.niuniu_srv_poker")
	proto.RegisterType((*NiuniuUserBill)(nil), "ddproto.niuniu_user_bill")
	proto.RegisterType((*NiuniuSrvDeskOption)(nil), "ddproto.niuniu_srv_desk_option")
	proto.RegisterType((*NiuniuClientDesk)(nil), "ddproto.niuniu_client_desk")
	proto.RegisterType((*NiuniuClientUser)(nil), "ddproto.niuniu_client_user")
	proto.RegisterType((*NiuniuSrvDesk)(nil), "ddproto.niuniu_srv_desk")
	proto.RegisterType((*NiuniuSrvUser)(nil), "ddproto.niuniu_srv_user")
	proto.RegisterType((*NiuniuSrvRoom)(nil), "ddproto.niuniu_srv_room")
	proto.RegisterEnum("ddproto.NiuniuEnum_PokerType", NiuniuEnum_PokerType_name, NiuniuEnum_PokerType_value)
	proto.RegisterEnum("ddproto.NiuniuEnumUserState", NiuniuEnumUserState_name, NiuniuEnumUserState_value)
	proto.RegisterEnum("ddproto.NiuniuEnumDeskState", NiuniuEnumDeskState_name, NiuniuEnumDeskState_value)
	proto.RegisterEnum("ddproto.NiuniuEnumBankerRule", NiuniuEnumBankerRule_name, NiuniuEnumBankerRule_value)
}

var fileDescriptor29 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x55, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x66, 0xfc, 0x17, 0xbb, 0x1c, 0x27, 0x4d, 0xb3, 0xbb, 0xcc, 0x7a, 0x43, 0xb0, 0x7c, 0xb2,
	0x22, 0x11, 0x11, 0xf3, 0xb3, 0x20, 0x21, 0x21, 0x43, 0x8c, 0x33, 0x4b, 0x76, 0x1c, 0xda, 0xb6,
	0x42, 0x72, 0x69, 0x4d, 0xec, 0x16, 0x3b, 0xca, 0x78, 0xda, 0x9a, 0x9f, 0x24, 0x7b, 0xe4, 0x1d,
	0x78, 0x01, 0xae, 0x3c, 0x08, 0x6f, 0xc2, 0x89, 0x3b, 0x12, 0x37, 0x54, 0x3d, 0xdd, 0xe3, 0x1f,
	0x0c, 0x37, 0x4e, 0xec, 0x69, 0xaa, 0xea, 0xfb, 0xaa, 0xbb, 0xea, 0xab, 0x9a, 0x19, 0x78, 0x27,
	0xf4, 0xd3, 0xd0, 0x4f, 0x79, 0x2c, 0xa2, 0x3b, 0x11, 0x1d, 0x2f, 0x22, 0x99, 0x48, 0xba, 0x33,
	0x9b, 0x29, 0xa3, 0xf9, 0x74, 0x2a, 0xe7, 0x73, 0x19, 0x6a, 0x94, 0x2f, 0xe4, 0xad, 0xe1, 0xb4,
	0x1f, 0x80, 0x98, 0xd4, 0xe8, 0x2e, 0x43, 0xe8, 0x87, 0x50, 0x5a, 0x78, 0x7e, 0x6c, 0x17, 0x5a,
	0xc5, 0x4e, 0xbd, 0x7b, 0x70, 0xac, 0x8f, 0x39, 0x36, 0xa7, 0x18, 0xe2, 0x85, 0xe7, 0x33, 0xc5,
	0xa4, 0x5d, 0x28, 0x25, 0xaf, 0x17, 0xc2, 0x2e, 0xb6, 0xac, 0xce, 0x5e, 0xf7, 0x30, 0xcf, 0xd0,
	0x47, 0x8b, 0x30, 0x9d, 0xf3, 0x0b, 0x4c, 0x19, 0xbf, 0x5e, 0x08, 0xa6, 0xb8, 0xed, 0x5f, 0xac,
	0xfc, 0xea, 0x34, 0x16, 0x11, 0xbf, 0xf1, 0x83, 0x80, 0x3e, 0x82, 0x72, 0x3c, 0x95, 0x91, 0xb0,
	0xad, 0x96, 0xd5, 0x29, 0xb3, 0xcc, 0xa1, 0x6d, 0x68, 0x4c, 0x65, 0x1a, 0x26, 0xfc, 0x95, 0x17,
	0xf3, 0xd0, 0x4f, 0xed, 0x82, 0x42, 0xeb, 0x2a, 0x78, 0xe6, 0xc5, 0xae, 0x9f, 0xd2, 0x16, 0xec,
	0x66, 0x9c, 0x50, 0x2a, 0x4a, 0x51, 0x51, 0x40, 0xc5, 0x5c, 0x89, 0x8c, 0x67, 0x50, 0xcb, 0x18,
	0xf7, 0x7e, 0x68, 0x97, 0x14, 0x5c, 0x55, 0x81, 0x4b, 0x3f, 0xa4, 0xef, 0x41, 0x46, 0xe5, 0x81,
	0x8c, 0x13, 0xbb, 0xac, 0xd0, 0x8c, 0x7e, 0x2e, 0xe3, 0xa4, 0xfd, 0xab, 0x05, 0x4f, 0x56, 0x74,
	0x9a, 0x89, 0xf8, 0x96, 0xcb, 0x45, 0xe2, 0xcb, 0x90, 0xda, 0xb0, 0x33, 0xf7, 0xc3, 0x49, 0x2c,
	0x22, 0x5d, 0xb4, 0x71, 0x15, 0xe2, 0x3d, 0x28, 0xa4, 0xa0, 0x91, 0xcc, 0xa5, 0x07, 0x50, 0x9b,
	0x7b, 0x0f, 0x5f, 0xfb, 0xd1, 0x34, 0x10, 0xba, 0xd2, 0x65, 0x00, 0xd1, 0x57, 0x5e, 0xfc, 0x4d,
	0x20, 0xef, 0x45, 0xa4, 0x0a, 0xad, 0xb2, 0x65, 0x80, 0x7e, 0x01, 0xd5, 0x1b, 0x2f, 0xbc, 0x65,
	0x69, 0x20, 0x54, 0x9d, 0x7b, 0xdd, 0xd6, 0x56, 0xbd, 0x91, 0x24, 0x22, 0x1e, 0xa5, 0x81, 0x60,
	0x79, 0x46, 0xfb, 0xe7, 0x22, 0x50, 0xcd, 0x9a, 0x06, 0xbe, 0x08, 0x13, 0xd5, 0x0b, 0x7d, 0x02,
	0x15, 0x7c, 0x3a, 0x33, 0xdd, 0x83, 0xf6, 0xe8, 0x21, 0x00, 0x5a, 0x6e, 0x3a, 0xbf, 0xd1, 0x5d,
	0xd4, 0xd8, 0x4a, 0x04, 0xf1, 0x1f, 0xbc, 0xb9, 0xd0, 0xb8, 0xd6, 0x7c, 0x19, 0xc1, 0x73, 0x23,
	0x29, 0xe7, 0xce, 0x4c, 0x0b, 0xae, 0x3d, 0xfa, 0x1c, 0x2a, 0x71, 0xe2, 0x25, 0x69, 0xac, 0x5b,
	0x78, 0x7f, 0x6b, 0x0b, 0x4a, 0x66, 0xe4, 0x09, 0xa6, 0xe9, 0xa8, 0x4d, 0xe0, 0xc5, 0x38, 0x32,
	0x11, 0xd9, 0x95, 0x96, 0xd5, 0x69, 0xb0, 0x65, 0x80, 0x36, 0xa1, 0x3a, 0x55, 0x1a, 0xba, 0xd2,
	0xae, 0xea, 0x09, 0x6b, 0x1f, 0x57, 0x4b, 0xde, 0x63, 0x16, 0xa8, 0xac, 0xcc, 0xc1, 0x06, 0xa6,
	0x69, 0x14, 0x7d, 0xa5, 0xc4, 0xb2, 0xeb, 0x0a, 0x5a, 0x89, 0xd0, 0x2f, 0x33, 0x01, 0x86, 0x6a,
	0xd6, 0xf6, 0x6e, 0xcb, 0xea, 0xd4, 0xff, 0x5e, 0xec, 0xc6, 0x4a, 0xb0, 0x95, 0x14, 0x7a, 0x02,
	0x65, 0x5c, 0xef, 0xd8, 0x6e, 0xa8, 0xb7, 0xe9, 0xd9, 0x66, 0xae, 0x9e, 0x02, 0x72, 0x58, 0xc6,
	0x6c, 0xff, 0x69, 0x6d, 0xce, 0x08, 0xe3, 0xa8, 0x25, 0x3e, 0xf5, 0x8c, 0x1a, 0x4c, 0x7b, 0xf4,
	0x03, 0x28, 0xe1, 0xbb, 0xa3, 0xd4, 0xa8, 0x77, 0x9f, 0x6e, 0x5e, 0x90, 0xbf, 0x5c, 0x4c, 0xd1,
	0x56, 0xa4, 0xdf, 0xf9, 0x17, 0xe9, 0x55, 0xd6, 0xba, 0xf4, 0x4d, 0xa8, 0xfa, 0xf1, 0x30, 0x0c,
	0xfc, 0x50, 0x28, 0x0d, 0xab, 0x2c, 0xf7, 0x51, 0x5c, 0x3f, 0x9c, 0x89, 0x07, 0xa5, 0x50, 0x99,
	0x65, 0x0e, 0x3d, 0x81, 0x8a, 0xfa, 0x50, 0x60, 0xf3, 0x5b, 0x6b, 0xcb, 0x3f, 0x25, 0x4c, 0x13,
	0xdb, 0xbf, 0x17, 0x60, 0x7f, 0x43, 0xd5, 0x37, 0xcb, 0xf9, 0x5f, 0x2d, 0x67, 0xfb, 0x0f, 0x6b,
	0x4d, 0xed, 0xff, 0xcf, 0x9a, 0xfd, 0xb8, 0xde, 0x38, 0x0e, 0x7e, 0x65, 0x1d, 0xac, 0xb5, 0x75,
	0x68, 0x42, 0x15, 0x2d, 0xfc, 0x75, 0xe9, 0xef, 0x78, 0xee, 0xe3, 0xc4, 0xd1, 0x3e, 0x17, 0x77,
	0x22, 0x30, 0x1f, 0xf2, 0x3c, 0x60, 0xd0, 0xb1, 0x9f, 0x04, 0x42, 0xed, 0x58, 0x8d, 0x2d, 0x03,
	0x47, 0xbf, 0x59, 0xf0, 0x78, 0xeb, 0x0f, 0x92, 0x02, 0x54, 0xdc, 0x21, 0x77, 0x9d, 0x09, 0xb1,
	0x68, 0x0d, 0xca, 0xae, 0x33, 0xe1, 0x27, 0xa4, 0x60, 0xcc, 0x2e, 0x29, 0x1a, 0xf3, 0x23, 0x52,
	0x32, 0xe6, 0xc7, 0xa4, 0x6c, 0xcc, 0x4f, 0x48, 0xc5, 0x98, 0x9f, 0x92, 0x1d, 0x63, 0x3e, 0x27,
	0x55, 0x63, 0x7e, 0x46, 0x6a, 0xc6, 0xfc, 0x9c, 0x00, 0xad, 0xc3, 0x0e, 0x9a, 0x78, 0x5f, 0x1d,
	0x9d, 0x2b, 0xc7, 0x55, 0xce, 0x2e, 0x3a, 0x2f, 0xb4, 0xd3, 0xa0, 0xfb, 0x50, 0xbf, 0x9c, 0xf0,
	0xef, 0x9d, 0x5e, 0x56, 0xda, 0x1e, 0x06, 0x30, 0xef, 0xfa, 0xac, 0xc7, 0x4f, 0x7b, 0x2e, 0xd9,
	0xa7, 0x8f, 0x80, 0x60, 0xe0, 0xca, 0xe1, 0x63, 0x64, 0x9d, 0x0f, 0xdd, 0x01, 0x21, 0x47, 0x49,
	0xfe, 0xeb, 0xdc, 0x18, 0x39, 0x3d, 0x00, 0x1b, 0xf9, 0x93, 0x51, 0x9f, 0xf1, 0xd1, 0xb8, 0x37,
	0x9e, 0x8c, 0xb8, 0x3b, 0xe4, 0xac, 0xdf, 0x3b, 0xbd, 0x22, 0xd6, 0x36, 0xd4, 0x19, 0x69, 0xb4,
	0x40, 0x0f, 0xa1, 0xb9, 0x05, 0x1d, 0xf4, 0x5e, 0xbe, 0x74, 0xdc, 0x01, 0x29, 0x1e, 0xfd, 0x64,
	0xad, 0x5f, 0xbb, 0x7c, 0x5d, 0x4d, 0xea, 0x69, 0x7f, 0xf4, 0xad, 0x49, 0xbd, 0xec, 0x39, 0x63,
	0xde, 0x77, 0xc7, 0x7d, 0xb6, 0xbc, 0x78, 0x15, 0xcf, 0xcb, 0x2a, 0x6c, 0x43, 0xf3, 0xb2, 0x8a,
	0xdb, 0xce, 0x5e, 0x29, 0xab, 0x74, 0x34, 0x82, 0x77, 0xff, 0xe1, 0x27, 0x8d, 0x72, 0x9e, 0x3a,
	0xee, 0x80, 0x5f, 0x9f, 0x4d, 0x7a, 0xee, 0x80, 0x58, 0xf4, 0x31, 0xbc, 0x3d, 0x9a, 0x38, 0xfc,
	0x85, 0xc3, 0xaf, 0x27, 0x43, 0x13, 0x2e, 0x50, 0x02, 0xbb, 0xdf, 0x39, 0xbd, 0x25, 0xb1, 0x78,
	0xf1, 0xd6, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x30, 0xed, 0x92, 0x00, 0x0a, 0x00, 0x00,
}
