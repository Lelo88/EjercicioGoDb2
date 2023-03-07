package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type productSqlMock struct {
	mock.Mock
}

func NewServiceProductSQLMock() *productSqlMock {
	return &productSqlMock{}
}

func (pm *productSqlMock) GetAll() ([]domain.Product, error) {
	args := pm.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (pm *productSqlMock) GetByID(id int) (domain.Product, error) {
	args := pm.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (pm *productSqlMock) Create(product domain.Product) (domain.Product, error) {
	args := pm.Called(product)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (pm *productSqlMock) Update(id int, product domain.Product) error {
	args := pm.Called(id, product)
	return args.Error(0)
}

func (pm *productSqlMock) Delete(id int) error {
	args := pm.Called(id)
	return args.Error(0)
}

func CreateServerControllerProductsTests(pm *productSqlMock) (engine *gin.Engine) {

	handler := NewProductSQLHandler(pm)

	engine = gin.Default()

	routerProduct := engine.Group("/products")
	{
		routerProduct.GET("", handler.Getall())
		routerProduct.GET("/:id", handler.GetByID())
		routerProduct.POST("", handler.Post())
		routerProduct.PUT("/:id", handler.Put())
		routerProduct.DELETE("/:id", handler.Delete())
	}

	return engine
}

func CreateReqProd(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	return req, httptest.NewRecorder()
}

//-----------------TESTS-------------------

func TestProductSQLHandler_Getall(t *testing.T) {

	type responseProduct struct {
		Data []domain.Product
	}

	products := []domain.Product{
		{Id: 1, Name: "Milanesa", Quantity: 12, IsPublished: true, Expiration: "2023/12/12", Price: 12.3},
		{Id: 2, Name: "Papas", Quantity: 5, IsPublished: false, Expiration: "2023/12/1", Price: 11.3},
	}

	productsData := responseProduct{
		Data: products,
	}

	t.Run("GetAll ok", func(t *testing.T) {
		
	})
}

func TestControllerProductCreate(t *testing.T) {

	t.Run("Should create a new product with a HTTP 201 status code", func(t *testing.T) {
		//arrange

		//act

		//assert
	})
}

//1. hay que crear un server para los tests, lo creamos mas arriba
