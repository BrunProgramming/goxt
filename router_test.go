package goxt

import (
  "testing"
  "fmt"
)

func TestHelloByParam(t *testing.T) {
  router := NewRouter()
  router.Get("/:name",func(c Ctx) {
    name := c.Param("name")
    c.View("test",HbsCtx{
      "name":name,
    },""/*this is parameter is for you want to change the default views dir put "" if you want to use the default dir*/)
  })
  fmt.Println("Server listening in http://localhost:8080")
  router.Run(":8080")
}

