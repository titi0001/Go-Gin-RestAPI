package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/titi0001/Go-Gin-RestAPI/database"
	"github.com/titi0001/Go-Gin-RestAPI/models"
)

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleze?",
	})
}

func ExibeAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorId(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	c.JSON(http.StatusOK, aluno)
}