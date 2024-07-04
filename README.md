<h1 style="text-align">
    Goxt a laravel inspired framework written in <a href="https://go.dev">Golang</a>
    <br>Very thanks to gin-tonic and aymerick I use gin for the http and raymond for the template parsing without his modules create this project was not possible
</h1>

## Quick start

install 
> [!IMPORTANT]
> the module is in diapers in the future was a project generator but in the present you use the module like any other http module for golang

    $ go get github.com/BrunProgramming/goxt

This is a example for a basic use

main.go
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

test.hbs
```hbs
<!DOCTYPE html>
<html>
  <head>
    <title>{{title}}</title>
  </head>
  <body>
   <h1>{{title}}</h1>
  </body>
</html>
```

