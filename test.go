package SimulatedAnnealing

import (
	"strconv"
	"strings"
)

var cityNum []int

func ReadCityNum() []int {
	return cityNum
}

func GetBest() float64 {
	InitTestData()
	data := ReadData()
	dataDest := make([]Data, 0)
	for _, i := range cityNum {
		for _, d := range data {
			if d.CityNum == i {
				dataDest = append(dataDest, d)
			}
		}
	}
	return GetResult(dataDest)
}

func InitTestData() {
	data := make([]int, 0)
	str := strings.Split(DataFileName, ".")
	file, _ := GetFileLines("./data/" + str[0] + "." + str[1] + ".opt.tour.txt")
	file = file[5 : len(file)-1]
	for _, s := range file {
		i, _ := strconv.Atoi(s)
		data = append(data, i)
	}
	cityNum = append(data[:len(data)-1], data[0])
}
