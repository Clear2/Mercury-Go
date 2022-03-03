package main

import (
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 字段标签 tag, 在运行期间，可用反射获取标签信息，它常被用作格式校验，数据库关系映射等
func main() {
	//u := User{Name: "tom", Age: 19 }
	//v := reflect.ValueOf(u)
	//t := v.Type()

	//for i, n := 0, t.NumField(); i < n; i++ {
	//	fmt.Printf("%s, %v\n", t.Field(i).Tag, v.Field(i))
	//}
	initMethod()

}

// 初始化
func initMethod() {
	var u = User{Name: "jerry", Age: 18}
	// 如果初始化时没有给属性 设置对应的初始值，那么对应属性就是其类型的默认值
	var u1 = User{}
	// 对结构体指针进行初始化
	var u2 = new(User)
	u2.Name = "Hay"
	var u3 = &User{Name: "Boy", Age: 10}

	fmt.Println(u.Name, u1.Name, u1.Age, u2.Name, u3.Name)
	fmt.Println(&(u.Name), &(u.Age))

	u4 := u
	u4.Name = "Change Value"

	u5 := &u
	u5.Name = "Change Point"
	fmt.Println(u.Name, u4.Name, u5.Name)
}
