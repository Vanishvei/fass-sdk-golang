package volume

// File       : volume_test.go
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
	"time"

	"github.com/google/uuid"

	fassSDK "github.com/Vanishvei/fass-sdk-golang"
	_ "github.com/Vanishvei/fass-sdk-golang/test"
)

var (
	poolName        = "pool1"
	srcSubsysName   = "s1000"
	srcVolumeName   = "v1000"
	srcSnapshotName = "snap1000"
	newSubsysName1  = "s2000"
	newVolumeName1  = "v2000"
	newSubsysName2  = "s3000"
	newVolumeName2  = "v3000"
	taskId          = ""
)

func setup() {
	fmt.Printf("create source subsys %s\n", srcSubsysName)
	createSubsysParameter := fassSDK.CreateSubsysParameter{}
	createSubsysParameter.SetPoolName(poolName)
	createSubsysParameter.SetCapacityGB(10)
	createSubsysParameter.SetSectorSize4096()
	createSubsysParameter.SetName(srcSubsysName)
	createSubsysParameter.SetVolumeName(srcVolumeName)
	createSubsysParameter.EnableISCSI()
	createSubsysParameter.SetFormatROW()
	_, err := fassSDK.CreateSubsysRequest(&createSubsysParameter, uuid.New().String())
	if err != nil {
		panic(fmt.Sprintf("create source subsys %s failed due to %s\n", srcSubsysName, err.Error()))
	}
	fmt.Printf("create source subsys %s success\n", srcSubsysName)

	time.Sleep(3 * time.Second)

	fmt.Printf("create source snapshot %s\n", srcSnapshotName)
	createSnapshotParameter := fassSDK.CreateSnapshotParameter{}
	createSnapshotParameter.SetVolumeName(srcVolumeName)
	createSnapshotParameter.SetSnapshotName(srcSnapshotName)

	_, err = fassSDK.CreateSnapshotRequest(&createSnapshotParameter, uuid.New().String())
	if err != nil {
		panic(fmt.Sprintf("create source snapshot %s failed due to %s\n", srcSnapshotName, err.Error()))
	}

	fmt.Printf("create source snapshot %s success\n", srcSnapshotName)

	err = createVolumeFromSnapshot(newSubsysName2, newVolumeName2)
	if err != nil {
		panic(fmt.Sprintf("create volume %s success\n", newVolumeName2))
	}
}

func teardown() {
	fmt.Printf("delete source snapshot %s\n", srcSnapshotName)
	deleteSnapshotParameter := fassSDK.DeleteSnapshotParameter{}
	deleteSnapshotParameter.SetVolumeName(srcVolumeName)
	deleteSnapshotParameter.SetSnapshotName(srcSnapshotName)

	err := fassSDK.DeleteSnapshotRequest(&deleteSnapshotParameter, uuid.New().String())
	if err != nil {
		fmt.Printf("delete source snapshot %s failed due to %s\n", srcSnapshotName, err.Error())
	} else {
		fmt.Printf("delete source snapshot %s success\n", srcSnapshotName)
	}

	fmt.Printf("delete source subsys %s\n", srcSubsysName)
	deleteSubsysParameter := fassSDK.DeleteSubsysParameter{}
	deleteSubsysParameter.SetSubsysName(srcSubsysName)
	deleteSubsysParameter.ForceDelete()
	deleteSubsysParameter.DeleteVolume()

	err = fassSDK.DeleteSubsysRequest(&deleteSubsysParameter, uuid.New().String())
	if err != nil {
		fmt.Printf("delete source subsys %s failed due to %s\n", srcSubsysName, err.Error())
	} else {
		fmt.Printf("delete source subsys %s success\n", srcSubsysName)
	}

	err = deleteVolume(newVolumeName2, 3)
	if err != nil {
		fmt.Printf("delete volume %s failed due to %s\n", newVolumeName2, err.Error())
	} else {
		fmt.Printf("delete volume %s success\n", newVolumeName2)
	}
}

