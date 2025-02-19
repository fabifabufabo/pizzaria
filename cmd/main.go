package main

import (
	"pizzaria/internal/data"

	"pizzaria/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()
	r := gin.Default()
	r.GET("/pizzas", handler.GetPizzas)
	r.POST("/pizzas", handler.PostPizzas)
	r.GET("/pizzas/:id", handler.GetPizzasByID)
	r.DELETE("/pizzas/:id", handler.DeletePizzaByID)
	r.PUT("/pizzas/:id", handler.UpdatePizzaByID)
	r.POST("/pizzas/:id/reviews", handler.PostReview)
	r.Run()
}
