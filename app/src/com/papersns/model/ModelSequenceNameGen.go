package model

import (
	"strings"
)

func GetMasterSequenceName(dataSource DataSource) string {
	byte0 := dataSource.Id[0]
	return strings.ToLower(string(byte0)) + dataSource.Id[1:] + "Id"
}

func GetDetailSequenceName(dataSource DataSource, detailData DetailData) string {
	return GetMasterSequenceName(dataSource)
//	byte0 := dataSource.Id[0]
//	return strings.ToLower(string(byte0)) + dataSource.Id[1:] + detailData.Id + "Id"
}
