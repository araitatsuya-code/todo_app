package main

import (
	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

var todos []Todo

func initData() {
	todos = []Todo{
		{ID: 1, Title: "Go言語を学ぶ", Completed: false},
		{ID: 2, Title: "買い物に行く", Completed: true},
	}
}

func getTodos(c *gin.Context) {
	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func main() {
	initData()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/todos", getTodos)

	r.Run(":8080")
}
