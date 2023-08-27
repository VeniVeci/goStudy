package main

import "fmt"

func main() {
	// for true  {
	//     fmt.Printf("这是无限循环。\n")
	// }

	var a int = 5
	var b int = 7
	var maxNum = myMax(a, b)
	fmt.Printf("max number is %d\n", maxNum)

	n1 := 2
	n2 := 4
	maxNum = myMax(n1, n2)
	fmt.Printf("max number is %d\n", maxNum)

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

/* 函数返回两个数的最大值 */
func myMax(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
