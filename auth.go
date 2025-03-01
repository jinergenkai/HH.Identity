package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// print username and password
	// fmt.Println("đăng nhập", username, password)

	var hash string
	err := db.QueryRow("SELECT password_hash FROM account WHERE username=$1", username).Scan(&hash)
	if err == sql.ErrNoRows {
		http.Error(w, "Sai tài khoản hoặc mật khẩu", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Lỗi server", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		http.Error(w, "Sai tài khoản hoặc mật khẩu", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 365).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Không tạo được token", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, tokenString)
}

func RegisterUser(username, password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := db.Exec("INSERT INTO account (username, password_hash, fullname, role, status) VALUES ($1, $2, $1, 'Staff', 'Active')", username, string(hash))
	if err != nil {
		log.Fatal("❌ Không thể tạo tài khoản:", err)
	}
	fmt.Println("✅ Tạo tài khoản thành công!")
}
