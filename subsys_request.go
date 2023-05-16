package fass_sdk_golang

// File       : subsys.go
// Path       : requests
// Time       : CST 2023/4/26 10:57
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk-golang/horizontal"
)

func CreateSubsysRequest(parameter *CreateSubsysParameter, requestId string) (
	*CreateSubsysResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetBody(parameter)
	_request.SetPath("subsys")
	_request.SetMethodPOST()
	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &CreateSubsysResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func ListSubsysRequest(parameter *ListSubsysParameter, requestId string) (
	*ListSubsysResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("pool")
	_request.SetQuery(parameter.GetQuery())

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &ListSubsysResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveSubsysRequest(parameter *RetrieveSubsysParameter, requestId string) (
	*RetrieveSubsysResponse, error) {
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

	data := &RetrieveSubsysResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteSubsysRequest(parameter *DeleteSubsysParameter, requestId string) error {
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

func ExportSubsysRequest(parameter *ExportSubsysParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	_, err = _client.CallApi(_request)
	return err
}

func UnexportSubsysRequest(parameter *UnexportSubsysParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	_, err = _client.CallApi(_request)
	return err
}

func RetrieveSubsysAuthRequest(parameter *RetrieveSubsysAuthParameter,
	requestId string) (*RetrieveSubsysAuthResponse, error) {
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

	data := &RetrieveSubsysAuthResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func SetSubsysAuthRequest(parameter *SetSubsysAuthParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	_, err = _client.CallApi(_request)
	return err
}

func RemoveSubsysAuthRequest(parameter *RemoveSubsysAuthParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()

	_, err = _client.CallApi(_request)
	return err
}

func RetrieveSubsysChapRequest(parameter *RetrieveSubsysChapParameter, requestId string) (
	*RetrieveSubsysChapResponse, error) {
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

	data := &RetrieveSubsysChapResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func SetSubsysChapRequest(parameter *SetSubsysChapParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	_, err = _client.CallApi(_request)
	return err
}

func RemoveSubsysChapRequest(parameter *RemoveSubsysChapParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()

	_, err = _client.CallApi(_request)
	return err
}
