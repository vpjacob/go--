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
