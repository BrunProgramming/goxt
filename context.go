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

func (c Ctx) View(view string,ctx HbsCtx,viewDir string) {
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

func TransformContext(c *gin.Context) Ctx {
  return Ctx{
    context:c,
  }
}

