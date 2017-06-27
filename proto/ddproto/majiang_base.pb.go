// Code generated by protoc-gen-go.
// source: majiang_base.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of CardInfo from common_mj.proto

// Ignoring public import of RoomTypeInfo from common_mj.proto

// Ignoring public import of PlayOptions from common_mj.proto

// Ignoring public import of ChangShaPlayOptions from common_mj.proto

// Ignoring public import of BaiShanPlayOptions from common_mj.proto

// Ignoring public import of ZhuanZhuanPlayOptions from common_mj.proto

// Ignoring public import of ComposeCard from common_mj.proto

// Ignoring public import of PlayerCard from common_mj.proto

// Ignoring public import of ErrorCode from common_mj.proto

// Ignoring public import of MahjongColor from common_mj.proto

// Ignoring public import of GangType from common_mj.proto

// Ignoring public import of HuType from common_mj.proto

// Ignoring public import of ComposeCardType from common_mj.proto

// Ignoring public import of PaiType from common_mj.proto

// Ignoring public import of MJUserGameStatus from common_mj.proto

// Ignoring public import of DeskGameStatus from common_mj.proto

// Ignoring public import of MJRoomType from common_mj.proto

type MjEnumProtoId int32

const (
	MjEnumProtoId_PID_QUICK_CONN                     MjEnumProtoId = 1
	MjEnumProtoId_PID_QUICK_CONN_ACK                 MjEnumProtoId = 2
	MjEnumProtoId_PID_GAME_LOGIN                     MjEnumProtoId = 3
	MjEnumProtoId_PID_GAME_LOGIN_ACK                 MjEnumProtoId = 4
	MjEnumProtoId_PID_CREATEROOM                     MjEnumProtoId = 5
	MjEnumProtoId_PID_CREATEROOM_ACK                 MjEnumProtoId = 6
	MjEnumProtoId_PID_ENTER_ROOM                     MjEnumProtoId = 7
	MjEnumProtoId_PID_ENTER_ROOM_ACK                 MjEnumProtoId = 8
	MjEnumProtoId_PID_SEND_GAMEINFO                  MjEnumProtoId = 9
	MjEnumProtoId_PID_READY                          MjEnumProtoId = 10
	MjEnumProtoId_PID_READY_ACK                      MjEnumProtoId = 11
	MjEnumProtoId_PID_EXCHANGECARDS                  MjEnumProtoId = 12
	MjEnumProtoId_PID_EXCHANGECARDS_ACK              MjEnumProtoId = 13
	MjEnumProtoId_PID_DINGQUE                        MjEnumProtoId = 14
	MjEnumProtoId_PID_OPENING                        MjEnumProtoId = 15
	MjEnumProtoId_PID_DEAL_CARDS                     MjEnumProtoId = 16
	MjEnumProtoId_PID_GET_IN_CARD                    MjEnumProtoId = 17
	MjEnumProtoId_PID_SEND_OUT_CARD                  MjEnumProtoId = 18
	MjEnumProtoId_PID_SEND_OUT_CARD_ACK              MjEnumProtoId = 19
	MjEnumProtoId_PID_PENG_CARD                      MjEnumProtoId = 20
	MjEnumProtoId_PID_PENG_CARD_ACK                  MjEnumProtoId = 21
	MjEnumProtoId_PID_GANG_CARD                      MjEnumProtoId = 22
	MjEnumProtoId_PID_GANG_CARD_ACK                  MjEnumProtoId = 23
	MjEnumProtoId_PID_GUO_CARD                       MjEnumProtoId = 24
	MjEnumProtoId_PID_GUO_CARD_ACK                   MjEnumProtoId = 25
	MjEnumProtoId_PID_HU_CARD                        MjEnumProtoId = 26
	MjEnumProtoId_PID_HU_CARD_ACK                    MjEnumProtoId = 27
	MjEnumProtoId_PID_BROADCAST_BEGIN_DINGQUE        MjEnumProtoId = 28
	MjEnumProtoId_PID_BROADCAST_BEGIN_EXCHANGE       MjEnumProtoId = 29
	MjEnumProtoId_PID_OVERTURN                       MjEnumProtoId = 30
	MjEnumProtoId_PID_CURRENTRESULT                  MjEnumProtoId = 31
	MjEnumProtoId_PID_SENDENDLOTTERY                 MjEnumProtoId = 32
	MjEnumProtoId_PID_DISSOLVE_DESK                  MjEnumProtoId = 33
	MjEnumProtoId_PID_DISSOLVE_DESK_ACK              MjEnumProtoId = 34
	MjEnumProtoId_PID_LEAVE_DESK                     MjEnumProtoId = 35
	MjEnumProtoId_PID_LEAVE_DESK_ACK                 MjEnumProtoId = 36
	MjEnumProtoId_PID_MESSAGE                        MjEnumProtoId = 37
	MjEnumProtoId_PID_SEND_MESSAGE                   MjEnumProtoId = 38
	MjEnumProtoId_PID_GAME_DINGQUEEND                MjEnumProtoId = 39
	MjEnumProtoId_PID_GAME_GAMERECORD                MjEnumProtoId = 40
	MjEnumProtoId_PID_GAME_ACKGAMERECORD             MjEnumProtoId = 41
	MjEnumProtoId_PID_GAME_ACKGAMERECORDEDN          MjEnumProtoId = 42
	MjEnumProtoId_PID_GAME_NOTICE                    MjEnumProtoId = 43
	MjEnumProtoId_PID_GAME_ACKNOTICE                 MjEnumProtoId = 44
	MjEnumProtoId_PID_LOGOUT_REQ                     MjEnumProtoId = 45
	MjEnumProtoId_PID_LOGOUT_ACK                     MjEnumProtoId = 46
	MjEnumProtoId_PID_AWARD_ONLINE_REQ               MjEnumProtoId = 47
	MjEnumProtoId_PID_AWARD_ONLINE_ACK               MjEnumProtoId = 48
	MjEnumProtoId_PID_HALL_TASK_REQ                  MjEnumProtoId = 49
	MjEnumProtoId_PID_HALL_TASK_ACK                  MjEnumProtoId = 50
	MjEnumProtoId_PID_ENTER_AGENTMODE_REQ            MjEnumProtoId = 51
	MjEnumProtoId_PID_ENTER_AGENTMODE_ACK            MjEnumProtoId = 52
	MjEnumProtoId_PID_QUIT_AGENTMODE_REQ             MjEnumProtoId = 53
	MjEnumProtoId_PID_QUIT_AGENTMODE_ACK             MjEnumProtoId = 54
	MjEnumProtoId_PID_REG_REQ                        MjEnumProtoId = 55
	MjEnumProtoId_PID_REG_ACK                        MjEnumProtoId = 56
	MjEnumProtoId_PID_GAMESTATE_REQ                  MjEnumProtoId = 57
	MjEnumProtoId_PID_GAMESTATE_ACK                  MjEnumProtoId = 58
	MjEnumProtoId_PID_FEEDBACK_REQ                   MjEnumProtoId = 59
	MjEnumProtoId_PID_APPLYDISSOLVE_REQ              MjEnumProtoId = 60
	MjEnumProtoId_PID_APPLYDISSOLVE_ACK              MjEnumProtoId = 61
	MjEnumProtoId_PID_APPLYDISSOLVEBACK_REQ          MjEnumProtoId = 62
	MjEnumProtoId_PID_APPLYDISSOLVEBACK_ACK          MjEnumProtoId = 63
	MjEnumProtoId_PID_COMMONBCKICKOUT                MjEnumProtoId = 64
	MjEnumProtoId_PID_ACTCHI_REQ                     MjEnumProtoId = 65
	MjEnumProtoId_PID_ACTCHI_ACK                     MjEnumProtoId = 66
	MjEnumProtoId_PID_CHANGSHA_GANG_CARD_ACK         MjEnumProtoId = 67
	MjEnumProtoId_PID_ACTCHANGSHAQISHOUHU            MjEnumProtoId = 68
	MjEnumProtoId_PID_ACTCHANGSHAQISHOUHU_ACK        MjEnumProtoId = 69
	MjEnumProtoId_PID_GAME_CHANGSHQISHOUHUOVERTURN   MjEnumProtoId = 70
	MjEnumProtoId_PID_GAME_CHANGSHAOVERTURNAFTERGANG MjEnumProtoId = 71
	MjEnumProtoId_PID_GAME_ACKACTHUCHANGSHA          MjEnumProtoId = 72
	MjEnumProtoId_PID_GAME_DEALHAIDICARDS            MjEnumProtoId = 73
	MjEnumProtoId_PID_GAME_REQDEALHAIDICARDS         MjEnumProtoId = 74
	MjEnumProtoId_PID_GAME_ACKDEALHAIDICARDS         MjEnumProtoId = 75
	MjEnumProtoId_PID_COMMONBCUSERBREAK              MjEnumProtoId = 76
	MjEnumProtoId_PID_COMMONREQRECONNECT             MjEnumProtoId = 77
	MjEnumProtoId_PID_GAMEREQBUXIAZI                 MjEnumProtoId = 78
	MjEnumProtoId_PID_GAMEACKBUXIAZI                 MjEnumProtoId = 79
	MjEnumProtoId_PID_OFFLINE_REQ                    MjEnumProtoId = 80
	MjEnumProtoId_PID_RECONNECT_ACK                  MjEnumProtoId = 81
	MjEnumProtoId_PID_BC_BAOTING                     MjEnumProtoId = 82
	MjEnumProtoId_PID_REQ_BAOTING                    MjEnumProtoId = 83
	MjEnumProtoId_PID_ACK_BAOTING                    MjEnumProtoId = 84
	MjEnumProtoId_PID_BC_FENZHANG                    MjEnumProtoId = 85
	// 宜宾麻将
	MjEnumProtoId_PID_BC_PIAO          MjEnumProtoId = 86
	MjEnumProtoId_PID_REQ_PIAO         MjEnumProtoId = 87
	MjEnumProtoId_PID_ACK_PIAO         MjEnumProtoId = 88
	MjEnumProtoId_PID_REQ_FLY          MjEnumProtoId = 89
	MjEnumProtoId_PID_ACK_FLY          MjEnumProtoId = 90
	MjEnumProtoId_PID_REQ_TI           MjEnumProtoId = 91
	MjEnumProtoId_PID_ACK_TI           MjEnumProtoId = 92
	MjEnumProtoId_PID_BC_LEAVE_TIMEOUT MjEnumProtoId = 93
)

