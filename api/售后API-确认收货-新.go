package api

import json "github.com/bytedance/sonic"

//https://open.xiaohongshu.com/document/api?apiNavigationId=207&id=49&gatewayId=103&gatewayVersionId=1661&apiId=17602&apiParentNavigationId=16

type (
	ReqAfterSaleConfirmReceive struct {
		BaseRequest
		ReturnsId   string `json:"returnsId"`
		Action      int64  `json:"action"`
		Reason      int64  `json:"reason"`
		Description string `json:"description"`
	}
)

var _ Request = new(ReqAfterSaleConfirmReceive)

func (r ReqAfterSaleConfirmReceive) Method() string {
	return "afterSale.confirmReceive"
}

func (r ReqAfterSaleConfirmReceive) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespAfterSaleConfirmReceive struct {
		BaseResponse
	}
)

var _ Response = new(RespAfterSaleConfirmReceive)

func (r RespAfterSaleConfirmReceive) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespAfterSaleConfirmReceive) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespAfterSaleConfirmReceive) Success() bool {
	return r.BaseResponse.Success
}
