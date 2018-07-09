package ginserver

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetPeople returns a slice of Person objects
func (s *Service) GetPeople(c *gin.Context) {
	file, err := ioutil.ReadFile("./mockDB.json")
	fmt.Println(string(file))
	check(err)
	c.String(200, "this will get all the people in the database")
}

// CreatePerson creates a new person and adds them to the db
func (s *Service) CreatePerson(c *gin.Context) {
	c.String(200, "this will create a new person and add them to the db")
}

// GetPerson returns a person with a matching id
func (s *Service) GetPerson(c *gin.Context) {
	c.String(200, "this will get a person from the database")
}

// UpdatePerson seraches the db and finds the person with the matching id
// the data for this person is then updated then saved back to the db
func (s *Service) UpdatePerson(c *gin.Context) {
	c.String(200, "this will update a given person in the db")
}

// DeletePerson finds the person with the matching id and removes them from the db
func (s *Service) DeletePerson(c *gin.Context) {
	c.String(200, "this will remove a person from the db")
}
