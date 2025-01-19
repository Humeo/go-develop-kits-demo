package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRegistry 结构体，用于渲染 HTML 模板
type TemplateRegistry struct {
	templates *template.Template
}

// Render 实现 Echo 的 Renderer 接口
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// User 结构体示例
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// 创建 Echo 实例
	e := echo.New()

	// 设置自定义的渲染器
	t := &TemplateRegistry{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	// 中间件

	// 日志中间件
	e.Use(middleware.Logger())

	// 恢复中间件，防止程序崩溃
	e.Use(middleware.Recover())

	// 压缩中间件，压缩响应内容
	e.Use(middleware.Gzip())

	// 静态文件服务
	e.Static("/static", "static")

	// 路由

	// 基础路由，渲染 HTML 页面
	e.GET("/", func(c echo.Context) error {
		data := map[string]interface{}{
			"title": "Echo Demo",
		}
		return c.Render(http.StatusOK, "index.html", data)
	})

	// API 路由组
	api := e.Group("/api")

	// 获取所有用户
	api.GET("/users", getUsers)

	// 根据 ID 获取单个用户
	api.GET("/users/:id", getUserByID)

	// 创建新用户
	api.POST("/users", createUser)

	// 更新用户
	api.PUT("/users/:id", updateUser)

	// 删除用户
	api.DELETE("/users/:id", deleteUser)

	// 启动服务器
	e.Logger.Fatal(e.Start(":8080"))
}

// 获取所有用户
func getUsers(c echo.Context) error {
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	return c.JSON(http.StatusOK, users)
}

// 根据 ID 获取单个用户
func getUserByID(c echo.Context) error {
	id := c.Param("id")
	// 这里可以添加实际的数据库查询逻辑
	user := User{
		ID:   1,
		Name: "Alice",
	}
	// 简单示例，假设只有 ID 为 1 的用户存在
	if id != "1" {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

// 创建新用户
func createUser(c echo.Context) error {
	var newUser User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// 这里可以添加实际的数据库创建逻辑
	newUser.ID = 3 // 示例 ID
	return c.JSON(http.StatusCreated, newUser)
}

// 更新用户
func updateUser(c echo.Context) error {
	id := c.Param("id")
	var updatedUser User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// 这里可以添加实际的数据库更新逻辑
	updatedUser.ID = 1 // 示例 ID
	if id != "1" {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, updatedUser)
}

// 删除用户
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	// 这里可以添加实际的数据库删除逻辑
	if id != "1" {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}
