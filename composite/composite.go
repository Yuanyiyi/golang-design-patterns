package composite

import (
	"container/list"
	"reflect"
	"strconv"
)

/*
组合模式定义： 组合是指使用组合和继承关系将聚合体及其组成元素分解成树状结构，以便客户端在不需要区分聚合或组成元素类型的情况下使用统一的接口操作类型
优点：1. 高层模块调用简单； 2. 更容易在组合体内加入新的对象：客户端不会因为加入了新的对象而更改源代码，满足“开闭原则”
缺点：1. 设计较复杂：客户端需要花更多时间理清类之间的层次关系
应用场景：1. 在需要表示一个对象整体与部分的层次结构的场合； 2. 为了简化代码结构，客户端要以统一的方式操作聚合体及组成元素

Go组合模式实现方式：
抽象构建角色：主要作用是为树叶构件和树枝构件声明公共接口，并实现它们的默认行为。在透明式的组合模式中抽象构件还声明访问和管理子类的接口；
           在安全式的组合模式中不声明访问和管理子类的接口，管理工作由数据构件完成。（总的抽象或接口，定义一些通用的方法，比如新增、删除）
树叶构件角色：是组合中的叶节点对象，它没有子节点，用于继承或实现抽象构件。
树枝构件角色/中间构件：是组合中的分支节点对象，它有子节点，用于继承和实现抽象构件。他的主要作用是存储和管理子部件，它的主要作用是存储和管理子部件，
                   同程包含Add()、Remove()、GetChild()等方法
*/

// 管理等级结构
type PersonLevel struct {
	Name        string
	Role        string
	IncomeLevel int
	SubList     *list.List
}

// 添加子等级
func (p *PersonLevel) Add(o *PersonLevel) {
	p.SubList.PushBack(o)
}

// 删除一个等级
func (p *PersonLevel) Remove(o *PersonLevel) {
	for i := p.SubList.Front(); i != nil; i = i.Next() {
		if reflect.DeepEqual(i.Value, o) {
			p.SubList.Remove(i)
		}
	}
}

// 获取等级列表
func (p *PersonLevel) GetSubList() *list.List {
	return p.SubList
}

// 获取等级的string 信息
func (p *PersonLevel) ToString() string {
	return "[ Name: " + p.Name + ", Role: " + p.Role + ", IncomeLevel: " + strconv.Itoa(p.IncomeLevel) + " ]"
}

// 实例化 管理等级对象
func NewPersonLevel(name, role string, income int) *PersonLevel {
	sub := list.New()
	return &PersonLevel{
		Name:        name,
		Role:        role,
		IncomeLevel: income,
		SubList:     sub,
	}
}
