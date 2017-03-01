// Code generated by protoc-gen-go.
// source: common_client_award.proto
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

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 请求得到在线奖励
type AwardReqOnline struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AwardReqOnline) Reset()                    { *m = AwardReqOnline{} }
func (m *AwardReqOnline) String() string            { return proto.CompactTextString(m) }
func (*AwardReqOnline) ProtoMessage()               {}
func (*AwardReqOnline) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *AwardReqOnline) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type WardAckOnline struct {
	Coin             *int64 `protobuf:"varint,1,opt,name=coin" json:"coin,omitempty"`
	Duration         *int64 `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
	ChangeCoin       *int64 `protobuf:"varint,3,opt,name=changeCoin" json:"changeCoin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *WardAckOnline) Reset()                    { *m = WardAckOnline{} }
func (m *WardAckOnline) String() string            { return proto.CompactTextString(m) }
func (*WardAckOnline) ProtoMessage()               {}
func (*WardAckOnline) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *WardAckOnline) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *WardAckOnline) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *WardAckOnline) GetChangeCoin() int64 {
	if m != nil && m.ChangeCoin != nil {
		return *m.ChangeCoin
	}
	return 0
}

// 请求得到在线奖励
type AwardReqOnlineInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AwardReqOnlineInfo) Reset()                    { *m = AwardReqOnlineInfo{} }
func (m *AwardReqOnlineInfo) String() string            { return proto.CompactTextString(m) }
func (*AwardReqOnlineInfo) ProtoMessage()               {}
func (*AwardReqOnlineInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *AwardReqOnlineInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AwardAckOnlineInfo struct {
	Duration         *int64 `protobuf:"varint,1,opt,name=duration" json:"duration,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AwardAckOnlineInfo) Reset()                    { *m = AwardAckOnlineInfo{} }
func (m *AwardAckOnlineInfo) String() string            { return proto.CompactTextString(m) }
func (*AwardAckOnlineInfo) ProtoMessage()               {}
func (*AwardAckOnlineInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *AwardAckOnlineInfo) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func init() {
	proto.RegisterType((*AwardReqOnline)(nil), "ddproto.award_req_online")
	proto.RegisterType((*WardAckOnline)(nil), "ddproto.ward_ack_online")
	proto.RegisterType((*AwardReqOnlineInfo)(nil), "ddproto.award_req_onlineInfo")
	proto.RegisterType((*AwardAckOnlineInfo)(nil), "ddproto.award_ack_onlineInfo")
}

var fileDescriptor1 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x89, 0x4f, 0x2c, 0x4f, 0x2c, 0x4a, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0xa4, 0x84, 0x51, 0xd4, 0x40,
	0x64, 0x95, 0x2c, 0xb8, 0x04, 0xc0, 0x8a, 0xe3, 0x8b, 0x52, 0x0b, 0xe3, 0xf3, 0xf3, 0x72, 0x32,
	0xf3, 0x52, 0x85, 0x54, 0xb8, 0xd8, 0x32, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0x44, 0xf4, 0xa0, 0x46, 0xe8, 0x05, 0x80, 0x48, 0x0f, 0xb0, 0x9c, 0x92, 0x2b,
	0x17, 0x3f, 0x58, 0x63, 0x62, 0x72, 0x36, 0x4c, 0x23, 0x0f, 0x17, 0x4b, 0x72, 0x7e, 0x66, 0x1e,
	0x58, 0x1b, 0xb3, 0x90, 0x00, 0x17, 0x47, 0x4a, 0x69, 0x51, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x04,
	0x13, 0x58, 0x44, 0x88, 0x8b, 0x2b, 0x39, 0x23, 0x31, 0x2f, 0x3d, 0xd5, 0x19, 0xa4, 0x8a, 0x19,
	0x24, 0xa6, 0x64, 0xc3, 0x25, 0x82, 0xee, 0x00, 0xcf, 0xbc, 0xb4, 0x7c, 0x22, 0x1d, 0xa1, 0x01,
	0xd3, 0x8d, 0x70, 0x05, 0x58, 0x37, 0xb2, 0xdd, 0x60, 0xd7, 0x04, 0x30, 0x00, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x21, 0xc9, 0x6e, 0x68, 0x24, 0x01, 0x00, 0x00,
}
