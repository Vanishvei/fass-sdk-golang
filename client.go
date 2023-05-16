package fass_sdk_golang

// File       : client.go
// Path       : client
// Time       : CST 2023/4/10 14:05
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk-golang/horizontal"
)

type Client struct {
	Port       *int
	Endpoint   *string
	Method     *string
	Headers    map[string]*string
	RunTime    *horizontal.RuntimeObject
	ApiVersion *string
}

func NewClient() (*Client, error) {
	client := new(Client)
	err := client.Init(&GlobalConfig)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if horizontal.BoolValue(horizontal.IsUnset(config)) {
		_err = horizontal.NewSDKError(map[string]interface{}{
			"code":       2,
			"message":    "'config' can not be unset",
			"data":       nil,
			"request_id": "",
		})
		return _err
	}

	client.RunTime = &horizontal.RuntimeObject{
		ConnectTimeout: config.ConnectTimeout,
		ReadTimeout:    config.ReadTimeout,
		Backoff:        config.Backoff,
		Retry:          config.RetryCount,
	}

	client.Port = config.Port
	client.Endpoint = config.CurrentEndpoint
	client.ApiVersion = config.ApiVersion
	return nil
}

func (client *Client) DoRequest(request *horizontal.Request) (_result *SuzakuResponse, _err error) {
	_err = horizontal.Validate(request)
	if _err != nil {
		return _result, _err
	}

	_runtime := map[string]interface{}{
		"retry":          horizontal.IntValue(horizontal.DefaultNumber(client.RunTime.Retry, horizontal.Int(2))),
		"backoff":        horizontal.IntValue(horizontal.DefaultNumber(client.RunTime.Backoff, horizontal.Int(1))),
		"readTimeout":    horizontal.IntValue(horizontal.DefaultNumber(client.RunTime.ReadTimeout, horizontal.Int(60))),
		"connectTimeout": horizontal.IntValue(horizontal.DefaultNumber(client.RunTime.ConnectTimeout, horizontal.Int(5))),
	}

	globalQueries := make(map[string]*string)
	globalHeaders := make(map[string]*string)
	request.UpdateQuery(horizontal.Merge(globalQueries, request.GetQuery()))
	request.Headers = horizontal.Merge(map[string]*string{
		"host":         request.GetEndpoint(),
		"requestId":    request.GetRequestId(),
		"user-agent":   horizontal.String("fass_sdk-golang/v1.0"),
		"accept":       horizontal.String("application/json"),
		"content-type": horizontal.String("application/json; charset=utf-8"),
	},
		globalHeaders,
		request.Headers,
	)

	GlobalConfig.SwitchEndpoint()

	_resp := &SuzakuResponse{}
	for _retryTimes := 0; horizontal.BoolValue(horizontal.AllowRetry(_runtime["retry"], horizontal.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := horizontal.GetBackoffTime(_runtime["backoff"], horizontal.Int(_retryTimes))
			if horizontal.IntValue(_backoffTime) > 0 {
				horizontal.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*SuzakuResponse, error) {
			response_, _err := horizontal.DoRequest(request, _runtime)
			if _err != nil {
				return _result, _err
			}

			if *response_.StatusCode == 200 {
				_res, _err := horizontal.ReadAsJSON(response_.Body)
				if _err != nil {
					return _result, _err
				}

				_, _err = horizontal.AssertAsMap(_res)
				if _err != nil {
					return _result, _err
				}

				_ = horizontal.Convert(_res, &_result)
				return _result, _err
			}

			_err = horizontal.Convert(map[string]interface{}{
				"headers":    response_.Headers,
				"statusCode": horizontal.IntValue(response_.StatusCode),
			}, &_result)
			return _result, _err

		}()
		if !horizontal.BoolValue(horizontal.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) CallApi(request *horizontal.Request) (_result *SuzakuResponse, _err error) {
	if horizontal.BoolValue(horizontal.IsUnset(request)) {
		_err = horizontal.NewSDKError(map[string]interface{}{
			"code":       2,
			"message":    "'params' can not be unset",
			"request_id": "",
			"data":       nil,
		})
		return _result, _err
	}

	request.SetPort(client.Port)
	request.SetEndpoint(client.Endpoint)
	request.SetApiVersion(client.ApiVersion)

	_resp, _err := client.DoRequest(request)
	if _err != nil {
		_err = horizontal.NewSDKError(map[string]interface{}{
			"code":       1,
			"message":    _err.Error(),
			"request_id": _resp.RequestId.String(),
			"data":       nil,
		})
		return _result, _err
	}

	if _resp.Code != 0 {
		_err = horizontal.NewSDKError(map[string]interface{}{
			"code":       _resp.Code,
			"message":    _resp.Message,
			"request_id": _resp.RequestId.String(),
			"data":       _resp.Data,
		})
		return _result, _err
	}

	return _resp, _err
}
