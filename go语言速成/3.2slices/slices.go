package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("原始切片:", numbers)

	evens := make([]int, 0, len(numbers))
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}

	// 打印新切片
	fmt.Println("偶数切片:", evens)
	// Print the new slice
	fmt.Println("奇数切片:"	, evens)

	//make 创建切片
	makenum := make([]int, 5)
	fmt.Println("使用make创建的切片:", makenum)

	//test
	//如果我告诉你 Go 使用 2x 算法来增加数组长度，你猜下面将会打印什么?
	scores := make([]int, 0, 5)
	c := cap(scores) // 译者注：cap()可以用来查看数组或 slice 的容量
	fmt.Println("初始容量:",c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		// 如果容量改变了
		// Go 必须增加数组长度来容纳新的数据
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println("容量改变:",c)
		}
	}

}
