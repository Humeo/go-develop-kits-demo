package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// User 用户结构体
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required,gte=0,lte=130"`
}

// 模拟数据库
var users = []User{
	{ID: 1, Username: "user1", Password: "pass1", Age: 20},
	{ID: 2, Username: "user2", Password: "pass2", Age: 25},
}

func main() {
	// 创建默认的 gin 引擎
	r := gin.Default()

	// 使用自定义中间件
	r.Use(Logger())

	// 静态文件服务
	r.Static("/static", "./static")

	// 加载HTML模板
	r.LoadHTMLGlob("templates/*")

	// 基础路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Gin Demo",
		})
	})

	// API 路由组
	api := r.Group("/api")
	{
		// GET 请求
		api.GET("/users", getUsers)
		api.GET("/users/:id", getUserByID)

		// POST 请求
		api.POST("/users", createUser)

		// PUT 请求
		api.PUT("/users/:id", updateUser)

		// DELETE 请求
		api.DELETE("/users/:id", deleteUser)
	}

	// 启动服务器
	r.Run(":8080")
}

// 中间件：日志记录
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 请求前
		c.Next()

		// 请求后
		latency := time.Since(t)
		status := c.Writer.Status()

		// 打印日志
		gin.DefaultWriter.Write([]byte(fmt.Sprintf(
			"[%s] %s %s %d %v\n",
			t.Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			status,
			latency,
		)))
	}
}

// 处理函数：获取所有用户
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// 处理函数：根据ID获取用户
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if fmt.Sprint(user.ID) == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

// 处理函数：创建用户
func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	c.JSON(http.StatusCreated, newUser)
}

// 处理函数：更新用户
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updateUser User

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if fmt.Sprint(user.ID) == id {
			updateUser.ID = user.ID
			users[i] = updateUser
			c.JSON(http.StatusOK, updateUser)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

// 处理函数：删除用户
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, user := range users {
		if fmt.Sprint(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
