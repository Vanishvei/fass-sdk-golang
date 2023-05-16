package fass_sdk_golang

// File       : snapshot.go
// Path       : requests
// Time       : CST 2023/4/26 10:57
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk-golang/horizontal"
)

func RetrieveSnapshotRequest(parameter *RetrieveSnapshotParameter, requestId string) (
	*RetrieveSnapshotResponse, error) {
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

	data := &RetrieveSnapshotResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func ListSnapshotRequest(parameter *ListSnapshotParameter, requestId string) (
	*ListSnapshotResponse, error) {
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

	data := &ListSnapshotResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreateSnapshotRequest(parameter *CreateSnapshotParameter, requestId string) (
	*CreateSnapshotResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("snapshot")
	_request.SetBody(parameter)
	_request.SetMethodPOST()

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &CreateSnapshotResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RevertSnapshotRequest(parameter *RevertSnapshotParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetBody(parameter)
	_request.SetPath("snapshot")
	_request.SetMethodPUT()

	_, err = _client.CallApi(_request)
	return err
}

func DeleteSnapshotRequest(parameter *DeleteSnapshotParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodDELETE()

	_, err = _client.CallApi(_request)
	return err
}
