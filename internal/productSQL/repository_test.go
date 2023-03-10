package productSQL

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRepository_ReadAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	dbExpected := []domain.Product{
		{
			Id:          1,
			Name:        "Milanesa",
			Quantity:    12,
			CodeValue:   "abc123",
			IsPublished: true,
			Expiration:  "2002-12-12",
			Price:       123.12,
		},

		{
			Id:          2,
			Name:        "Papas",
			Quantity:    13,
			CodeValue:   "abcaasd",
			IsPublished: true,
			Expiration:  "2002-1-12",
			Price:       12.12,
		},
	}
	rows := mock.NewRows([]string{"id", "name", "quantity", "code_value", "is_published", "expiration", "price"})

	for _, data := range dbExpected {
		rows.AddRow(data.Id, data.Name, data.Quantity, data.CodeValue, data.IsPublished, data.Expiration, data.Price)
	}
	rep := NewSQLRepository(db)
	//ctx := context.Background()

	query := "Select * from products"

	t.Run("Read All Products", func(t *testing.T) {
		//arrange
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
		//act
		products, err := rep.ReadAll()
		assert.NoError(t, err)
		assert.Equal(t, dbExpected, products)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Read All Products InternalError", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(ErrDatabaseNotFound)

		products, err := rep.ReadAll()

		assert.Error(t, err)
		assert.Empty(t, products, products)
		assert.Equal(t, ErrInternal, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Read All Products Scan Internal Error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		dbExpected := []domain.Product{
			{
				Id:          1,
				Name:        "Milanesa",
				Quantity:    12,
				CodeValue:   "abc123",
				IsPublished: true,
				Expiration:  "2002-12-12",
			},
			{
				Id:          2,
				Name:        "Papas",
				Quantity:    13,
				CodeValue:   "abcaasd",
				IsPublished: true,
				Expiration:  "2002-1-12",
			},
		}
		rows := mock.NewRows([]string{"id", "name", "quantity", "code_value", "is_published", "expiration"})

		for _, data := range dbExpected {
			rows.AddRow(data.Id, data.Name, data.Quantity, data.CodeValue, data.IsPublished, data.Expiration)
		}
		rep := NewSQLRepository(db)
		//ctx := context.Background()

		query := "Select * from products"

		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		products, err := rep.ReadAll()

		assert.Error(t, err)
		assert.Equal(t, ErrInternal, err)
		assert.Empty(t, products, products)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepository_Read(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	dbExpected := domain.Product{
		Id:          1,
		Name:        "Milanesa",
		Quantity:    12,
		CodeValue:   "abc123",
		IsPublished: true,
		Expiration:  "2002-12-12",
		Price:       123.12,
	}

	row := mock.NewRows([]string{"id", "name", "quantity", "code_value", "is_published", "expiration", "price"})
	row.AddRow(dbExpected.Id, dbExpected.Name, dbExpected.Quantity, dbExpected.CodeValue, dbExpected.IsPublished, dbExpected.Expiration, dbExpected.Price)

	rep := NewSQLRepository(db)

	query := "SELECT * FROM products WHERE id = ?;"

	t.Run("Get Product By ID OK", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(row)

		product, err := rep.Read(1)

		assert.NoError(t, err)
		assert.Equal(t, product, dbExpected)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetById ErrNotFound", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(sql.ErrNoRows)

		product, err := rep.Read(1)

		assert.Error(t, err)
		assert.Equal(t, ErrNotFound, err)
		assert.Empty(t, product, product)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByID ErrorInternal", func(t *testing.T) {
		// arrange
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(ErrInternal)

		// act
		buyer, err := rep.Read(1)

		// assert
		assert.Error(t, err)
		assert.Equal(t, ErrInternal, err)
		assert.Empty(t, buyer, buyer)
		//assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepository_Delete(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	rep := NewSQLRepository(db)

	query := "DELETE FROM products WHERE id =?"

	t.Run("Delete Product", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

		err := rep.Delete(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("ErrInternal Prepare on Delete Method", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(query)).WillReturnError(ErrInternal)

		err := rep.Delete(1)

		assert.Error(t, err)
		assert.Equal(t, ErrInternal, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Err Internal Exec on Delete Method", func(t *testing.T) {

		//arrange
		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WithArgs(1).WillReturnError(ErrInternal)

		//act
		err := rep.Delete(1)

		//assert
		assert.Error(t, err)
		assert.Equal(t, ErrInternal, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error Internal RowsAffected", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 0))

		err := rep.Delete(1)

		//aca me quede
		assert.Error(t, err)
		assert.Equal(t, err, ErrNotFound)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error Internal RowsAffected", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WillReturnResult(sqlmock.NewErrorResult(ErrDatabaseNotFound))

		err := rep.Delete(1)

		//aca me quede
		assert.Error(t, err)
		assert.Equal(t, ErrDatabaseNotFound, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	query := "INSERT INTO products(name,quantity,code_value,is_published,expiration,price) VALUES (?,?,?,?,?,?);"
	t.Run("create ok", func(t *testing.T) {

		product := domain.Product{Name: "Milanesa", Quantity: 10, CodeValue: "12345", IsPublished: true, Expiration: "2023/12/12", Price: 12.2}
		//expected :=1

		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WithArgs(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price).WillReturnResult(sqlmock.NewResult(1, 1))

		rep := NewSQLRepository(db)

		err := rep.Create(product)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.True(t, true)

	})

	t.Run("Error prepare", func(t *testing.T) {
		product := domain.Product{Name: "Milanesa", Quantity: 10, CodeValue: "12345", IsPublished: true, Expiration: "2023/12/12", Price: 12.2}

		mock.ExpectPrepare(regexp.QuoteMeta(query)).WillReturnError(ErrInternal)

		rep := NewSQLRepository(db)

		err := rep.Create(product)

		assert.Error(t, err)
		assert.Equal(t, ErrInternal, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Err Internal Exec", func(t *testing.T) {

		product := domain.Product{Name: "Milanesa", Quantity: 10, CodeValue: "12345", IsPublished: true, Expiration: "2023/12/12", Price: 12.2}

		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WillReturnError(ErrDatabaseNotFound)

		rep := NewSQLRepository(db)

		err := rep.Create(product)

		assert.Error(t, err)
		assert.Equal(t, ErrDatabaseNotFound, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("err case 1062", func(t *testing.T) {
		product := domain.Product{Name: "Milanesa", Quantity: 10, CodeValue: "12345", IsPublished: true, Expiration: "2023/12/12", Price: 12.2}

		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WillReturnError(&mysql.MySQLError{Number: 1062})

		rep := NewSQLRepository(db)

		err := rep.Create(product)

		assert.Error(t, err)
		assert.Equal(t, ErrDuplicate, err)
		//assert.Equal(t, 0)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("err default", func(t *testing.T) {
		product := domain.Product{Name: "Milanesa", Quantity: 10, CodeValue: "12345", IsPublished: true, Expiration: "2023/12/12", Price: 12.2}

		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WillReturnError(&mysql.MySQLError{Number: 0})

		rep := NewSQLRepository(db)

		err := rep.Create(product)

		assert.Error(t, err)
		assert.Equal(t, ErrInternal, err)
		//assert.Equal(t, 0)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("err row affected", func(t *testing.T) {

		product := domain.Product{Name: "Milanesa", Quantity: 10, CodeValue: "12345", IsPublished: true, Expiration: "2023/12/12", Price: 12.2}

		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WithArgs(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price).WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		rep := NewSQLRepository(db)

		err := rep.Create(product)

		assert.Error(t, err)
		assert.Equal(t, ErrInternal, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error last id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		product := domain.Product{Name: "Milanesa", Quantity: 10, CodeValue: "12345", IsPublished: true, Expiration: "2023/12/12", Price: 12.2}

		mock.ExpectPrepare(regexp.QuoteMeta(query)).ExpectExec().WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		rep := NewSQLRepository(db)

		err = rep.Create(product)

		assert.Error(t, err)
		assert.Equal(t, ErrInternal, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestRepository_Exists(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := mock.NewRows([]string{"code_value"})
	row.AddRow(1)

	query := "SELECT code_value FROM products WHERE code_value=?;"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("abc").WillReturnRows(row)

	rep := NewSQLRepository(db)

	result := rep.Exists("abc")

	assert.NoError(t, err)
	assert.Equal(t, true, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Update(t *testing.T) {

}
