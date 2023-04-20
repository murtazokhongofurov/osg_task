package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osg_task/api/models"
	"github.com/osg_task/internal/controller/storage/repo"
)

// @Summary Login an employee
// @Description Through this api, can login an employee
// @Tags 	Employee Login
// @Accept 	json
// @Produce json
// @Param  	phone query string true "phone_number"
// @Success 200 {object} models.GetEmployee
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router 	/login/{phone} [GET]
func (h *handlerV1) Login(c *gin.Context) {
	phone := c.Query("phone")

	res, err := h.storage.Task().GetEmployeeByPhone(&repo.PhoneNumber{
		Phone: phone,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while getting employe by phone: ", err.Error())
		return
	}
	if res.Role == "admin" {
		h.jwthandler.Role = "admin"
	} else if res.Role == "direktor" {
		h.jwthandler.Role = "direktor"
	} else if res.Role == "developer" {
		h.jwthandler.Role = "developer"
	} else {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Invalid role_name",
		})
		h.log.Error("Error role checking: ", err.Error())
		return
	}

	h.jwthandler.Sub = res.Id
	h.jwthandler.Aud = []string{"osg.task"}
	h.jwthandler.SigninKey = h.cfg.SigninKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessToken := tokens[0]
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("error occured while generating tokens ", err.Error())
		return
	}

	data := models.GetEmployee{
		Id:           res.Id,
		FullName:     res.FullName,
		ProfilePhoto: res.ProfilePhoto,
		BirthDate:    res.BirthDate,
		Phone:        res.Phone,
		Position:     res.Position,
		Role:         res.Role,
		AccessToken:  accessToken,
	}

	c.JSON(http.StatusOK, data)
}
