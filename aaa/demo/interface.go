package main

import "fmt"

type Animal interface {
	Bark() string
	Walk() string
}

// 必须同时实现 Bark() 和 Walk()方法 否则都不算实现了Animal接口
type Dog struct {
	name string
}

func (dog Dog) Bark() string {
	fmt.Println(dog.name + "Bark Bark")
	return "Bark"
}

func (dog Dog) Walk() string {
	fmt.Println(dog.name + "Walk Walk")
	return "Walk"
}

func main() {
	var animal Animal
	// animal = nil
	fmt.Println("animal values is", animal)
	// animal type = nil
	fmt.Printf("animal type %T\n", animal)

	animal = Dog{name: "Jerry"}
	animal.Bark()
	animal.Walk()

	//  {Jerry}
	fmt.Println("animal value is:", animal)
	// type = main.Dog
	fmt.Printf("animal type is: %T\n", animal)
}
