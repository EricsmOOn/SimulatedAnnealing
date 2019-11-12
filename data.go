package SimulatedAnnealing

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//var DataFileName = "1.att48.tsp.txt"
var DataFileName = "3.eil76.tsp.txt"

//var DataFileName = "4.kroA100.tsp.txt"

type Data struct {
	CityNum int
	PosX    float64
	PosY    float64
}

func (d Data) String() string {
	//return fmt.Sprintf("[CityNum:%3d, PosX:%1f, PosY:%1f]\n", d.CityNum, d.PosX, d.PosY)
	return fmt.Sprintf("%2d ->", d.CityNum)
}

var data []Data

func ReadData() []Data {
	return data
}

func InitSampleData() {
	data = make([]Data, 0)
	var d Data
	file, _ := GetFileLines("./data/" + DataFileName)
	file = file[6 : len(file)-1]
	for _, s := range file {
		sp := strings.Split(s, " ")
		sp0, _ := strconv.Atoi(sp[0])
		sp1, _ := strconv.Atoi(sp[1])
		sp2, _ := strconv.Atoi(sp[2])
		d = Data{sp0, float64(sp1), float64(sp2)}
		data = append(data, d)
	}
	data = append(data, data[0])
}

func GetFileLines(filePath string) ([]string, error) {
	var result []string
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file: %v error: %v", filePath, err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	return result, nil
}

func GetDataNum() int {
	data := ReadData()
	return len(data)
}
