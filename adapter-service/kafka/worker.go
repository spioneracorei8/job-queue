package kafka

import (
	"adapter-service/models"
	"adapter-service/proto/proto_models"
	"context"
	"encoding/json"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/mail.v2"
)

func (k *kafkaQueue) sendMailWorker(data []byte) error {
	var form = new(models.MailForm)
	if err := json.Unmarshal(data, &form); err != nil {
		return err
	}
	msg := mail.NewMessage()
	msg.SetAddressHeader("From", k.SMTP_ADDRESS, k.SENDER_NAME)
	msg.SetHeader("To", form.To)
	msg.SetHeader("Subject", form.Subject)
	msg.SetBody("text/html", form.Body)

	d := mail.NewDialer(k.SMTP_HOST, k.SMTP_PORT, k.SMTP_USERNAME, k.SMTP_PASSWORD)
	if err := d.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func (k *kafkaQueue) sendLogWorker(data []byte) error {
	conn, err := grpc.NewClient(k.SERVICE_CLIENT_LOG_GRPC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto_models.NewLogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(k.GRPC_TIMEOUT*int(time.Second)))
	defer cancel()

	var request = new(proto_models.LogRequest)
	if err := json.Unmarshal(data, &request); err != nil {
		return err
	}
	_, err = client.SaveLog(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
