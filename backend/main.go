package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	index int = 0
	db []gin.H = []gin.H{}
)

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt string `json:"createdAt"`
}

func main() {
	g := gin.Default()

	g.GET("/", Index)
	g.POST("/todo", createTodo)

	g.GET("/todo", getAllTodo)

	g.Run(":8080")
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Hello World!",
	})
}

func createTodo(c *gin.Context) {
	newTodo := map[string]string{}
	c.ShouldBindJSON(newTodo)
	fmt.Println(newTodo)
	c.JSON(201, gin.H{
		"id": 1,
		"title": "田野と話す",
		"description": "電話で、Webアプリケーションがどのようなものかを説明する",
		"createdAt": "2023-06-29T20:19:53+09:00",
	})
}

func getAllTodo(c *gin.Context) {
	c.JSON(200, []gin.H{
		{
			"id": 1,
		"title": "田野と話す",
		"description": "電話で、Webアプリケーションがどのようなものかを説明する",
		"createdAt": "2023-06-29T20:19:53+09:00",
		},{
			"id": 2,
		"title": "田野と喋る",
		"description": "電話で、Webアプリケーションがどのようなものかを説明する",
		"createdAt": "2023-06-29T20:19:53+09:00",
		},
	},)
}
