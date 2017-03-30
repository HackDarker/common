package api

import (
	"github.com/golang/protobuf/proto"
	"casino_common/common/userService"
	"github.com/name5566/leaf/gate"
	"casino_common/utils/agentUtils"
)

type MJUser interface {
	//所有基本信息的优化
	GetUserId() uint32
	GetCoin() int64
	GetSex() int32
	GetHeadUrl() string
	GetOpenId() string
	GetNickname() string
	GetUserSkeleton() interface{}
	GetGameData() interface{}
	GetDesk() MJDesk //得到桌子

	//基本功能
	SendOverTurn(p proto.Message) error //发送overTurn
	WriteMsg(p proto.Message) error     //发送信息
	DoReady() error                     //准备

	//业务
	AddBillBean(interface{}) error //增加账单
}

//核心User
type MJUserCore struct {
	UserId uint32
	uRedis userService.U_REDIS //这里可以使用redis的方法
	Desk   MJDesk
	coin   int64
	gate.Agent
}

//New一个CoreUser
func NewMJUserCore(userId uint32, a gate.Agent) *MJUserCore {
	return &MJUserCore{
		UserId: userId,
		uRedis: userService.U_REDIS(userId),
		Agent:  a,
	}
}

func (u *MJUserCore) GetUserId() uint32 {
	return u.UserId
}

func (u *MJUserCore) GetSex() int32 {
	return u.uRedis.GetSex()
}

func (u *MJUserCore) GetHeadUrl() string {
	return u.uRedis.GetHeadUrl()
}

func (u *MJUserCore) GetOpenId() string {
	return u.uRedis.GetOpenId()
}

func (u *MJUserCore) GetNickname() string {
	return u.uRedis.GetNickName()
}

func (u *MJUserCore) GetIp() string {
	return agentUtils.GetIP(u.Agent)
}

func (u *MJUserCore) GetCoin() int64 {
	return u.coin
}
