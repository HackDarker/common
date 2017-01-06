// Code generated by protoc-gen-go.
// source: zjh_serever.proto
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

// 三张牌的类型
type ZjhEnum_ZJHTYPE int32

const (
	ZjhEnum_ZJHTYPE_ZJHTYPE_GAOPAI   ZjhEnum_ZJHTYPE = 1
	ZjhEnum_ZJHTYPE_ZJHTYPE_DUIZI    ZjhEnum_ZJHTYPE = 2
	ZjhEnum_ZJHTYPE_ZJHTYPE_LIANZI   ZjhEnum_ZJHTYPE = 3
	ZjhEnum_ZJHTYPE_ZJHTYPE_QING     ZjhEnum_ZJHTYPE = 4
	ZjhEnum_ZJHTYPE_ZJHTYPE_QINGLIAN ZjhEnum_ZJHTYPE = 5
	ZjhEnum_ZJHTYPE_ZJHTYPE_BAOZI    ZjhEnum_ZJHTYPE = 6
)

var ZjhEnum_ZJHTYPE_name = map[int32]string{
	1: "ZJHTYPE_GAOPAI",
	2: "ZJHTYPE_DUIZI",
	3: "ZJHTYPE_LIANZI",
	4: "ZJHTYPE_QING",
	5: "ZJHTYPE_QINGLIAN",
	6: "ZJHTYPE_BAOZI",
}
var ZjhEnum_ZJHTYPE_value = map[string]int32{
	"ZJHTYPE_GAOPAI":   1,
	"ZJHTYPE_DUIZI":    2,
	"ZJHTYPE_LIANZI":   3,
	"ZJHTYPE_QING":     4,
	"ZJHTYPE_QINGLIAN": 5,
	"ZJHTYPE_BAOZI":    6,
}

func (x ZjhEnum_ZJHTYPE) Enum() *ZjhEnum_ZJHTYPE {
	p := new(ZjhEnum_ZJHTYPE)
	*p = x
	return p
}
func (x ZjhEnum_ZJHTYPE) String() string {
	return proto.EnumName(ZjhEnum_ZJHTYPE_name, int32(x))
}
func (x *ZjhEnum_ZJHTYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnum_ZJHTYPE_value, data, "ZjhEnum_ZJHTYPE")
	if err != nil {
		return err
	}
	*x = ZjhEnum_ZJHTYPE(value)
	return nil
}
func (ZjhEnum_ZJHTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor23, []int{0} }

// 用户游戏状态
type ZjhEnumUserStatus int32

const (
	ZjhEnumUserStatus_zjh_S_GAMING   ZjhEnumUserStatus = 1
	ZjhEnumUserStatus_zjh_S_STANDING ZjhEnumUserStatus = 2
	ZjhEnumUserStatus_zjh_S_SITED    ZjhEnumUserStatus = 3
	ZjhEnumUserStatus_zjh_S_READY    ZjhEnumUserStatus = 4
	ZjhEnumUserStatus_zjh_S_OFFLINE  ZjhEnumUserStatus = 5
)

var ZjhEnumUserStatus_name = map[int32]string{
	1: "zjh_S_GAMING",
	2: "zjh_S_STANDING",
	3: "zjh_S_SITED",
	4: "zjh_S_READY",
	5: "zjh_S_OFFLINE",
}
var ZjhEnumUserStatus_value = map[string]int32{
	"zjh_S_GAMING":   1,
	"zjh_S_STANDING": 2,
	"zjh_S_SITED":    3,
	"zjh_S_READY":    4,
	"zjh_S_OFFLINE":  5,
}

func (x ZjhEnumUserStatus) Enum() *ZjhEnumUserStatus {
	p := new(ZjhEnumUserStatus)
	*p = x
	return p
}
func (x ZjhEnumUserStatus) String() string {
	return proto.EnumName(ZjhEnumUserStatus_name, int32(x))
}
func (x *ZjhEnumUserStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumUserStatus_value, data, "ZjhEnumUserStatus")
	if err != nil {
		return err
	}
	*x = ZjhEnumUserStatus(value)
	return nil
}
func (ZjhEnumUserStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor23, []int{1} }

