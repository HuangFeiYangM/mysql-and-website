# 车辆管理系统全栈项目示例

下面我将为你展示一个完整的车辆管理系统项目架构，使用：
- **前端**: HTML/CSS + JavaScript (纯前端，不使用框架)
- **后端**: Go (Golang)
- **数据库**: MySQL

## 项目结构

```
vehicle-management-system/
├── frontend/          # 前端代码
│   ├── index.html    # 主页面
│   ├── style.css     # 样式表
│   ├── script.js     # 前端逻辑
│   └── auth.js       # 认证相关函数
├── backend/          # 后端代码
│   ├── main.go       # 主程序入口
│   ├── handlers/     # 请求处理器
│   ├── models/       # 数据模型
│   └── utils/        # 工具函数
├── database/         # 数据库脚本
│   └── schema.sql    # MySQL表结构
└── README.md         # 项目说明
```

## 1. 数据库设计 (MySQL)

```sql
-- database/schema.sql
CREATE DATABASE IF NOT EXISTS vehicle_management;
USE vehicle_management;

-- 用户表
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    role ENUM('admin', 'user', 'company') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 车辆信息表
CREATE TABLE vehicles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    make VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    year INT NOT NULL,
    license_plate VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

## 2. 前端实现 (HTML/CSS/JS)

```html
<!-- frontend/index.html -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Vehicle Management System</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div id="app">
        <header>
            <h1>Vehicle Management</h1>
            <div id="auth-buttons">
                <button id="login-btn">Login</button>
                <button id="register-btn">Register</button>
                <button id="logout-btn" style="display:none">Logout</button>
            </div>
        </header>

        <div id="content">
            <!-- 动态内容将在这里加载 -->
        </div>
    </div>

    <!-- 模态框 -->
    <div id="modal" class="modal">
        <div class="modal-content">
            <span class="close">&times;</span>
            <div id="modal-body"></div>
        </div>
    </div>

    <script src="auth.js"></script>
    <script src="script.js"></script>
</body>
</html>
```

```css
/* frontend/style.css */
body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
    background-color: #f4f4f9;
}

header {
    background-color: #333;
    color: white;
    padding: 1rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

#auth-buttons button {
    padding: 0.5rem 1rem;
    margin-left: 0.5rem;
    cursor: pointer;
}

.modal {
    display: none;
    position: fixed;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0,0,0,0.4);
}

.modal-content {
    background-color: #fefefe;
    margin: 15% auto;
    padding: 20px;
    border: 1px solid #888;
    width: 50%;
}

.close {
    color: #aaa;
    float: right;
    font-size: 28px;
    font-weight: bold;
    cursor: pointer;
}

.vehicle-card {
    background: white;
    margin: 1rem;
    padding: 1rem;
    border-radius: 5px;
    box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}
```

```javascript
// frontend/script.js
document.addEventListener('DOMContentLoaded', () => {
    // 检查用户是否已登录
    checkAuthStatus();
  
    // 事件监听
    document.getElementById('login-btn').addEventListener('click', showLoginForm);
    document.getElementById('register-btn').addEventListener('click', showRegisterForm);
    document.getElementById('logout-btn').addEventListener('click', logout);
  
    // 关闭模态框
    document.querySelector('.close').addEventListener('click', () => {
        document.getElementById('modal').style.display = 'none';
    });
});

function checkAuthStatus() {
    const token = localStorage.getItem('token');
    if (token) {
        document.getElementById('login-btn').style.display = 'none';
        document.getElementById('register-btn').style.display = 'none';
        document.getElementById('logout-btn').style.display = 'block';
        loadUserVehicles();
    }
}

function showLoginForm() {
    document.getElementById('modal-body').innerHTML = `
        <h2>Login</h2>
        <form id="login-form">
            <div>
                <label>Username:</label>
                <input type="text" name="username" required>
            </div>
            <div>
                <label>Password:</label>
                <input type="password" name="password" required>
            </div>
            <button type="submit">Login</button>
        </form>
    `;
  
    document.getElementById('modal').style.display = 'block';
  
    document.getElementById('login-form').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const data = Object.fromEntries(formData.entries());
        login(data);
    });
}

// 类似实现 register 和 vehicle 表单
```

## 3. 后端实现 (Go)

```go
// backend/main.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var jwtKey = []byte("your_secret_key")

func main() {
	// 初始化数据库连接
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/vehicle_management")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建路由
	router := mux.NewRouter()
	
	// 公开路由
	router.HandleFunc("/api/register", registerHandler).Methods("POST")
	router.HandleFunc("/api/login", loginHandler).Methods("POST")
	
	// 受保护路由
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(authMiddleware)
	authRouter.HandleFunc("/vehicles", getVehiclesHandler).Methods("GET")
	authRouter.HandleFunc("/vehicles", addVehicleHandler).Methods("POST")
	
	// 静态文件服务
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend")))
	
	// 启动服务器
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// 用户模型
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// 车辆模型
type Vehicle struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	LicensePlate string `json:"license_plate"`
}

// JWT声明
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// 注册处理器
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// 这里应该添加密码哈希处理
	_, err = db.Exec("INSERT INTO users (username, password, email, role) VALUES (?, ?, ?, ?)", 
		user.Username, user.Password, user.Email, user.Role)
	
	if err != nil {
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// 登录处理器
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// 验证用户
	var user User
	err = db.QueryRow("SELECT id, username, role FROM users WHERE username = ? AND password = ?", 
		credentials.Username, credentials.Password).Scan(&user.ID, &user.Username, &user.Role)
	
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	
	// 创建JWT令牌
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	
	// 返回令牌
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
		"role":  user.Role,
	})
}

// 认证中间件
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		
		tokenStr := cookie.Value
		claims := &Claims{}
		
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		// 将用户信息添加到请求上下文
		r = r.WithContext(context.WithValue(r.Context(), "user", claims))
		next.ServeHTTP(w, r)
	})
}

// 获取车辆处理器
func getVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("user").(*Claims)
	
	rows, err := db.Query("SELECT id, make, model, year, license_plate FROM vehicles WHERE user_id = ?", claims.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	var vehicles []Vehicle
	for rows.Next() {
		var v Vehicle
		if err := rows.Scan(&v.ID, &v.Make, &v.Model, &v.Year, &v.LicensePlate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vehicles = append(vehicles, v)
	}
	
	json.NewEncoder(w).Encode(vehicles)
}
```

## 4. 如何运行项目

1. **设置数据库**:
   - 创建MySQL数据库并执行`schema.sql`
   - 更新Go代码中的数据库连接字符串

2. **安装依赖**:
   ```bash
   go get github.com/gorilla/mux
   go get github.com/dgrijalva/jwt-go
   go get github.com/go-sql-driver/mysql
   ```

3. **启动后端**:
   ```bash
   cd backend
   go run main.go
   ```

4. **访问前端**:
   - 打开浏览器访问 `http://localhost:8080`

## 5. 功能扩展建议

1. **前端改进**:
   - 添加表单验证
   - 实现车辆CRUD操作的完整UI
   - 根据用户角色显示不同内容

2. **后端改进**:
   - 添加密码哈希（使用bcrypt）
   - 实现更精细的权限控制
   - 添加输入验证和清理

3. **数据库改进**:
   - 添加索引优化查询
   - 实现软删除

这个示例提供了完整的全栈实现基础，你可以根据需求进一步扩展功能。



