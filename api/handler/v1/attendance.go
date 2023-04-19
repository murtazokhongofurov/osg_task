package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osg_task/api/models"
	"github.com/osg_task/internal/controller/storage/repo"
)

// @Summary Create attendance
// @Description Through this api, can create attendance
// @Tags Attendance
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	body body models.AttendanceReq true "CreateAttendance"
// @Success 201 {object} models.AttendanceRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /attendance [POST]
func (h *handlerV1) CreateAttendance(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	var body models.AttendanceReq
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request ", err.Error())
		return
	}

	res, err := h.storage.Task().CreateAttendance(&repo.AttendanceReq{
		EmployeeId: body.EmployeeId,
		Type:       body.Type,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while creating attendance", err.Error())
		return
	}
	c.JSON(http.StatusCreated, res)
}

// @Summary Update attendance
// @Description Through this api, can update attendance
// @Tags Attendance
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	body body models.AttendanceUpdateReq true "UpdateAttendance"
// @Success 201 {object} models.AttendanceRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /	 [PUT]
func (h *handlerV1) UpdateAttendance(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	var body models.AttendanceUpdateReq

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request ", err.Error())
		return
	}
	res, err := h.storage.Task().UpdateAttendance(&repo.AttendanceUpdateReq{
		Id:         body.Id,
		EmployeeId: body.EmployeeId,
		Type:       body.Type,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while creating attendance", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
