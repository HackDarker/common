// Code generated by protoc-gen-go.
// source: common_mj.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ErrorCode int32

const (
	ErrorCode_MJ_EC_SUCCESS ErrorCode = 0
	// -101   -200	游戏异常
	ErrorCode_MJ_EC_CREATE_DESK_DIAMOND_NOTENOUGH ErrorCode = -101
	ErrorCode_MJ_EC_CREATE_DESK_USER_NOTFOUND     ErrorCode = -102
	ErrorCode_MJ_EC_INTO_DESK_NOTFOUND            ErrorCode = -103
	ErrorCode_MJ_EC_GAME_READY_REPEAT             ErrorCode = -110
	ErrorCode_MJ_EC_GAME_READY_CHIP_NOT_ENOUGH    ErrorCode = -111
)

var ErrorCode_name = map[int32]string{
	0:    "MJ_EC_SUCCESS",
	-101: "MJ_EC_CREATE_DESK_DIAMOND_NOTENOUGH",
	-102: "MJ_EC_CREATE_DESK_USER_NOTFOUND",
	-103: "MJ_EC_INTO_DESK_NOTFOUND",
	-110: "MJ_EC_GAME_READY_REPEAT",
	-111: "MJ_EC_GAME_READY_CHIP_NOT_ENOUGH",
}
var ErrorCode_value = map[string]int32{
	"MJ_EC_SUCCESS":                       0,
	"MJ_EC_CREATE_DESK_DIAMOND_NOTENOUGH": -101,
	"MJ_EC_CREATE_DESK_USER_NOTFOUND":     -102,
	"MJ_EC_INTO_DESK_NOTFOUND":            -103,
	"MJ_EC_GAME_READY_REPEAT":             -110,
	"MJ_EC_GAME_READY_CHIP_NOT_ENOUGH":    -111,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}
func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// 麻将花色
type MjEnumColor int32

const (
	MjEnumColor_WAN  MjEnumColor = 1
	MjEnumColor_TIAO MjEnumColor = 2
	MjEnumColor_TONG MjEnumColor = 3
)

var MjEnumColor_name = map[int32]string{
	1: "WAN",
	2: "TIAO",
	3: "TONG",
}
var MjEnumColor_value = map[string]int32{
	"WAN":  1,
	"TIAO": 2,
	"TONG": 3,
}

func (x MjEnumColor) Enum() *MjEnumColor {
	p := new(MjEnumColor)
	*p = x
	return p
}
func (x MjEnumColor) String() string {
	return proto.EnumName(MjEnumColor_name, int32(x))
}
func (x *MjEnumColor) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjEnumColor_value, data, "MjEnumColor")
	if err != nil {
		return err
	}
	*x = MjEnumColor(value)
	return nil
}
func (MjEnumColor) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

// 杠的类型
type MjEnumGangType int32

const (
	MjEnumGangType_G_MINGGANG MjEnumGangType = 1
	MjEnumGangType_G_BAGANG   MjEnumGangType = 2
	MjEnumGangType_G_ANGANG   MjEnumGangType = 3
)

var MjEnumGangType_name = map[int32]string{
	1: "G_MINGGANG",
	2: "G_BAGANG",
	3: "G_ANGANG",
}
var MjEnumGangType_value = map[string]int32{
	"G_MINGGANG": 1,
	"G_BAGANG":   2,
	"G_ANGANG":   3,
}

