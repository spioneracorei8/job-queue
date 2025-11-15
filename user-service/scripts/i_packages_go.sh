# !/bin/bash

# FRAMEWORK
go get -u github.com/gin-gonic/gin

# UUID
go get github.com/gofrs/uuid

# ENV
go get github.com/joho/godotenv

# LOG
go get github.com/sirupsen/logrus

# CAST DATA TYPE
go get github.com/spf13/cast

# GORM
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

# GRPC
go get google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
