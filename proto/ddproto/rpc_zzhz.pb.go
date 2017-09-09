// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_zzhz.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mjproto1 "casino_common/proto/ddproto/mjproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ServerInfo from public import mjproto/mahjong_hall.proto
type ServerInfo mjproto1.ServerInfo

func (m *ServerInfo) Reset()           { (*mjproto1.ServerInfo)(m).Reset() }
func (m *ServerInfo) String() string   { return (*mjproto1.ServerInfo)(m).String() }
func (*ServerInfo) ProtoMessage()      {}
func (m *ServerInfo) GetIp() string    { return (*mjproto1.ServerInfo)(m).GetIp() }
func (m *ServerInfo) GetPort() int32   { return (*mjproto1.ServerInfo)(m).GetPort() }
func (m *ServerInfo) GetStatus() int32 { return (*mjproto1.ServerInfo)(m).GetStatus() }

// game_QuickConn from public import mjproto/mahjong_hall.proto
type Game_QuickConn mjproto1.Game_QuickConn

func (m *Game_QuickConn) Reset()                { (*mjproto1.Game_QuickConn)(m).Reset() }
func (m *Game_QuickConn) String() string        { return (*mjproto1.Game_QuickConn)(m).String() }
func (*Game_QuickConn) ProtoMessage()           {}
func (m *Game_QuickConn) GetChannelId() string  { return (*mjproto1.Game_QuickConn)(m).GetChannelId() }
func (m *Game_QuickConn) GetGameId() int32      { return (*mjproto1.Game_QuickConn)(m).GetGameId() }
func (m *Game_QuickConn) GetCurrVersion() int32 { return (*mjproto1.Game_QuickConn)(m).GetCurrVersion() }
func (m *Game_QuickConn) GetLanguageId() int32  { return (*mjproto1.Game_QuickConn)(m).GetLanguageId() }
func (m *Game_QuickConn) GetUserId() uint32     { return (*mjproto1.Game_QuickConn)(m).GetUserId() }

// game_AckQuickConn from public import mjproto/mahjong_hall.proto
type Game_AckQuickConn mjproto1.Game_AckQuickConn

func (m *Game_AckQuickConn) Reset()         { (*mjproto1.Game_AckQuickConn)(m).Reset() }
func (m *Game_AckQuickConn) String() string { return (*mjproto1.Game_AckQuickConn)(m).String() }
func (*Game_AckQuickConn) ProtoMessage()    {}
func (m *Game_AckQuickConn) GetGameServer() *ServerInfo {
	return (*ServerInfo)((*mjproto1.Game_AckQuickConn)(m).GetGameServer())
}
func (m *Game_AckQuickConn) GetReleaseTag() int32 {
	return (*mjproto1.Game_AckQuickConn)(m).GetReleaseTag()
}
func (m *Game_AckQuickConn) GetCurrVersion() int32 {
	return (*mjproto1.Game_AckQuickConn)(m).GetCurrVersion()
}
func (m *Game_AckQuickConn) GetIsUpdate() int32 { return (*mjproto1.Game_AckQuickConn)(m).GetIsUpdate() }
func (m *Game_AckQuickConn) GetDownloadUrl() string {
	return (*mjproto1.Game_AckQuickConn)(m).GetDownloadUrl()
}
func (m *Game_AckQuickConn) GetVersionInfo() string {
	return (*mjproto1.Game_AckQuickConn)(m).GetVersionInfo()
}
func (m *Game_AckQuickConn) GetIsMaintain() int32 {
	return (*mjproto1.Game_AckQuickConn)(m).GetIsMaintain()
}
func (m *Game_AckQuickConn) GetMaintainMsg() string {
	return (*mjproto1.Game_AckQuickConn)(m).GetMaintainMsg()
}

// game_Login from public import mjproto/mahjong_hall.proto
type Game_Login mjproto1.Game_Login

func (m *Game_Login) Reset()                 { (*mjproto1.Game_Login)(m).Reset() }
func (m *Game_Login) String() string         { return (*mjproto1.Game_Login)(m).String() }
func (*Game_Login) ProtoMessage()            {}
func (m *Game_Login) GetUserId() uint32      { return (*mjproto1.Game_Login)(m).GetUserId() }
func (m *Game_Login) GetProtoVersion() int32 { return (*mjproto1.Game_Login)(m).GetProtoVersion() }

// game_AckLogin from public import mjproto/mahjong_hall.proto
type Game_AckLogin mjproto1.Game_AckLogin

