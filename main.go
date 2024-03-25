package main

import (
	"github.com/Smartdevs17/go_crud/controllers"
	"github.com/Smartdevs17/go_crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetAllPost)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run() // listen and serve on localhost:3000
}
