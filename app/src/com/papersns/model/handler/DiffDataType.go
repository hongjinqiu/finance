package handler

import (
	. "com/papersns/model"
)

//ADD("新增"),BEFORE_UPDATE("修改前"),AFTER_UPDATE("修改后"),DELETE("删除")
const ADD int = 1
const BEFORE_UPDATE int = 2
const AFTER_UPDATE int = 3
const DELETE int = 4

type DiffDataRow struct {
	FieldGroupLi []FieldGroup
	DestBo *map[string]interface{}
	DestData *map[string]interface{}
	SrcData map[string]interface{}
	SrcBo map[string]interface{}
}
