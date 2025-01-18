package main
import (
    "fmt"
    "os"
    "encoding/json"
)

var JSONrecord = `{
    "Flag": true,
    "Array": ["a","b","c"],
    "Entity": {
    "a1": "b1",
    "a2": "b2",
    "Value": -456,
    "Null": null
    },
    "Message": "Hello Go!"
}`

func main() {
    if len(os.Args) == 1 {
    fmt.Println("*** Using default JSON record ***")
    } else {
        JSONrecord = os.Args[1]
    }
    JSONMap := make(map[string]any)
    err := json.Unmarshal([]byte(JSONrecord), &JSONMap)
    if err != nil {
        fmt.Println(err)
        return
    }

    exploreMap(JSONMap)
    typeSwitch(JSONMap)
}

func typeSwitch(m map[string]any) {
    for k, v := range m {
        switch c := v.(type) {
        case string:
            fmt.Println("Is a string!", k, c)
        case float64:
            fmt.Println("Is a float64!", k, c)
        case bool:
            fmt.Println("Is a Boolean!", k, c)
        case map[string]any:
            fmt.Println("Is a map!", k, c)
            typeSwitch(v.(map[string]any))
        default:
            fmt.Printf("...Is %v: %T!\n", k, c)
        }
    }
}

func exploreMap(m map[string]any) {
    for k, v := range m {
        embMap, ok := v.(map[string]any)
        if ok {
            fmt.Printf("{\"%v\": \n", k)
            exploreMap(embMap)
            fmt.Printf("}\n")
        } else {
            fmt.Printf("%v: %v\n", k, v)
        }
    }
}