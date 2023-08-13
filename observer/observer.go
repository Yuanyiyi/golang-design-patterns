package observer

import (
	"context"
	"fmt"
	"sync"
)

/*
原文：https://zhuanlan.zhihu.com/p/637532728
在观察者模式中，核心的角色包含三类：
	Observer: 观察者，指的是关注食物动态的角色
	Event：事物的变更事件，其中Topic标识了事物的身份以及变更的类型，Val是变更详情
	EventBus：事件总线。位于观察者与事物之间承上启下的代理层。负责维护管理观察者，并且在事物发生变更时，将情况同步给每个观察者

观察者模式的核心在于建立EventBus的角色。由于EventBus模块的诞生，实现了观察者与具体被观察事物之间的解耦：
	* 针对于观察者而言，需要向EventBus完成注册操作，注册时需要声明自己关心的变更事件类型（调用EventBus的Subscribe方法），不在需要直接和事物打交道
	* 针对于事物而言，在其发生变更时，只需要将变更情况向EventBus统一汇报即可（调用EventBus的Publish方法），不再需要和每个观察者直接交互
	* 对于EventBus，需要提前维护好每个观察者和被关注事物之间的映射关系，保证在变更事件到达时，能找到所有的观察者逐一进行通知（调用Observer 的OnChange方法）

优点：1）高内聚：不同业务代码变动互不影响；2）可复用：新的业务订阅不同接口（主题）；3）极易扩展：新增接口（主题），新增业务（订阅者）
*/

// 如下是实现

type Event struct {
	Topic string
	Val   interface{}
}

type Observer interface {
	OnChange(ctx context.Context, e *Event) error
}

type EventBus interface {
	Subscribe(topic string, o Observer)
	Unsubscribe(topic string, o Observer)
	Publish(ctx context.Context, e *Event)
}

type BaseObserver struct {
	name string
}

func NewBaseObserver(name string) *BaseObserver {
	return &BaseObserver{name: name}
}

func (b *BaseObserver) OnChange(ctx context.Context, e *Event) error {
	fmt.Printf("observer: %s, event key: %s, event val: %v\n", b.name, e.Topic, e.Val)
	// ...
	return nil
}

type BaseEventBus struct {
	mux       sync.RWMutex
	observers map[string]map[Observer]struct{}
}

func NewBaseEventBus() BaseEventBus {
	return BaseEventBus{
		observers: make(map[string]map[Observer]struct{}),
	}
}

// 订阅：注册观察者
func (b *BaseEventBus) Subscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	_, ok := b.observers[topic]
	if !ok {
		b.observers[topic] = make(map[Observer]struct{})
	}
	b.observers[topic][o] = struct{}{}
}

// 取消订阅：删除观察者
func (b *BaseEventBus) Unsubscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	delete(b.observers[topic], o)
}

// 同步模式： 将变更同步给观察者
type SyncEventBus struct {
	BaseEventBus
}

func NewSyncEventBus() *SyncEventBus {
	return &SyncEventBus{BaseEventBus: NewBaseEventBus()}
}

// 通知观察者变更事件
func (s *SyncEventBus) Publish(ctx context.Context, e *Event) {
	s.mux.Lock()
	defer s.mux.Unlock()

	subscribes := s.observers[e.Topic]

	errs := make(map[Observer]error)
	for subscribe := range subscribes {
		if err := subscribe.OnChange(ctx, e); err != nil {
			errs[subscribe] = err
		}
	}

	s.handleErr(ctx, errs)
}

func (s *SyncEventBus) handleErr(ctx context.Context, errs map[Observer]error) {
	for o, err := range errs {
		// 处理 publish失败的observer
		fmt.Printf("observer: %v, err: %v\n", o, err)
	}
}

// 异步处理方式
type observerWithErr struct {
	o   Observer
	err error
}

type AsyncEventBus struct {
	BaseEventBus
	errC chan *observerWithErr
	ctx  context.Context
	stop context.CancelFunc
}

func NewAsyncEventBus() *AsyncEventBus {
	aBus := AsyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}
	aBus.ctx, aBus.stop = context.WithCancel(context.Background())
	// 处理错误的异步守护协程
	go aBus.handleErr()
	return &aBus
}

func (a *AsyncEventBus) Stop() {
	a.stop()
}

func (a *AsyncEventBus) Publish(ctx context.Context, e *Event) {
	a.mux.RLock()
	defer a.mux.RUnlock()

	subscribes := a.observers[e.Topic]
	for subscribe := range subscribes {
		// shadow
		subscribe := subscribe
		go func() {
			if err := subscribe.OnChange(ctx, e); err != nil {
				select {
				case <-a.ctx.Done():
				case a.errC <- &observerWithErr{
					o:   subscribe,
					err: err}:
				}
			}
		}()
	}
}

func (a *AsyncEventBus) handleErr() {
	for {
		select {
		case <-a.ctx.Done():
			return
		case resp := <-a.errC:
			// 处理publish失败的observe
			fmt.Printf("observer: %v, err: %v\n", resp.o, resp.err)
		}
	}
}
