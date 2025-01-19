1. Gin
   
简介：

Gin 是一个高性能的 Go Web 框架，以其快速和简洁而著称。它受到了 Martini 框架的启发，但进行了大量优化以提高性能。

主要特性：

* 高性能： Gin 使用 httprouter 作为路由引擎，性能非常优秀，适合对性能要求高的应用。
* 中间件支持： 内置了丰富的中间件，并且支持自定义中间件的编写。
* JSON 验证： 提供强大的 JSON 验证功能，简化数据验证过程。
* 错误处理： 集成了高级的错误处理机制，便于捕获和处理运行时错误。
* 路由分组： 支持路由分组，便于组织和管理路由。

适用场景：

适合构建高性能的 API 服务，特别是在性能关键型应用中表现出色。

2. Echo

简介：
   
Echo 是另一个高性能的 Go Web 框架，强调简洁性和易用性，同时也提供了丰富的功能。

主要特性：

* 高性能： 与 Gin 类似，Echo 也采用了高效的路由引擎，性能优越。
* 极简设计： API 设计简洁，易于上手和使用。
* 中间件丰富： 内置了多种常用中间件，并支持自定义中间件。
* 数据绑定与验证： 提供强大的数据绑定和验证功能，支持多种格式（如 JSON、XML、表单）。
* WebSocket 支持： 内置对 WebSocket 的支持，适合实时应用。

适用场景：

适合需要快速开发且性能要求较高的 Web 应用和 API 服务，尤其适用于实时性要求较高的应用场景。


3. Beego

简介：

Beego 是一个全功能的 Go Web 框架，类似于 Django 或 Ruby on Rails，提供了丰富的功能集，适合构建大型应用。

主要特性：

* 全栈框架： 提供了路由、中间件、ORM（Bee ORM）、缓存、日志、配置管理等一站式解决方案。
* 自动生成文档： 支持自动生成 API 文档，便于维护和协作。
* 模块化设计： 支持模块化开发，便于代码组织和复用。
* 内置工具： 提供了命令行工具用于项目生成、代码自动化等，提升开发效率。
* MVC 架构： 支持 MVC（模型-视图-控制器）架构，适合大型复杂项目的开发。

适用场景：

适合构建大型、复杂的 Web 应用，需要一站式解决方案和丰富功能支持的项目。

选择建议
1. 性能优先且需要快速开发 API： 选择 Gin 或 Echo。
2. 需要丰富内置功能和全栈支持： 选择 Beego。
3. 简单项目或对性能要求极高： Gin 和 Echo 均是不错的选择，具体可以根据个人偏好选择。


```shell
curl http://localhost:8080

curl.exe http://localhost:8080/api/users
```

