package main
import (
    "encoding/csv"
    "log"
    "os"
)

type Record struct {
    Name string
    Surname string
    Number string
    LastAccess string
}

var myData = []Record{}

func main() {
    if len(os.Args)!=3 {
        log.Println("csvData input output!")
        os.Exit(1)
    }

    input := os.Args[1]
    output := os.Args[2]
    lines, err := readCSVFile(input)
    if err!=nil {
        log.Println(err)
        os.Exit(1)
    }

    for _, line := range lines {
        temp := Record{
            Name: line[0],
            Surname: line[1],
            Number: line[2],
            LastAccess: line[3],
        }
        myData = append(myData, temp)
        log.Println(temp)
    }
    err = saveCSVFile(output)
    if err!=nil {
        log.Println(err)
        os.Exit(1)
    }
}

func readCSVFile(filePath string) ([][]string, error) {
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

    return lines, nil
}

func saveCSVFile(filePath string) error {
    csvFile, err := os.Create(filePath)
    if err!=nil {
        return err
    }
    defer csvFile.Close()

    csvWriter := csv.NewWriter(csvFile)
    csvWriter.Comma = '\t'
    for _, row := range myData {
        temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
        err = csvWriter.Write(temp)
        if err!=nil {
            return err
        }
    }
    csvWriter.Flush()
    return nil
}