var MjEnumProtoId_name = map[int32]string{
	1:  "PID_QUICK_CONN",
	2:  "PID_QUICK_CONN_ACK",
	3:  "PID_GAME_LOGIN",
	4:  "PID_GAME_LOGIN_ACK",
	5:  "PID_CREATEROOM",
	6:  "PID_CREATEROOM_ACK",
	7:  "PID_ENTER_ROOM",
	8:  "PID_ENTER_ROOM_ACK",
	9:  "PID_SEND_GAMEINFO",
	10: "PID_READY",
	11: "PID_READY_ACK",
	12: "PID_EXCHANGECARDS",
	13: "PID_EXCHANGECARDS_ACK",
	14: "PID_DINGQUE",
	15: "PID_OPENING",
	16: "PID_DEAL_CARDS",
	17: "PID_GET_IN_CARD",
	18: "PID_SEND_OUT_CARD",
	19: "PID_SEND_OUT_CARD_ACK",
	20: "PID_PENG_CARD",
	21: "PID_PENG_CARD_ACK",
	22: "PID_GANG_CARD",
	23: "PID_GANG_CARD_ACK",
	24: "PID_GUO_CARD",
	25: "PID_GUO_CARD_ACK",
	26: "PID_HU_CARD",
	27: "PID_HU_CARD_ACK",
	28: "PID_BROADCAST_BEGIN_DINGQUE",
	29: "PID_BROADCAST_BEGIN_EXCHANGE",
	30: "PID_OVERTURN",
	31: "PID_CURRENTRESULT",
	32: "PID_SENDENDLOTTERY",
	33: "PID_DISSOLVE_DESK",
	34: "PID_DISSOLVE_DESK_ACK",
	35: "PID_LEAVE_DESK",
	36: "PID_LEAVE_DESK_ACK",
	37: "PID_MESSAGE",
	38: "PID_SEND_MESSAGE",
	39: "PID_GAME_DINGQUEEND",
	40: "PID_GAME_GAMERECORD",
	41: "PID_GAME_ACKGAMERECORD",
	42: "PID_GAME_ACKGAMERECORDEDN",
	43: "PID_GAME_NOTICE",
	44: "PID_GAME_ACKNOTICE",
	45: "PID_LOGOUT_REQ",
	46: "PID_LOGOUT_ACK",
	47: "PID_AWARD_ONLINE_REQ",
	48: "PID_AWARD_ONLINE_ACK",
	49: "PID_HALL_TASK_REQ",
	50: "PID_HALL_TASK_ACK",
	51: "PID_ENTER_AGENTMODE_REQ",
	52: "PID_ENTER_AGENTMODE_ACK",
	53: "PID_QUIT_AGENTMODE_REQ",
	54: "PID_QUIT_AGENTMODE_ACK",
	55: "PID_REG_REQ",
	56: "PID_REG_ACK",
	57: "PID_GAMESTATE_REQ",
	58: "PID_GAMESTATE_ACK",
	59: "PID_FEEDBACK_REQ",
	60: "PID_APPLYDISSOLVE_REQ",
	61: "PID_APPLYDISSOLVE_ACK",
	62: "PID_APPLYDISSOLVEBACK_REQ",
	63: "PID_APPLYDISSOLVEBACK_ACK",
	64: "PID_COMMONBCKICKOUT",
	65: "PID_ACTCHI_REQ",
	66: "PID_ACTCHI_ACK",
	67: "PID_CHANGSHA_GANG_CARD_ACK",
	68: "PID_ACTCHANGSHAQISHOUHU",
	69: "PID_ACTCHANGSHAQISHOUHU_ACK",
	70: "PID_GAME_CHANGSHQISHOUHUOVERTURN",
	71: "PID_GAME_CHANGSHAOVERTURNAFTERGANG",
	72: "PID_GAME_ACKACTHUCHANGSHA",
	73: "PID_GAME_DEALHAIDICARDS",
	74: "PID_GAME_REQDEALHAIDICARDS",
	75: "PID_GAME_ACKDEALHAIDICARDS",
	76: "PID_COMMONBCUSERBREAK",
	77: "PID_COMMONREQRECONNECT",
	78: "PID_GAMEREQBUXIAZI",
	79: "PID_GAMEACKBUXIAZI",
	80: "PID_OFFLINE_REQ",
	81: "PID_RECONNECT_ACK",
	82: "PID_BC_BAOTING",
	83: "PID_REQ_BAOTING",
	84: "PID_ACK_BAOTING",
	85: "PID_BC_FENZHANG",
	86: "PID_BC_PIAO",
	87: "PID_REQ_PIAO",
	88: "PID_ACK_PIAO",
	89: "PID_REQ_FLY",
	90: "PID_ACK_FLY",
	91: "PID_REQ_TI",
	92: "PID_ACK_TI",
	93: "PID_BC_LEAVE_TIMEOUT",
}
var MjEnumProtoId_value = map[string]int32{
	"PID_QUICK_CONN":                     1,
	"PID_QUICK_CONN_ACK":                 2,
	"PID_GAME_LOGIN":                     3,
	"PID_GAME_LOGIN_ACK":                 4,
	"PID_CREATEROOM":                     5,
	"PID_CREATEROOM_ACK":                 6,
	"PID_ENTER_ROOM":                     7,
	"PID_ENTER_ROOM_ACK":                 8,
	"PID_SEND_GAMEINFO":                  9,
	"PID_READY":                          10,
	"PID_READY_ACK":                      11,
	"PID_EXCHANGECARDS":                  12,
	"PID_EXCHANGECARDS_ACK":              13,
	"PID_DINGQUE":                        14,
	"PID_OPENING":                        15,
	"PID_DEAL_CARDS":                     16,
	"PID_GET_IN_CARD":                    17,
	"PID_SEND_OUT_CARD":                  18,
	"PID_SEND_OUT_CARD_ACK":              19,
	"PID_PENG_CARD":                      20,
	"PID_PENG_CARD_ACK":                  21,
	"PID_GANG_CARD":                      22,
	"PID_GANG_CARD_ACK":                  23,
	"PID_GUO_CARD":                       24,
	"PID_GUO_CARD_ACK":                   25,
	"PID_HU_CARD":                        26,
	"PID_HU_CARD_ACK":                    27,
	"PID_BROADCAST_BEGIN_DINGQUE":        28,
	"PID_BROADCAST_BEGIN_EXCHANGE":       29,
	"PID_OVERTURN":                       30,
	"PID_CURRENTRESULT":                  31,
	"PID_SENDENDLOTTERY":                 32,
	"PID_DISSOLVE_DESK":                  33,
	"PID_DISSOLVE_DESK_ACK":              34,
	"PID_LEAVE_DESK":                     35,
	"PID_LEAVE_DESK_ACK":                 36,
	"PID_MESSAGE":                        37,
	"PID_SEND_MESSAGE":                   38,
	"PID_GAME_DINGQUEEND":                39,
	"PID_GAME_GAMERECORD":                40,
	"PID_GAME_ACKGAMERECORD":             41,
	"PID_GAME_ACKGAMERECORDEDN":          42,
	"PID_GAME_NOTICE":                    43,
	"PID_GAME_ACKNOTICE":                 44,
	"PID_LOGOUT_REQ":                     45,
	"PID_LOGOUT_ACK":                     46,
	"PID_AWARD_ONLINE_REQ":               47,
	"PID_AWARD_ONLINE_ACK":               48,
	"PID_HALL_TASK_REQ":                  49,
	"PID_HALL_TASK_ACK":                  50,
	"PID_ENTER_AGENTMODE_REQ":            51,
	"PID_ENTER_AGENTMODE_ACK":            52,
	"PID_QUIT_AGENTMODE_REQ":             53,
	"PID_QUIT_AGENTMODE_ACK":             54,
	"PID_REG_REQ":                        55,
	"PID_REG_ACK":                        56,
	"PID_GAMESTATE_REQ":                  57,
	"PID_GAMESTATE_ACK":                  58,
	"PID_FEEDBACK_REQ":                   59,
	"PID_APPLYDISSOLVE_REQ":              60,
	"PID_APPLYDISSOLVE_ACK":              61,
	"PID_APPLYDISSOLVEBACK_REQ":          62,
	"PID_APPLYDISSOLVEBACK_ACK":          63,
	"PID_COMMONBCKICKOUT":                64,
	"PID_ACTCHI_REQ":                     65,
	"PID_ACTCHI_ACK":                     66,
	"PID_CHANGSHA_GANG_CARD_ACK":         67,
	"PID_ACTCHANGSHAQISHOUHU":            68,
	"PID_ACTCHANGSHAQISHOUHU_ACK":        69,
	"PID_GAME_CHANGSHQISHOUHUOVERTURN":   70,
	"PID_GAME_CHANGSHAOVERTURNAFTERGANG": 71,
	"PID_GAME_ACKACTHUCHANGSHA":          72,
	"PID_GAME_DEALHAIDICARDS":            73,
	"PID_GAME_REQDEALHAIDICARDS":         74,
	"PID_GAME_ACKDEALHAIDICARDS":         75,
	"PID_COMMONBCUSERBREAK":              76,
	"PID_COMMONREQRECONNECT":             77,
	"PID_GAMEREQBUXIAZI":                 78,
	"PID_GAMEACKBUXIAZI":                 79,
	"PID_OFFLINE_REQ":                    80,
	"PID_RECONNECT_ACK":                  81,
	"PID_BC_BAOTING":                     82,
	"PID_REQ_BAOTING":                    83,
	"PID_ACK_BAOTING":                    84,
	"PID_BC_FENZHANG":                    85,
	"PID_BC_PIAO":                        86,
	"PID_REQ_PIAO":                       87,
	"PID_ACK_PIAO":                       88,
	"PID_REQ_FLY":                        89,
	"PID_ACK_FLY":                        90,
	"PID_REQ_TI":                         91,
	"PID_ACK_TI":                         92,
	"PID_BC_LEAVE_TIMEOUT":               93,
}

