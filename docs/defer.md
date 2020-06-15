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