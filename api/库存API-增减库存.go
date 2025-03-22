package api

import (
	json "github.com/bytedance/sonic"
)

type (
	ReqInventoryIncSkuStock struct {
		BaseRequest
		SkuId string `json:"skuId"`
		Qty   int64  `json:"qty"`
	}
)

var _ Request = new(ReqInventoryIncSkuStock)

func (r ReqInventoryIncSkuStock) Method() string {
	return "inventory.incSkuStock"
}

func (r ReqInventoryIncSkuStock) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespInventoryIncSkuStock struct {
		BaseResponse
		Data struct {
			SkuStock struct {
				Available int64 `json:"available"`
				Total     int64 `json:"total"`
			} `json:"skuStock"`
		} `json:"data"`
	}
)

var _ Response = new(RespInventoryIncSkuStock)

func (r RespInventoryIncSkuStock) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespInventoryIncSkuStock) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}

func (r RespInventoryIncSkuStock) Success() bool {
	return r.BaseResponse.Success
}
