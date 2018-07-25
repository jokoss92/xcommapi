package route

import (
	"xcommapi/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(Cors())

	v1 := router.Group("api/v1")
	{
		v1.GET("/ping", GetPing)
	}

	department := router.Group("api/department")
	{
		department.GET("/", controller.GetAllDepartment)
		department.GET("/:id", controller.GetDepartmentByID)
		department.POST("/", controller.CreateNewDepartment)
		department.PUT("/", controller.UpdateExistingDepartment)
		department.DELETE("/:id", controller.DeleteExistingDepartment)
	}

	return router
}

func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}