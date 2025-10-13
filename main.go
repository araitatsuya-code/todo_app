package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title" validate:"required,min=1,max=50"`
	Completed bool   `json:"completed"`
}

var validate *validator.Validate

var todos []Todo
var nextID = 1

func initData() {
	validate = validator.New()

	todos = []Todo{
		{ID: 1, Title: "Go言語を学ぶ", Completed: false},
		{ID: 2, Title: "買い物に行く", Completed: true},
	}
	nextID = 3
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
			"error":   "無効な入力です",
			"details": err.Error(),
		})
		return
	}

	if err := validate.Struct(newTodo); err != nil {
		errors := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errors = append(errors, err.Field()+"は必須です")
			case "min":
				errors = append(errors, err.Field()+"は最低"+err.Param()+"文字必要です")
			case "max":
				errors = append(errors, err.Field()+"は最大"+err.Param()+"文字までです")
			}
		}

		c.JSON(400, gin.H{
			"error":   "バリデーションエラー",
			"details": errors,
		})
		return
	}

	newTodo.ID = nextID
	nextID++
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

	c.JSON(404, gin.H{
		"error": "Todoが見つかりません",
	})
}

func updateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
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

func deleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "無効なIDです",
		})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)

			c.JSON(200, gin.H{
				"message": "Todoを削除しました",
			})
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
	r.DELETE("/todos/:id", deleteTodo)

	r.Run(":8080")
}
