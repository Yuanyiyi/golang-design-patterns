package template

import "fmt"

// 父类
type DocSuper struct {
	GetContent func() string
}

func (d DocSuper) DoOperate() {
	fmt.Println("对这个文档做了一些处理，文档是：", d.GetContent())
}

// 子类
type LocalDoc struct {
	DocSuper
}

func NewLocalDoc() *LocalDoc {
	c := new(LocalDoc)
	c.DocSuper.GetContent = c.GetContent
	return c
}

func (e *LocalDoc) GetContent() string {
	return "this is a LocalDoc."
}

// 子类
type NetDoc struct {
	DocSuper
}

func NewNetDoc() *NetDoc {
	c := new(NetDoc)
	c.DocSuper.GetContent = c.GetContent
	return c
}

func (c *NetDoc) GetContent() string {
	return "this is a net doc."
}
