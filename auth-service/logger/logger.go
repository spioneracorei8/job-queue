package logger

import (
	"auth-service/constants"
	"auth-service/helper"
	"auth-service/models"
	"auth-service/services/adapter"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

func GormLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)
}

type Logger struct {
	service string
	adapter adapter.AdapterUsecase
	logger  zerolog.Logger
}

type Event struct {
}

func NewLogger(service string, adapter adapter.AdapterUsecase) *Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	return &Logger{
		service: service,
		adapter: adapter,
		logger: zerolog.New(output).
			With().
			Timestamp().
			Str("service", service).
			Logger(),
	}
	// OUTPUT:
	/*
		2025-08-09T12:05:42+07:00 INF server/server.go:82 > Starting app
		2025-08-09T12:05:42+07:00 ERR server/server.go:83 > Something went wrong error=eiei
		2025-08-09T12:05:42+07:00 DBG server/server.go:84 > User info age=30 user=johndoe
	*/
	// w := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	// zlog.Logger = zerolog.New(w).With().Timestamp().Caller().Logger()

	// OUTPUT:
	/*
		12:06PM INF Starting app
		12:06PM ERR Something went wrong error=eiei
		12:06PM DBG User info age=30 user=johndoe
	*/
	// zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func (l *Logger) Trace(g *gin.Context, msg string, metadata map[string]any) {
	method, _ := g.Get("method")
	path, _ := g.Get("path")
	ip, _ := g.Get("ip")
	e := l.logger.Trace()
	for k, v := range metadata {
		e.Interface(k, v)
	}
	e.Msg(msg)
	logForm := l.formatLogForm(constants.LEVEL_TRACE, msg, metadata, method, path, ip)
	l.adapter.SendLog(logForm)
}

func (l *Logger) Debug(g *gin.Context, msg string, metadata map[string]any) {
	method, _ := g.Get("method")
	path, _ := g.Get("path")
	ip, _ := g.Get("ip")
	e := l.logger.Debug()
	for k, v := range metadata {
		e.Interface(k, v)
	}
	e.Msg(msg)
	logForm := l.formatLogForm(constants.LEVEL_DEBUG, msg, metadata, method, path, ip)
	l.adapter.SendLog(logForm)
}

func (l *Logger) Info(g *gin.Context, msg string, metadata map[string]any) {
	method, _ := g.Get("method")
	path, _ := g.Get("path")
	ip, _ := g.Get("ip")
	e := l.logger.Info()
	for k, v := range metadata {
		e.Interface(k, v)
	}
	e.Msg(msg)
	logForm := l.formatLogForm(constants.LEVEL_INFO, msg, metadata, method, path, ip)
	l.adapter.SendLog(logForm)
}

func (l *Logger) Warn(g *gin.Context, msg string, metadata map[string]any) {
	method, _ := g.Get("method")
	path, _ := g.Get("path")
	ip, _ := g.Get("ip")
	e := l.logger.Warn()
	for k, v := range metadata {
		e.Interface(k, v)
	}
	e.Msg(msg)
	logForm := l.formatLogForm(constants.LEVEL_WARN, msg, metadata, method, path, ip)
	l.adapter.SendLog(logForm)
}
func (l *Logger) Error(g *gin.Context, msg string, metadata map[string]any) {
	method, _ := g.Get("method")
	path, _ := g.Get("path")
	ip, _ := g.Get("ip")
	e := l.logger.Error()
	for k, v := range metadata {
		e.Interface(k, v)
	}
	e.Msg(msg)
	logForm := l.formatLogForm(constants.LEVEL_ERROR, msg, metadata, method, path, ip)
	l.adapter.SendLog(logForm)
}
func (l *Logger) Fatal(g *gin.Context, msg string, metadata map[string]any) {
	method, _ := g.Get("method")
	path, _ := g.Get("path")
	ip, _ := g.Get("ip")
	e := l.logger.Fatal()
	for k, v := range metadata {
		e.Interface(k, v)
	}
	e.Msg(msg)
	logForm := l.formatLogForm(constants.LEVEL_FATAL, msg, metadata, method, path, ip)
	l.adapter.SendLog(logForm)
}

func (l *Logger) formatLogForm(level, msg string, metadata map[string]any, method, path, ip any) *models.LogForm {
	return &models.LogForm{
		Time:     helper.NewTimestampRFC3339FromTime(time.Now()),
		Service:  l.service,
		Level:    level,
		Message:  msg,
		Metadata: metadata,
		Method:   method,
		Path:     path,
		Ip:       ip,
	}
}
