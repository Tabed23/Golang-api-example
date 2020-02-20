package main

import (
	"gin_pratice_api/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)
var (
	port =":8080"
)
func main(){
	r := gin.Default()

	v1 := r.Group("/api/company")
	{
		v1.GET("/",router.GetAllEmployee)
		v1.POST("/",router.CreateUser)
		v1.GET("/:id", router.GetEmployee)
		v1.DELETE("/:id",router.DeleteEmployee)
		v1.PUT("/:id", router.UpdateEmployee)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
 	r.Run(port)
}
