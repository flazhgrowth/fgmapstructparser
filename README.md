# FGMAPSTRUCTPARSER [WIP]
Simple package to parse map string map(and many map type along the future. It is already on the roadmap) to designated struct
Refer the example directory on the how to use the package
Ex:
// destination struct
```
type DestStruct struct {
    ID      int64   `justmap:"id"`
    Name    string  `justmap:"name"`
}

func main() {
    // source data
    src := map[string]interface{}{
        "id":       123,
        "name":     "shn ndr nt",
    }

    dest := &DestStruct{}
    if err := New("justmap").Parse(src, dest); err != nil {
        panic(err)
    }
    fmt.Println("dest.ID: ", dest.ID)
    fmt.Println("dest.Name: ", dest.Name)
}
```

## explanation
The argument passed to the New function is the tag for struct field annotation.
As per the example above, field ID, and Name are annotated with tag `justmap`, hence,
the argument passed to function New is "justmap"