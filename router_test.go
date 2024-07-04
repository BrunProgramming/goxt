package goxt

import ( 
  "testing"
  "fmt" 
)

func TestHelloWorld(t *testing.T) {
  router := NewRouter()
  router.Get("/",func(c Context) { 
    c.view("test",HbsContext{
      // the handlebars file is not submit into the repository if you wanna testing you need to create the test.hbs file
      "title":"Test",
    })
})
  fmt.Println("Server listening in http://localhost:8080")
  router.Run(":8080")
} 

