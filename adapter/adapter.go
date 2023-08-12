package adapter

import "fmt"

/*
适配器定义：适配器（Adapter）指将某种接口或数据结构转换为客户端期望的类型，使得不兼容的类或对象能够一起协作
优点：1）将目标类和适配者解耦：为解决目标类和适配者接口不一致的问题。这样通过适配器可以透明地调用目标接口，在很多业务场景中符合开闭原则。
     2）复用现存的类：解决了目标类和适配者类的不一致问题。
缺点：1）增加了系统的复杂性：适配器编写过程需要结合业务场景全面考虑。
     2）增加代码的阅读复杂度：降低代码可读性，过多适配器会让系统越来越复杂。

Go适配器模式实现方
适配器模式（Adapter）包含以下主要角色。
目标（Target）接口：当前系统业务所期待的接口，它可以是抽象类或接口。
适配者（Adaptee）类：它是被访问和适配的现存组件库中的组件接口。
适配器（Adapter）类：它是一个转换器，通过继承或引用适配者的对象，把适配者接口转换成目标接口，让客户按目标接口的格式访问适配者。

适配器模式在我所接触的业务中，一个是支付SDK的集成形成同一个支付接口调用，聚合广告SDK的集成形成统一广告接口调用。下面我们来具体看示例应用。通过适配器实现支付宝SDK和微信SDK的集成。
*/

// 支付宝支付SDK
type AlipayInterfce interface {
	Pay(money int)
}

type AlipayPay struct{}

func (a *AlipayPay) Pay(money int) {
	fmt.Println("支付宝支付, 费用是: ", money)
}

// 微信支付
type WeChatPayInterface interface {
	WXPay(money int)
}

type WeChatPay struct {
}

func (w *WeChatPay) WXPay(money int) {
	fmt.Println("微信支付, 费用是: ", money)
}

// 目标接口，能支持传入支付宝或者微信支付进行支付
type TargetInterface interface {
	DealDiffPay(payType string, money int)
}

// adapter, 实现微信和支付宝支付
type NewAdapter struct {
	AlipayInterfce
	WeChatPayInterface
}

func (n *NewAdapter) DealDiffPay(payType string, money int) {
	if payType == "alipay" {
		n.AlipayInterfce.Pay(money)
	} else if payType == "weixin" {
		n.WeChatPayInterface.WXPay(money)
	}
}
