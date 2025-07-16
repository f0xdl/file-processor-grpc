package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
	"time"
)

const RequestInfoTimeout = time.Second * 25
const UploadFileTimeout = time.Minute

func isJson(c *gin.Context) bool {
	ct := c.GetHeader("Content-Type")
	return ct != "" && strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0])) == "application/json"
}

func (s *Server) fileinfoHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), RequestInfoTimeout)
	defer cancel()

	if !isJson(c) {
		c.String(http.StatusUnsupportedMediaType, "Content-Type header is not application/json")
		return
	}

	data := struct {
		Filenames []string `json:"filenames"`
	}{}
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Error().Err(err).Msg("Error parsing json")
		c.String(http.StatusBadRequest, "Error parsing json")
		return
	}

	if len(data.Filenames) == 0 {
		c.String(http.StatusBadRequest, "filenames is empty")
		return
	}

	results, err := s.uc.GetFileInfo(ctx, data.Filenames)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to process files")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}

func (s *Server) uploadHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), UploadFileTimeout)
	defer cancel()

	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get file")
		c.String(http.StatusBadRequest, "Missing file")
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Error().Err(err).Msg("Failed to open uploaded file")
		c.String(http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer file.Close()

	err = s.uc.UploadFile(ctx, fileHeader.Filename, file)
	if err != nil {
		log.Error().Err(err).Msg("Failed to upload file")
		c.String(http.StatusInternalServerError, "Failed to upload file")
		return
	}

	c.Status(http.StatusOK)
}

func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func checkPanic(c *gin.Context) {
	panic("test recovery")
}
