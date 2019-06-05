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
	ntheta := gradientDescent(X, y, m, theta)
	fmt.Println("Current theta: ", ntheta, " cost: ", cost(X, y, m, ntheta))
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

func individualCost(X float64, y float64, theta []float64) float64 {
	hval := X*theta[0] + theta[1]
	J := math.Pow((hval - y), 2)
	return J / float64(2)
}

func individualPredict(X float64, theta []float64) float64 {
	return X
}

func gradientDescent(X []float64, y []float64, m int, theta []float64) []float64 {
	learningRate := .001
	iterations := 1000

	currentCost := cost(X, y, m, theta)
	lowestCost := currentCost
	lowestCostTheta := []float64{0, 0}

	for i := 0; i <= iterations; i++ {

		tempTheta := []float64{0, 0}
		for j := 0; j <= m; j++ {
			iCost := individualCost(X[i], y[i], theta)
			tempTheta[0] = tempTheta[0] + float64(learningRate)*iCost*X[i]
			tempTheta[1] = tempTheta[1] + float64(learningRate)*iCost

		}
		theta[0] = tempTheta[0] / float64(m)
		theta[1] = tempTheta[1] / float64(m)
		currentCost = cost(X, y, m, theta)
		if currentCost < lowestCost {
			lowestCostTheta = theta
			lowestCost = currentCost
			break
		}
	}
	return lowestCostTheta
}
