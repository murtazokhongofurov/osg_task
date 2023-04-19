package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/osg_task/api/models"
	"github.com/osg_task/internal/controller/storage/repo"
)

// @Summary Create project
// @Description Through this api, can create project
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param  	body body models.ProjectReq true "CreateProject"
// @Success 201 {object} models.ProjectRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /prject [POST]
func (h *handlerV1) CreateProject(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	var body models.ProjectReq
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request", err.Error())
		return
	}
	res, err := h.storage.Task().CreateProject(&repo.ProjectReq{
		DeveloperId:  body.DeveloperId,
		ProjectName:  body.ProjectName,
		StartedDate:  body.StartedDate,
		FinishedDate: body.FinishedDate,
		FileUrl:      body.FileUrl,
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
}

// @Summary Get a project
// @Description Through this api, can get a project
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param  	id path int true "prject_id"
// @Success 200 {object} models.ProjectRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /prject/{id} [GET]
func (h *handlerV1) GetProject(c *gin.Context) {
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
	project_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while parsing id", err.Error())
		return
	}
	res, err := h.storage.Task().GetProject(int(project_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while getting project id ", err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get projects
// @Description Through this api, can get projects
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} models.AllProject
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /prjects [GET]
func (h *handlerV1) GetAllProjects(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}
	res, err := h.storage.Task().GetAllProject()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while getting projects ", err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update project
// @Description Through this api, can update project
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	body body models.ProjectRes true "UpdateProject"
// @Success 200 {object} models.ProjectRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /prject [PUT]
func (h *handlerV1) UpdateProject(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}

	var body models.ProjectRes
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request", err.Error())
		return
	}
	res, err := h.storage.Task().UpdateProject(&repo.ProjectRes{
		Id:           body.Id,
		DeveloperId:  body.DeveloperId,
		ProjectName:  body.ProjectName,
		StartedDate:  body.StartedDate,
		FinishedDate: body.FinishedDate,
		Status:       body.Status,
		FileUrl:      body.FileUrl,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while update project ", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)

}

// @Summary DELETE a project
// @Description Through this api, can delete a project
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param  	id path int true "prject_id"
// @Success 200 {object} models.Success
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /prject/{id} [DELETE]
func (h *handlerV1) DeleteProject(c *gin.Context) {
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
	project_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while parsing id", err.Error())
		return
	}

	err = h.storage.Task().DeleteProject(int(project_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while delete project ", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Success{
		Message:    "Successfully delete project",
		StatusCode: http.StatusOK,
	})

}

// @Summary Update project status
// @Description Through this api, can update project status
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	body body models.StatusUpdate true "UpdateProjectStatus"
// @Success 200 {object} models.StatusUpdate
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /prject/status [PUT]
func (h *handlerV1) UpdateProjectStatus(c *gin.Context) {
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
		h.log.Error("Error while parsing id", err.Error())
		return
	}
	res, err := h.storage.Task().UpdateProjectStatus(&repo.StatusUpdate{
		Id:     body.Id,
		Status: body.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while update project ", err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
