package main
import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("Please provide an argument!")
    }

    file := arguments[1]
    path := os.Getenv("PATH")
    pathSplit := filepath.SplitList(path)
    for _, directory := range pathSplit {
        fullPath := filepath.Join(directory, file)

        // Using error to verify if file exists
        fileInfo, err := os.Stat(fullPath)
        if err!=nil {
            continue
        }

        mode := fileInfo.Mode()
        // Checking if the file is a regular file
        if !mode.IsRegular() {
            continue
        }

        // If true it is an executable
        if mode&0111 != 0 {
            fmt.Println(fullPath)
            return
        }
    }
}