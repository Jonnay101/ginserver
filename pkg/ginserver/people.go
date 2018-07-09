package ginserver

import (
	"math/rand"
	"strconv"
)

// Person struct
type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	ID        string `json:"id"`
}

// NewPerson creates a new person
func NewPerson(f string, l string) Person {
	return Person{
		Firstname: f,
		Lastname:  l,
		ID:        strconv.Itoa(rand.Intn(999999)),
	}
}
