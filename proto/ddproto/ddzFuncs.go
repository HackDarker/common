package ddproto

import (
	"casino_common/utils/pokerUtils"
	"casino_common/utils/numUtils"
)

//得到牌的张数
func (out *DdzSrvOutPokerPais) GetPaiCount() int32 {
	return int32(len(out.PokerPais))
}

//牌相关的方法

func ( p *CommonSrvPokerPai) GetClientPoker() *ClientBasePoker {
	ret := new(ClientBasePoker)
	ret.Suit = new(CommonEnumPokerColor)
	ret.Id = p.Id
	ret.Num = p.Value
	*ret.Suit = p.GetSuit()
	return ret
}

func (p *CommonSrvPokerPai) GetSuit() CommonEnumPokerColor {
	if p.GetFlower() == pokerUtil.FLOWER_DIAMOND {
		return CommonEnumPokerColor_FANGKUAI
	} else if p.GetFlower() == pokerUtil.FLOWER_CLUB {
		return CommonEnumPokerColor_MEIHUA
	} else if p.GetFlower() == pokerUtil.FLOWER_HEART {
		return CommonEnumPokerColor_HONGTAO
	} else if p.GetFlower() == pokerUtil.FLOWER_SPADE {
		return CommonEnumPokerColor_HEITAO
	} else if p.GetFlower() == pokerUtil.FLOWER_BLACKJOKER {
		return CommonEnumPokerColor_BLACKBIGJOKER
	} else {
		return CommonEnumPokerColor_REDJOKER
	}
}

func (p *CommonSrvPokerPai) GetLogDes() string {
	suit, _ := numUtils.Int2String(p.GetId())
	switch p.GetSuit() {
	case CommonEnumPokerColor_FANGKUAI:
		suit = "方块"
	case CommonEnumPokerColor_BLACKBIGJOKER:
		suit = "小王"
	case CommonEnumPokerColor_HEITAO:
		suit = "黑桃"
	case CommonEnumPokerColor_HONGTAO:
		suit = "红桃"
	case CommonEnumPokerColor_REDJOKER:
		suit = "大王"
	case CommonEnumPokerColor_MEIHUA:
		suit = "梅花"
	default:
	}
	suit += p.GetName()
	return suit
}
