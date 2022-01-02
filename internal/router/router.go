package router

import (
	"github.com/gin-gonic/gin"
	"test/internal/usecase"
	"test/model"
	"time"
)

type router struct {
	usecase usecase.Usecase
}

type Router interface {
	Create(c *gin.Context)
	LotsAccess(c *gin.Context)
	LastAccess(c *gin.Context)
}

func NewRouter(usecase usecase.Usecase) Router {
	return &router{usecase:usecase}
}


func (r *router) Create(c *gin.Context) {
	ip := model.Ip{
		Name:    c.ClientIP(),
		Address: c.Request.Host,
		Time:    time.Now(),
	}
	if err := r.usecase.Create(ip); err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"message":"can't create data ip",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 400,
		"message":"create data ip success",
		"data": ip,
	})
}

func (r *router) LotsAccess(c *gin.Context) {
	ips, err := r.usecase.LotsAccess()
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"message":"can't get data ip lots access 10",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"message":"get data ip success lots access 10",
		"data": ips,
	})
}

func (r *router) LastAccess(c *gin.Context) {
	ips, err := r.usecase.LastAccess()
	if err != nil {
		c.JSON(500, gin.H{
			"status": 500,
			"message":"can't get data ip last access 10",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"message":"get data ip success last access 10",
		"data": ips,
	})
}