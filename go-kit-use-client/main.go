package main

import (
	"fmt"
	"go-kit-use-client/Client"
	"go-kit-use-client/EndPoint"
	"go-kit-use-client/Transport"
)

// 调用我们在client封装的函数就好了
func main() {
	//i, err := Client.Direct("GET", "http://127.0.0.1:8000", Transport.HelloEncodeRequestFunc, Transport.HelloDecodeResponseFunc, EndPoint.HelloRequest{Name: "wangpan"})
	i, err := Client.ServiceDiscovery("GET", "http://127.0.0.1:8500", Transport.ByeEncodeRequestFunc, Transport.ByeDecodeResponseFunc, EndPoint.HelloRequest{Name: "wangpan"}, "测试", true, "test")

	if err != nil {
		fmt.Println(err)
		return
	}
	res, ok := i.(EndPoint.HelloResponse)
	if !ok {
		fmt.Println("no ok")
		return
	}
	fmt.Println(res)
}
