package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	// 只能读取一次，意味着你读了，别热就不能读了。别人读了，你也不能读
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}
	// 类型转换
	fmt.Fprintf(w, "ready the data %s\n", string(body))

	//尝试再次读取，啥也读不到，但是也不会报错
	//body, err := io.ReadAll(r.Body)
}

// Query传参
func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v \n", values)
}

type ToyDuck struct {
	Name string
}

func (ToyDuck) swim() {}

func testStruct() {
	var pointer *ToyDuck = &ToyDuck{}
	// 解引用，得到结构体
	var duck ToyDuck = *pointer
	duck.swim()

	// 只声明

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I Love %s", r.URL.Path[1:])
}

// http://localhost:8080/golang
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
