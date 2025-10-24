package grpc_client

import (
	"context"
	"errors"
	"io"
	"time"

	pb "github.com/f0xdl/file-processor-grpc/api/generated/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	GetTimeout    = time.Minute
	UploadTimeout = 5 * time.Minute
)

type Handler struct {
	fileClient pb.FileProcessorClient
}

func NewHandler(client *grpc.ClientConn) *Handler {
	fileClient := pb.NewFileProcessorClient(client)
	return &Handler{fileClient: fileClient}
}

func (h *Handler) GetFileInfo(ctx context.Context, names []string) ([]domain.FileStats, error) {
	getCtx, cancel := context.WithTimeout(ctx, GetTimeout)
	defer cancel()
	stream, err := h.fileClient.GetFileStats(getCtx, &pb.FileList{Paths: names})

	if err != nil {
		return nil, gErr(err)
	}

	var results []domain.FileStats
	for {
		fileStats, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error().Err(err).Msg("Error receiving file stats")
			return nil, gErr(err)
		}
		results = append(results, fileStatsPbToD(fileStats))
	}
	return results, nil
}

func (h *Handler) UploadFile(ctx context.Context, name string, data []byte) error {
	if len(data) == 0 {
		return errors.New("data is empty")
	}

	uploadCtx, cancel := context.WithTimeout(ctx, UploadTimeout)
	defer cancel()

	req := &pb.UploadFileReq{
		Filename: name,
	}

	stream, err := h.fileClient.UploadFile(uploadCtx)
	if err != nil {
		return gErr(err)
	}

	chunk := 1000
	for i := 0; i < len(data); i += chunk {
		if len(data) < i+chunk {
			chunk = len(data) - i
		}
		req.Content = data[i : i+chunk]
		err = stream.Send(req)
		if err != nil {
			return gErr(err)
		}
	}
	//final empty array
	req.Content = []byte{}
	err = stream.Send(req)
	if err != nil {
		return gErr(err)
	}

	res, err := stream.CloseAndRecv()
	if res != nil {
		log.Warn().Str("res", res.Hash).Str("filename", name).Msg("")
	}
	if err != nil {
		return gErr(err)
	}
	return nil
}

func gErr(err error) error {
	if err == nil {
		return nil
	}
	statusErr, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch statusErr.Code() {
	case codes.DeadlineExceeded:
		log.Warn().Err(err).Any("code", statusErr.Code()).Msg("deadline exceeded")
		return domain.ErrTimeout
	case codes.Canceled:
		log.Warn().Err(err).Msg("canceled")
		return domain.ErrCanceled
	case codes.Unavailable:
		log.Error().Err(err).Msg("service unavailable")
		return domain.ErrServiceUnavailable
	default:
		log.Warn().Err(err).Msg("unknown gRPC error")
		return domain.ErrInternal
	}
}
