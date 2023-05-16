package fass_sdk_golang

// File       : common.go
// Path       : parameters
// Time       : CST 2023/4/26 14:15
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"github.com/Vanishvei/fass-sdk-golang/horizontal"
	"strconv"
)

type listParameter struct {
	pageSize *int
	pageNum  *int
}

func (parameter *listParameter) SetPageSize(pageSize int) {
	parameter.pageSize = horizontal.Int(pageSize)
}

func (parameter *listParameter) SetPageNum(pageNum int) {
	parameter.pageNum = horizontal.Int(pageNum)
}

func (parameter *listParameter) GetQuery() map[string]*string {
	defaultPageNum := 1
	defaultPageSize := 20
	pageNum := strconv.Itoa(*horizontal.DefaultNumber(parameter.pageNum, horizontal.Int(defaultPageNum)))
	pageSize := strconv.Itoa(*horizontal.DefaultNumber(parameter.pageSize, horizontal.Int(defaultPageSize)))
	return map[string]*string{
		"page_num":  horizontal.String(pageNum),
		"page_size": horizontal.String(pageSize),
	}
}
