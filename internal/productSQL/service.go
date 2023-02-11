package productSQL

import (
	"errors"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
)

//Aca tendria que poner algunos errores comunes. 

var (
	ErrDatabaseNotFound = errors.New("database not found")
	ErrNotFound = errors.New("product not found")
)


type Service interface{
	Create(p domain.Product) (domain.Product, error)
	GetByID(id int) (domain.Product, error)
	Update(id int, p domain.Product) (error)
	GetAll() ([]domain.Product, error)
	Delete(id int) (error)
}

type service struct{
	 r Repository
}

func NewSqlService(r Repository) Service {
	return &service{
		r,
	}
}

func (s *service) Create(p domain.Product) (domain.Product, error) {
	
	errExiste := errors.New("esto ya existe")

	if s.r.Exists(p.CodeValue){
		return p, errExiste
	}

	err := s.r.Create(p)
	if err!= nil {
		return p,ErrDatabaseNotFound
	}

	return p, nil
	
}

func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.Read(id)
	if err != nil {
		return domain.Product{}, ErrNotFound
	}
	return p, nil

}

func (s *service) Update(id int, p domain.Product) (error){

	product,_ := s.GetByID(id)

	if s.r.Exists(p.CodeValue) && product.CodeValue!=p.CodeValue{
		return errors.New("este codigo de producto ya existe")
	}

	if p.Name != "" {
		product.Name = p.Name
	}
	if p.CodeValue != "" {
		product.CodeValue = p.CodeValue
	}
	if p.Expiration != "" {
		product.Expiration = p.Expiration
	}
	if p.Quantity > 0 {
		product.Quantity = p.Quantity
	}
	if p.Price > 0 {
		product.Price = p.Price
	}


	err := s.r.Update(product)

	if err!= nil {
		return  errors.New("no se puede actualizar el producto")
	}

	return  nil
}

func (s *service) GetAll() ([]domain.Product, error){

	products , err := s.r.ReadAll()
	if err!= nil {
		return []domain.Product{}, ErrDatabaseNotFound
	}

	return products, nil
}

func (s *service) Delete(id int) error{
	
	err := s.r.Delete(id)

	switch err {
		case ErrNotFound:
			return ErrNotFound
		case ErrDatabaseNotFound:
			return ErrDatabaseNotFound
	}

	return nil
}