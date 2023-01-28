package productSQL

import (
	"errors"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
)

type Service interface{
	Create(p domain.Product) (domain.Product, error)
	GetByID(id int) (domain.Product, error)
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

	if s.r.Exists(p.CodeValue) {
		return domain.Product{}, errExiste
	}

	err := s.r.Create(p)
	
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
	
}

func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.Read(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil

}