```go
// 创建一个带有默认中间件的 Gin 引擎
r := gin.Default()

// Default() 包含了以下两个默认中间件：
// - Logger() 用于日志记录
// - Recovery() 用于panic恢复


// 使用全局中间件
r.Use(Logger())

// 中间件示例
func Logger() gin.HandlerFunc {
return func(c *gin.Context) {
// 请求前的处理
t := time.Now()

// 处理请求
c.Next()

// 请求后的处理
latency := time.Since(t)
// 记录日志
log.Printf("请求耗时: %v", latency)
}
}

// 提供静态文件服务
r.Static("/static", "./static")

// 参数说明：
// - "/static": URL路径
// - "./static": 本地文件夹路径

// 使用示例：
// 访问 http://localhost:8080/static/image.jpg 
// 将返回 ./static/image.jpg 文件


// 加载HTML模板文件
r.LoadHTMLGlob("templates/*")

// 加载所有模板文件
// templates/
// ├── index.html
// ├── user.html
// └── about.html


// 创建路由组
api := r.Group("/api")
// 所有组内路由都会带有 "/api" 前缀

{
// GET 请求示例
api.GET("/users", getUsers)  // 完整路径: /api/users

// 带参数的路由
api.GET("/users/:id", getUserByID)  // 完整路径: /api/users/123

// 参数获取示例
func getUserByID(c *gin.Context) {
id := c.Param("id")  // 获取URL参数
// ...
}
}


// GET 请求 - 获取资源
api.GET("/users", getUsers)

// POST 请求 - 创建资源
api.POST("/users", createUser)

// PUT 请求 - 更新资源
api.PUT("/users/:id", updateUser)

// DELETE 请求 - 删除资源
api.DELETE("/users/:id", deleteUser)

// 处理函数示例
func createUser(c *gin.Context) {
var user User
// 绑定JSON请求体
if err := c.ShouldBindJSON(&user); err != nil {
c.JSON(400, gin.H{"error": err.Error()})
return
}
// ...
}

// 启动HTTP服务器
r.Run(":8080")

// 相当于 http.ListenAndServe(":8080", r)
// 默认监听 localhost:8080

// context常用方法


func handler(c *gin.Context) {
// 获取参数
id := c.Param("id")            // 路径参数
name := c.Query("name")        // 查询参数 (?name=value)
page := c.DefaultQuery("page", "1")  // 带默认值的查询参数

// 获取请求体
var json struct{}
c.ShouldBindJSON(&json)        // 绑定JSON

// 响应
c.JSON(200, gin.H{})           // 返回JSON
c.String(200, "Hello")         // 返回字符串
c.HTML(200, "template.html", data)  // 返回HTML

// 设置Header
c.Header("Content-Type", "application/json")

// 获取Header
userAgent := c.GetHeader("User-Agent")
}

```

完整的请求处理流程
```go
func ProcessRequest(c *gin.Context) {
    // 1. 参数验证
    var req RequestData
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // 2. 业务逻辑处理
    result, err := service.Process(req)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // 3. 返回响应
    c.JSON(200, result)
}

```

当然，下面我将详细解释你提供的 Gin 框架 `main` 函数中使用的各个函数及其作用。这将帮助你更好地理解每个部分的功能及其在整个应用中的作用。

### 完整的 `main` 函数代码

```go
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
```

### 逐行解析

#### 1. 创建默认的 Gin 引擎

```go
r := gin.Default()
```

- **函数**：`gin.Default()`
- **作用**：创建一个默认的 Gin 路由引擎实例。这包括两个默认的中间件：
    - **Logger 中间件**：记录每个请求的详细信息（如方法、路径、状态码、响应时间等）。
    - **Recovery 中间件**：在发生 panic 时恢复并返回 500 错误，防止整个服务器崩溃。

- **用途**：作为整个应用的核心路由引擎，处理所有的 HTTP 请求和路由匹配。

#### 2. 使用自定义中间件

```go
r.Use(Logger())
```

- **函数**：`r.Use(...)`
- **作用**：将一个或多个中间件函数添加到路由引擎中。中间件是在路由处理器之前或之后执行的函数，可以用于处理认证、日志记录、错误处理等。

- **自定义中间件**：`Logger()`
    - 这是一个用户自定义的中间件函数，用于特定的日志记录或其他自定义功能。
    - **示例实现**：
      ```go
      func Logger() gin.HandlerFunc {
          return func(c *gin.Context) {
              // 在请求之前执行
              startTime := time.Now()
  
              // 处理请求
              c.Next()
  
              // 在请求之后执行
              duration := time.Since(startTime)
              status := c.Writer.Status()
              log.Printf("Status: %d | Duration: %v | Path: %s", status, duration, c.Request.URL.Path)
          }
      }
      ```

- **用途**：在处理实际请求之前或之后执行一些通用逻辑，如记录日志、处理身份验证等。

#### 3. 静态文件服务

```go
r.Static("/static", "./static")
```

- **函数**：`r.Static(relativePath, root string)`
- **作用**：将指定的文件夹作为静态文件服务器提供。例如，将本地的 `./static` 文件夹映射到 URL 路径 `/static`。

