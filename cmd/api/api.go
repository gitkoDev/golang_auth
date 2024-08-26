package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gitkoDev/medods_task/service/user"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	addr string
	db   *pgx.Conn
}

func NewAPIServer(addr string, db *pgx.Conn) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := gin.Default()

	storage := user.NewStorage(s.db)
	handler := user.NewHandler(storage)
	handler.RegisterRoutes(router)

	if err := router.Run(s.addr); err != nil {
		return err
	}

	logrus.Printf("server running on port %s", s.addr)
	return nil
}
