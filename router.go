package goxt

import (
  "github.com/gin-gonic/gin"
)
type Router struct {
  router *gin.Engine
}

func (r Router) Get(path string, listener func(Ctx)) {
  r.router.GET(path,func(c *gin.Context){
    u := TransformContext(c)
    listener(u)
  })
}

func (r Router) Post(path string, listener func(Ctx)) {
  r.router.POST(path,func(c *gin.Context){
    u := TransformContext(c)
    listener(u)
  })
}

func (r Router) Put(path string, listener func(Ctx)) {
  r.router.PUT(path,func(c *gin.Context){
    u := TransformContext(c)
    listener(u)
  })
}

func (r Router) Delete(path string, listener func(Ctx)) {
  r.router.DELETE(path,func(c *gin.Context){
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

