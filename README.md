# golang_learn_note   
  
## 2020.6.2
### append
make创建两个int型切片，一个长度为0，容量为10，一个长度为10，容量为10，代码参见make.go   

    v:= make([]int,0,10)
 	v=append(v,1)
 	fmt.Println(v)
 	
 	x:= make([]int,10,10)
 	x=append(x,1)
 	fmt.Println(x)

显示结果如下图   
![image](https://github.com/flyingkoala/golang_learn_note/blob/master/image/20200603105945.png)

原因是append函数将元素附加到切片的末尾。   
### 关于append对切片地址的影响。
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
显示结果如下图   
![image](https://github.com/flyingkoala/golang_learn_note/blob/master/image/20200603135526.png)   
不难得出结论，append时，切片没有扩容，内存地址不会变，存在动态扩容，则分配新的内存地址

### make源码注释翻译
make内置函数分配并初始化一个类型的对象 slice，map或chan。像new一样，第一个参数是类型，而不是值。与new不同，make的返回类型与其返回类型相同参数，而不是指向它的指针。结果的规格取决于类型：   
1. slice：大小指定长度。 切片的容量等于其长度。 可以提供第二个整数参数来指定不同的容量； 它必须不小于长度。 例如，make（[] int，0，10）分配一个大小为10的基础数组，并返回一个长度为0且容量为10的切片，该切片由该基础数组支持。   
2. map：为空的map分配足够的空间以容纳指定数量的元素。 该大小可以省略，在这种情况下，分配的起始大小较小。   
3. channel：使用指定的缓冲区容量初始化通道的缓冲区。 如果为零或忽略大小，则通道不缓冲。   

## 2020.6.3
### 逃逸分析机制   
逃逸分析是编译器用来确定由程序创建的值所处位置的过程。具体来说，编译器执行静态代码分析，以确定是否可以将值放在构造函数的栈上，或者该值是否必须“逃逸”到堆上。    
查看逃逸相关信息命令   
`go tool compile -m make.go`  

1. 堆   
堆是除栈之外的第二个内存区域，用于存储值。堆不像栈那样是自清理的，因此使用这个内存的成本更大。首先，成本与垃圾收集器(GC)有关，垃圾收集器必须参与进来以保持该区域的清洁。当GC运行时，它将使用25%的可用CPU资源。此外，它可能会产生微秒级的“stop the world”延迟。拥有GC的好处是你不需要担心内存的管理问题，因为内存管理是相当复杂、也容易出错的。
堆上的值构成Go中的内存分配。这些分配对GC造成压力，因为堆中不再被指针引用的每个值都需要删除。需要检查和删除的值越多，GC每次运行时必须执行的工作就越多。因此，GC算法一直在努力在堆的大小分配和运行速度之间寻求平衡。



