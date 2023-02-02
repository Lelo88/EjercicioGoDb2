package productSQL

import (
	

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
)

type dummyRepo struct {
	ID 			int
	IsProduct 	bool
	Products 	[]domain.Product
	Err 		error
	GetProduct 	domain.Product
	code_value	string
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


