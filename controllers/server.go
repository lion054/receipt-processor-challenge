package controllers

import (
	"net/http"

	"receipt-processor-challenge/models"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	DataMap map[string]models.Receipt
}

func (s *Server) Initialize() error {
	s.Router = gin.Default()
	s.DataMap = make(map[string]models.Receipt)
	s.SetUpRoutes()
	return nil
}

func (s *Server) SetUpRoutes() {
	s.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ping"})
	})

	// receipts routes
	s.Router.POST("/receipts/process", s.StoreReceipt)
}
