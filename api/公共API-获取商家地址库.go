package api

import json "github.com/bytedance/sonic"

type ReqGetAddressRecord struct {
	BaseRequest
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
}

var _ Request = new(ReqGetAddressRecord)

func (r ReqGetAddressRecord) Method() string {
	return "common.getAddressRecord"
}

func (r ReqGetAddressRecord) Params(base BaseRequest) []byte {
	r.BaseRequest = base
	ret, _ := json.Marshal(r)
	return ret
}

type (
	RespGetAddressRecord struct {
		BaseResponse
		Data RespGetAddressRecordData `json:"data"`
	}
	RespGetAddressRecordData struct {
		Total                   int64           `json:"total"`
		SellerAddressRecordList []AddressRecord `json:"sellerAddressRecordList"`
	}

	AddressRecord struct {
		SellerAddressRecordId   int64  `json:"sellerAddressRecordId"`
		ContactName             string `json:"contactName"`
		PhoneAreaCode           string `json:"phoneAreaCode"`
		Phone                   string `json:"phone"`
		LandlineAreaCode        string `json:"landlineAreaCode"`
		LandlinePhone           string `json:"landlinePhone"`
		LandlineExtensionNumber string `json:"landlineExtensionNumber"`
		CountryCode             string `json:"countryCode"`
		ProvinceCode            string `json:"provinceCode"`
		CityCode                string `json:"cityCode"`
		CountyCode              string `json:"countyCode"`
		TownCode                string `json:"townCode"`
		CountryName             string `json:"countryName"`
		ProvinceName            string `json:"provinceName"`
		CityName                string `json:"cityName"`
		CountyName              string `json:"countyName"`
		TownName                string `json:"townName"`
		Address                 string `json:"address"`
		FullAddress             string `json:"fullAddress"`
		DeliveryDefault         string `json:"deliveryDefault"`
		AftersaleDefault        string `json:"aftersaleDefault"`
		Version                 int    `json:"version"`
		Active                  string `json:"active"`
		CreateTime              int64  `json:"createTime"`
		UpdateTime              int64  `json:"updateTime"`
	}
)

var _ Response = new(RespGetAddressRecord)

func (r RespGetAddressRecord) ErrorCode() int64 {
	return r.BaseResponse.ErrorCode
}

func (r RespGetAddressRecord) ErrorMsg() string {
	return r.BaseResponse.ErrorMsg

}

func (r RespGetAddressRecord) Success() bool {
	return r.BaseResponse.Success
}
