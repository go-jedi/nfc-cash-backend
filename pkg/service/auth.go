package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/rob-bender/nfc-cash-backend/appl_row"
	"github.com/rob-bender/nfc-cash-backend/pkg/email"
	"github.com/rob-bender/nfc-cash-backend/pkg/hash"
	"github.com/rob-bender/nfc-cash-backend/pkg/repository"
)

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.TodoAuth
}

func NewAuthService(r repository.TodoAuth) *AuthService {
	return &AuthService{
		repo: r,
	}
}

func (s *AuthService) CreateUser(userForm appl_row.CreateUser) (int, error) {
	resGeneratePasswordHash, err := hash.GeneratePasswordHash(userForm.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	userForm.Password = resGeneratePasswordHash
	resCreateUser, statusCode, err := s.repo.CreateUser(userForm)
	if err != nil {
		return statusCode, err
	}
	statusCodeSendActivationMail, err := email.SendActivationMail(userForm.Email, resCreateUser) // отправка письма активации на почту пользователя
	if err != nil {
		return statusCodeSendActivationMail, err
	}
	return statusCodeSendActivationMail, nil
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	resGeneratePasswordHash, err := hash.GeneratePasswordHash(password)
	if err != nil {
		return "", err
	}
	user, err := s.repo.GetUser(username, resGeneratePasswordHash)
	if err != nil {
		return "", err
	}
	if len(user) == 0 {
		return "", fmt.Errorf("неправильный логин или пароль")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user[0].Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) CheckEmailExist(userForm appl_row.CheckEmailExist) (bool, int, error) {
	return s.repo.CheckEmailExist(userForm)
}

func (s *AuthService) CheckUsernameExist(userForm appl_row.CheckUsernameExist) (bool, int, error) {
	return s.repo.CheckUsernameExist(userForm)
}

func (s *AuthService) CheckConfirmAccount(userForm appl_row.CheckConfirmAccount) (bool, int, error) {
	resGeneratePasswordHash, err := hash.GeneratePasswordHash(userForm.Password)
	if err != nil {
		return false, http.StatusInternalServerError, err
	}
	userForm.Password = resGeneratePasswordHash
	return s.repo.CheckConfirmAccount(userForm)
}