func (x MjEnumProtoId) Enum() *MjEnumProtoId {
	p := new(MjEnumProtoId)
	*p = x
	return p
}
func (x MjEnumProtoId) String() string {
	return proto.EnumName(MjEnumProtoId_name, int32(x))
}
func (x *MjEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjEnumProtoId_value, data, "MjEnumProtoId")
	if err != nil {
		return err
	}
	*x = MjEnumProtoId(value)
	return nil
}
func (MjEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

type MJOption int32

const (
	MJOption_EXCHANGE_CARDS       MJOption = 1
	MJOption_ZIMO_JIA_DI          MJOption = 2
	MJOption_ZIMO_JIA_FAN         MJOption = 3
	MJOption_DIANGANG_HUA_DIANPAO MJOption = 4
	MJOption_DIANGANG_HUA_ZIMO    MJOption = 5
	MJOption_YAOJIU_JIANGDUI      MJOption = 6
	MJOption_MENQING_MID_CARD     MJOption = 7
	MJOption_TIAN_DI_HU           MJOption = 8
	MJOption_KA_ER_TIAO           MJOption = 9
	// 三人两房牌
	MJOption_DIANPAO_CAN_HU MJOption = 10
	MJOption_DUIDUIHU_2_FAN MJOption = 11
	MJOption_JIA_XIN_WU     MJOption = 12
	// 倒倒胡
	MJOption_DDH_QIDUI_HU   MJOption = 13
	MJOption_DDH_ZIMO_HU    MJOption = 14
	MJOption_DDH_DIANPAO_HU MJOption = 15
)

var MJOption_name = map[int32]string{
	1:  "EXCHANGE_CARDS",
	2:  "ZIMO_JIA_DI",
	3:  "ZIMO_JIA_FAN",
	4:  "DIANGANG_HUA_DIANPAO",
	5:  "DIANGANG_HUA_ZIMO",
	6:  "YAOJIU_JIANGDUI",
	7:  "MENQING_MID_CARD",
	8:  "TIAN_DI_HU",
	9:  "KA_ER_TIAO",
	10: "DIANPAO_CAN_HU",
	11: "DUIDUIHU_2_FAN",
	12: "JIA_XIN_WU",
	13: "DDH_QIDUI_HU",
	14: "DDH_ZIMO_HU",
	15: "DDH_DIANPAO_HU",
}
var MJOption_value = map[string]int32{
	"EXCHANGE_CARDS":       1,
	"ZIMO_JIA_DI":          2,
	"ZIMO_JIA_FAN":         3,
	"DIANGANG_HUA_DIANPAO": 4,
	"DIANGANG_HUA_ZIMO":    5,
	"YAOJIU_JIANGDUI":      6,
	"MENQING_MID_CARD":     7,
	"TIAN_DI_HU":           8,
	"KA_ER_TIAO":           9,
	"DIANPAO_CAN_HU":       10,
	"DUIDUIHU_2_FAN":       11,
	"JIA_XIN_WU":           12,
	"DDH_QIDUI_HU":         13,
	"DDH_ZIMO_HU":          14,
	"DDH_DIANPAO_HU":       15,
}

func (x MJOption) Enum() *MJOption {
	p := new(MJOption)
	*p = x
	return p
}
func (x MJOption) String() string {
	return proto.EnumName(MJOption_name, int32(x))
}
func (x *MJOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MJOption_value, data, "MJOption")
	if err != nil {
		return err
	}
	*x = MJOption(value)
	return nil
}
func (MJOption) EnumDescriptor() ([]byte, []int) { return fileDescriptor30, []int{1} }

