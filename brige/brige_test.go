package brige

import "testing"

func TestBrige(t *testing.T) {
	// 送一个蓝色的带水晶的皇冠
	color := Blue{}
	crystal := Crystal{}
	c := Crown{
		Color:   &color,
		DressUp: &crystal,
	}
	c.SendGift()

	// 送一个红色的飞机
	color2 := Red{}
	p := Plane{
		Color:   &color2,
		DressUp: nil,
	}
	p.SendGift()
}
