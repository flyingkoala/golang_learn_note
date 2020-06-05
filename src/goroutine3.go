/*
@Time : 2020/6/5 14:43
@Author : wkang
@File : goroutine3
@Description:
*/
package main

import (
	"fmt"
	"sync"
)
//生产者
func Productor(mychan chan int,data int,wait *sync.WaitGroup)  {
	mychan <- data
	fmt.Println("product data：",data)
	wait.Done()
}
//消费者
func Consumer(mychan chan int,wait *sync.WaitGroup)  {
	a := <- mychan
	fmt.Println("consumer data：",a)
	wait.Done()
}
//主函数
func main() {

	datachan := make(chan int, 100)   //通讯数据管道
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go Productor(datachan, i,&wg) //生产数据
		wg.Add(1)
	}
	for j := 0; j < 10; j++ {
		go Consumer(datachan,&wg)  //消费数据
		wg.Add(1)
	}
	wg.Wait()
}