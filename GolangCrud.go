package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	dbDriver = "mysql"
	dbUser   = "jack"
	dbPass   = "Avin@123"
	dbName   = "gocrud_app"
)

type User struct {
	Name  string
	Email string
}

func main() {
	// Create a new route
	r := mux.NewRouter()

	// Define your HTTP routes using the router
	r.HandleFunc("/user", createUserHandler).Methods("POST")
	// r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	// r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
	// r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	// Start the http server on port 8090
	log.Println("server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Parse Json data from the request body
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")

}

func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (?,?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
	return nil
}
