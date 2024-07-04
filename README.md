<h1 style="text-align">
    Goxt a laravel inspired framework written in [golang](https://golang.org)
    Very thanks to gin-tonic and aymerick but without his module Gin for the http and raymond for the template parsing this project was not possible
</h1>
<hr>
## Quick start

install 
> [!IMPORTANT]
> the module is in diapers in the future was a project generator but in the present you use the module like any other http module for golang

    $ go get github.com/BrunProgramming/goxt

This is a example for a basic use
```go
package main

import (
    "github.com/BrunProgramming/goxt"
    "fmt"
)

func main() {
  router := goxt.NewRouter()
  router.Get("/",func(c goxt.Context) { 
    c.view("test",goxt.HbsContext{
      "title":"Test",
    })
  })
  fmt.Println("Server listening in http://localhost:8080")
  router.Run(":8080")
}
```
