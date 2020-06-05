# 2020.6.4
## goroutine
### go语言使用MPG模式  
在传统的并发中起很多线程只会加大CPU和内存的开销，太多的线程会大量的消耗计算机硬件资源，造成并发量的瓶颈  

![image](../img/1.png)

M指的是machine，一个M直接关联了一个内核线程。   

- M代表真正的执行计算资源可以认为它是os thread(系统线程)  

- M在绑定有效的P后，进入调度循环，而且M并不保留G的状态，这是G可以跨M调度的基础   
  
P指的是processor  

- 代表了M所需的上下文环境，也是处理用户级代码逻辑的处理器。  
- 拥有的各种G对象队列、链表、cache和状态  
  
G指的是goroutine，其实本质上也是一种轻量级的线程。  

- 调度系统的基本单位，存储了goroutine的执行stack信息、状态以及任务函数  

- G的眼里只有P P就是G的cpu  

- 相当于两级线程  
  
### Tips

1. 服务启动时 默认分配好cpu内核数量的P，并用一个slice维护  
2. processor和goroutine的创建都在proc.go中  
3. getg()为获取到的当前的g，这个看不到源码，是因为逻辑是在编译器执行时由编译器重写的  
4. 调度时，先从当前p中找可用的g，找不到就从别的p中偷，每61次，从全局的所有g中取一个


### proc.go
        // Create a new g running fn with siz bytes of arguments.
    // Put it on the queue of g's waiting to run.
    // The compiler turns a go statement into a call to this.
    // Cannot split the stack because it assumes that the arguments
    // are available sequentially after &fn; they would not be
    // copied if a stack split occurred.
    //go:nosplit
    func newproc(siz int32, fn *funcval) {
    	argp := add(unsafe.Pointer(&fn), sys.PtrSize)
    	gp := getg()
    	pc := getcallerpc()
    	systemstack(func() {
    		newproc1(fn, argp, siz, gp, pc)
    	})
    }
    
    // Create a new g running fn with narg bytes of arguments starting
    // at argp. callerpc is the address of the go statement that created
    // this. The new g is put on the queue of g's waiting to run.
    func newproc1(fn *funcval, argp unsafe.Pointer, narg int32, callergp *g, callerpc uintptr) {
    	_g_ := getg()//获取到当前goroutine，要创建goroutine，肯定是已经有一个最初的goroutine
    
    	if fn == nil {//空值检查
    		_g_.m.throwing = -1 // do not dump full stacks
    		throw("go of nil func value")
    	}
    	acquirem() // disable preemption because it can be holding p in a local var
    	siz := narg
    	siz = (siz + 7) &^ 7
    
    	// We could allocate a larger initial stack if necessary.
    	// Not worth it: this is almost always an error.
    	// 4*sizeof(uintreg): extra space added below
    	// sizeof(uintreg): caller's LR (arm) or return address (x86, in gostartcall).
    	//函数参数大小限制 太大就抛异常
    	if siz >= _StackMin-4*sys.RegSize-sys.RegSize {
    		throw("newproc: function arguments too large for new goroutine")
    	}
    	//从m中获取p
    	_p_ := _g_.m.p.ptr()
    	//从gfreelist（空闲的）获取g
    	newg := gfget(_p_)
    	//如果没有获取到g 就新建一个
    	if newg == nil {
    		newg = malg(_StackMin)
    		casgstatus(newg, _Gidle, _Gdead)//将g的状态设置为Gdead
    		//添加到allg数组中，防止gc扫描清除
    		allgadd(newg) // publishes with a g->status of Gdead so GC scanner doesn't look at uninitialized stack.
    	}
    	if newg.stack.hi == 0 {
    		throw("newproc1: newg missing stack")
    	}
    
    	if readgstatus(newg) != _Gdead {
    		throw("newproc1: new g is not Gdead")
    	}
    
    	totalSize := 4*sys.RegSize + uintptr(siz) + sys.MinFrameSize // extra space in case of reads slightly beyond frame
    	totalSize += -totalSize & (sys.SpAlign - 1)  // align to spAlign
    	sp := newg.stack.hi - totalSize
    	spArg := sp
    	if usesLR {
    		// caller's LR
    		*(*uintptr)(unsafe.Pointer(sp)) = 0
    		prepGoExitFrame(sp)
    		spArg += sys.MinFrameSize
    	}
    	if narg > 0 {
    		memmove(unsafe.Pointer(spArg), argp, uintptr(narg))
    		// This is a stack-to-stack copy. If write barriers
    		// are enabled and the source stack is grey (the
    		// destination is always black), then perform a
    		// barrier copy. We do this *after* the memmove
    		// because the destination stack may have garbage on
    		// it.
    		if writeBarrier.needed && !_g_.m.curg.gcscandone {
    			f := findfunc(fn.fn)
    			stkmap := (*stackmap)(funcdata(f, _FUNCDATA_ArgsPointerMaps))
    			if stkmap.nbit > 0 {
    				// We're in the prologue, so it's always stack map index 0.
    				bv := stackmapdata(stkmap, 0)
    				bulkBarrierBitmap(spArg, spArg, uintptr(bv.n)*sys.PtrSize, 0, bv.bytedata)
    			}
    		}
    	}
    
    	memclrNoHeapPointers(unsafe.Pointer(&newg.sched), unsafe.Sizeof(newg.sched))
    	newg.sched.sp = sp
    	newg.stktopsp = sp
    	newg.sched.pc = funcPC(goexit) + sys.PCQuantum // +PCQuantum so that previous instruction is in same function
    	newg.sched.g = guintptr(unsafe.Pointer(newg))
    	gostartcallfn(&newg.sched, fn)
    	newg.gopc = callerpc
    	newg.ancestors = saveAncestors(callergp)
    	newg.startpc = fn.fn
    	if _g_.m.curg != nil {
    		newg.labels = _g_.m.curg.labels
    	}
    	if isSystemGoroutine(newg, false) {
    		atomic.Xadd(&sched.ngsys, +1)
    	}
    	//更改当前g的状态为_Grunnable（可运行的）
    	casgstatus(newg, _Gdead, _Grunnable)
    
    	if _p_.goidcache == _p_.goidcacheend {
    		// Sched.goidgen is the last allocated id,
    		// this batch must be [sched.goidgen+1, sched.goidgen+GoidCacheBatch].
    		// At startup sched.goidgen=0, so main goroutine receives goid=1.
    		_p_.goidcache = atomic.Xadd64(&sched.goidgen, _GoidCacheBatch)
    		_p_.goidcache -= _GoidCacheBatch - 1
    		_p_.goidcacheend = _p_.goidcache + _GoidCacheBatch
    	}
    	//生成唯一的goid
    	newg.goid = int64(_p_.goidcache)
    	_p_.goidcache++
    	if raceenabled {
    		newg.racectx = racegostart(callerpc)
    	}
    	if trace.enabled {
    		traceGoCreate(newg, newg.startpc)
    	}
    	//将当前新生成的g，放入队列
    	runqput(_p_, newg, true)
    
    	if atomic.Load(&sched.npidle) != 0 && atomic.Load(&sched.nmspinning) == 0 && mainStarted {
    		wakep()
    	}
    	releasem(_g_.m)
    }
    
