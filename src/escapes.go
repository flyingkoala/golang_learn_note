/*
@Time : 2020/6/3 15:25
@Author : wkang
@File : escapes
@Description:
*/
package main

import "fmt"

//编译命令 go tool compile -m escapes.go

func main() {
	test()

	run()

	var size int =10
	t:= make([]int,size)
	for i:=0;i<size ;i++  {
		t[i]=i
	}
	unknownSize(size)
	return
}
//指针逃逸
func test() *int{
	var a = 10
	return &a
}
//栈空间不足逃逸
func run(){
	t:=make([]int,1024,1024)
	s:=make([]int,10000,10000)
	for i:=0;i<len(t);i++{
		s[i]=i
		t[i]=i
	}
	return
}
//动态类型逃逸
func unknownSize(p interface{}){
	fmt.Println(p)
}