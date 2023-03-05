package abstract

// 抽象工厂模式，它和简单工厂模式的唯一区别，就是它返回的是接口而不是结构体。
//在你不公开内部实现的情况下，让调用者使用你提供的各种功能。

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

func NewPerson(name string, age int) Person {
	return &person{
		name: name,
		age:  age,
	}
}
