package server

import (
	"excercise2/internal/handler"
	"excercise2/internal/provider/db"
	"excercise2/internal/provider/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uc     manager.UsecaseManager
	engine *gin.Engine
	host   string
}

func (s *Server) setupControllers() {
	group := s.engine.Group("/api/v1")
	handler.NewHandlerEvent(s.uc.EventUc(), group).Route()
	handler.NewHandlerTransaction(s.uc.TransactionUc(), group).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {

	db, err := db.NewConnectionDatabase()
	if err != nil {
		fmt.Println(err)
	}

	repo := manager.NewRepoManager(db)
	uc := manager.NewUcManager(repo, db)
	engine := gin.Default()
	return &Server{uc: uc, engine: engine, host: ":8080"}
}
