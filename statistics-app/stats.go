package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "math"
    "math/rand"
    "os"
    "sort"
    "strconv"
)

const stdMin float64 = -10
const stdMax float64 = 10
const stdNumberValues = 10

func main() {
    var values []float64
    if len(os.Args) == 1 {
        fmt.Println("Need onde argument with the file path!")
        fmt.Printf("Generating %v random values from %v to %v (not included).\n\n", stdNumberValues, stdMin, stdMax)

        for i:=0; i<stdNumberValues; i++ {
            values = append(values, randomFloat(stdMin, stdMax))
        }
    } else {
        file := os.Args[1]
        var err error
        values, err = readFile(file)
        if err!=nil {
            log.Println("Error reading:", file, err)
            os.Exit(0)
        }
    }

    sort.Float64s(values)
    nValues := len(values)

    // Number of values, Min and Max
    fmt.Println("Number of values:", nValues)
    fmt.Printf("Min: %.4f\n", values[0])
    fmt.Printf("Max: %.4f\n", values[nValues-1])

    meanValue, standardDeviation := stdDev(values)

    normalized := normalize(values, meanValue, standardDeviation)
    fmt.Println("Normalized:", normalized)
}

func stdDev(x []float64) (float64, float64) {
    nValues := len(x)
    sum := 0.
    for _, val := range x {
        sum += val
    }

    meanValue := sum / float64(nValues)
    fmt.Printf("Mean value: %.5f\n", meanValue)

    var squared float64
    for i:=0; i<nValues; i++ {
        squared += math.Pow((x[i]-meanValue), 2)
    }
    standardDeviation := math.Sqrt(squared / float64(nValues))
    return meanValue, standardDeviation
}

func normalize(data [] float64, mean, stdDev float64) []float64 {
    if stdDev == 0 {
        return data
    }

    normalized := make([]float64, len(data))
    for i, val := range data {
        normalized[i] = math.Floor((val-mean) / stdDev*10000) / 10000
    }
    return normalized
}

func randomFloat(min, max float64) float64 {
    return min + rand.Float64()*(max-min)
}

func readFile(filePath string) ([]float64, error) {
    _, err := os.Stat(filePath)
    if err!=nil {
        return nil, err
    }

    f, err := os.Open(filePath)
    if err!=nil {
        return nil, err
    }
    defer f.Close()

    lines, err := csv.NewReader(f).ReadAll()
    if err!=nil {
        return nil, err
    }

    values := make([]float64, 0)
    for _, line := range lines {
        tmp, err := strconv.ParseFloat(line[0], 64)
        if err!=nil {
            log.Println("Error reading:", line[0], err)
            continue
        }
        values = append(values, tmp)
    }
    return values, nil
}