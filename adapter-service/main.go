package main

import (
	"adapter-service/helper"
	my_kafka "adapter-service/kafka"
	"adapter-service/routes"
	_grcp_adapter_handler "adapter-service/services/adapter/grpc"
	_mail_repo "adapter-service/services/mail/repository"
	_mail_us "adapter-service/services/mail/usecase"
	"fmt"
	"net"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
)

var (
	ROOT_PATH    string
	GRPC_PORT    string
	GRPC_TIMEOUT int
	SERVER_READY chan bool

	KAFKA_BROKER_URL string

	SMTP_ADDRESS  string
	SMTP_HOST     string
	SMTP_PORT     int
	SMTP_USERNAME string
	SMTP_PASSWORD string

	MAIL_SENDER_NAME string
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	index := strings.LastIndex(basepath, "/config")
	if index != -1 {
		ROOT_PATH = strings.Replace(basepath, "/config", "", index)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error loading .env file: %v", err)
	}
	GRPC_PORT = helper.GetENV("GRPC_PORT", "")
	GRPC_TIMEOUT = cast.ToInt(helper.GetENV("GRPC_TIMEOUT", ""))
	KAFKA_BROKER_URL = helper.GetENV("KAFKA_BROKER_URL", "")

	SMTP_ADDRESS = helper.GetENV("SMTP_ADDRESS", "")
	SMTP_HOST = helper.GetENV("SMTP_HOST", "")
	SMTP_PORT = cast.ToInt(helper.GetENV("SMTP_PORT", ""))
	SMTP_USERNAME = helper.GetENV("SMTP_USERNAME", "")
	SMTP_PASSWORD = helper.GetENV("SMTP_PASSWORD", "")

	MAIL_SENDER_NAME = helper.GetENV("MAIL_SENDER_NAME", "")
}

func GetPath(dir string) string {
	return fmt.Sprintf("%s/%s", ROOT_PATH, strings.Trim(dir, "/"))
}

func main() {
	var (
		grpcServ = grpc.NewServer()
	)
	defer grpcServ.GracefulStop()

	queueProducer := my_kafka.NewQueueProducerImpl()
	queueConsumer := my_kafka.NewQueueConsumerImpl()

	//==============================================================
	// # REPOSITORIES
	//==============================================================
	mailRepo := _mail_repo.NewMailRepositoryImpl(queueProducer)

	//==============================================================
	// # USECASES
	//==============================================================
	mailUs := _mail_us.NewMailUsecaseImpl(mailRepo)

	//==============================================================
	// # HANDLERS
	//==============================================================
	grpcAdapterHandler := _grcp_adapter_handler.NewGrpcAdapterHandlerImpl(mailUs)

	//==============================================================
	// # KAFKA
	//==============================================================
	kafkaRoute := my_kafka.NewKafkaQueue(queueProducer, queueConsumer, SMTP_ADDRESS, SMTP_HOST, SMTP_PORT, SMTP_USERNAME, SMTP_PASSWORD, MAIL_SENDER_NAME)
	kafkaRoute.StartKafkaQueue(KAFKA_BROKER_URL)

	//==============================================================
	// # GRPC
	//==============================================================
	grpcRoute := routes.NewGrpcRoute(grpcServ)
	grpcRoute.RegisterAdapterHandler(grpcAdapterHandler)

	go func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered from panic: %v", r)
		}
		startGrpcServer(grpcServ)
	}()
	SERVER_READY <- true
}

func startGrpcServer(grpcServ *grpc.Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", GRPC_PORT))
	if err != nil {
		logrus.Fatalf("Error starting gRPC server: %v \n", err)
		return
	}
	logrus.Infoln("gRPC server listening on port", GRPC_PORT)
	if err := grpcServ.Serve(lis); err != nil {
		logrus.Fatalf("Error serving gRPC server: %v \n", err)
		return
	}
}
