package grpc

import (
	"adapter-service/models"
	"adapter-service/proto/proto_models"
	"adapter-service/services/log"
	"adapter-service/services/mail"
	"context"
)

type grpcAdapterHandler struct {
	proto_models.UnimplementedAdapterServer
	mailUs mail.MailUsecase
	logUs  log.LogUsecase
}

func NewGrpcAdapterHandlerImpl(mailUs mail.MailUsecase, logUs log.LogUsecase) proto_models.AdapterServer {
	return &grpcAdapterHandler{
		mailUs: mailUs,
		logUs:  logUs,
	}
}

func (g *grpcAdapterHandler) SendMail(ctx context.Context, request *proto_models.SendMailRequest) (*proto_models.SendMailResponse, error) {
	if request == nil {
		return nil, nil
	}

	form := &models.MailForm{
		To:      request.GetTo(),
		ToName:  request.GetToName(),
		Subject: request.GetSubject(),
		Body:    request.GetBody(),
	}

	if err := g.mailUs.SendMail(ctx, form); err != nil {
		return nil, err
	}

	return nil, nil
}

func (g *grpcAdapterHandler) SendLog(ctx context.Context, request *proto_models.SendLogRequest) (*proto_models.SendLogResponse, error) {
	if request == nil {
		return nil, nil
	}

	if err := g.logUs.SaveLog(ctx, request); err != nil {
		return nil, err
	}

	return nil, nil
}
