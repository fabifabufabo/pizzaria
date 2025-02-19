package main

import (
	"pizzaria/internal/data"

	"github.com/gin-gonic/gin"
	"pizzaria/internal/handler"
)

func main() {
	data.LoadPizzas()
	r := gin.Default()
	r.GET("/pizzas", handler.GetPizzas)
	r.POST("/pizzas", handler.PostPizzas)
	r.GET("/pizzas/:id", handler.GetPizzasByID)
	r.DELETE("/pizzas/:id", handler.DeletePizzaByID)
	r.PUT("/pizzas/:id", handler.UpdatePizzaByID)
	r.Run()
}
