package productSQL

import (
	"testing"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct {
	ID 			int
	IsProduct 	bool
	Products 	[]domain.Product
	Err 		error
	GetProduct 	domain.Product
	Code_value	string
}

func createTestService(dumm dummyRepo) (Service){
	return &service{r: dumm}
}

func (dumm dummyRepo) Read(id int) (domain.Product, error) {
	return dumm.GetProduct, dumm.Err
}

func (dumm dummyRepo) ReadAll() ([]domain.Product, error) {
	return dumm.Products, dumm.Err
}

func (dumm dummyRepo) Create(p *domain.Product) error {
	return dumm.Err
}

func (dumm dummyRepo) Update(p domain.Product) error {
	return dumm.Err
}

func (dumm dummyRepo) Delete(id int) error {
	return dumm.Err
}

func (dumm dummyRepo) Exists(code_value string) bool {
	return dumm.IsProduct
}


var newDummyProduct = domain.Product{
	Name: "Carne",
	Quantity: 20,
	CodeValue: "abc123",
	IsPublished: true,
	Expiration: "2020-12-12",
	Price: 11.11,
}

var dummyInBBDD = domain.Product{
	Id: 1,
	Name: "Carne",
	Quantity: 20,
	CodeValue: "abc123",
	IsPublished: true,
	Expiration: "2020-12-12",
	Price: 11.11,
}

var dummyInBBDD2 = domain.Product{
	Id: 2,
	Name: "Vegetal",
	Quantity: 20,
	CodeValue: "xyz123",
	IsPublished: true,
	Expiration: "2020-12-12",
	Price: 11.11,
}

var dummyProducts = []domain.Product{
	dummyInBBDD, dummyInBBDD2,
}

func TestGetByIDOK(t *testing.T){
	ser := createTestService(dummyRepo{
		GetProduct: dummyInBBDD,
		Err: nil,
	})

	product, err := ser.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, dummyInBBDD, product)
	
}

func TestGetByIDFail(t *testing.T){
	ser := createTestService(dummyRepo{
		GetProduct: dummyInBBDD,
		Err: ErrNotFound,
	})

	product, err := ser.GetByID(2)
	assert.Error(t, err)
	assert.NotEqual(t, dummyInBBDD, product)
}

func TestGetAllOK(t *testing.T){
	ser:= createTestService(dummyRepo{
		Products: dummyProducts,
		Err: nil,
	})

	products, err := ser.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, dummyProducts, products)
}

