package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)
type user struct {
	User string `json:"id"`
	Email string `json:"title"`
	Completed bool `json:"completed"`
}

var Users = []user{
	{User: "Michael Owen", Email: "Clean Room", Completed: false},
	{User: "David Beckham", Email: "Dirty Room", Completed: true},
	{User: "David Seaman", Email: "Clean windows", Completed: true},
}
func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Users)
}
func addUsers(context *gin.Context) {
	var newTodo user
	if err := context.BindJSON(&newTodo); err != nil {
		return 
	}
	Users = append(Users, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
func getUser(context *gin.Context) {
	id := context.Param("id")
	todo, err := getUserByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message":"user not found"})
		return 
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleUserStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getUserByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
} 

func getUserByID(id string) (*user, error) {
	for i,t := range Users {
		if t.User == id {
			return &Users[i], nil
		}
	}
	return nil, errors.New("users not found")
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.PATCH("users/:id", toggleUserStatus)
	router.POST("/users", addUsers)
	router.Run("localhost:9090")
}

