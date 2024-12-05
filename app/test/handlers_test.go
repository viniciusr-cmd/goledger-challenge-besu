package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"goledger-challenge/vinicius/besu/handlers"
	"goledger-challenge/vinicius/besu/models"
	"goledger-challenge/vinicius/besu/routes"
)

var db *gorm.DB

func setupRouter() *gin.Engine {
	r := gin.Default()
	handlers.SetDB(db)
	routes.InitRoutes(r, db)
	return r
}

func setupMockDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to open gorm database connection")
	}

	db.AutoMigrate(&models.Contract{})

	db.Create(&models.Contract{
		Value:     123,
		UpdatedAt: time.Now(),
	})
}

func TestMain(m *testing.M) {
	setupMockDatabase()

	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	m.Run()
}

func TestSetContractHandler(t *testing.T) {
	MockSetContract = func(value uint) (*types.Receipt, error) {
		return &types.Receipt{
			Status: 1,
			TxHash: common.HexToHash("0x123"),
		}, nil
	}

	router := setupRouter()

	requestBody, _ := json.Marshal(map[string]interface{}{
		"value": 123,
	})

	req, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Success")
	assert.Contains(t, w.Body.String(), "123")
}

func TestGetContractHandler(t *testing.T) {
	MockGetContract = func() (*types.Receipt, error) {
		return &types.Receipt{
			Status: 1,
			TxHash: common.HexToHash("0x123"),
		}, nil
	}

	router := setupRouter()

	req, _ := http.NewRequest("GET", "/get", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Success")
	assert.Contains(t, w.Body.String(), "123")
}

func TestSyncContractHandler(t *testing.T) {
	MockSyncContract = func() (*types.Receipt, error) {
		return &types.Receipt{
			Status: 1,
			TxHash: common.HexToHash("0x123"),
		}, nil
	}

	router := setupRouter()

	req, _ := http.NewRequest("GET", "/sync", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckContractHandler(t *testing.T) {
	MockCheckContract = func() (*types.Receipt, error) {
		return &types.Receipt{
			Status: 1,
			TxHash: common.HexToHash("0x123"),
		}, nil
	}

	router := setupRouter()

	req, _ := http.NewRequest("GET", "/check", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "isEqual")
}
