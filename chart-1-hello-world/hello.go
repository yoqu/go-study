package main

import "fmt"

//闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 这个函数使用任意数目的 `int` 作为参数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 我们将通过两个函数：`zeroval` 和 `zeroptr` 来比较指针和
// 值类型的不同。`zeroval` 有一个 `int` 型参数，所以使用值
// 传递。`zeroval` 将从调用它的那个函数中得到一个 `ival`
// 形参的拷贝。
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` 有一和上面不同的 `*int` 参数，意味着它用了一
// 个 `int`指针。函数体内的 `*iptr` 接着_解引用_ 这个指针，
// 从它内存地址得到这个地址对应的当前值。对一个解引用的指
// 针赋值将会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	var a string = "init"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Println("hello world")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	sum(1, 2)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	seq := intSeq()
	fmt.Println(seq())

}
