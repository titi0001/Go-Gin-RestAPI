package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/titi0001/Go-Gin-RestAPI/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controller.ExibeAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.Run(":9000")
}
