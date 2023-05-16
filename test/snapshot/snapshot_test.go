package snapshot

// File       : snapshot_test.go
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
	poolName     = "pool1"
	subsysName   = "s1000"
	volumeName   = "v1000"
	snapshotName = "snap1000"
)

func setup() {
	fmt.Printf("create subsys %s\n", subsysName)
	createSubsysParameter := fassSDK.CreateSubsysParameter{}
	createSubsysParameter.SetPoolName(poolName)
	createSubsysParameter.SetCapacityGB(10)
	createSubsysParameter.SetSectorSize4096()
	createSubsysParameter.SetName(subsysName)
	createSubsysParameter.SetVolumeName(volumeName)
	createSubsysParameter.EnableISCSI()
	createSubsysParameter.SetFormatROW()
	_, err := fassSDK.CreateSubsysRequest(&createSubsysParameter, uuid.New().String())
	if err != nil {
		panic(fmt.Sprintf("create subsys %s failed due to %s\n", subsysName, err.Error()))
	}

	fmt.Printf("create subsys %s success\n", subsysName)
}

func teardown() {
	fmt.Printf("delete subsys %s\n", subsysName)
	deleteSubsysParameter := fassSDK.DeleteSubsysParameter{}
	deleteSubsysParameter.SetSubsysName(subsysName)
	deleteSubsysParameter.ForceDelete()
	deleteSubsysParameter.DeleteVolume()
	err := fassSDK.DeleteSubsysRequest(&deleteSubsysParameter, uuid.New().String())
	if err != nil {
		panic(fmt.Sprintf("delete subsys %s failed due to %s\n", subsysName, err.Error()))
	}

	fmt.Printf("delete subsys %s success\n", subsysName)
}

func TestCreateSnapshot(t *testing.T) {
	parameter := fassSDK.CreateSnapshotParameter{}
	parameter.SetVolumeName(volumeName)
	parameter.SetSnapshotName(snapshotName)
	_, err := fassSDK.CreateSnapshotRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveSnapshot(t *testing.T) {
	parameter := fassSDK.RetrieveSnapshotParameter{}
	parameter.SetVolumeName(volumeName)
	parameter.SetSnapshotName(snapshotName)
	_, err := fassSDK.RetrieveSnapshotRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestListSnapshot(t *testing.T) {
	parameter := fassSDK.ListSnapshotParameter{}
	parameter.SetVolumeName(volumeName)
	_, err := fassSDK.ListSnapshotRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestRevertSnapshot(t *testing.T) {
	parameter := fassSDK.RevertSnapshotParameter{}
	parameter.SetVolumeName(volumeName)
	parameter.SetSnapshotName(snapshotName)
	err := fassSDK.RevertSnapshotRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestDeleteSnapshot(t *testing.T) {
	parameter := fassSDK.DeleteSnapshotParameter{}
	parameter.SetVolumeName(volumeName)
	parameter.SetSnapshotName(snapshotName)
	err := fassSDK.DeleteSnapshotRequest(&parameter, uuid.New().String())
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
