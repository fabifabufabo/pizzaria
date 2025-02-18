package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()
	r := gin.Default()
	r.GET("/pizzas", getPizzas)
	r.POST("/pizzas", postPizzas)
	r.GET("/pizzas/:id", getPizzasByID)
	r.DELETE("/pizzas/:id", deletePizzaByID)
	r.PUT("/pizzas/:id", updatePizzaByID)
	r.Run()
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": "Algo deu errado :(",
		})
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	c.JSON(201, newPizza)
}

func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pizza not found",
	})
}

func loadPizzas() {
	file, err := os.Open("data/pizza.json")
	if err != nil {
		fmt.Println("File error:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error while decoding JSON: ", err)
	}
}

func savePizza() {
	file, err := os.Create("data/pizza.json")
	if err != nil {
		fmt.Println("File error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error while encoding JSON:", err)
	}

}

func deletePizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, p := range pizzas {
		if p.ID == id {
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			savePizza()
			c.JSON(200, gin.H{
				"message": "Pizza deleted",
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Pizza not found",
	})

}

func updatePizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updatedPizza models.Pizza

	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, p := range pizzas {
		if p.ID == id {
			pizzas[i] = updatedPizza
			pizzas[i].ID = id
			savePizza()
			c.JSON(200, pizzas[i])
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza not found"})
}
