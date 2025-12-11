package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type stream struct {
	streamId            int
	srcIp               string
	dstIp               string
	totalBytes          int
	streamDuration      float64
	packetCount         int
	avgInterarrivalTime float64
}

func GetWeights() ([4]float64, error) {
	var weights [4]float64
	//transform each weights from the args
	for i := range 4 {
		weight, err := strconv.ParseFloat(strings.TrimSpace(flag.Arg(i+2)), 64) // +2 is for the file and the number of clusters
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

func GetGoalClusters() (int, error) {
	numberClusters, err := strconv.Atoi(flag.Arg(1)) // 1 is to avoid the file
	if err != nil {
		return 0, err
	}
	if numberClusters < 0 {
		return 0, errors.New("Usage ./streams FILE [N WB WT WD WS]\n\tN should be a positive integer")
	}
	return numberClusters, nil
}

func GetNumberClusters(row string) (int, error) {
	splitRow := strings.Split(row, "=")
	lenSplit := len(splitRow)
	if lenSplit != 2 {
		return 0, errors.New("count=X // X is an positive integer")
	}
	numberCluster, err := strconv.Atoi(splitRow[1])
	if err != nil {
		return 0, errors.New("count=X // X is an positive integer")
	}
	if numberCluster <= 0 {
		return 0, errors.New("count=X // X is an positive integer")
	}
	if splitRow[0] != "count" {
		return 0, errors.New("count=X // X is an positive integer")
	}

	return numberCluster, nil
}

func GetSplitStream(flow string) ([]string, error) {
	f := func(c rune) bool {
		return unicode.IsSpace(c)
	}
	splitRow := strings.FieldsFunc(flow, f)
	lenSplit := len(splitRow)
	if lenSplit != 7 {
		return splitRow, errors.New("//each row in the file should be like this\n\tFLOWID SRC_IP DST_IP TOTAL_BYTES FLOW_DURATION PACKET_COUNT AVG_INTERARRIVAL ")
	}
	return splitRow, nil
}

func GetStreams(rows []string, streams *[]stream, numberClusters int) error {
	for i := range numberClusters {
		var stream stream
		splitRow, err := GetSplitStream(rows[i])
		if err != nil {
			return err
		}
		streamId, err := strconv.Atoi(splitRow[0])
		if err != nil {
			return err
		}
		stream.streamId = streamId
		*streams = append(*streams, stream)
	}
	return nil
}

func main() {
	//check the args
	flag.Parse()
	countArgs := len(flag.Args())
	if countArgs != 1 && countArgs != 6 {
		fmt.Println("Usage ./streams FILE [N WB WT WD WS]")
		return
	}
	//get the file name from the args
	fileName := flag.Arg(0)
	if fileName == "" {
		fmt.Println("Error with the file")
		return
	}

	//open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//get the amount of clusters in the file
	scanner.Scan()
	countRow := scanner.Text()
	numberClusters, err := GetNumberClusters(countRow)
	if err != nil {
		fmt.Println(err)
		return
	}

	//get each flow
	var rowsFile []string
	for scanner.Scan() {
		rowsFile = append(rowsFile, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var streams []stream

	err = GetStreams(rowsFile, &streams, numberClusters)
	if err != nil {
		fmt.Println(err)
		return
	}

	//when there's no weights
	if countArgs == 1 {
		fmt.Println(numberClusters)
		fmt.Println(rowsFile)
		return
	}

	//obtain the desire amount of clusters
	goalClusters, err := GetGoalClusters()
	if err != nil {
		fmt.Println(err)
		return
	}

	//get the weights from the args
	weights, err := GetWeights()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(goalClusters)
	fmt.Println(weights)
}
