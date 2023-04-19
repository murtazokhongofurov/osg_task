package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/osg_task/api/models"
	"github.com/osg_task/internal/controller/storage/repo"
)

// @Summary Create task
// @Description Through this api, can create task
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param  	body body models.TaskReq true "CreateTask"
// @Success 201 {object} models.TaskRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /task [POST]
func (h *handlerV1) CreateTask(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	var body models.TaskReq
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request", err.Error())
		return
	}
	data, err := h.storage.Task().GetDeveloper(body.DeveloperId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while getting developer info", err.Error())
		return
	}
	if data.DeveloperRole == "teamlead" {
		res, err := h.storage.Task().CreateTask(&repo.TaskReq{
			DeveloperId: body.DeveloperId,
			Title:       body.Title,
			Description: body.Description,
			FileUrl:     body.FileUrl,
			StartedDate: body.StartedDate,
			EndDate:     body.EndDate,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.FailureInfo{
				StatusCode:  http.StatusInternalServerError,
				Description: "Something went wrong",
			})
			h.log.Error("Error while create project", err.Error())
			return
		}

		c.JSON(http.StatusCreated, res)
	} else {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while checking developer role", err.Error())
		return
	}
}

// @Summary Create task
// @Description Through this api, can create task
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param  id path int true "task_id"
// @Success 200 {object} models.TaskRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /task/{id} [GET]
func (h *handlerV1) GetTask(c *gin.Context) {
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
	task_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error parsing request ", err.Error())
		return
	}
	res, err := h.storage.Task().GetTask(int(task_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while getting task ", err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary UPDATE task
// @Description Through this api, can update task
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param  	body body models.TaskUpdateReq true "UpdateTask"
// @Success 200 {object} models.TaskRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /task [PUT]
func (h *handlerV1) UpdateTask(c *gin.Context) {
	claim, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}
	var body models.TaskUpdateReq
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request", err.Error())
		return
	}
	data, err := h.storage.Task().GetDeveloper(claim.Sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while getting developer info", err.Error())
		return
	}
	if data.DeveloperRole == "teamlead" {
		res, err := h.storage.Task().UpdateTask(&repo.TaskRes{
			Id:          body.Id,
			Title:       body.Title,
			Description: body.Description,
			FileUrl:     body.FileUrl,
			StartedDate: body.StartedDate,
			EndDate:     body.EndDate,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.FailureInfo{
				StatusCode:  http.StatusInternalServerError,
				Description: "Something went wrong",
			})
			h.log.Error("Error while update task info", err.Error())
			return
		}
		c.JSON(http.StatusOK, res)

	} else {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Permission developer role",
		})
		h.log.Error("Error while checking role developer info", err.Error())
		return
	}
}

// @Summary Create task
// @Description Through this api, can create task
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} models.AllTask
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /tasks [GET]
func (h *handlerV1) GetAllTask(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	res, err := h.storage.Task().GetAllTask()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while getting tasks info", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Update task status
// @Description Through this api, can update task status
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body models.StatusUpdate true "Update status"
// @Success 200 {object} models.StatusUpdate
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /task/status [PUT]
func (h *handlerV1) UpdateTaskStatus(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	var body models.StatusUpdate

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request", err.Error())
		return
	}
	res, err := h.storage.Task().UpdateStatus(&repo.StatusUpdate{
		Id:     body.Id,
		Status: body.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while update task info", err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary DELETE task
// @Description Through this api, can delete task
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param  id path int true "task_id"
// @Success 200 {object} models.Success
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /task/{id} [DELETE]
func (h *handlerV1) DeleteTask(c *gin.Context) {
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
	task_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "ENter right info",
		})
		h.log.Error("Error while parsing id ", err.Error())
		return
	}
	err = h.storage.Task().DeleteTask(int(task_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while delete task ", err.Error())
		return
	}
	c.JSON(http.StatusOK, models.Success{
		Message:    "Successfully deleted task",
		StatusCode: http.StatusOK,
	})
}

// @Summary Create task
// @Description Through this api, can create task
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param  id path string true "GetTasks"
// @Success 200 {object} models.AllTask
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /tasks/{id} [GET]
func (h *handlerV1) GetTaskDevId(c *gin.Context) {
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
	res, err := h.storage.Task().GetTaskDeveloperId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while getting task ", err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
