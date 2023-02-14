package handler

import (
	"testing"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type errorResponse struct {
	Code string 	`json:"code"`
	Message string 	`json:"message"`
}

type productSqlMock struct{
	mock.Mock
}

func NewServiceProductSQLMock() *productSqlMock{
	return &productSqlMock{}
}

func (pm *productSqlMock) GetAll() ([]domain.Product, error) {
	args := pm.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func CreateServerControllerProductsTests() (engine *gin.Engine){
	return engine
}

func TestControllerProductCreate(t *testing.T){

	t.Run("Should create a new product with a HTTP 201 status code", func(t *testing.T){
		//arrange 

		//act

		//assert
	})
}

//1. hay que crear un server para los tests, lo creamos mas arriba