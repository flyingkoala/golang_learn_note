/*
@Time : 2020/6/17 17:22
@Author : wkang
@File : closure
@Description:闭包相关
*/
package main

//查看逃逸信息 go build -gcflags "-N -l -m" closure.go

import "fmt"

//闭包
func closure() func(int) int {
	var x int
	return func(a int) int {
		x++
		return a+x
	}
}

func main(){
	A:=closure()
	fmt.Println("x:",A(2))
}
