package api

import (
	json "github.com/bytedance/sonic"
)

type (
	ReqInventorySyncSkuStock struct {
		BaseRequest
		SkuId string `json:"skuId"`
		Qty   int64  `json:"qty"`
	}
)

var _ Request = new(ReqInventorySyncSkuStock)

func (r ReqInventorySyncSkuStock) Method() string {
	return "inventory.syncSkuStock"
}

func (r ReqInventorySyncSkuStock) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespInventorySyncSkuStock struct {
		BaseResponse
		Data struct {
			SkuStock struct {
				Available int64 `json:"available"`
				Total     int64 `json:"total"`
			} `json:"skuStock"`
		} `json:"data"`
	}
)

var _ Response = new(RespInventorySyncSkuStock)

func (r RespInventorySyncSkuStock) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespInventorySyncSkuStock) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}

func (r RespInventorySyncSkuStock) Success() bool {
	return r.BaseResponse.Success
}
