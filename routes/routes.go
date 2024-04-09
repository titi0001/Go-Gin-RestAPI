package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/titi0001/Go-Gin-RestAPI/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/:nome", controller.Saudacao)
	r.GET("/alunos", controller.ExibeAlunos)
	r.POST("/alunos", controller.CriarAluno)
	r.GET("/alunos/:id", controller.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	r.Run(":9000")
}
