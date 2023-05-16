# Development Kit for golang

```text
Fass-sdk-golang is the Taocloud block storage service (FASS) Software Development Kit
for golang, which work with golang 1.18 or above.
```

----
## Object description
**The relationship between the request and parameter of the resource operation depends on the reference**

<table>
    <tr>
        <td>Object</td>
        <td>Operation</td>
        <td>Request</td>
        <td>Parameter</td>
        <td>Result</td>
    </tr>
    <tr>
        <td rowspan="4">Account</td>
        <td>Create</td>
        <td>CreateAccountRequest</td>
        <td>CreateAccountParameter</td>
        <td>CreateAccountResponse, error</td>
    </tr>
    <tr>
        <td>Describe</td>
        <td>RetrieveAccountRequest</td>
        <td>RetrieveAccountParameter</td>
        <td>RetrieveAccountResponse, error</td>
    </tr>
    <tr>
        <td>List</td>
        <td>ListAccountRequest</td>
        <td>ListAccountParameter</td>
        <td>ListAccountResponse, error</td>
    </tr>
    <tr>
        <td>Delete</td>
        <td>DeleteAccountRequest</td>
        <td>DeleteAccountParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td rowspan="5">Group</td>
        <td>AddQualifier</td>
        <td>AddQualifierToGroupRequest</td>
        <td>AddQualifierToGroupParameter</td>
        <td>AddQualifierToGroupResponse, error</td>
    </tr>
    <tr>
        <td>RemoveQualifier</td>
        <td>RemoveQualifierFromGroupRequest</td>
        <td>RemoveQualifierFromGroupParameter</td>
        <td>RemoveQualifierFromGroupResponse, error</td>
    </tr>
    <tr>
        <td>Describe</td>
        <td>RetrieveGroupRequest</td>
        <td>RetrieveGroupParameter</td>
        <td>RetrieveGroupResponse, error</td>
    </tr>
    <tr>
        <td>List</td>
        <td>ListGroupRequest</td>
        <td>ListGroupParameter</td>
        <td>ListGroupResponse, error</td>
    </tr>
    <tr>
        <td>Delete</td>
        <td>DeleteGroupRequest</td>
        <td>DeleteGroupParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td rowspan="4">Pool</td>
        <td>Create</td>
        <td>CreatePoolRequest</td>
        <td>CreatePoolParameter</td>
        <td>CreatePoolResponse, error</td>
    </tr>
    <tr>
        <td>Describe</td>
        <td>RetrievePoolRequest</td>
        <td>RetrievePoolParameter</td>
        <td>RetrievePoolResponse, error</td>
    </tr>
    <tr>
        <td>List</td>
        <td>ListPoolRequest</td>
        <td>ListPoolParameter</td>
        <td>ListPoolResponse, error</td>
    </tr>
    <tr>
        <td>Delete</td>
        <td>DeletePoolRequest</td>
        <td>DeletePoolParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td rowspan="12">Subsys</td>
        <td>Create</td>
        <td>CreateSubsysRequest</td>
        <td>CreateSubsysParameter</td>
        <td>CreateSubsysResponse, error</td>
    </tr>
    <tr>
        <td>Describe</td>
        <td>RetrieveSubsysRequest</td>
        <td>RetrieveSubsysParameter</td>
        <td>RetrieveSubsysResponse, error</td>
    </tr>
    <tr>
        <td>List</td>
        <td>ListSubsysRequest</td>
        <td>ListSubsysParameter</td>
        <td>ListSubsysResponse, error</td>
    </tr>
    <tr>
        <td>Export</td>
        <td>ExportSubsysRequest</td>
        <td>ExportSubsysParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td>Unexport</td>
        <td>UnexportSubsysRequest</td>
        <td>UnexportSubsysParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td>SetAuth</td>
        <td>SetSubsysAuthRequest</td>
        <td>SetSubsysAuthParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td>AuthDescribe</td>
        <td>RetrieveSubsysAuthRequest</td>
        <td>RetrieveSubsysAuthParameter</td>
        <td>RetrieveSubsysAuthResponse, error</td>
    </tr>
    <tr>
        <td>RemoveAuth</td>
        <td>RemoveSubsysAuthRequest</td>
        <td>RemoveSubsysAuthParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td>SetChap</td>
        <td>SetSubsysChapRequest</td>
        <td>SetSubsysChapParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td>ChapDescribe</td>
        <td>RetrieveSubsysChapRequest</td>
        <td>RetrieveSubsysChapParameter</td>
        <td>RetrieveSubsysChapResponse, error</td>
    </tr>
    <tr>
        <td>RemoveChap</td>
        <td>RemoveSubsysChapRequest</td>
        <td>RemoveSubsysChapParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td>Delete</td>
        <td>DeleteSubsysRequest</td>
        <td>DeleteSubsysParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td rowspan="8">Volume</td>
        <td>Describe</td>
        <td>RetrieveVolumeRequest</td>
        <td>RetrieveVolumeParameter</td>
        <td>RetrieveVolumeResponse, error</td>
    </tr>
    <tr>
        <td>List</td>
        <td>ListVolumeRequest</td>
        <td>ListVolumeParameter</td>
        <td>ListVolumeResponse, error</td>
    </tr>
    <tr>
        <td>Flatten</td>
        <td>FlattenVolumeRequest</td>
        <td>FlattenVolumeParameter</td>
        <td>FlattenVolumeResponse, error</td>
    </tr>
    <tr>
        <td>FlattenProgress</td>
        <td>FlattenVolumeProgressRequest</td>
        <td>FlattenVolumeProgressParameter</td>
        <td>FlattenVolumeProgressResponse, error</td>
    </tr>
    <tr>
        <td>StopFlatten</td>
        <td>StopFlattenVolumeRequest</td>
        <td>StopFlattenVolumeParameter</td>
        <td>StopFlattenVolumeResponse, error</td>
    </tr>
    <tr>
        <td>SetQos</td>
        <td>SetQosOfVolumeRequest</td>
        <td>SetQosOfVolumeParameter</td>
        <td>SetQosOfVolumeResponse, error</td>
    </tr>
    <tr>
        <td>Expand</td>
        <td>ExpandVolumeRequest</td>
        <td>ExpandVolumeParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td>Delete</td>
        <td>DeleteVolumeRequest</td>
        <td>DeleteVolumeParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td rowspan="5">Snapshot</td>
        <td>Create</td>
        <td>CreateSnapshotRequest</td>
        <td>CreateSnapshotParameter</td>
        <td>CreateSnapshotResponse, error</td>
    </tr>
    <tr>
        <td>Describe</td>
        <td>RetrieveSnapshotRequest</td>
        <td>RetrieveSnapshotParameter</td>
        <td>RetrieveSnapshotResponse, error</td>
    </tr>
    <tr>
        <td>List</td>
        <td>ListSnapshotRequest</td>
        <td>ListSnapshotParameter</td>
        <td>ListSnapshotResponse, error</td>
    </tr>
    <tr>
        <td>Revert</td>
        <td>RevertSnapshotRequest</td>
        <td>RevertSnapshotParameter</td>
        <td>error</td>
    </tr>
    <tr>
        <td>Delete</td>
        <td>DeleteSnapshotRequest</td>
        <td>DeleteSnapshotParameter</td>
        <td>error</td>
    </tr>
