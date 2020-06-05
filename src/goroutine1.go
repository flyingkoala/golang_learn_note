package main

import (
	"fmt"
	"sync"
)

func cal(a int , b int ,n *sync.WaitGroup)  {
	c := a+b
	fmt.Printf("%d + %d = %d\n",a,b,c)
	defer n.Done() //goroutinue完成后, WaitGroup的计数-1
}

//使用sync包同步goroutine
func main() {
	var go_sync sync.WaitGroup //声明一个WaitGroup变量
	for i :=0 ; i<10 ;i++{
		go_sync.Add(1) // WaitGroup的计数加1
		go cal(i,i+1,&go_sync)
	}
	go_sync.Wait()  //等待所有goroutine执行完毕
}