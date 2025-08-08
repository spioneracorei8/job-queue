package repository

import (
	"auth-service/models"
	"auth-service/proto/proto_models"
	"auth-service/services/user"
	"context"
	"time"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcUserRepo struct {
	grpcAddr string
	timeout  int
}

func NewGrpcUserRepoImpl(grpcAddr string, timeout int) user.UserRepository {
	return &grpcUserRepo{
		grpcAddr: grpcAddr,
		timeout:  timeout,
	}
}

func (r *grpcUserRepo) RegisterUser(params *models.Register) (*uuid.UUID, error) {
	conn, err := grpc.NewClient(r.grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto_models.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout*int(time.Second)))
	defer cancel()

	var registerUserObj = new(proto_models.RegisterUserObj)
	registerUserObj.IdCardNumber = params.IdCardNumber
	registerUserObj.TitleNameTh = params.TitleNameTh
	registerUserObj.FirstNameTh = params.FirstNameTh
	registerUserObj.LastNameTh = params.LastNameTh
	registerUserObj.TitleNameEn = params.TitleNameEn
	registerUserObj.FirstNameEn = params.FirstNameEn
	registerUserObj.LastNameEn = params.LastNameEn
	registerUserObj.MobilePhone = params.MobilePhone
	registerUserObj.OfficePhone = params.OfficePhone
	registerUserObj.Email = params.Email
	registerUserObj.Bod = params.BOD
	registerUserObj.Gender = params.Gender

	var request = &proto_models.UserRequest{
		RegisterUser: registerUserObj,
	}

	response, err := client.RegisterUser(ctx, request)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	if response.UserId == "" {
		return nil, nil
	}

	var userId = uuid.FromStringOrNil(response.UserId)

	return &userId, nil
}
