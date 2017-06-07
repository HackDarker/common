// Code generated by protoc-gen-go.
// source: tdk_base.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of cm_offline from common_client.proto

// Ignoring public import of cm_hearbeat from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

type TdkEnumProtoId int32

const (
	TdkEnumProtoId_TDK_PID_HEATBEAT               TdkEnumProtoId = 0
	TdkEnumProtoId_TDK_PID_TdkCreateDeskReq       TdkEnumProtoId = 1
	TdkEnumProtoId_TDK_PID_TdkJoinDeskReq         TdkEnumProtoId = 2
	TdkEnumProtoId_TDK_PID_TdkJoinRoomReq         TdkEnumProtoId = 3
	TdkEnumProtoId_TDK_PID_TdkLeaveDeskReq        TdkEnumProtoId = 4
	TdkEnumProtoId_TDK_PID_TdkLeaveRoomReq        TdkEnumProtoId = 5
	TdkEnumProtoId_TDK_PID_TdkDissolveDeskReq     TdkEnumProtoId = 6
	TdkEnumProtoId_TDK_PID_TdkDisDeskDesicionReq  TdkEnumProtoId = 7
	TdkEnumProtoId_TDK_PID_TdkBetReq              TdkEnumProtoId = 8
	TdkEnumProtoId_TDK_PID_TdkQiJiaoReq           TdkEnumProtoId = 9
	TdkEnumProtoId_TDK_PID_TdkFanTiReq            TdkEnumProtoId = 10
	TdkEnumProtoId_TDK_PID_TdkFoldReq             TdkEnumProtoId = 11
	TdkEnumProtoId_TDK_PID_TdkUserReadyReq        TdkEnumProtoId = 12
	TdkEnumProtoId_TDK_PID_TdkCreateDeskRsp       TdkEnumProtoId = 13
	TdkEnumProtoId_TDK_PID_TdkJoinDeskRsp         TdkEnumProtoId = 14
	TdkEnumProtoId_TDK_PID_TdkJoinRoomRsp         TdkEnumProtoId = 15
	TdkEnumProtoId_TDK_PID_TdkLeaveDeskRsp        TdkEnumProtoId = 16
	TdkEnumProtoId_TDK_PID_TdkLeaveRoomRsp        TdkEnumProtoId = 17
	TdkEnumProtoId_TDK_PID_TdkDissolveDeskUserRsp TdkEnumProtoId = 18
	TdkEnumProtoId_TDK_PID_TdkDisDeskDesicionRsp  TdkEnumProtoId = 19
	TdkEnumProtoId_TDK_PID_TdkBetRsp              TdkEnumProtoId = 20
	TdkEnumProtoId_TDK_PID_TdkQiJiaoRsp           TdkEnumProtoId = 21
	TdkEnumProtoId_TDK_PID_TdkFanTiRsp            TdkEnumProtoId = 22
	TdkEnumProtoId_TDK_PID_TdkFoldRsp             TdkEnumProtoId = 23
	TdkEnumProtoId_TDK_PID_TdkUserReadyRsp        TdkEnumProtoId = 24
	TdkEnumProtoId_TDK_PID_TdkStartBetRsp         TdkEnumProtoId = 25
	TdkEnumProtoId_TDK_PID_TTdkReturnDeskRsp      TdkEnumProtoId = 26
	TdkEnumProtoId_TDK_PID_TdkDisDeskRsp          TdkEnumProtoId = 27
	TdkEnumProtoId_TDK_PID_TdkSendPokerRsp        TdkEnumProtoId = 28
	TdkEnumProtoId_TDK_PID_TdkRoundEndRsp         TdkEnumProtoId = 29
	TdkEnumProtoId_TDK_PID_TdkPass                TdkEnumProtoId = 30
	TdkEnumProtoId_TDK_PID_TdkPassRsp             TdkEnumProtoId = 31
	TdkEnumProtoId_TDK_PID_TdkKaiPai              TdkEnumProtoId = 32
	TdkEnumProtoId_TDK_PID_TdkKaiPaiRsp           TdkEnumProtoId = 33
	TdkEnumProtoId_TDK_PID_TdkDisDeskResult       TdkEnumProtoId = 34
	TdkEnumProtoId_TDK_PID_TdkEnterGame           TdkEnumProtoId = 35
	TdkEnumProtoId_TDK_PID_TdkEnterGameRsp        TdkEnumProtoId = 36
	TdkEnumProtoId_TDK_PID_TdkZhanJi              TdkEnumProtoId = 37
	TdkEnumProtoId_TDK_PID_TdkZhanJiRsp           TdkEnumProtoId = 38
)

