package routers

import (
	"database/sql"
	"rutikbhosale/handler"

	"github.com/gin-gonic/gin"
)

func RouterGroup(r *gin.Engine, db *sql.DB) {
	r.POST("/v1/crm/createuser", handler.CreateUser(db))
	r.GET("/v1/crm/getallusers", handler.GetAllUsers(db))
	r.DELETE("/v1/crm/deleteuser/:id", handler.DeleteUser(db))
}
