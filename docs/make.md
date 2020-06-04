# 2020.6.2
## append
make创建两个int型切片，一个长度为0，容量为10，一个长度为10，容量为10，代码参见make.go   

    v:= make([]int,0,10)
 	v=append(v,1)
 	fmt.Println(v)
 	
 	x:= make([]int,10,10)
 	x=append(x,1)
 	fmt.Println(x)

显示结果如下

    [1]
    [0 0 0 0 0 0 0 0 0 0 1]


原因是append函数将元素附加到切片的末尾。   
## 关于append对切片地址的影响。
代码参见make.go  
 
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
显示结果如下   

    append没有扩容
    [0 0] 0xc0000660f0
    [0 0 1] 0xc0000660f0
    append导致扩容
    [0 0 0 0 0 0 0 0 0 0] 0xc000066140
    [0 0 0 0 0 0 0 0 0 0 1] 0xc0000960a0

不难得出结论，append时，切片没有扩容，内存地址不会变，存在动态扩容，则分配新的内存地址   
slice这种数据结构便于使用和管理数据集合，可以理解为是一种“动态数组”，slice也是围绕动态数组的概念来构建的。   
以下两条规则：   
1. 如果切片的容量小于1024个元素，那么扩容的时候slice的cap就翻番，乘以2；一旦元素个数超过1024个元素，增长因子就变成1.25，即每次增加原来容量的四分之一。   
2. 如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组，如果扩容之后，超过了原数组的容量，那么，Go就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。   

## make源码注释翻译
make内置函数分配并初始化一个类型的对象 slice，map或chan。像new一样，第一个参数是类型，而不是值。与new不同，make的返回类型与其返回类型相同参数，而不是指向它的指针。结果的规格取决于类型：   
1. slice：大小指定长度。 切片的容量等于其长度。 可以提供第二个整数参数来指定不同的容量； 它必须不小于长度。 例如，make（[] int，0，10）分配一个大小为10的基础数组，并返回一个长度为0且容量为10的切片，该切片由该基础数组支持。   
2. map：为空的map分配足够的空间以容纳指定数量的元素。 该大小可以省略，在这种情况下，分配的起始大小较小。   
3. channel：使用指定的缓冲区容量初始化通道的缓冲区。 如果为零或忽略大小，则通道不缓冲。   