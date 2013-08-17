package component

import (
	"strconv"
	"strings"
	"time"
)

type QueryParameterBuilder struct {
}

type RestrictionEditorFunc func(queryParameter QueryParameter, value string) map[string]interface{}

func (o QueryParameterBuilder) buildQuery(queryParameter QueryParameter, value string) map[string]interface{} {
	// ------------------------------------
	{
		funcMap := map[string]map[string]RestrictionEditorFunc{
			"textfield":     nil,
			"textarea":      nil,
			"numberfield":   nil,
			"datefield":     nil,
			"combo":         nil,
			"combotree":     nil,
			"displayfield":  nil,
			"hidden":        nil,
			"htmleditor":    nil,
			"checkbox":      nil,
			"checkboxgroup": nil,
			"radio":         nil,
			"radiogroup":    nil,
			"trigger":       nil,
		}
		println(funcMap)
	}
	// ------------------------------------

	funcMap := map[string]map[string]RestrictionEditorFunc{
		"eq":             o.eqNqMap("$eq"),
		"nq":             o.eqNqMap("$ne"),
		"ge":             o.logicCmpMap("$gte"),
		"le":             o.logicCmpMap("$lte"),
		"gt":             o.logicCmpMap("$gt"),
		"lt":             o.logicCmpMap("$lt"),
		"null":           o.eqNqMap("$eq"), // don't forget to set value to ""
		"not_null":       o.eqNqMap("$ne"), // don't forget to set value to ""
		"exist":          o.existNotExistCmpMap(true),
		"not_exist":      o.existNotExistCmpMap(false),
		"in":             o.inNotInMap("$in"),
		"not_in":         o.inNotInMap("$nin"),
		"like":           o.regexpCmpMap("$regex"),
		"left_like":      o.regexpCmpMap("$regex"),
		"right_like":     o.regexpCmpMap("$regex"),
		"not_like":       o.regexpCmpMap("$regex"),
		"not_left_like":  o.regexpCmpMap("$regex"),
		"not_right_like": o.regexpCmpMap("$regex"),
	}
	if funcMap[queryParameter.Restriction] != nil {
		if funcMap[queryParameter.Restriction][queryParameter.Editor] != nil {
			var restrictionEditorFunc RestrictionEditorFunc
			restrictionEditorFunc = funcMap[queryParameter.Restriction][queryParameter.Editor]
			return restrictionEditorFunc(queryParameter, value)
		}
	}
	panic(queryParameter.Name + ",restriction is:" + queryParameter.Restriction + ",editor is:" + queryParameter.Editor + ", value is:" + value)
}

func (o QueryParameterBuilder) eqNqMap(operator string) map[string]RestrictionEditorFunc {
	inNotIn := "$in"
	if operator == "$ne" {
		inNotIn = "$nin"
	}
	return map[string]RestrictionEditorFunc{
		"textfield":     o.intOrFloatOrStringCmp(operator),
		"textarea":      o.intOrFloatOrStringCmp(operator),
		"numberfield":   o.intOrFloatOrStringCmp(operator),
		"datefield":     o.dateCmp(operator),
		"combo":         o.intOrStringOrFloatInCmp(inNotIn),
		"combotree":     o.intOrStringOrFloatInCmp(inNotIn),
		"displayfield":  nil,
		"hidden":        o.intOrFloatOrStringCmp(operator),
		"htmleditor":    o.intOrFloatOrStringCmp(operator),
		"checkbox":      o.intOrFloatOrStringCmp(operator),
		"checkboxgroup": o.intOrStringOrFloatInCmp(inNotIn),
		"radio":         o.intOrFloatOrStringCmp(operator),
		"radiogroup":    o.intOrFloatOrStringCmp(operator),
		"trigger":       o.intOrFloatOrStringCmp(operator),
	}
}

func (o QueryParameterBuilder) inNotInMap(operator string) map[string]RestrictionEditorFunc {
	return map[string]RestrictionEditorFunc{
		"textfield":     o.intOrStringOrFloatInCmp(operator),
		"textarea":      o.intOrStringOrFloatInCmp(operator),
		"numberfield":   o.intOrStringOrFloatInCmp(operator),
		"datefield":     nil,
		"combo":         o.intOrStringOrFloatInCmp(operator),
		"combotree":     o.intOrStringOrFloatInCmp(operator),
		"displayfield":  nil,
		"hidden":        o.intOrStringOrFloatInCmp(operator),
		"htmleditor":    o.intOrStringOrFloatInCmp(operator),
		"checkbox":      o.intOrStringOrFloatInCmp(operator),
		"checkboxgroup": o.intOrStringOrFloatInCmp(operator),
		"radio":         o.intOrStringOrFloatInCmp(operator),
		"radiogroup":    o.intOrStringOrFloatInCmp(operator),
		"trigger":       o.intOrStringOrFloatInCmp(operator),
	}
}

