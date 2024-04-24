package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/titi0001/Go-Gin-RestAPI/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:nome", controller.Saudacao)
	r.GET("/alunos", controller.ExibeAlunos)
	r.POST("/alunos", controller.CriarAluno)
	r.GET("/alunos/:id", controller.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	r.PATCH("/alunos/:id", controller.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCPF)
	r.GET("/index", controller.ExibePaginaIndex)
	r.Run(":9000")
}
