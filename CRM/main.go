package main

import (
	dbConnection "rutikbhosale/db"
	"rutikbhosale/model"
	"rutikbhosale/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := dbConnection.ConnectDb()
	r := gin.Default()
	model.InitializeDB(db)
	routers.RouterGroup(r, db)
	defer db.Close()

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
