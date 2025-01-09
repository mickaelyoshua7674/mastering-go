package main
import (
    "fmt"
)

func main() {
    slice := make([]byte, 3)

    // Slice to array pointer
    arrayPtr := (*[3]byte)(slice)
    fmt.Println("Print array pointer:", arrayPtr)
    fmt.Printf("Data type: %T\n", arrayPtr)
    fmt.Println("arrayPtr[0]:", arrayPtr[0])

    slice2 := []int{-1, -2, -3}

    // Slice to array
    array := [3]int(slice2)
    fmt.Println("Print array contents:", array)
    fmt.Printf("Data type: %T\n", array)
}