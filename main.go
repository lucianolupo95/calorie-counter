package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func main() {
    err := godotenv.Load()
    fmt.Println("Esta funcionando!")
    if err != nil {
        fmt.Println("Error cargando el archivo .env "+ err.Error())
    }
    fmt.Println(os.Getenv("DB_USER"))

}
