package api

import (
	json "github.com/bytedance/sonic"
)

type (
	ReqInventoryGetSkuStock struct {
		BaseRequest
		SkuId string `json:"skuId"`
	}
)

var _ Request = new(ReqInventoryGetSkuStock)

func (r ReqInventoryGetSkuStock) Method() string {
	return "inventory.getSkuStock"
}

func (r ReqInventoryGetSkuStock) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespInventoryGetSkuStock struct {
		BaseResponse
		Data struct {
			SkuId    string `json:"skuId"`
			SkuStock struct {
				Available int64 `json:"available"`
				Total     int64 `json:"total"`
			} `json:"skuStock"`
			Response struct {
				Success bool   `json:"success"`
				Code    int64  `json:"code"`
				Message string `json:"message"`
			} `json:"response"`
		} `json:"data"`
	}
)

var _ Response = new(RespInventoryGetSkuStock)

func (r RespInventoryGetSkuStock) ErrorCode() int64 {

	return r.BaseResponse.ErrorCode
}

func (r RespInventoryGetSkuStock) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}

func (r RespInventoryGetSkuStock) Success() bool {
	return r.BaseResponse.Success
}
