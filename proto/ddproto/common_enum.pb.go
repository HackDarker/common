// Code generated by protoc-gen-go.
// source: common_enum.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommonEnumGame int32

const (
	CommonEnumGame_GID_SRC     CommonEnumGame = 0
	CommonEnumGame_GID_TEXAS   CommonEnumGame = 1
	CommonEnumGame_GID_MAHJONG CommonEnumGame = 2
	CommonEnumGame_GID_DDZ     CommonEnumGame = 3
	CommonEnumGame_GID_ZJH     CommonEnumGame = 4
	CommonEnumGame_GID_HALL    CommonEnumGame = 5
)

var CommonEnumGame_name = map[int32]string{
	0: "GID_SRC",
	1: "GID_TEXAS",
	2: "GID_MAHJONG",
	3: "GID_DDZ",
	4: "GID_ZJH",
	5: "GID_HALL",
}
var CommonEnumGame_value = map[string]int32{
	"GID_SRC":     0,
	"GID_TEXAS":   1,
	"GID_MAHJONG": 2,
	"GID_DDZ":     3,
	"GID_ZJH":     4,
	"GID_HALL":    5,
}

func (x CommonEnumGame) Enum() *CommonEnumGame {
	p := new(CommonEnumGame)
	*p = x
	return p
}
func (x CommonEnumGame) String() string {
	return proto.EnumName(CommonEnumGame_name, int32(x))
}
func (x *CommonEnumGame) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommonEnumGame_value, data, "CommonEnumGame")
	if err != nil {
		return err
	}
	*x = CommonEnumGame(value)
	return nil
}
func (CommonEnumGame) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterEnum("ddproto.CommonEnumGame", CommonEnumGame_name, CommonEnumGame_value)
}

var fileDescriptor3 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x4f, 0xcd, 0x2b, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f,
	0x49, 0x01, 0x33, 0xb4, 0xd2, 0xb9, 0x04, 0x90, 0x64, 0xe3, 0xd3, 0x13, 0x73, 0x53, 0x85, 0xb8,
	0xb9, 0xd8, 0xdd, 0x3d, 0x5d, 0xe2, 0x83, 0x83, 0x9c, 0x05, 0x18, 0x84, 0x78, 0xb9, 0x38, 0x41,
	0x9c, 0x10, 0xd7, 0x08, 0xc7, 0x60, 0x01, 0x46, 0x21, 0x7e, 0x2e, 0x6e, 0x10, 0xd7, 0xd7, 0xd1,
	0xc3, 0xcb, 0xdf, 0xcf, 0x5d, 0x80, 0x09, 0xa6, 0xd8, 0xc5, 0x25, 0x4a, 0x80, 0x19, 0xc6, 0x89,
	0xf2, 0xf2, 0x10, 0x60, 0x11, 0xe2, 0xe1, 0xe2, 0x00, 0x71, 0x3c, 0x1c, 0x7d, 0x7c, 0x04, 0x58,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x73, 0x9b, 0x22, 0x19, 0x85, 0x00, 0x00, 0x00,
}
