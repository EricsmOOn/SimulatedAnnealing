package main

import (
	"fmt"
	"github.com/EricsmOOn/SimulatedAnnealing"
	"math"
	"testing"
)

func main() {
	// 实验证明使用 goroutine 处理较快的函数时会得不偿失
	// 确保 goroutine 使用在耗时函数上以保证协程额外时间使用远小于函数执行时间
	// 否则无需使用 goroutine
	fmt.Println("no sa", testing.Benchmark(BenchmarkSA).String())
	fmt.Println("go sa", testing.Benchmark(BenchmarkSAGo).String())
}

func BenchmarkSA(b *testing.B) {
	data := SimulatedAnnealing.ReadData()
	for i := 0; i < b.N; i++ {
		SimulatedAnnealing.GetResult(data)
	}
}

func BenchmarkSAGo(b *testing.B) {
	data := SimulatedAnnealing.ReadData()
	for i := 0; i < b.N; i++ {
		GetResult(data)
	}
}

//得到这组排列总里程数
func GetResult(data []SimulatedAnnealing.Data) float64 {
	sum := 0.0
	num := 0
	cityDistanceCh := make(chan float64)
	for i := 0; i < SimulatedAnnealing.GetDataNum()-1; i++ {
		getCityDistance(data[i], data[i+1], cityDistanceCh)
	}
	for num < SimulatedAnnealing.GetDataNum()-1 {
		sum += <-cityDistanceCh
		num++
	}
	return sum
}

//两个城市间的距离 Map
func getCityDistance(d1 SimulatedAnnealing.Data, d2 SimulatedAnnealing.Data, result chan float64) {
	go func() {
		result <- math.Hypot(d1.PosX-d2.PosX, d1.PosY-d2.PosY)
	}()
}
