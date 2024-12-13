package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
)
var db *sql.DB

func Connect() {
	var err error
	dbPassword := os.Getenv("DB_PASSWORD")
	serviceURI := fmt.Sprint("postgres://avnadmin:", dbPassword, "@calorie-counter-calorie-counter.b.aivencloud.com:13577/calorie_counter?sslmode=require")
	conn, _ := url.Parse(serviceURI)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	db, err = sql.Open("postgres", conn.String())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func PingDb() {
	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

}