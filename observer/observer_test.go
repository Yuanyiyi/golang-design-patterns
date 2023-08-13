package observer

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSyncEventBus(t *testing.T) {
	observerA := NewBaseObserver("a")
	observerB := NewBaseObserver("b")
	observerC := NewBaseObserver("c")
	observerD := NewBaseObserver("d")

	sbus := NewSyncEventBus()
	topic := "order_finish"
	sbus.Subscribe(topic, observerA)
	sbus.Subscribe(topic, observerB)
	sbus.Subscribe(topic, observerC)
	sbus.Subscribe(topic, observerD)

	sbus.Publish(context.Background(), &Event{
		Topic: topic,
		Val:   "order_id: xxx",
	})
}

func TestAsyncEventBus(t *testing.T) {
	observerA := NewBaseObserver("a")
	observerB := NewBaseObserver("b")
	observerC := NewBaseObserver("c")
	observerD := NewBaseObserver("d")

	abus := NewAsyncEventBus()
	defer abus.Stop()

	topic := "order_finish"
	abus.Subscribe(topic, observerA)
	abus.Subscribe(topic, observerB)
	abus.Subscribe(topic, observerC)
	abus.Subscribe(topic, observerD)

	abus.Publish(context.Background(), &Event{
		Topic: topic,
		Val:   "order_id: xxx",
	})

	<-time.After(time.Second)
}

func TestOrder(t *testing.T) {
	// 创建 未支付取消订单 “主题”
	fmt.Println("----------------------- 未支付取消订单 “主题”")
	orderUnPaidCancelSubject := &ObservableConcrete{}
	orderUnPaidCancelSubject.Attach(
		&OrderStatus{},
		&OrderStatusLog{},
		&CouponRefund{},
		&PromotionRefund{},
		&StockRefund{},
	)
	orderUnPaidCancelSubject.Notify()

	// 创建 超时关单 “主题”
	fmt.Println("----------------------- 超时关单 “主题”")
	orderOverTimeSubject := &ObservableConcrete{}
	orderOverTimeSubject.Attach(
		&OrderStatus{},
		&OrderStatusLog{},
		&CouponRefund{},
		&PromotionRefund{},
		&StockRefund{},
		&Email{},
		&Sms{},
		&WechatNotify{},
	)
	orderOverTimeSubject.Notify()

	// 创建 已支付取消订单 “主题”
	fmt.Println("----------------------- 已支付取消订单 “主题”")
	orderPaidCancelSubject := &ObservableConcrete{}
	orderPaidCancelSubject.Attach(
		&OrderStatus{},
		&OrderStatusLog{},
		&CouponRefund{},
		&PromotionRefund{},
		&StockRefund{},
		&GiftCardRefund{},
		&WalletRefund{},
		&Refund{},
		&Invoice{},
		&Email{},
		&Sms{},
		//&WechatNotify{},
	)
	orderPaidCancelSubject.Notify()

	// 创建 取消发货单 “主题”
	fmt.Println("----------------------- 取消发货单 “主题”")
	deliverBillCancelSubject := &ObservableConcrete{}
	deliverBillCancelSubject.Attach(
		&OrderStatus{},
		&OrderStatusLog{},
		&DeliverBillStatus{},
		&DeliverBillStatusLog{},
		&StockRefund{},
		&GiftCardRefund{},
		&WalletRefund{},
		&Refund{},
		&Invoice{},
		&Email{},
		&Sms{},
		//&WechatNotify{},
	)
	deliverBillCancelSubject.Notify()

	// 创建 拒收 “主题”
	fmt.Println("----------------------- 拒收 “主题”")
	deliverBillRejectSubject := &ObservableConcrete{}
	deliverBillRejectSubject.Attach(
		&OrderStatus{},
		&OrderStatusLog{},
		&DeliverBillStatus{},
		&DeliverBillStatusLog{},
		&StockRefund{},
		&GiftCardRefund{},
		&WalletRefund{},
		&Refund{},
		&Invoice{},
		&Email{},
		&Sms{},
		//&WechatNotify{},
	)
	deliverBillRejectSubject.Notify()

	// 未来可以快速的根据业务的变化 创建新的主题 从而快速构建新的业务接口
	fmt.Println("----------------------- 未来的扩展...")

}
