package fass_sdk_golang

// File       : acl.go
// Path       : requests
// Time       : CST 2023/5/5 9:53
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk-golang/horizontal"
	//"github.com/Vanishvei/fass-sdk-golang/parameters"
	//"github.com/Vanishvei/fass-sdk-golang/responses"
)

func ListAccountRequest(parameter *ListAccountParameter, requestId string) (
	*ListAccountResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetQuery(parameter.GetQuery())
	_request.SetPath("acl/account")

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &ListAccountResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func CreateAccountRequest(parameter *CreateAccountParameter, requestId string) (
	*CreateAccountResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("acl/account")
	_request.SetMethodPOST()
	_request.SetBody(parameter)

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &CreateAccountResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveAccountRequest(parameter *RetrieveAccountParameter, requestId string) (
	*RetrieveAccountResponse, error) {
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

	data := &RetrieveAccountResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteAccountRequest(parameter *DeleteAccountParameter, requestId string) error {
	_client, err := NewClient()
	if err != nil {
		return err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())

	_, err = _client.CallApi(_request)
	return err
}

func ListGroupRequest(parameter *ListAccountParameter, requestId string) (
	*ListGroupResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}

	_request := horizontal.NewRequest(requestId)
	_request.SetPath("acl/group")

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &ListGroupResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RetrieveGroupRequest(parameter *RetrieveGroupParameter, requestId string) (
	*RetrieveGroupResponse, error) {
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

	data := &RetrieveGroupResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func DeleteGroupRequest(parameter *DeleteGroupParameter, requestId string) error {
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

func AddQualifierToGroupRequest(parameter *AddQualifierToGroupParameter, requestId string) (
	*AddQualifierToGroupResponse, error) {
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

	data := &AddQualifierToGroupResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err
}

func RemoveQualifierFromGroupRequest(parameter *RemoveQualifierFromGroupParameter, requestId string) (
	*RemoveQualifierFromGroupResponse, error) {
	_client, err := NewClient()
	if err != nil {
		return nil, err
	}
	_request := horizontal.NewRequest(requestId)
	_request.SetPath(parameter.GetPath())
	_request.SetMethodPUT()
	_request.SetBody(parameter)

	resp, err := _client.CallApi(_request)
	if err != nil {
		return nil, err
	}

	data := &RemoveQualifierFromGroupResponse{}
	err = horizontal.ConvertToSuzakuResp(resp.Data, data)
	return data, err

}
