// Code generated by protoc-gen-go.
// source: common_server_redis.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 游戏服务器所处状态
type RedisGameServerStatus struct {
	IsOnMaintain     *bool  `protobuf:"varint,1,opt,name=isOnMaintain" json:"isOnMaintain,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RedisGameServerStatus) Reset()                    { *m = RedisGameServerStatus{} }
func (m *RedisGameServerStatus) String() string            { return proto.CompactTextString(m) }
func (*RedisGameServerStatus) ProtoMessage()               {}
func (*RedisGameServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *RedisGameServerStatus) GetIsOnMaintain() bool {
	if m != nil && m.IsOnMaintain != nil {
		return *m.IsOnMaintain
	}
	return false
}

func init() {
	proto.RegisterType((*RedisGameServerStatus)(nil), "ddproto.redis_game_server_status")
}

var fileDescriptor14 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x8a, 0x2f, 0x4a, 0x4d, 0xc9, 0x2c, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0x94, 0x0c, 0xb8, 0x24, 0xc0,
	0xe2, 0xf1, 0xe9, 0x89, 0xb9, 0xa9, 0x30, 0x95, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x42, 0x22,
	0x5c, 0x3c, 0x99, 0xc5, 0xfe, 0x79, 0xbe, 0x89, 0x99, 0x79, 0x25, 0x40, 0x2c, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x01, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xf0, 0x49, 0xc6, 0x56, 0x00, 0x00, 0x00,
}
