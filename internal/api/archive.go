package api

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (server *Server) archiveHandler(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	defer file.Close()

	extension := filepath.Ext(header.Filename)
	if extension != ".zip" {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tempFile, err := os.CreateTemp("", "*.zip")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arcInfo, err := server.service.GetInformation(tempFile.Name(), header.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, arcInfo)
}
