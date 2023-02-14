package productSQL

import (
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
}