- **参数**：
    - `relativePath`：客户端请求的路径前缀（如 `/static`）。
    - `root`：服务器上静态文件的存放位置（如 `./static`）。

- **用途**：提供静态资源（如 CSS、JavaScript、图片等）给客户端。例如，客户端可以通过 `http://localhost:8080/static/style.css` 访问 `./static/style.css` 文件。

- **示例目录结构**：
  ```
  project/
  ├── main.go
  ├── static/
  │   ├── style.css
  │   └── script.js
  └── templates/
      └── index.html
  ```

#### 4. 加载 HTML 模板

```go
r.LoadHTMLGlob("templates/*")
```

- **函数**：`r.LoadHTMLGlob(pattern string)`
- **作用**：加载符合指定模式的所有 HTML 模板文件，用于在处理请求时渲染 HTML 页面。

- **参数**：
    - `pattern`：匹配模板文件的路径模式，例如 `templates/*` 表示加载 `templates` 文件夹中的所有文件。

- **用途**：使得应用能够渲染动态 HTML 页面。例如，可以使用 `c.HTML` 方法在处理请求时渲染模板。

- **模板文件示例** (`templates/index.html`)：
  ```html
  <!DOCTYPE html>
  <html>
  <head>
      <title>{{ .title }}</title>
      <link rel="stylesheet" href="/static/style.css">
  </head>
  <body>
      <h1>{{ .title }}</h1>
      <script src="/static/script.js"></script>
  </body>
  </html>
  ```

#### 5. 基础路由

```go
r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "Gin Demo",
    })
})
```

- **函数**：`r.GET(path string, handlers ...gin.HandlerFunc)`
- **作用**：定义一个处理 GET 请求的路由。在此例中，根路径 `/` 被注册，用于渲染 `index.html` 模板。

- **参数**：
    - `path`：路由路径，这里是 `/`。
    - `handlers`：一组处理函数（可以是一个或多个）。这里使用了一个匿名函数作为处理器。

- **处理器函数**：
  ```go
  func(c *gin.Context) {
      c.HTML(http.StatusOK, "index.html", gin.H{
          "title": "Gin Demo",
      })
  }
  ```
    - **`c *gin.Context`**：Gin 提供的上下文对象，包含请求和响应的信息。
    - **`c.HTML`**：渲染 HTML 模板并返回给客户端。
        - `http.StatusOK`：HTTP 状态码 200。
        - `"index.html"`：要渲染的模板文件。
        - `gin.H{...}`：传递给模板的数据（键值对）。在模板中可以通过 `{{ .title }}` 使用 `title` 值。

- **用途**：当客户端访问 `http://localhost:8080/` 时，服务器将渲染 `index.html` 模板，并将 `"Gin Demo"` 作为标题传递给模板。

- **访问示例**：
    - 客户端访问 `http://localhost:8080/`。
    - 服务器响应渲染后的 HTML 页面，包含传递的数据。

#### 6. API 路由组

```go
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
```

- **函数**：`r.Group(relativePath string, handlers ...gin.HandlerFunc)`
- **作用**：创建一个路由组，通过共享一个共同的基础路径和中间件，来组织相关的路由。

- **参数**：
    - `relativePath`：路由组的基础路径，这里是 `/api`。
    - `handlers`：可选的一组中间件，这里未指定（可传入中间件函数来为整个组添加中间件）。

- **用途**：将相关的路由组织在一起，增强代码的可读性和可维护性。例如，所有 API 路由都在 `/api` 组下，路径如 `/api/users`。

