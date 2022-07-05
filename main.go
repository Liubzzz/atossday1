package main

import (
	"encoding/json"
	"ex4/entity"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

var person entity.Person

func main() {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {

		r.Get("/gperson", GetPerson)
		r.Post("/cperson", CreatePerson)
	})

	http.ListenAndServe(":8080", router)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &person)
	fmt.Print(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r.Method)
	person = entity.Person{
		Name:   "Liuba",
		ID:     "123",
		Gender: "Male",
	}
	fmt.Fprint(w, person)
}
