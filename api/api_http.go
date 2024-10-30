package api

import (
	"errors"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/valyala/fasthttp"
)

const HttpTTL = 1 * time.Minute

var FastClient = CreateFastHttpClient()

func CreateFastHttpClient() fasthttp.Client {
	var defaultDialer = &fasthttp.TCPDialer{Concurrency: 300}

	return fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			idx := 3 // 重试三次
			for {
				idx--
				conn, err := defaultDialer.DialTimeout(addr, 10*time.Second) // tcp连接超时时间10s
				if err != fasthttp.ErrDialTimeout || idx == 0 {
					return conn, err
				}
			}
		},
	}
}

var restyClient *resty.Client
var _once sync.Once

func Transport200() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,

		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          1000,             //针对整个Client的所有idle池中的连接数的和
		MaxIdleConnsPerHost:   20,               //在client进入闲时时要进入idle状态的最大连接数。
		MaxConnsPerHost:       100,              //某个host建立最大连接数的设置
		IdleConnTimeout:       60 * time.Second, //IdleConnTimeout字段用于超时清理idle池中的长连接
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func GetRestyClient() *resty.Client {
	_once.Do(func() {
		restyClient = resty.New()
		restyClient.SetTimeout(time.Millisecond * 1500)
		restyClient.SetRetryCount(1)
		restyClient.SetRetryWaitTime(time.Millisecond * 100)
		restyClient.SetRetryMaxWaitTime(time.Second)
		//restyClient.SetDebug(os.Getenv("xhs_debug") == "true")
		restyClient.SetTransport(Transport200())
	})
	return restyClient
}

func PostWithRestyClient(req []byte) ([]byte, error) {

	cli := GetRestyClient()

	cli.SetDebug(os.Getenv("xhs_debug") == "true")
	if cli.Debug {
		cli.EnableTrace()
		cli.EnableGenerateCurlOnDebug()
	}
	cli.SetHeader("Content-Type", "application/json")
	resp, err := cli.R().SetBody(req).Post(ApiUrl)
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.StatusCode() != http.StatusOK {
		return nil, errors.New("http resp is nil")
	}
	if len(resp.Body()) == 0 { // 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
		return nil, errors.New("http resp body is nil")
	}

	return resp.Body(), nil
}

func Post(req []byte) ([]byte, error) {
	return PostWithRestyClient(req)
	/*
		var resp []byte
		httpReq := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(httpReq)

		httpResp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(httpResp)

		httpReq.SetRequestURI(ApiUrl)
		httpReq.Header.SetContentType("application/json;charset=utf-8")
		httpReq.SetBody(req)
		httpReq.Header.SetMethod(http.MethodPost)

		if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
			return resp, err
		}

		resp = httpResp.Body()
		if len(resp) == 0 { // 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
			return resp, errors.New("http resp body is nil")
		}

		return resp, nil
		*
	*/
}