func (m *Game_AckLogin) Reset()              { (*mjproto1.Game_AckLogin)(m).Reset() }
func (m *Game_AckLogin) String() string      { return (*mjproto1.Game_AckLogin)(m).String() }
func (*Game_AckLogin) ProtoMessage()         {}
func (m *Game_AckLogin) GetUserId() uint32   { return (*mjproto1.Game_AckLogin)(m).GetUserId() }
func (m *Game_AckLogin) GetNickName() string { return (*mjproto1.Game_AckLogin)(m).GetNickName() }
func (m *Game_AckLogin) GetRoomPassword() string {
	return (*mjproto1.Game_AckLogin)(m).GetRoomPassword()
}
func (m *Game_AckLogin) GetCostCreateRoom() int64 {
	return (*mjproto1.Game_AckLogin)(m).GetCostCreateRoom()
}
func (m *Game_AckLogin) GetCostRebuy() int64   { return (*mjproto1.Game_AckLogin)(m).GetCostRebuy() }
func (m *Game_AckLogin) GetChampionship() bool { return (*mjproto1.Game_AckLogin)(m).GetChampionship() }
func (m *Game_AckLogin) GetChip() int64        { return (*mjproto1.Game_AckLogin)(m).GetChip() }
func (m *Game_AckLogin) GetMailCount() int32   { return (*mjproto1.Game_AckLogin)(m).GetMailCount() }
func (m *Game_AckLogin) GetNotice() string     { return (*mjproto1.Game_AckLogin)(m).GetNotice() }
func (m *Game_AckLogin) GetGameStatus() int32  { return (*mjproto1.Game_AckLogin)(m).GetGameStatus() }

// game_Notice from public import mjproto/mahjong_hall.proto
type Game_Notice mjproto1.Game_Notice

func (m *Game_Notice) Reset()               { (*mjproto1.Game_Notice)(m).Reset() }
func (m *Game_Notice) String() string       { return (*mjproto1.Game_Notice)(m).String() }
func (*Game_Notice) ProtoMessage()          {}
func (m *Game_Notice) GetNoticeType() int32 { return (*mjproto1.Game_Notice)(m).GetNoticeType() }
func (m *Game_Notice) GetChannelId() string { return (*mjproto1.Game_Notice)(m).GetChannelId() }

// game_AckNotice from public import mjproto/mahjong_hall.proto
type Game_AckNotice mjproto1.Game_AckNotice

func (m *Game_AckNotice) Reset()               { (*mjproto1.Game_AckNotice)(m).Reset() }
func (m *Game_AckNotice) String() string       { return (*mjproto1.Game_AckNotice)(m).String() }
func (*Game_AckNotice) ProtoMessage()          {}
func (m *Game_AckNotice) GetNoticeType() int32 { return (*mjproto1.Game_AckNotice)(m).GetNoticeType() }
func (m *Game_AckNotice) GetNoticeTitle() string {
	return (*mjproto1.Game_AckNotice)(m).GetNoticeTitle()
}
func (m *Game_AckNotice) GetNoticeContent() string {
	return (*mjproto1.Game_AckNotice)(m).GetNoticeContent()
}
func (m *Game_AckNotice) GetNoticeMemo() string { return (*mjproto1.Game_AckNotice)(m).GetNoticeMemo() }
func (m *Game_AckNotice) GetId() int32          { return (*mjproto1.Game_AckNotice)(m).GetId() }
func (m *Game_AckNotice) GetFileds() []string   { return (*mjproto1.Game_AckNotice)(m).GetFileds() }

// game_GameRecord from public import mjproto/mahjong_hall.proto
type Game_GameRecord mjproto1.Game_GameRecord

func (m *Game_GameRecord) Reset()            { (*mjproto1.Game_GameRecord)(m).Reset() }
func (m *Game_GameRecord) String() string    { return (*mjproto1.Game_GameRecord)(m).String() }
func (*Game_GameRecord) ProtoMessage()       {}
func (m *Game_GameRecord) GetId() int32      { return (*mjproto1.Game_GameRecord)(m).GetId() }
func (m *Game_GameRecord) GetGameId() int32  { return (*mjproto1.Game_GameRecord)(m).GetGameId() }
func (m *Game_GameRecord) GetUserId() uint32 { return (*mjproto1.Game_GameRecord)(m).GetUserId() }

// BeanUserRecord from public import mjproto/mahjong_hall.proto
type BeanUserRecord mjproto1.BeanUserRecord

func (m *BeanUserRecord) Reset()              { (*mjproto1.BeanUserRecord)(m).Reset() }
func (m *BeanUserRecord) String() string      { return (*mjproto1.BeanUserRecord)(m).String() }
func (*BeanUserRecord) ProtoMessage()         {}
func (m *BeanUserRecord) GetUserId() uint32   { return (*mjproto1.BeanUserRecord)(m).GetUserId() }
func (m *BeanUserRecord) GetNickName() string { return (*mjproto1.BeanUserRecord)(m).GetNickName() }
func (m *BeanUserRecord) GetWinAmount() int64 { return (*mjproto1.BeanUserRecord)(m).GetWinAmount() }

