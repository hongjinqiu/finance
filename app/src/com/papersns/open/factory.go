package open

import (

)

type ThirdOpen interface {
	GetUserInfo(url string) map[string]interface{}
	GetShopInfo(resStruct map[string]interface{}) map[string]interface{}
}

type ThirdOpenFactory struct {
	
}

func (o ThirdOpenFactory) GetThirdOpen() ThirdOpen {
	return JingdongInterface{}
}
