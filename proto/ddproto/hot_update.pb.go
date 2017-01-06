// Code generated by protoc-gen-go.
// source: hot_update.proto
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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

type VersionInfo struct {
	FileId           *int32 `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32 `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *VersionInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *VersionInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

type AssetInfo struct {
	FileId           *int32          `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32          `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	FilePath         *string         `protobuf:"bytes,3,opt,name=filePath" json:"filePath,omitempty"`
	FileSize         *int32          `protobuf:"varint,4,opt,name=fileSize" json:"fileSize,omitempty"`
	Md5              *string         `protobuf:"bytes,5,opt,name=md5" json:"md5,omitempty"`
	IsCompress       *bool           `protobuf:"varint,6,opt,name=isCompress" json:"isCompress,omitempty"`
	IsCode           *bool           `protobuf:"varint,7,opt,name=isCode" json:"isCode,omitempty"`
	GameId           *CommonEnumGame `protobuf:"varint,8,opt,name=gameId,enum=ddproto.CommonEnumGame" json:"gameId,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *AssetInfo) Reset()                    { *m = AssetInfo{} }
func (m *AssetInfo) String() string            { return proto.CompactTextString(m) }
func (*AssetInfo) ProtoMessage()               {}
func (*AssetInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

func (m *AssetInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *AssetInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

func (m *AssetInfo) GetFilePath() string {
	if m != nil && m.FilePath != nil {
		return *m.FilePath
	}
	return ""
}

func (m *AssetInfo) GetFileSize() int32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *AssetInfo) GetMd5() string {
	if m != nil && m.Md5 != nil {
		return *m.Md5
	}
	return ""
}

func (m *AssetInfo) GetIsCompress() bool {
	if m != nil && m.IsCompress != nil {
		return *m.IsCompress
	}
	return false
}

func (m *AssetInfo) GetIsCode() bool {
	if m != nil && m.IsCode != nil {
		return *m.IsCode
	}
	return false
}

func (m *AssetInfo) GetGameId() CommonEnumGame {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return CommonEnumGame_GID_SRC
}

// 获取资源文件的最新版本信息
type HotupdateReqVersionInfo struct {
	Header              *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId         *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	DownloadGameId      []int32      `protobuf:"varint,3,rep,name=downloadGameId" json:"downloadGameId,omitempty"`
	ClientAssetsVersion *int32       `protobuf:"varint,4,opt,name=clientAssetsVersion" json:"clientAssetsVersion,omitempty"`
	ClientResolution    *int32       `protobuf:"varint,5,opt,name=clientResolution" json:"clientResolution,omitempty"`
	XXX_unrecognized    []byte       `json:"-"`
}

func (m *HotupdateReqVersionInfo) Reset()                    { *m = HotupdateReqVersionInfo{} }
func (m *HotupdateReqVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqVersionInfo) ProtoMessage()               {}
func (*HotupdateReqVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func (m *HotupdateReqVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetDownloadGameId() []int32 {
	if m != nil {
		return m.DownloadGameId
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAssetsVersion() int32 {
	if m != nil && m.ClientAssetsVersion != nil {
		return *m.ClientAssetsVersion
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

type HotupdateAckVersionInfo struct {
	Header               *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Version              []*VersionInfo `protobuf:"bytes,2,rep,name=version" json:"version,omitempty"`
	LastestAssetsVersion *int32         `protobuf:"varint,3,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	ForceUpdate          *bool          `protobuf:"varint,4,opt,name=forceUpdate" json:"forceUpdate,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *HotupdateAckVersionInfo) Reset()                    { *m = HotupdateAckVersionInfo{} }
func (m *HotupdateAckVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckVersionInfo) ProtoMessage()               {}
func (*HotupdateAckVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

func (m *HotupdateAckVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetVersion() []*VersionInfo {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

func (m *HotupdateAckVersionInfo) GetForceUpdate() bool {
	if m != nil && m.ForceUpdate != nil {
		return *m.ForceUpdate
	}
	return false
}

// 获取资源文件的详细信息
type HotupdateReqAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	FileIds          []int32      `protobuf:"varint,3,rep,name=fileIds" json:"fileIds,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqAssetsInfo) Reset()                    { *m = HotupdateReqAssetsInfo{} }
func (m *HotupdateReqAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{4} }

func (m *HotupdateReqAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqAssetsInfo) GetFileIds() []int32 {
	if m != nil {
		return m.FileIds
	}
	return nil
}

type HotupdateAckAssetsInfo struct {
	Header               *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost            *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets               []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	LastestAssetsVersion *int32       `protobuf:"varint,4,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	XXX_unrecognized     []byte       `json:"-"`
}

func (m *HotupdateAckAssetsInfo) Reset()                    { *m = HotupdateAckAssetsInfo{} }
func (m *HotupdateAckAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{5} }

func (m *HotupdateAckAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

// 获取某个独立游戏的下载信息
type HotupdateReqGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqGameAssetsInfo) Reset()                    { *m = HotupdateReqGameAssetsInfo{} }
func (m *HotupdateReqGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{6} }

func (m *HotupdateReqGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqGameAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqGameAssetsInfo) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

type HotupdateAckGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost        *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets           []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	ClientResolution *int32       `protobuf:"varint,4,opt,name=clientResolution" json:"clientResolution,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateAckGameAssetsInfo) Reset()                    { *m = HotupdateAckGameAssetsInfo{} }
func (m *HotupdateAckGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{7} }

func (m *HotupdateAckGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckGameAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckGameAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *HotupdateAckGameAssetsInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

func init() {
	proto.RegisterType((*VersionInfo)(nil), "ddproto.VersionInfo")
	proto.RegisterType((*AssetInfo)(nil), "ddproto.AssetInfo")
	proto.RegisterType((*HotupdateReqVersionInfo)(nil), "ddproto.hotupdate_req_versionInfo")
	proto.RegisterType((*HotupdateAckVersionInfo)(nil), "ddproto.hotupdate_ack_versionInfo")
	proto.RegisterType((*HotupdateReqAssetsInfo)(nil), "ddproto.hotupdate_req_assetsInfo")
	proto.RegisterType((*HotupdateAckAssetsInfo)(nil), "ddproto.hotupdate_ack_assetsInfo")
	proto.RegisterType((*HotupdateReqGameAssetsInfo)(nil), "ddproto.hotupdate_req_gameAssetsInfo")
	proto.RegisterType((*HotupdateAckGameAssetsInfo)(nil), "ddproto.hotupdate_ack_gameAssetsInfo")
}

var fileDescriptor16 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9d, 0xa6, 0x4d, 0xdb, 0x57, 0x58, 0xea, 0xec, 0xb2, 0xcc, 0x96, 0x45, 0x42, 0x0e,
	0x12, 0x16, 0x29, 0x5a, 0xf0, 0x2c, 0xd5, 0x83, 0xed, 0xad, 0x8c, 0xb8, 0xd7, 0x10, 0x3a, 0x53,
	0x1b, 0x4c, 0x32, 0x31, 0x33, 0x55, 0xf0, 0x2c, 0x7e, 0x30, 0x0f, 0x5e, 0xfd, 0x2c, 0xfa, 0x09,
	0x64, 0x26, 0x93, 0x98, 0xac, 0xd9, 0xcb, 0xee, 0xf6, 0x12, 0xe6, 0xfd, 0xdf, 0x7b, 0xc3, 0x7f,
	0x7e, 0x2f, 0x0f, 0xa6, 0x7b, 0xa1, 0xc2, 0x43, 0xce, 0x22, 0xc5, 0xe7, 0x79, 0x21, 0x94, 0xc0,
	0x43, 0xc6, 0xcc, 0x61, 0x76, 0xba, 0x15, 0x69, 0x2a, 0xb2, 0x70, 0x9b, 0xc4, 0x3c, 0x53, 0x65,
	0x76, 0xf6, 0xd8, 0x8a, 0x3c, 0x3b, 0xa4, 0xa5, 0xe4, 0xbf, 0x82, 0xc9, 0x35, 0x2f, 0x64, 0x2c,
	0xb2, 0x75, 0xb6, 0x13, 0xf8, 0x1c, 0xdc, 0x5d, 0x9c, 0xf0, 0x35, 0x23, 0xc8, 0x43, 0xc1, 0x80,
	0xda, 0x08, 0x13, 0x18, 0xea, 0xd3, 0x35, 0x2f, 0x48, 0xcf, 0x24, 0xaa, 0xd0, 0xff, 0x83, 0x60,
	0xbc, 0x94, 0x92, 0xab, 0xbb, 0xf5, 0xe3, 0x19, 0x8c, 0xf4, 0x71, 0x13, 0xa9, 0x3d, 0x71, 0x3c,
	0x14, 0x8c, 0x69, 0x1d, 0x57, 0xb9, 0x77, 0xf1, 0x57, 0x4e, 0xfa, 0xa6, 0xad, 0x8e, 0xf1, 0x14,
	0x9c, 0x94, 0xbd, 0x24, 0x03, 0xd3, 0xa2, 0x8f, 0xf8, 0x09, 0x40, 0x2c, 0xdf, 0x88, 0x34, 0x2f,
	0xb8, 0x94, 0xc4, 0xf5, 0x50, 0x30, 0xa2, 0x0d, 0x45, 0x7b, 0xd3, 0x11, 0xe3, 0x64, 0x68, 0x72,
	0x36, 0xc2, 0x2f, 0xc0, 0xfd, 0x10, 0xa5, 0xda, 0xf3, 0xc8, 0x43, 0xc1, 0xc9, 0xe2, 0x62, 0x6e,
	0x21, 0xce, 0x1b, 0xb8, 0x42, 0x5d, 0x42, 0x6d, 0xa1, 0xff, 0x1b, 0xc1, 0xc5, 0x5e, 0xa8, 0x12,
	0x7d, 0x58, 0xf0, 0x4f, 0xe1, 0xe7, 0x06, 0xc4, 0x67, 0xe0, 0xee, 0x79, 0xc4, 0x78, 0x61, 0x20,
	0x4c, 0x16, 0x67, 0xf5, 0x85, 0x1b, 0xfd, 0x5d, 0x99, 0x1c, 0xb5, 0x35, 0xd8, 0x83, 0x49, 0x39,
	0xa4, 0x65, 0x9e, 0xaf, 0x99, 0xc5, 0xd3, 0x94, 0xf0, 0x53, 0x38, 0x61, 0xe2, 0x4b, 0x96, 0x88,
	0x88, 0xbd, 0x2d, 0x8d, 0x3a, 0x9e, 0x13, 0x0c, 0xe8, 0x0d, 0x15, 0x3f, 0x87, 0x53, 0xdb, 0xa6,
	0xe7, 0x21, 0xed, 0x5c, 0x2d, 0xb9, 0xae, 0x14, 0xbe, 0x82, 0x69, 0x29, 0x53, 0x2e, 0x45, 0x72,
	0x50, 0xba, 0x7c, 0x60, 0xca, 0xff, 0xd3, 0xfd, 0x5f, 0xad, 0x37, 0x47, 0xdb, 0x8f, 0xf7, 0x78,
	0xf3, 0x1c, 0x86, 0xb6, 0x99, 0xf4, 0x3c, 0xa7, 0x55, 0xde, 0xf8, 0x1b, 0x69, 0x55, 0x84, 0x17,
	0x70, 0x96, 0x44, 0x52, 0x71, 0x79, 0xe3, 0x69, 0x8e, 0xf1, 0xda, 0x99, 0xd3, 0x5c, 0x77, 0xa2,
	0xd8, 0xf2, 0xf7, 0xc6, 0xb0, 0xa1, 0x30, 0xa2, 0x4d, 0xc9, 0xff, 0x86, 0x80, 0xb4, 0xa7, 0x18,
	0x99, 0x1b, 0x8e, 0x32, 0x44, 0xbb, 0x01, 0x6b, 0x26, 0xed, 0xf4, 0xaa, 0xd0, 0xff, 0xd9, 0xb2,
	0xa1, 0xc1, 0xde, 0xd9, 0xc6, 0x25, 0x8c, 0x4d, 0xef, 0x4a, 0x48, 0x65, 0x4c, 0x8c, 0xe9, 0x3f,
	0x01, 0x5f, 0x81, 0x5b, 0xde, 0x6c, 0x1c, 0x4c, 0x16, 0xb8, 0xbe, 0xab, 0x5e, 0x60, 0x6a, 0x2b,
	0x6e, 0x25, 0xde, 0xbf, 0x9d, 0xb8, 0xff, 0x1d, 0xc1, 0x65, 0x9b, 0xa7, 0x5e, 0x97, 0xe5, 0xf1,
	0x98, 0x9e, 0xd7, 0x9b, 0x5b, 0xfe, 0x08, 0xd5, 0x7a, 0xfe, 0x68, 0x19, 0xd1, 0x44, 0xef, 0x65,
	0xe4, 0xe1, 0xa8, 0x76, 0xed, 0x5b, 0xbf, 0x7b, 0xdf, 0x5e, 0xf7, 0x56, 0xce, 0xe6, 0xd1, 0x06,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x6d, 0xaa, 0x5b, 0xe5, 0x05, 0x00, 0x00,
}
