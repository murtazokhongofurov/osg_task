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
// Security BearerAuth
// @Accept json
// @Produce json
// @Param  body body models.EmployeeReq true "CreateEmployee"
// @Success 201 {object} models.FailureInfo
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /employee [POST]
func (h *handlerV1) CreateEmployee(c *gin.Context) {
	// _, err := GetClaims(*h, c)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, models.FailureInfo{
	// 		StatusCode:  http.StatusInternalServerError,
	// 		Description: "Invalid access token",
	// 	})
	// 	h.log.Error("Error while getting claims of access token ", err.Error())
	// 	return
	// }

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

	IsPhone, err := h.storage.Task().CheckField(&repo.CheckfieldReq{
		Field: "phone",
		Value: body.Phone,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("error checkfield request", err.Error())
		return
	}
	if IsPhone.Exists {
		c.JSON(http.StatusConflict, models.FailureInfo{
			StatusCode:  http.StatusConflict,
			Description: "Phone Number already in use",
		})
		h.log.Error("error checking phone", err.Error())
		return
	}

	if body.Role == "admin" {
		h.jwthandler.Role = "admin"
	} else if body.Role == "direktor" {
		h.jwthandler.Role = "direktor"
	} else if body.Role == "developer" {
		h.jwthandler.Role = "developer"
	} else {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Invalid role_name",
		})
		h.log.Error("Error role checking: ", err.Error())
		return
	}
	id := uuid.New().String()

	h.jwthandler.Sub = id
	h.jwthandler.Aud = []string{"osg.task"}
	h.jwthandler.SigninKey = h.cfg.SigninKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	refreshToken := tokens[1]
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("error occured while generating tokens ", err.Error())
		return
	}

	res, err := h.storage.Task().CreateEmployee(&repo.Employee{
		Id:           id,
		FullName:     body.FullName,
		ProfilePhoto: body.ProfilePhoto,
		BirthDate:    body.BirthDate,
		Phone:        body.Phone,
		Position:     body.Position,
		Role:         body.Role,
		RefreshToken: refreshToken,
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
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param  	id path string true "id"
// @Success 200 {object} models.FailureInfo
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /employee/{id} [GET]
func (h *handlerV1) GetEmployee(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

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
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} models.EmployeeList
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /employees [GET]
func (h *handlerV1) GetAllEmployees(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

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
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	body body models.EmployeeRes true "UpdateEmployee"
// @Success 200 {object} models.EmployeeRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /employee [PUT]
func (h *handlerV1) UpdateEmployee(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	var body models.EmployeeRes
	err = c.ShouldBindJSON(&body)
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

// @Summary DELETE an employee
// @Description Through this api, can Delete an employee
// @Tags Employee
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	id path string true "id"
// @Success 200 {object} models.Success
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /employee/{id} [DELETE]
func (h *handlerV1) DeleteEmployee(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	id := c.Param("id")
	err = h.storage.Task().DeleteEmployee(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while delete employee ", err.Error())
		return
	}
	c.JSON(http.StatusOK, models.Success{
		Message:    "Successfully deleted employee",
		StatusCode: http.StatusOK,
	})

}

// @Summary Create an employee
// @Description Through this api, can create a developer
// @Tags Developer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	body body models.DeveloperReq true "CreateRole"
// @Success 201 {object} models.DeveloperRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /developer [POST]
func (h *handlerV1) CreateDeveloper(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	var body models.DeveloperReq
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request ", err.Error())
		return
	}
	res, err := h.storage.Task().CreateDeveloper(&repo.Developer{
		Id:            uuid.New().String(),
		EmployeeId:    body.EmployeeId,
		DeveloperRole: body.DeveloperRole,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while creating developer", err.Error())
		return
	}
	c.JSON(http.StatusCreated, res)
}

// @Summary GET all developer roles
// @Description Through this api, can get all developer roles
// @Tags Developer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} models.AllDeveloperRole
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /developer/roles [GET]
func (h *handlerV1) GetDeveloperRoles(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	res, err := h.storage.Task().GetAllDeveloperRole()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while getting developers roles", err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
