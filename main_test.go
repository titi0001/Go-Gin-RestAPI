package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	controller "github.com/titi0001/Go-Gin-RestAPI/controllers"
	"github.com/titi0001/Go-Gin-RestAPI/database"
	"github.com/titi0001/Go-Gin-RestAPI/models"
)

var ID int

func SetupRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas

}

func TestStatusCode(t *testing.T) {
	r := SetupRotasTeste()
	r.GET("/:nome", controller.Saudacao)
	req, _ := http.NewRequest("GET", "/thiago", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Deverian ser iguais")

	mockResponse := `{"API diz:":"E ai thiago, tudo beleza?"}`
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, mockResponse, string(responseBody))
}

func TestListarTodosOsAlunos(t *testing.T) {
	database.ConectaDb()
	CriarAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasTeste()
	r.GET("/alunos", controller.ExibeAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaDb()
	CriarAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasTeste()
	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/01234567890", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

}

func CriarAlunoMock() {
	aluno := models.Aluno{
		Nome: "Nome do Aluno teste",
		CPF:  "12345678901",
		RG:   "123456789",
	}

	database.DB.Create(&aluno)
	ID = int(aluno.ID)

}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}
