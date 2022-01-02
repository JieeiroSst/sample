package server

import (
	"github.com/gin-gonic/gin"
	"test/internal/repository"
	"test/internal/router"
	"test/internal/usecase"
	"test/pkg/postgres"

)

type server struct {
	engine *gin.Engine
}

type Server interface {
	Run() error
}

func NewServer(engine *gin.Engine) Server {
	return &server{
		engine:engine,
	}
}

func (s *server) Run() error {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	postgresConn := postgres.NewPostgresConn(dsn)
	repository := repository.NewRepository(postgresConn)
	usecase := usecase.NewUsecase(repository)
	router := router.NewRouter(usecase)

	s.engine.POST("/", router.Create)
	s.engine.GET("/last_access", router.LastAccess)
	s.engine.GET("/lots_access", router.LotsAccess)

	return nil
}