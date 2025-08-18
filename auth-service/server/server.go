package server

import (
	my_logger "auth-service/logger"
	"auth-service/middleware"
	"auth-service/routes"
	_adapter_repo "auth-service/services/adapter/repository"
	_adapter_us "auth-service/services/adapter/usecase"
	_register_handler "auth-service/services/register/handler"
	_register_repo "auth-service/services/register/repository"
	_register_us "auth-service/services/register/usecase"
	_user_repo "auth-service/services/user/repository"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	ROOT_PATH string

	APP_PORT     string
	GRPC_PORT    string
	GRPC_TIMEOUT int

	PSQL_CONNECTION string

	SERVICE_NAME string

	SERVICE_CLIENT_USER_GRPC_ADDRESS    string
	SERVICE_CLIENT_ADAPTER_GRPC_ADDRESS string
	SERVICE_CLIENT_LOG_GRPC_ADDRESS     string
}

func connectDatabase(PSQL_CONNECTION string) *gorm.DB {
	gormLogger := my_logger.GormLogger()
	database, err := gorm.Open(postgres.Open(PSQL_CONNECTION), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
		return nil
	}
	return database
}

func (s *Server) startGrpcServer(grpcServ *grpc.Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.GRPC_PORT))
	if err != nil {
		logrus.Fatalf("Error starting gRPC server: %v \n", err)
		return
	}
	logrus.Infoln("gRPC server listening on port", s.GRPC_PORT)
	if err := grpcServ.Serve(lis); err != nil {
		logrus.Fatalf("Error serving gRPC server: %v \n", err)
		return
	}
}

func (s *Server) Start() {
	var (
		app      = gin.Default()
		database = connectDatabase(s.PSQL_CONNECTION)
		middl    = middleware.InitMiddleware()
		grpcServ = grpc.NewServer()
	)
	defer grpcServ.GracefulStop()

	//==============================================================
	// # REPOSITORIES
	//==============================================================
	registerRepo := _register_repo.NewRegisterRepoImpl(database)
	userRepo := _user_repo.NewGrpcUserRepoImpl(s.SERVICE_CLIENT_USER_GRPC_ADDRESS, s.GRPC_TIMEOUT)
	adapterRepo := _adapter_repo.NewGrpcAdapterRepositoryImpl(s.SERVICE_CLIENT_ADAPTER_GRPC_ADDRESS, s.GRPC_TIMEOUT)
	//==============================================================
	// # USECASES
	//==============================================================
	registerUs := _register_us.NewRegisterUsImpl(registerRepo, userRepo, adapterRepo, s.ROOT_PATH)
	adapterUs := _adapter_us.NewAdapterUsecaseImpl(adapterRepo)

	//==============================================================
	// # LOGGER
	//==============================================================
	myZerolog := my_logger.NewLogger(s.SERVICE_NAME, adapterUs)

	//==============================================================
	// # HANDLERS
	//==============================================================
	registerHandler := _register_handler.NewRegisterHandlerImpl(myZerolog, registerUs)
	//==============================================================
	// # API
	//==============================================================
	app.GET("/", middl.Logger(), func(g *gin.Context) {
		myZerolog.Info(g, "test", map[string]any{"k1": "string", "k2": 2, "k3": 2.22})
		g.JSON(http.StatusOK, "Hello World!")
	})
	api := routes.NewRoute(app, middl)
	api.NewRegisterRoutes(registerHandler)
	//==============================================================
	// # GRPC
	//==============================================================
	routes.NewGrpcRoute(grpcServ)

	go func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered from panic: %v", r.(error))
		}
		s.startGrpcServer(grpcServ)
	}()

	if err := app.Run(fmt.Sprintf(":%s", s.APP_PORT)); err != nil {
		logrus.Fatalf("Error starting server: %v", err)
	}
}
