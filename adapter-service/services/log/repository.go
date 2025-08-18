package log

import (
	"adapter-service/proto/proto_models"
	"context"
)

type GrpcLogRepository interface {
	SaveLog(ctx context.Context, params *proto_models.SendLogRequest) error
}
