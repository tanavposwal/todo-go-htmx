package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	defer DB.Close()

	e := gin.Default()
    e.Static("/static", "./static")
	e.LoadHTMLGlob("templates/*")

	e.GET("/", func(c *gin.Context) {
		todos := ReadTodoList()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	e.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		status := c.PostForm("status")
		id, _ := CreateTodo(title, status)

		c.HTML(http.StatusOK, "todo.html", gin.H{
			"title":  title,
			"status": status,
			"id":     id,
		})
	})

	e.DELETE("/todos/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)
		DeleteTodo(id)
		c.Status(http.StatusOK)
	})

	e.Run(":8080")
}
