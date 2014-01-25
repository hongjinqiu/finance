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
