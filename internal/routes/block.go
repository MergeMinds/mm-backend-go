package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DataType struct {
	Format string `json:"format" `
	Text   string `json:"text"`
}

type BlockModel struct {
	Id        uuid.UUID `json:"id" binding:"required"`
	BlockType string    `json:"blockType" binding:"required"`
	Data      DataType  `json:"data" binding:"required"`
	CourseId  uuid.UUID `json:"courseId" binding:"required"`
}

type BlockModelResponse struct {
	BlockType string    `json:"blockType" binding:"required"`
	Data      DataType  `json:"data" binding:"required"`
	CourseId  uuid.UUID `json:"courseId" binding:"required"`
}

type CreateBlockType struct {
	BlockType string   `json:"blockType" binding:"required"`
	Data      DataType `json:"data" binding:"required"`
}

// @description Get block data
// @summary Get block data
// @tags blocks
// @produce json
// @param blockId path int true "Block ID"
// @success 201 {object} BlockModelResponse
// @failure 400 {object} apierr.ApiError "Invalid ID"
// @failure 404 {object} apierr.ApiError "Block not found"
// @failure 500 {object} apierr.ApiError "Internal server error"
// @router /block [GET]
func GetBlock(c *gin.Context, blockId string) {}

// @description Register a new account
// @summary Register a new account
// @tags blocks
// @accept json
// @produce json
// @param request body CreateBlockType true "Block payload"
// @success 201 {object} BlockModelResponse
// @failure 400 {object} apierr.ApiError "Invalid JSON"
// @failure 403 {object} apierr.ApiError "No permission"
// @failure 500 {object} apierr.ApiError "Internal server error"
// @router /block [POST]
func CreateBlock(c *gin.Context) {}

// @description Change single or multiple parameters of the block
// @summary Modify block
// @tags blocks
// @accept json
// @produce json
// @param blockId path int true "Block ID"
// @param request body CreateBlockType true "Block payload"
// @success 200 {object} BlockModelResponse
// @failure 400 {object} apierr.ApiError "Invalid ID"
// @failure 404 {object} apierr.ApiError "Parameter not found"
// @failure 404 {object} apierr.ApiError "Block not found"
// @failure 500 {object} apierr.ApiError "Internal server error"
// @router /block [PATCH]
func PatchBlock(c *gin.Context, blockId string) {}

// @description Will remove block from course but won't delete it from database
// @summary Remove block
// @tags blocks
// @produce json
// @param blockId path int true "Block ID"
// @success 204
// @failure 400 {object} apierr.ApiError "Invalid ID"
// @failure 404 {object} apierr.ApiError "Block not found"
// @failure 500 {object} apierr.ApiError "Internal server error"
// @router /block [DELETE]
func DeleteBlock(c *gin.Context, blockId string) {}
