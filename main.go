package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)



func main() {
    fmt.Println("La variable de entorno es: ")
    fmt.Println(os.Getenv("DB_USER"))

}