var TdkEnumProtoId_name = map[int32]string{
	0:  "TDK_PID_HEATBEAT",
	1:  "TDK_PID_TdkCreateDeskReq",
	2:  "TDK_PID_TdkJoinDeskReq",
	3:  "TDK_PID_TdkJoinRoomReq",
	4:  "TDK_PID_TdkLeaveDeskReq",
	5:  "TDK_PID_TdkLeaveRoomReq",
	6:  "TDK_PID_TdkDissolveDeskReq",
	7:  "TDK_PID_TdkDisDeskDesicionReq",
	8:  "TDK_PID_TdkBetReq",
	9:  "TDK_PID_TdkQiJiaoReq",
	10: "TDK_PID_TdkFanTiReq",
	11: "TDK_PID_TdkFoldReq",
	12: "TDK_PID_TdkUserReadyReq",
	13: "TDK_PID_TdkCreateDeskRsp",
	14: "TDK_PID_TdkJoinDeskRsp",
	15: "TDK_PID_TdkJoinRoomRsp",
	16: "TDK_PID_TdkLeaveDeskRsp",
	17: "TDK_PID_TdkLeaveRoomRsp",
	18: "TDK_PID_TdkDissolveDeskUserRsp",
	19: "TDK_PID_TdkDisDeskDesicionRsp",
	20: "TDK_PID_TdkBetRsp",
	21: "TDK_PID_TdkQiJiaoRsp",
	22: "TDK_PID_TdkFanTiRsp",
	23: "TDK_PID_TdkFoldRsp",
	24: "TDK_PID_TdkUserReadyRsp",
	25: "TDK_PID_TdkStartBetRsp",
	26: "TDK_PID_TTdkReturnDeskRsp",
	27: "TDK_PID_TdkDisDeskRsp",
	28: "TDK_PID_TdkSendPokerRsp",
	29: "TDK_PID_TdkRoundEndRsp",
	30: "TDK_PID_TdkPass",
	31: "TDK_PID_TdkPassRsp",
	32: "TDK_PID_TdkKaiPai",
	33: "TDK_PID_TdkKaiPaiRsp",
	34: "TDK_PID_TdkDisDeskResult",
	35: "TDK_PID_TdkEnterGame",
	36: "TDK_PID_TdkEnterGameRsp",
	37: "TDK_PID_TdkZhanJi",
	38: "TDK_PID_TdkZhanJiRsp",
}
var TdkEnumProtoId_value = map[string]int32{
	"TDK_PID_HEATBEAT":               0,
	"TDK_PID_TdkCreateDeskReq":       1,
	"TDK_PID_TdkJoinDeskReq":         2,
	"TDK_PID_TdkJoinRoomReq":         3,
	"TDK_PID_TdkLeaveDeskReq":        4,
	"TDK_PID_TdkLeaveRoomReq":        5,
	"TDK_PID_TdkDissolveDeskReq":     6,
	"TDK_PID_TdkDisDeskDesicionReq":  7,
	"TDK_PID_TdkBetReq":              8,
	"TDK_PID_TdkQiJiaoReq":           9,
	"TDK_PID_TdkFanTiReq":            10,
	"TDK_PID_TdkFoldReq":             11,
	"TDK_PID_TdkUserReadyReq":        12,
	"TDK_PID_TdkCreateDeskRsp":       13,
	"TDK_PID_TdkJoinDeskRsp":         14,
	"TDK_PID_TdkJoinRoomRsp":         15,
	"TDK_PID_TdkLeaveDeskRsp":        16,
	"TDK_PID_TdkLeaveRoomRsp":        17,
	"TDK_PID_TdkDissolveDeskUserRsp": 18,
	"TDK_PID_TdkDisDeskDesicionRsp":  19,
	"TDK_PID_TdkBetRsp":              20,
	"TDK_PID_TdkQiJiaoRsp":           21,
	"TDK_PID_TdkFanTiRsp":            22,
	"TDK_PID_TdkFoldRsp":             23,
	"TDK_PID_TdkUserReadyRsp":        24,
	"TDK_PID_TdkStartBetRsp":         25,
	"TDK_PID_TTdkReturnDeskRsp":      26,
	"TDK_PID_TdkDisDeskRsp":          27,
	"TDK_PID_TdkSendPokerRsp":        28,
	"TDK_PID_TdkRoundEndRsp":         29,
	"TDK_PID_TdkPass":                30,
	"TDK_PID_TdkPassRsp":             31,
	"TDK_PID_TdkKaiPai":              32,
	"TDK_PID_TdkKaiPaiRsp":           33,
	"TDK_PID_TdkDisDeskResult":       34,
	"TDK_PID_TdkEnterGame":           35,
	"TDK_PID_TdkEnterGameRsp":        36,
	"TDK_PID_TdkZhanJi":              37,
	"TDK_PID_TdkZhanJiRsp":           38,
}

