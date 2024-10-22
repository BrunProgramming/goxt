package goxt

import (
  "testing"
  "context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func TestHelloByParam(t *testing.T) {
  router := NewRouter()
  srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Router,
	}
  router.Static("/style","style")

  router.Get("/",func(c Ctx) {
    c.View("main",HbsCtx{},"")
  })
  router.Get("/:name",func(c Ctx) {
    name := c.Param("name")
    c.View("hello",HbsCtx{
      "name":name,
    },""/*this is parameter is for you want to change the default views dir put "" if you want to use the default dir*/)
  })
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

