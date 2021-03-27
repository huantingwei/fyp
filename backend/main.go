package main

import (
	// "github.com/gin-gonic/gin"

	//internal package
	"github.com/gin-gonic/gin"
	"github.com/huantingwei/fyp/kubebench"
	"github.com/huantingwei/fyp/kubescore"
	"github.com/huantingwei/fyp/login"
	"github.com/huantingwei/fyp/overview"
	"github.com/huantingwei/fyp/util"
	"github.com/huantingwei/fyp/network"
)

func main() {
	/*
		setup DB
	*/
	db, ctx := util.NewDatabase()
	defer db.Client.Disconnect(ctx)

	k8sClient := util.GetKubeClientSet()

	/*
		setup Gin route
	*/
	router := gin.Default()

	router.Use(CORSMiddleware())

	v1 := router.Group("/api/v1")
	{
		overview.NewService(v1, db, k8sClient)
		kubebench.NewService(v1, db)
		kubescore.NewService(v1, db)
		network.NewService(v1, db, k8sClient)
	}

	auth := router.Group("/")
	{
		login.NewService(auth, db)
	}

	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
