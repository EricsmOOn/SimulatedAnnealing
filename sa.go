package SimulatedAnnealing

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var Temp = 10000.0    //初始化温度
var TMin = 0.0001     //温度下界限
var Delta = 0.999     //温度下降率
var SearchTime = 1000 //内循环次数

var R = rand.New(rand.NewSource(time.Now().UnixNano()))

func Sa() float64 {
	t := Temp
	data := ReadData()
	for t > TMin {
		for i := 0; i < SearchTime; i++ {
			r := GetResult(data)
			temp := changCity(data)
			rn := GetResult(temp)
			if accept(rn-r, t) {
				data = temp
				//fmt.Printf("Min : %10.5f\n", rn)
			}
			//if rn < r {
			//	data = temp
			//	fmt.Printf("Min : %10.5f\n",rn )
			//}
		}
		t = t * Delta
	}
	fmt.Print(data)
	return GetResult(data)
}

//是否愿意接受新解
func accept(delta float64, temper float64) bool {
	if delta <= 0 {
		return true
	}
	return R.Float64() <= math.Exp((-delta)/temper)
}

//得到这组排列总里程数
func GetResult(data []Data) float64 {
	sum := 0.0
	for i := 0; i < GetDataNum()-1; i++ {
		sum += getCityDistance(data[i], data[i+1])
	}
	return sum
}

//两个城市间的距离 Map
func getCityDistance(d1 Data, d2 Data) float64 {
	return math.Hypot(d1.PosX-d2.PosX, d1.PosY-d2.PosY)
}

func changCity(dataSrc []Data) []Data {
	pos1 := R.Intn(GetDataNum()-3) + 1
	pos2 := R.Intn(GetDataNum()-pos1-1) + pos1
	dataDest := make([]Data, GetDataNum())
	copy(dataDest, dataSrc)

	dataDest[pos1], dataDest[pos2] = dataDest[pos2], dataDest[pos1]

	return dataDest
}
