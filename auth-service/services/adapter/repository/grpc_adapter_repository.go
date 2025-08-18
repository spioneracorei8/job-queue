package repository

import (
	"auth-service/models"
	"auth-service/proto/proto_models"
	"auth-service/services/adapter"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcAdapterRepository struct {
	grpcAddr string
	timeout  int
}

func NewGrpcAdapterRepositoryImpl(grpcAddr string, timeout int) adapter.GrpcAdapterRepository {
	return &grpcAdapterRepository{
		grpcAddr: grpcAddr,
		timeout:  timeout,
	}
}

func (g *grpcAdapterRepository) SendMail(mail *models.MailForm) (*proto_models.SendMailResponse, error) {
	conn, err := grpc.NewClient(g.grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto_models.NewAdapterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(g.timeout*int(time.Second)))
	defer cancel()

	var request = &proto_models.SendMailRequest{
		To:      mail.To,
		ToName:  mail.ToName,
		Subject: mail.Subject,
		Body:    mail.Body,
	}

	response, err := client.SendMail(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (g *grpcAdapterRepository) SendLog(params *models.LogForm) error {
	conn, err := grpc.NewClient(g.grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto_models.NewAdapterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(g.timeout*int(time.Second)))
	defer cancel()

	metadataJSON, _ := json.Marshal(params.Metadata)
	var logForm = new(proto_models.SendLogRequest)
	logForm.Service = params.Service
	logForm.Level = params.Level
	logForm.Message = params.Message
	logForm.Metadata = string(metadataJSON)
	logForm.Time = params.Time.String()
	logForm.Method = fmt.Sprint(params.Method)
	logForm.Path = fmt.Sprint(params.Path)
	logForm.Ip = fmt.Sprint(params.Ip)

	_, err = client.SendLog(ctx, logForm)
	if err != nil {
		return err
	}

	return nil
}
