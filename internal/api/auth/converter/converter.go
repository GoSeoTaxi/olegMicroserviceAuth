package converter

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	desc "github.com/GoSeoTaxi/olegMicroserviceAuth/grpc/pkg/user_v1"
	"github.com/GoSeoTaxi/olegMicroserviceAuth/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertProtoToModelDeleteRequest(user *desc.DeleteRequest) int64 {
	return user.Id
}

func ConvertProtoToModelGetRequest(user *desc.GetRequest) int64 {
	return user.Id
}

func ConvertModelToProtoGetRequest(user *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		UserInfo: &desc.UserInfo{
			Id:        user.ID,
			Name:      user.UserT.Name,
			Email:     user.UserT.Email,
			Role:      desc.Role(user.UserT.Role),
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdateAt:  timestamppb.New(user.UpdatedAt),
		},
	}
}

func ConvertProtoToModelUpdateRequest(user *desc.UpdateRequest) *model.User {
	var name, email string
	var role int64

	if user.Name != nil {
		name = user.Name.Value
	}
	if user.Email != nil {
		email = user.Email.Value
	}
	if user.Role != 0 {
		role = int64(user.Role)
	}

	return &model.User{
		ID: user.Id,
		UserT: struct {
			Name         string
			Email        string
			HashPassword string
			Role         int64
		}{Name: name, Email: email, HashPassword: "", Role: role},
	}
}

func ConvertProtoToModelCreateRequest(user *desc.CreateRequest) *model.User {
	hash := md5.Sum([]byte(user.UserCreate.Password))
	hashedPassword := hex.EncodeToString(hash[:])
	return &model.User{
		UserT: struct {
			Name         string
			Email        string
			HashPassword string
			Role         int64
		}{Name: user.UserCreate.Name, Email: user.UserCreate.Email, HashPassword: hashedPassword, Role: int64(user.UserCreate.Role)},
	}
}

func ConvertModelToProtoCreateRequest(id int64) (res *desc.CreateResponse, err error) {
	if id < 1 {
		err = fmt.Errorf("error create User")
	}
	return &desc.CreateResponse{Id: id}, err
}
