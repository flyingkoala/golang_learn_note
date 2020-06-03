package main

import "fmt"

func main()  {

	//初始化切片len不同 append数据结果不同
	v:= make([]int,0,10)
	v=append(v,1)
	fmt.Println(v)
	x:= make([]int,10,10)
	x=append(x,1)
	fmt.Println(x)

	//关于append是否会重新分配地址
	fmt.Println("append没有扩容")
	t:= make([]int,2,10)
	fmt.Println(t,&t[0])
	t=append(t,1)
	fmt.Println(t,&t[0])

	fmt.Println("append导致扩容")
	i:= make([]int,10,10)
	fmt.Println(i,&i[0])
	i=append(i,1)
	fmt.Println(i,&i[0])

}
