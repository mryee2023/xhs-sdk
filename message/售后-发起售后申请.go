package message

import json "github.com/bytedance/sonic"

type MsgAfterSaleCreate struct {
	ReturnsId   string  `json:"returnsId"`
	OrderId     string  `json:"orderId"`
	ReturnType  int64   `json:"returnType"`
	RequestFrom int64   `json:"request_from"`
	RefundFee   float64 `json:"refundFee"`
	UpdateTime  int64   `json:"updateTime"`
}

var _ MessageData = new(MsgAfterSaleCreate)

func (MsgAfterSaleCreate) MsgTag() string {
	return "msg_after_sale_create"
}

func (MsgAfterSaleCreate) DecodeData(data string) (MsgAfterSaleCreate, error) {
	var resp MsgAfterSaleCreate
	err := json.Unmarshal([]byte(data), &resp)
	return resp, err
}