- **路由定义**：
    - **`api.GET("/users", getUsers)`**：
        - **路径**：`/api/users`
        - **方法**：GET
        - **处理器**：`getUsers`（用于获取所有用户）

    - **`api.GET("/users/:id", getUserByID)`**：
        - **路径**：`/api/users/:id`
        - **方法**：GET
        - **处理器**：`getUserByID`（用于获取指定 ID 的用户）

    - **`api.POST("/users", createUser)`**：
        - **路径**：`/api/users`
        - **方法**：POST
        - **处理器**：`createUser`（用于创建新用户）

    - **`api.PUT("/users/:id", updateUser)`**：
        - **路径**：`/api/users/:id`
        - **方法**：PUT
        - **处理器**：`updateUser`（用于更新指定 ID 的用户）

    - **`api.DELETE("/users/:id", deleteUser)`**：
        - **路径**：`/api/users/:id`
        - **方法**：DELETE
        - **处理器**：`deleteUser`（用于删除指定 ID 的用户）

- **处理器函数示例**：
  由于没有提供具体的实现，下面是每个处理器函数的示例定义：

  ```go
  // 获取所有用户
  func getUsers(c *gin.Context) {
      // 示例数据
      users := []map[string]interface{}{
          {"id": 1, "name": "Alice"},
          {"id": 2, "name": "Bob"},
      }
      c.JSON(http.StatusOK, users)
  }

  // 根据 ID 获取单个用户
  func getUserByID(c *gin.Context) {
      id := c.Param("id")
      // 根据 ID 查询用户逻辑（这里使用示例数据）
      user := map[string]interface{}{
          "id":   id,
          "name": "Alice",
      }
      c.JSON(http.StatusOK, user)
  }

  // 创建新用户
  func createUser(c *gin.Context) {
      var newUser struct {
          Name string `json:"name" binding:"required"`
      }
      if err := c.ShouldBindJSON(&newUser); err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
      }
      // 保存新用户逻辑（这里仅返回示例数据）
      createdUser := map[string]interface{}{
          "id":   3,
          "name": newUser.Name,
      }
      c.JSON(http.StatusCreated, createdUser)
  }

  // 更新用户
  func updateUser(c *gin.Context) {
      id := c.Param("id")
      var updateData struct {
          Name string `json:"name"`
      }
      if err := c.ShouldBindJSON(&updateData); err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
      }
      // 更新用户逻辑（这里仅示例）
      updatedUser := map[string]interface{}{
          "id":   id,
          "name": updateData.Name,
      }
      c.JSON(http.StatusOK, updatedUser)
  }

  // 删除用户
  func deleteUser(c *gin.Context) {
      id := c.Param("id")
      // 删除用户逻辑（这里仅示例）
      c.JSON(http.StatusOK, gin.H{"status": "用户已删除", "id": id})
  }
  ```

- **总结**：通过路由组 `/api`，你可以集中管理与 API 相关的所有路由，路径前缀为 `/api`，并对相关的 RESTful API 进行统一管理。

#### 7. 启动服务器

```go
r.Run(":8080")
```

- **函数**：`r.Run(address ...string)`
- **作用**：启动 Gin HTTP 服务器，监听指定的地址和端口。

- **参数**：
    - `address`（可选）：服务器的监听地址，格式为 `"host:port"`。这里是 `":8080"`，表示监听所有可用的网络接口的 8080 端口。

- **默认地址**：如果未指定，默认监听 `:8080`。

- **用途**：使服务器开始接收并处理 HTTP 请求。

- **错误处理**：
    - `r.Run()` 返回一个错误，如果服务器启动失败，可以捕获并处理。
    - 示例：
      ```go
      if err := r.Run(":8080"); err != nil {
          log.Fatalf("Failed to run server: %v", err)
      }
      ```

- **启动服务器后**：
    - 服务器将在指定的端口上运行，等待并处理进入的 HTTP 请求。
    - 你可以通过浏览器或工具（如 curl、Postman）访问服务器的路由。

### 其他相关函数和概念

#### 1. 自定义中间件 `Logger()`

