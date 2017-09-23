// Code generated by protoc-gen-go. DO NOT EDIT.
// source: phz_desk.proto

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

// Ignoring public import of common_req_list_coin_desk from common_client.proto

// Ignoring public import of common_ack_list_coin_desk from common_client.proto

// Ignoring public import of CommonCoinDeskInfo from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of phz_base_createOption from phz_base.proto

// Ignoring public import of phz_base_roomInfo from phz_base.proto

// Ignoring public import of phz_base_deskInfo from phz_base.proto

// Ignoring public import of phz_base_playerInfo from phz_base.proto

// Ignoring public import of phz_enum_protoId from phz_base.proto

// Ignoring public import of phz_enum_roomType from phz_base.proto

// Ignoring public import of phz_enum_tiType from phz_base.proto

// Ignoring public import of phz_enum_pengType from phz_base.proto

// Ignoring public import of phz_enum_paoType from phz_base.proto

// 创建房间
type PhzReqCreateDesk struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomInfo         *PhzBaseCreateOption `protobuf:"bytes,2,opt,name=roomInfo" json:"roomInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PhzReqCreateDesk) Reset()                    { *m = PhzReqCreateDesk{} }
func (m *PhzReqCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzReqCreateDesk) ProtoMessage()               {}
func (*PhzReqCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{0} }

func (m *PhzReqCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzReqCreateDesk) GetRoomInfo() *PhzBaseCreateOption {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

// 进入房间
type PhzReqEnterDesk struct {
	Header           *ProtoHeader     `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomInfo         *PhzBaseRoomInfo `protobuf:"bytes,2,opt,name=roomInfo" json:"roomInfo,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PhzReqEnterDesk) Reset()                    { *m = PhzReqEnterDesk{} }
func (m *PhzReqEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzReqEnterDesk) ProtoMessage()               {}
func (*PhzReqEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{1} }

func (m *PhzReqEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzReqEnterDesk) GetRoomInfo() *PhzBaseRoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

// 进入、创建房间回复的ack
type PhzAckDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	IsCreateDeskOK   *bool        `protobuf:"varint,2,opt,name=isCreateDeskOK" json:"isCreateDeskOK,omitempty"`
	IsEnterDeskOK    *bool        `protobuf:"varint,3,opt,name=isEnterDeskOK" json:"isEnterDeskOK,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PhzAckDesk) Reset()                    { *m = PhzAckDesk{} }
func (m *PhzAckDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzAckDesk) ProtoMessage()               {}
func (*PhzAckDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{2} }

func (m *PhzAckDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzAckDesk) GetIsCreateDeskOK() bool {
	if m != nil && m.IsCreateDeskOK != nil {
		return *m.IsCreateDeskOK
	}
	return false
}

func (m *PhzAckDesk) GetIsEnterDeskOK() bool {
	if m != nil && m.IsEnterDeskOK != nil {
		return *m.IsEnterDeskOK
	}
	return false
}

type PhzReqGameInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PhzReqGameInfo) Reset()                    { *m = PhzReqGameInfo{} }
func (m *PhzReqGameInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzReqGameInfo) ProtoMessage()               {}
func (*PhzReqGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{3} }

func (m *PhzReqGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type PhzDeskGameInfo struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskInfo         *PhzBaseDeskInfo     `protobuf:"bytes,2,opt,name=deskInfo" json:"deskInfo,omitempty"`
	PlayerInfo       []*PhzBasePlayerInfo `protobuf:"bytes,3,rep,name=playerInfo" json:"playerInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PhzDeskGameInfo) Reset()                    { *m = PhzDeskGameInfo{} }
func (m *PhzDeskGameInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzDeskGameInfo) ProtoMessage()               {}
func (*PhzDeskGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{4} }

func (m *PhzDeskGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzDeskGameInfo) GetDeskInfo() *PhzBaseDeskInfo {
	if m != nil {
		return m.DeskInfo
	}
	return nil
}

func (m *PhzDeskGameInfo) GetPlayerInfo() []*PhzBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

// 房主直接解散房间
type PhzReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PhzReqDissolveDesk) Reset()                    { *m = PhzReqDissolveDesk{} }
func (m *PhzReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzReqDissolveDesk) ProtoMessage()               {}
func (*PhzReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{5} }

func (m *PhzReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzReqDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type PhzAckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32       `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string      `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	UserId           *uint32      `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PhzAckDissolveDesk) Reset()                    { *m = PhzAckDissolveDesk{} }
func (m *PhzAckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*PhzAckDissolveDesk) ProtoMessage()               {}
func (*PhzAckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{6} }

func (m *PhzAckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PhzAckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PhzAckDissolveDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *PhzAckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*PhzReqCreateDesk)(nil), "ddproto.phz_req_createDesk")
	proto.RegisterType((*PhzReqEnterDesk)(nil), "ddproto.phz_req_enterDesk")
	proto.RegisterType((*PhzAckDesk)(nil), "ddproto.phz_ack_desk")
	proto.RegisterType((*PhzReqGameInfo)(nil), "ddproto.phz_req_gameInfo")
	proto.RegisterType((*PhzDeskGameInfo)(nil), "ddproto.phz_desk_gameInfo")
	proto.RegisterType((*PhzReqDissolveDesk)(nil), "ddproto.phz_req_dissolveDesk")
	proto.RegisterType((*PhzAckDissolveDesk)(nil), "ddproto.phz_ack_dissolveDesk")
}

