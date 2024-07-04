package goxt

import (
  "github.com/gin-gonic/gin"
  "github.com/aymerick/raymond"
  "os"
)
type Context struct {
  context *gin.Context
}

type HbsContext map[string]interface{}

func (c Context) view(view string,ctx HbsContext) {
  c.context.Header("Content-Type","text/html")
  cwd,_ := os.Getwd()
  file,err := os.ReadFile(cwd+"/"+view+".hbs")
  if err != nil {
    c.context.String(500,err.Error())
  }
  result,_ := raymond.Render(string(file), ctx)
  c.context.String(200,result)
}

func TransformContext(c *gin.Context) Context {
  return Context{
    context:c,
  }
}


type Router struct {
  router *gin.Engine
}

func (r Router) Get(path string,listener func(Context)) {
  r.router.GET(path,func(c *gin.Context){
    u := TransformContext(c)
    listener(u)
  })
}

func (r Router) Run(addr string) {
  r.router.Run(addr)
}



func NewRouter() Router {
  router := gin.Default()
  return Router{
    router:router,
  }
}

