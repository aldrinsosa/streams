package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
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

func GetNumberClusters() (int, error) {
	numberClusters, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		return 0, err
	}
	if numberClusters < 0 {
		return 0, errors.New("Usage ./streams FILE [N WB WT WD WS]\n\tN should be a positive integer")
	}
	return numberClusters, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//check the args
	flag.Parse()
	countArgs := len(flag.Args())
	if countArgs != 1 && countArgs != 6 {
		fmt.Println("Usage ./streams FILE [N WB WT WD WS]")
		return
	}
	fileName := flag.Arg(0)
	if fileName == "" {
		fmt.Println("Error with the file")
		return
	}

	file, err := os.ReadFile(fileName)
	check(err)

	if countArgs == 1 {
		fmt.Println(string(file))
		return
	}
	numberClusters, err := GetNumberClusters()
	if err != nil {
		fmt.Println(err)
		return
	}

	weights, err := GetWeights()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(numberClusters)
	fmt.Println(weights)
}
