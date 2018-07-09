package ginserver

import (
	"github.com/gin-gonic/gin"
)

// Service contains entire service
type Service struct {
	Path        string
	Port        string
	DBConnect   string
	DefaultPort string
}

// New initializes the Service
func New(p string, pt string, db string) Service {
	return Service{
		Path:        p,
		Port:        pt,
		DBConnect:   db,
		DefaultPort: ":3000",
	}
}

// StartServer initilaizes the gin server and sets then port
func (s *Service) StartServer() {
	// create gin router
	r := gin.Default()

	s.SetRoutes(r, s.Path)

	// use default port if one is not previded
	if s.Port == "" {
		r.Run(s.DefaultPort)
	} else {
		r.Run(s.Port)
	}
}
