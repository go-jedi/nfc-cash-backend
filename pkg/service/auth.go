package service

import (
	"fmt"
	"math/rand"
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
	tokenTTL   = 15 * time.Minute
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

func newRefreshToken() (string, error) { // генерация refresh токена
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func GenerateTokens(id int) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	accessToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) SignIn(username string, password string) (string, string, error) {
	resGeneratePasswordHash, err := hash.GeneratePasswordHash(password)
	if err != nil {
		return "", "", err
	}
	user, err := s.repo.GetUser(username, resGeneratePasswordHash)
	if err != nil {
		return "", "", err
	}
	if len(user) == 0 {
		return "", "", fmt.Errorf("неправильный логин или пароль")
	}

	accessToken, refreshToken, err := GenerateTokens(user[0].Id)
	if err != nil {
		return "", "", err
	}

	_, err = s.repo.AddRefreshToken(user[0].Id, refreshToken, time.Now().Add(time.Hour*24*30))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshTokens(refreshToken string) (string, string, error) {
	resGetUserIdByRefreshToken, _, err := s.repo.GetUserIdByRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := GenerateTokens(resGetUserIdByRefreshToken)
	if err != nil {
		return "", "", err
	}

	_, err = s.repo.AddRefreshToken(resGetUserIdByRefreshToken, refreshToken, time.Now().Add(time.Hour*24*30))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
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