func (o QueryParameterBuilder) logicCmpMap(operator string) map[string]RestrictionEditorFunc {
	return map[string]RestrictionEditorFunc{
		"textfield":     o.intOrFloatOrStringCmp(operator),
		"textarea":      nil,
		"numberfield":   o.intOrFloatOrStringCmp(operator),
		"datefield":     o.dateCmp(operator),
		"combo":         o.intOrFloatOrStringCmp(operator),
		"combotree":     nil,
		"displayfield":  nil,
		"hidden":        o.intOrFloatOrStringCmp(operator),
		"htmleditor":    nil,
		"checkbox":      o.intOrFloatOrStringCmp(operator),
		"checkboxgroup": nil,
		"radio":         o.intOrFloatOrStringCmp(operator),
		"radiogroup":    nil,
		"trigger":       o.intOrFloatOrStringCmp(operator),
	}
}

func (o QueryParameterBuilder) existNotExistCmpMap(isExist bool) map[string]RestrictionEditorFunc {
	return map[string]RestrictionEditorFunc{
		"textfield":     o.existNotExistCmp(isExist),
		"textarea":      o.existNotExistCmp(isExist),
		"numberfield":   o.existNotExistCmp(isExist),
		"datefield":     o.existNotExistCmp(isExist),
		"combo":         o.existNotExistCmp(isExist),
		"combotree":     o.existNotExistCmp(isExist),
		"displayfield":  o.existNotExistCmp(isExist),
		"hidden":        o.existNotExistCmp(isExist),
		"htmleditor":    o.existNotExistCmp(isExist),
		"checkbox":      o.existNotExistCmp(isExist),
		"checkboxgroup": o.existNotExistCmp(isExist),
		"radio":         o.existNotExistCmp(isExist),
		"radiogroup":    o.existNotExistCmp(isExist),
		"trigger":       o.existNotExistCmp(isExist),
	}
}

func (o QueryParameterBuilder) regexpCmpMap(operator string) map[string]RestrictionEditorFunc {
	return map[string]RestrictionEditorFunc{
		"textfield":     o.regexpCmp(operator),
		"textarea":      o.regexpCmp(operator),
		"numberfield":   nil,
		"datefield":     nil,
		"combo":         nil,
		"combotree":     nil,
		"displayfield":  nil,
		"hidden":        o.regexpCmp(operator),
		"htmleditor":    o.regexpCmp(operator),
		"checkbox":      nil,
		"checkboxgroup": nil,
		"radio":         nil,
		"radiogroup":    nil,
		"trigger":       nil,
	}
}

func (o QueryParameterBuilder) intOrFloatOrStringCmp(operator string) RestrictionEditorFunc {
	return func(queryParameter QueryParameter, value string) map[string]interface{} {
		return o.intOrFloatOrStringOperator(queryParameter, value, operator)
	}
}

func (o QueryParameterBuilder) intOrFloatOrStringOperator(queryParameter QueryParameter, value string, operator string) map[string]interface{} {
	resultLi := []map[string]interface{}{}
	intValue, err := strconv.ParseInt(value, 10, 0)
	if err == nil {
		if operator == "$eq" {
			resultLi = append(resultLi, map[string]interface{}{
				o.getQueryName(queryParameter): intValue,
			})
		} else {
			resultLi = append(resultLi, map[string]interface{}{
				o.getQueryName(queryParameter): map[string]int{
					operator: int(intValue),
				},
			})
		}
	}

	floatValue, err := strconv.ParseFloat(value, 32)
	if err == nil {
		if operator == "$eq" {
			resultLi = append(resultLi, map[string]interface{}{
				o.getQueryName(queryParameter): float32(floatValue),
			})
		} else {
			resultLi = append(resultLi, map[string]interface{}{
				o.getQueryName(queryParameter): map[string]float32{
					operator: float32(floatValue),
				},
			})
		}
	}

	if operator == "$eq" {
		resultLi = append(resultLi, map[string]interface{}{
			o.getQueryName(queryParameter): value,
		})
	} else {
		resultLi = append(resultLi, map[string]interface{}{
			o.getQueryName(queryParameter): map[string]string{
				operator: value,
			},
		})
	}

	if len(resultLi) == 1 {
		return resultLi[0]
	}
	return map[string]interface{}{
		"$or": resultLi,
	}
}

func (o QueryParameterBuilder) dateCmp(operator string) RestrictionEditorFunc {
	return func(queryParameter QueryParameter, value string) map[string]interface{} {
		return o.dateOperator(queryParameter, value, operator)
	}
}

