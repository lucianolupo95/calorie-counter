package main

import (
	"calorie-counter/routers"
	"log"

	"github.com/joho/godotenv"
)

// CalorieRecord representa un registro en la tabla calorie_counter
type CalorieRecord struct {
	ID        int     `json:"id"`
	Foods     string  `json:"foods"`
	Name      string  `json:"name"`
	Calories  int     `json:"calories"`
	Photo     string  `json:"photo"`
	Notes     string  `json:"notes"`
	CreatedAt string  `json:"created_at"`
	Visible   bool    `json:"visible"`
}



func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error cargando el archivo .env")
    }
}

func main() {
	loadEnv()
	// services.FoodServices.CreateFood()
	routers.Routes()
}

// // getCalories maneja el endpoint /calories
// func getCalories(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Consulta a la base de datos
// 	rows, err := db.Query("SELECT id, foods, name, calories, photo, notes, created_at, visible FROM public.calorie_counter")
// 	if err != nil {
// 		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
// 		log.Printf("Error al ejecutar la consulta: %v", err)
// 		return
// 	}
// 	defer rows.Close()

// 	// Mapea los resultados
// 	var records []CalorieRecord
// 	for rows.Next() {
// 		var record CalorieRecord
// 		if err := rows.Scan(&record.ID, &record.Foods, &record.Name, &record.Calories, &record.Photo, &record.Notes, &record.CreatedAt, &record.Visible); err != nil {
// 			http.Error(w, "Error al procesar los resultados", http.StatusInternalServerError)
// 			log.Printf("Error al escanear la fila: %v", err)
// 			return
// 		}
// 		records = append(records, record)
// 	}

// 	// Verifica si hubo errores en el proceso
// 	if err = rows.Err(); err != nil {
// 		http.Error(w, "Error al leer los resultados", http.StatusInternalServerError)
// 		log.Printf("Error al iterar las filas: %v", err)
// 		return
// 	}

// 	// Envía la respuesta en formato JSON
// 	if err := json.NewEncoder(w).Encode(records); err != nil {
// 		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
// 		log.Printf("Error al codificar la respuesta JSON: %v", err)
// 	}
// }
