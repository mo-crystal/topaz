package database

import (
	"errors"

	"gorm.io/gorm"
)

type Server struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Data      string `json:"data"`
	Enabled   bool   `json:"enabled"`
	PublicKey string `json:"publicKey"`
}

func GetServerByName(name string) *Server {
	var server Server
	result := db.Where("name = ?", name).First(&server)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &server
}

func GetServer(id int) *Server {
	server := &Server{Id: id}
	result := db.First(server)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return server
}

func SetServer(server *Server) {
	db.Save(server)
}