前面提到的 `Logger()` 中间件，可以帮助记录请求的详细信息。自定义中间件可以用于多种目的，如认证、日志记录、请求修改等。

- **示例实现**：

  ```go
  import (
      "log"
      "time"
      
      "github.com/gin-gonic/gin"
  )

  func Logger() gin.HandlerFunc {
      return func(c *gin.Context) {
          // 请求开始时间
          startTime := time.Now()

          // 处理请求
          c.Next()

          // 请求结束时间
          endTime := time.Now()
          latency := endTime.Sub(startTime)

          // 获取请求信息
          method := c.Request.Method
          path := c.Request.URL.Path
          statusCode := c.Writer.Status()

          // 日志格式
          log.Printf("[GIN] %v | %3d | %13v | %-7s  %s\n",
              endTime.Format("2006-01-02 15:04:05"),
              statusCode,
              latency,
              method,
              path,
          )
      }
  }
  ```

- **作用**：记录每个请求的时间、状态码、请求方法和路径等信息。

#### 2. 路由处理器函数

在你的代码中，路由处理器函数如 `getUsers`、`getUserByID` 等负责处理具体的业务逻辑。它们接收 `*gin.Context` 参数，通过上下文对象读取请求信息、处理数据并返回响应。

- **通用模板**：

  ```go
  func handlerFunction(c *gin.Context) {
      // 获取参数
      param := c.Param("paramName")
      
      // 读取请求体
      var jsonData struct {
          Field1 string `json:"field1"`
          Field2 int    `json:"field2"`
      }
      if err := c.ShouldBindJSON(&jsonData); err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
      }
      
      // 业务逻辑处理
      result := processBusinessLogic(param, jsonData)
      
      // 返回响应
      c.JSON(http.StatusOK, result)
  }
  ```

- **详细解释**：
    - **获取路径参数**：使用 `c.Param("id")` 获取 URL 中的动态参数。
    - **读取请求体**：使用 `c.ShouldBindJSON(&struct)` 解析 JSON 请求体。
    - **业务逻辑**：执行具体的业务逻辑，如数据库操作、数据处理等。
    - **返回响应**：使用 `c.JSON(statusCode, data)` 返回 JSON 格式的响应。

### 综合示例

为了更好地理解整个流程，下面是一个完整的示例，包括中间件、路由处理器和服务器启动。

```go
package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// 自定义中间件：记录请求日志
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 请求开始时间
        startTime := time.Now()

        // 处理请求
        c.Next()

        // 请求结束时间
        endTime := time.Now()
        latency := endTime.Sub(startTime)

        // 获取请求信息
        method := c.Request.Method
        path := c.Request.URL.Path
        statusCode := c.Writer.Status()

        // 日志格式
        log.Printf("[GIN] %v | %3d | %13v | %-7s  %s\n",
            endTime.Format("2006-01-02 15:04:05"),
            statusCode,
            latency,
            method,
            path,
        )
    }
}

// 处理器函数：获取所有用户
func getUsers(c *gin.Context) {
    users := []map[string]interface{}{
        {"id": 1, "name": "Alice"},
        {"id": 2, "name": "Bob"},
    }
    c.JSON(http.StatusOK, users)
}

// 处理器函数：根据 ID 获取用户
func getUserByID(c *gin.Context) {
    id := c.Param("id")
    user := map[string]interface{}{
        "id":   id,
        "name": "Alice",
    }
    c.JSON(http.StatusOK, user)
}

// 处理器函数：创建新用户
func createUser(c *gin.Context) {
    var newUser struct {
        Name string `json:"name" binding:"required"`
    }
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdUser := map[string]interface{}{
        "id":   3,
        "name": newUser.Name,
    }
    c.JSON(http.StatusCreated, createdUser)
}

// 处理器函数：更新用户
func updateUser(c *gin.Context) {
    id := c.Param("id")
    var updateData struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    updatedUser := map[string]interface{}{
        "id":   id,
        "name": updateData.Name,
    }
    c.JSON(http.StatusOK, updatedUser)
}

// 处理器函数：删除用户
func deleteUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"status": "用户已删除", "id": id})
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
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
```