</table>

**SDKError contains the key information**
- RequestId :  request id
- StatusCode:  http request status code
- Code      :  fass API operation result status code
- Message   :  fass API operation failed description message
- Data      :  empty can be ignored

Example
```text
RequestId: b225b6e3-50d2-493b-9219-8443d2e0389b
StatusCode: 0
Code: 400003
Message: Pool pool3 Not Exist
Data: {}
```

----
## Example
*All operations need to be initialized before the client, it is recommended to put it in the init method to handle*
```go
import "fass-go/client"

...
defer func() {
if err := recover(); err != nil {
	// custom panic handle
    }
}()

endpointList := ["xxx.xxx.xxx.xxx", "xxx.xxx.xxx.xxx", ...]
port := 8000
backoff := 5
retryCount := 3
readTimeout := 30
connectTimeout := 1

client.InitConfig(&endpointList, &port, &readTimeout, &connectTimeout, &backoff, &retryCount)
...
```

### Pool

CreatePool
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.CreatePoolParameter{}
parameter.SetECRatio_4_2()
parameter.SetSectorSize4096()
parameter.SetPoolName(poolName)
result, err := requests.CreatePoolRequest(&parameter, requestId)
......
```

RetrievePool
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.RetrievePoolParameter{}
parameter.SetPoolName("fast_pool") // 设置要获取的存储池信息的池名
result, err := requests.RetrievePoolRequest(&parameter, requestId)
......
```

