package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type dataType struct {
	Format string `json:"format" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type blockType struct {
	Id        uuid.UUID `json:"id" binding:"required"`
	BlockType string    `json:"blockType" binding:"required"`
	Data      dataType  `json:"data" binding:"required"`
	CourseId  uuid.UUID `json:"courseId" binding:"required"`
}

type createBlockType struct {
	BlockType string
	Data      dataType
}

func GetBlock(c *gin.Context, blockId string) {}

func CreateBlock(c *gin.Context) {}

func PatchBlock(c *gin.Context, blockId string) {}

func DeleteBlock(c *gin.Context, blockId string) {}
