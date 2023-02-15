package productSQL

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
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
}
