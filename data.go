package SimulatedAnnealing

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var DataFileName = "4.kroA100.tsp.txt"

type Data struct {
	CityNum int
	PosX    float64
	PosY    float64
}

func (d Data) String() string {
	return fmt.Sprintf("[CityNum:%3d, PosX:%5f, PosY:%5f]\n", d.CityNum, d.PosX, d.PosY)
}

var data []Data

func ReadData() []Data {
	return data
}

func InitSampleData() {
	data = make([]Data, 0)
	var d Data
	file, _ := GetFileLines()
	for _, s := range file {
		sp := strings.Split(s, " ")
		sp0, _ := strconv.Atoi(sp[0])
		sp1, _ := strconv.Atoi(sp[1])
		sp2, _ := strconv.Atoi(sp[2])
		d = Data{sp0, float64(sp1), float64(sp2)}
		data = append(data, d)
	}
}

func GetFileLines() ([]string, error) {
	var result []string
	filePath := "./data/" + DataFileName
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
	return result[6 : len(result)-1], nil
}

func GetDataNum() int {
	data := ReadData()
	return len(data)
}
