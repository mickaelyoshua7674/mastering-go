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
    "slices"
)

type DataFile struct {
    FileName string
    Len int
    Minimum float64
    Maximum float64
    Mean float64
    stdDev float64
}
type DFslice []DataFile

func main() {
    if len(os.Args) == 1 {
        fmt.Println("Need one or mode file paths!")
        return
    }

    files := DFslice{}
    for i:= 1; i<len(os.Args); i++ {
        file := os.Args[i]
        currentFile := DataFile{}
        currentFile.FileName = file

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
        currentFile.stdDev = standardDeviation

        files = append(files, currentFile)
    }

    sort.Sort(files)
    for _, val := range files {
        f := val.FileName
        fmt.Println(f, ":", val.Len, val.Mean, val.Minimum, val.Maximum, val.stdDev)
    }
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

func (a DFslice) Len() int {
    return len(a)
}
func (a DFslice) Less(i, j int) bool {
    return a[i].Mean < a[j].Mean
}
func (a DFslice) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