// BeanGameRecord from public import mjproto/mahjong_hall.proto
type BeanGameRecord mjproto1.BeanGameRecord

func (m *BeanGameRecord) Reset()               { (*mjproto1.BeanGameRecord)(m).Reset() }
func (m *BeanGameRecord) String() string       { return (*mjproto1.BeanGameRecord)(m).String() }
func (*BeanGameRecord) ProtoMessage()          {}
func (m *BeanGameRecord) GetId() int32         { return (*mjproto1.BeanGameRecord)(m).GetId() }
func (m *BeanGameRecord) GetDeskId() int32     { return (*mjproto1.BeanGameRecord)(m).GetDeskId() }
func (m *BeanGameRecord) GetBeginTime() string { return (*mjproto1.BeanGameRecord)(m).GetBeginTime() }
func (m *BeanGameRecord) GetUsers() []*BeanUserRecord {
	o := (*mjproto1.BeanGameRecord)(m).GetUsers()
	if o == nil {
		return nil
	}
	s := make([]*BeanUserRecord, len(o))
	for i, x := range o {
		s[i] = (*BeanUserRecord)(x)
	}
	return s
}

// game_AckGameRecord from public import mjproto/mahjong_hall.proto
type Game_AckGameRecord mjproto1.Game_AckGameRecord

func (m *Game_AckGameRecord) Reset()            { (*mjproto1.Game_AckGameRecord)(m).Reset() }
func (m *Game_AckGameRecord) String() string    { return (*mjproto1.Game_AckGameRecord)(m).String() }
func (*Game_AckGameRecord) ProtoMessage()       {}
func (m *Game_AckGameRecord) GetUserId() uint32 { return (*mjproto1.Game_AckGameRecord)(m).GetUserId() }
func (m *Game_AckGameRecord) GetGameId() int32  { return (*mjproto1.Game_AckGameRecord)(m).GetGameId() }
func (m *Game_AckGameRecord) GetRecords() []*BeanGameRecord {
	o := (*mjproto1.Game_AckGameRecord)(m).GetRecords()
	if o == nil {
		return nil
	}
	s := make([]*BeanGameRecord, len(o))
	for i, x := range o {
		s[i] = (*BeanGameRecord)(x)
	}
	return s
}

// game_Feedback from public import mjproto/mahjong_hall.proto
type Game_Feedback mjproto1.Game_Feedback

func (m *Game_Feedback) Reset()             { (*mjproto1.Game_Feedback)(m).Reset() }
func (m *Game_Feedback) String() string     { return (*mjproto1.Game_Feedback)(m).String() }
func (*Game_Feedback) ProtoMessage()        {}
func (m *Game_Feedback) GetMessage() string { return (*mjproto1.Game_Feedback)(m).GetMessage() }

// game_CreateRoom from public import mjproto/mahjong_hall.proto
type Game_CreateRoom mjproto1.Game_CreateRoom

func (m *Game_CreateRoom) Reset()            { (*mjproto1.Game_CreateRoom)(m).Reset() }
func (m *Game_CreateRoom) String() string    { return (*mjproto1.Game_CreateRoom)(m).String() }
func (*Game_CreateRoom) ProtoMessage()       {}
func (m *Game_CreateRoom) GetIsDaiKai() bool { return (*mjproto1.Game_CreateRoom)(m).GetIsDaiKai() }

// game_AckCreateRoom from public import mjproto/mahjong_hall.proto
type Game_AckCreateRoom mjproto1.Game_AckCreateRoom

func (m *Game_AckCreateRoom) Reset()           { (*mjproto1.Game_AckCreateRoom)(m).Reset() }
func (m *Game_AckCreateRoom) String() string   { return (*mjproto1.Game_AckCreateRoom)(m).String() }
func (*Game_AckCreateRoom) ProtoMessage()      {}
func (m *Game_AckCreateRoom) GetDeskId() int32 { return (*mjproto1.Game_AckCreateRoom)(m).GetDeskId() }
func (m *Game_AckCreateRoom) GetPassword() string {
	return (*mjproto1.Game_AckCreateRoom)(m).GetPassword()
}
func (m *Game_AckCreateRoom) GetUserBalance() int64 {
	return (*mjproto1.Game_AckCreateRoom)(m).GetUserBalance()
}
func (m *Game_AckCreateRoom) GetCreateFee() int64 {
	return (*mjproto1.Game_AckCreateRoom)(m).GetCreateFee()
}

// game_EnterRoom from public import mjproto/mahjong_hall.proto
type Game_EnterRoom mjproto1.Game_EnterRoom