func (x TdkEnumProtoId) Enum() *TdkEnumProtoId {
	p := new(TdkEnumProtoId)
	*p = x
	return p
}
func (x TdkEnumProtoId) String() string {
	return proto.EnumName(TdkEnumProtoId_name, int32(x))
}
func (x *TdkEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumProtoId_value, data, "TdkEnumProtoId")
	if err != nil {
		return err
	}
	*x = TdkEnumProtoId(value)
	return nil
}
func (TdkEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

type TdkEnumDeskstatus int32

const (
	TdkEnumDeskstatus_DESKSTATUS_IDLE                TdkEnumDeskstatus = 0
	TdkEnumDeskstatus_DESKSTATUS_WAITINGJOIN         TdkEnumDeskstatus = 1
	TdkEnumDeskstatus_DESKSTATUS_GAMING_BET          TdkEnumDeskstatus = 2
	TdkEnumDeskstatus_DESKSTATUS_GAMING_QIJIAO       TdkEnumDeskstatus = 3
	TdkEnumDeskstatus_DESKSTATUS_GAMING_FANTI        TdkEnumDeskstatus = 4
	TdkEnumDeskstatus_DESKSTATUS_GAMING_BET_GEN      TdkEnumDeskstatus = 5
	TdkEnumDeskstatus_DESKSTATUS_GAMING_QIJIAO_GEN   TdkEnumDeskstatus = 6
	TdkEnumDeskstatus_DESKSTATUS_GAMING_FANTI_GEN    TdkEnumDeskstatus = 7
	TdkEnumDeskstatus_DESKSTATUS_WAITINGFORREADY     TdkEnumDeskstatus = 8
	TdkEnumDeskstatus_DESKSTATUS_GAMING_LANGUO       TdkEnumDeskstatus = 9
	TdkEnumDeskstatus_DESKSTATUS_GAMING_END          TdkEnumDeskstatus = 10
	TdkEnumDeskstatus_DESKSTATUS_GAMING_KAIPAI       TdkEnumDeskstatus = 11
	TdkEnumDeskstatus_DESKSTATUS_GAMING_SENDALLPOKER TdkEnumDeskstatus = 12
	TdkEnumDeskstatus_DESKSTATUS_GAMING_SCORECMP     TdkEnumDeskstatus = 13
	TdkEnumDeskstatus_DESKSTATUS_ZHANJI              TdkEnumDeskstatus = 14
)

var TdkEnumDeskstatus_name = map[int32]string{
	0:  "DESKSTATUS_IDLE",
	1:  "DESKSTATUS_WAITINGJOIN",
	2:  "DESKSTATUS_GAMING_BET",
	3:  "DESKSTATUS_GAMING_QIJIAO",
	4:  "DESKSTATUS_GAMING_FANTI",
	5:  "DESKSTATUS_GAMING_BET_GEN",
	6:  "DESKSTATUS_GAMING_QIJIAO_GEN",
	7:  "DESKSTATUS_GAMING_FANTI_GEN",
	8:  "DESKSTATUS_WAITINGFORREADY",
	9:  "DESKSTATUS_GAMING_LANGUO",
	10: "DESKSTATUS_GAMING_END",
	11: "DESKSTATUS_GAMING_KAIPAI",
	12: "DESKSTATUS_GAMING_SENDALLPOKER",
	13: "DESKSTATUS_GAMING_SCORECMP",
	14: "DESKSTATUS_ZHANJI",
}
var TdkEnumDeskstatus_value = map[string]int32{
	"DESKSTATUS_IDLE":                0,
	"DESKSTATUS_WAITINGJOIN":         1,
	"DESKSTATUS_GAMING_BET":          2,
	"DESKSTATUS_GAMING_QIJIAO":       3,
	"DESKSTATUS_GAMING_FANTI":        4,
	"DESKSTATUS_GAMING_BET_GEN":      5,
	"DESKSTATUS_GAMING_QIJIAO_GEN":   6,
	"DESKSTATUS_GAMING_FANTI_GEN":    7,
	"DESKSTATUS_WAITINGFORREADY":     8,
	"DESKSTATUS_GAMING_LANGUO":       9,
	"DESKSTATUS_GAMING_END":          10,
	"DESKSTATUS_GAMING_KAIPAI":       11,
	"DESKSTATUS_GAMING_SENDALLPOKER": 12,
	"DESKSTATUS_GAMING_SCORECMP":     13,
	"DESKSTATUS_ZHANJI":              14,
}

func (x TdkEnumDeskstatus) Enum() *TdkEnumDeskstatus {
	p := new(TdkEnumDeskstatus)
	*p = x
	return p
}
func (x TdkEnumDeskstatus) String() string {
	return proto.EnumName(TdkEnumDeskstatus_name, int32(x))
}
func (x *TdkEnumDeskstatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumDeskstatus_value, data, "TdkEnumDeskstatus")
	if err != nil {
		return err
	}
	*x = TdkEnumDeskstatus(value)
	return nil
}
func (TdkEnumDeskstatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

type TdkEnumRadyrsp int32

const (
	TdkEnumRadyrsp_DESKREADYRSP_SUCCESS         TdkEnumRadyrsp = 0
	TdkEnumRadyrsp_DESKREADYRSP_DESKSTATUSERROR TdkEnumRadyrsp = 1
)

var TdkEnumRadyrsp_name = map[int32]string{
	0: "DESKREADYRSP_SUCCESS",
	1: "DESKREADYRSP_DESKSTATUSERROR",
}
var TdkEnumRadyrsp_value = map[string]int32{
	"DESKREADYRSP_SUCCESS":         0,
	"DESKREADYRSP_DESKSTATUSERROR": 1,
}

func (x TdkEnumRadyrsp) Enum() *TdkEnumRadyrsp {
	p := new(TdkEnumRadyrsp)
	*p = x
	return p
}
func (x TdkEnumRadyrsp) String() string {
	return proto.EnumName(TdkEnumRadyrsp_name, int32(x))
}
func (x *TdkEnumRadyrsp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumRadyrsp_value, data, "TdkEnumRadyrsp")
	if err != nil {
		return err
	}
	*x = TdkEnumRadyrsp(value)
	return nil
}
func (TdkEnumRadyrsp) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

type TdkEnumDisdesicionrsp int32

const (
	TdkEnumDisdesicionrsp_DESKDISDESICIONRSP_SUCCESS    TdkEnumDisdesicionrsp = 0
	TdkEnumDisdesicionrsp_DESKDISDESICIONRSP_DECIDED    TdkEnumDisdesicionrsp = 1
	TdkEnumDisdesicionrsp_DESKDISDESICIONRSP_STATEERROR TdkEnumDisdesicionrsp = 2
)

var TdkEnumDisdesicionrsp_name = map[int32]string{
	0: "DESKDISDESICIONRSP_SUCCESS",
	1: "DESKDISDESICIONRSP_DECIDED",
	2: "DESKDISDESICIONRSP_STATEERROR",
}
var TdkEnumDisdesicionrsp_value = map[string]int32{
	"DESKDISDESICIONRSP_SUCCESS":    0,
	"DESKDISDESICIONRSP_DECIDED":    1,
	"DESKDISDESICIONRSP_STATEERROR": 2,
}

func (x TdkEnumDisdesicionrsp) Enum() *TdkEnumDisdesicionrsp {
	p := new(TdkEnumDisdesicionrsp)
	*p = x
	return p
}
func (x TdkEnumDisdesicionrsp) String() string {
	return proto.EnumName(TdkEnumDisdesicionrsp_name, int32(x))
}
func (x *TdkEnumDisdesicionrsp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumDisdesicionrsp_value, data, "TdkEnumDisdesicionrsp")
	if err != nil {
		return err
	}
	*x = TdkEnumDisdesicionrsp(value)
	return nil
}
func (TdkEnumDisdesicionrsp) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{3} }

type TdkEnumUserdisdesicion int32

const (
	TdkEnumUserdisdesicion_USERDISDESK_NOTDECIDE TdkEnumUserdisdesicion = 0
	TdkEnumUserdisdesicion_USERDISDESK_YES       TdkEnumUserdisdesicion = 1
	TdkEnumUserdisdesicion_USERDISDESK_NO        TdkEnumUserdisdesicion = 2
)