func (x MjEnumGangType) Enum() *MjEnumGangType {
	p := new(MjEnumGangType)
	*p = x
	return p
}
func (x MjEnumGangType) String() string {
	return proto.EnumName(MjEnumGangType_name, int32(x))
}
func (x *MjEnumGangType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjEnumGangType_value, data, "MjEnumGangType")
	if err != nil {
		return err
	}
	*x = MjEnumGangType(value)
	return nil
}
func (MjEnumGangType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

// 胡牌的类型
type MjEnumHuType int32

const (
	MjEnumHuType_H_NORMAL MjEnumHuType = 0
	// 附加番数(特定操作)
	MjEnumHuType_H_TianHu            MjEnumHuType = 1
	MjEnumHuType_H_DiHu              MjEnumHuType = 2
	MjEnumHuType_H_GangShangHua      MjEnumHuType = 3
	MjEnumHuType_H_GangShangPao      MjEnumHuType = 4
	MjEnumHuType_H_QiangGang         MjEnumHuType = 5
	MjEnumHuType_H_HaiDiLao          MjEnumHuType = 6
	MjEnumHuType_H_HaiDiPao          MjEnumHuType = 7
	MjEnumHuType_H_HaidiGangShangHua MjEnumHuType = 8
	MjEnumHuType_H_HaidiGangShangPao MjEnumHuType = 9
	//
	MjEnumHuType_H_JinGouDiao MjEnumHuType = 10
	// 可选附加玩法
	MjEnumHuType_H_ZiMoJiaFan        MjEnumHuType = 11
	MjEnumHuType_H_ZiMoJiaDi         MjEnumHuType = 12
	MjEnumHuType_H_changsha_qishouhu MjEnumHuType = 13
)

var MjEnumHuType_name = map[int32]string{
	0:  "H_NORMAL",
	1:  "H_TianHu",
	2:  "H_DiHu",
	3:  "H_GangShangHua",
	4:  "H_GangShangPao",
	5:  "H_QiangGang",
	6:  "H_HaiDiLao",
	7:  "H_HaiDiPao",
	8:  "H_HaidiGangShangHua",
	9:  "H_HaidiGangShangPao",
	10: "H_JinGouDiao",
	11: "H_ZiMoJiaFan",
	12: "H_ZiMoJiaDi",
	13: "H_changsha_qishouhu",
}
var MjEnumHuType_value = map[string]int32{
	"H_NORMAL":            0,
	"H_TianHu":            1,
	"H_DiHu":              2,
	"H_GangShangHua":      3,
	"H_GangShangPao":      4,
	"H_QiangGang":         5,
	"H_HaiDiLao":          6,
	"H_HaiDiPao":          7,
	"H_HaidiGangShangHua": 8,
	"H_HaidiGangShangPao": 9,
	"H_JinGouDiao":        10,
	"H_ZiMoJiaFan":        11,
	"H_ZiMoJiaDi":         12,
	"H_changsha_qishouhu": 13,
}

func (x MjEnumHuType) Enum() *MjEnumHuType {
	p := new(MjEnumHuType)
	*p = x
	return p
}
func (x MjEnumHuType) String() string {
	return proto.EnumName(MjEnumHuType_name, int32(x))
}
func (x *MjEnumHuType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjEnumHuType_value, data, "MjEnumHuType")
	if err != nil {
		return err
	}
	*x = MjEnumHuType(value)
	return nil
}
func (MjEnumHuType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

// 碰杠类型(客户端显示用)
type MjEnumComposeCardType int32

const (
	MjEnumComposeCardType_C_MINGGANG MjEnumComposeCardType = 1
	MjEnumComposeCardType_C_BAGANG   MjEnumComposeCardType = 2
	MjEnumComposeCardType_C_ANGANG   MjEnumComposeCardType = 3
	MjEnumComposeCardType_C_PENG     MjEnumComposeCardType = 4
	MjEnumComposeCardType_C_CHI      MjEnumComposeCardType = 5
)

var MjEnumComposeCardType_name = map[int32]string{
	1: "C_MINGGANG",
	2: "C_BAGANG",
	3: "C_ANGANG",
	4: "C_PENG",
	5: "C_CHI",
}
var MjEnumComposeCardType_value = map[string]int32{
	"C_MINGGANG": 1,
	"C_BAGANG":   2,
	"C_ANGANG":   3,
	"C_PENG":     4,
	"C_CHI":      5,
}

func (x MjEnumComposeCardType) Enum() *MjEnumComposeCardType {
	p := new(MjEnumComposeCardType)
	*p = x
	return p
}
func (x MjEnumComposeCardType) String() string {
	return proto.EnumName(MjEnumComposeCardType_name, int32(x))
}
func (x *MjEnumComposeCardType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjEnumComposeCardType_value, data, "MjEnumComposeCardType")
	if err != nil {
		return err
	}
	*x = MjEnumComposeCardType(value)
	return nil
}
func (MjEnumComposeCardType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

type MjEnumPaiType int32

const (
	MjEnumPaiType_H_DuiDuiHu   MjEnumPaiType = 1
	MjEnumPaiType_H_QingYiSe   MjEnumPaiType = 2
	MjEnumPaiType_H_QiDui      MjEnumPaiType = 3
	MjEnumPaiType_H_DaiYaoJiu  MjEnumPaiType = 4
	MjEnumPaiType_H_LongQiDui  MjEnumPaiType = 5
	MjEnumPaiType_H_JiangDui   MjEnumPaiType = 6
	MjEnumPaiType_H_MenQing    MjEnumPaiType = 7
	MjEnumPaiType_H_ZhongZhang MjEnumPaiType = 8
	// (内江麻将)
	MjEnumPaiType_H_KaErTiao MjEnumPaiType = 9
	// (德阳麻将)
	MjEnumPaiType_H_JiaXin5       MjEnumPaiType = 10
	MjEnumPaiType_H_QingLongQiDui MjEnumPaiType = 11
	MjEnumPaiType_H_QingQiDui     MjEnumPaiType = 12
	MjEnumPaiType_H_PingHu        MjEnumPaiType = 13
	// 长沙麻将
	MjEnumPaiType_H_CHANGSHA_DAXISI                  MjEnumPaiType = 14
	MjEnumPaiType_H_CHANGSHA_BANBANHU                MjEnumPaiType = 15
	MjEnumPaiType_H_CHANGSHA_QUEYISE                 MjEnumPaiType = 16
	MjEnumPaiType_H_CHANGSHA_LIULIUSHUN              MjEnumPaiType = 17
	MjEnumPaiType_H_CHANGSHA_PINGHU                  MjEnumPaiType = 18
	MjEnumPaiType_H_CHANGSHA_PENGPENGHU              MjEnumPaiType = 19
	MjEnumPaiType_H_CHANGSHA_JIANGJIANGHU            MjEnumPaiType = 20
	MjEnumPaiType_H_CHANGSHA_QINGYISE                MjEnumPaiType = 21
	MjEnumPaiType_H_CHANGSHA_QIXIAODUI               MjEnumPaiType = 22
	MjEnumPaiType_H_CHANGSHA_QIXIAODUI_HAOHUA        MjEnumPaiType = 23
	MjEnumPaiType_H_CHANGSHA_QIXIAODUI_HAOHUA_DOUBLE MjEnumPaiType = 24
	MjEnumPaiType_H_CHANGSHA_QUANQIUREN              MjEnumPaiType = 25
)

var MjEnumPaiType_name = map[int32]string{
	1:  "H_DuiDuiHu",
	2:  "H_QingYiSe",
	3:  "H_QiDui",
	4:  "H_DaiYaoJiu",
	5:  "H_LongQiDui",
	6:  "H_JiangDui",
	7:  "H_MenQing",
	8:  "H_ZhongZhang",
	9:  "H_KaErTiao",
	10: "H_JiaXin5",
	11: "H_QingLongQiDui",
	12: "H_QingQiDui",
	13: "H_PingHu",
	14: "H_CHANGSHA_DAXISI",
	15: "H_CHANGSHA_BANBANHU",
	16: "H_CHANGSHA_QUEYISE",
	17: "H_CHANGSHA_LIULIUSHUN",
	18: "H_CHANGSHA_PINGHU",
	19: "H_CHANGSHA_PENGPENGHU",
	20: "H_CHANGSHA_JIANGJIANGHU",
	21: "H_CHANGSHA_QINGYISE",
	22: "H_CHANGSHA_QIXIAODUI",
	23: "H_CHANGSHA_QIXIAODUI_HAOHUA",
	24: "H_CHANGSHA_QIXIAODUI_HAOHUA_DOUBLE",
	25: "H_CHANGSHA_QUANQIUREN",
}
var MjEnumPaiType_value = map[string]int32{
	"H_DuiDuiHu":                         1,
	"H_QingYiSe":                         2,
	"H_QiDui":                            3,
	"H_DaiYaoJiu":                        4,
	"H_LongQiDui":                        5,
	"H_JiangDui":                         6,
	"H_MenQing":                          7,
	"H_ZhongZhang":                       8,
	"H_KaErTiao":                         9,
	"H_JiaXin5":                          10,
	"H_QingLongQiDui":                    11,
	"H_QingQiDui":                        12,
	"H_PingHu":                           13,
	"H_CHANGSHA_DAXISI":                  14,
	"H_CHANGSHA_BANBANHU":                15,
	"H_CHANGSHA_QUEYISE":                 16,
	"H_CHANGSHA_LIULIUSHUN":              17,
	"H_CHANGSHA_PINGHU":                  18,
	"H_CHANGSHA_PENGPENGHU":              19,
	"H_CHANGSHA_JIANGJIANGHU":            20,
	"H_CHANGSHA_QINGYISE":                21,
	"H_CHANGSHA_QIXIAODUI":               22,
	"H_CHANGSHA_QIXIAODUI_HAOHUA":        23,
	"H_CHANGSHA_QIXIAODUI_HAOHUA_DOUBLE": 24,
	"H_CHANGSHA_QUANQIUREN":              25,
}

func (x MjEnumPaiType) Enum() *MjEnumPaiType {
	p := new(MjEnumPaiType)
	*p = x
	return p
}
func (x MjEnumPaiType) String() string {
	return proto.EnumName(MjEnumPaiType_name, int32(x))
}
func (x *MjEnumPaiType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjEnumPaiType_value, data, "MjEnumPaiType")
	if err != nil {
		return err
	}
	*x = MjEnumPaiType(value)
	return nil
}
func (MjEnumPaiType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

type MjEnumUserGameStatus int32

const (
	MjEnumUserGameStatus_MJ_U_INIT     MjEnumUserGameStatus = 0
	MjEnumUserGameStatus_MJ_U_DINGQUE  MjEnumUserGameStatus = 1
	MjEnumUserGameStatus_MJ_U_EXCHANGE MjEnumUserGameStatus = 2
	MjEnumUserGameStatus_MJ_U_PLAYING  MjEnumUserGameStatus = 3
	MjEnumUserGameStatus_MJ_U_FINISH   MjEnumUserGameStatus = 4
)

var MjEnumUserGameStatus_name = map[int32]string{
	0: "MJ_U_INIT",
	1: "MJ_U_DINGQUE",
	2: "MJ_U_EXCHANGE",
	3: "MJ_U_PLAYING",
	4: "MJ_U_FINISH",
}
var MjEnumUserGameStatus_value = map[string]int32{
	"MJ_U_INIT":     0,
	"MJ_U_DINGQUE":  1,
	"MJ_U_EXCHANGE": 2,
	"MJ_U_PLAYING":  3,
	"MJ_U_FINISH":   4,
}

func (x MjEnumUserGameStatus) Enum() *MjEnumUserGameStatus {
	p := new(MjEnumUserGameStatus)
	*p = x
	return p
}
func (x MjEnumUserGameStatus) String() string {
	return proto.EnumName(MjEnumUserGameStatus_name, int32(x))
}
func (x *MjEnumUserGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjEnumUserGameStatus_value, data, "MjEnumUserGameStatus")
	if err != nil {
		return err
	}
	*x = MjEnumUserGameStatus(value)
	return nil
}
func (MjEnumUserGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

type MjEnumDeskGameStatus int32

const (
	MjEnumDeskGameStatus_MJ_INIT     MjEnumDeskGameStatus = 0
	MjEnumDeskGameStatus_MJ_FAPAI    MjEnumDeskGameStatus = 1
	MjEnumDeskGameStatus_MJ_EXCHANGE MjEnumDeskGameStatus = 2
	MjEnumDeskGameStatus_MJ_DINGQUE  MjEnumDeskGameStatus = 3
	MjEnumDeskGameStatus_MJ_PLAYING  MjEnumDeskGameStatus = 4
	MjEnumDeskGameStatus_MJ_FINISH   MjEnumDeskGameStatus = 5
)

var MjEnumDeskGameStatus_name = map[int32]string{
	0: "MJ_INIT",
	1: "MJ_FAPAI",
	2: "MJ_EXCHANGE",
	3: "MJ_DINGQUE",
	4: "MJ_PLAYING",
	5: "MJ_FINISH",
}
var MjEnumDeskGameStatus_value = map[string]int32{
	"MJ_INIT":     0,
	"MJ_FAPAI":    1,
	"MJ_EXCHANGE": 2,
	"MJ_DINGQUE":  3,
	"MJ_PLAYING":  4,
	"MJ_FINISH":   5,
}

func (x MjEnumDeskGameStatus) Enum() *MjEnumDeskGameStatus {
	p := new(MjEnumDeskGameStatus)
	*p = x
	return p
}
func (x MjEnumDeskGameStatus) String() string {
	return proto.EnumName(MjEnumDeskGameStatus_name, int32(x))
}
func (x *MjEnumDeskGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjEnumDeskGameStatus_value, data, "MjEnumDeskGameStatus")
	if err != nil {
		return err
	}
	*x = MjEnumDeskGameStatus(value)
	return nil
}
func (MjEnumDeskGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

// 房间类型信息：包含房间类型和对应的局数、封顶、玩法等信息
// 房间类型枚举
type MJRoomType int32

const (
	MJRoomType_roomType_xueZhanDaoDi      MJRoomType = 0
	MJRoomType_roomType_sanRenLiangFang   MJRoomType = 1
	MJRoomType_roomType_siRenLiangFang    MJRoomType = 2
	MJRoomType_roomType_deYangMaJiang     MJRoomType = 3
	MJRoomType_roomType_daoDaoHu          MJRoomType = 4
	MJRoomType_roomType_xueLiuChengHe     MJRoomType = 5
	MJRoomType_roomType_liangRenLiangFang MJRoomType = 6
	MJRoomType_roomType_liangRenSanFang   MJRoomType = 7
	MJRoomType_roomType_sanRenSanFang     MJRoomType = 8
	MJRoomType_roomType_changSha          MJRoomType = 9
	MJRoomType_roomType_zhuoxiazi         MJRoomType = 10
	MJRoomType_roomType_mj_baishan        MJRoomType = 12
)

var MJRoomType_name = map[int32]string{
	0:  "roomType_xueZhanDaoDi",
	1:  "roomType_sanRenLiangFang",
	2:  "roomType_siRenLiangFang",
	3:  "roomType_deYangMaJiang",
	4:  "roomType_daoDaoHu",
	5:  "roomType_xueLiuChengHe",
	6:  "roomType_liangRenLiangFang",
	7:  "roomType_liangRenSanFang",
	8:  "roomType_sanRenSanFang",
	9:  "roomType_changSha",
	10: "roomType_zhuoxiazi",
	12: "roomType_mj_baishan",
}
var MJRoomType_value = map[string]int32{
	"roomType_xueZhanDaoDi":      0,
	"roomType_sanRenLiangFang":   1,
	"roomType_siRenLiangFang":    2,
	"roomType_deYangMaJiang":     3,
	"roomType_daoDaoHu":          4,
	"roomType_xueLiuChengHe":     5,
	"roomType_liangRenLiangFang": 6,
	"roomType_liangRenSanFang":   7,
	"roomType_sanRenSanFang":     8,
	"roomType_changSha":          9,
	"roomType_zhuoxiazi":         10,
	"roomType_mj_baishan":        12,
}

func (x MJRoomType) Enum() *MJRoomType {
	p := new(MJRoomType)
	*p = x
	return p
}
func (x MJRoomType) String() string {
	return proto.EnumName(MJRoomType_name, int32(x))
}
func (x *MJRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MJRoomType_value, data, "MJRoomType")
	if err != nil {
		return err
	}
	*x = MJRoomType(value)
	return nil
}
func (MJRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

// 麻将牌
type CardInfo struct {
	Type             *int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Value            *int32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Id               *int32 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CardInfo) Reset()                    { *m = CardInfo{} }
func (m *CardInfo) String() string            { return proto.CompactTextString(m) }
func (*CardInfo) ProtoMessage()               {}
func (*CardInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *CardInfo) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *CardInfo) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *CardInfo) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type RoomTypeInfo struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RoomTypeInfo) Reset()                    { *m = RoomTypeInfo{} }
func (m *RoomTypeInfo) String() string            { return proto.CompactTextString(m) }
func (*RoomTypeInfo) ProtoMessage()               {}
func (*RoomTypeInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type ComposeCard struct {
	Value            *int32  `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	ChiValue         []int32 `protobuf:"varint,3,rep,name=chiValue" json:"chiValue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ComposeCard) Reset()                    { *m = ComposeCard{} }
func (m *ComposeCard) String() string            { return proto.CompactTextString(m) }
func (*ComposeCard) ProtoMessage()               {}
func (*ComposeCard) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ComposeCard) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *ComposeCard) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *ComposeCard) GetChiValue() []int32 {
	if m != nil {
		return m.ChiValue
	}
	return nil
}

