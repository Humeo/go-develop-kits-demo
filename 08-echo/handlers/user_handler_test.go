package handlers

import (
	"08-echo/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserHandler_CreateUser(t *testing.T) {
	// 设置
	e := echo.New()
	userService := services.NewUserService()
	handler := NewUserHandler(userService)

	// 测试数据
	userJSON := `{"username":"test","email":"test@example.com","age":25}`

	// 创建请求
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// 执行
	if assert.NoError(t, handler.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
