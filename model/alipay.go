package model

import (
	"github.com/jinzhu/gorm"
	"github.com/smartwalle/alipay"
)

type AlipaySuccess struct {
	gorm.Model
	AuthAppId         string // App Id
	NotifyTime        string // 通知时间
	NotifyType        string // 通知类型
	NotifyId          string // 通知校验ID
	AppId             string // 开发者的app_id
	Charset           string // 编码格式
	Version           string // 接口版本
	SignType          string // 签名类型
	TradeNo           string // 支付宝交易号
	OutTradeNo        string // 商户订单号
	OutBizNo          string // 商户业务号
	BuyerId           string // 买家支付宝用户号
	BuyerLogonId      string // 买家支付宝账号
	SellerId          string // 卖家支付宝用户号
	SellerEmail       string // 卖家支付宝账号
	TradeStatus       string // 交易状态
	TotalAmount       string // 订单金额
	ReceiptAmount     string // 实收金额
	InvoiceAmount     string // 开票金额
	BuyerPayAmount    string // 付款金额
	PointAmount       string // 集分宝金额
	RefundFee         string // 总退款金额
	Subject           string // 总退款金额
	Body              string // 商品描述
	GmtCreate         string // 交易创建时间
	GmtPayment        string // 交易付款时间
	GmtRefund         string // 交易退款时间
	GmtClose          string // 交易结束时间
	FundBillList      string // 支付金额信息
	PassbackParams    string // 回传参数
	VoucherDetailList string // 优惠券信息
}

func AutoInsert(ok *alipay.TradeNotification) AlipaySuccess {
	return AlipaySuccess{
		AuthAppId:         ok.AuthAppId,
		NotifyTime:        ok.NotifyTime,
		NotifyType:        ok.NotifyType,
		NotifyId:          ok.NotifyId,
		AppId:             ok.AppId,
		Charset:           ok.Charset,
		Version:           ok.Version,
		SignType:          ok.SignType,
		TradeNo:           ok.TradeNo,
		OutTradeNo:        ok.OutTradeNo,
		OutBizNo:          ok.OutBizNo,
		BuyerId:           ok.BuyerId,
		BuyerLogonId:      ok.BuyerId,
		SellerId:          ok.SellerId,
		SellerEmail:       ok.SellerEmail,
		TradeStatus:       ok.TradeStatus,
		TotalAmount:       ok.TotalAmount,
		ReceiptAmount:     ok.ReceiptAmount,
		InvoiceAmount:     ok.InvoiceAmount,
		BuyerPayAmount:    ok.BuyerPayAmount,
		PointAmount:       ok.PointAmount,
		RefundFee:         ok.RefundFee,
		Subject:           ok.Subject,
		Body:              ok.Body,
		GmtCreate:         ok.GmtCreate,
		GmtPayment:        ok.GmtPayment,
		GmtRefund:         ok.GmtRefund,
		GmtClose:          ok.GmtClose,
		FundBillList:      ok.FundBillList,
		PassbackParams:    ok.PassbackParams,
		VoucherDetailList: ok.VoucherDetailList,
	}
}