func createVolumeFromSnapshot(subsysName, volumeName string) error {
	createSubsysParameter := fassSDK.CreateSubsysParameter{}
	createSubsysParameter.SetPoolName(poolName)
	createSubsysParameter.SetName(subsysName)
	createSubsysParameter.SetVolumeName(volumeName)
	createSubsysParameter.SetSrcVolumeName(srcVolumeName)
	createSubsysParameter.SetSrcSnapshotName(srcSnapshotName)
	createSubsysParameter.EnableISCSI()

	_, err := fassSDK.CreateSubsysRequest(&createSubsysParameter, uuid.New().String())
	if err != nil {
		return err
	}

	deleteSubsysParameter := fassSDK.DeleteSubsysParameter{}
	deleteSubsysParameter.SetSubsysName(subsysName)

	err = fassSDK.DeleteSubsysRequest(&deleteSubsysParameter, uuid.New().String())
	if err != nil {
		return err
	}

	return nil
}

func TestCreateVolume(t *testing.T) {
	err := createVolumeFromSnapshot(newSubsysName1, newVolumeName1)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestListVolume(t *testing.T) {
	parameter := fassSDK.ListVolumeParameter{}

	_, err := fassSDK.ListVolumeRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestRetrieveVolume(t *testing.T) {
	parameter := fassSDK.RetrieveVolumeParameter{}
	parameter.SetVolumeName(newVolumeName1)

	_, err := fassSDK.RetrieveVolumeRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.FailNow()
	}
}

func TestExpandVolume(t *testing.T) {
	parameter := fassSDK.ExpandVolumeParameter{}
	parameter.SetVolumeName(newVolumeName1)
	parameter.SetCapacityGB(20)

	err := fassSDK.ExpandVolumeRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func TestSetQosOfVolume(t *testing.T) {
	parameter := fassSDK.SetQosOfVolumeParameter{}
	parameter.SetVolumeName(newVolumeName1)
	parameter.SetBpsMB(100)
	parameter.SetIops(1000)

	err := fassSDK.SetQosOfVolumeRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func flattenVolume(volumeName string) (*fassSDK.FlattenVolumeResponse, error) {
	parameter := fassSDK.FlattenVolumeParameter{}
	parameter.SetVolumeName(volumeName)

	data, err := fassSDK.FlattenVolumeRequest(&parameter, uuid.New().String())
	return data, err
}

func TestFlattenVolume(t *testing.T) {
	data, err := flattenVolume(newVolumeName1)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}

	taskId = data.TaskId
	fmt.Printf("flatten volume task_id %s\n", taskId)
}

func TestFlattenVolumeProgress(t *testing.T) {
	parameter := fassSDK.GetFlattenVolumeProgress{}
	parameter.SetTaskId(taskId)

	for {
		data, err := fassSDK.FlattenVolumeProgressRequest(&parameter, uuid.New().String())
		if !reflect.DeepEqual(err, nil) {
			fmt.Printf("%s", err.Error())
			t.FailNow()
		}
		if data.IsDone() {
			fmt.Printf("volume flatten complete\n")
			return
		}

		fmt.Printf("wait voluem faltten end\n")
		time.Sleep(2 * time.Second)
	}
}

func TestStopFlattenVolume(t *testing.T) {
	data, err := flattenVolume(newVolumeName2)
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}

	parameter := fassSDK.StopFlattenVolumeParameter{}
	parameter.SetTaskId(data.TaskId)

	err = fassSDK.StopFlattenVolumeRequest(&parameter, uuid.New().String())
	if !reflect.DeepEqual(err, nil) {
		fmt.Printf("%s", err.Error())
		t.Fail()
	}
}

func deleteVolume(volumeName string, retry int) error {
	parameter := fassSDK.DeleteVolumeParameter{}
	parameter.SetVolumeName(volumeName)
	parameter.ForceDelete()

	var err error
	err = fassSDK.DeleteVolumeRequest(&parameter, uuid.New().String())

	if retry == 0 {
		return err
	}

	for i := 1; i < retry; i++ {
		err = fassSDK.DeleteVolumeRequest(&parameter, uuid.New().String())
		if err == nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}

	return err
}

func TestDeleteVolume(t *testing.T) {
	err := deleteVolume(newVolumeName1, 0)
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
