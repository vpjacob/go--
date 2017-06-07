// 表达式
package main

// 自增
func main() {
	a := 1
	// ++a //错误，不能前置

	// if (a++)>1 { //语句不能作为表达式

	// }

	p := &a
	println(p, *p) //0xc420039f68 1
	*p++           //相当于(*p)++
	println(p, *p) //0xc420039f68 2
}

// 初始化（数组，切片，字典，结构图）
// 左面花括号必须在尾部，不能另起一行
// 多个成员使用逗号分开
// 允许多行，但是要用逗号或者花括号结束

func a() {
	type data struct {
		x int
		s string
	}
	// var a data = data(1, "av")
	b := data{
		1,
		"av",
	}
	c := []int{
		1, 2}
	d := []int{1, 2, 3,
		4, 5,
	}
	println(b)
}


// switchCase
// 可以穿透，但是不能作为最后一个，穿透后面得有case，可以使用表达式，不能有重复的case条件
func main() {
	switch x := 5;x{
	case 1:
		x += 10
		println(x)
	case 5:
	case 6:
		x += 10
		println(x)
		fallthrough
	case 7:
		x += 4
	}
}


//for 循环
 for i := 0; i < count; i++ {
	x ++
}

for x < 10{
	x ++//相当于while x<10{}
}

for {
	break//相当于while ture{}  for ture{}
}

// break continue 前者中断switch for select 后者只能用在for循环，结束这一轮的循环
