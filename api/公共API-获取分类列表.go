package api

import "encoding/json"

// 文档：https://open.xiaohongshu.com/document/api?apiNavigationId=2&id=12&gatewayId=103&gatewayVersionId=1661&apiId=5747

type ReqGetCategories struct {
	BaseRequest
	CategoryId string `json:"categoryId"`
}

var _ Request = new(ReqGetCategories)

func (r ReqGetCategories) Method() string {
	return "common.getCategories"
}

func (r ReqGetCategories) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetCategories struct {
		BaseResponse
		Data RespGetCategoriesData `json:"data"`
	}
	RespGetCategoriesData struct {
		CategoryV3S []struct {
			Id     string `json:"id"`
			Name   string `json:"name"`
			EnName string `json:"enName"`
		} `json:"categoryV3s"`
	}
)

var _ Response = new(RespGetCategories)

func (r RespGetCategories) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetCategories) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetCategories) Success() bool {
	return r.BaseResponse.Success
}
