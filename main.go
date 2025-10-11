package main

import (
	"strconv"
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

func createTodo(c *gin.Context) {
	var newTodo Todo

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(400, gin.H{
			"error": "無効な入力です",
		})
	}

	newTodo.ID = len(todos) + 1

	todos = append(todos, newTodo)

	c.JSON(201, newTodo)
}

func getTodo(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "無効なIDです",
		})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(200, todo)
			return
		}
	}

	c.JSON(404, gin.H {
		"error": "Todoが見つかりません",
	})
}

func updateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H {
			"error": "無効なIDです",
		})
		return
	}

	var updateTodo Todo
	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		c.JSON(400, gin.H{
			"error": "無効な入力です",
		})
	}

	for i, todo := range todos {
		if todo.ID == id {
			updateTodo.ID = id
			todos[i] = updateTodo
			
			c.JSON(200, updateTodo)
			return
		}
	}

	c.JSON(404, gin.H{
		"error": "Todoが見つかりません",
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
	r.POST("/todos", createTodo)
	r.GET("/todos/:id", getTodo)
	r.PUT("/todos/:id", updateTodo)

	r.Run(":8080")
}
