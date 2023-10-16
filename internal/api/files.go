package api

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gogoalish/doodocs-test/utils"
)

func (server *Server) filesHandler(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var fileNames []string
	var filePaths []string
	headers := c.Request.MultipartForm.File["files[]"]

	for _, header := range headers {
		file, err := header.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		defer file.Close()

		allowedTypes, err := utils.LoadAllowedTypes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		if !utils.IsAllowedType(header.Filename, allowedTypes) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mimetype of uploaded file"})
			return
		}

		tempFile, err := os.CreateTemp("", "upload-*.tmp")
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		defer tempFile.Close()

		_, err = io.Copy(tempFile, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		filePaths = append(filePaths, tempFile.Name())
		fileNames = append(fileNames, header.Filename)
	}
	zipData, err := server.service.CompressFiles(filePaths, fileNames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename=archive.zip")
	c.Data(http.StatusOK, "application/zip", *zipData)
	for _, path := range filePaths {
		os.Remove(path)
	}
}
