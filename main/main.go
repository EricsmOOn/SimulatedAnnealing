package main

import (
	"fmt"
	. "github.com/EricsmOOn/SimulatedAnnealing"
)

func main() {
	InitSampleData() //文件预处理
	fmt.Printf("模拟退火启动:\n文件名称:%s\n初始化温度:%.1f\n温度下界限:%f\n温度下降率:%f\n内循环次数:%d\n最短路径:", DataFileName, Temp, TMin, Delta, SearchTime)
	//开始模拟退火
	r := Sa()
	best := GetBest()
	fmt.Printf("\n退火解为: %7.3f ,最优解为: %7.3f ,误差率为: %4.2f", r, best, (r-best)/best*100)
	fmt.Print("%")
}
