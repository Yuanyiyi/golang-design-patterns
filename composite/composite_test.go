package composite

import (
	"fmt"
	"testing"
)

func TestPersonLevel(t *testing.T) {
	// 创建现今社会的等级
	mainLevel := NewPersonLevel("高层领导干部", "社会上层", 1000000000)

	// 中上层
	level11 := NewPersonLevel("中层领导干部", "中上层", 10000000)
	level12 := NewPersonLevel("大企业中层管理人员", "中上层", 10000000)

	// 中中层
	level1111 := NewPersonLevel("小企业主", "中中层", 10000)
	level1112 := NewPersonLevel("办事人员", "中中层", 10000)

	// 中下层
	level111111 := NewPersonLevel("农民工程序员", "中下层", 1000)
	level111112 := NewPersonLevel("个体服务者", "中下层", 1000)

	// 底层
	level1111111 := NewPersonLevel("失业人员", "底层", 100)

	//组成当前人类等级
	// 上层添加中上层
	mainLevel.Add(level11)
	mainLevel.Add(level12)

	// 中上层添加中中层
	level11.Add(level1111)
	level12.Add(level1112)

	// 中中层添加中下层
	level1111.Add(level111111)
	level1112.Add(level111112)

	// 中下层添加底层
	level111111.Add(level1111111)

	// 打印今社会的等级
	fmt.Println(mainLevel.ToString())
	for i := mainLevel.SubList.Front(); i != nil; i = i.Next() {
		em := i.Value.(*PersonLevel)
		fmt.Println(em.ToString())
		for j := i.Value.(*PersonLevel).SubList.Front(); j != nil; j = j.Next() {
			em := j.Value.(*PersonLevel)
			fmt.Println(em.ToString())
			for k := j.Value.(*PersonLevel).SubList.Front(); k != nil; k = k.Next() {
				em := k.Value.(*PersonLevel)
				fmt.Println(em.ToString())
				for l := k.Value.(*PersonLevel).SubList.Front(); l != nil; l = l.Next() {
					em := l.Value.(*PersonLevel)
					fmt.Println(em.ToString())
				}
			}
		}
	}

}
