package api

import (
	json "github.com/bytedance/sonic"
)

type (
	ReqProductGetItemInfo struct {
		BaseRequest
		ItemId   string `json:"itemId"`
		PageSize int64  `json:"pageSize"`
		PageNo   int64  `json:"pageNo"`
	}
)

var _ Request = new(ReqProductGetItemInfo)

func (r ReqProductGetItemInfo) Method() string {
	return "product.getItemInfo"
}

func (r ReqProductGetItemInfo) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespProductGetItemInfo struct {
		BaseResponse
		Data RespProductGetItemInfoData `json:"data"`
	}
	RespProductGetItemInfoData struct {
		Total    int64            `json:"total"`
		SkuInfos []ProductSkuInfo `json:"skuInfos"`
	}
	ProductSkuInfo struct {
		ItemId  string `json:"itemId"`
		Id      string `json:"id"`
		Name    string `json:"name"`
		ErpCode string `json:"erpCode"`
		Price   int64  `json:"price"`
		Stock   int64  `json:"stock"`
		Whcode  string `json:"whcode"` // 仓库编码
		IsGift  bool   `json:"isGift"` // 是否是赠品
	}
)

var _ Response = new(RespProductGetItemInfo)

func (r RespProductGetItemInfo) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespProductGetItemInfo) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}

func (r RespProductGetItemInfo) Success() bool {
	return r.BaseResponse.Success
}
