package brige

import "fmt"

/*
桥接模式定义：桥（Bridge）使用组合关系将代码的实现层和抽象层分离，让实现层与抽象层代码可以分别自由变化。
优点：1）实现抽象和实现的分离，扩展能力强
	 2）提高了系统的可扩充性：在两个变化维度中任意扩展一个维度，都不需要修改原有的系统
缺点：1）桥接模式的引入会增加系统的理解与设计难度，由于聚合关联关系建立在抽象层，要求开发者针对抽象进行设计与编程
	 2）桥接模式要求正确识别系统中两个独立变化的维度，因此其使用范围具有一定的局限性

应用场景：
1. 抽象层代码和实现层代码分别需要自由扩展
2. 需要独立封装或复用实现层代码

桥接模式结构：
1. Abstraction(抽象类)：用于定义抽象类接口
2. RefinedAbstratction(扩展抽象类)：扩充由Abstraction定义的接口，通常情况下它不再是抽象类而是具体类，实现了在Abstraction中声明的抽象业务方法，在RefinedAbstraction中可以调用在Implementor中定义的业务方法
3. Implementor（实现类接口，比如例子中颜色，装饰等）：定义实现类的接口，一般而言，它不与Abstraction的接口一致。它只提供基本的或者简单的操作。
4. ConcreteImplementor（具体实现类，比如水晶，黄色，蓝色）：具体实现Implementor接口，在不同的ConcreteImplementor中提供基本操作的不同实现，在程序运行时，ConcreteImplentor将替换其父类对象，提供给抽象类具体的业务操作方法。
*/

// 颜色
type Color interface {
	Use()
}

// red
type Red struct {
}

func (r *Red) Use() {
	fmt.Println("use red color.")
}

// blue
type Blue struct {
}

func (b *Blue) Use() {
	fmt.Println("use blue color.")
}

// 装扮
type DressUp interface {
	Decorate()
}

// 水晶
type Crystal struct {
}

func (c *Crystal) Decorate() {
	fmt.Println("use crystal dress up.")
}

// 抽象类
type Plane struct {
	Color
	DressUp
}

func (p *Plane) SendGift() {
	p.Use()
	if p.DressUp != nil {
		p.DressUp.Decorate()
	}
	fmt.Println("送飞机礼物")
}

// 皇冠
type Crown struct {
	Color
	DressUp
}

func (c *Crown) SendGift() {
	c.Use()
	if c.DressUp != nil {
		c.DressUp.Decorate()
	}
	fmt.Println("送皇冠礼物")
}
