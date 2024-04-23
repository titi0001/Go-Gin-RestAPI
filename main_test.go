package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaDb()
	CriarAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasTeste()
	r.GET("/alunos/:id", controller.BuscaAlunoPorId)

	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Nome do Aluno teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)

}

 func TestDeletaAlunoHandler(t * testing.T){
	database.ConectaDb()
	CriarAlunoMock()

	r := SetupRotasTeste()
	r.DELETE("/alunos/:id", controller.DeletaAluno)

	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
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
