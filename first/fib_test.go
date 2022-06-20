package first

import "testing"

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib5(b *testing.B)  { benchmarkFib(b, 5) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }

//增加benchtime，使该函数在执行时增加可以重复执行的次数，否则因为时间太短执行次数太少，得到的平均执行函数用时就不准确
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

//benchmark基准测试，go里默认是一秒钟，即测试函数的时间大约在一秒钟左右，如果该函数的执行时间超过一秒，可能无法返回，此时需要延长
func BenchmarkFib400(b *testing.B) { benchmarkFib(b, 400) }