var TdkEnumUserdisdesicion_name = map[int32]string{
	0: "USERDISDESK_NOTDECIDE",
	1: "USERDISDESK_YES",
	2: "USERDISDESK_NO",
}
var TdkEnumUserdisdesicion_value = map[string]int32{
	"USERDISDESK_NOTDECIDE": 0,
	"USERDISDESK_YES":       1,
	"USERDISDESK_NO":        2,
}

func (x TdkEnumUserdisdesicion) Enum() *TdkEnumUserdisdesicion {
	p := new(TdkEnumUserdisdesicion)
	*p = x
	return p
}
func (x TdkEnumUserdisdesicion) String() string {
	return proto.EnumName(TdkEnumUserdisdesicion_name, int32(x))
}
func (x *TdkEnumUserdisdesicion) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumUserdisdesicion_value, data, "TdkEnumUserdisdesicion")
	if err != nil {
		return err
	}
	*x = TdkEnumUserdisdesicion(value)
	return nil
}
func (TdkEnumUserdisdesicion) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{4} }

type TdkEnumDeskstartrsp int32

const (
	TdkEnumDeskstartrsp_DESKSTARTRSP_SUCCESS     TdkEnumDeskstartrsp = 0
	TdkEnumDeskstartrsp_DESKSTARTRSP_NOTALLREADY TdkEnumDeskstartrsp = 1
	TdkEnumDeskstartrsp_DESKSTARTRSP_NOTOWNER    TdkEnumDeskstartrsp = 2
)

var TdkEnumDeskstartrsp_name = map[int32]string{
	0: "DESKSTARTRSP_SUCCESS",
	1: "DESKSTARTRSP_NOTALLREADY",
	2: "DESKSTARTRSP_NOTOWNER",
}
var TdkEnumDeskstartrsp_value = map[string]int32{
	"DESKSTARTRSP_SUCCESS":     0,
	"DESKSTARTRSP_NOTALLREADY": 1,
	"DESKSTARTRSP_NOTOWNER":    2,
}

func (x TdkEnumDeskstartrsp) Enum() *TdkEnumDeskstartrsp {
	p := new(TdkEnumDeskstartrsp)
	*p = x
	return p
}
func (x TdkEnumDeskstartrsp) String() string {
	return proto.EnumName(TdkEnumDeskstartrsp_name, int32(x))
}
func (x *TdkEnumDeskstartrsp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumDeskstartrsp_value, data, "TdkEnumDeskstartrsp")
	if err != nil {
		return err
	}
	*x = TdkEnumDeskstartrsp(value)
	return nil
}
func (TdkEnumDeskstartrsp) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{5} }

type TdkEnumBetrsp int32

const (
	TdkEnumBetrsp_DESKBETRSP_SUCCESS             TdkEnumBetrsp = 0
	TdkEnumBetrsp_DESKBETRSP_NOTACTUSER          TdkEnumBetrsp = 1
	TdkEnumBetrsp_DESKBETRSP_BETNUMERROR         TdkEnumBetrsp = 2
	TdkEnumBetrsp_DESKBETRSP_NOTINOPERATIONQUEUE TdkEnumBetrsp = 3
	TdkEnumBetrsp_DESKBETRSP_DESKSTATEERROR      TdkEnumBetrsp = 4
	TdkEnumBetrsp_DESKBETRSP_CANNOTALLIN         TdkEnumBetrsp = 5
)

var TdkEnumBetrsp_name = map[int32]string{
	0: "DESKBETRSP_SUCCESS",
	1: "DESKBETRSP_NOTACTUSER",
	2: "DESKBETRSP_BETNUMERROR",
	3: "DESKBETRSP_NOTINOPERATIONQUEUE",
	4: "DESKBETRSP_DESKSTATEERROR",
	5: "DESKBETRSP_CANNOTALLIN",
}
var TdkEnumBetrsp_value = map[string]int32{
	"DESKBETRSP_SUCCESS":             0,
	"DESKBETRSP_NOTACTUSER":          1,
	"DESKBETRSP_BETNUMERROR":         2,
	"DESKBETRSP_NOTINOPERATIONQUEUE": 3,
	"DESKBETRSP_DESKSTATEERROR":      4,
	"DESKBETRSP_CANNOTALLIN":         5,
}

func (x TdkEnumBetrsp) Enum() *TdkEnumBetrsp {
	p := new(TdkEnumBetrsp)
	*p = x
	return p
}
func (x TdkEnumBetrsp) String() string {
	return proto.EnumName(TdkEnumBetrsp_name, int32(x))
}
func (x *TdkEnumBetrsp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumBetrsp_value, data, "TdkEnumBetrsp")
	if err != nil {
		return err
	}
	*x = TdkEnumBetrsp(value)
	return nil
}
func (TdkEnumBetrsp) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{6} }

type TdkEnumQijiaorsp int32

const (
	TdkEnumQijiaorsp_DESKQIJIAORSP_SUCCESS        TdkEnumQijiaorsp = 0
	TdkEnumQijiaorsp_DESKQIJIAORSP_NOTACTUSER     TdkEnumQijiaorsp = 1
	TdkEnumQijiaorsp_DESKQIJIAORSP_DESKSTATEERROR TdkEnumQijiaorsp = 2
	TdkEnumQijiaorsp_DESKQIJIAORSP_CANNOTALLIN    TdkEnumQijiaorsp = 3
)

var TdkEnumQijiaorsp_name = map[int32]string{
	0: "DESKQIJIAORSP_SUCCESS",
	1: "DESKQIJIAORSP_NOTACTUSER",
	2: "DESKQIJIAORSP_DESKSTATEERROR",
	3: "DESKQIJIAORSP_CANNOTALLIN",
}
var TdkEnumQijiaorsp_value = map[string]int32{
	"DESKQIJIAORSP_SUCCESS":        0,
	"DESKQIJIAORSP_NOTACTUSER":     1,
	"DESKQIJIAORSP_DESKSTATEERROR": 2,
	"DESKQIJIAORSP_CANNOTALLIN":    3,
}

