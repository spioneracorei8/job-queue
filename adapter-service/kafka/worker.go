package kafka

import (
	"adapter-service/models"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
)

func (k *kafkaQueue) sendMailWorker(data []byte) {
	var form = new(models.MailForm)
	if err := json.Unmarshal(data, &form); err != nil {
		logrus.Fatalf("Error: %s", err.Error())
	}
	msg := mail.NewMessage()
	msg.SetAddressHeader("From", k.SMTP_ADDRESS, k.SENDER_NAME)
	msg.SetHeader("To", form.To)
	msg.SetHeader("Subject", form.Subject)
	msg.SetBody("text/html", form.Body)

	d := mail.NewDialer(k.SMTP_HOST, k.SMTP_PORT, k.SMTP_USERNAME, k.SMTP_PASSWORD)
	if err := d.DialAndSend(msg); err != nil {
		logrus.Errorf("Error: %s", err.Error())
	}
}

func (k *kafkaQueue) testWorker(data []byte) {
	logrus.Infof("✅ testWorker HANDLE! data: %s", string(data))
}
