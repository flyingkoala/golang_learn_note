# 2020.6.3
## 逃逸分析机制   
逃逸分析是编译器用来确定由程序创建的值所处位置的过程。具体来说，编译器执行静态代码分析，以确定是否可以将值放在构造函数的栈上，或者该值是否必须“逃逸”到堆上。    
查看逃逸相关信息命令   
`go tool compile -m make.go`  

1. 堆   
堆是除栈之外的第二个内存区域，用于存储值。堆不像栈那样是自清理的，因此使用这个内存的成本更大。首先，成本与垃圾收集器(GC)有关，垃圾收集器必须参与进来以保持该区域的清洁。当GC运行时，它将使用25%的可用CPU资源。此外，它可能会产生微秒级的“stop the world”延迟。拥有GC的好处是你不需要担心内存的管理问题，因为内存管理是相当复杂、也容易出错的。
堆上的值构成Go中的内存分配。这些分配对GC造成压力，因为堆中不再被指针引用的每个值都需要删除。需要检查和删除的值越多，GC每次运行时必须执行的工作就越多。因此，GC算法一直在努力在堆的大小分配和运行速度之间寻求平衡。   
2. 如果分配到栈上，待函数返回资源就被回收了。如果分配到堆上，函数返回后交给gc来管理该对象资源。   
3. 栈资源的分配及回收速度比堆要快，所以逃逸分析最大的好处应该是减少了GC的压力。   

##  指针逃逸
参考代码escapes.go   


    func main() {
		test()
		return   
     } 
    func test() *int{
		var a = 10
		return &a
    }
查看编译的逃逸相关信息如下   

    escapes.go:15:6: can inline test
    escapes.go:9:6: can inline main
    escapes.go:10:6: inlining call to test
    escapes.go:10:6: main &a does not escape
    escapes.go:17:9: &a escapes to heap
    escapes.go:16:6: moved to heap: a


典型的逃逸case，函数返回局部变量的指针。局部变量a被分配到堆上。   

## 栈空间不足逃逸
参考代码escapes.go 

    func main() {
    	run()
    	return
    }
    func run(){
    	t:=make([]int,1024,1024)
    	s:=make([]int,10000,10000)
    	for i:=0;i<len(t);i++{
    		s[i]=i
    		t[i]=i
    	}
    	return
    }
 查看编译的逃逸相关信息如下   

    escapes.go:12:6: can inline main
    escapes.go:25:9: make([]int, 10000, 10000) escapes to heap
    escapes.go:24:9: run make([]int, 1024, 1024) does not escape

当对象大小超过的栈帧大小时（详见go内存分配），变量对象发生逃逸被分配到堆上。   
当s的容量足够大时，s逃逸到堆上。t容量较小分配到栈上。
## 其他逃逸现象后续补充