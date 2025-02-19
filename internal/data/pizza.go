package data

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/internal/models"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("data/pizza.json")
	if err != nil {
		fmt.Println("File error:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Pizzas); err != nil {
		fmt.Println("Error while decoding JSON: ", err)
	}
}

func SavePizza() {
	file, err := os.Create("data/pizza.json")
	if err != nil {
		fmt.Println("File error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Pizzas); err != nil {
		fmt.Println("Error while encoding JSON:", err)
	}

}
