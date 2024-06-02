package app

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/aakashkaji/empolyee-go/app/internal/domain"
	"github.com/aakashkaji/empolyee-go/app/internal/service"
	"github.com/gin-gonic/gin"
)

type EmpHandler struct {
	service service.EmpolyeeService
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags Ping
// @Accept json
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {string} Pong
// @Router /ping [get]
func (u *EmpHandler) TestHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"value": "Pong"})

}

// Create an new empolyee record
// @summary Create new Empolyee
// @Description Create an new empolyee record.
// @Tags Empolyees
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param request body domain.EmpolyeeRequestDto true "User data" example={"name": "Aakash Gupta", "position": "Software Engineer", "salary": 1234.5}
// @Success 201 {string} string "Empolyee created successfully"
// @Failure 400 {object} string "Invalid request"
// @Router /empolyees [post]
func (u *EmpHandler) CreateEmpHandler(c *gin.Context) {

	var empolyeeRequestDto domain.EmpolyeeRequestDto

	if err := c.ShouldBindJSON(&empolyeeRequestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request!"})
		return
	}

	// handle context
	timeoutCtx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	createEmpolyee, apiErr := u.service.NewCreateEmpolyee(timeoutCtx, empolyeeRequestDto)

	if apiErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": apiErr})
		return
	}

	c.JSON(http.StatusCreated, createEmpolyee)

}

func (u *EmpHandler) UpdateEmpHandler(c *gin.Context) {

	id := c.Param("id")

	var empolyee domain.EmpolyeeUpdateRequestDto

	if err := c.ShouldBindJSON(&empolyee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// handle context
	timeoutCtx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	empID, _ := strconv.Atoi(id) // convert str to int
	empolyee.ID = empID

	emp, err := u.service.UpdateEmpolyee(timeoutCtx, empolyee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": emp})

}

func (u *EmpHandler) DeleteEmpHandler(c *gin.Context) {

	id := c.Param("id")

	empID, _ := strconv.Atoi(id)

	timeoutCtx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	err := u.service.DeleteEmpolyee(timeoutCtx, empID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record Deleted!"})

}

// GetEmpById retrieves an employee by ID
// @Summary Get an employee by ID
// @Schemes
// @Description Retrieve employee details by employee ID
// @Tags employees
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Param id path int true "Employee ID"
// @Success 200 {object} domain.Empolyee
// @Failure 400 {object} error
// @Router /empolyees/{id} [get]
func (u *EmpHandler) GetEmpById(c *gin.Context) {

	id := c.Param("id")
	// handle context
	timeoutCtx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	empID, _ := strconv.Atoi(id)

	emp, err := u.service.GetEmpolyeeId(timeoutCtx, empID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": emp})

}

func (u *EmpHandler) GetAllEmp(c *gin.Context) {

	pageNo, pageSize := c.Query("page_no"), c.Query("per_page")
	// handle context

	timeoutCtx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	emps, err := u.service.GetAllEmpolyee(timeoutCtx, pageNo, pageSize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": emps})
}