ListPool
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.ListPoolParameter{}
parameter.SetPageNum(1)// 设置第1页,可以不设置
parameter.SetPageSize(10)// 设置分页大小,可以不设置
result, err := requests.RetrievePoolRequest(&parameter, requestId)
......
```

DeletePool
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.DeletePoolParameter{}
parameter.SetPoolName("fast_pool") // 设置要删除的存储池池名为fast_pool
parameter.ForceDelete() // 强制删除
err := requests.DeletePoolRequest(&parameter, requestId)
......
```
### Subsys

CreateSubsys
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.CreateSubsysParameter{}
parameter.SetName("subsys001") //设置要创建的subsys名为subsys001(名称长度小于53个字符长度)
parameter.SetPoolName("fast_pool") //设置目标存储池为fast_pool(存储池fast_pool必须存在)
parameter.SetVolumeName("volume001") // 设置新卷名称为volume001(名称长度小于64个字符长度),可以不设置该卷默认卷名是subsys_subsys001_0001
parameter.SetSectorSize4096()// 设置新卷的SectorSize为4096,可以不设置默认使用目标池的配置
parameter.SetCapacityGB(10) // 设置新卷的容量为10GB
parameter.EnableISCSI()//启用iSCSI协议(可以不选但是iSCSI和NVMeoF不能都不启用)
parameter.EnableNVMeoF()//启用NVMeoF协议(可以不选但是iSCSI和NVMeoF不能都不启用)
parameter.SetBpsMB(1000) //设置新卷流控的IOPS为1000,可以不设置。默认不限制流控
parameter.SetIops(2000)// 设置新卷流控的BPS为2000,可以不设置。默认不限制流控
parameter.SetFormatROW()// 设置新卷的格式为ROW,可以不设置。默认使用池的配置

result, err := requests.CreateSubsysRequest(&parameter, requestId) // 创建subsys
......
```

CreateVolumeFromSnapshot
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
createSubsysParameter := parameters.CreateSubsysParameter{}
createSubsysParameter.SetPoolName("fast_pool")// 设置目标存储池为fast_pool
createSubsysParameter.SetName("subsys002") // 设置要创建的subsys名为subsys002
createSubsysParameter.SetVolumeName("volume002") // 设置新卷名称为volume002
createSubsysParameter.SetSrcVolumeName("volume001") // 设置源卷的卷名为volume001(使用快照建卷时源卷必须存在)
createSubsysParameter.SetSrcSnapshotName("snapshot001") //设置源卷的快照名为snapshot001(使用快照建卷时源卷的快照必须存在)
createSubsysParameter.EnableISCSI()// 新subsys启用iSCSI协议

result, err := requests.CreateSubsysRequest(&createSubsysParameter, requestId)
......
```

RetrieveSubsys
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()
requestId := uuid.New().String()
parameter := parameters.RetrieveSubsysParameter{}
parameter.SetSubsysName("subsys001") // 设置subsys名称为subsys001

result, err := requests.RetrieveSubsysRequest(&parameter, requestId)
......
```

ListSubsys
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()
requestId := uuid.New().String()
parameter := parameters.ListPoolParameter{}
parameter.SetPageNum(1)// 设置第1页,可以不设置
parameter.SetPageSize(10)// 设置分页大小,可以不设置
result, err := requests.ListPoolRequest(&parameter, requestId)
......
```

ExportSubsys
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()
requestId := uuid.New().String()
parameter := parameters.ExportSubsysParameter{}
parameter.SetSubsysName("subsys001") // 设置subsys名称为subsys001
parameter.ExportISCSI() // 导出iSCSI协议

