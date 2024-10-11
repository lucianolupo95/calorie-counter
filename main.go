package main

import (
	"fmt"
	"os"

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
}