func init() { proto.RegisterFile("phz_desk.proto", fileDescriptor46) }

var fileDescriptor46 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xff, 0xf9, 0x57, 0x6b, 0x9d, 0xda, 0xa2, 0x6b, 0x91, 0x12, 0x44, 0x4a, 0x10, 0xe9,
	0x41, 0x7a, 0xe8, 0xc1, 0x83, 0x78, 0x10, 0x54, 0xb0, 0xf4, 0xd0, 0xb2, 0x67, 0xa1, 0xac, 0xd9,
	0xd1, 0x86, 0x36, 0xd9, 0xb8, 0x1b, 0x95, 0x7a, 0xf0, 0x01, 0x3c, 0xf9, 0x36, 0xbe, 0x9e, 0xec,
	0x26, 0x1b, 0xa3, 0x05, 0xa1, 0xf1, 0x52, 0x3a, 0xbb, 0xbf, 0xf9, 0xbe, 0x99, 0x6f, 0x03, 0xcd,
	0x78, 0xfa, 0x32, 0xe1, 0xa8, 0x66, 0xbd, 0x58, 0x8a, 0x44, 0x90, 0x0d, 0xce, 0xcd, 0x1f, 0x77,
	0xd7, 0x17, 0x61, 0x28, 0xa2, 0x89, 0x3f, 0x0f, 0x30, 0x4a, 0xd2, 0x5b, 0xd7, 0xd0, 0xb7, 0x4c,
	0x61, 0x5a, 0x7b, 0xaf, 0x40, 0xf4, 0x89, 0xc4, 0x87, 0x89, 0x2f, 0x91, 0x25, 0x78, 0x89, 0x6a,
	0x46, 0x8e, 0xa1, 0x3a, 0x45, 0xc6, 0x51, 0xb6, 0x9d, 0x8e, 0xd3, 0xad, 0xf7, 0x5b, 0xbd, 0x4c,
	0xb4, 0x37, 0xd6, 0xbf, 0xd7, 0xe6, 0x8e, 0x66, 0x0c, 0x39, 0x85, 0x9a, 0x14, 0x22, 0x1c, 0x44,
	0x77, 0xa2, 0xfd, 0xdf, 0xf0, 0x07, 0x39, 0x6f, 0xed, 0x32, 0xf5, 0x51, 0x9c, 0x04, 0x22, 0xa2,
	0x39, 0xef, 0x2d, 0x60, 0xc7, 0xfa, 0x63, 0x94, 0xa0, 0x2c, 0x61, 0x7f, 0xb2, 0x64, 0xef, 0x2e,
	0xdb, 0x5b, 0xa2, 0x60, 0xfd, 0xe6, 0xc0, 0x96, 0xbe, 0x67, 0xfe, 0xcc, 0xe4, 0xb7, 0xa2, 0xed,
	0x11, 0x34, 0x03, 0x75, 0x91, 0x67, 0x36, 0x1a, 0x1a, 0xf3, 0x1a, 0xfd, 0x71, 0x4a, 0x0e, 0xa1,
	0x11, 0xa8, 0x2b, 0xbb, 0xdb, 0x68, 0xd8, 0xae, 0x18, 0xec, 0xfb, 0xa1, 0x77, 0x0e, 0xdb, 0x36,
	0x87, 0x7b, 0x16, 0xa2, 0x1e, 0x70, 0xb5, 0x79, 0xbc, 0x0f, 0x27, 0x8d, 0x52, 0xaf, 0x52, 0x52,
	0x43, 0x47, 0xa9, 0xdb, 0x7f, 0x8f, 0xd2, 0x12, 0x34, 0x67, 0xc9, 0x19, 0x40, 0x3c, 0x67, 0x0b,
	0x94, 0xa6, 0xb3, 0xd2, 0xa9, 0x74, 0xeb, 0xfd, 0xfd, 0xe5, 0xce, 0x2f, 0x86, 0x16, 0x78, 0xef,
	0x06, 0x5a, 0x76, 0x77, 0x1e, 0x28, 0x25, 0xe6, 0x4f, 0x65, 0xbe, 0xc2, 0x3d, 0xa8, 0x3e, 0x2a,
	0x94, 0x03, 0x6e, 0x26, 0x6f, 0xd0, 0xac, 0xf2, 0xde, 0x9d, 0x54, 0xde, 0x3c, 0xf3, 0x9f, 0xe4,
	0xcd, 0xba, 0xa9, 0xfc, 0x3a, 0xcd, 0x2a, 0xe2, 0x42, 0x2d, 0x66, 0x4a, 0x3d, 0x0b, 0xc9, 0xcd,
	0xcb, 0x6e, 0xd2, 0xbc, 0x2e, 0x8c, 0xb4, 0x56, 0x1c, 0x69, 0xfc, 0x6f, 0xec, 0x7c, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xc9, 0x2d, 0xe2, 0x11, 0xb7, 0x03, 0x00, 0x00,
}
