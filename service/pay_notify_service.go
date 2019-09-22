package service

import (
	"DuckyGo/conf"
	"DuckyGo/model"
	"fmt"
	"net/http"
)

type NotifyPayService struct {
}

func (NotifyPayService) Notify(req *http.Request) {
	ok, err := conf.AliClient.GetTradeNotification(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	var paydb model.AlipaySuccess

	// 查询
	if err := model.DB.Where(&model.AlipaySuccess{OutTradeNo: ok.OutTradeNo,}).First(&paydb).Error; err != nil {
		fmt.Println("查找订单失败")
		return
	}

	// 更新
	if err := model.DB.Model(&paydb).Update(model.AutoInsert(ok)).Error; err != nil {
		fmt.Println("更新订单失败")
		return
	}

	fmt.Println("成功付款！ 验证完毕！")
}
