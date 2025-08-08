package server

import (
	"fmt"
	"net"
	"net/http"
	my_logger "user-service/logger"
	"user-service/middleware"
	"user-service/routes"
	_grpc_user_handler "user-service/services/user/grpc"
	_user_repo "user-service/services/user/repository"
	_user_us "user-service/services/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	PSQL_CONNECTION string
	APP_PORT        string
	GRPC_PORT       string
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
	userRepo := _user_repo.NewUserRepoImpl(database)
	// grpcUserRepo := repository.NewGrpcUserRepoImpl()

	//==============================================================
	// # USECASES
	//==============================================================
	userUs := _user_us.NewUserUsecaseImpl(userRepo)

	//==============================================================
	// # HANDLERS
	//==============================================================
	// _ = _register_handler.NewRegisterHandlerImpl(registerUs)

	//==============================================================
	// # GRPC HANDLERS
	//==============================================================
	grpcUserHandler := _grpc_user_handler.NewGrpcUserHandler(userUs)

	//==============================================================
	// # API
	//==============================================================
	app.GET("/", func(g *gin.Context) {
		g.JSON(http.StatusOK, "Hello World!")
	})
	_ = routes.NewRoute(app, middl)

	//==============================================================
	// # GRPC
	//==============================================================
	grpcRoute := routes.NewGrpcRoute(grpcServ)
	grpcRoute.RegisterUserHandler(grpcUserHandler)

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
