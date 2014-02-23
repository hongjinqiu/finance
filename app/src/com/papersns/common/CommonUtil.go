package common

import (
	"regexp"
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
