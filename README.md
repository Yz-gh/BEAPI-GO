## Install
```
go get -u github.com/yz-gh/beapi-go
```

## Example
```go
package main
import (
    "fmt"

    "github.com/tidwall/gjson"
    beapi "github.com/yz-gh/beapi-go"
)

func main(){
    cl := beapi.DefaultClient
    
    //Storage Upload
    uploadedFileLink := gjson.Get(cl.FileUpload("img.jpg"), "result").String()
    fmt.Println(uploadedFileLink)
    
    //AlphaCoders
    ACList := gjson.Get(cl.AlphaCoders("naruto"), "result.#.url").Array()
    for _, v := range ACList{
        fmt.Println(v.String())
    }

    //Get Os Name List
    for i, osname := range beapi.OsNameList{
        fmt.Println(i, osname)
    }
}
