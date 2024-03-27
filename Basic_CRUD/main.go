package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID         int
	first_name string
	last_name  string
	address    string
	class_id   int
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:Trung@614@tcp(127.0.0.1:3306)/devmaster")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	// Parse JSON data from the request body
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	CreateUser(db, user.first_name, user.last_name, user.address, user.class_id)
	
	w.WriteHeader(http.StatusCreated)
}
func CreateUser(db *sql.DB, first_name, last_name, address string, class_id int) {
	query := "INSERT INTO student (first_name,last_name,address,clazz_id) VALUES (?, ?, ?, ?)"
	db.Exec(query, first_name, last_name, address, class_id)
}
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]
	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)
	//fetch the user data from the database
	user, err := GetUser(db, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func GetUser(db *sql.DB, id int) (*User, error) {
	query := "SELECT * FROM student WHERE id = ?"
	row := db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.first_name, &user.last_name, &user.address, &user.class_id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func main() {
	// Create a new router
	r := mux.NewRouter()
	// Define HTTP routes using the router
	r.HandleFunc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
