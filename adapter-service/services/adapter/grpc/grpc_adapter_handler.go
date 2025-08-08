package grpc

import (
	"adapter-service/models"
	"adapter-service/proto/proto_models"
	"adapter-service/services/mail"
	"context"
)

type grpcAdapterHandler struct {
	proto_models.UnimplementedAdapterServer
	mailUs mail.MailUsecase
}

func NewGrpcAdapterHandlerImpl(mailUs mail.MailUsecase) proto_models.AdapterServer {
	return &grpcAdapterHandler{
		mailUs: mailUs,
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
