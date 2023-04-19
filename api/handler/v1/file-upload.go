package v1

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/osg_task/api/models"
)

// @Summary upload media file
// @Description Through this api can upload media file
// @Tags 		Media
// @Security	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		file 	formData file true "image-upload"
// @Success  	201 	{object} models.FileResponse
// @Failure  	400 	{object} models.FailureInfo
// @Failure  	500 	{object} models.FailureInfo
// @Router 	 	/image-upload  	 [post]
func (h *handlerV1) UploadMedia(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			StatusCode:  http.StatusInternalServerError,
			Description: "Invalid access token",
			Error:       err,
		})
		h.log.Error("Error while getting claims of access token ", err)
		return
	}

	file := models.File{}

	err = c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			StatusCode:  http.StatusBadRequest,
			Description: "Please check your data",
		})
		h.log.Error("Error while binding request: ", err)
		return
	}

	ext := filepath.Ext(file.File.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".mp4" && ext != "pdf" {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			Description: "Unsupported file format",
		})
		return
	}

	mediaType := "image"
	if ext == ".mp4" {
		mediaType = "video"
	}

	fileName := uuid.New().String() + filepath.Ext(file.File.Filename)

	f, err := file.File.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.FailureInfo{
			Description: "Ooops something went wrong",
		})
		h.log.Error("while opening formed file", err.Error())
		return
	}

	// Aws started
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.FailureInfo{
			Description: "Ooops something went wrong",
		})
		h.log.Error("while Creating config to AWS", err.Error())
		return
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	res, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(h.cfg.OSGBucket),
		Key:    aws.String(fileName),
		ACL:    "public-read",
		Body:   f,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.FailureInfo{
			Description: "Ooops something went wrong",
		})
		h.log.Error("while Creating file on AWS S3", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.FileResponse{
		Url:       res.Location,
		MediaType: mediaType,
	})
}