func (x TdkEnumQijiaorsp) Enum() *TdkEnumQijiaorsp {
	p := new(TdkEnumQijiaorsp)
	*p = x
	return p
}
func (x TdkEnumQijiaorsp) String() string {
	return proto.EnumName(TdkEnumQijiaorsp_name, int32(x))
}
func (x *TdkEnumQijiaorsp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumQijiaorsp_value, data, "TdkEnumQijiaorsp")
	if err != nil {
		return err
	}
	*x = TdkEnumQijiaorsp(value)
	return nil
}
func (TdkEnumQijiaorsp) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{7} }

type TdkEnumFantirsp int32

const (
	TdkEnumFantirsp_DESKFANTIRSP_SUCCESS        TdkEnumFantirsp = 0
	TdkEnumFantirsp_DESKFANTIRSP_DESKSTATEERROR TdkEnumFantirsp = 1
	TdkEnumFantirsp_DESKFANTIRSP_NOTACTUSER     TdkEnumFantirsp = 2
	TdkEnumFantirsp_DESKFANTIRSP_BETNUMERROR    TdkEnumFantirsp = 3
	TdkEnumFantirsp_DESKFANTIRSP_CANNOTALLIN    TdkEnumFantirsp = 4
)

var TdkEnumFantirsp_name = map[int32]string{
	0: "DESKFANTIRSP_SUCCESS",
	1: "DESKFANTIRSP_DESKSTATEERROR",
	2: "DESKFANTIRSP_NOTACTUSER",
	3: "DESKFANTIRSP_BETNUMERROR",
	4: "DESKFANTIRSP_CANNOTALLIN",
}
var TdkEnumFantirsp_value = map[string]int32{
	"DESKFANTIRSP_SUCCESS":        0,
	"DESKFANTIRSP_DESKSTATEERROR": 1,
	"DESKFANTIRSP_NOTACTUSER":     2,
	"DESKFANTIRSP_BETNUMERROR":    3,
	"DESKFANTIRSP_CANNOTALLIN":    4,
}

func (x TdkEnumFantirsp) Enum() *TdkEnumFantirsp {
	p := new(TdkEnumFantirsp)
	*p = x
	return p
}
func (x TdkEnumFantirsp) String() string {
	return proto.EnumName(TdkEnumFantirsp_name, int32(x))
}
func (x *TdkEnumFantirsp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumFantirsp_value, data, "TdkEnumFantirsp")
	if err != nil {
		return err
	}
	*x = TdkEnumFantirsp(value)
	return nil
}
func (TdkEnumFantirsp) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{8} }

type TdkEnumFoldrsp int32

const (
	TdkEnumFoldrsp_DESKFOLDRSP_SUCCESS        TdkEnumFoldrsp = 0
	TdkEnumFoldrsp_DESKFOLDRSP_NOTACTUSER     TdkEnumFoldrsp = 1
	TdkEnumFoldrsp_DESKFOLDRSP_DESKSTATEERROR TdkEnumFoldrsp = 2
)

var TdkEnumFoldrsp_name = map[int32]string{
	0: "DESKFOLDRSP_SUCCESS",
	1: "DESKFOLDRSP_NOTACTUSER",
	2: "DESKFOLDRSP_DESKSTATEERROR",
}
var TdkEnumFoldrsp_value = map[string]int32{
	"DESKFOLDRSP_SUCCESS":        0,
	"DESKFOLDRSP_NOTACTUSER":     1,
	"DESKFOLDRSP_DESKSTATEERROR": 2,
}

func (x TdkEnumFoldrsp) Enum() *TdkEnumFoldrsp {
	p := new(TdkEnumFoldrsp)
	*p = x
	return p
}
func (x TdkEnumFoldrsp) String() string {
	return proto.EnumName(TdkEnumFoldrsp_name, int32(x))
}
func (x *TdkEnumFoldrsp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumFoldrsp_value, data, "TdkEnumFoldrsp")
	if err != nil {
		return err
	}
	*x = TdkEnumFoldrsp(value)
	return nil
}
func (TdkEnumFoldrsp) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{9} }

type TdkEnumDeskendtype int32

const (
	TdkEnumDeskendtype_ROUNDENDTYPE_WIN    TdkEnumDeskendtype = 0
	TdkEnumDeskendtype_ROUNDENDTYPE_LANGUO TdkEnumDeskendtype = 1
)

var TdkEnumDeskendtype_name = map[int32]string{
	0: "ROUNDENDTYPE_WIN",
	1: "ROUNDENDTYPE_LANGUO",
}
var TdkEnumDeskendtype_value = map[string]int32{
	"ROUNDENDTYPE_WIN":    0,
	"ROUNDENDTYPE_LANGUO": 1,
}

func (x TdkEnumDeskendtype) Enum() *TdkEnumDeskendtype {
	p := new(TdkEnumDeskendtype)
	*p = x
	return p
}
func (x TdkEnumDeskendtype) String() string {
	return proto.EnumName(TdkEnumDeskendtype_name, int32(x))
}
func (x *TdkEnumDeskendtype) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumDeskendtype_value, data, "TdkEnumDeskendtype")
	if err != nil {
		return err
	}
	*x = TdkEnumDeskendtype(value)
	return nil
}
func (TdkEnumDeskendtype) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{10} }

type TdkEnumReqdisdeskrsp int32

const (
	TdkEnumReqdisdeskrsp_DESKREQDISDESKRSP_SUCCESS    TdkEnumReqdisdeskrsp = 0
	TdkEnumReqdisdeskrsp_DESKREQDISDESKRSP_STATEERROR TdkEnumReqdisdeskrsp = 1
	TdkEnumReqdisdeskrsp_DESKREQDISDESKRSP_DISING     TdkEnumReqdisdeskrsp = 2
)

