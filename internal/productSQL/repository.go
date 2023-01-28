package productSQL

import (
	"database/sql"
	"errors"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	 "github.com/go-sql-driver/mysql"
)

type Repository interface {
	Create(product domain.Product) (int, error)
	Read(id int) (domain.Product, error)
	Exists(code_value string) bool
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

	result, err := statement.Exec(
								product.Name, 
								product.Quantity, 
								product.CodeValue, 
								product.IsPublished, 
								product.Expiration, 
								product.Price)

	if err!= nil {
		driverErr, ok := err.(*mysql.MySQLError)
        if !ok {
            return 0, err
        }
		switch driverErr.Number {
        case 1062:
            return 0, errors.New(driverErr.Message)
        default:
            return 0, errors.New("error por otra cosa")
        }
	}

	rowsAffected , err := result.RowsAffected()

	if err != nil{
		return 0, err
    }

	if rowsAffected != 1 {
		return 0, err
	}
	
	id, err := result.LastInsertId()
	if err!= nil {
		return 0, err
	}
	product.Id = int(id)
	return product.Id, nil
}


func (r *repository) Exists(code_value string) bool{
	
	var exists bool
    var id int
    query := "SELECT id FROM products WHERE code_value = ?;"
    row := r.db.QueryRow(query, code_value)
    err := row.Scan(&id)
    if err != nil {
        return false
    }
    if id > 0 {
        exists = true
		return exists
    }
    return exists
	
}

func (r *repository) Read(id int) (domain.Product, error){
	
	var product domain.Product

	query := "SELECT * FROM products WHERE id = ?;"
	
	row := r.db.QueryRow(query, id)
	err := row.Scan(&product.Id, 
					&product.Name, 
					&product.Quantity,
					&product.CodeValue,
					&product.IsPublished,
					&product.Expiration,
					&product.Price)

	if err!= nil {
		return domain.Product{}, err
	}
	
	return product,nil
}