package handler

import (
	"errors"
	"net/http"
	"strconv"

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

func (sqlH *productSQLHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		product, err := sqlH.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("product not found"))
			return
		}
		web.Success(c, 200, product)
	}
}

func (sqlH *productSQLHandler) Getall() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := sqlH.s.GetAll()
		if err!= nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusOK, products)
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

func (sqlH *productSQLHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}

		_, err = sqlH.s.GetByID(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("product not found"))
			return
		}
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		var product domain.Product
		err = ctx.ShouldBindJSON(&product)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid json"))
			return
		}
		err = sqlH.s.Update(id, product)
		if err != nil {
			web.Failure(ctx, 409, errors.New("error en handler2"))
			return
		}
		product.Id = id
		web.Success(ctx, 200, product)
	}
}

func (sqlH *productSQLHandler) Delete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err!= nil {
			web.Failure(ctx, 400, errors.New("invalid id"))
			return
		}
		
		err = sqlH.s.Delete(id)
		if err!= nil {
			web.Failure(ctx,404,err)
			return
		}

		web.Success(ctx, 204, "id deleted")
	}
}

