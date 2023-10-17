package api

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gogoalish/doodocs-test/internal/service"
	"github.com/gogoalish/doodocs-test/utils"
)

func (s *Server) mailHandler(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	emails := strings.Split(c.PostForm("emails"), ",")
	cc := strings.Split(c.PostForm("cc"), ",")
	bcc := strings.Split(c.PostForm("bcc"), ",")

	var fileNames, filePaths []string
	headers := c.Request.MultipartForm.File["file"]
	for _, header := range headers {
		file, err := header.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		defer file.Close()

		allowedTypes, err := utils.LoadAllowedTypes("for_sending")
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
	params := &service.MailParams{
		To:        emails,
		CC:        cc,
		BCC:       bcc,
		Subject:   c.PostForm("subject"),
		Body:      c.PostForm("body"),
		FilePaths: filePaths,
		FileNames: fileNames,
	}
	err = s.service.SendMessage(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
