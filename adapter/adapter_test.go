package adapter

import "testing"

func TestPay(t *testing.T) {
	// 初始化对接接口
	var target TargetInterface
	// 同时调用支付宝和微信支付
	target = &NewAdapter{
		AlipayInterfce:     &AlipayPay{},
		WeChatPayInterface: &WeChatPay{},
	}
	// 业务中基于一个用户同时只能调用一种支付方式
	target.DealDiffPay("weixin", 99)
	target.DealDiffPay("alipay", 99)
}
