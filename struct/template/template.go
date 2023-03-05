package template

// 模板模式 (Template Pattern) 定义一个操作中算法的骨架，而将一些步骤延迟到子类中。
// 这种方法让子类在不改变一个算法结构的情况下，就能重新定义该算法的某些特定步骤。

import "fmt"

type Cook interface {
	fire()
	cooke()
	outfire()
}

// 抽象类
type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("fire")
}

func (CookMenu) outfire() {
	fmt.Println("outfire")
}

// 抽象类，具体由子类实现
func (CookMenu) cooke() {

}

func DoCooke(cm Cook) {
	cm.fire()
	cm.cooke()
	cm.outfire()
}

type TomatoMenu struct {
	CookMenu
}

func (TomatoMenu) cooke() {
	fmt.Println("cooke tomato")
}
