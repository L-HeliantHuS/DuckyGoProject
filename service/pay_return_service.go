package service

import (
	"DuckyGo/conf"
	"DuckyGo/serializer"
	"net/url"
)

type ReturnPayService struct {
}

func (ReturnPayService) Return(data url.Values) *serializer.Response {

	if ok, _ := conf.AliClient.VerifySign(data); ok != true {
		return &serializer.Response{
			Status: 40002,
			Msg:    "fali, Error sign! Fuck!",
		}
	}

	return &serializer.Response{
		Data: nil,
		Msg:  "Success!",
	}
}
