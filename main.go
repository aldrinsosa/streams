package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func GetWeights() ([4]float64, error) {
	var weights [4]float64
	for i := range 4 {
		weight, err := strconv.ParseFloat(strings.TrimSpace(flag.Arg(i+2)), 64)
		weights[i] = weight
		if weight < 0 {
			return weights, errors.New("Heights can't be negative")
		}
		if err != nil || weight < 0 {
			return weights, err
		}
	}
	return weights, nil
}

func main() {
	flag.Parse()
	countFlags := len(flag.Args())
	if countFlags != 1 && countFlags != 6 {
		fmt.Println("Usage ./streams FILE [N WB WT WD WS]")
		return
	}
	fileName := flag.Arg(0)
	if fileName == "" {
		fmt.Println("Error with the file")
		return
	}
	if countFlags == 1 {
		fmt.Println(fileName)
		return
	}
	numberClusters, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println(err)
		return
	}
	weights, err := GetWeights()
	if err != nil {
		fmt.Println(err)
		return
	}
	println(numberClusters)
	fmt.Println(weights)
}
