package fass_sdk_golang

// File       : pool.go
// Path       : requests
// Time       : CST 2023/4/24 15:35
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk-golang/horizontal"
)

func ListPoolRequest(parameter *ListPoolParameter, requestId string) (
	*ListPoolResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetQuery(parameter.GetQuery())
	_request.SetPath("pool")

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &ListPoolResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrievePoolRequest(parameter *RetrievePoolParameter, requestId string) (
	*RetrievePoolResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &RetrievePoolResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreatePoolRequest(parameter *CreatePoolParameter, requestId string) (
	*CreatePoolResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("pool")
	_request.SetMethodPOST()
	_request.SetBody(parameter)

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &CreatePoolResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeletePoolRequest(parameter *DeletePoolParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetMethodDELETE()
	_request.SetPath(parameter.GetPath())
	_request.SetQuery(parameter.GetQuery())

	_, err = _client.CallApi(_request)
	return err
}