var TdkEnumReqdisdeskrsp_name = map[int32]string{
	0: "DESKREQDISDESKRSP_SUCCESS",
	1: "DESKREQDISDESKRSP_STATEERROR",
	2: "DESKREQDISDESKRSP_DISING",
}
var TdkEnumReqdisdeskrsp_value = map[string]int32{
	"DESKREQDISDESKRSP_SUCCESS":    0,
	"DESKREQDISDESKRSP_STATEERROR": 1,
	"DESKREQDISDESKRSP_DISING":     2,
}

func (x TdkEnumReqdisdeskrsp) Enum() *TdkEnumReqdisdeskrsp {
	p := new(TdkEnumReqdisdeskrsp)
	*p = x
	return p
}
func (x TdkEnumReqdisdeskrsp) String() string {
	return proto.EnumName(TdkEnumReqdisdeskrsp_name, int32(x))
}
func (x *TdkEnumReqdisdeskrsp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TdkEnumReqdisdeskrsp_value, data, "TdkEnumReqdisdeskrsp")
	if err != nil {
		return err
	}
	*x = TdkEnumReqdisdeskrsp(value)
	return nil
}
func (TdkEnumReqdisdeskrsp) EnumDescriptor() ([]byte, []int) { return fileDescriptor15, []int{11} }

func init() {
	proto.RegisterEnum("yjprotogo.TdkEnumProtoId", TdkEnumProtoId_name, TdkEnumProtoId_value)
	proto.RegisterEnum("yjprotogo.TdkEnumDeskstatus", TdkEnumDeskstatus_name, TdkEnumDeskstatus_value)
	proto.RegisterEnum("yjprotogo.TdkEnumRadyrsp", TdkEnumRadyrsp_name, TdkEnumRadyrsp_value)
	proto.RegisterEnum("yjprotogo.TdkEnumDisdesicionrsp", TdkEnumDisdesicionrsp_name, TdkEnumDisdesicionrsp_value)
	proto.RegisterEnum("yjprotogo.TdkEnumUserdisdesicion", TdkEnumUserdisdesicion_name, TdkEnumUserdisdesicion_value)
	proto.RegisterEnum("yjprotogo.TdkEnumDeskstartrsp", TdkEnumDeskstartrsp_name, TdkEnumDeskstartrsp_value)
	proto.RegisterEnum("yjprotogo.TdkEnumBetrsp", TdkEnumBetrsp_name, TdkEnumBetrsp_value)
	proto.RegisterEnum("yjprotogo.TdkEnumQijiaorsp", TdkEnumQijiaorsp_name, TdkEnumQijiaorsp_value)
	proto.RegisterEnum("yjprotogo.TdkEnumFantirsp", TdkEnumFantirsp_name, TdkEnumFantirsp_value)
	proto.RegisterEnum("yjprotogo.TdkEnumFoldrsp", TdkEnumFoldrsp_name, TdkEnumFoldrsp_value)
	proto.RegisterEnum("yjprotogo.TdkEnumDeskendtype", TdkEnumDeskendtype_name, TdkEnumDeskendtype_value)
	proto.RegisterEnum("yjprotogo.TdkEnumReqdisdeskrsp", TdkEnumReqdisdeskrsp_name, TdkEnumReqdisdeskrsp_value)
}

