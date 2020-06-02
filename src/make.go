package main

import "fmt"

func main()  {
	v:= make([]int,0,10)
	v=append(v,1)
	fmt.Println(v)

	x:= make([]int,10,10)
	x=append(x,1)
	fmt.Println(x)

}
