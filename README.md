<h1 style="text-align">
    Goxt a laravel inspired framework written in <a href="https://go.dev">Golang</a>
</h1>

## Very thanks to gin-gonic and aymerick I use gin for the http and raymond for the template parsing without his modules create this project was not possible
## Quick start

Create your first goxt project

```bash
go install github.com/BrunProgramming/goxtgen
```

only exec
```bash
goxtgen -name my-first-goxt-project -repo my-repo
```
and follow the instructions of the cli and happy coding


if you want to use the code without the project generator 
here is a basic example:

main.go
```go
package main

import (
    "github.com/BrunProgramming/goxt"
    "fmt"
)

func main() {
    router := goxt.NewRouter()
    router.Static("/style","./style")
    router.Get("/",func(c goxt.Ctx) {
        c.View("main",goxt.HbsCtx{},"")
    })
    router.Get("/:name",func(c goxt.Ctx) {
        name := c.Param("name")
        c.View("hello",goxt.HbsCtx{
            "name":name,
        },""/*this is parameter is for you want to change the default views dir put "" if you want to use the default dir*/)
    })
    fmt.Println("Server listening in http://localhost:8080")
    router.Run(":8080")
}
```

views/main.hbs
```hbs
<!DOCTYPE html>
<html>
  <head>
    <title>Goxt the best framework buffalo is noob</title>
  </head>
  <body>
   <h1>Plis enter your name below:</h1>
   <form id="form">
      <input type="text" id="input">
   </form>
    <script type="module">
      //midutrick copyright midudev©
      const $ = selector => document.querySelector(selector)
      
      $("#form").addEventListener("submit",e => {
        e.preventDefault()
        const value = $("#input").value
        history.pushState("",{},`${location.href}${value}`)
        location.reload()
      })
    </script>
  </body>
</html>
```


views/hello.hbs
```hbs
<!DOCTYPE html>
<html>
  <head>
    <title>Goxt the best framework buffalo is noob</title>
  </head>
  <body>
   <h1>Hello {{name}}</h1>
  </body>
</html>
```


