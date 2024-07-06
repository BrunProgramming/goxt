package goxt

import (
  "github.com/gin-gonic/gin"
  "github.com/aymerick/raymond"
  "os"
)

type Ctx struct {
  context *gin.Context
}

type HbsCtx map[string]interface{}

func (c Ctx) View(view string,ctx interface{},viewDir string) {
  var dir string
  cwd,_ := os.Getwd()
  if viewDir == "" { dir=cwd+"/views/"+view+".hbs" } else { dir=viewDir }
  c.context.Header("Content-Type","text/html")
  file,err := os.ReadFile(dir)
  if err != nil {
    c.context.String(500,err.Error())
  }
  result,_ := raymond.Render(string(file), ctx)
  c.context.String(200,result)
}

func (c Ctx) Param(key string) string {
  return c.context.Param(key)
}

func (c Ctx) Query(key string, defaultValue string) string {
  return c.context.DefaultQuery(key,defaultValue)
}

func (c Ctx) Set(key string, value string) {
  c.context.Set(key,value)
}

func (c Ctx) Get(key string) any {
  return c.context.MustGet(key)
}
// support for gin middleware nah is only a excuse because i dont make a form of creating middlewares with my Handler

func TransformContext(c *gin.Context) Ctx {
  return Ctx{
    context:c,
  }
}

