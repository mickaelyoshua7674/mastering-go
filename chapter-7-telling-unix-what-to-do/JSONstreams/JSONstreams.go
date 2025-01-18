package main
import (
    "bytes"
    "encoding/json"
    "fmt"
    "math/rand"
)

type Data struct {
    Key string `json:"key"`
    Val int `json:"value"`
}

var DataRecords []Data
var min = 0
var max = 26

func main() {
    var i int
    var t Data
    for i=0; i<2; i++ {
        t = Data{
            Key: getString(5),
            Val: random(1, 100),
        }
        DataRecords = append(DataRecords, t)
    }

    buf := new(bytes.Buffer)
    encoder := json.NewEncoder(buf)
    err := Serialize(encoder, DataRecords)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Print("After Serialize:", buf)

    decoder := json.NewDecoder(buf)
    var temp []Data
    err = DeSerialize(decoder, &temp)
    fmt.Println("After DeSerialize:")
    for i, v := range temp {
        fmt.Println(i, v)
    }
}

func random(min, max int) int {
    return rand.Intn(max-min) + min
}

func getString(l int64) string {
    startChar := "A"
    temp := ""
    var i int64 = 1
    for {
        myRand := random(min, max)
        newChar := string(startChar[0] + byte(myRand))
        temp += newChar
        if i == l {
            break
        }
        i++
    }
    return temp
}

func DeSerialize(e *json.Decoder, slice any) error {
    return e.Decode(slice)
}
func Serialize(e *json.Encoder, slice any) error {
    return e.Encode(slice)
}