package main
import (
    "encoding/json"
    "fmt"
)

type UseAll struct {
    Name string `json:"username"`
    Surname string `json:"surname"`
    Year int `json:"created,omitempty"`
    Pass string `json:"-"`
}

func main() {
    useall := UseAll{Name:"Mike", Pass: "1234"}
    t, err := json.Marshal(&useall)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Value %s\n", t)
    }

    str := `{"username":"M.", "surname":"Ts", "created":2024}`
    jsonRecord := []byte(str)

    // Create a structure variable to store the result
    temp := UseAll{}
    err = json.Unmarshal(jsonRecord, &temp)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Data type: %T with value %v\n", temp, temp)
    }
}