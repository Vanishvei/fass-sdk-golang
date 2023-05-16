package subsys

// File       : subsys_test.go
// Path       : requests
// Time       : CST 2023/4/26 11:06
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"

	fassSDK "github.com/Vanishvei/fass-sdk-golang"
	_ "github.com/Vanishvei/fass-sdk-golang/test"
)

var (
	poolName      = "pool1"
	subsysName    = "s1000"
	volumeName    = "v1000"
	accountName   = "account_1"
	password      = "admin@1234"
	groupName     = "group_1"
	hostName      = "client_1"
	qualifierList = []string{
		"nqn.2019-03.suzaku:s1000",
		"iqn.2019-03.cn.suzaku:s1000",
	}
)

func setup() {
	createAccountParameter := fassSDK.CreateAccountParameter{}
	createAccountParameter.SetAccountName(accountName)
	createAccountParameter.SetPassword(password)

	_, err := fassSDK.CreateAccountRequest(&createAccountParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		panic(fmt.Sprintf("create account %s failed\n", accountName))
	}

	createGroupParameter := fassSDK.AddQualifierToGroupParameter{}
	createGroupParameter.SetQualifierList(qualifierList)
	createGroupParameter.SetGroupName(groupName)
	createGroupParameter.SetHostName(hostName)

	_, err = fassSDK.AddQualifierToGroupRequest(&createGroupParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		panic(fmt.Sprintf("add qualifier to group %s failed\n", groupName))
	}
}

func teardown() {
	deleteGroupParameter := fassSDK.DeleteGroupParameter{}
	deleteGroupParameter.SetGroupName(groupName)
	err := fassSDK.DeleteGroupRequest(&deleteGroupParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("delete host group %s failed\n", groupName)
	}

	deleteAccountParameter := fassSDK.DeleteAccountParameter{}
	deleteAccountParameter.SetAccountName(accountName)

	err = fassSDK.DeleteAccountRequest(&deleteAccountParameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("delete account %s failed\n", accountName)
	}
}

func TestCreateSubsys(t *testing.T) {
	parameter := fassSDK.CreateSubsysParameter{}
	parameter.SetName(subsysName)
	parameter.SetPoolName(poolName)
	parameter.SetVolumeName(volumeName)
	parameter.SetSectorSize4096()
	parameter.SetCapacityGB(10)
	parameter.EnableISCSI()
	parameter.SetBpsMB(1000)
	parameter.SetIops(2000)
	parameter.SetFormatROW()

	_, err := fassSDK.CreateSubsysRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSubsys(t *testing.T) {
	parameter := fassSDK.RetrieveSubsysParameter{}
	parameter.SetSubsysName(subsysName)

	_, err := fassSDK.RetrieveSubsysRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestListSubsys(t *testing.T) {
	parameter := fassSDK.ListSubsysParameter{}

	_, err := fassSDK.ListSubsysRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestExportSubsys(t *testing.T) {
	parameter := fassSDK.ExportSubsysParameter{}
	parameter.SetSubsysName(subsysName)
	parameter.ExportISCSI()

	err := fassSDK.ExportSubsysRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestUnexportSubsys(t *testing.T) {
	parameter := fassSDK.UnexportSubsysParameter{}
	parameter.SetSubsysName(subsysName)
	parameter.UnexportISCSI()

	err := fassSDK.UnexportSubsysRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestSetSubsysAuth(t *testing.T) {
	parameter := fassSDK.SetSubsysAuthParameter{}
	parameter.SetSubsysName(subsysName)
	parameter.SetGroupName(groupName)

	err := fassSDK.SetSubsysAuthRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSubsysAuth(t *testing.T) {
	parameter := fassSDK.RetrieveSubsysAuthParameter{}
	parameter.SetSubsysName(subsysName)

	resp, err := fassSDK.RetrieveSubsysAuthRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.Auth, groupName) {
		fmt.Printf("auth result %s is not equal to %s\n", resp.Auth, groupName)
		t.FailNow()
	}
}

func TestSetSubsysChap(t *testing.T) {
	parameter := fassSDK.SetSubsysChapParameter{}
	parameter.SetSubsysName(subsysName)
	parameter.SetAccountName(accountName)

	err := fassSDK.SetSubsysChapRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSubsysChap(t *testing.T) {
	parameter := fassSDK.RetrieveSubsysChapParameter{}
	parameter.SetSubsysName(subsysName)

	resp, err := fassSDK.RetrieveSubsysChapRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}

	if !reflect.DeepEqual(resp.Chap, accountName) {
		fmt.Printf("chap result %s is not equal to %s\n", resp.Chap, accountName)
		t.FailNow()
	}
}

func TestRemoveSubsysAuth(t *testing.T) {
	parameter := fassSDK.RemoveSubsysAuthParameter{}
	parameter.SetSubsysName(subsysName)

	err := fassSDK.RemoveSubsysAuthRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRemoveSubsysChap(t *testing.T) {
	parameter := fassSDK.RemoveSubsysChapParameter{}
	parameter.SetSubsysName(subsysName)

	err := fassSDK.RemoveSubsysChapRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestDeleteSubsys(t *testing.T) {
	parameter := fassSDK.DeleteSubsysParameter{}
	parameter.SetSubsysName(subsysName)
	parameter.DeleteVolume()
	parameter.ForceDelete()

	err := fassSDK.DeleteSubsysRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
