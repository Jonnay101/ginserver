package ginserver

import (
	"github.com/gin-gonic/gin"
)

// SetRoutes sets the routes for the service
func (s *Service) SetRoutes(r *gin.Engine, path string) {
	// serve the list of people
	r.GET(path, s.GetPeople)

	// create new person
	r.POST(path, s.CreatePerson)

	// serve the list of people
	r.GET(path+"/:id", s.GetPerson)

	// update person
	r.PUT(path+"/:id", s.UpdatePerson)

	// delete person
	r.DELETE(path+"/:id", s.DeletePerson)
}