err := requests.ExportSubsysRequest(&parameter, requestId) //导出协议
......
```

UnexportSubsys
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()
requestId := uuid.New().String()

parameter := parameters.UnexportSubsysParameter{}
parameter.SetSubsysName("subsys001")// 设置subsys名称为subsys001
parameter.UnexportISCSI() // 关闭iSCSI协议的导出

err := requests.UnexportSubsysRequest(&parameter, requestId) 
......
```

DeleteSubsys
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()
requestId := uuid.New().String()
parameter := parameters.DeleteSubsysParameter{}
parameter.SetSubsysName("subsys001")
parameter.ForceDelete() // 使用强制删除,可以不设置。默认不强制删除
parameter.DeleteVolume() // 一起删除关联的卷,可以不设置。默认不删除关联的卷
err := requests.DeletePoolRequest(&parameter, requestId)
......
```

### Volume

RetrieveVolume
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.RetrieveVolumeParameter{}
parameter.SetVolumeName("volume001") //设置要获取的卷信息的卷名为volume001
result, err := requests.RetrieveVolumeRequest(&parameter, requestId)
......
```

ExpandVolume
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.ExpandVolumeParameter{}
parameter.SetVolumeName("volume001") //设置要扩容的卷的卷名为volume001(扩容操作的卷volume001必须存在)
parameter.SetCapacityGB(20) //设置卷要扩容的大小为20GB
err := requests.ExpandVolumeRequest(&parameter, requestId)
......
```

SetQosOfVolume
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.SetQosOfVolumeParameter{}
parameter.SetVolumeName("volume002") // 设置要设置流控的卷的卷名为volume002(设置流控的卷必须存在)
parameter.SetBpsMB(100)  // 设置bps限制为100MB/s
parameter.SetIops(1000) // 设置iops限制为1000/s

err := requests.SetQosOfVolumeRequest(&parameter, requestId)
......
```

FlattenVolume
**只有使用快照创建出来的卷才能做分离操**
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.FlattenVolumeParameter{}
parameter.SetVolumeName("volume002") // 设置要分离的卷的卷名为volume002(分离卷volume002必须存在)

data, err := requests.FlattenVolumeRequest(&parameter, requestId)
......
```

FlattenVolumeProgress
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.GetFlattenVolumeProgress{}
parameter.SetTaskId("d269f2ba-d3cf-4aa5-870a-667268e4a414") // 设置分离任务的任务id为d269f2ba-d3cf-4aa5-870a-667268e4a414(从FlattenVolume的结果中获取)
result, err := requests.FlattenVolumeProgressRequest(&parameter, requestId)
......
```

StopFlattenVolume
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.StopFlattenVolumeRequest{}
parameter.SetTaskId("d269f2ba-d3cf-4aa5-870a-667268e4a414") //设置要停止flatten任务的任务id为d269f2ba-d3cf-4aa5-870a-667268e4a414(从FlattenVolume的结果中获取)

err := requests.StopFlattenVolumeRequest(&parameter, requestId)
......
```


DeleteVolume
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.DeleteVolumeParameter{}
parameter.SetVolumeName("volume002") // 设置要分离的卷的卷名为volume002
parameter.ForceDelete() // 强制删除,可以不强制删除

err := requests.DeleteVolumeRequest(&parameter, requestId)
......
```
### Snapshot
CreateSnapshot
**一个卷最多能够创建256个快照**
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.CreateSnapshotParameter{}
parameter.SetVolumeName("volume001") // 设置要创建快照的卷的名称为volume001(卷volume001必须存在)
parameter.SetSnapshotName("snapshot001") // 设置快照名称为snapshot001
result, err := requests.CreateSnapshotRequest(&parameter, requestId)
......
```

RetrieveSnapshot
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.RevertSnapshotParameter{}
parameter.SetVolumeName("volume001") 
parameter.SetSnapshotName("snapshot001")
err := requests.RevertSnapshotRequest(&parameter, requestId)
......
```

