package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type IFileService interface {
	GetFileInfo(ctx context.Context, names []string) ([]domain.FileStats, error)
	UploadFile(ctx context.Context, name string, data []byte) error
}

type Server struct {
	httpServer *http.Server
	errCh      chan error
	uc         IFileService
}

func NewHttpServer(host string, uc IFileService) *Server {
	s := &Server{errCh: make(chan error), uc: uc}
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(RecoveryWithZerolog(&log.Logger))
	// r.Use(gin.Logger())

	r.GET("/status", healthHandler)
	r.GET("/file/info", s.fileinfoHandler)
	r.GET("/file/upload", s.uploadHandler)
	r.GET("/internal/test_recovery", checkPanic)

	s.httpServer = &http.Server{
		Addr:    host,
		Handler: r,
	}
	return s
}

func (s *Server) GetAddr() string {
	return s.httpServer.Addr
}

func (s *Server) Start() {
	go func() {
		err := s.httpServer.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			s.errCh <- err
		}
	}()
}

func (s *Server) Stop() error {
	err := s.httpServer.Shutdown(context.Background())
	if err != nil {
		return err
	}
	return nil
}

type ErrorResult struct {
	Code uuid.UUID `json:"code"`
	Msg  string    `json:"msg"`
}

// RecoveryWithZerolog returns a middleware that recovers from any panics and logs the error using zerolog.
func RecoveryWithZerolog(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var (
					err error
					ok  bool
				)

				if err, ok = r.(error); !ok {
					err = errors.New(r.(string))
				}
				data := ErrorResult{Code: uuid.New(), Msg: "internal server error"}
				logger.Error().
					Err(err).
					Stack().
					Str("uuid", data.Code.String()).
					Msg("Recovered from panic")

				c.AbortWithStatusJSON(http.StatusInternalServerError, data)
			}
		}()
		c.Next()
	}
}
