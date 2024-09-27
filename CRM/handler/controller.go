package handler

import (
	"database/sql"
	"net/http"
	"rutikbhosale/model"

	"github.com/gin-gonic/gin"
)

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request model.User
		if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": "Invalid request payload"})
			return
		}

		user := model.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
		}
		_, err := db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})

	}
}

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		usr, err := db.Query("SELECT * FROM users WHERE id = $1", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			return

		}
		result, err := db.Exec("DELETE FROM users WHERE id = $1", id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		if rowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		} else {
			var users []model.User
			for usr.Next() {
				var user model.User
				err := usr.Scan(&user.ID, &user.Name, &user.Email, &user.Password) // Exclude password
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user"})
					return
				}
				users = append(users, user)
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "users": users})
		}

	}
}

// GetAllUsers handles fetching all users
func GetAllUsers(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query, err := db.Query("select * from users") // Exclude password
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}
		defer query.Close()

		var users []model.User
		for query.Next() {
			var user model.User
			err := query.Scan(&user.ID, &user.Name, &user.Email, &user.Password) // Exclude password
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user"})
				return
			}
			users = append(users, user)
		}

		// Check for errors after the loop
		if err = query.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate users"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"users": users})
	}
}
