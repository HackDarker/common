package handlers

import (
	"casino_common/proto/funcsInit"
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
	"casino_common/common/log"
	"casino_common/common/consts"
	"casino_common/common/service/loginService"
)

func HandlerGame_Login(args []interface{}) {
	m := args[0].(*ddproto.CommonReqGameLogin)
	a := args[1].(gate.Agent)
	log.T("请求handlerGame_Login  m[%v]", m)


	//调用登录方法
	user, err := loginService.DoLogin(m.GetWxInfo(), m.GetUserId())
	if err != nil || user == nil {
		log.E("登录出现出错..")
		ack := commonNewPorot.NewCommonAckGameLogin()
		*ack.Header.Code = consts.ACK_RESULT_ERROR
		a.WriteMsg(ack)
	} else {
		//返回登陆成功的结果
		ack := commonNewPorot.NewCommonAckGameLogin()
		*ack.Header.Code = consts.ACK_RESULT_SUCC
		*ack.UserId = user.GetId()
		*ack.NickName = user.GetNickName()
		*ack.Chip = user.GetDiamond()
		//*ack.IsSignToday = signService.IsUserSignedToday(user)
		a.WriteMsg(ack)

		loginService.DoLoginSuccess(user.GetId())
	}
	return

}
