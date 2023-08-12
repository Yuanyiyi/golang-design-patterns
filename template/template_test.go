package template

import "testing"

func TestDocSuper(t *testing.T) {
	localD := NewLocalDoc()
	netD := NewNetDoc()

	localD.DoOperate()
	netD.DoOperate()
}