// 打出去的牌
type ZjhSrvPoker struct {
	KeyValue         []int32              `protobuf:"varint,1,rep,name=keyValue" json:"keyValue,omitempty"`
	Pais             []*CommonSrvPokerPai `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *ZjhEnum_ZJHTYPE     `protobuf:"varint,3,opt,name=type,enum=ddproto.ZjhEnum_ZJHTYPE" json:"type,omitempty"`
	UserId           *uint32              `protobuf:"varint,9,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhSrvPoker) Reset()                    { *m = ZjhSrvPoker{} }
func (m *ZjhSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvPoker) ProtoMessage()               {}
func (*ZjhSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{0} }

func (m *ZjhSrvPoker) GetKeyValue() []int32 {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *ZjhSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *ZjhSrvPoker) GetType() ZjhEnum_ZJHTYPE {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ZjhEnum_ZJHTYPE_ZJHTYPE_GAOPAI
}

func (m *ZjhSrvPoker) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 账单
type ZjhSrvBill struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ZjhSrvBill) Reset()                    { *m = ZjhSrvBill{} }
func (m *ZjhSrvBill) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvBill) ProtoMessage()               {}
func (*ZjhSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{1} }

// 用户的游戏数据
type ZjhSrv_GameData struct {
	Pai              *ZjhSrvPoker `protobuf:"bytes,1,opt,name=pai" json:"pai,omitempty"`
	Bill             *int64       `protobuf:"varint,2,opt,name=bill" json:"bill,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhSrv_GameData) Reset()                    { *m = ZjhSrv_GameData{} }
func (m *ZjhSrv_GameData) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrv_GameData) ProtoMessage()               {}
func (*ZjhSrv_GameData) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{2} }

func (m *ZjhSrv_GameData) GetPai() *ZjhSrvPoker {
	if m != nil {
		return m.Pai
	}
	return nil
}

func (m *ZjhSrv_GameData) GetBill() int64 {
	if m != nil && m.Bill != nil {
		return *m.Bill
	}
	return 0
}

// 用户属性
type ZjhSrvUser struct {
	UserId           *uint32            `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Coin             *int64             `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	RoomId           *int32             `protobuf:"varint,3,opt,name=roomId" json:"roomId,omitempty"`
	DeskId           *int32             `protobuf:"varint,4,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32             `protobuf:"varint,5,opt,name=gameNumber" json:"gameNumber,omitempty"`
	Data             *ZjhSrv_GameData   `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	Status           *ZjhEnumUserStatus `protobuf:"varint,7,opt,name=status,enum=ddproto.ZjhEnumUserStatus" json:"status,omitempty"`
	IsDiscard        *bool              `protobuf:"varint,8,opt,name=isDiscard" json:"isDiscard,omitempty"`
	IsWatch          *bool              `protobuf:"varint,9,opt,name=isWatch" json:"isWatch,omitempty"`
	IsOnline         *bool              `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	IsLost           *bool              `protobuf:"varint,11,opt,name=isLost" json:"isLost,omitempty"`
	Index            *int32             `protobuf:"varint,12,opt,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ZjhSrvUser) Reset()                    { *m = ZjhSrvUser{} }
func (m *ZjhSrvUser) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvUser) ProtoMessage()               {}
func (*ZjhSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{3} }

func (m *ZjhSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ZjhSrvUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *ZjhSrvUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhSrvUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ZjhSrvUser) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *ZjhSrvUser) GetData() *ZjhSrv_GameData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ZjhSrvUser) GetStatus() ZjhEnumUserStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ZjhEnumUserStatus_zjh_S_GAMING
}

func (m *ZjhSrvUser) GetIsDiscard() bool {
	if m != nil && m.IsDiscard != nil {
		return *m.IsDiscard
	}
	return false
}

func (m *ZjhSrvUser) GetIsWatch() bool {
	if m != nil && m.IsWatch != nil {
		return *m.IsWatch
	}
	return false
}

func (m *ZjhSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *ZjhSrvUser) GetIsLost() bool {
	if m != nil && m.IsLost != nil {
		return *m.IsLost
	}
	return false
}

func (m *ZjhSrvUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

// desk 的信息
type ZjhSrvDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32               `protobuf:"varint,2,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32               `protobuf:"varint,3,opt,name=roomId" json:"roomId,omitempty"`
	AllPokers        []*CommonSrvPokerPai `protobuf:"bytes,4,rep,name=allPokers" json:"allPokers,omitempty"`
	LastUser         *uint32              `protobuf:"varint,5,opt,name=lastUser" json:"lastUser,omitempty"`
	LastWiner        *uint32              `protobuf:"varint,6,opt,name=lastWiner" json:"lastWiner,omitempty"`
	IsGamming        *bool                `protobuf:"varint,7,opt,name=isGamming" json:"isGamming,omitempty"`
	CuurBaseValue    *int64               `protobuf:"varint,8,opt,name=cuurBaseValue" json:"cuurBaseValue,omitempty"`
	MinUser          *int32               `protobuf:"varint,9,opt,name=minUser" json:"minUser,omitempty"`
	BaseValue        *int64               `protobuf:"varint,10,opt,name=baseValue" json:"baseValue,omitempty"`
	MaxValue         *int64               `protobuf:"varint,11,opt,name=maxValue" json:"maxValue,omitempty"`
	GameStatus       *int32               `protobuf:"varint,12,opt,name=gameStatus" json:"gameStatus,omitempty"`
	GamerNum         *int32               `protobuf:"varint,13,opt,name=gamerNum" json:"gamerNum,omitempty"`
	CircleNo         *int32               `protobuf:"varint,14,opt,name=circleNo" json:"circleNo,omitempty"`
	XuepinBaseValue  *int64               `protobuf:"varint,15,opt,name=xuepinBaseValue" json:"xuepinBaseValue,omitempty"`
	CoinCount        *int64               `protobuf:"varint,16,opt,name=coinCount" json:"coinCount,omitempty"`
	XuepinStartIndex *int32               `protobuf:"varint,17,opt,name=xuepinStartIndex" json:"xuepinStartIndex,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhSrvDesk) Reset()                    { *m = ZjhSrvDesk{} }
func (m *ZjhSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvDesk) ProtoMessage()               {}
func (*ZjhSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{4} }

func (m *ZjhSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ZjhSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *ZjhSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhSrvDesk) GetAllPokers() []*CommonSrvPokerPai {
	if m != nil {
		return m.AllPokers
	}
	return nil
}

func (m *ZjhSrvDesk) GetLastUser() uint32 {
	if m != nil && m.LastUser != nil {
		return *m.LastUser
	}
	return 0
}

func (m *ZjhSrvDesk) GetLastWiner() uint32 {
	if m != nil && m.LastWiner != nil {
		return *m.LastWiner
	}
	return 0
}

func (m *ZjhSrvDesk) GetIsGamming() bool {
	if m != nil && m.IsGamming != nil {
		return *m.IsGamming
	}
	return false
}

func (m *ZjhSrvDesk) GetCuurBaseValue() int64 {
	if m != nil && m.CuurBaseValue != nil {
		return *m.CuurBaseValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetMinUser() int32 {
	if m != nil && m.MinUser != nil {
		return *m.MinUser
	}
	return 0
}

func (m *ZjhSrvDesk) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetMaxValue() int64 {
	if m != nil && m.MaxValue != nil {
		return *m.MaxValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *ZjhSrvDesk) GetGamerNum() int32 {
	if m != nil && m.GamerNum != nil {
		return *m.GamerNum
	}
	return 0
}

func (m *ZjhSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *ZjhSrvDesk) GetXuepinBaseValue() int64 {
	if m != nil && m.XuepinBaseValue != nil {
		return *m.XuepinBaseValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetCoinCount() int64 {
	if m != nil && m.CoinCount != nil {
		return *m.CoinCount
	}
	return 0
}

func (m *ZjhSrvDesk) GetXuepinStartIndex() int32 {
	if m != nil && m.XuepinStartIndex != nil {
		return *m.XuepinStartIndex
	}
	return 0
}

// room 的信息
type ZjhSrvRoom struct {
	RoomId           *int32  `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	BaseValue        *int64  `protobuf:"varint,5,opt,name=baseValue" json:"baseValue,omitempty"`
	MaxValue         *int64  `protobuf:"varint,6,opt,name=maxValue" json:"maxValue,omitempty"`
	XuepinBaseValue  *int64  `protobuf:"varint,7,opt,name=xuepinBaseValue" json:"xuepinBaseValue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ZjhSrvRoom) Reset()                    { *m = ZjhSrvRoom{} }
func (m *ZjhSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvRoom) ProtoMessage()               {}
func (*ZjhSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{5} }

func (m *ZjhSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *ZjhSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *ZjhSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func (m *ZjhSrvRoom) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *ZjhSrvRoom) GetMaxValue() int64 {
	if m != nil && m.MaxValue != nil {
		return *m.MaxValue
	}
	return 0
}

func (m *ZjhSrvRoom) GetXuepinBaseValue() int64 {
	if m != nil && m.XuepinBaseValue != nil {
		return *m.XuepinBaseValue
	}
	return 0
}

func init() {
	proto.RegisterType((*ZjhSrvPoker)(nil), "ddproto.zjh_srv_poker")
	proto.RegisterType((*ZjhSrvBill)(nil), "ddproto.zjh_srv_bill")
	proto.RegisterType((*ZjhSrv_GameData)(nil), "ddproto.zjh_srv_GameData")
	proto.RegisterType((*ZjhSrvUser)(nil), "ddproto.zjh_srv_user")
	proto.RegisterType((*ZjhSrvDesk)(nil), "ddproto.zjh_srv_desk")
	proto.RegisterType((*ZjhSrvRoom)(nil), "ddproto.zjh_srv_room")
	proto.RegisterEnum("ddproto.ZjhEnum_ZJHTYPE", ZjhEnum_ZJHTYPE_name, ZjhEnum_ZJHTYPE_value)
	proto.RegisterEnum("ddproto.ZjhEnumUserStatus", ZjhEnumUserStatus_name, ZjhEnumUserStatus_value)
}

var fileDescriptor23 = []byte{
	// 796 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x5b, 0x4e, 0xe3, 0x48,
	0x14, 0x1d, 0x27, 0xce, 0xab, 0x42, 0x82, 0xa9, 0x41, 0xa8, 0x40, 0x68, 0x14, 0x45, 0xf3, 0x61,
	0x21, 0x0d, 0x1a, 0xa1, 0xf9, 0x9a, 0xbf, 0x30, 0x09, 0x19, 0x8f, 0x32, 0x49, 0xda, 0x09, 0x8d,
	0xc8, 0x4f, 0x54, 0xc4, 0x25, 0xa8, 0xc6, 0x8f, 0xc8, 0x65, 0x47, 0xd0, 0x2b, 0xe8, 0x85, 0xf4,
	0x22, 0x7a, 0x3d, 0xbd, 0x83, 0xde, 0x41, 0xeb, 0x56, 0xf9, 0x15, 0x42, 0xa3, 0xfe, 0xab, 0x73,
	0xee, 0xfb, 0xdc, 0x6b, 0xa3, 0x83, 0x8f, 0x1f, 0x1e, 0x96, 0x82, 0x85, 0x6c, 0xc3, 0xc2, 0xf3,
	0x75, 0x18, 0x44, 0x01, 0xae, 0x39, 0x8e, 0x7c, 0x9c, 0x1c, 0xaf, 0x02, 0xcf, 0x0b, 0x7c, 0x30,
	0x6f, 0x58, 0xb8, 0x5c, 0x07, 0x8f, 0xa9, 0x4f, 0xf7, 0xb3, 0x86, 0x5a, 0x32, 0x32, 0xdc, 0x28,
	0x1e, 0x9f, 0xa0, 0xfa, 0x23, 0x7b, 0x7e, 0x4f, 0xdd, 0x98, 0x11, 0xad, 0x53, 0x36, 0x2b, 0x76,
	0x86, 0xf1, 0x9f, 0x48, 0x5f, 0x53, 0x2e, 0x48, 0xa9, 0x53, 0x36, 0x9b, 0x17, 0xa7, 0xe7, 0x49,
	0x81, 0xf3, 0x34, 0x7f, 0x9a, 0x64, 0x4a, 0xb9, 0x2d, 0x3d, 0xf1, 0x1f, 0x48, 0x8f, 0x9e, 0xd7,
	0x8c, 0x94, 0x3b, 0x9a, 0xd9, 0xbe, 0x38, 0xce, 0x22, 0xa0, 0x26, 0xf3, 0x63, 0x6f, 0xb9, 0xf8,
	0xef, 0xdf, 0xf9, 0xed, 0x74, 0x60, 0x4b, 0x37, 0x7c, 0x84, 0xaa, 0xb1, 0x60, 0xa1, 0xe5, 0x90,
	0x46, 0x47, 0x33, 0x5b, 0x76, 0x82, 0xba, 0x6d, 0xb4, 0x97, 0x76, 0x79, 0xc7, 0x5d, 0xb7, 0x3b,
	0x45, 0x46, 0x8a, 0x87, 0xd4, 0x63, 0x7d, 0x1a, 0x51, 0x6c, 0xa2, 0xf2, 0x9a, 0x72, 0xa2, 0x75,
	0x34, 0xb3, 0x79, 0x71, 0xb4, 0x55, 0x29, 0x6b, 0xcc, 0x06, 0x17, 0x8c, 0x91, 0x0e, 0x59, 0x48,
	0xa9, 0xa3, 0x99, 0x65, 0x5b, 0xbe, 0xbb, 0xdf, 0x4a, 0x79, 0x09, 0x28, 0x5a, 0x68, 0x45, 0x2b,
	0xb6, 0x02, 0xc1, 0xab, 0x80, 0xfb, 0x69, 0x30, 0xbc, 0xc1, 0x37, 0x0c, 0x02, 0xcf, 0x72, 0xe4,
	0x9c, 0x15, 0x3b, 0x41, 0xc0, 0x3b, 0x4c, 0x3c, 0x5a, 0x0e, 0xd1, 0x15, 0xaf, 0x10, 0xfe, 0x0d,
	0xa1, 0x7b, 0xea, 0xb1, 0x71, 0xec, 0xdd, 0xb1, 0x90, 0x54, 0xa4, 0xad, 0xc0, 0x80, 0x6a, 0x0e,
	0x8d, 0x28, 0xa9, 0xca, 0x59, 0x8e, 0x77, 0x66, 0x49, 0x67, 0xb6, 0xa5, 0x1b, 0xfe, 0x0b, 0x55,
	0x45, 0x44, 0xa3, 0x58, 0x90, 0x9a, 0x94, 0xf9, 0x74, 0x57, 0x66, 0x68, 0x7e, 0x26, 0x7d, 0xec,
	0xc4, 0x17, 0x9f, 0xa2, 0x06, 0x17, 0x7d, 0x2e, 0x56, 0x34, 0x74, 0x48, 0xbd, 0xa3, 0x99, 0x75,
	0x3b, 0x27, 0x30, 0x41, 0x35, 0x2e, 0x6e, 0x68, 0xb4, 0x7a, 0x90, 0xab, 0xa8, 0xdb, 0x29, 0x84,
	0x03, 0xe1, 0x62, 0xe2, 0xbb, 0xdc, 0x67, 0x04, 0x49, 0x53, 0x86, 0x61, 0x60, 0x2e, 0x46, 0x81,
	0x88, 0x48, 0x53, 0x5a, 0x12, 0x84, 0x0f, 0x51, 0x85, 0xfb, 0x0e, 0x7b, 0x22, 0x7b, 0x72, 0x56,
	0x05, 0xba, 0x5f, 0xf4, 0x5c, 0x73, 0x50, 0xa6, 0xa0, 0x97, 0xf6, 0x86, 0x5e, 0xa5, 0x1d, 0xbd,
	0x7e, 0xa4, 0xff, 0xdf, 0xa8, 0x41, 0x5d, 0x77, 0x0a, 0x9b, 0x17, 0x44, 0xff, 0x89, 0xa3, 0xcd,
	0xdd, 0x61, 0x4c, 0x97, 0x8a, 0xe8, 0x5a, 0x24, 0x1b, 0x6a, 0xd9, 0x19, 0x06, 0xe9, 0xe0, 0x7d,
	0xc3, 0x7d, 0x16, 0xca, 0x25, 0xb5, 0xec, 0x9c, 0x50, 0xc2, 0x0e, 0xa9, 0xe7, 0x71, 0xff, 0x5e,
	0x6e, 0x44, 0x0a, 0x9b, 0x10, 0xf8, 0x77, 0xd4, 0x5a, 0xc5, 0x71, 0x78, 0x49, 0x05, 0x53, 0x1f,
	0x59, 0x5d, 0x1e, 0xd2, 0x36, 0x09, 0xf2, 0x7b, 0xdc, 0x97, 0xc5, 0x1b, 0x72, 0xa4, 0x14, 0x42,
	0xf6, 0xbb, 0x2c, 0x16, 0xc9, 0xd8, 0x9c, 0x80, 0xae, 0x3d, 0xfa, 0xa4, 0x8c, 0x4d, 0x69, 0xcc,
	0x70, 0xaa, 0xa2, 0x3a, 0x83, 0x64, 0x13, 0x05, 0x06, 0x62, 0x01, 0x85, 0xe3, 0xd8, 0x23, 0x2d,
	0x69, 0xcd, 0x30, 0xd8, 0x56, 0x3c, 0x5c, 0xb9, 0x6c, 0x1c, 0x90, 0xb6, 0xb2, 0xa5, 0x18, 0x9b,
	0x68, 0xff, 0x29, 0x66, 0x6b, 0xee, 0xe7, 0x33, 0xed, 0xcb, 0xd2, 0x2f, 0x69, 0xe8, 0x1d, 0xbe,
	0x97, 0x7f, 0x82, 0xd8, 0x8f, 0x88, 0xa1, 0x7a, 0xcf, 0x08, 0x7c, 0x86, 0x0c, 0x15, 0x30, 0x8b,
	0x68, 0x18, 0x59, 0xf2, 0x5e, 0x0e, 0x64, 0xad, 0x1d, 0xbe, 0xfb, 0x55, 0xcb, 0x4f, 0x07, 0x96,
	0x5d, 0x38, 0x01, 0x6d, 0xeb, 0x04, 0x4e, 0x50, 0x1d, 0x5e, 0x73, 0xf8, 0x09, 0xa9, 0xc3, 0xc9,
	0x30, 0xb4, 0x03, 0xef, 0x11, 0xdb, 0x30, 0x37, 0xb9, 0x9c, 0x9c, 0x48, 0xad, 0x73, 0x1e, 0xb9,
	0x4c, 0x7e, 0xbf, 0x0d, 0x3b, 0x27, 0xb6, 0xd7, 0x50, 0x79, 0x6b, 0x0d, 0xd5, 0x17, 0x6b, 0x78,
	0x45, 0xae, 0xda, 0xab, 0x72, 0x9d, 0x7d, 0xd2, 0xd4, 0x6f, 0xae, 0xf8, 0xa3, 0xc4, 0x18, 0xb5,
	0x93, 0xe7, 0x72, 0xd8, 0x9b, 0x4c, 0x7b, 0x96, 0xa1, 0xe1, 0x03, 0xd4, 0x4a, 0xb9, 0xfe, 0xb5,
	0xb5, 0xb0, 0x8c, 0x52, 0xd1, 0x6d, 0x64, 0xf5, 0xc6, 0x0b, 0xcb, 0x28, 0x63, 0x03, 0xed, 0xa5,
	0xdc, 0x3b, 0x6b, 0x3c, 0x34, 0x74, 0x7c, 0x88, 0x8c, 0x22, 0x03, 0x9e, 0x46, 0xa5, 0x98, 0xee,
	0xb2, 0x37, 0x59, 0x58, 0x46, 0xf5, 0x6c, 0x8d, 0x7e, 0x7d, 0xe5, 0x5f, 0x02, 0x19, 0x81, 0x9e,
	0x2d, 0x87, 0xbd, 0xff, 0x21, 0xa3, 0x06, 0x75, 0x15, 0x33, 0x9b, 0xf7, 0xc6, 0x7d, 0xe0, 0x4a,
	0x78, 0x1f, 0x35, 0x13, 0xce, 0x9a, 0x0f, 0xfa, 0x46, 0x39, 0x27, 0xec, 0x41, 0xaf, 0x7f, 0x6b,
	0xe8, 0x50, 0x51, 0x11, 0x93, 0xab, 0xab, 0x91, 0x35, 0x1e, 0x18, 0x95, 0xe9, 0x2f, 0xdf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0x4a, 0xae, 0x36, 0xd3, 0x06, 0x00, 0x00,
}
