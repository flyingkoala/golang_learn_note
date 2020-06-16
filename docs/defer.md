# 2020-06-15

## defer
1.参考defer.go  

    func f1() (r int){
    	t:=5
    	defer func() {
    		t=t+1
    	}()
    	return t
    }
    
    func main(){
    	fmt.Println(f1())
    }
得到响应如下  

    5
原因是在返回r值时，需要将t的值赋值给r，而这个赋值行为是在defer之前执行的，所以defer中的赋值操作对r值没有影响  
解释代码如下  
    
    func f1_1() (r int){
    	t:=5
    	//1.赋值命令
    	r=t
    	//2.defer被插入到赋值和返回之间执行
    	func() {
    		t=t+1
    	}()
    	//3.空的return指令
    	return
    }
  
2.代码如下  

    func f2()(r int){
    	defer func(r int) {
    		r=r+5
    	}(r)
    	return 1
    }
    
    func main(){
    	fmt.Println(f2())
    }

 响应如下  

    1   
值不会改变的原因依旧是 这里改变的r是传值传进去的，是copy了一份，不会影响到原来返回的r  

3.代码如下  

    func f3()(r int){
    	defer func(r *int) {
    
    		*r=*r+5
    	}(&r)
    	return 1
    }
    
    func main(){
    	fmt.Println(f3())
    }  

得到响应  

    6  

传指针，defer执行的函数可以改变r的值