ListSnapshot
**获取卷的快照列表**
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.ListSnapshotParameter{}
parameter.SetVolumeName("volume001") // 设置要获取的快照列表所属的卷的卷名为volume001
result, err := requests.ListSnapshotRequest(&parameter, requestId)
......
```

RevertSnapshot
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.RevertSnapshotParameter{}
parameter.SetVolumeName("volume001")
parameter.SetSnapshotName("snapshot001")
err := requests.RevertSnapshotRequest(&parameter, requestId)
......
```

DeleteSnapshot
```go
......
defer func() {
if err := recover(); err != nil {
// custom panic handle
    }
}()

requestId := uuid.New().String()
parameter := parameters.DeleteSnapshotParameter{}
parameter.SetVolumeName("volume001")
parameter.SetSnapshotName("snapshot001")
err := requests.DeleteSnapshotRequest(&parameter, requestId)
......
```
----

## Parameter description

- **Pool**
<table>
    <tr>
        <td>ParameterStruct</td>
        <td>Method</td>
        <td>Required</td>
        <td>Description</td>
    </tr>
    <tr>
        <td rowspan="4">CreatePoolParameter</td>
        <td>SetPoolName</td>
        <td>yes</td>
        <td>Set storage pool name</td>
    </tr>
    <tr>
        <td>SetReplicaNum[2|3|4|5]</td>
        <td>yes</td>
        <td>Set the redundancy type to [2|3|4|5] replicas</td>
    </tr>
    <tr>
        <td>SetECRatio_[4|6|8]_2</td>
        <td>yse</td>
        <td>Set the redundancy type to ec [4|6|8]+2</td>
    </tr>
    <tr>
        <td>SetSectorSize[512|4096]</td>
        <td>yes</td>
        <td>Set the sector size to [512|4096]</td>
    </tr>
    <tr>
        <td rowspan="1">RetrievePoolParameter</td>
        <td>SetPoolName</td>
        <td>yes</td>
        <td>Set storage pool name</td>
    </tr>
    <tr>
        <td rowspan="2">DeletePoolParameter</td>
        <td>SetPoolName</td>
        <td>yes</td>
        <td>Set storage pool name</td>
    </tr>
    <tr>
        <td>ForceDelete</td>
        <td>no</td>
        <td>Force delete storage pool</td>
    </tr>
</table>

- **ACL**
<table>
    <tr>
        <td>ParameterStruct</td>
        <td>Method</td>
        <td>Required</td>
        <td>Description</td>
    </tr>
    <tr>
        <td rowspan="2">CreateAccountParameter</td>
        <td>SetAccountName</td>
        <td>yes</td>
        <td>Set account name</td>
    </tr>
    <tr>
        <td>SetPassword</td>
        <td>yes</td>
        <td>Set account password</td>
    </tr>
    <tr>
        <td rowspan="1">RetrieveAccountParameter</td>
        <td>SetAccountName</td>
        <td>yes</td>
        <td>Set account name</td>
    </tr>
    <tr>
        <td rowspan="1">DeleteAccountParameter</td>
        <td>SetAccountName</td>
        <td>yes</td>
        <td>Set account name</td>
    </tr>
    <tr>
        <td rowspan="3">AddQualifierToGroupParameter</td>
        <td>SetGroupName</td>
        <td>yes</td>
        <td>Set group name(Automatically created when group does not exist)</td>
    </tr>
    <tr>
        <td>SetHostName</td>
        <td>yes</td>
        <td>Set client host</td>
    </tr>
    <tr>
        <td>SetQualifierList</td>
        <td>yes</td>
        <td>Set qualifiers to of client host</td>
    </tr>
    <tr>
        <td rowspan="3">RemoveQualifierFromGroupParameter</td>
        <td>SetGroupName</td>
        <td>yes</td>
        <td>Set group name</td>
    </tr>
    <tr>
        <td>SetHostName</td>
        <td>yes</td>
        <td>Set client host</td>
    </tr>
    <tr>
        <td>SetQualifierList</td>
        <td>yes</td>
        <td>Set qualifiers removed from client host</td>
    </tr>
    
</table>

- **Subsys**

