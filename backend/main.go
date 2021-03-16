package main

import (
	// "github.com/gin-gonic/gin"

	//internal package
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/huantingwei/fyp/kubebench"
	"github.com/huantingwei/fyp/kubescore"
	"github.com/huantingwei/fyp/login"
	"github.com/huantingwei/fyp/overview"
	"github.com/huantingwei/fyp/util"
)

func main() {
	/*
		setup DB
	*/
	db, ctx := util.NewDatabase()
	defer db.Client.Disconnect(ctx)

	/*
		setup Gin route
	*/
	router := gin.Default()

	router.Use(cors.Default())

	v1 := router.Group("/api/v1")
	{
		overview.NewService(v1, db)
		kubebench.NewService(v1, db)
		kubescore.NewService(v1, db)
	}

	auth := router.Group("/")
	{
		login.NewService(auth, db)
	}

	router.Run()
}
