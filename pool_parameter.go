package fass_sdk_golang

// File       : pool.go
// Path       : parameters
// Time       : CST 2023/4/25 14:59
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"encoding/json"
	"fmt"

	"github.com/Vanishvei/fass-sdk-golang/horizontal"
)

type ListPoolParameter = listParameter

type RetrievePoolParameter struct {
	poolName *string
}

func (parameter *RetrievePoolParameter) SetPoolName(poolName string) {
	parameter.poolName = horizontal.String(poolName)
}

func (parameter *RetrievePoolParameter) GetPath() string {
	if parameter.poolName == nil {
		panic("parameter poolName not set")
	}
	return fmt.Sprintf("pool/%s", *parameter.poolName)
}

type DeletePoolParameter struct {
	poolName *string
	isForce  *bool
}

func (parameter *DeletePoolParameter) SetPoolName(poolName string) {
	parameter.poolName = horizontal.String(poolName)
}

func (parameter *DeletePoolParameter) ForceDelete() {
	parameter.isForce = horizontal.Bool(true)
}

func (parameter DeletePoolParameter) GetPath() string {
	if parameter.poolName == nil {
		panic("parameter poolName not set")
	}
	return fmt.Sprintf("pool/%s", *parameter.poolName)
}

func (parameter DeletePoolParameter) GetQuery() map[string]*string {
	if parameter.isForce == nil {
		return map[string]*string{"is_force": horizontal.String("false")}
	}

	return map[string]*string{"is_force": horizontal.String("true")}
}

type CreatePoolParameter struct {
	ecRatio        *string
	poolName       *string
	redundancyType *string
	replicaNum     *int
	sectorSize     *int
}

func (parameter CreatePoolParameter) MarshalJSON() ([]byte, error) {
	_map := map[string]interface{}{
		"pool_name":       *parameter.poolName,
		"sector_size":     *parameter.sectorSize,
		"redundancy_type": *parameter.redundancyType,
	}
	if parameter.ecRatio != nil {
		_map["ec_ratio"] = parameter.ecRatio
	}
	if parameter.replicaNum != nil {
		_map["replica_num"] = parameter.replicaNum
	}
	return json.Marshal(_map)
}

func (parameter *CreatePoolParameter) SetPoolName(poolName string) {
	parameter.poolName = horizontal.String(poolName)
}

func (parameter *CreatePoolParameter) SetReplicaNum2() {
	parameter.redundancyType = horizontal.String("replication")
	parameter.replicaNum = horizontal.Int(2)
}

func (parameter *CreatePoolParameter) SetReplicaNum3() {
	parameter.redundancyType = horizontal.String("replication")
	parameter.replicaNum = horizontal.Int(3)
}

func (parameter *CreatePoolParameter) SetReplicaNum4() {
	parameter.redundancyType = horizontal.String("replication")
	parameter.replicaNum = horizontal.Int(4)
}

func (parameter *CreatePoolParameter) SetReplicaNum5() {
	parameter.redundancyType = horizontal.String("replication")
	parameter.replicaNum = horizontal.Int(5)
}

func (parameter *CreatePoolParameter) SetECRatio_4_2() {
	parameter.redundancyType = horizontal.String("ec")
	parameter.ecRatio = horizontal.String("4+2")
}

func (parameter *CreatePoolParameter) SetECRatio_6_2() {
	parameter.redundancyType = horizontal.String("ec")
	parameter.ecRatio = horizontal.String("6+2")
}

func (parameter *CreatePoolParameter) SetECRatio_8_2() {
	parameter.redundancyType = horizontal.String("ec")
	parameter.ecRatio = horizontal.String("8+2")
}

func (parameter *CreatePoolParameter) SetSectorSize512() {
	parameter.sectorSize = horizontal.Int(512)
}

func (parameter *CreatePoolParameter) SetSectorSize4096() {
	parameter.sectorSize = horizontal.Int(4096)
}
