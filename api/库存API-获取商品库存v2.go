package api

import (
	json "github.com/bytedance/sonic"
)

type (
	ReqInventoryGetSkuStockV2 struct {
		BaseRequest
		SkuId string `json:"skuId"`
	}
)

var _ Request = new(ReqInventoryGetSkuStockV2)

func (r ReqInventoryGetSkuStockV2) Method() string {
	return "inventory.getSkuStockV2"
}

func (r ReqInventoryGetSkuStockV2) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespInventoryGetSkuStockV2 struct {
		BaseResponse
		Data struct {
			SkuStockInfo struct {
				SkuId        string `json:"skuId"`
				SkuStockInfo struct {
					Available int64 `json:"available"` // 可售库存
					Total     int64 `json:"total"`     // 总库存
				} `json:"skuStockInfo"`
			} `json:"skuStockInfo"`
		} `json:"data"`
	}
)

var _ Response = new(RespInventoryGetSkuStockV2)

func (r RespInventoryGetSkuStockV2) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespInventoryGetSkuStockV2) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}

func (r RespInventoryGetSkuStockV2) Success() bool {
	return r.BaseResponse.Success
}
