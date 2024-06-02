package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aakashkaji/empolyee-go/app/internal/domain"
	"github.com/aakashkaji/empolyee-go/app/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateEmpHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)

	// setup default
	r := gin.Default()

	// create mock service
	mockEmpolyeeService := new(service.EmpolyeeMockService)

	expectedResponse := &domain.Empolyee{
		ID:       1,
		Name:     "Aakash Gupta",
		Position: "Software Engineer",
		Salary:   120.0,
	}

	mockEmpolyeeService.On("CreateEmpolyee", mock.AnythingOfType("domain.Empolyee")).Return(expectedResponse, nil)

	// emph := EmpHandler{service: interface{}}

	newRequest := &domain.Empolyee{
		ID:       1,
		Name:     "Aakash Gupta",
		Position: "Software Engineer",
		Salary:   120.0,
	}

	data, _ := json.Marshal(newRequest)

	// r.POST("/api/v1/empolyees", "".CreateEmpHandler)

	req, _ := http.NewRequest("POST", "/api/v1/empolyees", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

}
