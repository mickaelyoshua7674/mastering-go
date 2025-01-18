package main
import (
    "fmt"
    "bufio"
    "io"
    "os"
    "regexp"
)

func main() {
    wordByWord("csv.data")
}

func wordByWord(filePath string) error {
    f, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer f.Close()

    r := bufio.NewReader(f)
    re := regexp.MustCompile(`^[\s]+`)
    for {
        line, err := r.ReadString('\n')
        if err == io.EOF {
            if len(line) != 0 {
                words := re.FindAllString(line, -1)
                for i:=0; i<len(words); i++ {
                    fmt.Println(words[i])
                }
            }
            break
        } else if err != nil {
            fmt.Printf("error reading file %s", err)
            return err
        }
        words := re.FindAllString(line, -1)
        for i:=0; i<len(words); i++ {
            fmt.Println(words[i])
        }
    }
    return nil
}