package api

import (
	json "github.com/bytedance/sonic"
)

type (
	ReqProductSearchItemList struct {
		BaseRequest
		PageNo      int64 `json:"pageNo"`
		PageSize    int64 `json:"pageSize"`
		SearchParam struct {
			Keywords []string `json:"keywords"`
		} `json:"searchParam"`
	}
)

var _ Request = new(ReqProductSearchItemList)

func (r ReqProductSearchItemList) Method() string {
	return "product.searchItemList"
}

func (r ReqProductSearchItemList) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespProductSearchItemList struct {
		BaseResponse
		Data RespProductSearchItemListData `json:"data"`
	}
	RespProductSearchItemListData struct {
		CurrentPage   int64                                        `json:"currentPage"`
		PageSize      int64                                        `json:"pageSize"`
		Total         int64                                        `json:"total"`
		ItemDetailV3s []RespProductSearchItemListDataItemDetailV3s `json:"itemDetailV3s"`
	}
	RespProductSearchItemListDataItemDetailV3s struct {
		Name           string `json:"name"`
		Id             string `json:"id"` // itemId
		ItemShortTitle string `json:"itemShortTitle"`
	}
)

var _ Response = new(RespProductSearchItemList)

func (r RespProductSearchItemList) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespProductSearchItemList) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg
}

func (r RespProductSearchItemList) Success() bool {
	return r.BaseResponse.Success
}
