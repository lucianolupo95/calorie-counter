package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"

// 	"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"
// 	_ "github.com/lib/pq"
// 	"github.com/rs/cors"
// )

// var db *sql.DB

// func mainOriginal() {
//     var err error
//     err = godotenv.Load()
//     fmt.Println("Esta funcionando!")
//     if err != nil {
//         fmt.Println("Error cargando el archivo .env "+ err.Error())
//     }
//     dbType := os.Getenv("DB")
//     dbName := os.Getenv("DB_NAME")
//     username := os.Getenv("DB_USER")
//     password := os.Getenv("DB_PASSWORD")
//     sslMode := os.Getenv("SSL_MODE")
//     db, err = sql.Open(dbType, fmt.Sprint("user=",username, " dbname=",dbName, " password=",password, " sslmode=", sslMode))
//     if err != nil {
//         log.Fatal(err)
//     }

//     defer db.Close()
//     err = db.Ping()
//     if err != nil {
//         log.Fatal("Error al conectar a la base de datos:", err)
//     }
//     router := mux.NewRouter()
//     router.HandleFunc("/users", getUsers).Methods("GET")
//     router.HandleFunc("/user", createUser).Methods("POST")
//     router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
//     router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

//     log.Println("Server running on port 8080")
//     // log.Fatal(http.ListenAndServe(":8080", router))
//     handler := cors.AllowAll().Handler(router)
//     //Si es en modo debug, le pongo la IP de localhost para que no me salte el firewall:
//     var debugModeString string
//     debugMode := os.Getenv("DEBUG_MODE")
//     debugModeBool, err := strconv.ParseBool(debugMode)
// 		if err != nil || debugModeBool {
//             debugModeString = "127.0.0.1"
//         }

//     log.Fatal(http.ListenAndServe(fmt.Sprint(debugModeString, ":8080"), handler))
// }
// type User struct {
//     ID    int    `json:"id"`
//     Username  string `json:"username"`
// }

// func getUsers(w http.ResponseWriter, r *http.Request) {
//     var users []User
//     rows, err := db.Query("SELECT id, username FROM users")
//     if err != nil {
//         http.Error(w, fmt.Sprint("Error obteniendo los elementos:", err.Error()), http.StatusInternalServerError)
//         return
//     }
//     defer rows.Close()

//     for rows.Next() {
//         var User User
//         err := rows.Scan(&User.ID, &User.Username)
//         if err != nil {
//             http.Error(w, "Error escaneando los elementos", http.StatusInternalServerError)
//             return
//         }
//         users = append(users, User)
//     }

//     if err := rows.Err(); err != nil {
//         http.Error(w, "Error finalizando el recorrido de los elementos", http.StatusInternalServerError)
//         return
//     }

//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(users)
// }

// func createUser(w http.ResponseWriter, r *http.Request) {
//     var user User
//     err := json.NewDecoder(r.Body).Decode(&user)
//     if err != nil {
//         http.Error(w, fmt.Sprint("Datos inválidos:", err.Error()), http.StatusBadRequest)
//         return
//     }

//     err = db.QueryRow("INSERT INTO users (username) VALUES ($1) RETURNING id", user.Username).Scan(&user.ID)
//     if err != nil {
//         http.Error(w, "Error creando el elemento", http.StatusInternalServerError)
//         return
//     }

//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(user)
// }

// func updateUser(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     id := vars["id"]

//     var User User
//     err := json.NewDecoder(r.Body).Decode(&User)
//     if err != nil {
//         http.Error(w, "Datos inválidos", http.StatusBadRequest)
//         return
//     }

//     result, err := db.Exec("UPDATE Users SET username = $1 WHERE id = $2", User.Username, id)
//     if err != nil {
//         http.Error(w, fmt.Sprint("Error actualizando el usuario:", err.Error()), http.StatusInternalServerError)
//         return
//     }

//     rowsAffected, err := result.RowsAffected()
//     if err != nil {
//         http.Error(w, "Error obteniendo las filas afectadas", http.StatusInternalServerError)
//         return
//     }

//     if rowsAffected == 0 {
//         http.Error(w, "Elemento no encontrado", http.StatusNotFound)
//         return
//     }

//     w.WriteHeader(http.StatusOK)
//     json.NewEncoder(w).Encode(User)
// }

// func deleteUser(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     id := vars["id"]

//     result, err := db.Exec("DELETE FROM Users WHERE id = $1", id)
//     if err != nil {
//         http.Error(w, "Error eliminando el elemento", http.StatusInternalServerError)
//         return
//     }

//     rowsAffected, err := result.RowsAffected()
//     if err != nil {
//         http.Error(w, "Error obteniendo las filas afectadas", http.StatusInternalServerError)
//         return
//     }

//     if rowsAffected == 0 {
//         http.Error(w, "Elemento no encontrado", http.StatusNotFound)
//         return
//     }

//     w.WriteHeader(http.StatusOK)
//     fmt.Fprint(w, "Elemento eliminado correctamente")
// }
