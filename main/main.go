package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var R = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	InitSampleData()//文件预处理
	fmt.Print(getResult())
}

func getResult() float64 {
	sum := 0.0
	data := ReadData()
	for i:=0;i<GetDataNum() - 1;i++{
		sum += getCityDistance(data[i],data[i+1])
	}
	return sum
}

//两个城市间的距离 Map
func getCityDistance(d1 Data,d2 Data) float64 {
	return math.Hypot(d1.PosX - d2.PosX,d1.PosY - d2.PosY)
}

//是否愿意接受新解
func accept(delta float64, temper float64) bool{
	if delta <= 0{
		return true
	}
	return R.Float64() <= math.Exp((-delta) / temper)
}
