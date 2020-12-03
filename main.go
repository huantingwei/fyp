package main

import (
	"github.com/gin-gonic/gin"

	//internal package
	"github.com/huantingwei/fyp/util"
	"github.com/huantingwei/fyp/overview"
)

func main() {
	/*
		setup DB
	*/
	db, ctx := util.NewDatabase();
	defer db.Client.Disconnect(ctx);

	/*
		setup Gin route
	*/
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		overview.NewService(v1, db);
	}

	router.Run();
}