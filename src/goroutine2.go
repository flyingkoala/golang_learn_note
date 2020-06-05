/*
@Time : 2020/6/5 14:25
@Author : wkang
@File : goroutine2
@Description:
*/
package main

import (
	"fmt"
	"time"
)

func cal(a int , b int ,exitchan chan bool)  {
	c := a+b
	fmt.Printf("%d + %d = %d\n",a,b,c)
	time.Sleep(time.Second*2)
	exitchan <- true
}

func main() {

	exitchan := make(chan bool,10)  //声明并分配管道内存
	for i :=0 ; i<10 ;i++{
		go cal(i,i+1,exitchan)
	}
	for j :=0; j<10; j++{
		<- exitchan  //取信号数据，如果取不到则会阻塞
	}
	close(exitchan) // 关闭管道
}