var fileDescriptor15 = []byte{
	// 1054 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x96, 0x5b, 0x72, 0xdb, 0x36,
	0x17, 0xc7, 0x25, 0xd9, 0x4e, 0x6c, 0x24, 0xb1, 0x8f, 0x21, 0xdf, 0x6f, 0xb9, 0x7c, 0x5f, 0xfb,
	0xa0, 0x87, 0xee, 0x81, 0x16, 0x61, 0x05, 0x92, 0x0c, 0xd2, 0x24, 0x35, 0x1e, 0xe7, 0xa1, 0x1c,
	0xc6, 0x44, 0x52, 0x5a, 0x36, 0x49, 0x0b, 0x54, 0x3b, 0x9e, 0xe9, 0x1a, 0xba, 0x8d, 0x2e, 0xa3,
	0x0f, 0xdd, 0x40, 0x97, 0xd4, 0x01, 0x6f, 0x02, 0x29, 0x6a, 0xfa, 0xa6, 0xc1, 0xef, 0x9c, 0x83,
	0xff, 0xf9, 0x03, 0x38, 0x14, 0xda, 0x4e, 0xfc, 0xa9, 0xfb, 0xd5, 0x13, 0xfc, 0xa7, 0x78, 0x16,
	0x25, 0x11, 0xde, 0x7a, 0x79, 0x48, 0x7f, 0x7c, 0x8f, 0x4e, 0xba, 0xf7, 0xd1, 0xd3, 0x53, 0x14,
	0xba, 0xf7, 0x8f, 0x01, 0x0f, 0x93, 0x8c, 0xf7, 0xfe, 0xde, 0x44, 0x20, 0x53, 0x78, 0x38, 0x7f,
	0x72, 0xd3, 0x25, 0xea, 0xe3, 0x3d, 0x04, 0x8e, 0x3e, 0x72, 0x4d, 0xaa, 0xbb, 0x9f, 0x89, 0xe6,
	0x5c, 0x12, 0xcd, 0x81, 0x16, 0x3e, 0x43, 0x47, 0xc5, 0xaa, 0xe3, 0x4f, 0xfb, 0x33, 0xee, 0x25,
	0x5c, 0xe7, 0x62, 0x6a, 0xf1, 0x67, 0x68, 0xe3, 0x13, 0x74, 0xa0, 0xd0, 0x61, 0x14, 0x84, 0x05,
	0xeb, 0x34, 0x30, 0x2b, 0x8a, 0x9e, 0x24, 0x5b, 0xc3, 0xa7, 0xe8, 0x50, 0x61, 0x63, 0xee, 0xfd,
	0x5a, 0x16, 0x5d, 0x6f, 0x82, 0x45, 0xe6, 0x06, 0xbe, 0x40, 0x27, 0x0a, 0xd4, 0x03, 0x21, 0xa2,
	0xc7, 0x45, 0xf2, 0x2b, 0xfc, 0x11, 0x9d, 0x57, 0xb9, 0x44, 0x3a, 0x17, 0xc1, 0x7d, 0x10, 0x85,
	0x32, 0xe4, 0x35, 0xde, 0x47, 0xbb, 0x4a, 0xc8, 0x25, 0x4f, 0xe4, 0xf2, 0x26, 0x3e, 0x42, 0x7b,
	0xca, 0xf2, 0x4d, 0x30, 0x0c, 0xbc, 0x48, 0x92, 0x2d, 0x7c, 0x88, 0xba, 0x0a, 0xb9, 0xf2, 0x42,
	0x27, 0x90, 0x00, 0xe1, 0x03, 0x84, 0x55, 0x10, 0x3d, 0xfa, 0x72, 0xfd, 0x4d, 0xad, 0x83, 0x89,
	0xe0, 0x33, 0x8b, 0x7b, 0xfe, 0x8b, 0x84, 0x6f, 0x57, 0x3b, 0x2a, 0x62, 0x78, 0xb7, 0xca, 0x51,
	0x11, 0xc3, 0xf6, 0x2a, 0x47, 0x45, 0x0c, 0x3b, 0x2b, 0x1d, 0x15, 0x31, 0xc0, 0x4a, 0x47, 0x45,
	0x0c, 0xbb, 0xf8, 0x13, 0xba, 0x58, 0xe1, 0x68, 0x2a, 0x5c, 0xc4, 0x80, 0xff, 0xc3, 0x55, 0x11,
	0x43, 0xb7, 0xc1, 0x55, 0x11, 0xc3, 0x5e, 0xb3, 0xab, 0x22, 0x86, 0xfd, 0x46, 0x57, 0x45, 0x0c,
	0x07, 0x4d, 0xae, 0x8a, 0x18, 0x0e, 0x57, 0xba, 0x2a, 0x62, 0x38, 0xaa, 0x79, 0x63, 0x27, 0xde,
	0x2c, 0xc9, 0x35, 0x1c, 0xe3, 0x73, 0x74, 0x5c, 0x32, 0xc7, 0x9f, 0x5a, 0x3c, 0x99, 0xcf, 0x4a,
	0x5b, 0x4f, 0xf0, 0x31, 0xda, 0x5f, 0x6e, 0x4e, 0xa2, 0xd3, 0xda, 0x96, 0x36, 0x0f, 0x7d, 0x33,
	0x9a, 0x66, 0xa6, 0x9c, 0xd5, 0xb6, 0xb4, 0xa2, 0x79, 0xe8, 0x93, 0x30, 0xd5, 0x7a, 0x8e, 0xbb,
	0x68, 0x47, 0x61, 0xa6, 0x27, 0x04, 0x5c, 0xd4, 0x1a, 0x93, 0x8b, 0x32, 0xf8, 0x7d, 0xcd, 0xba,
	0x91, 0x17, 0x98, 0x5e, 0x00, 0x1f, 0x6a, 0xd6, 0x65, 0xcb, 0x32, 0xe1, 0x63, 0xed, 0x0a, 0x15,
	0x8a, 0xb9, 0x98, 0x3f, 0x26, 0xf0, 0xa9, 0x96, 0x47, 0xc2, 0x84, 0xcf, 0x06, 0xde, 0x13, 0x87,
	0xff, 0xd5, 0xda, 0x29, 0x89, 0x2c, 0xfa, 0xff, 0x9a, 0x8a, 0x2f, 0xbf, 0x78, 0xe1, 0x30, 0x80,
	0x1f, 0x6a, 0xd5, 0xb2, 0x65, 0x99, 0xf0, 0x63, 0xef, 0x9f, 0x35, 0xd4, 0x2d, 0xa7, 0x88, 0xcf,
	0xc5, 0x54, 0x24, 0x5e, 0x32, 0x17, 0xb2, 0x77, 0x9d, 0xd8, 0x23, 0xdb, 0xd1, 0x9c, 0x89, 0xed,
	0x52, 0x7d, 0x4c, 0xa0, 0x25, 0xcd, 0x52, 0x16, 0x6f, 0x35, 0xea, 0x50, 0x36, 0x18, 0x1a, 0x94,
	0x41, 0x5b, 0x1e, 0x80, 0xc2, 0x06, 0xda, 0x35, 0x65, 0x03, 0xf7, 0x92, 0x38, 0xd0, 0x91, 0x9d,
	0x2e, 0xa3, 0x1b, 0x3a, 0xa4, 0x9a, 0x91, 0x8d, 0x91, 0x65, 0x7a, 0xa5, 0x31, 0x87, 0xc2, 0xba,
	0x3c, 0xf5, 0xc6, 0xaa, 0xee, 0x80, 0x30, 0xd8, 0xc0, 0x1f, 0xd0, 0xd9, 0xaa, 0xca, 0x69, 0xc4,
	0x2b, 0xfc, 0x1e, 0x9d, 0xae, 0xa8, 0x9e, 0x06, 0xbc, 0x96, 0xb3, 0x68, 0xb9, 0xa7, 0x2b, 0xc3,
	0xb2, 0x88, 0xa6, 0xdf, 0xc1, 0x66, 0xb3, 0xf8, 0xb1, 0xc6, 0x06, 0x13, 0x03, 0xb6, 0x9a, 0xbb,
	0x26, 0x4c, 0x07, 0xd4, 0x9c, 0x38, 0xd2, 0xa8, 0xa9, 0x51, 0x78, 0x23, 0x1f, 0xec, 0x32, 0xb5,
	0x09, 0xd3, 0xb5, 0xf1, 0xd8, 0x34, 0x46, 0xc4, 0x82, 0xb7, 0x35, 0x69, 0x45, 0x4c, 0xdf, 0xb0,
	0x48, 0xff, 0xda, 0x84, 0x77, 0xf2, 0xb0, 0x15, 0xfe, 0xe5, 0xb3, 0xc6, 0x86, 0x14, 0xb6, 0x7b,
	0x4c, 0xf9, 0x2e, 0xcc, 0x3c, 0xff, 0x65, 0x26, 0x62, 0x79, 0x01, 0x64, 0x68, 0xda, 0x94, 0x65,
	0x9b, 0xae, 0x3d, 0xe9, 0xf7, 0x89, 0x6d, 0x43, 0xab, 0xb0, 0xb0, 0x24, 0x8b, 0x8a, 0xc4, 0xb2,
	0x0c, 0x0b, 0xda, 0xbd, 0xdf, 0xd1, 0xe1, 0xe2, 0x86, 0x04, 0xc2, 0xcf, 0x27, 0x86, 0x2c, 0x9b,
	0x2b, 0xd4, 0xa9, 0xad, 0x13, 0x9b, 0xf6, 0xa9, 0xc1, 0xaa, 0xc5, 0x9b, 0xb9, 0x4e, 0xfa, 0x54,
	0x27, 0x3a, 0xb4, 0xe5, 0x48, 0x6a, 0xca, 0x77, 0x34, 0x87, 0x64, 0xbb, 0x77, 0x7a, 0x3f, 0xa3,
	0xa3, 0x72, 0xf7, 0xb9, 0xe0, 0x33, 0x45, 0x81, 0x74, 0x5f, 0xca, 0xcc, 0xd2, 0x47, 0x2e, 0x33,
	0x9c, 0xac, 0x34, 0xb4, 0xe4, 0xfd, 0x55, 0xd1, 0x1d, 0xb1, 0xa1, 0x8d, 0x31, 0xda, 0xae, 0xc6,
	0x43, 0xa7, 0xf7, 0x80, 0xf6, 0xeb, 0xf7, 0x7f, 0x96, 0x28, 0x96, 0xd9, 0x8e, 0x66, 0x39, 0xd5,
	0xae, 0x16, 0x27, 0x9b, 0x11, 0x66, 0x38, 0xda, 0x78, 0x9c, 0x5d, 0x18, 0xf5, 0x21, 0x94, 0xd4,
	0xb8, 0x65, 0x44, 0xf6, 0xf2, 0x57, 0x1b, 0xed, 0x94, 0x9b, 0x7d, 0xe5, 0xe9, 0x36, 0x07, 0x08,
	0xcb, 0xf0, 0x4b, 0x52, 0xdb, 0x24, 0x2f, 0x93, 0xaf, 0xcb, 0x2d, 0xfa, 0xe9, 0x99, 0x64, 0x1f,
	0x6c, 0x05, 0x5d, 0x12, 0x87, 0x4d, 0xae, 0x73, 0xbb, 0x8a, 0x7b, 0xb5, 0x48, 0xa3, 0xcc, 0x30,
	0x89, 0xa5, 0x39, 0xd4, 0x60, 0x37, 0x13, 0x32, 0x21, 0xb0, 0x56, 0x3c, 0xaa, 0x3c, 0xa6, 0x38,
	0xf0, 0xdc, 0xf1, 0xf5, 0x5a, 0xf9, 0xbe, 0xc6, 0xb2, 0xfe, 0x28, 0x83, 0x8d, 0xde, 0x1f, 0x6d,
	0x84, 0xcb, 0x0e, 0x9e, 0x83, 0x87, 0xc0, 0x8b, 0x64, 0x13, 0xb9, 0xd8, 0xec, 0xe5, 0x35, 0x9a,
	0xb5, 0x40, 0x95, 0x56, 0xf2, 0xdb, 0xb7, 0xa0, 0x35, 0x35, 0x9d, 0x42, 0xec, 0x22, 0x42, 0x15,
	0xb4, 0xd6, 0xfb, 0xb3, 0x8d, 0x76, 0x4b, 0x41, 0xdf, 0xbc, 0x30, 0x09, 0x94, 0xb3, 0x4b, 0xdf,
	0x79, 0x55, 0x4e, 0x3e, 0x0f, 0x4a, 0x52, 0xdb, 0xaf, 0x5d, 0x8c, 0xa3, 0x32, 0x40, 0x91, 0x5b,
	0x4e, 0xb2, 0x12, 0xaa, 0xde, 0xaf, 0x2d, 0x51, 0x55, 0xe9, 0x7a, 0xef, 0xbb, 0xf2, 0x2c, 0xbf,
	0x45, 0x8f, 0xbe, 0xd4, 0x79, 0x88, 0xba, 0x69, 0x86, 0x31, 0xd6, 0xab, 0x32, 0xf3, 0x33, 0x28,
	0x40, 0xc5, 0xb3, 0xfc, 0x51, 0x15, 0xac, 0xee, 0x58, 0x8f, 0xa0, 0xbd, 0xca, 0x8d, 0xe6, 0xa1,
	0x9f, 0xbc, 0xc4, 0x5c, 0xfe, 0x37, 0xb4, 0x8c, 0x09, 0xd3, 0x09, 0xd3, 0x9d, 0x3b, 0x93, 0xb8,
	0xb7, 0x94, 0x41, 0x4b, 0x4a, 0xa8, 0xac, 0xe6, 0xa3, 0xad, 0xdd, 0xfb, 0x0d, 0x1d, 0x2c, 0xc6,
	0x08, 0x7f, 0xce, 0xde, 0xdd, 0x54, 0xaa, 0xce, 0x8f, 0xc4, 0x22, 0x37, 0xf9, 0x4b, 0x5a, 0x31,
	0x51, 0x2a, 0x58, 0xf5, 0x38, 0x37, 0xaa, 0x1a, 0xa1, 0x53, 0x9b, 0xb2, 0x01, 0x74, 0xcc, 0xd6,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x68, 0x30, 0x49, 0xd0, 0x0b, 0x0b, 0x00, 0x00,
}