### 运行和测试

1. **项目结构**：

   ```
   project/
   ├── main.go
   ├── static/
   │   ├── style.css
   │   └── script.js
   └── templates/
       └── index.html
   ```

2. **创建 `templates/index.html` 示例文件**：

   ```html
   <!DOCTYPE html>
   <html>
   <head>
       <title>{{ .title }}</title>
       <link rel="stylesheet" href="/static/style.css">
   </head>
   <body>
       <h1>{{ .title }}</h1>
       <script src="/static/script.js"></script>
   </body>
   </html>
   ```

3. **运行服务器**：

   在终端中导航到项目目录并执行：

   ```bash
   go run main.go
   ```

   你将看到类似如下的输出，表示服务器已启动：

   ```
   [GIN-debug] Listening and serving HTTP on :8080
   ```

4. **访问路由**：

    - **根路由 `/`**：
        - **URL**：`http://localhost:8080/`
        - **功能**：渲染 `index.html` 模板，并显示页面标题 "Gin Demo"。

    - **静态文件 `/static/style.css`**：
        - **URL**：`http://localhost:8080/static/style.css`
        - **功能**：提供 `./static/style.css` 文件内容。

    - **API 路由**：
        - **获取所有用户**：
            - **URL**：`GET http://localhost:8080/api/users`
            - **响应**：
              ```json
              [
                  {"id": 1, "name": "Alice"},
                  {"id": 2, "name": "Bob"}
              ]
              ```

        - **根据 ID 获取用户**：
            - **URL**：`GET http://localhost:8080/api/users/1`
            - **响应**：
              ```json
              {
                  "id": "1",
                  "name": "Alice"
              }
              ```

        - **创建新用户**：
            - **URL**：`POST http://localhost:8080/api/users`
            - **请求体**：
              ```json
              {
                  "name": "Charlie"
              }
              ```
            - **响应**：
              ```json
              {
                  "id": 3,
                  "name": "Charlie"
              }
              ```

        - **更新用户**：
            - **URL**：`PUT http://localhost:8080/api/users/1`
            - **请求体**：
              ```json
              {
                  "name": "Alice Updated"
              }
              ```
            - **响应**：
              ```json
              {
                  "id": "1",
                  "name": "Alice Updated"
              }
              ```

        - **删除用户**：
            - **URL**：`DELETE http://localhost:8080/api/users/1`
            - **响应**：
              ```json
              {
                  "status": "用户已删除",
                  "id": "1"
              }
              ```

### 小结

在你的 `main` 函数中，使用了 Gin 的多个核心功能和扩展功能，包括：

1. **创建路由引擎**：`gin.Default()` 初始化默认的路由引擎，包含常用的中间件。
2. **添加中间件**：`r.Use(Logger())` 添加自定义日志中间件，用于记录请求信息。
3. **提供静态文件服务**：`r.Static("/static", "./static")` 使得 `./static` 目录下的文件可以通过 `/static` 路径访问。
4. **加载 HTML 模板**：`r.LoadHTMLGlob("templates/*")` 加载指定目录下的所有模板文件，用于渲染动态 HTML 页面。
5. **定义基础路由**：根路由 `/` 渲染 `index.html` 模板，显示欢迎页面。
6. **创建路由组**：使用 `r.Group("/api")` 将相关的 API 路由组织在一起，路径前缀为 `/api`，包括各种 CRUD 操作。
7. **启动服务器**：`r.Run(":8080")` 启动 HTTP 服务器，监听 8080 端口，开始处理请求。

通过这种结构化的方法，你可以轻松地组织和管理你的 Web 应用，确保代码的可读性和可维护性。如果有更多具体的问题或需要进一步的解释，请随时提问！