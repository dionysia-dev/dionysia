package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/learn-video/streaming-platform/internal/model"
	"github.com/learn-video/streaming-platform/internal/service"
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
// @Success 201 {object} model.Input
// @Router /inputs [post]
func (c *InputController) CreateInput(ctx *gin.Context) {
	var inputData model.Input
	if err := ctx.BindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input, err := c.inputHandler.CreateInput(ctx, &inputData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, input)
}

func (c *InputController) GetInput(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	input, err := c.inputHandler.GetInput(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, input)
}

func (c *InputController) DeleteInput(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := c.inputHandler.DeleteInput(ctx, id); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func NewNotificationController(nh service.NotificationHandler) *NotificationController {
	return &NotificationController{
		notificationHandler: nh,
	}
}

func (n *NotificationController) EnqueuePackaging(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := n.notificationHandler.PackageStream(id); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}
