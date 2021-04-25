package main

/*
正整数分解质因数。例如：输入90,打印出90=2*3*3*5
(1)如果这个质数恰等于n，则说明分解质因数的过程已经结束，打印出即可。

(2)如果n<>k，但n能被k整除，则应打印出k的值，并用n除以k的商,作为新的正整数你n,重复执行第一步。

(3)如果n不能被k整除，则用k+1作为k的值,重复执行第一步。
*/
func factor(n int) []int {
	res := make([]int, 0)
	for i := 2; i <= n; i++ {
		for n >= i {
			if n%i == 0 {
				res = append(res, i)
				n = n / i
			} else {
				break
			}
		}
	}
	return res
}
