package config

import (
	"log"

	"github.com/Tirivashe/book_management/routes"
	"github.com/gofiber/fiber/v2"
)

type APIServer struct {
	addr string
	mux *fiber.App
}

func NewAPIServer(addr string) *APIServer {
	app := fiber.New()
	routes.RegisterRoutes(app)
	return &APIServer{
		addr: addr,
		mux: app,
	}
}

func (s *APIServer) Start() error {
	err := s.mux.Listen(s.addr)
	if err != nil {
		log.Fatal("Could not start the server: ", err)
		return err
	}
	log.Println("Server started at: ", s.addr)
	return nil
}


