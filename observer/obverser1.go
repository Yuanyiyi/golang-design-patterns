package observer

import (
	"fmt"
	"reflect"
	"runtime"
)

/*
观察者案例2
*/

// Observable 被观察者
type Observable interface {
	Attach(observer ...ObserverInterface) Observable
	Detach(observer ObserverInterface) Observable
	Notify() error
}

// ObservableConcrete 一个具体的 订单状态变化的被观察者
type ObservableConcrete struct {
	observerList []ObserverInterface
}

// Attach 注册观察者
// @param $observer ObserverInterface 观察者列表
func (o *ObservableConcrete) Attach(observer ...ObserverInterface) Observable {
	o.observerList = append(o.observerList, observer...)
	return o
}

// Detach 注销观察者
// @param $observer ObserverInterface 待注销的观察者
func (o *ObservableConcrete) Detach(observer ObserverInterface) Observable {
	if len(o.observerList) == 0 {
		return o
	}
	for k, observerItem := range o.observerList {
		if observer == observerItem {
			fmt.Println(runFuncName(), "注销:", reflect.TypeOf(observer))
			o.observerList = append(o.observerList[:k], o.observerList[k+1:]...)
		}
	}
	return o
}

// Notify 通知观察者
func (o *ObservableConcrete) Notify() (err error) {
	for _, observer := range o.observerList {
		if err = observer.Do(o); err != nil {
			return err
		}
	}
	return nil
}

// ObserverInterface 定义一个观察者的接口
type ObserverInterface interface {
	// 自身的业务
	Do(o Observable) error
}

// 修改订单状态
type OrderStatus struct {
}

// Do 具体业务
func (observer *OrderStatus) Do(o Observable) (err error) {
	fmt.Println(runFuncName(), "修改订单状态...")
	return nil
}

// OrderStatusLog 记录订单状态变更日志
type OrderStatusLog struct {
}

// Do 具体业务
func (observer *OrderStatusLog) Do(o Observable) (err error) {
	fmt.Println(runFuncName(), "记录订单状态变更日志...")
	return nil
}

// CouponRefund 退优惠券
type CouponRefund struct {
}

// Do 具体业务
func (observer *CouponRefund) Do(o Observable) (err error) {
	fmt.Println(runFuncName(), "退优惠券...")
	return nil
}

// PromtionRefund 换优惠活动资格
type PromotionRefund struct {
}

// Do
func (observer *PromotionRefund) Do(o Observable) (err error) {
	fmt.Println(runFuncName(), "还优惠活动资格...")
	return nil
}

// StockRefund 还库存
type StockRefund struct {
}

// Do 具体业务
func (observer *StockRefund) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "还库存...")
	return
}

// GiftCardRefund 还礼品卡
type GiftCardRefund struct {
}

// Do 具体业务
func (observer *GiftCardRefund) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "还礼品卡...")
	return
}

// WalletRefund 退钱包余额
type WalletRefund struct {
}

// Do 具体业务
func (observer *WalletRefund) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "退钱包余额...")
	return
}

// DeliverBillStatus 修改发货单状态
type DeliverBillStatus struct {
}

// Do 具体业务
func (observer *DeliverBillStatus) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "修改发货单状态...")
	return
}

// DeliverBillStatusLog 记录发货单状态变更日志
type DeliverBillStatusLog struct {
}

// Do 具体业务
func (observer *DeliverBillStatusLog) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "记录发货单状态变更日志...")
	return
}

// Refund 生成退款单
type Refund struct {
}

// Do 具体业务
func (observer *Refund) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "生成退款单...")
	return
}

// Invoice 生成发票-红票
type Invoice struct {
}

// Do 具体业务
func (observer *Invoice) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "生成发票-红票...")
	return
}

// Email 发邮件
type Email struct {
}

// Do 具体业务
func (observer *Email) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "发邮件...")
	return
}

// Sms 发短信
type Sms struct {
}

// Do 具体业务
func (observer *Sms) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "发短信...")
	return
}

// WechatNotify 发微信消息
type WechatNotify struct {
}

// Do 具体业务
func (observer *WechatNotify) Do(o Observable) (err error) {
	// code...
	fmt.Println(runFuncName(), "发微信消息...")
	return
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
