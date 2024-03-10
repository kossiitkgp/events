package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type FormData struct {
	Year       string `json:"year"`
	CourseName string `json:"course_name"`
	Link       string `json:"file_link"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/form", FormHandler).Methods("POST")
	r.HandleFunc("/get/{id}", GetHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	// TODO read data from the request
	var data = FormData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Invalid Form data", 400)
		return
	}

	// TODO save the data
	err = saveData(data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, fmt.Sprintf("error saving"), 500)
		return
	}

	// TODO respond to the request Excercise for you

}

func saveData(entry FormData) error {
	// open the file
	file, err := os.OpenFile("data.json", os.O_CREATE|os.O_RDWR, 644)
	if err != nil {
		return err
	}
	defer file.Close()

	// read the file data
	var data []FormData
	json.NewDecoder(file).Decode(&data)

	// append the new entry
	data = append(data, entry)

	// write back to the file
	file.Seek(0, 0) // go to the begining of the file
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func readAllData(data *[]FormData) error {
	// open the file
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// read the file data
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return err
	}

	return nil
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	// read data from file
	var data []FormData
	readAllData(&data)

	// get id from url
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if len(data) <= id {
		http.Error(w, "ID out of bounds", http.StatusBadRequest)
		return
	}

	// return json
	res, err := json.Marshal(data[id])
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
