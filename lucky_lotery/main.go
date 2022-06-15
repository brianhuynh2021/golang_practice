package main

import (
	"encoding/json"
	"fmt"
	"golang_practice/lucky_lotery/data"
	"golang_practice/lucky_lotery/utils"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type Customer struct {
	CustomerID  string
	FullName    string
	PhoneNumber string
}

type Number struct {
	ID           uint      `json:"id"`
	Numbers      string    `json:"numbers"`
	Head         float64   `json:"head"`
	Tail         float64   `json:"tail"`
	Chanel       string    `json:"chanel"`
	CustomerName string    `json:"customer_name"`
	PhoneNumber  string    `json:"phone_number"`
	Date         time.Time `json:"date"`
}

type Response struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

var numbers []Number

func GetNumbers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := data.SetupDatabase()
	fmt.Println("Getting numbers...")
	rows, err := db.Query("SELECT * FROM numbers")
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		number := Number{}
		err = rows.Scan(&number.ID, &number.Numbers, &number.Head,
			&number.Tail, &number.Chanel, &number.CustomerName, &number.PhoneNumber, &number.Date)
		utils.CheckErr(err)
		numbers = append(numbers, number)
	}

	json.NewEncoder(w).Encode(Response{Result: numbers})
}

func addNumber(number *Number) error {
	db := data.SetupDatabase()
	return db.
	 QueryRow("INSERT INTO numbers(numbers, head, tail, chanel,customer_name, phone_number) VALUES($1,$2, $3, $4, $5, $6) RETURNING id",
	 number.Numbers, number.Head, number.Tail, number.Chanel, number.CustomerName, number.PhoneNumber).
	 Scan(&number.ID)
}

func PostNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		number Number
	)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: err})
		return
	}
	json.Unmarshal(body, &number)

	if err := addNumber(&number); err != nil {
		json.NewEncoder(w).Encode(Response{Error: err})
		return
	}

	json.NewEncoder(w).Encode(Response{Result: number})

}

func updateNumber(number *Number)(err error){
	db := data.SetupDatabase()
	_, err =  db.Exec("UPDATE numbers SET numbers = $2 where id = $1", number.ID, number.Numbers)
	return
}

func PutNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		number Number
	)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: err})
		return
	}
	json.Unmarshal(body, &number)

	if err := updateNumber(&number); err != nil {
		json.NewEncoder(w).Encode(Response{Error: err})
		return
	}
	json.NewEncoder(w).Encode(Response{Result: number.ID})

}

func deleteNumber(number *Number)(err error){
	db := data.SetupDatabase()
	_, err = db.Exec("DELETE from numbers where id = $1 returning id", number.ID)
	return
}

func DeleteNumber(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var (
		number Number
	)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: err})
		return
	}
	json.Unmarshal(body, &number)

	if err := deleteNumber(&number); err != nil {
		json.NewEncoder(w).Encode(Response{Error: err})
		return
	}
	json.NewEncoder(w).Encode(Response{Result: number.ID})
}

func main() {
	fmt.Println("Good luck to you with a lucky number!")
	data.SetupDatabase()
	router := mux.NewRouter()
	router.HandleFunc("/numbers", GetNumbers).Methods("GET")
	router.HandleFunc("/number", PostNumber).Methods("POST")
	router.HandleFunc("/number", PutNumber).Methods("PUT")
	router.HandleFunc("/number", DeleteNumber).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
