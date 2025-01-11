package main
import (
    "os"
    "fmt"
    "regexp"
)

func main() {
    args := os.Args
    if len(args)==1 {
        fmt.Println("Must give one argument!")
        return
    }

    s := args[1]
    ret := matchNameSur(s)
    fmt.Println(ret)
}

func matchNameSur(s string) bool {
    t := []byte(s)
    re := regexp.MustCompile(`^[A-Z][a-z]*$`)
    return re.Match(t)
}