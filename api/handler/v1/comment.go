package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/osg_task/api/models"
	"github.com/osg_task/internal/controller/storage/repo"
)

// @Summary Create comment
// @Description Through this api, can create comment
// @Tags Comment
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	body body models.CommentReq true "CreateComment"
// @Success 201 {object} models.CommentRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /comment [POST]
func (h *handlerV1) CreateComment(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}
	var body models.CommentReq
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request ", err.Error())
		return
	}
	res, err := h.storage.Task().CreateComment(&repo.CommentReq{
		DeveloperId: body.DeveloperId,
		TaskId:      body.TaskId,
		Text:        body.Text,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while create comment ", err.Error())
		return
	}
	c.JSON(http.StatusCreated, res)
}

// @Summary GET comments of task
// @Description Through this api, can get comments by task_id
// @Tags Comment
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	id path int true "task_id"
// @Success 201 {object} models.AllComment
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /comment/{id} [GET]
func (h *handlerV1) GetComments(c *gin.Context) {
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
	taskId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Enter right info",
		})
		h.log.Error("Error while parsing request", err.Error())
		return
	}
	res, err := h.storage.Task().GetComments(int(taskId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while get comment ", err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary UPDATE comment
// @Description Through this api, can UPDATE comments
// @Tags Comment
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	body body models.CommentUpdateReq true "Update comment"
// @Success 201 {object} models.AllComment
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.CommentRes
// @Router /comment [PUT]
func (h *handlerV1) UpdateComment(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
		})
		h.log.Error("Error while getting claims of access token ", err.Error())
		return
	}
	var body models.CommentUpdateReq
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while binding request ", err.Error())
		return
	}

	res, err := h.storage.Task().UpdateComment(&repo.CommentRes{
		Id:   body.Id,
		Text: body.Text,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while update comment ", err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary DELETE comment
// @Description Through this api, can delete comment
// @Tags Comment
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	id path int true "commentId"
// @Success 201 {object} models.Success
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /comment/{id} [DELETE]
func (h *handlerV1) DeleteComment(c *gin.Context) {
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
	commentId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Enter right info",
		})
		h.log.Error("Error while parsing id", err.Error())
		return
	}
	err = h.storage.Task().DeleteComment(commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Something went wrong",
		})
		h.log.Error("Error while delete comment ", err.Error())
		return
	}
	c.JSON(http.StatusOK, models.Success{
		Message: "Successfully deleted commet",
		StatusCode: http.StatusOK,
	})
}
