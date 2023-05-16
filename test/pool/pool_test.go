package pool

// File       : pool_test.go
// Path       : requests
// Time       : CST 2023/4/25 17:28
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

var poolName = "pool3"

func TestCreatePool(t *testing.T) {
	parameter := fassSDK.CreatePoolParameter{}
	parameter.SetECRatio_4_2()
	parameter.SetSectorSize4096()
	parameter.SetPoolName(poolName)
	_, err := fassSDK.CreatePoolRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrievePool(t *testing.T) {
	parameter := fassSDK.RetrievePoolParameter{}
	parameter.SetPoolName(poolName)
	_, err := fassSDK.RetrievePoolRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestListPoolRequest(t *testing.T) {
	parameter := fassSDK.ListPoolParameter{}
	parameter.SetPageNum(1)
	parameter.SetPageSize(10)
	_, err := fassSDK.ListPoolRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestDeletePoolRequest(t *testing.T) {
	parameter := fassSDK.DeletePoolParameter{}
	parameter.SetPoolName(poolName)
	parameter.ForceDelete()
	err := fassSDK.DeletePoolRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}
