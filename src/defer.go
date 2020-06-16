package main

import "fmt"

func f1() (r int){
	t:=5
	defer func() {
		t=t+1
	}()
	return t
}

//解释f1()
func f1_1() (r int){
	t:=5
	//1.赋值命令
	r=t
	//2.defer被插入到赋值和返回之间执行
	func() {
		t=t+1
	}()
	//3.空的return指令
	return
}

func f2()(r int){
	defer func(r int) {
		fmt.Println("呵呵",r)
		r=r+5
		fmt.Println("呵呵",r)
	}(r)
	return 1
}

func f3()(r int){
	defer func(r *int) {

		*r=*r+5
	}(&r)
	return 1
}

func main(){
	fmt.Println(f3())
}