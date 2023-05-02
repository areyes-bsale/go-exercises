package main

func fibonacci(x int) []int {
	result := make([]int, x)
	result[1] = 1
	for i := 2; i < x; i++ {
		result[i] = result[i-2] + result[i-1]
	}
	return result
}

func fibonacciRecur(x int) []int {
	if x == 1 {
		return []int{0}
	} else if x == 2 {
		return []int{0, 1}
	} else {
		arr := fibonacciRecur(x - 1)
		return append(arr, arr[x-2]+arr[x-3])
	}
}
