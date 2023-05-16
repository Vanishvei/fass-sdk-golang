package acl

// File       : acl_test.go
// Path       : test/acl
// Time       : CST 2023/5/5 14:55
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"

	fassSDK "github.com/Vanishvei/fass-sdk-golang"
	_ "github.com/Vanishvei/fass-sdk-golang/test"
)

var (
	hostName        = "client_1"
	groupName       = "group_1"
	accountName     = "account_1"
	accountPassword = "admin@123"

	addQualifierList = []string{
		"nqn.2019-06.suzaku:subsys001",
		"iqn.2019-03.cn.suzaku:subsys001",
		"iqn.2019-03.cn.suzaku:subsys002",
	}
	removeQualifierList = []string{
		"nqn.2019-06.suzaku:subsys001",
		"iqn.2019-03.cn.suzaku:subsys001",
	}
)

func TestCreateAccount(t *testing.T) {
	parameter := fassSDK.CreateAccountParameter{}
	parameter.SetAccountName(accountName)
	parameter.SetPassword(accountPassword)
	_, err := fassSDK.CreateAccountRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestListAccount(t *testing.T) {
	parameter := fassSDK.ListAccountParameter{}
	_, err := fassSDK.ListAccountRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveAccount(t *testing.T) {
	parameter := fassSDK.RetrieveAccountParameter{}
	parameter.SetAccountName(accountName)

	resp, err := fassSDK.RetrieveAccountRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.AccountName, accountName) {
		fmt.Printf("account result %s is not equal to %s\n", resp.AccountName, accountName)
		t.FailNow()
	}
}

func TestDeleteAccount(t *testing.T) {
	parameter := fassSDK.DeleteAccountParameter{}
	parameter.SetAccountName(accountName)
	err := fassSDK.DeleteAccountRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestAddQualifierToGroup(t *testing.T) {
	parameter := fassSDK.AddQualifierToGroupParameter{}
	parameter.SetGroupName(groupName)
	parameter.SetHostName(hostName)
	parameter.SetQualifierList(addQualifierList)
	_, err := fassSDK.AddQualifierToGroupRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRemoveQualifierFromGroup(t *testing.T) {
	parameter := fassSDK.RemoveQualifierFromGroupParameter{}
	parameter.SetGroupName(groupName)
	parameter.SetHostName(hostName)
	parameter.SetQualifierList(removeQualifierList)
	_, err := fassSDK.RemoveQualifierFromGroupRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveGroup(t *testing.T) {
	parameter := fassSDK.RetrieveGroupParameter{}
	parameter.SetGroupName(groupName)

	resp, err := fassSDK.RetrieveGroupRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.GroupName, groupName) {
		fmt.Printf("group result %s is not equal to %s\n", resp.GroupName, groupName)
		t.FailNow()
	}
}

func TestListGroup(t *testing.T) {
	parameter := fassSDK.ListGroupParameter{}
	_, err := fassSDK.ListGroupRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestDeleteGroup(t *testing.T) {
	parameter := fassSDK.DeleteGroupParameter{}
	parameter.SetGroupName(groupName)
	err := fassSDK.DeleteGroupRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}
