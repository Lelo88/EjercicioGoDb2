package productSQL

import (
	"database/sql"
	"errors"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	 "github.com/go-sql-driver/mysql"
)

type Repository interface {
	Create(product domain.Product) (error)
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

func (r *repository) Create(product domain.Product) (error) {

	query:=`INSERT INTO products(id,name, quantity, code_value, is_published, expiration, price) 
			VALUES (?,?,?,?,?,?,?);`

	statement, err := r.db.Prepare(query)
	if err!= nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(product.Id,
								product.Name, 
								product.Quantity, 
								product.CodeValue, 
								product.IsPublished, 
								product.Expiration, 
								product.Price)

	if err!= nil {
		driverErr, ok := err.(*mysql.MySQLError)
        if !ok {
            return err
        }
		switch driverErr.Number {
        case 1062:
            return  errors.New(driverErr.Message)
		default:
			return  errors.New("error aca")
		}
		
	}

	rowsAffected , err := result.RowsAffected()

	if err != nil{
		return  errors.New("error1")
    }

	if rowsAffected != 1 {
		return  errors.New("error2")
	}
	
	_, err = result.LastInsertId()
	if err!= nil {
		return  errors.New("error3")
	}
	
	return  nil
}


func (r *repository) Exists(code_value string) bool{
	
	query := "SELECT card_number_id FROM buyers WHERE card_number_id=?;"
	row := r.db.QueryRow(query, code_value)
	err := row.Scan(&code_value)
	return err == nil
	
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