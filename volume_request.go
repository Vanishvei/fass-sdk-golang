package fass_sdk_golang

// File       : volume.go
// Path       : requests
// Time       : CST 2023/4/26 10:57
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk-golang/horizontal"
)

func ListVolumeRequest(parameter *ListVolumeParameter, requestId string) (
	*ListVolumeResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("volume")
	_request.SetQuery(parameter.GetQuery())

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &ListVolumeResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveVolumeRequest(parameter *RetrieveVolumeParameter, requestId string) (
	*RetrieveVolumeResponse, error) {
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

	data := &RetrieveVolumeResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteVolumeRequest(parameter *DeleteVolumeParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetQuery(parameter.GetQuery())
	_request.SetMethodDELETE()

	_, err = _client.CallApi(_request)
	return err
}

func ExpandVolumeRequest(parameter *ExpandVolumeParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)
	_request.SetMethodPUT()

	_, err = _client.CallApi(_request)
	return err
}

func FlattenVolumeRequest(parameter *FlattenVolumeParameter, requestId string) (
	*FlattenVolumeResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)
	_request.SetMethodPUT()

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &FlattenVolumeResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func SetQosOfVolumeRequest(parameter *SetQosOfVolumeParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}
	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)
	_request.SetMethodPUT()

	_, err = _client.CallApi(_request)
	return err
}

func FlattenVolumeProgressRequest(parameter *GetFlattenVolumeProgress, requestId string) (
	*FlattenVolumeProgressResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetBody(parameter)

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &FlattenVolumeProgressResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func StopFlattenVolumeRequest(parameter *StopFlattenVolumeParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())

	_, err = _client.CallApi(_request)
	return err
}
