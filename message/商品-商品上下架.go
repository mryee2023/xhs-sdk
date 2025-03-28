package message

import json "github.com/bytedance/sonic"

type MsgSkuBuyable struct {
	SkuId      string `json:"skuId"`
	UpdateTime int64  `json:"updateTime"`
}

var _ MessageData = new(MsgSkuBuyable)

func (MsgSkuBuyable) MsgTag() string {
	return "msg_sku_buyable"
}

func (MsgSkuBuyable) DecodeData(data string) (MsgSkuBuyable, error) {
	var resp MsgSkuBuyable
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
