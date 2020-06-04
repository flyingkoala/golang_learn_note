# 2020.6.4
## goroutine
### go语言使用MPG模式  
在传统的并发中起很多线程只会加大CPU和内存的开销，太多的线程会大量的消耗计算机硬件资源，造成并发量的瓶颈  
M指的是machine，一个M直接关联了一个内核线程。  
P指的是processor，代表了M所需的上下文环境，也是处理用户级代码逻辑的处理器。  
G指的是goroutine，其实本质上也是一种轻量级的线程。
1. 服务启动时 默认分配好cpu内核数量的P，并用一个slice维护  
2. processor和goroutine的创建都在proc.go中  
