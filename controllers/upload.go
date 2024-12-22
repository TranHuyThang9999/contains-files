package controllers

import (
	"fmt"
	"log"
	"path/filepath"
	"virtual/configs"
	"virtual/utils"

	"github.com/gin-gonic/gin"
)

type UploadControllerResponse struct {
	Code     int      `json:"code"`
	Message  string   `json:"message"`
	ListFile []string `json:"list_file,omitempty"`
}

func (u *UploadControllerResponse) Upload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	newNameFile := utils.GenUUID()
	var uploadedFiles []string

	files := form.File["upload[]"]
	if len(files) == 0 {
		ctx.JSON(400, gin.H{"message": "No files uploaded"})
		return
	}

	for _, file := range files {
		fileExtension := filepath.Ext(file.Filename)
		newFileName := fmt.Sprintf("%s%s", newNameFile, fileExtension)

		err := ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s", "public", newFileName))
		if err != nil {
			log.Println("Failed to save file:", err)
			ctx.JSON(500, gin.H{"message": "Failed to save file"})
			return
		}

		uploadedFiles = append(uploadedFiles, configs.GetConfig().PathFile+newFileName)
	}

	ctx.JSON(200, UploadControllerResponse{
		Code:     200,
		Message:  "success",
		ListFile: uploadedFiles,
	})
}