func (m *Game_EnterRoom) Reset()              { (*mjproto1.Game_EnterRoom)(m).Reset() }
func (m *Game_EnterRoom) String() string      { return (*mjproto1.Game_EnterRoom)(m).String() }
func (*Game_EnterRoom) ProtoMessage()         {}
func (m *Game_EnterRoom) GetMatchId() int32   { return (*mjproto1.Game_EnterRoom)(m).GetMatchId() }
func (m *Game_EnterRoom) GetTableId() int32   { return (*mjproto1.Game_EnterRoom)(m).GetTableId() }
func (m *Game_EnterRoom) GetPassWord() string { return (*mjproto1.Game_EnterRoom)(m).GetPassWord() }
func (m *Game_EnterRoom) GetUserId() uint32   { return (*mjproto1.Game_EnterRoom)(m).GetUserId() }
func (m *Game_EnterRoom) GetRoomType() int32  { return (*mjproto1.Game_EnterRoom)(m).GetRoomType() }
func (m *Game_EnterRoom) GetRoomLevel() int32 { return (*mjproto1.Game_EnterRoom)(m).GetRoomLevel() }
func (m *Game_EnterRoom) GetEnterType() int32 { return (*mjproto1.Game_EnterRoom)(m).GetEnterType() }

// game_AckEnterRoom from public import mjproto/mahjong_hall.proto
type Game_AckEnterRoom mjproto1.Game_AckEnterRoom

func (m *Game_AckEnterRoom) Reset()         { (*mjproto1.Game_AckEnterRoom)(m).Reset() }
func (m *Game_AckEnterRoom) String() string { return (*mjproto1.Game_AckEnterRoom)(m).String() }
func (*Game_AckEnterRoom) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ZzHz service

type ZzHzClient interface {
	// 创建房间
	CreateRoom(ctx context.Context, in *mjproto1.Game_CreateRoom, opts ...grpc.CallOption) (*mjproto1.Game_AckCreateRoom, error)
}

type zzHzClient struct {
	cc *grpc.ClientConn
}

func NewZzHzClient(cc *grpc.ClientConn) ZzHzClient {
	return &zzHzClient{cc}
}

func (c *zzHzClient) CreateRoom(ctx context.Context, in *mjproto1.Game_CreateRoom, opts ...grpc.CallOption) (*mjproto1.Game_AckCreateRoom, error) {
	out := new(mjproto1.Game_AckCreateRoom)
	err := grpc.Invoke(ctx, "/ddproto.ZzHz/CreateRoom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ZzHz service

type ZzHzServer interface {
	// 创建房间
	CreateRoom(context.Context, *mjproto1.Game_CreateRoom) (*mjproto1.Game_AckCreateRoom, error)
}

func RegisterZzHzServer(s *grpc.Server, srv ZzHzServer) {
	s.RegisterService(&_ZzHz_serviceDesc, srv)
}

func _ZzHz_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(mjproto1.Game_CreateRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZzHzServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddproto.ZzHz/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZzHzServer).CreateRoom(ctx, req.(*mjproto1.Game_CreateRoom))
	}
	return interceptor(ctx, in, info, handler)
}

var _ZzHz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ddproto.ZzHz",
	HandlerType: (*ZzHzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _ZzHz_CreateRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_zzhz.proto",
}

func init() { proto.RegisterFile("rpc_zzhz.proto", fileDescriptor48) }

var fileDescriptor48 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2a, 0x48, 0x8e,
	0xaf, 0xaa, 0xca, 0xa8, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33,
	0xa4, 0xa4, 0x72, 0xb3, 0xc0, 0x0c, 0xfd, 0xdc, 0xc4, 0x8c, 0xac, 0xfc, 0xbc, 0xf4, 0xf8, 0x8c,
	0xc4, 0x9c, 0x1c, 0x88, 0x22, 0x23, 0x5f, 0x2e, 0x96, 0xa8, 0x2a, 0x8f, 0x2a, 0x21, 0x57, 0x2e,
	0x2e, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0xd4, 0xa0, 0xfc, 0xfc, 0x5c, 0x21, 0x09, 0x3d, 0xa8, 0x16,
	0xbd, 0xf4, 0xc4, 0xdc, 0xd4, 0x78, 0x84, 0x8c, 0x94, 0x34, 0xaa, 0x8c, 0x63, 0x72, 0x36, 0x42,
	0x52, 0x89, 0xc1, 0x89, 0xc9, 0x83, 0x39, 0x80, 0x21, 0x89, 0x0d, 0xac, 0xc2, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x79, 0x60, 0x60, 0x68, 0x92, 0x00, 0x00, 0x00,
}
