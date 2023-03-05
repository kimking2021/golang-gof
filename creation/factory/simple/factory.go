package simple

//简单工厂模式是最常用、最简单的。
//它就是一个接受一些参数，然后返回 Person 实例的函数。
//简单工厂模式可以确保我们创建的实例具有需要的参数，进而保证实例的方法可以按预期执行。

//在实际开发中，我建议返回非指针的实例，因为我们主要是想通过创建实例，调用其提供的方法，而不是对实例做更改。
//如果需要对实例做更改，可以实现SetXXX的方法。
//通过返回非指针的实例，可以确保实例的属性，避免属性被意外 / 任意修改。

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
