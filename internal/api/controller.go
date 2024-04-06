package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/learn-video/streaming-platform/internal/model"
	"github.com/learn-video/streaming-platform/internal/queue"
	"github.com/learn-video/streaming-platform/internal/service"
	"github.com/learn-video/streaming-platform/internal/task"
)

type InputController struct {
	inputHandler service.InputHandler
}

type NotificationController struct{}

func NewInputController(inputHandler service.InputHandler) *InputController {
	return &InputController{inputHandler: inputHandler}
}

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
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	input, err := c.inputHandler.GetInput(ctx, uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, input)
}

func NewNotificationController() *NotificationController {
	return &NotificationController{}
}

func (n *NotificationController) EnqueuePackaging(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	client := queue.NewClient("localhost:6379")

	task, err := task.NewPackageTask(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	info, err := client.Enqueue(task)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	log.Printf("Task enqueued: %s", info.ID, info.Queue)

	ctx.Status(http.StatusCreated)
}
