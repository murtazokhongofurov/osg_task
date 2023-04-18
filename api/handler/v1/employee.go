package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/osg_task/api/models"
	"github.com/osg_task/internal/controller/storage/repo"
)

// @Summary Create new employee
// @Description Through this api, can create an employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param  body body models.EmployeeReq true "CreateEmployee"
// @Success 201 {object} models.FailureInfo
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /employee [POST]
func (h *handlerV1) CreateEmployee(c *gin.Context) {
	var (
		body models.EmployeeReq
	)
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("error binding request", err.Error())
		return
	}

	res, err := h.storage.Task().CreateEmployee(&repo.Employee{
		Id:           uuid.New().String(),
		FullName:     body.FullName,
		ProfilePhoto: body.ProfilePhoto,
		BirthDate:    body.BirthDate,
		Phone:        body.Phone,
		Position:     body.Position,
		Role:         body.Role,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("error create request", err.Error())
		return
	}
	c.JSON(http.StatusCreated, res)
}

// @Summary GET an employee
// @Description Through this api, can GET an employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param  	id path string true "id"
// @Success 200 {object} models.FailureInfo
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /employee/{id} [GET]
func (h *handlerV1) GetEmployee(c *gin.Context) {
	id := c.Param("id")

	res, err := h.storage.Task().GetEmployee(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error getting employee", err.Error())
		return
	}
	data := models.EmployeeRes{
		Id:           res.Id,
		FullName:     res.FullName,
		ProfilePhoto: res.ProfilePhoto,
		BirthDate:    res.BirthDate,
		Phone:        res.Phone,
		Position:     res.Position,
		Role:         res.Role,
	}

	c.JSON(http.StatusOK, data)
}

// @Summary GET employees
// @Description Through this api, can GET employees
// @Tags Employee
// @Accept json
// @Produce json
// @Success 200 {object} models.FailureInfo
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /employees [GET]
func (h *handlerV1) GetAllEmployees(c *gin.Context) {
	res, err := h.storage.Task().GetAllEmployee()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error getting employee", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary UPDATE an employee
// @Description Through this api, can UPDATE an employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param 	body body models.EmployeeRes true "UpdateEmployee"
// @Success 200 {object} models.FailureInfo
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /employee [PUT]
func (h *handlerV1) UpdateEmployee(c *gin.Context) {
	var body models.EmployeeRes
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while updating employee", err.Error())
		return
	}
	res, err := h.storage.Task().UpdateEmployee(&repo.Employee{
		Id:           body.Id,
		FullName:     body.FullName,
		ProfilePhoto: body.ProfilePhoto,
		BirthDate:    body.BirthDate,
		Phone:        body.Phone,
		Position:     body.Position,
		Role:         body.Role,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while updating employee", err.Error())
		return
	}

	data := models.EmployeeRes{
		Id:           res.Id,
		FullName:     res.FullName,
		ProfilePhoto: res.ProfilePhoto,
		BirthDate:    res.BirthDate,
		Phone:        res.Phone,
		Position:     res.Position,
		Role:         res.Role,
	}

	c.JSON(http.StatusOK, data)
}
