package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) bool {
	if e != nil {
		return true
	} else {
		return false
	}
}

func loadData(fileName string) ([]float64, []float64, int) {
	var Xslice []float64
	var yslice []float64
	var length int

	dat, _ := ioutil.ReadFile(fileName)
	bytesReader := bytes.NewReader(dat)
	bufReader := bufio.NewReader(bytesReader)

	for {
		length++
		data, _, err := bufReader.ReadLine()
		if check(err) == true {
			break
		} else {
			splitData := strings.Split(string(data), ",")
			Xdat, _ := strconv.ParseFloat(splitData[0], 64)
			ydat, _ := strconv.ParseFloat(splitData[1], 64)
			Xslice = append(Xslice, Xdat)
			yslice = append(yslice, ydat)
		}

	}

	return Xslice, yslice, length
}

func main() {
	X, y, m := loadData("data.txt")
	theta := []float64{0, 0}

	fmt.Println(cost(X, y, m, theta))
}

func cost(X []float64, y []float64, m int, theta []float64) float64 {
	var J float64

	for i := range X {
		hval := X[i]*theta[0] + theta[1]
		J += math.Pow((hval - y[i]), 2)
	}
	J = J / float64(2*m)
	return J
}

func gradientDescent() {
	//need to add gradient descent
	//will do later
	//probably
}
