package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"test/internal/server"
)

func main(){
	router := gin.Default()

	if err :=server.NewServer(router); err != nil {
		log.Println(err)
	}

	if err := router.Run(":1234"); err != nil {
		log.Println(err)
	}
}