package main

import (
	"bytes"
	"encoding/json"
	"flea-market-app/dto"
	"flea-market-app/infra"
	"flea-market-app/models"
	"flea-market-app/services"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatal("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "test item 1", Price: 1000, Description: "test item 1 description", SoldOut: false, UserID: 1},
		{Name: "test item 2", Price: 2000, Description: "test item 2 description", SoldOut: true, UserID: 1},
		{Name: "test item 3", Price: 3000, Description: "test item 3 description", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "password1"},
		{Email: "test2@example.com", Password: "password2"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	router.ServeHTTP(w, req)

	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	router := setup()

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "test item4",
		Price:       4000,
		Description: "test item4 description",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer((reqBody)))
	req.Header.Set("Authorization", "Bearer "+*token)

	router.ServeHTTP(w, req)

	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}
