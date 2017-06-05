// 基础学习


// 变量

var x = int // 自动初始化为0
var y = false // 具有自动推断类型

var x,y int //可以定义多个变量 而且类型可不一样
var x,y = 100,"abs"
var {
	x,y int
	a,s = 100,"abs"
}

//：=作用创建并赋值，和var没有区别
func main() {
	x := 100;
	a,s := 1,"abc"
}

// 但是要注意。不能提供数据类型。只能用在函数内部。
var x = 10
func main() {
	x := "av"//这个在函数内部重新创建了一个局部变量x，地址并不一样，值也不一样
}

// 简短模式并不是都是重新定义变量，也可能是退化赋值操作，内存中地址不变，值会变
func main() {
	x := 100
	x,y := 200,"av"
}

// 退化操作至少有一个变量重新定义，而且作用域要相同
func main() {
	x := 100
	// 下面错误，退化没有重新定义变量
	x := 200
	{
		// 不同作用域是重新定义，地址会变
		x := 300
	}
}

// 全局变量不使用没事，局部变量没使用，编译器会报错 declared and not used

var x int 
func main() {
	y := 10
	// 会报错
}







// 常量
// 一般使用常量来代替魔法数字
func main() {
	const a = 200

	const b,c = 200
	const s = "hello world"
	// 不会报错，没使用的b，c，s不会报错
	{
		const a = 300
		println(x)  // 结果是300
		// 在不同作用域，常量的值会变
	}
}

import "fmt"
func main() {
	const {
		x uint16 = 11
		y
		s = "av"
		z
	}
	// x 和y，s和z数值和类型相同
}


// 枚举
const {
	x = iota //0
	y		//1
	z		//2

}

const {
	_,_ = iota,iota * 10  //0,0 * 10
	a,b						//1,1 * 10
	c,d						//2,2*10
}

const {
	a = iota 	//0  
	b			//1
	c = 100		//100  	中断
	d			//100	和上面的一样
	e = iota    //4		恢复，但是包括cd，
	f 			//5		
}


