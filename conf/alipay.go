package conf

import (
	"github.com/smartwalle/alipay"
)

const (
	APPID          = ""
	GATWAY         = "https://openapi.alipaydev.com/gateway.do"
	ALIPUBLIC_KEY  = ""
	ALIPRIVATE_KEY = ""
)

var AliClient *alipay.Client

func InitAliClient() {
	client, err := alipay.New(APPID, ALIPUBLIC_KEY, ALIPRIVATE_KEY, false)
	if err != nil {
		panic(err)
	}

	AliClient = client

}
