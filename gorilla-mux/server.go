package main
 
import (
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)
  
var people []Person

 
func main() {
    people = append(people, Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
    people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})
    
    router := mux.NewRouter()
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":4000", router))
}