type MJPlayerInfo struct {
	IsBanker         *bool       `protobuf:"varint,1,opt,name=isBanker" json:"isBanker,omitempty"`
	PlayerCard       *PlayerCard `protobuf:"bytes,2,opt,name=playerCard" json:"playerCard,omitempty"`
	Coin             *int64      `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	NickName         *string     `protobuf:"bytes,4,opt,name=nickName" json:"nickName,omitempty"`
	Sex              *int32      `protobuf:"varint,5,opt,name=sex" json:"sex,omitempty"`
	UserId           *uint32     `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool       `protobuf:"varint,7,opt,name=isOwner" json:"isOwner,omitempty"`
	BReady           *int32      `protobuf:"varint,8,opt,name=bReady" json:"bReady,omitempty"`
	BDingQue         *int32      `protobuf:"varint,9,opt,name=bDingQue" json:"bDingQue,omitempty"`
	BExchanged       *int32      `protobuf:"varint,10,opt,name=bExchanged" json:"bExchanged,omitempty"`
	NHuPai           *int32      `protobuf:"varint,11,opt,name=nHuPai" json:"nHuPai,omitempty"`
	QuePai           *int32      `protobuf:"varint,12,opt,name=quePai" json:"quePai,omitempty"`
	WxInfo           *WeixinInfo `protobuf:"bytes,13,opt,name=wxInfo" json:"wxInfo,omitempty"`
	GameStatus       *int32      `protobuf:"varint,14,opt,name=GameStatus" json:"GameStatus,omitempty"`
	AgentMode        *bool       `protobuf:"varint,15,opt,name=agentMode" json:"agentMode,omitempty"`
	Ip               *string     `protobuf:"bytes,16,opt,name=ip" json:"ip,omitempty"`
	XiaCount         *int32      `protobuf:"varint,17,opt,name=xiaCount" json:"xiaCount,omitempty"`
	IsBaoTing        *bool       `protobuf:"varint,18,opt,name=isBaoTing" json:"isBaoTing,omitempty"`
	IsQianniu        *bool       `protobuf:"varint,19,opt,name=isQianniu" json:"isQianniu,omitempty"`
	Address          *string     `protobuf:"bytes,20,opt,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MJPlayerInfo) Reset()                    { *m = MJPlayerInfo{} }
func (m *MJPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*MJPlayerInfo) ProtoMessage()               {}
func (*MJPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

func (m *MJPlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *MJPlayerInfo) GetPlayerCard() *PlayerCard {
	if m != nil {
		return m.PlayerCard
	}
	return nil
}

func (m *MJPlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *MJPlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *MJPlayerInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *MJPlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *MJPlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *MJPlayerInfo) GetBReady() int32 {
	if m != nil && m.BReady != nil {
		return *m.BReady
	}
	return 0
}

func (m *MJPlayerInfo) GetBDingQue() int32 {
	if m != nil && m.BDingQue != nil {
		return *m.BDingQue
	}
	return 0
}

func (m *MJPlayerInfo) GetBExchanged() int32 {
	if m != nil && m.BExchanged != nil {
		return *m.BExchanged
	}
	return 0
}

func (m *MJPlayerInfo) GetNHuPai() int32 {
	if m != nil && m.NHuPai != nil {
		return *m.NHuPai
	}
	return 0
}

func (m *MJPlayerInfo) GetQuePai() int32 {
	if m != nil && m.QuePai != nil {
		return *m.QuePai
	}
	return 0
}

func (m *MJPlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *MJPlayerInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *MJPlayerInfo) GetAgentMode() bool {
	if m != nil && m.AgentMode != nil {
		return *m.AgentMode
	}
	return false
}

func (m *MJPlayerInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *MJPlayerInfo) GetXiaCount() int32 {
	if m != nil && m.XiaCount != nil {
		return *m.XiaCount
	}
	return 0
}

func (m *MJPlayerInfo) GetIsBaoTing() bool {
	if m != nil && m.IsBaoTing != nil {
		return *m.IsBaoTing
	}
	return false
}

func (m *MJPlayerInfo) GetIsQianniu() bool {
	if m != nil && m.IsQianniu != nil {
		return *m.IsQianniu
	}
	return false
}

func (m *MJPlayerInfo) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type MJDeskGameInfo struct {
	GameStatus       *int32        `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	RoomTypeInfo     *RoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	PlayerNum        *int32        `protobuf:"varint,3,opt,name=playerNum" json:"playerNum,omitempty"`
	ActiveUserId     *uint32       `protobuf:"varint,4,opt,name=activeUserId" json:"activeUserId,omitempty"`
	ActionTime       *int32        `protobuf:"varint,5,opt,name=actionTime" json:"actionTime,omitempty"`
	DelayTime        *int32        `protobuf:"varint,6,opt,name=delayTime" json:"delayTime,omitempty"`
	NInitActionTime  *int32        `protobuf:"varint,7,opt,name=nInitActionTime" json:"nInitActionTime,omitempty"`
	NInitDelayTime   *int32        `protobuf:"varint,8,opt,name=nInitDelayTime" json:"nInitDelayTime,omitempty"`
	InitRoomCoin     *int64        `protobuf:"varint,9,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32        `protobuf:"varint,10,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32        `protobuf:"varint,11,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	RoomNumber       *string       `protobuf:"bytes,12,opt,name=roomNumber" json:"roomNumber,omitempty"`
	RemainCards      *int32        `protobuf:"varint,13,opt,name=remainCards" json:"remainCards,omitempty"`
	Banker           *uint32       `protobuf:"varint,14,opt,name=Banker" json:"Banker,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *MJDeskGameInfo) Reset()                    { *m = MJDeskGameInfo{} }
func (m *MJDeskGameInfo) String() string            { return proto.CompactTextString(m) }
func (*MJDeskGameInfo) ProtoMessage()               {}
func (*MJDeskGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{1} }

func (m *MJDeskGameInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *MJDeskGameInfo) GetRoomTypeInfo() *RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *MJDeskGameInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *MJDeskGameInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *MJDeskGameInfo) GetActionTime() int32 {
	if m != nil && m.ActionTime != nil {
		return *m.ActionTime
	}
	return 0
}

func (m *MJDeskGameInfo) GetDelayTime() int32 {
	if m != nil && m.DelayTime != nil {
		return *m.DelayTime
	}
	return 0
}

func (m *MJDeskGameInfo) GetNInitActionTime() int32 {
	if m != nil && m.NInitActionTime != nil {
		return *m.NInitActionTime
	}
	return 0
}

func (m *MJDeskGameInfo) GetNInitDelayTime() int32 {
	if m != nil && m.NInitDelayTime != nil {
		return *m.NInitDelayTime
	}
	return 0
}

func (m *MJDeskGameInfo) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *MJDeskGameInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *MJDeskGameInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *MJDeskGameInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *MJDeskGameInfo) GetRemainCards() int32 {
	if m != nil && m.RemainCards != nil {
		return *m.RemainCards
	}
	return 0
}

func (m *MJDeskGameInfo) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

func init() {
	proto.RegisterType((*MJPlayerInfo)(nil), "ddproto.MJPlayerInfo")
	proto.RegisterType((*MJDeskGameInfo)(nil), "ddproto.MJDeskGameInfo")
	proto.RegisterEnum("ddproto.MjEnumProtoId", MjEnumProtoId_name, MjEnumProtoId_value)
	proto.RegisterEnum("ddproto.MJOption", MJOption_name, MJOption_value)
}

var fileDescriptor30 = []byte{
	// 1625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x96, 0xeb, 0x76, 0xdb, 0xb8,
	0x11, 0xc7, 0x2b, 0xc7, 0x57, 0xd8, 0x96, 0x61, 0xd8, 0x49, 0x18, 0x27, 0x9b, 0x55, 0xdd, 0x34,
	0x55, 0xb3, 0x6d, 0xda, 0x66, 0x7b, 0xdb, 0xde, 0x21, 0x12, 0x12, 0x61, 0x89, 0xa0, 0x04, 0x92,
	0x49, 0x9c, 0xb6, 0x87, 0x87, 0xb6, 0x58, 0x97, 0x49, 0x44, 0xb9, 0xba, 0x74, 0x9d, 0x47, 0xe8,
	0x7b, 0xf4, 0x63, 0x1f, 0xa3, 0x0f, 0xb6, 0x67, 0x40, 0x0d, 0x75, 0xb1, 0xf3, 0xc5, 0xc7, 0xf8,
	0xcd, 0x1f, 0x83, 0x99, 0xc1, 0x0c, 0x28, 0xc2, 0x06, 0xc9, 0xfb, 0x2c, 0xc9, 0xaf, 0xe2, 0x8b,
	0x64, 0x9c, 0xbe, 0xbc, 0x1e, 0x0d, 0x27, 0x43, 0xb6, 0xd5, 0xef, 0x9b, 0x7f, 0x4e, 0x0e, 0x2e,
	0x87, 0x83, 0xc1, 0x30, 0x8f, 0x07, 0xef, 0x0b, 0xcb, 0xe9, 0xff, 0xd7, 0xc9, 0x9e, 0x77, 0xd6,
	0xfd, 0x98, 0x7c, 0x4a, 0x47, 0x32, 0xff, 0xc7, 0x90, 0x9d, 0x90, 0xed, 0x6c, 0xdc, 0x48, 0xf2,
	0x0f, 0xe9, 0xc8, 0xaa, 0xd4, 0x2a, 0xf5, 0x6d, 0x5d, 0xae, 0xd9, 0xd7, 0x84, 0x5c, 0x1b, 0xa5,
	0x9d, 0x8c, 0xfa, 0xd6, 0x5a, 0xad, 0x52, 0xdf, 0x7d, 0x75, 0xf4, 0x72, 0xe6, 0xfb, 0x65, 0xb7,
	0x34, 0xe9, 0x05, 0x19, 0x63, 0x64, 0xfd, 0x72, 0x98, 0xe5, 0xd6, 0xbd, 0x5a, 0xa5, 0x7e, 0x4f,
	0x9b, 0xff, 0xe1, 0x90, 0x3c, 0xbb, 0xfc, 0xa0, 0x92, 0x41, 0x6a, 0xad, 0xd7, 0x2a, 0xf5, 0x1d,
	0x5d, 0xae, 0x19, 0x25, 0xf7, 0xc6, 0xe9, 0x8d, 0xb5, 0x51, 0xab, 0xd4, 0x37, 0x34, 0xfc, 0xcb,
	0x1e, 0x90, 0xcd, 0xe9, 0x38, 0x1d, 0xc9, 0xbe, 0xb5, 0x59, 0xab, 0xd4, 0xf7, 0xf5, 0x6c, 0xc5,
	0x2c, 0xb2, 0x95, 0x8d, 0xfd, 0x6f, 0xf3, 0x74, 0x64, 0x6d, 0x99, 0x48, 0x71, 0x09, 0x3b, 0x2e,
	0x74, 0x9a, 0xf4, 0x3f, 0x59, 0xdb, 0xc6, 0xcd, 0x6c, 0x05, 0xe7, 0x5e, 0x38, 0x59, 0x7e, 0xd5,
	0x9b, 0xa6, 0xd6, 0x8e, 0xb1, 0x94, 0x6b, 0xf6, 0x94, 0x90, 0x0b, 0x71, 0x73, 0xf9, 0xcf, 0x24,
	0xbf, 0x4a, 0xfb, 0x16, 0x31, 0xd6, 0x05, 0x02, 0x3e, 0x73, 0x77, 0xda, 0x4d, 0x32, 0x6b, 0xb7,
	0xf0, 0x59, 0xac, 0x80, 0xff, 0x6b, 0x9a, 0x02, 0xdf, 0x2b, 0x78, 0xb1, 0x62, 0x5f, 0x91, 0xcd,
	0x6f, 0x6f, 0xa0, 0xa4, 0xd6, 0xfe, 0x4a, 0xa1, 0xde, 0xa4, 0xd9, 0x4d, 0x96, 0x83, 0x49, 0xcf,
	0x24, 0x70, 0x78, 0x2b, 0x19, 0xa4, 0xc1, 0x24, 0x99, 0x4c, 0xc7, 0x56, 0xb5, 0x38, 0x7c, 0x4e,
	0xd8, 0x13, 0xb2, 0x93, 0x5c, 0xa5, 0xf9, 0xc4, 0x1b, 0xf6, 0x53, 0xeb, 0xc0, 0x24, 0x3b, 0x07,
	0xac, 0x4a, 0xd6, 0xb2, 0x6b, 0x8b, 0x9a, 0x42, 0xae, 0x65, 0xd7, 0x90, 0xe6, 0x4d, 0x96, 0xd8,
	0xc3, 0x69, 0x3e, 0xb1, 0x0e, 0x8b, 0x34, 0x71, 0x0d, 0x9e, 0xe0, 0x3e, 0x87, 0x61, 0x96, 0x5f,
	0x59, 0xac, 0xf0, 0x54, 0x82, 0xc2, 0xda, 0xcb, 0x92, 0x3c, 0xcf, 0xa6, 0xd6, 0x11, 0x5a, 0x67,
	0x00, 0x0a, 0x9e, 0xf4, 0xfb, 0xa3, 0x74, 0x3c, 0xb6, 0x8e, 0xcd, 0x61, 0xb8, 0x3c, 0xfd, 0xcf,
	0x3a, 0xa9, 0x7a, 0x67, 0x4e, 0x3a, 0xfe, 0x00, 0x41, 0xdf, 0x91, 0x52, 0xe5, 0x56, 0x4a, 0xdf,
	0x90, 0xbd, 0xd1, 0x70, 0x38, 0x08, 0x3f, 0x5d, 0x1b, 0xfd, 0xac, 0x9d, 0xee, 0x97, 0x55, 0xd2,
	0x0b, 0x46, 0xbd, 0x24, 0x85, 0x28, 0x8b, 0x06, 0x53, 0xd3, 0x81, 0xe9, 0xab, 0x0d, 0x3d, 0x07,
	0xec, 0x94, 0xec, 0x25, 0x97, 0x93, 0xec, 0xdf, 0x69, 0x54, 0x34, 0xcd, 0xba, 0x69, 0x9a, 0x25,
	0x06, 0xc1, 0xc1, 0x7a, 0x98, 0x87, 0xd9, 0x20, 0x9d, 0xf5, 0xda, 0x02, 0x81, 0x13, 0xfa, 0xe9,
	0xc7, 0xe4, 0x93, 0x31, 0x6f, 0x16, 0x27, 0x94, 0x80, 0xd5, 0xc9, 0x41, 0x2e, 0xf3, 0x6c, 0xc2,
	0xe7, 0x2e, 0xb6, 0x8c, 0x66, 0x15, 0xb3, 0xe7, 0xa4, 0x6a, 0x90, 0x53, 0x3a, 0x2b, 0x1a, 0x72,
	0x85, 0x42, 0xcc, 0x59, 0x9e, 0x4d, 0x20, 0x67, 0x1b, 0x86, 0x65, 0xc7, 0x0c, 0xcb, 0x12, 0x63,
	0xcf, 0xc8, 0xfe, 0xe5, 0x74, 0x34, 0x82, 0x31, 0x2b, 0xae, 0xb6, 0xe8, 0xd1, 0x65, 0x08, 0x27,
	0x4e, 0x86, 0x93, 0xe4, 0xe3, 0x5c, 0x56, 0xb4, 0xeb, 0x0a, 0x85, 0x0a, 0x40, 0x4d, 0xd5, 0x74,
	0x70, 0x91, 0x8e, 0x4c, 0xeb, 0xee, 0xe8, 0x05, 0xc2, 0x6a, 0x64, 0x77, 0x94, 0x0e, 0x92, 0x2c,
	0x87, 0x21, 0x1e, 0x9b, 0x1e, 0xde, 0xd0, 0x8b, 0x08, 0x1a, 0x7f, 0xf6, 0x4e, 0x54, 0x8b, 0xb1,
	0x2c, 0x56, 0x2f, 0xfe, 0x7b, 0x48, 0x0e, 0x06, 0xef, 0xe3, 0x34, 0x9f, 0x0e, 0x62, 0x73, 0x95,
	0x12, 0x1e, 0x81, 0x6a, 0x57, 0x3a, 0x71, 0x2f, 0x92, 0x76, 0x3b, 0xb6, 0x7d, 0xa5, 0x68, 0x85,
	0x3d, 0x20, 0x6c, 0x99, 0xc5, 0xdc, 0x6e, 0xd3, 0x35, 0xd4, 0xb6, 0xb8, 0x27, 0xe2, 0x8e, 0xdf,
	0x92, 0x8a, 0xde, 0x43, 0xed, 0x9c, 0x19, 0xed, 0x3a, 0x6a, 0x6d, 0x2d, 0x78, 0x28, 0xb4, 0xef,
	0x7b, 0x74, 0x03, 0xb5, 0x73, 0x66, 0xb4, 0x9b, 0xa8, 0x15, 0x2a, 0x14, 0x3a, 0x36, 0xda, 0x2d,
	0xd4, 0xce, 0x99, 0xd1, 0x6e, 0xb3, 0xfb, 0xe4, 0x10, 0x78, 0x20, 0x54, 0x71, 0xa8, 0x54, 0x4d,
	0x9f, 0xee, 0xb0, 0x7d, 0xb2, 0x03, 0x58, 0x0b, 0xee, 0x9c, 0x53, 0xc2, 0x0e, 0xc9, 0x7e, 0xb9,
	0x34, 0x1b, 0x77, 0x71, 0xa3, 0x78, 0x6b, 0xbb, 0x5c, 0xb5, 0x84, 0xcd, 0xb5, 0x13, 0xd0, 0x3d,
	0xf6, 0x88, 0xdc, 0xbf, 0x85, 0xcd, 0x8e, 0x7d, 0x76, 0x40, 0x76, 0xc1, 0xe4, 0x48, 0xd5, 0xea,
	0x45, 0x82, 0x56, 0x11, 0xf8, 0x5d, 0xa1, 0xa4, 0x6a, 0xd1, 0x03, 0x0c, 0xdc, 0x11, 0xbc, 0x13,
	0x17, 0x0e, 0x29, 0x3b, 0x22, 0x07, 0xa6, 0x20, 0x22, 0x8c, 0xa5, 0x32, 0x94, 0x1e, 0x2e, 0x45,
	0xed, 0x47, 0x61, 0x81, 0x19, 0x1e, 0xbe, 0x84, 0xcd, 0xe1, 0x47, 0x98, 0x41, 0x57, 0xa8, 0x56,
	0xa1, 0x3e, 0x46, 0x27, 0x25, 0x32, 0xca, 0xfb, 0xa8, 0x6c, 0x71, 0x54, 0x3e, 0x40, 0x65, 0x89,
	0x8c, 0xf2, 0x21, 0xa3, 0x64, 0xcf, 0xe0, 0xc8, 0x2f, 0x84, 0x16, 0x3b, 0x26, 0x74, 0x91, 0x18,
	0xdd, 0x23, 0xcc, 0xd3, 0x8d, 0x0a, 0xd9, 0x09, 0xe6, 0x34, 0x03, 0x46, 0xf5, 0x98, 0x7d, 0x49,
	0x1e, 0x03, 0x6c, 0x68, 0x9f, 0x3b, 0x36, 0x0f, 0xc2, 0xb8, 0x21, 0xe0, 0xfa, 0xb1, 0x5c, 0x4f,
	0x58, 0x8d, 0x3c, 0xb9, 0x4b, 0x80, 0xa5, 0xa6, 0x5f, 0x60, 0x40, 0xfe, 0x6b, 0xa1, 0xc3, 0x48,
	0x2b, 0xfa, 0x14, 0x23, 0xb7, 0x23, 0xad, 0x85, 0x0a, 0xb5, 0x08, 0xa2, 0x4e, 0x48, 0xbf, 0xc4,
	0x6e, 0x80, 0x42, 0x09, 0xe5, 0x74, 0xfc, 0x30, 0x14, 0xfa, 0x9c, 0xd6, 0x50, 0xee, 0xc8, 0x20,
	0xf0, 0x3b, 0xaf, 0x45, 0xec, 0x88, 0xa0, 0x4d, 0xbf, 0x8f, 0x75, 0x5d, 0xc2, 0x26, 0xea, 0x53,
	0xbc, 0xb2, 0x8e, 0xe0, 0x28, 0xff, 0x01, 0x7a, 0x9f, 0x33, 0xa3, 0x7d, 0x86, 0x75, 0xf0, 0x44,
	0x10, 0xf0, 0x96, 0xa0, 0x3f, 0xc4, 0x72, 0x99, 0xfb, 0x42, 0xfa, 0x9c, 0x3d, 0x24, 0x47, 0xe5,
	0x08, 0xcc, 0xb2, 0x17, 0xca, 0xa1, 0x3f, 0x5a, 0x32, 0xc0, 0x1f, 0x2d, 0x6c, 0x5f, 0x3b, 0xb4,
	0xce, 0x4e, 0xc8, 0x83, 0xd2, 0xc0, 0xed, 0xf6, 0x82, 0xed, 0xc7, 0xec, 0x0b, 0xf2, 0xe8, 0x6e,
	0x9b, 0x70, 0x14, 0x7d, 0x51, 0xb6, 0x17, 0x98, 0x95, 0x1f, 0x4a, 0x5b, 0xd0, 0xaf, 0x96, 0x86,
	0x90, 0xdb, 0xed, 0x19, 0xff, 0x49, 0x99, 0xac, 0xdf, 0x82, 0xee, 0xd2, 0xa2, 0x47, 0x7f, 0xba,
	0xc2, 0x20, 0xd1, 0x97, 0xcc, 0x22, 0xc7, 0xc0, 0xf8, 0x1b, 0xb8, 0x5d, 0x5f, 0x75, 0xa4, 0x12,
	0x46, 0xfd, 0xb3, 0x3b, 0x2d, 0xb0, 0xe7, 0xe7, 0x58, 0x7a, 0x97, 0x77, 0x3a, 0x71, 0xc8, 0x83,
	0xb6, 0xd9, 0xf0, 0x8b, 0xdb, 0x18, 0xd4, 0xaf, 0xd8, 0x63, 0xf2, 0x70, 0x3e, 0xce, 0xbc, 0x25,
	0x54, 0xe8, 0xf9, 0x4e, 0x71, 0xc8, 0xd7, 0x9f, 0x33, 0xc2, 0xce, 0x5f, 0x62, 0xad, 0x7a, 0x91,
	0x0c, 0x57, 0x36, 0xfe, 0xea, 0x33, 0x36, 0xd8, 0xf7, 0x6b, 0xbc, 0x3c, 0x2d, 0x5a, 0x46, 0xfc,
	0x9b, 0x45, 0x00, 0x8a, 0xdf, 0xce, 0xa7, 0xc4, 0x13, 0x41, 0xc8, 0xc3, 0xc2, 0xe9, 0x37, 0xb7,
	0x31, 0xa8, 0x7f, 0x87, 0x77, 0xdf, 0x14, 0xc2, 0x69, 0x70, 0xbb, 0x48, 0xf7, 0xf7, 0xd8, 0x69,
	0xbc, 0xdb, 0xed, 0x9c, 0x97, 0xed, 0x06, 0xa6, 0x3f, 0xdc, 0x6d, 0x02, 0x5f, 0x7f, 0xc4, 0x3b,
	0x5e, 0x32, 0x95, 0x4e, 0xff, 0xf4, 0x79, 0x33, 0xec, 0xfe, 0x33, 0xb6, 0x95, 0xed, 0x7b, 0x9e,
	0xaf, 0x1a, 0x76, 0x5b, 0xda, 0x6d, 0x3f, 0x0a, 0xe9, 0x5f, 0xf0, 0x6a, 0xb9, 0x1d, 0xda, 0xae,
	0x34, 0xbe, 0xf8, 0x0a, 0x03, 0x07, 0x0d, 0xf6, 0x94, 0x9c, 0x18, 0x07, 0x30, 0x86, 0x81, 0xcb,
	0x57, 0xde, 0x09, 0x1b, 0xef, 0xc3, 0xec, 0x29, 0x24, 0x3d, 0x19, 0xb8, 0x7e, 0xe4, 0x46, 0xd4,
	0xc1, 0xb1, 0xbf, 0xc3, 0x68, 0x76, 0x0b, 0xf6, 0x8c, 0xd4, 0xca, 0x66, 0x9c, 0x49, 0x50, 0x51,
	0x0e, 0x7a, 0x93, 0x3d, 0x27, 0xa7, 0xab, 0x2a, 0x8e, 0x66, 0xde, 0x0c, 0x85, 0x86, 0xa0, 0x68,
	0x6b, 0x75, 0x1c, 0xb8, 0x1d, 0xba, 0x11, 0xea, 0xa9, 0x8b, 0xa1, 0x16, 0xb3, 0x27, 0x78, 0xc7,
	0xe5, 0xd2, 0x91, 0xc5, 0x53, 0x2c, 0x31, 0x4f, 0x63, 0xd4, 0xa2, 0xb7, 0x62, 0x3f, 0x5b, 0xb2,
	0x73, 0xbb, 0xbd, 0x62, 0x6f, 0xe3, 0x0d, 0x62, 0xa1, 0xa3, 0x40, 0xe8, 0x86, 0x16, 0xbc, 0x4d,
	0x3b, 0xd8, 0x79, 0x85, 0x49, 0x8b, 0x1e, 0x8c, 0xa8, 0x52, 0xc2, 0x0e, 0xa9, 0xb7, 0x38, 0x8d,
	0x5a, 0xf4, 0x1a, 0xd1, 0x5b, 0xc9, 0xdf, 0x49, 0xaa, 0x16, 0x39, 0xb7, 0xdb, 0xc8, 0x7d, 0x1c,
	0x69, 0xbf, 0xd9, 0x2c, 0x07, 0xaf, 0x8b, 0x5d, 0x58, 0xfa, 0x35, 0xc5, 0xed, 0xe1, 0x75, 0x36,
	0xec, 0xb8, 0xc1, 0xfd, 0x10, 0xbe, 0x42, 0x1a, 0xf7, 0x6b, 0xd1, 0x2b, 0x61, 0x80, 0x10, 0xba,
	0x06, 0x61, 0x88, 0xb0, 0x61, 0xc7, 0x4d, 0xa1, 0xde, 0x41, 0x15, 0x69, 0x84, 0x73, 0xd1, 0xb0,
	0xe3, 0xae, 0xe4, 0x3e, 0x7d, 0x8d, 0xaf, 0x32, 0xf8, 0x33, 0xe4, 0x0d, 0x12, 0x70, 0x66, 0xc8,
	0xdb, 0xf9, 0x30, 0xf5, 0xe2, 0x66, 0xe7, 0x9c, 0x9e, 0x23, 0x00, 0x09, 0x80, 0x77, 0xac, 0x4a,
	0x08, 0x2a, 0x42, 0x49, 0xff, 0x8a, 0x6b, 0x10, 0x84, 0x92, 0xfe, 0x0d, 0x5f, 0x96, 0x86, 0x3d,
	0x7b, 0x77, 0x43, 0xe9, 0x09, 0x68, 0xe3, 0xbf, 0xbf, 0xf8, 0xdf, 0x1a, 0xd9, 0xf6, 0xce, 0xfc,
	0x6b, 0xf8, 0xad, 0x06, 0x09, 0xe3, 0x07, 0x63, 0xf6, 0x89, 0xad, 0xc0, 0x59, 0xef, 0xa4, 0xe7,
	0xc7, 0x67, 0x92, 0xc7, 0x8e, 0xa4, 0x6b, 0x10, 0x5f, 0x09, 0x9a, 0x1c, 0x7e, 0x96, 0x58, 0xe4,
	0xd8, 0x91, 0x5c, 0x99, 0xce, 0x76, 0x23, 0x90, 0x71, 0xd5, 0xe5, 0x3e, 0x5d, 0x87, 0xc2, 0x2e,
	0x59, 0x60, 0x23, 0xdd, 0x80, 0xd2, 0x9c, 0x73, 0xff, 0x4c, 0x46, 0xe0, 0x44, 0xb5, 0x9c, 0x48,
	0xd2, 0x4d, 0x98, 0x79, 0x4f, 0xa8, 0x9e, 0x54, 0xad, 0xd8, 0x83, 0xdb, 0x86, 0xaf, 0xe1, 0x16,
	0x64, 0x12, 0x4a, 0x0e, 0x5f, 0xba, 0xd8, 0x8d, 0xe8, 0x36, 0xac, 0xdb, 0x3c, 0x16, 0x3a, 0x0e,
	0xa1, 0x36, 0x3b, 0x10, 0xf2, 0xec, 0xb8, 0xd8, 0xe6, 0x0a, 0x34, 0xc4, 0xb0, 0x48, 0x3a, 0x91,
	0x74, 0xa3, 0xf8, 0x95, 0x89, 0x71, 0x17, 0xf6, 0x41, 0xc0, 0x6f, 0xa5, 0x8a, 0xdf, 0x44, 0x74,
	0x0f, 0xb2, 0x70, 0x1c, 0x37, 0xee, 0x81, 0x0c, 0x76, 0x99, 0x5f, 0x20, 0x40, 0x4c, 0x6e, 0x6e,
	0x44, 0xab, 0xc6, 0x8d, 0xe3, 0x62, 0x36, 0xc0, 0x0e, 0xba, 0xdf, 0xfb, 0x2e, 0x00, 0x00, 0xff,
	0xff, 0x08, 0xdd, 0x4a, 0xc5, 0x59, 0x0e, 0x00, 0x00,
}
