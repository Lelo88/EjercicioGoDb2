package productSQL

import (
	"database/sql"
	"errors"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	"github.com/go-sql-driver/mysql"
	//"github.com/go-sql-driver/mysql"
)

type Repository interface {
	Create(product domain.Product) error
	Read(id int) (domain.Product, error)
	ReadAll() ([]domain.Product, error)
	Exists(codevalue string) bool
	Delete(id int) error
	Update(product domain.Product) error
}

type repository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(product domain.Product) (err error) {

	query := `INSERT INTO products(name,quantity,code_value,is_published,expiration,price) 
			VALUES (?,?,?,?,?,?);`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return ErrInternal
	}

	defer statement.Close()

	result, err := statement.Exec(
		&product.Name,
		&product.Quantity,
		&product.CodeValue,
		&product.IsPublished,
		&product.Expiration,
		&product.Price)

	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			err = ErrDatabaseNotFound
			return err
		}

		switch driverErr.Number {
		case 1062:
			err = ErrDuplicate
		default:
			err = ErrInternal
		}
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil || rowsAffected != 1 {
		return ErrInternal
	}

	id, err := result.LastInsertId()
	if err != nil {
		return ErrInternal
	}

	product.Id = int(id)

	return
}

func (r *repository) Exists(codevalue string) bool {

	query := "SELECT code_value FROM products WHERE code_value=?;"
	row := r.db.QueryRow(query, codevalue)
	err := row.Scan(&codevalue)
	return err == nil

}

func (r *repository) Read(id int) (domain.Product, error) {

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

	if err != nil {

		switch err {
		case sql.ErrNoRows:
			return domain.Product{}, ErrNotFound
		default:
			return domain.Product{}, ErrInternal
		}

	}

	return product, nil
}

func (r *repository) Update(product domain.Product) error {
	query := `UPDATE products SET name=?, quantity=?, code_value=?, is_published=?, expiration=?, price=? WHERE id=?`
	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	result, err := statement.Exec(&product.Name,
		&product.Quantity,
		&product.CodeValue,
		&product.IsPublished,
		&product.Expiration,
		&product.Price,
		&product.Id)

	if err != nil {
		return errors.New("error1")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("error2")
	}

	if rowsAffected != 1 {
		return errors.New("no se actualiza nada")
	}

	return nil
}

func (r *repository) ReadAll() ([]domain.Product, error) {
	query := `Select * from products`

	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.Product{}, ErrInternal
	}

	var producto []domain.Product

	for rows.Next() {
		prod := domain.Product{}
		err := rows.Scan(&prod.Id, &prod.Name,
			&prod.Quantity, &prod.CodeValue,
			&prod.IsPublished, &prod.Expiration,
			&prod.Price)
		if err != nil {
			return []domain.Product{}, ErrInternal
		}
		producto = append(producto, prod)
	}

	return producto, nil
}

func (r *repository) Delete(id int) error {
	query := "DELETE FROM products WHERE id =?"
	statement, err := r.db.Prepare(query)

	if err != nil {
		return ErrInternal
	}

	res, err := statement.Exec(id)
	if err != nil {
		return ErrInternal
	}

	affect, err := res.RowsAffected()

	if err != nil {
		return ErrDatabaseNotFound
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}