<table>
    <tr>
        <td>ParameterStruct</td>
        <td>Method</td>
        <td>Required</td>
        <td>Description</td>
    </tr>
    <tr>
        <td rowspan="17">CreateSubsysParameter</td>
        <td>SetPoolName</td>
        <td>yes</td>
        <td>Set storage pool name</td>
    </tr>
    <tr>
        <td>SetName</td>
        <td>yes</td>
        <td>Set the subsys name</td>
    </tr>
    <tr>
        <td>SetPoolName</td>
        <td>yes</td>
        <td>Set destination pool name</td>
    </tr>
    <tr>
        <td>SetBpsMB</td>
        <td>no</td>
        <td>Set volume bps(unit MB)</td>
    </tr>
    <tr>
        <td>SetIops</td>
        <td>no</td>
        <td>Set volume iops</td>
    </tr>
    <tr>
        <td>SetCapacityGB</td>
        <td>no</td>
        <td>Set volume capacity(unit GB, SetCapacityGB and SetCapacityTB are not allowed both are not set)</td>
    </tr>
    <tr>
        <td>SetCapacityTB</td>
        <td>no</td>
        <td>Set volume capacity(unit TB, SetCapacityGB and SetCapacityTB are not allowed both are not set)</td>
    </tr>
    <tr>
        <td>SetFormatROW</td>
        <td>no</td>
        <td>Set volume format ROW</td>
    </tr>
    <tr>
        <td>SetFormatRAW</td>
        <td>no</td>
        <td>Set volume format RAW</td>
    </tr>
    <tr>
        <td>InheritQos(Using the snapshot created)</td>
        <td>no</td>
        <td>Inherit source volume Qos config</td>
    </tr>
    <tr>
        <td>EnableISCSI</td>
        <td>no</td>
        <td>subsys enable iSCSI protocol(EnableISCSI and EnableNVMeoF are not allowed both are not set)</td>
    </tr>
    <tr>
        <td>EnableNVMeoF</td>
        <td>no</td>
        <td>subsys enable NVMeoF protocol(EnableISCSI and EnableNVMeoF are not allowed both are not set)</td>
    </tr>
    <tr>
        <td>SetSectorSize[512|4096]</td>
        <td>no</td>
        <td>Set subsys sector size to [512|4096]</td>
    </tr>
    <tr>
        <td>SetSharding</td>
        <td>no</td>
        <td>Set volume sharding size</td>
    </tr>
    <tr>
        <td>SetVolumeName</td>
        <td>no</td>
        <td>Set volume name(There is no specified for subsys_{subsysName}_{index})</td>
    </tr>
    <tr>
        <td>SetSrcVolumeName</td>
        <td>no</td>
        <td>Set source volume name(Using the snapshot created)</td>
    </tr>
    <tr>
        <td>SetSrcSnapshotName</td>
        <td>no</td>
        <td>Set source volume snapshot name(Using the snapshot created)</td>
    </tr>
    <tr>
        <td rowspan="1">RetrieveSubsysParameter</td>
        <td>SetSubsysName</td>
        <td>yes</td>
        <td>Set subsys name</td>
    </tr>
    <tr>
        <td rowspan="3">ExportSubsysParameter</td>
        <td>SetSubsysName</td>
        <td>yes</td>
        <td>Set subsys name</td>
    </tr>
    <tr>
        <td>ExportISCSI</td>
        <td>no</td>
        <td>subsys export iSCSI(ExportISCSI and ExportNVMeoF are not allowed both are not set)</td>
    </tr>
    <tr>
        <td>ExportNVMeoF</td>
        <td>no</td>
        <td>subsys export NVMeoF(ExportISCSI and ExportNVMeoF are not allowed both are not set)</td>
    </tr>
    <tr>
        <td rowspan="3">UnexportSubsysParameter</td>
        <td>SetSubsysName</td>
        <td>yes</td>
        <td>Set subsys name</td>
    </tr>
    <tr>
        <td>UnexportISCSI</td>
        <td>no</td>
        <td>Subsys unexport iSCSI(UnexportISCSI and UnexportNVMeoF are not allowed both are not set)</td>
    </tr>
    <tr>
        <td>UnexportNVMeoF</td>
        <td>no</td>
        <td>Subsys unexport NVMeoF(UnexportISCSI and UnexportNVMeoF are not allowed both are not set)</td>
    </tr>
    <tr>
        <td rowspan="3">DeleteSubsysParameter</td>
        <td>SetSubsysName</td>
        <td>yes</td>
        <td>Set subsys name</td>
    </tr>
    <tr>
        <td>ForceDelete</td>
        <td>no</td>
        <td>Force delete subsys</td>
    </tr>
    <tr>
        <td>DeleteVolume</td>
        <td>no</td>
        <td>Delete volume when deleting subsys</td>
    </tr>
    <tr>
        <td rowspan="2">SetSubsysAuthParameter</td>
        <td>SetSubsysName</td>
        <td>yes</td>
        <td>Set subsys name</td>
    </tr>
    <tr>
        <td>SetGroupName</td>
        <td>yes</td>
        <td>Set group name</td>
    </tr>
    <tr>
        <td rowspan="1">RetrieveSubsysAuthParameter</td>
        <td>SetSubsysName</td>
        <td>yes</td>
        <td>Set subsys name</td>
    </tr>
    <tr>
        <td rowspan="1">RemoveSubsysAuthParameter</td>
        <td>SetSubsysName</td>
        <td>yes</td>
        <td>Set subsys name</td>
    </tr>
    <tr>
        <td rowspan="2">SetSubsysChapParameter</td>
        <td>SetSubsysName</td>
        <td>yes</td>
        <td>Set subsys name</td>
    </tr>
    <tr>
        <td>SetAccountName</td>
        <td>yes</td>
        <td>Set account name</td>
    </tr>
    <tr>
        <td rowspan="1">RemoveSubsysChapParameter</td>
        <td>SetSubsysName</td>
        <td>yes</td>
        <td>Set subsys name</td>
    </tr>
