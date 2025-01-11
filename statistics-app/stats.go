package main
import (
    "fmt"
    "math"
    "math/rand"
    "os"
    "strconv"
    "sort"
)

const stdMin float64 = -10
const stdMax float64 = 10
const stdNumberValues = 10

func main() {
    arguments := os.Args
    var values []float64
    if len(arguments) == 1 {
        fmt.Println("No arguments given!")
        fmt.Printf("Generating %v random values from %v to %v (not included).\n\n", stdNumberValues, stdMin, stdMax)

        for i:=0; i<stdNumberValues; i++ {
            values = append(values, randomFloat(stdMin, stdMax))
        }
    } else {
        for i:=1; i<len(arguments); i++ {
            // Using error to verify if the argument can be parsed to float.
            n, err := strconv.ParseFloat(arguments[i], 64)
            if err!=nil {
                continue
            }
            values = append(values, n)
        }
    }

    sort.Float64s(values)
    nValues := len(values)

    // Number of values, Min and Max
    fmt.Println("Number of values:", nValues)
    fmt.Printf("Min: %.4f\n", values[0])
    fmt.Printf("Max: %.4f\n", values[nValues-1])

    // Mean value
    var sum float64
    for _, value := range values {
        sum += value
    }
    meanValue := sum / float64(nValues)
    fmt.Printf("Mean value: %.5f\n", meanValue)

    // Standard deviation
    var squared float64
    for _, value := range values {
        squared = squared + math.Pow((value-meanValue), 2)
    }
    standardDeviation := math.Sqrt(squared / float64(nValues))
    fmt.Printf("Standard deviation: %.5f\n", standardDeviation)

    normalized := normalize(values, meanValue, standardDeviation)
    fmt.Println("Normalized:", normalized)
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