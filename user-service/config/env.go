package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"user-service/helper"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	ROOT_PATH       string
	PSQL_CONNECTION string
	APP_PORT        string
	GRPC_PORT       string
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
	PSQL_CONNECTION = helper.GetENV("PSQL_CONNECTION", "")
	APP_PORT = helper.GetENV("APP_PORT", "")
	GRPC_PORT = helper.GetENV("GRPC_PORT", "")
}

func GetPath(dir string) string {
	return fmt.Sprintf("%s/%s", ROOT_PATH, strings.Trim(dir, "/"))
}
