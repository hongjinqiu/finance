package common

import (
	"regexp"
	"strconv"
	"fmt"
)

type CommonUtil struct{}

func (o CommonUtil) FilterJsonEmptyAttr(jsonString string) string {
	result := jsonString

	regx := regexp.MustCompile(`,"[^"]*?":(""|null)`)
	result = regx.ReplaceAllString(result, "")
	
	regx = regexp.MustCompile(`"[^"]*?":(""|null),?`)
	result = regx.ReplaceAllString(result, "")
	
	return result
}

func (o CommonUtil) IsNumber(str string) bool {
	regx := regexp.MustCompile(`^\d*$`)
	return regx.MatchString(str)
}

func (o CommonUtil) IsFloat(str string) bool {
	regx := regexp.MustCompile(`^-?\d*(\.\d*)?$`)
	return regx.MatchString(str)
}

func (o CommonUtil) GetIntFromMap(data map[string]interface{}, name string) int {
	return o.GetIntFromString(fmt.Sprint(data[name]))
}

func (o CommonUtil) GetIntFromString(str string) int {
	amtStr := str
	if amtStr == "" {
		amtStr = "0"
	}
	amt, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return amt
}

func (o CommonUtil) GetFloat64FromString(str string) float64 {
	amtStr := str
	if amtStr == "" {
		amtStr = "0"
	}
	amt, err := strconv.ParseFloat(amtStr, 64)
	if err != nil {
		panic(err)
	}
	return amt
}
