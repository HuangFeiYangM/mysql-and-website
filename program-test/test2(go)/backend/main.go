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
