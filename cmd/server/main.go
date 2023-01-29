package main

import (
	"database/sql"

	"github.com/Lelo88/EjercicioGoDb2/cmd/server/handler"
	//"github.com/Lelo88/EjercicioGoDb2/internal/product"
	"github.com/Lelo88/EjercicioGoDb2/internal/productSQL"
	//"github.com/Lelo88/EjercicioGoDb2/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {

	/*storage := store.NewJsonStore("./products.json")

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)*/

	databaseConfig := &mysql.Config{
		User: "root",
		Passwd: "",
		Addr: "localhost:3306",
		DBName: "my_db",
		ParseTime: true,
	} 

	db, err:= sql.Open("mysql", databaseConfig.FormatDSN())

	if err!= nil {
		panic(err.Error())

	}

	repo:= productSQL.NewSQLRepository(db)
	service := productSQL.NewSqlService(repo)
	productHandler := handler.NewProductSQLHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		//products.DELETE(":id", productHandler.Delete())
		//products.PATCH("/:id", productHandler.Patch())
		//products.PUT("/:id", productHandler.Put())
	}

	r.Run(":8080")
}
