package message

import json "github.com/bytedance/sonic"

type MsgFulfillmentStatusChange struct {
	OrderId     string `json:"orderId"`
	UpdateTime  int64  `json:"updateTime"`
	OrderStatus int64  `json:"orderStatus"`
}

var _ MessageData = new(MsgFulfillmentStatusChange)

func (MsgFulfillmentStatusChange) MsgTag() string {
	return "msg_fulfillment_status_change"
}

func (MsgFulfillmentStatusChange) DecodeData(data string) (MsgFulfillmentStatusChange, error) {
	var resp MsgFulfillmentStatusChange
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
