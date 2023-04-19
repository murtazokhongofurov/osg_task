package v1


// @Summary Create attendance
// @Description Through this api, can create attendance
// @Tags Attendance
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param 	body body models.
// @Success 201 {object} models.DeveloperRes
// @Failure 400 {object} models.FailureInfo
// @Failure 500 {object} models.FailureInfo
// @Router /developer [POST]