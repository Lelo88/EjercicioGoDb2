package productSQL

import (
	"errors"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
)

//Aca tendria que poner algunos errores comunes. 


type Service interface{
	Create(p *domain.Product) (*domain.Product, error)
	GetByID(id int) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product,error)
}

type service struct{
	 r Repository
}

func NewSqlService(r Repository) Service {
	return &service{
		r,
	}
}

func (s *service) Create(p *domain.Product) (*domain.Product, error) {
	
	errExiste := errors.New("esto ya existe")

	if s.r.Exists(p.CodeValue){
		return p, errExiste
	}

	err := s.r.Create(p)
	if err!= nil {
		return p,err
	}

	return p, nil
	
}

func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.Read(id)
	if err != nil {
		return domain.Product{}, errors.New("este producto no existe")
	}
	return p, nil

}

func (s *service) Update(id int, p domain.Product) (domain.Product,error){

	product,_ := s.GetByID(id)

	if s.r.Exists(p.CodeValue) && product.CodeValue!=p.CodeValue{
		return domain.Product{}, errors.New("este codigo de producto ya existe")
	}

	err := s.r.Update(p)

	if err!= nil {
		return domain.Product{}, errors.New("no se puede actualizar el producto")
	}

	return p, nil
}