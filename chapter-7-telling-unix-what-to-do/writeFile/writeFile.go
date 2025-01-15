package main
import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
    buffer := []byte("Data to write\n")

    f1, err := os.Create("f1.txt")
    if err != nil {
        fmt.Println("Can not create file", err)
        return
    }
    defer f1.Close()
    fmt.Fprintf(f1, string(buffer))

    f2, err := os.Create("f2.txt")
    if err != nil {
        fmt.Println("Can not create file", err)
        return
    }
    defer f2.Close()
    n, err := f2.WriteString(string(buffer))
    fmt.Printf("Wrote %d bytes\n", n)

    f3, err := os.Create("f3.txt")
    if err != nil {
        fmt.Println("Can not create file", err)
        return
    }
    defer f3.Close()

    w := bufio.NewWriter(f3)
    n, err = w.WriteString(string(buffer))
    fmt.Printf("Wrote %d bytes\n", n)
    w.Flush()

    f := "f4.txt"
    f4, err := os.Create(f)
    if err != nil {
        fmt.Println("Can not create file", err)
        return
    }
    defer f4.Close()
    for i:=0; i<5; i++ {
        n, err = io.WriteString(f4, string(buffer))
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("Wrote %d bytes\n", n)
    }

    // Append to a file
    f4, err = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Can not create file", err)
        return
    }
    defer f4.Close()
    n, err = f4.Write([]byte("Put some mode data at the end.\n"))
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Wrote %d bytes\n", n)
}