func (o QueryParameterBuilder) dateOperator(queryParameter QueryParameter, value string, operator string) map[string]interface{} {
	inFormat := "yyyy-MM-dd"
	queryFormat := "yyyyMMdd"
	for _, parameterAttribute := range queryParameter.ParameterAttributeLi {
		if parameterAttribute.Name == "inFormat" {
			inFormat = parameterAttribute.Value
		}
		if parameterAttribute.Name == "queryFormat" {
			queryFormat = parameterAttribute.Value
		}
	}
	t, err := time.Parse(inFormat, value)
	if err != nil {
		panic(err)
	}

	queryDataStr := t.Format(queryFormat)
	queryData, err := strconv.ParseInt(queryDataStr, 10, 0)
	if err != nil {
		panic(err)
	}

	if operator == "$eq" {
		return map[string]interface{}{
			o.getQueryName(queryParameter): int(queryData),
		}
	}
	return map[string]interface{}{
		o.getQueryName(queryParameter): map[string]int{
			operator: int(queryData),
		},
	}
}

func (o QueryParameterBuilder) intOrStringOrFloatInCmp(operator string) RestrictionEditorFunc {
	return func(queryParameter QueryParameter, value string) map[string]interface{} {
		return o.intOrStringOrFloatInOperator(queryParameter, value, operator)
	}
}

func (o QueryParameterBuilder) intOrStringOrFloatInOperator(queryParameter QueryParameter, value string, operator string) map[string]interface{} {
	valueLi := strings.Split(value, ",")
	resultLi := []map[string]interface{}{}

	resultLi = append(resultLi, map[string]interface{}{
		o.getQueryName(queryParameter): map[string]interface{}{
			operator: valueLi,
		},
	})

	intValueLi := []int{}
	for _, valueItem := range valueLi {
		intValue, err := strconv.ParseInt(valueItem, 10, 0)
		if err == nil {
			intValueLi = append(intValueLi, int(intValue))
		}
	}

	if len(intValueLi) > 0 {
		resultLi = append(resultLi, map[string]interface{}{
			o.getQueryName(queryParameter): map[string]interface{}{
				operator: intValueLi,
			},
		})
	}

	floatValueLi := []float32{}
	for _, valueItem := range valueLi {
		float32Value, err := strconv.ParseFloat(valueItem, 32)
		if err == nil {
			floatValueLi = append(floatValueLi, float32(float32Value))
		}
	}

	if len(floatValueLi) > 0 {
		resultLi = append(resultLi, map[string]interface{}{
			o.getQueryName(queryParameter): map[string]interface{}{
				operator: floatValueLi,
			},
		})
	}

	if len(resultLi) == 1 {
		return resultLi[0]
	}
	return map[string]interface{}{
		"$or": resultLi,
	}
}

func (o QueryParameterBuilder) existNotExistCmp(isExist bool) RestrictionEditorFunc {
	return func(queryParameter QueryParameter, value string) map[string]interface{} {
		return o.existNotExistOperator(queryParameter, isExist)
	}
}

func (o QueryParameterBuilder) existNotExistOperator(queryParameter QueryParameter, isExist bool) map[string]interface{} {
	return map[string]interface{}{
		o.getQueryName(queryParameter): map[string]bool{
			"$exists": isExist,
		},
	}
}

func (o QueryParameterBuilder) regexpCmp(operator string) RestrictionEditorFunc {
	return func(queryParameter QueryParameter, value string) map[string]interface{} {
		return o.regexpOperator(queryParameter, value, operator)
	}
}

func (o QueryParameterBuilder) regexpOperator(queryParameter QueryParameter, value string, operator string) map[string]interface{} {
	regex := ""
	switch queryParameter.Restriction {
	case "like", "not_like":
		regex = "^.*?" + value + ".*?$"
	case "left_like", "not_left_like":
		regex = "^" + value + ".*?$"
	case "right_like", "not_right_like":
		regex = "^.*?" + value + "$"
	}
	switch queryParameter.Restriction {
	case "like", "left_like", "right_like":
		return map[string]interface{}{
			o.getQueryName(queryParameter): map[string]string{
				operator: regex,
			},
		}
	case "not_like", "not_left_like", "not_right_like":
		return map[string]interface{}{
			o.getQueryName(queryParameter): map[string]map[string]string{
				"$not": map[string]string{
					operator: regex,
				},
			},
		}
	}
	panic("input queryParameter.Editor:" + queryParameter.Editor + ", must be one of 'like', 'left_like', 'right_like', 'not_like', 'not_left_like', 'not_left_like'")
}

func (o QueryParameterBuilder) getQueryName(queryParameter QueryParameter) string {
	if queryParameter.ColumnName != "" {
		return queryParameter.ColumnName
	}
	return queryParameter.Name
}
