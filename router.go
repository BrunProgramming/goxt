package goxt

import (
  "github.com/gin-gonic/gin"
  
  "os/exec"
  "runtime"
)

func open(url string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd.exe"
        args = []string{"/c","start"}
    case "darwin":
        cmd = "open"
    default: // "linux", "freebsd", "openbsd", "netbsd"
        cmd = "xdg-open"
    }
    args = append(args, url)
    return exec.Command(cmd, args...).Run()
}

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

func(r Router) Use(middleware ...gin.HandlerFunc) {
  r.Router.Use(middleware...)
}

func (r Router) Run(addr string) {
  go open("http://localhost"+addr)
  r.Router.Run(addr)
}

func NewRouter() Router {
  router := gin.New()
  return Router{
    Router:router,
  }
}

