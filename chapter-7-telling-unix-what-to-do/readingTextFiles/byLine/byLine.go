package main
import (
    "fmt"
    "os"
    "io"
    "bufio"
)

func main() {
    lineByLine("csv.data")
}

func lineByLine(filePath string) error {
    f, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer f.Close()

    r := bufio.NewReader(f)
    for {
        line, err := r.ReadString('\n')
        if err == io.EOF {
            if len(line) != 0 {
                fmt.Println(line)
            }
            break
        }
        if err != nil {
            fmt.Printf("Error reading file %s", err)
            return err
        }

        fmt.Print(line)
    }
    return nil
}