package foodservices

import (
	"calorie-counter/models"
	"fmt"
	"time"
)

func (*foodServicesStruct) CreateFood(newFood models.Food) {
	timestamp := time.Now()
	pointer := &timestamp
	newFood.CreatedAt = pointer
	fmt.Println("Este servicio creará una comida")
}
