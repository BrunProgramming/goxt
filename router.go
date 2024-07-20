package goxt

import (
  "github.com/gin-gonic/gin"
)

type Router struct {
  Router *gin.Engine
}

func (r Router) Get(path string, listener func(Ctx)) {
  r.Router.GET(path,func(c *gin.Context){
    u := TransformContext(c)
    listener(u)
  })
}

func (r Router) Post(path string, listener func(Ctx)) {
  r.Router.POST(path,func(c *gin.Context){
    u := TransformContext(c)
    listener(u)
  })
}

func (r Router) Put(path string, listener func(Ctx)) {
  r.Router.PUT(path,func(c *gin.Context){
    u := TransformContext(c)
    listener(u)
  })
}

func (r Router) Delete(path string, listener func(Ctx)) {
  r.Router.DELETE(path,func(c *gin.Context){
    u := TransformContext(c)
    listener(u)
  })
}

func (r Router) Static(route string, dir string) {
  r.Router.Static(route,dir)
}

func(r Router) Use(middleware ...gin.HandlerFunc) {
  r.Router.Use(middleware...)
}

func (r Router) Run(addr string) {
  r.Router.Run(addr)
}

func NewRouter() Router {
  router := gin.Default()
  return Router{
    Router:router,
  }
}

