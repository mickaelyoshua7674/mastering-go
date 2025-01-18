package main
import (
    "fmt"
)

func main() {
    aSlice := []float64{}
    fmt.Println(aSlice, len(aSlice), cap(aSlice))

    aSlice = append(aSlice, 1234.56)
    aSlice = append(aSlice, -34.)
    fmt.Println(aSlice, "with length", len(aSlice))

    t := make([]int, 4)
    for i:=0; i<len(t); i++ {
        t[i] = -(i+1)
    }
    t = append(t, -5)
    fmt.Println(t)

    twoD := [][]int{{1,2,3}, {4,5,6}}
    for _, i := range twoD {
        for _, j := range i {
            fmt.Print(j, " ")
        }
        fmt.Println()
    }

    make2D := make([][]int, 2)
    fmt.Println(make2D)
    make2D[0] = []int{1,2,3,4}
    make2D[1] = []int{-1,-2,-3,-4}
    fmt.Println(make2D)
}