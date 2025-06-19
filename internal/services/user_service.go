package services

import (
	"lolymarsh/internal/entity"
	"lolymarsh/internal/request"
	"lolymarsh/pkg/common"
	"lolymarsh/pkg/util"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(req *request.RegisterRequest) (*entity.UserEntity, error)
	LoginUser(req *request.LoginRequest) (*entity.UserEntity, *string, error)
}

func (sv *service) UserService() UserService {
	return sv
}

func (sv *service) RegisterUser(req *request.RegisterRequest) (*entity.UserEntity, error) {

	reqGetUser := &common.Filters{
		Field: "username",
		Value: req.Username,
	}

	getUser, err := sv.repo.UserRepository().GetUserByFilter(reqGetUser)
	if err != nil {
		return nil, common.HandleErrorService("RegisterUser", http.StatusBadRequest, "Username not found", err)
	}

	if getUser != nil {
		return nil, common.HandleErrorService("RegisterUser", http.StatusBadRequest, "Username already exists", nil)
	}

	newUID := util.NewUUID("UID_")
	currentEpochTime := util.GetCurrentEpochTimeMillisecond()

	hashPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, common.HandleErrorService("RegisterUser", http.StatusBadRequest, "Failed to hash password", err)
	}

	listInput := &entity.UserEntity{
		UserID:    newUID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashPassword,
		Role:      "USER",
		IsActive:  true,
		CreatedAt: 0,
		UpdatedAt: currentEpochTime,
		UpdatedBy: "",
	}

	err = sv.repo.UserRepository().InsertUser(listInput)
	if err != nil {
		return nil, common.HandleErrorService("RegisterUser", http.StatusBadRequest, "Failed to insert user", err)
	}

	listInput.Password = "********"

	return listInput, nil
}

func (sv *service) LoginUser(req *request.LoginRequest) (*entity.UserEntity, *string, error) {

	reqGetUser := &common.Filters{
		Field: "username",
		Value: req.Username,
	}

	getUser, err := sv.repo.UserRepository().GetUserByFilter(reqGetUser)
	if err != nil {
		return nil, nil, common.HandleErrorService("LoginUser", http.StatusBadRequest, "Username not found", err)
	}

	if getUser == nil {
		return nil, nil, common.HandleErrorService("LoginUser", http.StatusBadRequest, "Username not found", nil)
	}

	if !getUser.IsActive {
		return nil, nil, common.HandleErrorService("LoginUser", http.StatusBadRequest, "User is not active", nil)
	}

	if !checkPasswordHash(req.Password, getUser.Password) {
		return nil, nil, common.HandleErrorService("LoginUser", http.StatusBadRequest, "Wrong password", nil)
	}

	authToken, err := sv.generateTokenAuth(getUser)
	if err != nil {
		return nil, nil, common.HandleErrorService("Login", fiber.StatusInternalServerError, "failed to generate token", err)
	}

	getUser.Password = "********"

	return getUser, authToken, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *service) generateTokenAuth(user *entity.UserEntity) (*string, error) {
	conf := s.conf.Auth
	claims := jwt.MapClaims{
		"user_id":    user.UserID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"role":       user.Role,
		"is_active":  user.IsActive,
		"exp":        time.Now().Add(time.Hour * time.Duration(conf.TokenExpire)).Unix(),
		"issued_at":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(conf.SecretKey))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
