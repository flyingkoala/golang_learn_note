# golang_learn_note
### 2020.6.2
make创建两个int型切片，一个长度为0，容量为10，一个长度为10，容量为10，代码参见make.go   

    v:= make([]int,0,10)
 	v=append(v,1)
 	fmt.Println(v)
 	
 	x:= make([]int,10,10)
 	x=append(x,1)
 	fmt.Println(x)

显示结果如下图   
![image](https://github.com/flyingkoala/golang_learn_note/blob/master/image/20200603105945.png)