</table>


- **Volume**
<table>
    <tr>
        <td>ParameterStruct</td>
        <td>Method</td>
        <td>Required</td>
        <td>Description</td>
    </tr>
    <tr>
        <td rowspan="1">RetrieveVolumeParameter</td>
        <td>SetVolumeName</td>
        <td>yes</td>
        <td>Set volume name</td>
    </tr>
    <tr>
        <td rowspan="3">ExpandVolumeParameter</td>
        <td>SetVolumeName</td>
        <td>yes</td>
        <td>Set volume name</td>
    </tr>
    <tr>
        <td>SetCapacityGB</td>
        <td>no</td>
        <td>Expand volume capacity(unit GB, SetCapacityGB and SetCapacityTB are not allowed both are not set)</td>
    </tr>
    <tr>
        <td>SetCapacityTB</td>
        <td>no</td>
        <td>Expand volume capacity(unit TB, SetCapacityGB and SetCapacityTB are not allowed both are not set)</td>
    </tr>
    <tr>
        <td rowspan="1">FlattenVolumeParameter</td>
        <td>SetVolumeName</td>
        <td>yes</td>
        <td>Set volume name</td>
    </tr>
    <tr>
        <td rowspan="1">GetFlattenVolumeProgress</td>
        <td>SetTaskId</td>
        <td>yes</td>
        <td>Set volume flatten task id</td>
    </tr>
    <tr>
        <td rowspan="1">StopFlattenVolumeRequest</td>
        <td>SetTaskId</td>
        <td>yes</td>
        <td>Set volume flatten task id</td>
    </tr>
    <tr>
        <td rowspan="3">SetQosOfVolumeParameter</td>
        <td>SetVolumeName</td>
        <td>yes</td>
        <td>Set volume name</td>
    </tr>
    <tr>
        <td>SetIops</td>
        <td>yes</td>
        <td>Set volume iops</td>
    </tr>
    <tr>
        <td>SetBpsMB</td>
        <td>yes</td>
        <td>Set volume bps(unit MB)</td>
    </tr>
    <tr>
        <td rowspan="2">DeleteVolumeParameter</td>
        <td>SetVolumeName</td>
        <td>yes</td>
        <td>Set volume name</td>
    </tr>
    <tr>
        <td>ForceDelete</td>
        <td>no</td>
        <td>Force delete volume</td>
    </tr>
</table>
