package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_ = fibonacci(1000000)
	}
}

// 16 primeros f numbers BenchmarkFibonacci-16    	 7461058	       150.1 ns/op	     128 B/op	       1 allocs/op
// 100                   BenchmarkFibonacci-16    	 1599524	       863.2 ns/op	     896 B/op	       1 allocs/op
// 1000000               BenchmarkFibonacci-16    	     123	   9420438 ns/op	 8068671 B/op	       1 allocs/op

func BenchmarkFibonacciRecur(b *testing.B) {

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_ = fibonacciRecur(1000000)
	}
}

// 16  BenchmarkFibonacciRecur-16    	 3526744	       322.0 ns/op	     240 B/op	       4 allocs/op
// 100 BenchmarkFibonacciRecur-16    	  711152	      2379 ns/op	    2032 B/op	       7 allocs/op
// 1M  BenchmarkFibonacciRecur-16    	       2	 512457446 ns/op	62517160 B/op	      58 allocs/op
