package grpc

import (
	"context"
	"time"
	"user-service/helper"
	"user-service/models"
	"user-service/proto/proto_models"
	"user-service/services/user"
)

type grpcUserHandler struct {
	userUs user.UserUsecase
	proto_models.UnimplementedUserServer
}

func NewGrpcUserHandler(userUs user.UserUsecase) proto_models.UserServer {
	return &grpcUserHandler{
		userUs: userUs,
	}

}

func (g *grpcUserHandler) RegisterUser(ctx context.Context, request *proto_models.UserRequest) (*proto_models.UserResponse, error) {
	if request == nil {
		return &proto_models.UserResponse{}, nil
	}

	user, err := g.userUs.FetchUserByIdCardNumber(request.RegisterUser.IdCardNumber)
	if err != nil {
		return nil, err
	}

	if user != nil {
		var now = helper.NewTimestampFromTime(time.Now())
		user.IdCardNumber = request.RegisterUser.IdCardNumber
		user.TitleNameTH = request.RegisterUser.TitleNameTh
		user.FirstNameTH = request.RegisterUser.FirstNameTh
		user.LastNameTH = request.RegisterUser.LastNameTh
		user.TitleNameEN = request.RegisterUser.TitleNameEn
		user.FirstNameEN = request.RegisterUser.FirstNameEn
		user.LastNameEN = request.RegisterUser.LastNameEn
		user.MobilePhone = request.RegisterUser.MobilePhone
		user.OfficePhone = request.RegisterUser.OfficePhone
		user.Email = request.RegisterUser.Email
		user.BOD = helper.NewTimestampFromString((request.RegisterUser.Bod))
		user.Gender = request.RegisterUser.Gender
		user.UpdatedBy = user.Id.String()
		user.UpdatedAt = now
	} else {
		var now = helper.NewTimestampFromTime(time.Now())
		user = new(models.User)
		user.GenUUID()
		user.IdCardNumber = request.RegisterUser.IdCardNumber
		user.TitleNameTH = request.RegisterUser.TitleNameTh
		user.FirstNameTH = request.RegisterUser.FirstNameTh
		user.LastNameTH = request.RegisterUser.LastNameTh
		user.TitleNameEN = request.RegisterUser.TitleNameEn
		user.FirstNameEN = request.RegisterUser.FirstNameEn
		user.LastNameEN = request.RegisterUser.LastNameEn
		user.MobilePhone = request.RegisterUser.MobilePhone
		user.OfficePhone = request.RegisterUser.OfficePhone
		user.Email = request.RegisterUser.Email
		user.BOD = helper.NewTimestampFromString((request.RegisterUser.Bod))
		user.Gender = request.RegisterUser.Gender
		user.CreatedBy = user.Id.String()
		user.UpdatedBy = user.Id.String()
		user.CreatedAt = now
		user.UpdatedAt = now
	}

	if err := g.userUs.UpsertUser(user); err != nil {
		return nil, err
	}

	return &proto_models.UserResponse{UserId: user.Id.String()}, nil
}
