package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	db "go-gin/gen"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type updateRequest struct {
	Name        string `form:"name" json:"name" binding:"omitzero,max=40"`
	Email       string `form:"email" json:"email" binding:"omitzero,email"`
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"omitzero,phone11"`
}

type createRequest struct {
	Name        string `form:"name" json:"name" binding:"required,max=40"`
	Email       string `form:"email" json:"email" binding:"required,email"`
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"omitzero,phone11"`
	Password    string `form:"password" json:"password" binding:"required,min=8"`
}

type response struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

var q *db.Queries

func main() {
	conn, err := sql.Open("postgres", generateDsn())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	q = db.New(conn)

	setupValidator()

	r := gin.Default()
	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
	r.Run()
}

func generateDsn() string {
	// 環境変数から取得
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// DSN組み立て
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, port, dbname, sslmode,
	)

	return dsn
}

func setupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// カスタムバリデーション登録
		// 11桁の数字のみ許可するバリデーション（例: 08000001111）
		v.RegisterValidation("phone11", func(fl validator.FieldLevel) bool {
			phone := fl.Field().String()
			re := regexp.MustCompile(`^\d{11}$`)
			return re.MatchString(phone)
		})
	}
}

func createUser(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードをハッシュ化
	hashedP, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	params := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedP),
	}
	if req.PhoneNumber != "" {
		params.PhoneNumber = sql.NullString{String: req.PhoneNumber, Valid: true}
	}

	user, err := q.CreateUser(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := response{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber.String,
		CreatedAt:   user.CreatedAt.Time.String(),
		UpdatedAt:   user.UpdatedAt.Time.String(),
	}
	c.JSON(http.StatusCreated, res)
}

func getUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := q.GetUser(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	res := response{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber.String,
		CreatedAt:   user.CreatedAt.Time.String(),
		UpdatedAt:   user.UpdatedAt.Time.String(),
	}
	c.JSON(http.StatusOK, res)
}

func updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := db.UpdateUserParams{
		ID: int32(id),
	}
	if req.Name != "" {
		params.Name = sql.NullString{String: req.Name, Valid: true}
	}
	if req.Email != "" {
		params.Email = sql.NullString{String: req.Email, Valid: true}
	}
	if req.PhoneNumber != "" {
		params.PhoneNumber = sql.NullString{String: req.PhoneNumber, Valid: true}
	}

	user, err := q.UpdateUser(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := response{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber.String,
		CreatedAt:   user.CreatedAt.Time.String(),
		UpdatedAt:   user.UpdatedAt.Time.String(),
	}
	c.JSON(http.StatusOK, res)
}

func deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = q.DeleteUser(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
