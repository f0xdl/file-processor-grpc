package grpc_client

import (
	"context"
	"io"
	"time"

	pb "github.com/f0xdl/file-processor-grpc/api/generated/fileprocessor"
	"github.com/f0xdl/file-processor-grpc/internal/domain"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const GetTimeout = time.Minute

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
