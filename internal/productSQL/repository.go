package productSQL

import (
	"database/sql"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	_"github.com/go-sql-driver/mysql"
)

type Repository interface {
	Create(product domain.Product) (int, error)
	Read()([]*domain.Product, error)
}


type repository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) Repository{
	return &repository{
		db,
	}
}

func (r *repository) Create(product domain.Product) (int, error) {

	query:=`INSERT INTO products(name, quantity, code_value, is_published, expiration, price) 
			VALUES (?,?,?,?,?,?);`

	statement, err := r.db.Prepare(query)
	if err!= nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(product.Name, 
								product.Quantity, 
								product.CodeValue, 
								product.IsPublished, 
								product.Expiration, 
								product.Price)

	if err!= nil {
		return 0, err	
	}
	
	id, err := result.LastInsertId()
	if err!= nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Read()([]*domain.Product, error) {

	return ([]*domain.Product{}), nil
}