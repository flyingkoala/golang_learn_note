/*
@Time : 2020/6/16 10:29
@Author : wkang
@File : defer2
@Description:
*/
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func e1(){
	var err error
	defer fmt.Println(err)
	err = errors.New("defer1 error")
	return
}
func e2(){
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("defer2 error")
	return
}

func e3(){
	var err error
	//执行到这一行时 func中的err已赋值为nil，所以之后的输出为nil
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("defer2 error")
	return
}

func main(){
	//defer的执行速度为先入后出，所以依次执行的defer是 e3->e3->e1

	e1()
	e2()
	e3()
}
