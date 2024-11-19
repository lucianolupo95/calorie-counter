package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)



func main() {
    fmt.Println("Esta funcionando!")
    dbUser := os.Getenv("DB_USER")
    if dbUser == "" {
        fmt.Println("La variable de entorno DB_USER no está configurada.")
    } else {
        fmt.Println("DB_USER:", dbUser)

    }
    router := mux.NewRouter()
    router.HandleFunc("/env", getEnv).Methods("GET")
    router.HandleFunc("/health", health).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	//Manejador de permisos para las transacciones, sin cors se revotarian las solicitudes externas
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func getEnv(w http.ResponseWriter, r *http.Request) {
    fmt.Println("ENV: ")
    fmt.Println(os.Getenv("DB"))

}
func health(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["status"] = "pass"
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
