package main

import (
	"github.com/Smartdevs17/go_crud/initializers"
	"github.com/Smartdevs17/go_crud/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
