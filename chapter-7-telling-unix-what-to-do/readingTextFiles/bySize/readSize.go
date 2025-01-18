package main
import (
    "os"
    "io"
    "fmt"
    "strconv"
)

func main() {
    size, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }

    f, err := os.Open(os.Args[2])
    if err != nil {
        panic(err)
    }
    fmt.Println(string(readSize(f, size)))
}

func readSize(f *os.File, size int) []byte {
    buffer := make([]byte, size)
    n, err := f.Read(buffer)
    if err == io.EOF {
        return nil
    }
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return buffer[0:n]
}