package handler

import (
	"errors"

	"github.com/Lelo88/EjercicioGoDb2/internal/domain"
	"github.com/Lelo88/EjercicioGoDb2/internal/productSQL"
	"github.com/Lelo88/EjercicioGoDb2/pkg/web"
	"github.com/gin-gonic/gin"
)

type productSQLHandler struct {
	s productSQL.Service
}

func NewProductSQLHandler(s productSQL.Service) *productSQLHandler {
	return &productSQLHandler{
		s,
	}
}

func (sqlH *productSQLHandler) Post() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var product domain.Product

		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}

		p, err := sqlH.s.Create(product)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, p)
	}
}