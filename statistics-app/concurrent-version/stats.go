package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need one or mode file paths!")
		return
	}

	var waitGroup sync.WaitGroup
	files = make(DFslice, len(os.Args))

	for i:=1; i<len(os.Args); i++ {
		waitGroup.Add(1)
		go func(x int) {
			process(os.Args[x], x)
			defer waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
}

type DataFile struct {
	Filename string
	Len int
	Minimum float64
	Maximum float64
	Mean float64
	StdDev float64
}
type DFslice []DataFile

var files DFslice

func stdDev(x []float64) (float64, float64) {
	sum := float64(0)
	for _, val := range x {
		sum = sum + val
	}

	meanValue := sum / float64(len(x))

	// Standard deviation
	var squared float64
	for i:=0; i<len(x); i++ {
		squared += math.Pow((x[i]-meanValue), 2)
	}

	standardDeviation := math.Sqrt(squared / float64(len(x)))
	return meanValue, standardDeviation
}

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}

	normalized := make([]float64, len(data))
	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}
	return normalized
}

func readFile(filePath string) ([]float64, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	values := make([]float64, 0)
	for _, line := range lines {
		tmp, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			fmt.Println("Error reading:", line[0], err)
			continue
		}
		values = append(values, tmp)
	}
	return values, nil
}

func process(file string, place int) {
	currentFile := DataFile{}
	currentFile.Filename = file

	values, err := readFile(file)
	if err != nil {
		fmt.Println("Error reading:", file, err)
		os.Exit(0)
	}

	currentFile.Len = len(values)
	currentFile.Minimum = slices.Min(values)
	currentFile.Maximum = slices.Max(values)
	meanValue, standardDeviation := stdDev(values)
	currentFile.Mean = meanValue
	currentFile.StdDev = standardDeviation

	files[place] = currentFile

	normalized := normalize(values, meanValue, standardDeviation)
	fmt.Println("Normalized length:", len(normalized), values[:5], normalized[:5])
}