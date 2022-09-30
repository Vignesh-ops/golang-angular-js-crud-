package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/cors"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "vicky"
	dbname   = "postgres"
)

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port = %d user= %s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	fmt.Println("The database is connected and error is ", err)
	return db
}

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func getstudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Student
	db := connect()
	result, err := db.Query("SELECT * from student")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post Student
		err := result.Scan(&post.ID, &post.Name, &post.Phone)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	fmt.Print(posts)
	json.NewEncoder(w).Encode(posts)
}

func viewstudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := connect()
	result, err := db.Query("SELECT * FROM student WHERE id = $1", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post Student
	for result.Next() {
		err := result.Scan(&post.ID, &post.Name, &post.Phone)
		if err != nil {
			panic(err.Error())
		}
		print(err)
	}
	json.NewEncoder(w).Encode(post)
}

func updatestudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db := connect()
	stmt, err := db.Prepare("UPDATE student SET name = $2 , phone = $3 WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("data", body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	Name := keyVal["name"]
	Phone := keyVal["phone"]

	fmt.Printf("data  %s", body)
	_, err = stmt.Exec(params["id"], Name, Phone)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "%s ", params["id"])
	//w.WriteHeader(http.StatusOK)

	defer db.Close()
}

func deletestudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := connect()
	stmt, err := db.Prepare("DELETE FROM student WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "data deleted successfully %s ", params["id"])
	defer db.Close()
}

func addstudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")

	db := connect()
	sqlStatement := `INSERT INTO student (name,phone)VALUES ($1, $2)RETURNING id`
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	//ID := keyVal["id"]
	Name := keyVal["name"]
	Phone := keyVal["phone"]

	fmt.Print(Name, Phone)
	id := 0
	err = db.QueryRow(sqlStatement, Name, Phone).Scan(&id)

	if err != nil {
		panic(err.Error())
	}
	var response = Student{}
	response = Student{}
	//var a interface{}
	//fmt.Print("interface", a)
	//response, err := json.Marshal(a)

	fmt.Print(err)
	json.NewEncoder(w).Encode(response)

	//fmt.Fprintf(w, "%s", response)
	//w.Write(response)
	//fmt.Fprintf(w, "values are inserted")

	defer db.Close()
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/GetStudents/", getstudent).Methods("GET")
	router.HandleFunc("/savestudent/", addstudent).Methods("POST")
	router.HandleFunc("/student/{id}", updatestudent).Methods("PUT")
	router.HandleFunc("/deletestudent/{id}", deletestudent).Methods("DELETE")
	router.HandleFunc("/getstdbyid/{id}", viewstudent).Methods("GET")
	//cors := handlers.AllowedMethods([]string{"*", "PUT", "POST", "GET", "DELETE"})
	handler := cors.AllowAll().Handler(router)
	http.ListenAndServe(":8080", handler) // handlers.CORS(cors)(router)
}
