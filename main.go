package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"main.go/cantroller"
	"main.go/models"
)

func main() {
	r := gin.Default()

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	models.ConnectDataBase()
	//r.POST("/login",services.Login)
	r.GET("/getuser", cantroller.ShowUserData)
	r.PATCH("/update/:id", cantroller.Updatedata)
	//r.GET("/get",   cantroller.findDetails)
	r.POST("/post", cantroller.Createdata)
	r.DELETE("/delete/:id", cantroller.DeleteData)
	r.GET("/Querystring_parameters/:first_name", cantroller.Querystring_parameters)
	r.Run()
}
