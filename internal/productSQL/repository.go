package productSQL

import (
	"database/sql"
	"errors"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	//"github.com/go-sql-driver/mysql"
)

type Repository interface {
	Create(product *domain.Product) (error)
	Read(id int) (domain.Product, error)
	Exists(code_value string) bool
	Delete(id int) error

}


type repository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) Repository{
	return &repository{
		db:db,
	}
}

func (r *repository) Create(product *domain.Product) (error) {

	query:=`INSERT INTO products(name,quantity,code_value,is_published,expiration,price) 
			VALUES (?,?,?,?,?,?);`

	statement, err := r.db.Prepare(query)
	if err!= nil {
		return  err
	}


	result,err := statement.Exec(
								&product.Name, 
								&product.Quantity, 
								&product.CodeValue, 
								&product.IsPublished, 
								&product.Expiration, 
								&product.Price)

	if err!=nil{
		return err
	}

	rowsAffected , err := result.RowsAffected()

	if err != nil{
		return  errors.New("error1")
    }

	if rowsAffected < 1 {
		return  errors.New("error2")
	}
	
	id, err := result.LastInsertId()
	if err!= nil {
		return  errors.New("error3")
	}
	
	product.Id = int(id)

	return  nil
}


func (r *repository) Exists(code_value string) bool{
	
	query := "SELECT code_value FROM products WHERE code_value=?;"
	row := r.db.QueryRow(query, code_value)
	err := row.Scan(&code_value)
	return err==nil
	
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

func (r *repository) Update(product domain.Product) error{

	return nil
}

func (r *repository) Delete(id int) error{

	return nil
}