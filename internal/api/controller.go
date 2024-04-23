package api

import (
	"net/http"

	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputController struct {
	inputHandler service.InputHandler
}

type NotificationController struct {
	notificationHandler service.NotificationHandler
}

func NewInputController(inputHandler service.InputHandler) *InputController {
	return &InputController{inputHandler: inputHandler}
}

// @Summary Create an input
// @Description Create an input ready to be ingested
// @Accept json
// @Produce json
// @Param input body model.Input true "Input data"
// @Success 201 {object} api.SuccessResponse
// @Failure 400 {object} api.ErrorResponse "Invalid input data"
// @Failure 500 {object} api.ErrorResponse "Internal server error"
// @Router /inputs [post]
func (c *InputController) CreateInput(ctx *gin.Context) {
	var inputData InputData
	if err := ctx.BindJSON(&inputData); err != nil {
		statusCode, response := HandleValidationError(err)
		if statusCode == 0 && response == nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Error: Error{Message: "InternalServerError: handle validation failed"},
			})

			return
		}

		ctx.JSON(int(statusCode), response)

		return
	}

	audioProfiles := make([]service.AudioProfile, 0, len(inputData.AudioProfiles))
	for _, audioProfileData := range inputData.AudioProfiles {
		audioProfiles = append(audioProfiles, service.AudioProfile{
			Codec:   audioProfileData.Codec,
			Bitrate: audioProfileData.Bitrate,
		})
	}

	videoProfiles := make([]service.VideoProfile, 0, len(inputData.VideoProfiles))
	for _, videoProfileData := range inputData.VideoProfiles {
		videoProfiles = append(videoProfiles, service.VideoProfile{
			Codec:          videoProfileData.Codec,
			Bitrate:        videoProfileData.Bitrate,
			MaxKeyInterval: videoProfileData.MaxKeyInterval,
			Framerate:      videoProfileData.Framerate,
			Width:          videoProfileData.Width,
			Height:         videoProfileData.Height,
		})
	}

	input, err := c.inputHandler.CreateInput(ctx, &service.Input{
		Name:          inputData.Name,
		Format:        inputData.Format,
		AudioProfiles: audioProfiles,
		VideoProfiles: videoProfiles,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: Error{Message: "InternalServerError: failed creating input"},
		})

		return
	}

	responseData := InputData{
		ID:            input.ID,
		Name:          input.Name,
		Format:        input.Format,
		AudioProfiles: inputData.AudioProfiles,
		VideoProfiles: inputData.VideoProfiles,
	}
	ctx.JSON(http.StatusCreated, SuccessResponse{
		Message: "Input created successfully",
		Data:    responseData,
	})
}

// @Summary Get input
// @Description Get input information by ID
// @Produce json
// @Param id path string true "Input ID"
// @Success 200 {object} api.SuccessResponse
// @Failure 400 {object} api.ErrorResponse "Invalid UUID format"
// @Failure 500 {object} api.ErrorResponse "Internal server error"
// @Router /inputs/{id} [get]
func (c *InputController) GetInput(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error: Error{Message: "BadRequest: invalid UUID format"},
		})

		return
	}

	input, err := c.inputHandler.GetInput(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: Error{Message: "InternalServerError: failed creating input"},
		})

		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{
		Data: input,
	})
}

// @Summary Delete input
// @Description Delete input by ID
// @Produce json
// @Param id path string true "Input ID"
// @Success 200 {object} api.SuccessResponse
// @Failure 400 {object} api.ErrorResponse "Invalid UUID format"
// @Failure 500 {object} api.ErrorResponse "Internal server error"
// @Router /inputs/{id} [delete]
func (c *InputController) DeleteInput(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error: Error{Message: "BadRequest: invalid UUID format"},
		})

		return
	}

	if err := c.inputHandler.DeleteInput(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: Error{Message: "InternalServerError: failed deleting input"},
		})

		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{
		Message: "Input deleted successfully",
	})
}

func (c *InputController) Authenticate(ctx *gin.Context) {
	var authData IngestAuthData
	if err := ctx.BindJSON(&authData); err != nil {
		statusCode, response := HandleValidationError(err)
		if statusCode == 0 && response == nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse{
				Error: Error{Message: "InternalServerError: failed to authenticate"},
			})

			return
		}

		ctx.JSON(int(statusCode), response)

		return
	}

	err := c.inputHandler.Authenticate(ctx, service.IngestAuth{
		Path:   authData.Path,
		Action: authData.Action,
	})

	switch {
	case err == service.ErrFailedAuth:
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error: Error{Message: "BadRequest: invalid credentials"},
		})

		return
	case err != nil:
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: Error{Message: "InternalServerError: failed to authenticate"},
		})

		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{
		Message: "Ingest authenticated successfully",
	})
}

func NewNotificationController(nh service.NotificationHandler) *NotificationController {
	return &NotificationController{
		notificationHandler: nh,
	}
}

// @Summary Enqueue packacing job
// @Description Enqueue packaging job using input URL, format and ID
// @Produce json
// @Param id query string true "Input ID"
// @Success 201 {object} api.SuccessResponse
// @Failure 400 {object} api.ErrorResponse "Invalid UUID format"
// @Failure 500 {object} api.ErrorResponse "Internal server error"
// @Router /notifications/package [post]
func (n *NotificationController) EnqueuePackaging(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Error: Error{Message: "Invalid UUID format"},
		})

		return
	}

	if err := n.notificationHandler.PackageStream(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: Error{Message: "InternalServerError: while creating input"},
		})

		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse{
		Message: "Packaging job enqueued successfully",
	})
}
