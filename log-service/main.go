package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log-service/constants"
	"log-service/proto/proto_models"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
)

var (
	ROOT_PATH        string
	GRPC_PORT        string
	GRPC_TIMEOUT     int
	SERVICE_NAME     string
	SERVER_READY     chan bool
	PROD_ENVIRONMENT bool
)

var (
	buildInfo, _ = debug.ReadBuildInfo()
	_, b, _, _   = runtime.Caller(0)
	basepath     = filepath.Dir(b)

	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.
			Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()
)

func init() {
	index := strings.LastIndex(basepath, "/config")
	if index != -1 {
		ROOT_PATH = strings.Replace(basepath, "/config", "", index)
	}
	if err := godotenv.Load(); err != nil {
		logger.Fatal().Msgf("Error loading .env file: %v", err)
	}
	GRPC_PORT = GetENV("GRPC_PORT", "")
	GRPC_TIMEOUT = cast.ToInt(GetENV("GRPC_TIMEOUT", ""))
	SERVICE_NAME = GetENV("SERVICE_NAME", "")
	PROD_ENVIRONMENT = cast.ToBool(GetENV("PROD_ENVIRONMENT", ""))
}

func GetPath(dir string) string {
	return fmt.Sprintf("%s/%s", ROOT_PATH, strings.Trim(dir, "/"))
}

func GetENV(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getCurrentFileName() string {
	var y, m, d = time.Now().Year(), time.Now().Month(), time.Now().Day()
	return fmt.Sprintf("%d_%02d_%02d", y, m, d)
}

// ==============================================================
// # gRPC route
// ==============================================================
type grpcRoute struct {
	proto_models.UnimplementedLogServer
	server *grpc.Server
}

func NewgRPCRoute(server *grpc.Server) *grpcRoute {
	return &grpcRoute{server: server}
}

func (g *grpcRoute) RegisterLogHandler(handler proto_models.LogServer) {
	proto_models.RegisterLogServer(g.server, handler)
}

// ==============================================================
// # gRPC handler
// ==============================================================
type grpcLogHandler struct {
	proto_models.UnimplementedLogServer
}

func NewgRPCLogHandlerImpl() proto_models.LogServer {
	return &grpcLogHandler{}
}

// ==============================================================
// # Start GRPC server
// ==============================================================
func startgRPCServer(grpcServ *grpc.Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", GRPC_PORT))
	if err != nil {
		logger.Fatal().Msgf("Error starting gRPC server: %v \n", err)
		return
	}
	logger.Info().Msgf("gRPC server listening on port %s", GRPC_PORT)
	if err := grpcServ.Serve(lis); err != nil {
		logger.Fatal().Msgf("Error serving gRPC server: %v \n", err)
		return
	}
}

func main() {
	var (
		grpcServ = grpc.NewServer()
	)
	defer grpcServ.GracefulStop()
	//==============================================================
	// # HANDLERS
	//==============================================================
	grpcLogHandler := NewgRPCLogHandlerImpl()
	//==============================================================
	// # GRPC
	//==============================================================
	grpcRoute := NewgRPCRoute(grpcServ)
	grpcRoute.RegisterLogHandler(grpcLogHandler)

	go func() {
		if r := recover(); r != nil {
			logger.Panic().Msgf("Recovered from panic: %v", r)
		}
		startgRPCServer(grpcServ)
	}()
	SERVER_READY <- true
}

func (g *grpcLogHandler) SaveLog(ctx context.Context, request *proto_models.LogRequest) (*proto_models.LogResponse, error) {
	if request == nil {
		return nil, nil
	}

	file, err := os.OpenFile(fmt.Sprintf("./logs/%s.log", getCurrentFileName()), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var metadata map[string]any
	if err := json.Unmarshal([]byte(request.Metadata), &metadata); err != nil {
		return nil, err
	}
	logger := zerolog.New(file).With().
		Str("time", request.Time).
		Str("service", request.Service).
		Str("level", request.Level).
		Interface("metadata", request.Metadata).
		Str("method", request.Method).
		Str("path", request.Path).
		Str("ip", request.Ip).
		Logger()

	switch request.Level {
	case constants.LEVEL_FATAL:
		logger.Fatal().Msg(request.Message)
	case constants.LEVEL_ERROR:
		logger.Error().Msg(request.Message)
	case constants.LEVEL_WARN:
		logger.Warn().Msg(request.Message)
	case constants.LEVEL_INFO:
		logger.Info().Msg(request.Message)
	case constants.LEVEL_DEBUG:
		logger.Debug().Msg(request.Message)
	case constants.LEVEL_TRACE:
		logger.Trace().Msg(request.Message)
	default:
		logger.Info().Msg(request.Message)
	}

	return &proto_models.LogResponse{Message: "successful"}, nil
}