type PlayerCard struct {
	HandCard         []*CardInfo    `protobuf:"bytes,1,rep,name=handCard" json:"handCard,omitempty"`
	ComposeCard      []*ComposeCard `protobuf:"bytes,2,rep,name=composeCard" json:"composeCard,omitempty"`
	OutCard          []*CardInfo    `protobuf:"bytes,3,rep,name=outCard" json:"outCard,omitempty"`
	HuCard           []*CardInfo    `protobuf:"bytes,4,rep,name=huCard" json:"huCard,omitempty"`
	UserId           *uint32        `protobuf:"varint,5,opt,name=UserId" json:"UserId,omitempty"`
	HandCardCount    *int32         `protobuf:"varint,6,opt,name=handCardCount" json:"handCardCount,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *PlayerCard) Reset()                    { *m = PlayerCard{} }
func (m *PlayerCard) String() string            { return proto.CompactTextString(m) }
func (*PlayerCard) ProtoMessage()               {}
func (*PlayerCard) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *PlayerCard) GetHandCard() []*CardInfo {
	if m != nil {
		return m.HandCard
	}
	return nil
}

func (m *PlayerCard) GetComposeCard() []*ComposeCard {
	if m != nil {
		return m.ComposeCard
	}
	return nil
}

func (m *PlayerCard) GetOutCard() []*CardInfo {
	if m != nil {
		return m.OutCard
	}
	return nil
}

func (m *PlayerCard) GetHuCard() []*CardInfo {
	if m != nil {
		return m.HuCard
	}
	return nil
}

func (m *PlayerCard) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PlayerCard) GetHandCardCount() int32 {
	if m != nil && m.HandCardCount != nil {
		return *m.HandCardCount
	}
	return 0
}

func init() {
	proto.RegisterType((*CardInfo)(nil), "ddproto.CardInfo")
	proto.RegisterType((*RoomTypeInfo)(nil), "ddproto.RoomTypeInfo")
	proto.RegisterType((*ComposeCard)(nil), "ddproto.ComposeCard")
	proto.RegisterType((*PlayerCard)(nil), "ddproto.PlayerCard")
	proto.RegisterEnum("ddproto.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("ddproto.MjEnumColor", MjEnumColor_name, MjEnumColor_value)
	proto.RegisterEnum("ddproto.MjEnumGangType", MjEnumGangType_name, MjEnumGangType_value)
	proto.RegisterEnum("ddproto.MjEnumHuType", MjEnumHuType_name, MjEnumHuType_value)
	proto.RegisterEnum("ddproto.MjEnumComposeCardType", MjEnumComposeCardType_name, MjEnumComposeCardType_value)
	proto.RegisterEnum("ddproto.MjEnumPaiType", MjEnumPaiType_name, MjEnumPaiType_value)
	proto.RegisterEnum("ddproto.MjEnumUserGameStatus", MjEnumUserGameStatus_name, MjEnumUserGameStatus_value)
	proto.RegisterEnum("ddproto.MjEnumDeskGameStatus", MjEnumDeskGameStatus_name, MjEnumDeskGameStatus_value)
	proto.RegisterEnum("ddproto.MJRoomType", MJRoomType_name, MJRoomType_value)
}

var fileDescriptor4 = []byte{
	// 1200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x6e, 0xdb, 0xc6,
	0x12, 0xb6, 0xa8, 0x5f, 0x8f, 0x7e, 0xbc, 0xde, 0x38, 0xb6, 0x92, 0x1c, 0x9c, 0x18, 0x3e, 0x39,
	0x07, 0x3e, 0x6a, 0x12, 0x14, 0x05, 0xda, 0xcb, 0x02, 0x34, 0x49, 0x73, 0xa9, 0x48, 0x94, 0x44,
	0x8a, 0xad, 0x1d, 0x14, 0x58, 0x6c, 0x2d, 0x56, 0xdc, 0xd4, 0xe2, 0xba, 0x92, 0x59, 0x24, 0x79,
	0x83, 0x5e, 0xa6, 0x77, 0x6d, 0x9f, 0xa1, 0x4f, 0xd3, 0xb7, 0xe9, 0x4d, 0x8b, 0x5d, 0xfe, 0x48,
	0x72, 0x92, 0x1a, 0xb0, 0xb1, 0x33, 0xdf, 0x37, 0xdf, 0xcc, 0xce, 0xcc, 0xd2, 0xb0, 0x77, 0x25,
	0x16, 0x0b, 0x11, 0xd3, 0xc5, 0xab, 0xe7, 0x37, 0x4b, 0x71, 0x2b, 0x70, 0x7d, 0x36, 0x53, 0x87,
	0x13, 0x13, 0x1a, 0x06, 0x5b, 0xce, 0x9c, 0xf8, 0x3b, 0x81, 0x31, 0x54, 0x6e, 0xdf, 0xdc, 0x84,
	0xdd, 0xd2, 0x71, 0xe9, 0xb4, 0xea, 0xa9, 0x33, 0x3e, 0x80, 0xea, 0x8f, 0xec, 0x3a, 0x09, 0xbb,
	0x9a, 0x72, 0xa6, 0x06, 0xee, 0x80, 0xc6, 0x67, 0xdd, 0xb2, 0x72, 0x69, 0x7c, 0x76, 0xd2, 0x81,
	0x96, 0x27, 0xc4, 0x62, 0xfa, 0xe6, 0x26, 0x94, 0x4a, 0x27, 0x3e, 0x34, 0x0d, 0xb1, 0xb8, 0x11,
	0xab, 0x50, 0x8a, 0xaf, 0x45, 0x4a, 0x9b, 0x22, 0x79, 0x3a, 0x6d, 0x23, 0xdd, 0x43, 0x68, 0x5c,
	0x45, 0xfc, 0x2b, 0x45, 0x2e, 0x1f, 0x97, 0x4f, 0xab, 0x5e, 0x61, 0x9f, 0xfc, 0xa4, 0x01, 0x8c,
	0xaf, 0xd9, 0x9b, 0x70, 0xa9, 0x44, 0x9f, 0x41, 0x23, 0x62, 0xf1, 0x4c, 0x9e, 0xbb, 0xa5, 0xe3,
	0xf2, 0x69, 0xf3, 0xb3, 0xfd, 0xe7, 0xd9, 0xad, 0x9e, 0xe7, 0x57, 0xf2, 0x0a, 0x0a, 0xfe, 0x02,
	0x9a, 0x57, 0xeb, 0x92, 0xba, 0x9a, 0x8a, 0x38, 0x58, 0x47, 0xac, 0x31, 0x6f, 0x93, 0x88, 0x3f,
	0x81, 0xba, 0x48, 0x6e, 0x55, 0x4c, 0xf9, 0x63, 0x59, 0x72, 0x06, 0xfe, 0x3f, 0xd4, 0xa2, 0x44,
	0x71, 0x2b, 0x1f, 0xe3, 0x66, 0x04, 0x7c, 0x08, 0xb5, 0x60, 0x15, 0x2e, 0x9d, 0x59, 0xb7, 0x7a,
	0x5c, 0x3a, 0x6d, 0x7b, 0x99, 0x85, 0x9f, 0x40, 0x3b, 0xaf, 0xd9, 0x10, 0x49, 0x7c, 0xdb, 0xad,
	0xa9, 0xf6, 0x6c, 0x3b, 0x7b, 0x7f, 0x96, 0x60, 0xd7, 0x5a, 0x2e, 0xc5, 0xd2, 0x10, 0xb3, 0x10,
	0xef, 0x43, 0x7b, 0xd8, 0xa7, 0x96, 0x41, 0xfd, 0xc0, 0x30, 0x2c, 0xdf, 0x47, 0x3b, 0xf8, 0x53,
	0xf8, 0x4f, 0xea, 0x32, 0x3c, 0x4b, 0x9f, 0x5a, 0xd4, 0xb4, 0xfc, 0x17, 0xd4, 0x74, 0xf4, 0xe1,
	0xc8, 0x35, 0xa9, 0x3b, 0x9a, 0x5a, 0xee, 0x28, 0xb0, 0x09, 0xfa, 0xed, 0xaf, 0xec, 0xa7, 0x84,
	0x9f, 0xc2, 0xe3, 0xf7, 0x23, 0x02, 0xdf, 0xf2, 0x24, 0xfd, 0x7c, 0x14, 0xb8, 0x26, 0xfa, 0x75,
	0xcd, 0xfe, 0x2f, 0x74, 0x53, 0xb6, 0xe3, 0x4e, 0x47, 0x29, 0xb7, 0xa0, 0xfd, 0xb2, 0xa6, 0x3d,
	0x81, 0xa3, 0x94, 0x66, 0xeb, 0x43, 0x8b, 0x7a, 0x96, 0x6e, 0x5e, 0x52, 0xcf, 0x1a, 0x5b, 0xfa,
	0x14, 0xfd, 0xbc, 0x66, 0x3d, 0x83, 0xe3, 0xf7, 0x58, 0x06, 0x71, 0xc6, 0x52, 0x92, 0x66, 0x95,
	0xbe, 0x2b, 0xe8, 0xbd, 0xa7, 0xd0, 0x5e, 0xbc, 0xa2, 0x61, 0x9c, 0x2c, 0xe8, 0x95, 0xb8, 0x16,
	0x4b, 0x5c, 0x87, 0xf2, 0xd7, 0xba, 0x8b, 0x4a, 0xb8, 0x01, 0x95, 0xa9, 0xa3, 0x8f, 0x90, 0xa6,
	0x4e, 0x23, 0xd7, 0x46, 0xe5, 0xde, 0x97, 0x80, 0x72, 0xf6, 0x9c, 0xc5, 0x73, 0xb9, 0xa3, 0xb8,
	0x03, 0x60, 0xd3, 0xa1, 0xe3, 0xda, 0xb6, 0xee, 0xda, 0xa8, 0x84, 0x5b, 0xd0, 0xb0, 0xe9, 0x99,
	0xae, 0x2c, 0x2d, 0xb5, 0x74, 0x57, 0x59, 0xe5, 0xde, 0x3b, 0x0d, 0x3a, 0xb9, 0x40, 0x94, 0xa8,
	0xf0, 0x16, 0x34, 0x08, 0x75, 0x47, 0xde, 0x50, 0x1f, 0xa0, 0x9d, 0xd4, 0x9a, 0x72, 0x16, 0x93,
	0x04, 0x95, 0x30, 0x40, 0x8d, 0x50, 0x93, 0x93, 0x04, 0x69, 0x18, 0x43, 0x87, 0x50, 0x9b, 0xc5,
	0x73, 0x3f, 0x62, 0xf1, 0x9c, 0x24, 0x0c, 0x95, 0xef, 0xf8, 0xc6, 0x4c, 0xa0, 0x0a, 0xde, 0x83,
	0x26, 0xa1, 0x13, 0xce, 0xe2, 0xb9, 0x04, 0x50, 0x55, 0xd6, 0x47, 0x28, 0x61, 0xdc, 0xe4, 0x03,
	0x26, 0x50, 0x6d, 0xc3, 0x96, 0x01, 0x75, 0x7c, 0x04, 0xf7, 0x94, 0x3d, 0xe3, 0x5b, 0xea, 0x8d,
	0x0f, 0x01, 0x32, 0x62, 0x17, 0x23, 0x68, 0x11, 0xda, 0xe7, 0xb1, 0x2d, 0x12, 0x93, 0x33, 0x81,
	0x20, 0xf5, 0xbc, 0xe4, 0x43, 0xd1, 0xe7, 0xec, 0x9c, 0xc5, 0xa8, 0x99, 0x96, 0x91, 0x79, 0x4c,
	0x8e, 0x5a, 0xa9, 0xda, 0x95, 0x14, 0x59, 0x45, 0x8c, 0xfe, 0xc0, 0x57, 0x91, 0x48, 0xa2, 0x04,
	0xb5, 0x7b, 0xdf, 0xc0, 0xd1, 0x7a, 0x02, 0xc5, 0x5b, 0xc9, 0x5b, 0x6b, 0xdc, 0x69, 0xad, 0xb1,
	0xd5, 0x5a, 0xa3, 0x68, 0xad, 0xec, 0x95, 0x41, 0xc7, 0x96, 0x6b, 0xa3, 0x0a, 0xde, 0x85, 0xaa,
	0x21, 0x87, 0x8e, 0xaa, 0xbd, 0xdf, 0x2b, 0xb0, 0x97, 0xcb, 0xdf, 0x30, 0x9e, 0xcb, 0x12, 0x6a,
	0x26, 0xdc, 0x4c, 0xb8, 0x6a, 0xb3, 0xb2, 0x27, 0x3c, 0x9e, 0x5f, 0x72, 0x3f, 0x44, 0x1a, 0x6e,
	0x42, 0x5d, 0xda, 0x66, 0xc2, 0x51, 0x39, 0xbd, 0x88, 0xc9, 0xf8, 0x25, 0x13, 0x7d, 0x9e, 0xe4,
	0x0d, 0x1e, 0x88, 0x78, 0x9e, 0x32, 0xb2, 0x06, 0xf7, 0x65, 0xc7, 0xa5, 0x5d, 0xc3, 0x6d, 0xd8,
	0x25, 0x74, 0x18, 0xc6, 0x52, 0x11, 0xd5, 0xb3, 0xde, 0x44, 0x22, 0x9e, 0xbf, 0x94, 0xb7, 0x47,
	0x8d, 0x34, 0xe0, 0x05, 0xb3, 0x96, 0x53, 0xae, 0xfa, 0xa9, 0x02, 0xfa, 0x9c, 0x5d, 0xf0, 0xf8,
	0x73, 0x04, 0xf8, 0x1e, 0xec, 0xa5, 0xe5, 0xac, 0x93, 0x34, 0xf3, 0xb1, 0xe6, 0x8e, 0x56, 0xba,
	0x29, 0x63, 0x2e, 0x87, 0x85, 0xda, 0xf8, 0x3e, 0xec, 0x13, 0x6a, 0x10, 0xdd, 0xb5, 0x7d, 0xa2,
	0x53, 0x53, 0xbf, 0x70, 0x7c, 0x07, 0x75, 0xd2, 0xa6, 0x17, 0xee, 0x33, 0xdd, 0x3d, 0xd3, 0x5d,
	0x12, 0xa0, 0x3d, 0x7c, 0x08, 0x78, 0x03, 0x98, 0x04, 0xd6, 0xa5, 0xe3, 0x5b, 0x08, 0xe1, 0x07,
	0x70, 0x7f, 0xc3, 0x3f, 0x70, 0x82, 0x81, 0x13, 0xf8, 0x24, 0x70, 0xd1, 0xfe, 0x9d, 0x14, 0x63,
	0xc7, 0xb5, 0x49, 0x80, 0xf0, 0x9d, 0x08, 0x39, 0x00, 0xf9, 0x4b, 0x02, 0x74, 0x0f, 0x3f, 0x82,
	0xa3, 0x0d, 0xa8, 0xef, 0xe8, 0xae, 0xad, 0xfe, 0x90, 0x00, 0x1d, 0xdc, 0x29, 0x6d, 0xe2, 0xb8,
	0xb6, 0x2a, 0xe1, 0x3e, 0xee, 0xc2, 0xc1, 0x16, 0x70, 0xe1, 0xe8, 0x23, 0x33, 0x70, 0xd0, 0x21,
	0x7e, 0x0c, 0x8f, 0x3e, 0x84, 0x50, 0xa2, 0x8f, 0x48, 0xa0, 0xa3, 0x23, 0xfc, 0x3f, 0x38, 0xf9,
	0x07, 0x02, 0x35, 0x47, 0xc1, 0xd9, 0xc0, 0x42, 0xdd, 0x3b, 0x35, 0x4f, 0x02, 0xdd, 0x9d, 0x38,
	0x81, 0x67, 0xb9, 0xe8, 0x41, 0x4f, 0xc0, 0x61, 0xbe, 0x2e, 0xc9, 0x2a, 0x5c, 0xda, 0x6c, 0x11,
	0xfa, 0xb7, 0xec, 0x36, 0x59, 0xc9, 0x29, 0x0d, 0xfb, 0x34, 0xa0, 0x8e, 0xeb, 0x4c, 0xd1, 0x8e,
	0x1c, 0xab, 0x32, 0x4d, 0xc7, 0xb5, 0x27, 0x81, 0x85, 0x4a, 0xd9, 0x97, 0x33, 0xa0, 0xd6, 0x85,
	0x92, 0xb6, 0x90, 0x56, 0x90, 0xc6, 0x03, 0xfd, 0xd2, 0x51, 0x6b, 0xba, 0x07, 0x4d, 0xe5, 0x39,
	0x77, 0x5c, 0xc7, 0x27, 0xa8, 0xd2, 0x5b, 0xad, 0x13, 0xce, 0xc2, 0xd5, 0xf7, 0x1b, 0x09, 0x9b,
	0x50, 0x1f, 0xf6, 0xf3, 0x74, 0x2d, 0x68, 0x0c, 0xfb, 0xf4, 0x5c, 0x1f, 0xeb, 0x0e, 0x2a, 0x65,
	0x2a, 0x1b, 0x89, 0x3a, 0x00, 0xc3, 0x7e, 0x51, 0x4b, 0x39, 0xb3, 0xf3, 0xb4, 0x95, 0xac, 0xf8,
	0x2c, 0x69, 0xb5, 0xf7, 0x87, 0x26, 0xf1, 0xfc, 0xdf, 0xac, 0xec, 0xc7, 0x32, 0x3b, 0xd3, 0xd7,
	0x49, 0x28, 0xf7, 0xd4, 0x64, 0xc2, 0xe4, 0x68, 0x07, 0xff, 0x0b, 0xba, 0x05, 0xb4, 0x62, 0xb1,
	0x17, 0xc6, 0x03, 0xb9, 0xe9, 0xe7, 0x72, 0x93, 0x4b, 0x72, 0xc2, 0x6b, 0x94, 0x6f, 0x81, 0x1a,
	0x7e, 0x08, 0x87, 0x05, 0x38, 0x0b, 0x2f, 0x59, 0x3c, 0x1f, 0x32, 0xf5, 0x4c, 0x50, 0x59, 0x2e,
	0xd3, 0x1a, 0x63, 0xc2, 0x64, 0x82, 0xc8, 0xb7, 0xb5, 0x19, 0xf2, 0x3a, 0x09, 0x07, 0x3c, 0x31,
	0xa2, 0x30, 0x9e, 0x93, 0x10, 0x55, 0xf1, 0xbf, 0xe1, 0x61, 0x81, 0x5d, 0x4b, 0x99, 0xad, 0x74,
	0xb5, 0xad, 0x4a, 0x73, 0xdc, 0x67, 0xb1, 0x42, 0xeb, 0x5b, 0xca, 0xe9, 0x3d, 0x72, 0xac, 0xb1,
	0x55, 0x8c, 0xfa, 0x42, 0xf9, 0x11, 0x43, 0xbb, 0xf2, 0x8d, 0x14, 0xee, 0xb7, 0x51, 0x22, 0x5e,
	0x73, 0xf6, 0x96, 0x23, 0x90, 0x9b, 0x5b, 0xf8, 0x17, 0xaf, 0xe8, 0xb7, 0x8c, 0xaf, 0x22, 0x16,
	0xa3, 0xd6, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x44, 0xc5, 0x44, 0x1a, 0x09, 0x00, 0x00,
}