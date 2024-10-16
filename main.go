package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))

}

func getEnv(w http.ResponseWriter, r *http.Request) {
    fmt.Println("ENV: ")
    fmt.Println(os.Getenv("DB"))

}
